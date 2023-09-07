
package main

import (
	"embed"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
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
		Table *Table

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
				switch name { // do not support to generate these versions
				case "1.13.1":
					loger.Warnf("Skipping %s ...", name)
					return
				}
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

	// ensure protocols is sorted from high to low
	sort.Slice(protocols, func(i, j int)(bool){ return protocols[i].Version > protocols[j].Version })

	// generate protocols values in the root package
	if err = ExecuteTemplateTo("../protocols.go", "protocols.go.gotemp", map[string]any{
		"protocols": protocols,
	}); err != nil {
		loger.Fatal(err)
	}
	// panic("do fixing")

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
							Name: addPrefix(joinWords(cur.PrevAllFiltered("h4").First().Text()), strings.Title(state)),
							Table: tb,
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
									field := packetField{
										Name: joinWords(strings.ReplaceAll(r[3], "/", " Or ")),
										WikiType: r[4],
										Note: r[5],
									}
									if r[4] == "(See below)" {
										if strings.HasSuffix(field.Name, "Tags") {
											field.Type = "[]Tag"
										}else{
											field.Type = "TODO_See_Below"
										}
									}else{
										field.Type = wikiTypeAsGoType(r[4])
									}
									pinfo.Fields = append(pinfo.Fields, field)
								}
							}
							pinfo.Matched = true
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
						registerPacketInfo(pinfo, p.Version)
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
	if typ == "Angle" {
		typ = "UByte"
	}
	if isBasicType(typ) {
		fmt.Fprintf(w, "b.%s(%s)\n", typ, name)
	}else if typ == "Object" {
		fmt.Fprintf(w, "b.JSON(%s)\n", name)
	}else if typ == "nbt.NBT" {
		fmt.Fprintf(w, "nbt.WriteNBT(b, %s)\n", name)
	}else if strings.HasPrefix(typ, "[]") {
		fmt.Fprintf(w, "TODO_Encode_Array(%s)\n", name)
	}else if typ == "ByteArray" {
		fmt.Fprintf(w, "b.ByteArray(%s)\n", name)
	}else if cutted, ok := strings.CutPrefix(typ, "Optional["); ok {
		cutted = cutted[:len(cutted) - 1] // remove ']'
		fmt.Fprintf(w, "if %s.Ok = TODO; %s.Ok {\n", name, name)
		generateValueEncodeCode(w, name + ".V", cutted, indent + 1)
		fmt.Fprint(w, indents)
		fmt.Fprintln(w, "}")
	}else{
		fmt.Fprintf(w, "%s.Encode(b)\n", name)
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
					ok = strings.HasPrefix(next.Name, arrName) || strings.HasSuffix(next.Name, arrName)
				}
				if !ok {
					arrName, ok = strings.CutSuffix(field.Name, "Size")
					if ok {
						ok = strings.HasPrefix(next.Name, arrName) || strings.HasSuffix(next.Name, arrName)
					}
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
						fmt.Fprintf(&s, "\tp.%s = (%s)(len(p.%s))\n", field.Name, field.Type, next.Name)
						fmt.Fprintf(&s, "\tb.%s(p.%s)\n", field.Type, field.Name)
						fmt.Fprintf(&s, "\tb.ByteArray(p.%s)\n", next.Name)
						i++
						continue
					}else if elem, ok := strings.CutPrefix(next.Type, "[]"); ok {
						fmt.Fprintf(&s, "\tp.%s = (%s)(len(p.%s))\n", field.Name, field.Type, next.Name)
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
		}else{
			generateValueEncodeCode(&s, "p." + field.Name, field.Type, 1)
		}
	}
	return s.String()
}

