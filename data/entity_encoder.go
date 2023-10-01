
package data

import (
	"fmt"
	"io"

	. "github.com/kmcsr/go-liter"
)

type Rotation struct {
	X, Y, Z Float
}

// Down = 0, Up = 1, North = 2, South = 3, West = 4, East = 5
type Direction VarInt

const (
	DirDown  Direction = 0
	DirUp    Direction = 1
	DirNorth Direction = 2
	DirSouth Direction = 3
	DirWest  Direction = 4
	DirEast  Direction = 5
)

func (d Direction)String()(string){
	switch d {
	case DirDown:  return "Down"
	case DirUp:    return "Up"
	case DirNorth: return "North"
	case DirSouth: return "South"
	case DirWest:  return "West"
	case DirEast:  return "East"
	default: return "<Invalid Direction>"
	}
}

type (
	optMeta[T EntityPropEncoder] struct{
		typ VarInt
	}
	boolMeta struct{}
	byteMeta struct{}
	varIntMeta struct{}
	varLongMeta struct{}
	floatMeta struct{}
	stringMeta struct{}
	chatMeta struct{}
	slotMeta struct{}
	rotationMeta struct{}
	positionMeta struct{}
	directionMeta struct{}
)

var _ EntityPropEncoder = (*optMeta[EntityPropEncoder])(nil)

func (m *optMeta[T])Type()(VarInt){ return m.typ }

func (*optMeta[T])Encode(b *PacketBuilder, value any)(err error){
	vok := value != nil
	b.Bool(vok)
	if vok {
		var encoder2 T
		encoder2.Encode(b, value)
	}
	return
}

func (*optMeta[T])Decode(r *PacketReader)(value any, err error){
	var ok bool
	var vok bool
	if vok, ok = r.Bool(); !ok {
		return nil, io.EOF
	}
	if vok {
		var encoder2 T
		return encoder2.Decode(r)
	}
	return
}

func (boolMeta)   Type()(VarInt){ return 8 }
func (byteMeta)   Type()(VarInt){ return 0 }
func (varIntMeta) Type()(VarInt){ return 1 }
func (varLongMeta)Type()(VarInt){ return 2 }
func (floatMeta)  Type()(VarInt){ return 3 }
func (stringMeta) Type()(VarInt){ return 4 }
func (chatMeta)   Type()(VarInt){ return 5 }

func (boolMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(Bool)
	b.Bool(v)
	return
}

func (byteMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(Byte)
	b.Byte(v)
	return
}

func (varIntMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(VarInt)
	b.VarInt(v)
	return
}

func (varLongMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(VarLong)
	b.VarLong(v)
	return
}

func (floatMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(Float)
	b.Float(v)
	return
}

func (stringMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(String)
	b.String(v)
	return
}

func (chatMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(Object)
	b.JSON(v)
	return
}

func (boolMeta)Decode(r *PacketReader)(value any, err error){
	var ok bool
	if value, ok = r.Bool(); !ok {
		return nil, io.EOF
	}
	return
}

func (byteMeta)Decode(r *PacketReader)(value any, err error){
	var ok bool
	if value, ok = r.Byte(); !ok {
		return nil, io.EOF
	}
	return
}

func (varIntMeta)Decode(r *PacketReader)(value any, err error){
	var ok bool
	if value, ok = r.VarInt(); !ok {
		return nil, io.EOF
	}
	return
}

func (varLongMeta)Decode(r *PacketReader)(value any, err error){
	var ok bool
	if value, ok = r.VarLong(); !ok {
		return nil, io.EOF
	}
	return
}

func (floatMeta)Decode(r *PacketReader)(value any, err error){
	var ok bool
	if value, ok = r.Float(); !ok {
		return nil, io.EOF
	}
	return
}

func (stringMeta)Decode(r *PacketReader)(value any, err error){
	var ok bool
	if value, ok = r.String(); !ok {
		return nil, io.EOF
	}
	return
}

func (chatMeta)Decode(r *PacketReader)(value any, err error){
	var v Object
	if err = r.JSON(&v); err != nil {
		return
	}
	return v, nil
}

func (slotMeta)Type()(VarInt){ return 7 }

func (slotMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(*Slot)
	v.Encode(b)
	return
}

func (slotMeta)Decode(r *PacketReader)(value any, err error){
	var v Slot
	if err = v.DecodeFrom(r); err != nil {
		return
	}
	return &v, nil
}

func (rotationMeta)Type()(VarInt){ return 9 }

func (rotationMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(Rotation)
	b.Float(v.X)
	b.Float(v.Y)
	b.Float(v.Z)
	return
}

func (rotationMeta)Decode(r *PacketReader)(value any, err error){
	var ok bool
	var v Rotation
	if v.X, ok = r.Float(); !ok {
		return nil, io.EOF
	}
	if v.Y, ok = r.Float(); !ok {
		return nil, io.EOF
	}
	if v.Z, ok = r.Float(); !ok {
		return nil, io.EOF
	}
	return v, nil
}

func (positionMeta)Type()(VarInt){ return 10 }

func (positionMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(Position)
	v.Encode(b)
	return
}

func (positionMeta)Decode(r *PacketReader)(value any, err error){
	var v Position
	if err = v.DecodeFrom(r); err != nil {
		return
	}
	return v, nil
}

func (directionMeta)Type()(VarInt){ return 12 }

func (directionMeta)Encode(b *PacketBuilder, value any)(err error){
	v := value.(Direction)
	b.VarInt((VarInt)(v))
	return
}

func (directionMeta)Decode(r *PacketReader)(value any, err error){
	var ok bool
	var v VarInt
	if v, ok = r.VarInt(); !ok {
		return nil, io.EOF
	}
	return (Direction)(v), nil
}

// func (uuidMeta)Type()(VarInt){ panic("unreachable") }

func init(){
	vanillaEncoders := []EntityPropEncoder{
		byteMeta{},
		varIntMeta{},
		varLongMeta{},
		floatMeta{},
		stringMeta{},
		chatMeta{},
		&optMeta[chatMeta]{6},
		slotMeta{},
		boolMeta{},
		rotationMeta{},
		positionMeta{},
		&optMeta[positionMeta]{11},
		directionMeta{},
	}
	for _, e := range vanillaEncoders {
		if !RegisterEntityEncoder(e) {
			panic(fmt.Errorf("Cannot register entity encoder %d", e.Type()))
		}
	}
}
