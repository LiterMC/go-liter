
package data

import (
	"github.com/kmcsr/go-liter"
)

type EntityMetadata struct {
	Index liter.UByte
	Type liter.VarInt
	Value any
}

var _ liter.Packet = (*EntityMetadata)(nil)

func (c EntityMetadata)Encode(b *liter.PacketBuilder){
	//
}

func (c *EntityMetadata)DecodeFrom(r *liter.PacketReader)(err error){
	return
}

