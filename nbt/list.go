
package nbt

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type NBTList struct {
	name string
	Data []NBT
}

func (n *NBTList)Type()(Byte){ return NbtList }
func (n *NBTList)Name()(string){ return n.name }
func (n *NBTList)SetName(name string){ n.name = name }

func (n *NBTList)String()(string){
	var s strings.Builder
	s.WriteString("Tag_List(")
	if len(n.name) == 0 {
		s.WriteString("None")
	}else{
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	fmt.Fprintf(&s, "): %d ", len(n.Data))
	if len(n.Data) == 1 {
		s.WriteString("entry")
	}else{
		s.WriteString("entries")
	}
	s.WriteString("\n{\n")
	for _, v := range n.Data {
		s.WriteString(addIndent(v.String()))
		s.WriteByte('\n')
	}
	s.WriteString("}")
	return s.String()
}

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

type NBTByteArray struct {
	name string
	Data ByteArray
}

func (n *NBTByteArray)Type()(Byte){ return NbtByteArray }
func (n *NBTByteArray)Name()(string){ return n.name }
func (n *NBTByteArray)SetName(name string){ n.name = name }

func (n *NBTByteArray)String()(string){
	var s strings.Builder
	s.WriteString("Tag_ByteArray(")
	if len(n.name) == 0 {
		s.WriteString("None")
	}else{
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	fmt.Fprintf(&s, "): [%d byte", len(n.Data))
	if len(n.Data) != 1 {
		s.WriteByte('s')
	}
	s.WriteByte(']')
	return s.String()
}

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

type NBTIntArray struct {
	name string
	Data []Int
}

func (n *NBTIntArray)Type()(Byte){ return NbtIntArray }
func (n *NBTIntArray)Name()(string){ return n.name }
func (n *NBTIntArray)SetName(name string){ n.name = name }

func (n *NBTIntArray)String()(string){
	var s strings.Builder
	s.WriteString("Tag_IntArray(")
	if len(n.name) == 0 {
		s.WriteString("None")
	}else{
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	fmt.Fprintf(&s, "): %d ", len(n.Data))
	if len(n.Data) == 1 {
		s.WriteString("entry")
	}else{
		s.WriteString("entries")
	}
	s.WriteString("\n{\n")
	for _, v := range n.Data {
		fmt.Fprintf(&s, "  %d\n", v)
	}
	s.WriteString("}")
	return s.String()
}

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

func (n *NBTLongArray)String()(string){
	var s strings.Builder
	s.WriteString("Tag_LongArray(")
	if len(n.name) == 0 {
		s.WriteString("None")
	}else{
		s.WriteString(strconv.QuoteToASCII(n.name))
	}
	fmt.Fprintf(&s, "): %d ", len(n.Data))
	if len(n.Data) == 1 {
		s.WriteString("entry")
	}else{
		s.WriteString("entries")
	}
	s.WriteString("\n{\n")
	for _, v := range n.Data {
		fmt.Fprintf(&s, "  %dL\n", v)
	}
	s.WriteString("}")
	return s.String()
}

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
