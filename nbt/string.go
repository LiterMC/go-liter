package nbt

import (
	"io"
	"strconv"
	"strings"
)

type NBTString struct {
	name string
	Data String
}

func (n *NBTString) Type() Byte          { return NbtString }
func (n *NBTString) Name() string        { return n.name }
func (n *NBTString) SetName(name string) { n.name = name }

func (n *NBTString) String() string {
	var s strings.Builder
	s.WriteString("Tag_String(")
	if len(n.name) == 0 {
		s.WriteString("None")
	} else {
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	s.WriteString("): ")
	s.WriteString(strconv.QuoteToASCII(n.Data))
	return s.String()
}

func (n *NBTString) Encode(b *PacketBuilder) {
	writeNBTString(b, n.Data)
}

func (n *NBTString) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if n.Data, ok = readNBTString(r); !ok {
		return io.EOF
	}
	return
}

func writeNBTString(b *PacketBuilder, s string) {
	buf := EncodeUtf8m(s)
	b.UShort((UShort)(len(buf)))
	b.ByteArray(buf)
}

func readNBTString(r *PacketReader) (s string, ok bool) {
	var l UShort
	if l, ok = r.UShort(); !ok {
		return
	}
	buf := make([]byte, l)
	if ok = r.ByteArray(buf); !ok {
		return
	}
	return DecodeUtf8m(buf), true
}
