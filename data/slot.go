package data

import (
	"io"

	. "github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/nbt"
)

type Slot struct {
	Present   bool
	ItemId    int
	ItemCount Byte
	NBT       nbt.NBT
}

var _ Packet = (*Slot)(nil)

func (c Slot) Encode(b *PacketBuilder) {
	if b.Protocol() >= V1_13_2 {
		b.Bool(c.Present)
		if c.Present {
			b.VarInt((VarInt)(c.ItemId))
			b.Byte(c.ItemCount)
			nbt.WriteNBT(b, c.NBT)
		}
	} else {
		b.Short((Short)(c.ItemId))
		if c.ItemId >= 0 {
			b.Byte(c.ItemCount)
			nbt.WriteNBT(b, c.NBT)
		}
	}
}

func (c *Slot) DecodeFrom(r *PacketReader) (err error) {
	var ok bool
	if r.Protocol() >= V1_13_2 {
		if c.Present, ok = r.Bool(); !ok {
			return io.EOF
		}
		if c.Present {
			var id VarInt
			if id, ok = r.VarInt(); !ok {
				return io.EOF
			}
			c.ItemId = (int)(id)
			if c.ItemCount, ok = r.Byte(); !ok {
				return io.EOF
			}
			if c.NBT, err = nbt.ReadNBT(r); err != nil {
				return
			}
		}
	} else {
		var id Short
		if id, ok = r.Short(); !ok {
			return io.EOF
		}
		c.ItemId = (int)(id)
		if c.Present = c.ItemId >= 0; c.Present {
			if c.ItemCount, ok = r.Byte(); !ok {
				return io.EOF
			}
			if c.NBT, err = nbt.ReadNBT(r); err != nil {
				return
			}
		}
	}
	return
}
