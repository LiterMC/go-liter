
package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/kmcsr/go-logger"
	logrusl "github.com/kmcsr/go-logger/logrus"
)

const DEBUG = false
const (
	PROTOCOLS_URL = "https://wiki.vg/Protocol_version_numbers"
	BASE_URL = "https://wiki.vg/index.php?title=Protocol&oldid="
)

var (
	StableVersionMatcher = regexp.MustCompile(`^\d+(?:\.\d+)+$`)
)

var loger logger.Logger = logrusl.Logger

//go:embed template
var templateFS embed.FS
var genTime = time.Now().Format("2006-01-02 15:04:05.000 -07:00")

var numberTypes = []string{"Byte", "UByte", "Short", "Int", "Long", "Float", "Double", "VarInt", "VarLong"}
var basicTypes = []string{"Bool", "Byte", "UByte", "Short", "Int", "Long", "Float", "Double", "VarInt", "VarLong", "String", "UUID"}

func isNumberType(typ string)(bool){
	for _, t := range numberTypes {
		if typ == t {
			return true
		}
	}
	return false
}

func isBasicType(typ string)(bool){
	for _, t := range basicTypes {
		if typ == t {
			return true
		}
	}
	return false
}

type (
	packetField struct {
		Name string
		Type string
		WikiType string
		Note string
	}
	packetInfo struct {
		Matched bool

		State string
		Bound string
		Id int
		Name string
		Fields []packetField
		Table string

		RepeatInfo *repeatPacketInfo
	}
)

var protocolNameMap = make(map[int]string)

var tmpl = template.Must(template.New("").
	Funcs(template.FuncMap{
		"add": func(a, b int)(int){ return a + b },
		"wrap": func(args ...any)([]any){ return args },
		"join": func(sep string, args ...any)(string){
			arg := make([]string, len(args))
			for i, v := range args {
				switch w := v.(type) {
				case string:
					arg[i] = w
				default:
					arg[i] = fmt.Sprint(v)
				}
			}
			return strings.Join(arg, sep)
		},
		"replace": strings.ReplaceAll,
		"hasPrefix": strings.HasPrefix,
		"trimPrefix": strings.TrimPrefix,

		"gentime": func()(string){ return genTime },
		"isBasicType": isBasicType,
		"protocolName": func(v int)(string){ return protocolNameMap[v] },

		"render_packet_encode_method": renderPacketEncodeMethod,
		"render_packet_decode_method": renderPacketDecodeMethod,
	}).
	ParseFS(templateFS, "*/**.gotemp"),
)

func ExecuteTemplateTo(filename string, tmplname string, data any)(err error){
	fd, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Cannot create %q: %v", filename, err)
	}
	defer fd.Close()
	if err = tmpl.ExecuteTemplate(fd, tmplname, data); err != nil {
		return fmt.Errorf("Cannot execute template %q: %v", tmplname, err)
	}
	return nil
}

type Protocol struct {
	Name string
	Version int
	Doc string
	docid string
}

