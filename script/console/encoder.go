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

func encode(w writer, v goja.Value) {
	switch v := v.(type) {
	case goja.String:
		w.WriteString(v.String())
	default:
		encode0(w, v, "", make(map[*goja.Object]struct{}))
	}
}

const indent = "  "

func encode0(w writer, v goja.Value, indents string, emap map[*goja.Object]struct{}) {
	var buf [128]byte
	switch v := v.(type) {
	case *goja.Object:
		indents2 := indents + indent

		switch e := v.Export().(type) {
		case goja.ArrayBuffer:
			w.WriteString("ArrayBuffer { length = ")
			bts := e.Bytes()
			leng := len(bts)
			w.Write(strconv.AppendUint(buf[:0], (uint64)(leng), 10))
			if leng == 0 {
				w.WriteString(" ")
			} else {
				encodeBytes(w, bts, indents)
				w.WriteString(indents)
			}
			w.WriteString("}")
		default:
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
		}
	case goja.String:
		w.Write(strconv.AppendQuote(buf[:0], v.String()))
	default:
		w.WriteString(v.String())
	}
}

func encodeBytes(w writer, buf []byte, indents string) {
	var bf [2]byte
	const step = 0x10
	indents2 := indents + indent
	w.WriteString("\n")
	for i := 0; i < len(buf); i += step {
		w.WriteString(indents2)
		for j, b := range buf[i:] {
			if j >= step {
				break
			}
			if j == 0x8 {
				w.WriteString("  ")
			} else {
				w.WriteString(" ")
			}
			w.Write(formatByte(bf[:], b))
		}
		w.WriteString("\n")
	}
}

func formatByte(buf []byte, b byte) []byte {
	const hexNumbers = "0123456789abcdef"
	buf[0] = hexNumbers[b>>4]
	buf[1] = hexNumbers[b&0xf]
	return buf[:2]
}
