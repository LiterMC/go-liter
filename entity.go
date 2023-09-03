
package liter

type EntityMetadata struct {
	Index UByte
	Type VarInt
	Value any
}

var _ Packet = (*EntityMetadata)(nil)

func (c EntityMetadata)Encode(b *PacketBuilder){
	//
}

func (c *EntityMetadata)DecodeFrom(r *PacketReader)(err error){
	return
}