func main(){
	if DEBUG {
		loger.SetLevel(logger.DebugLevel)
	}else{
		loger.SetLevel(logger.InfoLevel)
	}

	loger.Info("==> Running packet protocol generator")
	defer loger.Info("==> Done packet protocol generator")

	loger.Infof("Generate time: %s", genTime)

	res, err := GetOrUseCached(PROTOCOLS_URL, "generator/cache/protocols.html")
	if err != nil {
		loger.Fatalf("Cannot get protocols page: %v", err)
	}
	doc, err := goquery.NewDocumentFromReader(res)
	res.Close()
	if err != nil {
		loger.Fatalf("Cannot parse page %q: %v", PROTOCOLS_URL, err)
	}

	protocols := make([]Protocol, 0)
	// Release name | Version number | Last known documentation
	doc.FindMatcher(goquery.Single(".wikitable")).ChildrenFiltered("tbody").Children().Each(func(i int, s *goquery.Selection){
		s = s.Children()
		if s.Length() != 3 {
			return
		}
		if !s.Is("td") {
			return
		}
		name := strings.TrimSpace(s.Slice(0, 1).Text())
		version, _ := strconv.Atoi(strings.TrimSpace(s.Slice(1, 2).Text()))
		doc, _ := s.Slice(2, 3).Find("a").Attr("href")
		doc = strings.TrimSpace(doc)
		if StableVersionMatcher.MatchString(name) && version > 0 {
			if oldid, ok := strings.CutPrefix(doc, BASE_URL); ok {
				if lasti := len(protocols) - 1; lasti >= 0 {
					if lastp := protocols[lasti]; lastp.Version == version {
						loger.Warnf("Skipped same protocol version %d between %s and %s", version, lastp.Name, name)
						return
					}
				}
				protocolNameMap[version] = name
				loger.Debugf("name=%s; version=%d; doc=%s", name, version, doc)
				protocols = append(protocols, Protocol{ Name: name, Version: version, Doc: doc, docid: oldid })
			}
		}
	})

	if err = ExecuteTemplateTo("protocols.go", "protocols.go.gotemp", map[string]any{
		"protocols": protocols,
	}); err != nil {
		loger.Fatal(err)
	}

	{
		files, err := os.ReadDir(".")
		if err != nil {
			loger.Fatalf("Cannot read current dir: %v", err)
		}
		for _, f := range files {
			if f.IsDir() && strings.HasPrefix(f.Name(), "v") {
				if err := os.RemoveAll(f.Name()); err != nil {
					loger.Warnf("Cannot remove old generated directory %s: %v", f.Name(), err)
				}else{
					loger.Infof("Removed %s", f.Name())
				}
			}
		}
	}

	var wg, syncWg sync.WaitGroup
	wg.Add(len(protocols))
	syncWg.Add(len(protocols))
	for _, p := range protocols {
		go func(p Protocol){
			defer wg.Done()
			loger.Infof("-> Generating %s from %q...", p.Name, p.Doc)
			defer loger.Infof("-> Done generate %s", p.Name)
			if err := genProtocol(p, &syncWg); err != nil {
				loger.Errorf("Cannot generate %s: %v", p.Name, err)
			}
		}(p)
	}
	syncWg.Wait()
	loger.Infof("-> Generating \"internal/repeated.go\" ...")
	if err = ExecuteTemplateTo("internal/repeated.go", "repeated.go.gotemp", map[string]any{
		"repeats": repeatPackets,
	}); err != nil {
		loger.Fatal(err)
	}
	wg.Wait()
}

var packetTableHeads = []string{"Packet ID", "State", "Bound To", "Field Name", "Field Type", "Notes"}

func genProtocol(p Protocol, syncWg *sync.WaitGroup)(err error){
	needDoneWg := true
	defer func(){
		if needDoneWg {
			syncWg.Done()
		}
	}()

	dirname := "v" + p.Name
	if err = os.MkdirAll(dirname, 0755); err != nil {
		return
	}

	res, err := GetOrUseCached(p.Doc, "generator/cache/oldid_" + p.docid + ".html")
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(res)
	res.Close()
	if err != nil {
		return
	}

	type (
		boundMap struct {
			Client []*packetInfo
			Server []*packetInfo
		}
		stateMap map[string]*boundMap
	)

	states := make(stateMap, 2)
	states["login"] = new(boundMap)
	states["play"] = new(boundMap)
	{
		var (
			state string
			boundTo int // 1: client; 2: server
			nameSet = make(map[string]*packetInfo)
		)
		cur := doc.Find("#mw-content-text > .mw-parser-output").Children().First();
		for ; cur.Length() > 0; cur = cur.Next() {
			if cur.Is("h2") {
				switch s := strings.ToLower(cur.Text()); s {
				case "handshaking", "status": // should be defined in common.go
					state = ""
				case "login", "play":
					state = s
				default:
					state = ""
				}
			}else if cur.Is("h3") {
				if state != "" {
					switch strings.ToLower(cur.Text()) {
					case "clientbound":
						boundTo = 1
					case "serverbound":
						boundTo = 2
					default:
						boundTo = 0
					}
				}
			}else if cur.Is("table") {
				if state != "" && boundTo != 0 {
					if tb := ParseTableFromSelection(cur); tb != nil {
						bm := states[state]
						pinfo := &packetInfo{
							State: state,
							Name: joinWords(addPrefix(cur.PrevAllFiltered("h4").First().Text(), state + " ")),
							Table: tb.String(),
						}
						if boundTo == 1 {
							pinfo.Bound = "client"
						}else{
							pinfo.Bound = "server"
						}
						// Packet ID | State | Bound To | Field Name | Field Type | Notes
						if tb.HasHeads(packetTableHeads) {
							_, err := fmt.Sscanf(tb.Rows[0][0], "0x%x\n", &pinfo.Id)
							if err != nil {
								loger.Fatalf("Cannot parse packet ID %q: %v", tb.Rows[0][0], err)
							}
							loger.Debugf("State = %s; BoundTo = %d; Packet Name = %s; Packet ID = 0x%x", pinfo.Name, state, boundTo, pinfo.Id)
							if tb.Rows[0][3] != "no fields" {
								for _, r := range tb.Rows {
									if r[3] == "" {
										continue
									}
									pinfo.Fields = append(pinfo.Fields, packetField{
										Name: joinWords(strings.ReplaceAll(r[3], "/", " Or ")),
										Type: wikiTypeAsGoType(r[4]),
										WikiType: r[4],
										Note: r[5],
									})
								}
							}
							pinfo.Matched = true
							registerPacketInfo(pinfo, p.Version)
						}else if strings.ToLower(tb.Heads[0]) == strings.ToLower(packetTableHeads[0]) {
							_, err := fmt.Sscanf(tb.Rows[0][0], "0x%x\n", &pinfo.Id)
							if err != nil {
								loger.Fatalf("Cannot parse packet ID %q: %v", tb.Rows[0][0], err)
							}
						}else{
							continue
						}
						if other, ok := nameSet[pinfo.Name]; ok {
							if boundTo == 1 {
								pinfo.Name = addSuffix(pinfo.Name, "Client")
								other.Name = addSuffix(other.Name, "Server")
							}else{
								pinfo.Name = addSuffix(pinfo.Name, "Server")
								other.Name = addSuffix(other.Name, "Client")
							}
						}
						nameSet[pinfo.Name] = pinfo
						if boundTo == 1 {
							bm.Client = append(bm.Client, pinfo)
						}else{
							bm.Server = append(bm.Server, pinfo)
						}
					}
				}
			}
		}
	}

	syncWg.Done()
	needDoneWg = false
	syncWg.Wait()

	return ExecuteTemplateTo(filepath.Join(dirname, "packet.go"), "packet.go.gotemp", map[string]any{
		"protocol": p,
		"states": states,
	})
}