func genRegularDecodeCode(w io.Writer, name, typ string, indent int){
	indents := strings.Repeat("\t", indent)
	if cutted, ok := strings.CutPrefix(typ, "*"); ok {
		fmt.Fprint(w, indents)
		fmt.Fprintf(w, "p.%s = new(%s)\n", name, cutted)
		typ = cutted
	}
	fmt.Fprint(w, indents)
	if typ == "Angle" {
		typ = "UByte"
	}
	if isBasicType(typ) {
		fmt.Fprintf(w, "if p.%s, ok = r.%s(); !ok {\n", name, typ)
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "\treturn io.EOF\n")
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "}\n")
	}else if typ == "Object" {
		fmt.Fprintf(w, "if err = r.JSON(&p.%s); err != nil {\n", name)
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "\treturn err\n")
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "}\n")
	}else if typ == "nbt.NBT" {
		fmt.Fprintf(w, "if p.%s, err = nbt.ReadNBT(r); err != nil {\n", name)
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "\treturn err\n")
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "}\n")
	}else if strings.HasPrefix(typ, "[]") {
		fmt.Fprintf(w, "TODO_Decode_Array(p.%s)\n", name)
	}else if typ == "ByteArray" {
		fmt.Fprintf(w, "TODO_Decode_ByteArray(p.%s)\n", name)
	}else if cutted, ok := strings.CutPrefix(typ, "Optional["); ok {
		cutted = cutted[:len(cutted) - 1] // remove ']'
		fmt.Fprintf(w, "if p.%s.Ok = TODO; p.%s.Ok {\n", name, name)
		genRegularDecodeCode(w, name + ".V", cutted, indent + 1)
		fmt.Fprint(w, indents)
		fmt.Fprintln(w, "}")
	}else{
		fmt.Fprintf(w, "if err = p.%s.DecodeFrom(r); err != nil {\n", name)
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "\treturn err\n")
		fmt.Fprint(w, indents)
		fmt.Fprint(w, "}\n")
	}
}

func renderPacketDecodeMethod(fields []packetField)(string){
	if len(fields) == 0 {
		return " return nil "
	}
	var s strings.Builder
	fmt.Fprintln(&s, "\n\tvar ok bool")
	fmt.Fprintln(&s, "\t_ = ok")
	fmt.Fprintln(&s, "\tvar err error")
	fmt.Fprintln(&s, "\t_ = err")
	for i := 0; i < len(fields); i++ {
		field := fields[i]
		if cutted, ok := strings.CutPrefix(field.Type, "*"); ok {
			fmt.Fprintf(&s, "\tp.%s = new(%s)\n", field.Name, cutted)
			field.Type = cutted
		}
		if isnum := isNumberType(field.Type); isnum || isBasicType(field.Type) {
			if isnum && i + 1 < len(fields) {
				next := fields[i + 1]
				arrName, ok := strings.CutSuffix(field.Name, "Count")
				if ok {
					ok = strings.HasPrefix(next.Name, arrName) || strings.HasSuffix(next.Name, arrName)
				}
				if !ok {
					arrName, ok = strings.CutSuffix(field.Name, "Size")
					if ok {
						ok = strings.HasPrefix(next.Name, arrName) || strings.HasSuffix(next.Name, arrName)
					}
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
						fmt.Fprintf(&s, "\tif ok = r.ByteArray(p.%s); !ok {\n", next.Name)
						fmt.Fprintln(&s, "\t\treturn io.EOF")
						fmt.Fprintln(&s, "\t}")
						i++
						continue
					}else if elem, ok := strings.CutPrefix(next.Type, "[]"); ok {
						genRegularDecodeCode(&s, field.Name, field.Type, 1)
						fmt.Fprintf(&s, "\tp.%s = make(%s, p.%s)\n", next.Name, next.Type, field.Name)
						fmt.Fprintf(&s, "\tfor i, _ := range p.%s {\n", next.Name)
						genRegularDecodeCode(&s, next.Name + "[i]", elem, 2)
						fmt.Fprintln(&s, "\t}")
						i++
						continue
					}
				}
			}
			genRegularDecodeCode(&s, field.Name, field.Type, 1)
		}else if field.Type == "ByteArray" {
			if i == len(fields) - 1 {
				fmt.Fprintf(&s, "\tp.%s = r.ReadAll()\n", field.Name)
			}else{
				fmt.Fprintf(&s, "\tTODO_Decode_ByteArray(p.%s)\n", field.Name)
			}
		}else{
			genRegularDecodeCode(&s, field.Name, field.Type, 1)
		}
	}
	fmt.Fprint(&s, "\treturn nil\n")
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
