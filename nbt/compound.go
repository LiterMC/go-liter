
package nbt

import (
	"fmt"
	"strings"
)

type NBTCompound struct {
	name string
	Data []NBT
}

func (n *NBTCompound)Type()(Byte){ return NbtCompound }
func (n *NBTCompound)Name()(string){ return n.name }
func (n *NBTCompound)SetName(name string){ n.name = name }

func (n *NBTCompound)String()(string){
	var s strings.Builder
	s.WriteString("Tag_Compound(")
	if len(n.name) == 0 {
		s.WriteString("None")
	}else{
		fmt.Fprintf(&s, "%q", n.name)
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

func (n *NBTCompound)Encode(b *PacketBuilder){
	for _, v := range n.Data {
		WriteNBT(b, v)
	}
}

func (n *NBTCompound)DecodeFrom(r *PacketReader)(err error){
	n.Data = n.Data[:0]
	for {
		var v NBT
		if v, err = ReadNBT(r); err != nil {
			return
		}
		if _, done := v.(*NBTEnd); done {
			break
		}
		n.Data = append(n.Data, v)
	}
	return
}

func (n *NBTCompound)Get(name string)(NBT){
	for _, v := range n.Data {
		if v.Name() == name {
			return v
		}
	}
	return nil
}

func (n *NBTCompound)Set(name string, tag NBT){
	for i, v := range n.Data {
		if v.Name() == name {
			n.Data[i] = tag
		}
	}
	n.Data = append(n.Data, tag)
}
