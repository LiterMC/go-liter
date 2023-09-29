
package console

import (
	"io"
	"strconv"

	"github.com/dop251/goja"
)

type writer interface {
	io.Writer
	io.StringWriter
}

func encode(w writer, v goja.Value){
	encode0(w, v, "", make(map[*goja.Object]struct{}))
}

const indent = "  "

func encode0(w writer, v goja.Value, indents string, emap map[*goja.Object]struct{}){
	var buf [128]byte
	switch v := v.(type) {
	case *goja.Object:
		keys := v.Keys()
		if len(keys) == 0 {
			if _, ok := goja.AssertFunction(v); ok {
				w.WriteString(v.String())
				return
			}
			w.WriteString(v.ClassName())
			w.WriteString(" {}")
			return
		}
		if _, ok := emap[v]; ok {
			w.WriteString(v.String())
			return
		}
		w.WriteString(v.ClassName())
		w.WriteString(" {\n")
		emap[v] = struct{}{}
		indents2 := indents + indent
		for _, k := range keys {
			c := v.Get(k)
			w.WriteString(indents2)
			w.WriteString(k)
			w.WriteString(": ")
			encode0(w, c, indents2, emap)
			w.WriteString(",\n")
		}
		delete(emap, v)
		w.WriteString(indents)
		w.WriteString("}")
	case goja.String:
		w.Write(strconv.AppendQuote(buf[:0], v.String()))
	default:
		w.WriteString(v.String())
	}
}
