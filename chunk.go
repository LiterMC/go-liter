
package liter

type ChunkSection struct { // 1.9.4 ~ 1.11.2
	//
}

var _ Packet = (*ChunkSection)(nil)

func (c *ChunkSection)Encode(b *PacketBuilder){
	//
}

func (c *ChunkSection)DecodeFrom(r *PacketReader)(err error){
	return
}
