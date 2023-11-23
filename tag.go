package liter

import (
	"io"
)

type Tag struct {
	Name    Identifier
	Entries []VarInt
}

var _ Packet = (*Tag)(nil)

func (t Tag) Encode(b *PacketBuilder) {
	b.String(t.Name)
	b.VarInt((VarInt)(len(t.Entries)))
	for _, v := range t.Entries {
		b.VarInt(v)
	}
}

func (t *Tag) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	var n VarInt
	if t.Name, ok = r.String(); !ok {
		return io.EOF
	}
	if n, ok = r.VarInt(); !ok {
		return io.EOF
	}
	t.Entries = make([]VarInt, n)
	for i, _ := range t.Entries {
		if t.Entries[i], ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
	return
}
