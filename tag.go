
package liter

type Tag struct {
	//
}

var _ Packet = (*Tag)(nil)

func (c Tag)Encode(b *PacketBuilder){
	//
}

func (c *Tag)DecodeFrom(r *PacketReader)(err error){
	return
}
