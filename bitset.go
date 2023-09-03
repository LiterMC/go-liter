
package liter

import (
	"io"
)

type BitSet []uint64

var _ Packet = (*BitSet)(nil)

func NewBitSet(size int)(s *BitSet){
	leng := size / 64
	if size % 64 != 0 {
		leng++
	}
	s = new(BitSet)
	*s = make(BitSet, leng)
	return
}

func (s BitSet)Encode(p *PacketBuilder){
	p.VarInt((VarInt)(len(s)))
	for _, v := range s {
		p.Long((Long)(v))
	}
}

func (s *BitSet)DecodeFrom(r *PacketReader)(err error){
	var (
		ok bool
		l VarInt
		v Long
	)
	if l, ok = r.VarInt(); !ok {
		return io.EOF
	}
	*s = make(BitSet, l)
	for i, _ := range *s {
		if v, ok = r.Long(); !ok {
			return io.EOF
		}
		(*s)[i] = (uint64)(v)
	}
	return
}
