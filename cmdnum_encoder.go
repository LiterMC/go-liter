
package liter

import (
	"io"
)

const (
	NumPropMin Byte = 0x01
	NumPropMax Byte = 0x02
)

type numbers interface {
	Byte | Short | Int | Long | Float | Double
}

type NumRangeProp[T numbers] struct {
	flags Byte
	min, max T
}

func (p *NumRangeProp[T])Clear(){
	p.flags = 0
}

func (p *NumRangeProp[T])SetMin(min T)(*NumRangeProp[T]){
	p.flags |= NumPropMin
	p.min = min
	return p
}

func (p *NumRangeProp[T])SetMax(max T)(*NumRangeProp[T]){
	p.flags |= NumPropMax
	p.max = max
	return p
}

func (p *NumRangeProp[T])HasMin()(bool){
	return p.flags & NumPropMin != 0
}

func (p *NumRangeProp[T])HasMax()(bool){
	return p.flags & NumPropMax != 0
}

func (p *NumRangeProp[T])Min()(T){
	if !p.HasMin() {
		panic("Min value is not exists")
	}
	return p.min
}

func (p *NumRangeProp[T])Max()(T){
	if !p.HasMax() {
		panic("Max value is not exists")
	}
	return p.max
}

func (p *NumRangeProp[T])InRange(v T)(bool){
	if p.HasMin() && v < p.min {
		return false
	}
	if p.HasMax() && v > p.max {
		return false
	}
	return true
}

type numRangePropEncoder_Int struct {
	id VarInt
	sid string
}

var _ CmdPropEncoder = (*numRangePropEncoder_Int)(nil)

func (p *numRangePropEncoder_Int)Id()(VarInt){ return p.id }
func (p *numRangePropEncoder_Int)String()(String){ return p.sid }

func (p *numRangePropEncoder_Int)Encode(b *PacketBuilder, value any)(err error){
	v := value.(*NumRangeProp[Int])
	b.Byte(v.flags)
	if v.flags & NumPropMin != 0 {
		b.Int(v.min)
	}
	if v.flags & NumPropMax != 0 {
		b.Int(v.max)
	}
	return
}

func (p *numRangePropEncoder_Int)Decode(r *PacketReader)(value any, err error){
	var ok bool
	v := new(NumRangeProp[Int])
	if v.flags, ok = r.Byte(); !ok {
		return nil, io.EOF
	}
	if v.flags & NumPropMin != 0 {
		if v.min, ok = r.Int(); !ok {
			return nil, io.EOF
		}
	}
	if v.flags & NumPropMax != 0 {
		if v.max, ok = r.Int(); !ok {
			return nil, io.EOF
		}
	}
	return v, nil
}

type numRangePropEncoder_Long struct {
	id VarInt
	sid string
}

var _ CmdPropEncoder = (*numRangePropEncoder_Long)(nil)

func (p *numRangePropEncoder_Long)Id()(VarInt){ return p.id }
func (p *numRangePropEncoder_Long)String()(String){ return p.sid }

func (p *numRangePropEncoder_Long)Encode(b *PacketBuilder, value any)(err error){
	v := value.(*NumRangeProp[Long])
	b.Byte(v.flags)
	if v.flags & NumPropMin != 0 {
		b.Long(v.min)
	}
	if v.flags & NumPropMax != 0 {
		b.Long(v.max)
	}
	return
}

func (p *numRangePropEncoder_Long)Decode(r *PacketReader)(value any, err error){
	var ok bool
	v := new(NumRangeProp[Long])
	if v.flags, ok = r.Byte(); !ok {
		return nil, io.EOF
	}
	if v.flags & NumPropMin != 0 {
		if v.min, ok = r.Long(); !ok {
			return nil, io.EOF
		}
	}
	if v.flags & NumPropMax != 0 {
		if v.max, ok = r.Long(); !ok {
			return nil, io.EOF
		}
	}
	return v, nil
}

type numRangePropEncoder_Float struct {
	id VarInt
	sid string
}

var _ CmdPropEncoder = (*numRangePropEncoder_Float)(nil)

func (p *numRangePropEncoder_Float)Id()(VarInt){ return p.id }
func (p *numRangePropEncoder_Float)String()(String){ return p.sid }

func (p *numRangePropEncoder_Float)Encode(b *PacketBuilder, value any)(err error){
	v := value.(*NumRangeProp[Float])
	b.Byte(v.flags)
	if v.flags & NumPropMin != 0 {
		b.Float(v.min)
	}
	if v.flags & NumPropMax != 0 {
		b.Float(v.max)
	}
	return
}

func (p *numRangePropEncoder_Float)Decode(r *PacketReader)(value any, err error){
	var ok bool
	v := new(NumRangeProp[Float])
	if v.flags, ok = r.Byte(); !ok {
		return nil, io.EOF
	}
	if v.flags & NumPropMin != 0 {
		if v.min, ok = r.Float(); !ok {
			return nil, io.EOF
		}
	}
	if v.flags & NumPropMax != 0 {
		if v.max, ok = r.Float(); !ok {
			return nil, io.EOF
		}
	}
	return v, nil
}

type numRangePropEncoder_Double struct {
	id VarInt
	sid string
}

var _ CmdPropEncoder = (*numRangePropEncoder_Double)(nil)

func (p *numRangePropEncoder_Double)Id()(VarInt){ return p.id }
func (p *numRangePropEncoder_Double)String()(String){ return p.sid }

func (p *numRangePropEncoder_Double)Encode(b *PacketBuilder, value any)(err error){
	v := value.(*NumRangeProp[Double])
	b.Byte(v.flags)
	if v.flags & NumPropMin != 0 {
		b.Double(v.min)
	}
	if v.flags & NumPropMax != 0 {
		b.Double(v.max)
	}
	return
}

func (p *numRangePropEncoder_Double)Decode(r *PacketReader)(value any, err error){
	var ok bool
	v := new(NumRangeProp[Double])
	if v.flags, ok = r.Byte(); !ok {
		return nil, io.EOF
	}
	if v.flags & NumPropMin != 0 {
		if v.min, ok = r.Double(); !ok {
			return nil, io.EOF
		}
	}
	if v.flags & NumPropMax != 0 {
		if v.max, ok = r.Double(); !ok {
			return nil, io.EOF
		}
	}
	return v, nil
}
