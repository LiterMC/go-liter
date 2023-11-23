package nbt

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type NBTByte struct {
	name string
	Data Byte
}

func (n *NBTByte) Type() Byte          { return NbtByte }
func (n *NBTByte) Name() string        { return n.name }
func (n *NBTByte) SetName(name string) { n.name = name }

func (n *NBTByte) String() string {
	var s strings.Builder
	s.WriteString("Tag_Byte(")
	if len(n.name) == 0 {
		s.WriteString("None")
	} else {
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	fmt.Fprintf(&s, "): %d", n.Data)
	return s.String()
}

func (n *NBTByte) Encode(b *PacketBuilder) {
	b.Byte(n.Data)
}

func (n *NBTByte) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if n.Data, ok = r.Byte(); !ok {
		return io.EOF
	}
	return
}

type NBTShort struct {
	name string
	Data Short
}

func (n *NBTShort) Type() Byte          { return NbtShort }
func (n *NBTShort) Name() string        { return n.name }
func (n *NBTShort) SetName(name string) { n.name = name }

func (n *NBTShort) String() string {
	var s strings.Builder
	s.WriteString("Tag_Short(")
	if len(n.name) == 0 {
		s.WriteString("None")
	} else {
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	fmt.Fprintf(&s, "): %d", n.Data)
	return s.String()
}

func (n *NBTShort) Encode(b *PacketBuilder) {
	b.Short(n.Data)
}

func (n *NBTShort) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if n.Data, ok = r.Short(); !ok {
		return io.EOF
	}
	return
}

type NBTInt struct {
	name string
	Data Int
}

func (n *NBTInt) Type() Byte          { return NbtInt }
func (n *NBTInt) Name() string        { return n.name }
func (n *NBTInt) SetName(name string) { n.name = name }

func (n *NBTInt) String() string {
	var s strings.Builder
	s.WriteString("Tag_Int(")
	if len(n.name) == 0 {
		s.WriteString("None")
	} else {
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	fmt.Fprintf(&s, "): %d", n.Data)
	return s.String()
}

func (n *NBTInt) Encode(b *PacketBuilder) {
	b.Int(n.Data)
}

func (n *NBTInt) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if n.Data, ok = r.Int(); !ok {
		return io.EOF
	}
	return
}

type NBTLong struct {
	name string
	Data Long
}

func (n *NBTLong) Type() Byte          { return NbtLong }
func (n *NBTLong) Name() string        { return n.name }
func (n *NBTLong) SetName(name string) { n.name = name }

func (n *NBTLong) String() string {
	var s strings.Builder
	s.WriteString("Tag_Long(")
	if len(n.name) == 0 {
		s.WriteString("None")
	} else {
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	fmt.Fprintf(&s, "): %dL", n.Data)
	return s.String()
}

func (n *NBTLong) Encode(b *PacketBuilder) {
	b.Long(n.Data)
}

func (n *NBTLong) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if n.Data, ok = r.Long(); !ok {
		return io.EOF
	}
	return
}

type NBTFloat struct {
	name string
	Data Float
}

func (n *NBTFloat) Type() Byte          { return NbtFloat }
func (n *NBTFloat) Name() string        { return n.name }
func (n *NBTFloat) SetName(name string) { n.name = name }

func (n *NBTFloat) String() string {
	var s strings.Builder
	s.WriteString("Tag_Float(")
	if len(n.name) == 0 {
		s.WriteString("None")
	} else {
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	s.WriteString("): ")
	s.WriteString(strconv.FormatFloat((float64)(n.Data), 'f', -1, 32))
	return s.String()
}

func (n *NBTFloat) Encode(b *PacketBuilder) {
	b.Float(n.Data)
}

func (n *NBTFloat) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if n.Data, ok = r.Float(); !ok {
		return io.EOF
	}
	return
}

type NBTDouble struct {
	name string
	Data Double
}

func (n *NBTDouble) Type() Byte          { return NbtDouble }
func (n *NBTDouble) Name() string        { return n.name }
func (n *NBTDouble) SetName(name string) { n.name = name }

func (n *NBTDouble) String() string {
	var s strings.Builder
	s.WriteString("Tag_Double(")
	if len(n.name) == 0 {
		s.WriteString("None")
	} else {
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	s.WriteString("): ")
	s.WriteString(strconv.FormatFloat(n.Data, 'f', -1, 64))
	return s.String()
}

func (n *NBTDouble) Encode(b *PacketBuilder) {
	b.Double(n.Data)
}

func (n *NBTDouble) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if n.Data, ok = r.Double(); !ok {
		return io.EOF
	}
	return
}
