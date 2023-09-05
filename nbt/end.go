
package nbt

type NBTEnd struct {}
var _NBTEndIns = new(NBTEnd)
func (n *NBTEnd)Type()(Byte){ return NbtEnd }
func (n *NBTEnd)SetName(name string){}
func (n *NBTEnd)Name()(string){ return "" }
func (n *NBTEnd)Encode(b *PacketBuilder){}
func (n *NBTEnd)DecodeFrom(r *PacketReader)(error){ return nil }
func (n *NBTEnd)String()(string){
	return "Tag_End(None)"
}
