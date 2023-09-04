
package nbt

import "io"

type NBTList struct {
	name string
	Data []NBT
}

func (n *NBTList)Type()(Byte){ return NbtList }
func (n *NBTList)Name()(string){ return n.name }
func (n *NBTList)SetName(name string){ n.name = name }

func (n *NBTList)Elem()(Byte){
	if len(n.Data) == 0 {
		return NbtEnd
	}
	return n.Data[0].Type()
}

func (n *NBTList)Encode(b *PacketBuilder){
	b.Byte(n.Elem())
	b.Int((Int)(len(n.Data)))
	for _, v := range n.Data {
		v.Encode(b)
	}
}

func (n *NBTList)DecodeFrom(r *PacketReader)(err error){
	var (
		ok bool
		id Byte
		l Int
	)
	if id, ok = r.Byte(); !ok {
		return io.EOF
	}
	if l, ok = r.Int(); !ok {
		return io.EOF
	}
	if l <= 0 {
		n.Data = nil
		return
	}
	var newer func()(NBT)
	if newer, ok = NBTNewer[id]; !ok {
		return &NBTIdNotExistsErr{ id }
	}
	n.Data = make([]NBT, l)
	for i, _ := range n.Data {
		v := newer()
		if err = v.DecodeFrom(r); err != nil {
			return
		}
		n.Data[i] = v
	}
	return
}

type NBTIntArray struct {
	name string
	Data []Int
}

func (n *NBTIntArray)Type()(Byte){ return NbtIntArray }
func (n *NBTIntArray)Name()(string){ return n.name }
func (n *NBTIntArray)SetName(name string){ n.name = name }

func (n *NBTIntArray)Encode(b *PacketBuilder){
	b.Int((Int)(len(n.Data)))
	for _, v := range n.Data {
		b.Int(v)
	}
}

func (n *NBTIntArray)DecodeFrom(r *PacketReader)(err error){
	var (
		ok bool
		l Int
	)
	if l, ok = r.Int(); !ok {
		return io.EOF
	}
	if l <= 0 {
		n.Data = nil
		return
	}
	n.Data = make([]Int, l)
	for i, _ := range n.Data {
		if n.Data[i], ok = r.Int(); !ok {
			return io.EOF
		}
	}
	return
}

type NBTLongArray struct {
	name string
	Data []Long
}

func (n *NBTLongArray)Type()(Byte){ return NbtLongArray }
func (n *NBTLongArray)Name()(string){ return n.name }
func (n *NBTLongArray)SetName(name string){ n.name = name }

func (n *NBTLongArray)Encode(b *PacketBuilder){
	b.Long((Long)(len(n.Data)))
	for _, v := range n.Data {
		b.Long(v)
	}
}

func (n *NBTLongArray)DecodeFrom(r *PacketReader)(err error){
	var (
		ok bool
		l Int
	)
	if l, ok = r.Int(); !ok {
		return io.EOF
	}
	if l <= 0 {
		n.Data = nil
		return
	}
	n.Data = make([]Long, l)
	for i, _ := range n.Data {
		if n.Data[i], ok = r.Long(); !ok {
			return io.EOF
		}
	}
	return
}
