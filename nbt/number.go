
package nbt

import "io"

type NBTByte struct {
	name string
	Data Byte
}

func (n *NBTByte)Type()(Byte){ return NbtByte }
func (n *NBTByte)Name()(string){ return n.name }
func (n *NBTByte)SetName(name string){ n.name = name }

func (n *NBTByte)Encode(b *PacketBuilder){
	b.Byte(n.Data)
}

func (n *NBTByte)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if n.Data, ok = r.Byte(); !ok {
		return io.EOF
	}
	return
}

type NBTShort struct {
	name string
	Data Short
}

func (n *NBTShort)Type()(Byte){ return NbtShort }
func (n *NBTShort)Name()(string){ return n.name }
func (n *NBTShort)SetName(name string){ n.name = name }

func (n *NBTShort)Encode(b *PacketBuilder){
	b.Short(n.Data)
}

func (n *NBTShort)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if n.Data, ok = r.Short(); !ok {
		return io.EOF
	}
	return
}

type NBTInt struct {
	name string
	Data Int
}

func (n *NBTInt)Type()(Byte){ return NbtInt }
func (n *NBTInt)Name()(string){ return n.name }
func (n *NBTInt)SetName(name string){ n.name = name }

func (n *NBTInt)Encode(b *PacketBuilder){
	b.Int(n.Data)
}

func (n *NBTInt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if n.Data, ok = r.Int(); !ok {
		return io.EOF
	}
	return
}

type NBTLong struct {
	name string
	Data Long
}

func (n *NBTLong)Type()(Byte){ return NbtLong }
func (n *NBTLong)Name()(string){ return n.name }
func (n *NBTLong)SetName(name string){ n.name = name }

func (n *NBTLong)Encode(b *PacketBuilder){
	b.Long(n.Data)
}

func (n *NBTLong)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if n.Data, ok = r.Long(); !ok {
		return io.EOF
	}
	return
}

type NBTFloat struct {
	name string
	Data Float
}

func (n *NBTFloat)Type()(Byte){ return NbtFloat }
func (n *NBTFloat)Name()(string){ return n.name }
func (n *NBTFloat)SetName(name string){ n.name = name }

func (n *NBTFloat)Encode(b *PacketBuilder){
	b.Float(n.Data)
}

func (n *NBTFloat)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if n.Data, ok = r.Float(); !ok {
		return io.EOF
	}
	return
}

type NBTDouble struct {
	name string
	Data Double
}

func (n *NBTDouble)Type()(Byte){ return NbtDouble }
func (n *NBTDouble)Name()(string){ return n.name }
func (n *NBTDouble)SetName(name string){ n.name = name }

func (n *NBTDouble)Encode(b *PacketBuilder){
	b.Double(n.Data)
}

func (n *NBTDouble)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if n.Data, ok = r.Double(); !ok {
		return io.EOF
	}
	return
}

type NBTByteArray struct {
	name string
	Data ByteArray
}

func (n *NBTByteArray)Type()(Byte){ return NbtByteArray }
func (n *NBTByteArray)Name()(string){ return n.name }
func (n *NBTByteArray)SetName(name string){ n.name = name }

func (n *NBTByteArray)Encode(b *PacketBuilder){
	b.Int((Int)(len(n.Data)))
	b.ByteArray(n.Data)
}

func (n *NBTByteArray)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	var l Int
	if l, ok = r.Int(); !ok {
		return io.EOF
	}
	if l <= 0 {
		n.Data = nil
	}else{
		n.Data = make(ByteArray, l)
		if ok = r.ByteArray(n.Data); !ok {
			return io.EOF
		}
	}
	return
}
