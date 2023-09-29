
package liter

import (
	"io"
)

type BitSet []uint64

var _ Packet = (*BitSet)(nil)

func NewBitSet()(s *BitSet){
	s = new(BitSet)
	*s = make(BitSet, 0)
	return
}

func NewBitSetWith(size int)(s *BitSet){
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

func (s *BitSet)Set(index int, value bool){
	i, b := index / 64, (uint64)(1) << (index % 64)
	arr := *s
	if i >= len(arr) {
		for i >= len(arr) {
			arr = append(arr, 0)
		}
		*s = arr
	}
	if value {
		arr[i] |= b
	}else{
		arr[i] &^= b
	}
}

func (s *BitSet)Open(index int){
	s.Set(index, true)
}

func (s *BitSet)Close(index int){
	s.Set(index, true)
}

func (s *BitSet)Get(index int)(value bool){
	i, b := index / 64, (uint64)(1) << (index % 64)
	if i >= len(*s) {
		return false
	}
	return (*s)[i] & b != 0
}

func (s *BitSet)Clear(){
	for i, _ := range *s {
		(*s)[i] = 0
	}
}
