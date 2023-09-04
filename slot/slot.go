
package slot

type Slot struct {
	Present bool
	ItemId VarInt
	ItemCount Byte
	NBT NBT
}

var _ Packet = (*Slot)(nil)

func (c Slot)Encode(b *PacketBuilder){
	//
}

func (c *Slot)DecodeFrom(r *PacketReader)(err error){
	return
}
