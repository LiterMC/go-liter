
package liter

type EmptyCmdPropParser struct {
	id VarInt
	sid string
}

var _ CmdPropParser = (*EmptyCmdPropParser)(nil)

func NewEmptyCmdPropParser(id VarInt, sid string)(*EmptyCmdPropParser){
	return &EmptyCmdPropParser{
		id: id,
		sid: sid,
	}
}

func (p *EmptyCmdPropParser)Id()(VarInt){ return p.id }
func (p *EmptyCmdPropParser)String()(String){ return p.sid }
func (p *EmptyCmdPropParser)Encode(b *PacketBuilder, value any)(err error){ return }
func (p *EmptyCmdPropParser)Decode(r *PacketReader)(value any, err error){ return }

func init(){
	vanillaParsers := []CmdPropParser{
		NewEmptyCmdPropParser(0,  "brigadier:bool"),
		NewEmptyCmdPropParser(1,  "brigadier:float"),
		NewEmptyCmdPropParser(2,  "brigadier:double"),
		NewEmptyCmdPropParser(3,  "brigadier:integer"),
		NewEmptyCmdPropParser(4,  "brigadier:long"),
		NewEmptyCmdPropParser(5,  "brigadier:string"),
		NewEmptyCmdPropParser(6,  "minecraft:entity"),
		NewEmptyCmdPropParser(7,  "minecraft:game_profile"),
		NewEmptyCmdPropParser(8,  "minecraft:block_pos"),
		NewEmptyCmdPropParser(9,  "minecraft:column_pos"),
		NewEmptyCmdPropParser(10, "minecraft:vec3with"),
		NewEmptyCmdPropParser(11, "minecraft:vec2with"),
		NewEmptyCmdPropParser(12, "minecraft:block_state"),
		NewEmptyCmdPropParser(13, "minecraft:block_predicate"),
		NewEmptyCmdPropParser(14, "minecraft:item_stack"),
		NewEmptyCmdPropParser(15, "minecraft:item_predicate"),
		NewEmptyCmdPropParser(16, "minecraft:color"),
		NewEmptyCmdPropParser(17, "minecraft:component"),
		NewEmptyCmdPropParser(18, "minecraft:message"),
		NewEmptyCmdPropParser(19, "minecraft:nbt"),
		NewEmptyCmdPropParser(20, "minecraft:nbt_tag"),
		NewEmptyCmdPropParser(21, "minecraft:nbt_path"),
		NewEmptyCmdPropParser(22, "minecraft:objective"),
		NewEmptyCmdPropParser(23, "minecraft:objective_criteria"),
		NewEmptyCmdPropParser(24, "minecraft:operation"),
		NewEmptyCmdPropParser(25, "minecraft:particle"),
		NewEmptyCmdPropParser(26, "minecraft:angle"),
		NewEmptyCmdPropParser(27, "minecraft:rotationwith"),
		NewEmptyCmdPropParser(28, "minecraft:scoreboard_slot"),
		NewEmptyCmdPropParser(29, "minecraft:score_holder"),
		NewEmptyCmdPropParser(30, "minecraft:swizzle"),
		NewEmptyCmdPropParser(31, "minecraft:team"),
		NewEmptyCmdPropParser(32, "minecraft:item_slot"),
		NewEmptyCmdPropParser(33, "minecraft:resource_location"),
		NewEmptyCmdPropParser(34, "minecraft:function"),
		NewEmptyCmdPropParser(35, "minecraft:entity_anchor"),
		NewEmptyCmdPropParser(36, "minecraft:int_range"),
		NewEmptyCmdPropParser(37, "minecraft:float_range"),
		NewEmptyCmdPropParser(38, "minecraft:dimension"),
		NewEmptyCmdPropParser(39, "minecraft:gamemode"),
		NewEmptyCmdPropParser(40, "minecraft:time"),
		NewEmptyCmdPropParser(41, "minecraft:resource_or_tag"),
		NewEmptyCmdPropParser(42, "minecraft:resource_or_tag_key"),
		NewEmptyCmdPropParser(43, "minecraft:resource"),
		NewEmptyCmdPropParser(44, "minecraft:resource_key"),
		NewEmptyCmdPropParser(45, "minecraft:template_mirror"),
		NewEmptyCmdPropParser(46, "minecraft:template_rotation"),
		NewEmptyCmdPropParser(47, "minecraft:uuid"),
	}
	for _, p := range vanillaParsers {
		RegisterCmdParser(p)
	}
}