func generateValueEncodeCode(w io.Writer, name string, typ string, indent int){
	indents := strings.Repeat("\t", indent)
	fmt.Fprint(w, indents)
	if isBasicType(typ) {
		fmt.Fprintf(w, "b.%s(%s)\n", typ, name)
	}else if typ == "Object" {
		fmt.Fprintf(w, "b.JSON(%s)\n", name)
	}else if strings.HasPrefix(typ, "[]") {
		fmt.Fprintf(w, "TODO_Encode_Array(%s)\n", name)
	}else if typ == "ByteArray" {
		fmt.Fprintf(w, "b.ByteArray(%s)\n", name)
	}else{
		fmt.Fprintf(w, "b.Encode(%s)\n", name)
	}
}

func renderPacketEncodeMethod(fields []packetField)(string){
	var s strings.Builder
	s.WriteByte('\n')
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		if isnum := isNumberType(field.Type); isnum || isBasicType(field.Type) {
			if isnum && i + 1 < len(fields) {
				next := fields[i + 1]
				arrName, ok := strings.CutSuffix(field.Name, "Count")
				if ok {
					ok = strings.HasPrefix(next.Name, arrName)
				}
				if !ok {
					arrName, ok = strings.CutSuffix(field.Name, "Length")
					if ok {
						ok = strings.HasPrefix(next.Name, arrName)
					}
				}
				if !ok {
					arrName, ok = strings.CutPrefix(field.Name, "NumberOf")
					if ok {
						ok = strings.HasPrefix(next.Name, arrName)
					}
				}
				if ok {
					if next.Type == "ByteArray" {
						fmt.Fprintf(&s, "\tp.%s = len(p.%s)\n", field.Name, next.Name)
						fmt.Fprintf(&s, "\tb.%s(p.%s)\n", field.Type, field.Name)
						fmt.Fprintf(&s, "\tb.ByteArray(p.%s)\n", next.Name)
						i++
						continue
					}else if elem, ok := strings.CutPrefix(next.Type, "[]"); ok {
						fmt.Fprintf(&s, "\tp.%s = len(p.%s)\n", field.Name, next.Name)
						fmt.Fprintf(&s, "\tb.%s(p.%s)\n", field.Type, field.Name)
						fmt.Fprintf(&s, "\tfor _, v := range p.%s {\n", next.Name)
						generateValueEncodeCode(&s, "v", elem, 2)
						fmt.Fprintf(&s, "\t}\n")
						i++
						continue
					}
				}
			}
			fmt.Fprintf(&s, "\tb.%s(p.%s)\n", field.Type, field.Name)
		}else if field.Type == "Object" {
			fmt.Fprintf(&s, "\tb.JSON(p.%s)\n", field.Name)
		}else if strings.HasPrefix(field.Type, "[]") {
			fmt.Fprintf(&s, "\tTODO_Encode_Array(p.%s)\n", field.Name)
		}else if field.Type == "ByteArray" {
			if i == len(fields) - 1 {
				fmt.Fprintf(&s, "\tb.ReadAll(p.%s)\n", field.Name)
			}else{
				fmt.Fprintf(&s, "\tb.ByteArray(p.%s)\n", field.Name)
			}
		}else{
			fmt.Fprintf(&s, "\tb.Encode(p.%s)\n", field.Name)
		}
	}
	return s.String()
}

