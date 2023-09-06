
package data

import (
	"io"

	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/nbt"
)

type Slot struct {
	Present bool
	ItemId int
	ItemCount liter.Byte
	NBT nbt.NBT
}

var _ liter.Packet = (*Slot)(nil)

func (c Slot)Encode(b *liter.PacketBuilder){
	if b.Protocol() >= liter.V1_13_2 {
		b.Bool(c.Present)
		if c.Present {
			b.VarInt((liter.VarInt)(c.ItemId))
			b.Byte(c.ItemCount)
			nbt.WriteNBT(b, c.NBT)
		}
	}else{
		b.Short((liter.Short)(c.ItemId))
		if c.ItemId >= 0 {
			b.Byte(c.ItemCount)
			nbt.WriteNBT(b, c.NBT)
		}
	}
}

func (c *Slot)DecodeFrom(r *liter.PacketReader)(err error){
	var ok bool
	if r.Protocol() >= liter.V1_13_2 {
		if c.Present, ok = r.Bool(); !ok {
			return io.EOF
		}
		if c.Present {
			var id liter.VarInt
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
	}else{
		var id liter.Short
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
