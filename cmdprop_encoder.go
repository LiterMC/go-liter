
package liter

import (
	"fmt"
)

type EmptyCmdPropEncoder struct {
	id VarInt
	sid string
}

var _ CmdPropEncoder = (*EmptyCmdPropEncoder)(nil)

func NewEmptyCmdPropEncoder(id VarInt, sid string)(*EmptyCmdPropEncoder){
	return &EmptyCmdPropEncoder{
		id: id,
		sid: sid,
	}
}

func (p *EmptyCmdPropEncoder)Id()(VarInt){ return p.id }
func (p *EmptyCmdPropEncoder)String()(String){ return p.sid }
func (p *EmptyCmdPropEncoder)Encode(b *PacketBuilder, value any)(err error){ return }
func (p *EmptyCmdPropEncoder)Decode(r *PacketReader)(value any, err error){ return }


func init(){
	vanillaParsers := []CmdPropEncoder{
		NewEmptyCmdPropEncoder(0,  "brigadier:bool"),
		&numRangePropEncoder_Float{1,  "brigadier:float"},
		&numRangePropEncoder_Double{2,  "brigadier:double"},
		&numRangePropEncoder_Int{3,  "brigadier:integer"},
		&numRangePropEncoder_Long{4,  "brigadier:long"},
		NewEmptyCmdPropEncoder(5,  "brigadier:string"),
		NewEmptyCmdPropEncoder(6,  "minecraft:entity"),
		NewEmptyCmdPropEncoder(7,  "minecraft:game_profile"),
		NewEmptyCmdPropEncoder(8,  "minecraft:block_pos"),
		NewEmptyCmdPropEncoder(9,  "minecraft:column_pos"),
		NewEmptyCmdPropEncoder(10, "minecraft:vec3with"),
		NewEmptyCmdPropEncoder(11, "minecraft:vec2with"),
		NewEmptyCmdPropEncoder(12, "minecraft:block_state"),
		NewEmptyCmdPropEncoder(13, "minecraft:block_predicate"),
		NewEmptyCmdPropEncoder(14, "minecraft:item_stack"),
		NewEmptyCmdPropEncoder(15, "minecraft:item_predicate"),
		NewEmptyCmdPropEncoder(16, "minecraft:color"),
		NewEmptyCmdPropEncoder(17, "minecraft:component"),
		NewEmptyCmdPropEncoder(18, "minecraft:message"),
		NewEmptyCmdPropEncoder(19, "minecraft:nbt"),
		NewEmptyCmdPropEncoder(20, "minecraft:nbt_tag"),
		NewEmptyCmdPropEncoder(21, "minecraft:nbt_path"),
		NewEmptyCmdPropEncoder(22, "minecraft:objective"),
		NewEmptyCmdPropEncoder(23, "minecraft:objective_criteria"),
		NewEmptyCmdPropEncoder(24, "minecraft:operation"),
		NewEmptyCmdPropEncoder(25, "minecraft:particle"),
		NewEmptyCmdPropEncoder(26, "minecraft:angle"),
		NewEmptyCmdPropEncoder(27, "minecraft:rotationwith"),
		NewEmptyCmdPropEncoder(28, "minecraft:scoreboard_slot"),
		NewEmptyCmdPropEncoder(29, "minecraft:score_holder"),
		NewEmptyCmdPropEncoder(30, "minecraft:swizzle"),
		NewEmptyCmdPropEncoder(31, "minecraft:team"),
		NewEmptyCmdPropEncoder(32, "minecraft:item_slot"),
		NewEmptyCmdPropEncoder(33, "minecraft:resource_location"),
		NewEmptyCmdPropEncoder(34, "minecraft:function"),
		NewEmptyCmdPropEncoder(35, "minecraft:entity_anchor"),
		NewEmptyCmdPropEncoder(36, "minecraft:int_range"),
		NewEmptyCmdPropEncoder(37, "minecraft:float_range"),
		NewEmptyCmdPropEncoder(38, "minecraft:dimension"),
		NewEmptyCmdPropEncoder(39, "minecraft:gamemode"),
		NewEmptyCmdPropEncoder(40, "minecraft:time"),
		NewEmptyCmdPropEncoder(41, "minecraft:resource_or_tag"),
		NewEmptyCmdPropEncoder(42, "minecraft:resource_or_tag_key"),
		NewEmptyCmdPropEncoder(43, "minecraft:resource"),
		NewEmptyCmdPropEncoder(44, "minecraft:resource_key"),
		NewEmptyCmdPropEncoder(45, "minecraft:template_mirror"),
		NewEmptyCmdPropEncoder(46, "minecraft:template_rotation"),
		NewEmptyCmdPropEncoder(47, "minecraft:uuid"),
	}
	for _, p := range vanillaParsers {
		if !RegisterCmdEncoder(p) {
			panic(fmt.Errorf("Cannot register command parser %s(%d)", p.String(), p.Id()))
		}
	}
}