func genRegularDecodeCode(w io.Writer, name, typ string, indent int){
	indents := strings.Repeat("\t", indent)
	fmt.Fprint(w, indents)
	if isBasicType(typ) {
		fmt.Fprintf(w, "if p.%s, ok = r.%s(); !ok {\n", name, typ)
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "\treturn io.EOF\n")
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "}\n")
	}else if typ == "Object" {
		fmt.Fprintf(w, "if err = r.JSON(&p.%s); err != nil {\n", name)
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "\treturn\n")
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "}\n")
	}else if strings.HasPrefix(typ, "[]") {
		fmt.Fprintf(w, "TODO_Decode_Array(%s)\n", name)
	}else if typ == "ByteArray" {
		fmt.Fprintf(w, "TODO_Decode_ByteArray(%s)\n", name)
	}else{
		fmt.Fprintf(w, "if err = p.%s.DecodeFrom(r); err != nil {\n", name)
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "\treturn\n")
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "}\n")
	}
}

func renderPacketDecodeMethod(fields []packetField)(string){
	if len(fields) == 0 {
		return " return "
	}
	var s strings.Builder
	fmt.Fprintf(&s, "\n\tvar ok bool\n")
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		if isnum := isNumberType(field.Type); isnum || isBasicType(field.Type) {
			if isnum && i + 1 < len(fields) {
				next := fields[i + 1]
				arrName, ok := strings.CutSuffix(field.Name, "Count")
				if ok {
					ok = strings.HasPrefix(next.Name, arrName) || strings.HasSuffix(next.Name, arrName)
				}
				if !ok {
					arrName, ok = strings.CutSuffix(field.Name, "Length")
					if ok {
						ok = strings.HasPrefix(next.Name, arrName) || strings.HasSuffix(next.Name, arrName)
					}
				}
				if !ok {
					arrName, ok = strings.CutPrefix(field.Name, "NumberOf")
					if ok {
						ok = strings.HasPrefix(next.Name, arrName) || strings.HasSuffix(next.Name, arrName)
					}
				}
				if ok {
					if next.Type == "ByteArray" {
						genRegularDecodeCode(&s, field.Name, field.Type, 1)
						fmt.Fprintf(&s, "\tp.%s = make(ByteArray, p.%s)\n", next.Name, field.Name)
						fmt.Fprintf(&s,
`	if ok = r.ByteArray(p.%s); !ok {
		return io.EOF
	}
`, next.Name)
						i++
						continue
					}else if elem, ok := strings.CutPrefix(next.Type, "[]"); ok {
						genRegularDecodeCode(&s, field.Name, field.Type, 1)
						fmt.Fprintf(&s, "\tp.%s = make(%s, p.%s)\n", next.Name, next.Type, field.Name)
						fmt.Fprintf(&s, "\tfor i, _ := range p.%s {\n", next.Name)
						genRegularDecodeCode(&s, next.Name + "[i]", elem, 2)
						fmt.Fprintf(&s, "\t}\n")
						i++
						continue
					}
				}
			}
			genRegularDecodeCode(&s, field.Name, field.Type, 1)
		}else if field.Type == "Object" {
			fmt.Fprintf(&s,
`	if err = r.JSON(&p.%s); err != nil {
		return
	}
`, field.Name)
		}else if strings.HasPrefix(field.Type, "[]") {
			fmt.Fprintf(&s, "\tTODO_Decode_Array(p.%s)\n", field.Name)
		// }else if field.Type == "ByteArray" {
		// 	if i == len(fields) - 1 {
		// 		fmt.Fprintf(&s, "\tb.ReadAll(p.%s)\n", field.Name)
		// 	}else{
		// 		fmt.Fprintf(&s, "\tr.ByteArray(p.%s)\n", field.Name)
		// 	}
		}else{
			fmt.Fprintf(&s,
`	if err = p.%s.DecodeFrom(r); err != nil {
		return
	}
`, field.Name)
		}
	}
	return s.String()
}

func split(s string, b byte)(l, r string){
	i := strings.IndexByte(s, b)
	if i < 0 {
		return s, ""
	}
	return s[:i], s[i:]
}

func addPrefix(s string, prefix string)(string){
	if strings.HasPrefix(s, prefix) {
		return s
	}
	return prefix + s
}

func addSuffix(s string, suffix string)(string){
	if strings.HasSuffix(s, suffix) {
		return s
	}
	return s + suffix
}

func joinWords(line string)(string){
	line, _ = split(line, '(')
	line = strings.ReplaceAll(line, "?", "")
	line = strings.ReplaceAll(line, "-", " ")
	words := strings.Fields(line)
	for i, v := range words {
		words[i] = strings.Title(v)
	}
	return strings.Join(words, "")
}
