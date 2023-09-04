
package nbt

type NBTCompound struct {
	name string
	Data map[string]NBT
}

func (n *NBTCompound)Type()(Byte){ return NbtCompound }
func (n *NBTCompound)Name()(string){ return n.name }
func (n *NBTCompound)SetName(name string){ n.name = name }

func (n *NBTCompound)Encode(b *PacketBuilder){
	for _, v := range n.Data {
		WriteNBT(b, v)
	}
}

func (n *NBTCompound)DecodeFrom(r *PacketReader)(err error){
	for {
		var v NBT
		if v, err = ReadNBT(r); err != nil {
			return
		}
		if _, done := v.(*NBTEnd); done {
			break
		}
		n.Data[v.Name()] = v
	}
	return
}
