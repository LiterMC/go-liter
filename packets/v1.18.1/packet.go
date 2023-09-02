
// Generated at 2023-09-01 20:57:33.569 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=17341
// Protocol: 757
// Protocol Name: 1.18.1

package packet_1_18_1

import (
	"io"
	. "github.com/kmcsr/go-liter"
	internal "github.com/kmcsr/go-liter/packets/internal"
)

// ======== BEGIN login ========
// ---- login: serverbound ----

// ID=0x0
type LoginLoginStartPkt = internal.LoginLoginStart_758_2

// ID=0x1
type LoginEncryptionResponsePkt = internal.LoginEncryptionResponse_763_0

// ID=0x2
type LoginLoginPluginResponsePkt = internal.LoginLoginPluginResponse_763_0

// ---- login: clientbound ----

// ID=0x0
type LoginDisconnectPkt = internal.LoginDisconnect_763_0

// ID=0x1
type LoginEncryptionRequestPkt = internal.LoginEncryptionRequest_763_0

// ID=0x2
type LoginLoginSuccessPkt = internal.LoginLoginSuccess_758_0

// ID=0x3
type LoginSetCompressionPkt = internal.LoginSetCompression_763_0

// ID=0x4
type LoginLoginPluginRequestPkt = internal.LoginLoginPluginRequest_763_0

// ======== END login ========

// ======== BEGIN play ========
// ---- play: serverbound ----

// ID=0x0
type PlayTeleportConfirmPkt = internal.PlayTeleportConfirm_758_0

// ID=0x1
type PlayQueryBlockNBTPkt = internal.PlayQueryBlockNBT_758_0

// ID=0x2
type PlaySetDifficultyPkt = internal.PlaySetDifficulty_758_0

// ID=0x3
type PlayChatMessageServerPkt = internal.PlayChatMessage_758_3

// ID=0x4
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x5
type PlayClientSettingsPkt = internal.PlayClientSettings_758_0

// ID=0x6
type PlayTabCompleteServerPkt = internal.PlayTabComplete_758_0

// ID=0x7
type PlayClickWindowButtonPkt = internal.PlayClickWindowButton_758_0

// ID=0x8
type PlayClickWindowPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type    | Field Type    | Notes                                                                                                          |
	 * |-----------|-------|----------|---------------------|---------------------|---------------|---------------|----------------------------------------------------------------------------------------------------------------|
	 * | 0x08      | Play  | Server   | Window ID           | Window ID           | Unsigned Byte | Unsigned Byte | The ID of the window which was clicked. 0 for player inventory.                                                |
	 * | 0x08      | Play  | Server   | State ID            | State ID            | VarInt        | VarInt        | The last recieved State ID from either a Set Slot or a Window Items packet                                     |
	 * | 0x08      | Play  | Server   | Slot                | Slot                | Short         | Short         | The clicked slot number, see below.                                                                            |
	 * | 0x08      | Play  | Server   | Button              | Button              | Byte          | Byte          | The button used in the click, see below.                                                                       |
	 * | 0x08      | Play  | Server   | Mode                | Mode                | VarInt Enum   | VarInt Enum   | Inventory operation mode, see below.                                                                           |
	 * | 0x08      | Play  | Server   | Length of the array | Length of the array | VarInt        | VarInt        |                                                                                                                |
	 * | 0x08      | Play  | Server   | Array of slots      | Slot number         | Array         | Short         |                                                                                                                |
	 * | 0x08      | Play  | Server   | Array of slots      | Slot data           | Array         | Slot          | New data for this slot                                                                                         |
	 * | 0x08      | Play  | Server   | Clicked item        | Clicked item        | Slot          | Slot          | The clicked slot. Has to be empty (item ID = -1) for drop mode. Is always empty for mode 2 and mode 5 packets. |
	 * 
	 */
}

// ID=0x9
type PlayCloseWindowServerPkt = internal.PlayCloseWindow_758_0

// ID=0xa
type PlayPluginMessageServerPkt = internal.PlayPluginMessage_763_0

// ID=0xb
type PlayEditBookPkt struct {
	/* 0: Main hand, 1: Off hand. */
	Hand VarInt // VarInt enum
	/* Number of elements in the following array */
	Count VarInt // VarInt
	/* Text from each page. */
	Entries [][]String // Array of Strings
	/* If true, the next field is present. */
	HasTitle Bool // Boolean
	/* Title of book. */
	Title String // String
}

var _ Packet = (*PlayEditBookPkt)(nil)

func (p PlayEditBookPkt)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
	p.Count = len(p.Entries)
	b.VarInt(p.Count)
	for _, v := range p.Entries {
		TODO_Encode_Array(v)
	}
	b.Bool(p.HasTitle)
	b.String(p.Title)
}

func (p *PlayEditBookPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Entries = make([][]String, p.Count)
	for i, _ := range p.Entries {
		TODO_Decode_Array(Entries[i])
	}
	if p.HasTitle, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Title, ok = r.String(); !ok {
		return io.EOF
	}
}

// ID=0xc
type PlayQueryEntityNBTPkt = internal.PlayQueryEntityNBT_758_0

// ID=0xd
type PlayInteractEntityPkt = internal.PlayInteractEntity_758_0

// ID=0xe
type PlayGenerateStructurePkt = internal.PlayGenerateStructure_758_0

// ID=0xf
type PlayKeepAliveServerPkt = internal.PlayKeepAlive_763_0

// ID=0x10
type PlayLockDifficultyPkt = internal.PlayLockDifficulty_763_0

// ID=0x11
type PlayPlayerPositionPkt = internal.PlayPlayerPosition_758_0

// ID=0x12
type PlayPlayerPositionAndRotationPkt = internal.PlayPlayerPositionAndRotation_758_0

// ID=0x13
type PlayPlayerRotationPkt = internal.PlayPlayerRotation_758_0

// ID=0x14
type PlayPlayerMovementPkt = internal.PlayPlayerMovement_758_0

// ID=0x15
type PlayVehicleMoveServerPkt = internal.PlayVehicleMove_758_0

// ID=0x16
type PlaySteerBoatPkt = internal.PlaySteerBoat_758_0

// ID=0x17
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x18
type PlayCraftRecipeRequestPkt = internal.PlayCraftRecipeRequest_758_0

// ID=0x19
type PlayPlayerAbilitiesServerPkt = internal.PlayPlayerAbilities_760_3

// ID=0x1a
type PlayPlayerDiggingPkt = internal.PlayPlayerDigging_758_0

// ID=0x1b
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x1c
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x1d
type PlayPongPkt = internal.PlayPong_763_0

// ID=0x1e
type PlaySetRecipeBookStatePkt = internal.PlaySetRecipeBookState_758_0

// ID=0x1f
type PlaySetDisplayedRecipePkt = internal.PlaySetDisplayedRecipe_758_0

// ID=0x20
type PlayNameItemPkt = internal.PlayNameItem_758_0

// ID=0x21
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_758_0

// ID=0x22
type PlayAdvancementTabPkt = internal.PlayAdvancementTab_758_0

// ID=0x23
type PlaySelectTradePkt = internal.PlaySelectTrade_763_0

// ID=0x24
type PlaySetBeaconEffectPkt = internal.PlaySetBeaconEffect_758_2

// ID=0x25
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChange_758_1

// ID=0x26
type PlayUpdateCommandBlockPkt = internal.PlayUpdateCommandBlock_758_0

// ID=0x27
type PlayUpdateCommandBlockMinecartPkt = internal.PlayUpdateCommandBlockMinecart_758_0

// ID=0x28
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x29
type PlayUpdateJigsawBlockPkt = internal.PlayUpdateJigsawBlock_758_0

// ID=0x2a
type PlayUpdateStructureBlockPkt = internal.PlayUpdateStructureBlock_758_0

// ID=0x2b
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x2c
type PlayAnimationPkt = internal.PlayAnimation_758_0

// ID=0x2d
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x2e
type PlayPlayerBlockPlacementPkt = internal.PlayPlayerBlockPlacement_758_0

// ID=0x2f
type PlayUseItemPkt = internal.PlayUseItem_758_1

// ---- play: clientbound ----

// ID=0x0
type PlaySpawnEntityPkt = internal.PlaySpawnEntity_758_1

// ID=0x1
type PlaySpawnExperienceOrbPkt = internal.PlaySpawnExperienceOrb_763_0

// ID=0x2
type PlaySpawnLivingEntityPkt = internal.PlaySpawnLivingEntity_757_1

// ID=0x3
type PlaySpawnPaintingPkt = internal.PlaySpawnPainting_758_0

// ID=0x4
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_763_0

// ID=0x5
type PlaySculkVibrationSignalPkt = internal.PlaySculkVibrationSignal_758_0

// ID=0x6
type PlayEntityAnimationPkt = internal.PlayEntityAnimation_763_0

// ID=0x7
type PlayStatisticsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name   | Field Type | Field Type | Notes                                      |
	 * |-----------|-------|----------|------------|--------------|------------|------------|--------------------------------------------|
	 * | 0x07      | Play  | Client   | Count      | Count        | VarInt     | VarInt     | Number of elements in the following array. |
	 * | 0x07      | Play  | Client   | Statistic  | Category ID  | Array      | VarInt     | See below.                                 |
	 * | 0x07      | Play  | Client   | Statistic  | Statistic ID | Array      | VarInt     | See below.                                 |
	 * | 0x07      | Play  | Client   | Statistic  | Value        | Array      | VarInt     | The amount to set it to.                   |
	 * 
	 */
}

// ID=0x8
type PlayAcknowledgePlayerDiggingPkt = internal.PlayAcknowledgePlayerDigging_758_0

// ID=0x9
type PlayBlockBreakAnimationPkt = internal.PlayBlockBreakAnimation_758_0

// ID=0xa
type PlayBlockEntityDataPkt = internal.PlayBlockEntityData_763_0

// ID=0xb
type PlayBlockActionPkt = internal.PlayBlockAction_758_1

// ID=0xc
type PlayBlockChangePkt = internal.PlayBlockChange_758_0

// ID=0xd
type PlayBossBarPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name       | Field Name | Field Type    | Notes                                                                                                                                     |
	 * |-----------|-------|----------|------------------|------------|---------------|-------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0D      | Play  | Client   | UUID             | UUID       | UUID          | Unique ID for this bar.                                                                                                                   |
	 * | 0x0D      | Play  | Client   | Action           | Action     | VarInt Enum   | Determines the layout of the remaining packet.                                                                                            |
	 * | 0x0D      | Play  | Client   | Action           | Field Name |               |                                                                                                                                           |
	 * | 0x0D      | Play  | Client   | 0: add           | Title      | Chat          |                                                                                                                                           |
	 * | 0x0D      | Play  | Client   | 0: add           | Health     | Float         | From 0 to 1. Values greater than 1 do not crash a Notchian client, and start rendering part of a second health bar at around 1.5.         |
	 * | 0x0D      | Play  | Client   | 0: add           | Color      | VarInt Enum   | Color ID (see below).                                                                                                                     |
	 * | 0x0D      | Play  | Client   | 0: add           | Division   | VarInt Enum   | Type of division (see below).                                                                                                             |
	 * | 0x0D      | Play  | Client   | 0: add           | Flags      | Unsigned Byte | Bit mask. 0x1: should darken sky, 0x2: is dragon bar (used to play end music), 0x04: create fog (previously was also controlled by 0x02). |
	 * | 0x0D      | Play  | Client   | 1: remove        | no fields  | no fields     | Removes this boss bar.                                                                                                                    |
	 * | 0x0D      | Play  | Client   | 2: update health | Health     | Float         | as above                                                                                                                                  |
	 * | 0x0D      | Play  | Client   | 3: update title  | Title      | Chat          |                                                                                                                                           |
	 * | 0x0D      | Play  | Client   | 4: update style  | Color      | VarInt Enum   | Color ID (see below).                                                                                                                     |
	 * | 0x0D      | Play  | Client   | 4: update style  | Dividers   | VarInt Enum   | as above                                                                                                                                  |
	 * | 0x0D      | Play  | Client   | 5: update flags  | Flags      | Unsigned Byte | as above                                                                                                                                  |
	 * 
	 */
}

// ID=0xe
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_758_0

// ID=0xf
type PlayChatMessageClientPkt = internal.PlayChatMessage_758_2

// ID=0x10
type PlayClearTitlesPkt = internal.PlayClearTitles_763_0

// ID=0x11
type PlayTabCompleteClientPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name  | Field Type | Field Type     | Notes                                                                                                                                                                                                  |
	 * |-----------|-------|----------|------------|-------------|------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x11      | Play  | Client   |            |             |            |                |                                                                                                                                                                                                        |
	 * | 0x11      | Play  | Client   | ID         | ID          | VarInt     | VarInt         | Transaction ID.                                                                                                                                                                                        |
	 * | 0x11      | Play  | Client   | Start      | Start       | VarInt     | VarInt         | Start of the text to replace.                                                                                                                                                                          |
	 * | 0x11      | Play  | Client   | Length     | Length      | VarInt     | VarInt         | Length of the text to replace.                                                                                                                                                                         |
	 * | 0x11      | Play  | Client   | Count      | Count       | VarInt     | VarInt         | Number of elements in the following array.                                                                                                                                                             |
	 * | 0x11      | Play  | Client   | Matches    | Match       | Array      | String (32767) | One eligible value to insert, note that each command is sent separately instead of in a single string, hence the need for Count.  Note that for instance this doesn't include a leading / on commands. |
	 * | 0x11      | Play  | Client   | Matches    | Has tooltip | Array      | Boolean        | True if the following is present.                                                                                                                                                                      |
	 * | 0x11      | Play  | Client   | Matches    | Tooltip     | Array      | Optional Chat  | Tooltip to display; only present if previous boolean is true.                                                                                                                                          |
	 * 
	 */
}

// ID=0x12
type PlayDeclareCommandsPkt = internal.PlayDeclareCommands_758_0

// ID=0x13
type PlayCloseWindowClientPkt = internal.PlayCloseWindow_758_0

// ID=0x14
type PlayWindowItemsPkt = internal.PlayWindowItems_758_0

// ID=0x15
type PlayWindowPropertyPkt = internal.PlayWindowProperty_758_0

// ID=0x16
type PlaySetSlotPkt = internal.PlaySetSlot_758_0

// ID=0x17
type PlaySetCooldownPkt = internal.PlaySetCooldown_763_0

// ID=0x18
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_763_0

// ID=0x19
type PlayNamedSoundEffectPkt = internal.PlayNamedSoundEffect_758_0

// ID=0x1a
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x1b
type PlayEntityStatusPkt = internal.PlayEntityStatus_758_0

// ID=0x1c
type PlayExplosionPkt = internal.PlayExplosion_760_1

// ID=0x1d
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x1e
type PlayChangeGameStatePkt = internal.PlayChangeGameState_758_0

// ID=0x1f
type PlayOpenHorseWindowPkt = internal.PlayOpenHorseWindow_758_0

// ID=0x20
type PlayInitializeWorldBorderPkt = internal.PlayInitializeWorldBorder_763_0

// ID=0x21
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x22
type PlayChunkDataAndUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name               | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
	 * |-----------|-------|----------|--------------------------|--------------------------|------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x22      | Play  | Client   | Chunk X                  | Chunk X                  | Int        | Int                 | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x22      | Play  | Client   | Chunk Z                  | Chunk Z                  | Int        | Int                 | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x22      | Play  | Client   | Heightmaps               | Heightmaps               | NBT        | NBT                 | Compound containing one long array named MOTION_BLOCKING, which is a heightmap for the highest solid block at each position in the chunk (as a compacted long array with 256 entries, with the number of bits per entry varying depending on the world's height, defined by the formula ceil(log2(height + 1))). The Notchian server also adds a WORLD_SURFACE long array, the purpose of which is unknown, but it's not required for the chunk to be accepted. |
	 * | 0x22      | Play  | Client   | Size                     | Size                     | VarInt     | VarInt              | Size of Data in bytes                                                                                                                                                                                                                                                                                                                                                                                                                                           |
	 * | 0x22      | Play  | Client   | Data                     | Data                     | Byte array | Byte array          | See data structure in Chunk Format                                                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x22      | Play  | Client   | Number of block entities | Number of block entities | VarInt     | VarInt              | Number of elements in the following array                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x22      | Play  | Client   | Block Entity             | Packed XZ                | Array      | Byte                | The packed section coordinates, calculated from ((blockX & 15) << 4) | (blockZ & 15)                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x22      | Play  | Client   | Block Entity             | Y                        | Array      | Short               | The height relative to the world                                                                                                                                                                                                                                                                                                                                                                                                                                |
	 * | 0x22      | Play  | Client   | Block Entity             | Type                     | Array      | VarInt              | The type of block entity                                                                                                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x22      | Play  | Client   | Block Entity             | Data                     | Array      | NBT                 | The block entity's data, without the X, Y, and Z values                                                                                                                                                                                                                                                                                                                                                                                                         |
	 * | 0x22      | Play  | Client   | Trust Edges              | Trust Edges              | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x22      | Play  | Client   | Sky Light Mask           | Sky Light Mask           | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Sky Light array below.  The least significant bit is for blocks 16 blocks to 1 block below the min world height (one section below the world), while the most significant bit covers blocks 1 to 16 blocks above the max world height (one section above the world).                                          |
	 * | 0x22      | Play  | Client   | Block Light Mask         | Block Light Mask         | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Block Light array below.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                 |
	 * | 0x22      | Play  | Client   | Empty Sky Light Mask     | Empty Sky Light Mask     | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Sky Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                    |
	 * | 0x22      | Play  | Client   | Empty Block Light Mask   | Empty Block Light Mask   | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Block Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                  |
	 * | 0x22      | Play  | Client   | Sky Light array count    | Sky Light array count    | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Sky Light Mask                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x22      | Play  | Client   | Sky Light arrays         | Length                   | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x22      | Play  | Client   | Sky Light arrays         | Sky Light array          | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value. Indexed ((y<<8) | (z<<4) | x) / 2  If there's a remainder, masked 0xF0 else 0x0F.                                                                                                                                                                                                                                                |
	 * | 0x22      | Play  | Client   | Block Light array count  | Block Light array count  | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Block Light Mask                                                                                                                                                                                                                                                                                                                                                               |
	 * | 0x22      | Play  | Client   | Block Light arrays       | Length                   | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x22      | Play  | Client   | Block Light arrays       | Block Light array        | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value. Indexed ((y<<8) | (z<<4) | x) / 2  If there's a remainder, masked 0xF0 else 0x0F.                                                                                                                                                                                                                                              |
	 * 
	 */
}

// ID=0x23
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x24
type PlayParticlePkt = internal.PlayParticle_758_1

// ID=0x25
type PlayUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name              | Field Name              | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                  |
	 * |-----------|-------|----------|-------------------------|-------------------------|------------|---------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x25      | Play  | Client   | Chunk X                 | Chunk X                 | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x25      | Play  | Client   | Chunk Z                 | Chunk Z                 | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x25      | Play  | Client   | Trust Edges             | Trust Edges             | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                                                          |
	 * | 0x25      | Play  | Client   | Sky Light Mask          | Sky Light Mask          | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Sky Light array below.  The least significant bit is for blocks 16 blocks to 1 block below the min world height (one section below the world), while the most significant bit covers blocks 1 to 16 blocks above the max world height (one section above the world). |
	 * | 0x25      | Play  | Client   | Block Light Mask        | Block Light Mask        | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Block Light array below.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                        |
	 * | 0x25      | Play  | Client   | Empty Sky Light Mask    | Empty Sky Light Mask    | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Sky Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                           |
	 * | 0x25      | Play  | Client   | Empty Block Light Mask  | Empty Block Light Mask  | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Block Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                         |
	 * | 0x25      | Play  | Client   | Sky Light array count   | Sky Light array count   | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Sky Light Mask                                                                                                                                                                                                                                                                                                                        |
	 * | 0x25      | Play  | Client   | Sky Light arrays        | Length                  | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x25      | Play  | Client   | Sky Light arrays        | Sky Light array         | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                                                         |
	 * | 0x25      | Play  | Client   | Block Light array count | Block Light array count | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Block Light Mask                                                                                                                                                                                                                                                                                                                      |
	 * | 0x25      | Play  | Client   | Block Light arrays      | Length                  | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x25      | Play  | Client   | Block Light arrays      | Block Light array       | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                                                       |
	 * 
	 */
}

// ID=0x26
type PlayJoinGamePkt = internal.PlayJoinGame_758_0

// ID=0x27
type PlayMapDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type                      | Field Type                      | Notes                                                                                                      |
	 * |-----------|-------|----------|-------------------|-------------------|---------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x27      | Play  | Client   | Map ID            | Map ID            | VarInt                          | VarInt                          | Map ID of the map being modified                                                                           |
	 * | 0x27      | Play  | Client   | Scale             | Scale             | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x27      | Play  | Client   | Locked            | Locked            | Boolean                         | Boolean                         | True if the map has been locked in a cartography table                                                     |
	 * | 0x27      | Play  | Client   | Tracking Position | Tracking Position | Boolean                         | Boolean                         | Specifies whether player and item frame icons are shown                                                    |
	 * | 0x27      | Play  | Client   | Icon Count        | Icon Count        | VarInt                          | VarInt                          | Number of elements in the following array. Only present if "Tracking Position" is true.                    |
	 * | 0x27      | Play  | Client   | Icon              | Type              | Array                           | VarInt enum                     | See below                                                                                                  |
	 * | 0x27      | Play  | Client   | Icon              | X                 | Array                           | Byte                            | Map coordinates: -128 for furthest left, +127 for furthest right                                           |
	 * | 0x27      | Play  | Client   | Icon              | Z                 | Array                           | Byte                            | Map coordinates: -128 for highest, +127 for lowest                                                         |
	 * | 0x27      | Play  | Client   | Icon              | Direction         | Array                           | Byte                            | 0-15                                                                                                       |
	 * | 0x27      | Play  | Client   | Icon              | Has Display Name  | Array                           | Boolean                         |                                                                                                            |
	 * | 0x27      | Play  | Client   | Icon              | Display Name      | Array                           | Optional Chat                   | Only present if previous Boolean is true                                                                   |
	 * | 0x27      | Play  | Client   | Columns           | Columns           | Unsigned Byte                   | Unsigned Byte                   | Number of columns updated                                                                                  |
	 * | 0x27      | Play  | Client   | Rows              | Rows              | Optional Unsigned Byte          | Optional Unsigned Byte          | Only if Columns is more than 0; number of rows updated                                                     |
	 * | 0x27      | Play  | Client   | X                 | X                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column                                         |
	 * | 0x27      | Play  | Client   | Z                 | Z                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row                                           |
	 * | 0x27      | Play  | Client   | Length            | Length            | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array                                              |
	 * | 0x27      | Play  | Client   | Data              | Data              | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format                                                        |
	 * 
	 */
}

// ID=0x28
type PlayTradeListPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name                   | Field Type | Field Type    | Notes                                                                                                                                                                             |
	 * |-----------|-------|----------|---------------------|------------------------------|------------|---------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x28      | Play  | Client   |                     |                              |            |               |                                                                                                                                                                                   |
	 * | 0x28      | Play  | Client   | Window ID           | Window ID                    | VarInt     | VarInt        | The ID of the window that is open; this is an int rather than a byte.                                                                                                             |
	 * | 0x28      | Play  | Client   | Size                | Size                         | Byte       | Byte          | The number of trades in the following array.                                                                                                                                      |
	 * | 0x28      | Play  | Client   | Trades              | Input item 1                 | Array      | Slot          | The first item the player has to supply for this villager trade. The count of the item stack is the default "price" of this trade.                                                |
	 * | 0x28      | Play  | Client   | Trades              | Output item                  | Array      | Slot          | The item the player will receive from this villager trade.                                                                                                                        |
	 * | 0x28      | Play  | Client   | Trades              | Has second item              | Array      | Boolean       | Whether there is a second item.                                                                                                                                                   |
	 * | 0x28      | Play  | Client   | Trades              | Input item 2                 | Array      | Optional Slot | The second item the player has to supply for this villager trade; only present if has a second item is true.                                                                      |
	 * | 0x28      | Play  | Client   | Trades              | Trade disabled               | Array      | Boolean       | True if the trade is disabled; false if the trade is enabled.                                                                                                                     |
	 * | 0x28      | Play  | Client   | Trades              | Number of trade uses         | Array      | Integer       | Number of times the trade has been used so far. If equal to the maximum number of trades, the client will display a red X.                                                        |
	 * | 0x28      | Play  | Client   | Trades              | Maximum number of trade uses | Array      | Integer       | Number of times this trade can be used before it's exhausted.                                                                                                                     |
	 * | 0x28      | Play  | Client   | Trades              | XP                           | Array      | Integer       | Amount of XP the villager will earn each time the trade is used.                                                                                                                  |
	 * | 0x28      | Play  | Client   | Trades              | Special Price                | Array      | Integer       | Can be zero or negative. The number is added to the price when an item is discounted due to player reputation or other effects.                                                   |
	 * | 0x28      | Play  | Client   | Trades              | Price Multiplier             | Array      | Float         | Can be low (0.05) or high (0.2). Determines how much demand, player reputation, and temporary effects will adjust the price.                                                      |
	 * | 0x28      | Play  | Client   | Trades              | Demand                       | Array      | Integer       | If positive, causes the price to increase. Negative values seem to be treated the same as zero.                                                                                   |
	 * | 0x28      | Play  | Client   | Villager level      | Villager level               | VarInt     | VarInt        | Appears on the trade GUI; meaning comes from the translation key merchant.level. + level.
	 * 1: Novice, 2: Apprentice, 3: Journeyman, 4: Expert, 5: Master.                          |
	 * | 0x28      | Play  | Client   | Experience          | Experience                   | VarInt     | VarInt        | Total experience for this villager (always 0 for the wandering trader).                                                                                                           |
	 * | 0x28      | Play  | Client   | Is regular villager | Is regular villager          | Boolean    | Boolean       | True if this is a regular villager; false for the wandering trader.  When false, hides the villager level and some other GUI elements.                                            |
	 * | 0x28      | Play  | Client   | Can restock         | Can restock                  | Boolean    | Boolean       | True for regular villagers and false for the wandering trader. If true, the "Villagers restock up to two times per day." message is displayed when hovering over disabled trades. |
	 * 
	 */
}

// ID=0x29
type PlayEntityPositionPkt = internal.PlayEntityPosition_758_0

// ID=0x2a
type PlayEntityPositionAndRotationPkt = internal.PlayEntityPositionAndRotation_758_0

// ID=0x2b
type PlayEntityRotationPkt = internal.PlayEntityRotation_758_0

// ID=0x2c
type PlayVehicleMoveClientPkt = internal.PlayVehicleMove_758_0

// ID=0x2d
type PlayOpenBookPkt = internal.PlayOpenBook_763_0

// ID=0x2e
type PlayOpenWindowPkt = internal.PlayOpenWindow_758_0

// ID=0x2f
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x30
type PlayPingPkt = internal.PlayPing_763_0

// ID=0x31
type PlayCraftRecipeResponsePkt = internal.PlayCraftRecipeResponse_758_0

// ID=0x32
type PlayPlayerAbilitiesClientPkt = internal.PlayPlayerAbilities_760_2

// ID=0x33
type PlayEndCombatEventPkt = internal.PlayEndCombatEvent_758_0

// ID=0x34
type PlayEnterCombatEventPkt = internal.PlayEnterCombatEvent_758_0

// ID=0x35
type PlayDeathCombatEventPkt = internal.PlayDeathCombatEvent_758_0

// ID=0x36
type PlayPlayerInfoPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type              | Notes                                                    |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-------------------------|----------------------------------------------------------|
	 * | 0x36      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt                  | Determines the rest of the Player format after the UUID. |
	 * | 0x36      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt                  | Number of elements in the following array.               |
	 * | 0x36      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID                    |                                                          |
	 * | 0x36      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                         |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String (16)   | String (16)             |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt                  | Number of elements in the following array.               |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String (32767)          |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String (32767)          |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean                 |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String (32767) | Only if Is Signed is true.                               |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds.                                |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only if Has Display Name is true.                        |
	 * | 0x36      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds.                                |
	 * | 0x36      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                          |
	 * | 0x36      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true.                   |
	 * | 0x36      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields               |                                                          |
	 * 
	 */
}

// ID=0x37
type PlayFacePlayerPkt = internal.PlayFacePlayer_757_0

// ID=0x38
type PlayPlayerPositionAndLookPkt = internal.PlayPlayerPositionAndLook_758_0

// ID=0x39
type PlayUnlockRecipesPkt = internal.PlayUnlockRecipes_758_0

// ID=0x3a
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x3b
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x3c
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_758_0

// ID=0x3d
type PlayRespawnPkt = internal.PlayRespawn_758_3

// ID=0x3e
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x3f
type PlayMultiBlockChangePkt = internal.PlayMultiBlockChange_758_0

// ID=0x40
type PlaySelectAdvancementTabPkt = internal.PlaySelectAdvancementTab_758_0

// ID=0x41
type PlayActionBarPkt = internal.PlayActionBar_758_0

// ID=0x42
type PlayWorldBorderCenterPkt = internal.PlayWorldBorderCenter_758_0

// ID=0x43
type PlayWorldBorderLerpSizePkt = internal.PlayWorldBorderLerpSize_758_0

// ID=0x44
type PlayWorldBorderSizePkt = internal.PlayWorldBorderSize_758_0

// ID=0x45
type PlayWorldBorderWarningDelayPkt = internal.PlayWorldBorderWarningDelay_758_0

// ID=0x46
type PlayWorldBorderWarningReachPkt = internal.PlayWorldBorderWarningReach_758_0

// ID=0x47
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x48
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_758_0

// ID=0x49
type PlayUpdateViewPositionPkt = internal.PlayUpdateViewPosition_758_0

// ID=0x4a
type PlayUpdateViewDistancePkt = internal.PlayUpdateViewDistance_758_0

// ID=0x4b
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_758_0

// ID=0x4c
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x4d
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x4e
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x4f
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x50
type PlayEntityEquipmentPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name | Field Type | Field Type | Notes                                                                                                                                                                                                                          |
	 * |-----------|-------|----------|------------|------------|------------|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x50      | Play  | Client   | Entity ID  | Entity ID  | VarInt     | VarInt     | Entity's EID.                                                                                                                                                                                                                  |
	 * | 0x50      | Play  | Client   | Equipment  | Slot       | Array      | Byte Enum  | Equipment slot. 0: main hand, 1: off hand, 2–5: armor slot (2: boots, 3: leggings, 4: chestplate, 5: helmet).  Also has the top bit set if another entry follows, and otherwise unset if this is the last item in the array. |
	 * | 0x50      | Play  | Client   | Equipment  | Item       | Array      | Slot       |                                                                                                                                                                                                                                |
	 * 
	 */
}

// ID=0x51
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x52
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x53
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_758_0

// ID=0x54
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x55
type PlayTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                   | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|------------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x55      | Play  | Client   | Team Name                    | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x55      | Play  | Client   | Mode                         | Mode                | Byte                 | Determines the layout of the remaining packet.                                                                         |
	 * | 0x55      | Play  | Client   | 0: create team               | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x55      | Play  | Client   | 0: create team               | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team.                                     |
	 * | 0x55      | Play  | Client   | 0: create team               | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never.                                                                      |
	 * | 0x55      | Play  | Client   | 0: create team               | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never.                                                                            |
	 * | 0x55      | Play  | Client   | 0: create team               | Team Color          | VarInt enum          | Used to color the name of players on the team; see below.                                                              |
	 * | 0x55      | Play  | Client   | 0: create team               | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team.                                                      |
	 * | 0x55      | Play  | Client   | 0: create team               | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team.                                                       |
	 * | 0x55      | Play  | Client   | 0: create team               | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x55      | Play  | Client   | 0: create team               | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x55      | Play  | Client   | 1: remove team               | no fields           | no fields            |                                                                                                                        |
	 * | 0x55      | Play  | Client   | 2: update team info          | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x55      | Play  | Client   | 2: update team info          | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team.                                    |
	 * | 0x55      | Play  | Client   | 2: update team info          | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x55      | Play  | Client   | 2: update team info          | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x55      | Play  | Client   | 2: update team info          | Team Color          | VarInt enum          | Used to color the name of players on the team; see below.                                                              |
	 * | 0x55      | Play  | Client   | 2: update team info          | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team.                                                      |
	 * | 0x55      | Play  | Client   | 2: update team info          | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team.                                                       |
	 * | 0x55      | Play  | Client   | 3: add entities to team      | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x55      | Play  | Client   | 3: add entities to team      | Entities            | Array of String (40) | Identifiers for the added entities.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x55      | Play  | Client   | 4: remove entities from team | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x55      | Play  | Client   | 4: remove entities from team | Entities            | Array of String (40) | Identifiers for the removed entities.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

// ID=0x56
type PlayUpdateScorePkt = internal.PlayUpdateScore_757_1

// ID=0x57
type PlayUpdateSimulationDistancePkt = internal.PlayUpdateSimulationDistance_758_0

// ID=0x58
type PlaySetTitleSubTitlePkt = internal.PlaySetTitleSubTitle_758_0

// ID=0x59
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x5a
type PlaySetTitleTextPkt = internal.PlaySetTitleText_763_0

// ID=0x5b
type PlaySetTitleTimesPkt = internal.PlaySetTitleTimes_758_0

// ID=0x5c
type PlayEntitySoundEffectPkt = internal.PlayEntitySoundEffect_760_1

// ID=0x5d
type PlaySoundEffectPkt = internal.PlaySoundEffect_758_2

// ID=0x5e
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x5f
type PlayPlayerListHeaderAndFooterPkt = internal.PlayPlayerListHeaderAndFooter_758_0

// ID=0x60
type PlayNBTQueryResponsePkt = internal.PlayNBTQueryResponse_758_0

// ID=0x61
type PlayCollectItemPkt = internal.PlayCollectItem_758_0

// ID=0x62
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x63
type PlayAdvancementsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name    | Field Type          | Field Type           | Notes                                                       |
	 * |-----------|-------|----------|---------------------|---------------|---------------------|----------------------|-------------------------------------------------------------|
	 * | 0x63      | Play  | Client   | Reset/Clear         | Reset/Clear   | Boolean             | Boolean              | Whether to reset/clear the current advancements.            |
	 * | 0x63      | Play  | Client   | Mapping size        | Mapping size  | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x63      | Play  | Client   | Advancement mapping | Key           | Array               | Identifier           | The identifier of the advancement.                          |
	 * | 0x63      | Play  | Client   | Advancement mapping | Value         | Array               | Advancement          | See below                                                   |
	 * | 0x63      | Play  | Client   | List size           | List size     | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x63      | Play  | Client   | Identifiers         | Identifiers   | Array of Identifier | Array of Identifier  | The identifiers of the advancements that should be removed. |
	 * | 0x63      | Play  | Client   | Progress size       | Progress size | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x63      | Play  | Client   | Progress mapping    | Key           | Array               | Identifier           | The identifier of the advancement.                          |
	 * | 0x63      | Play  | Client   | Progress mapping    | Value         | Array               | Advancement progress | See below.                                                  |
	 * 
	 */
}

// ID=0x64
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x64      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x64      | Play  | Client   | Number Of Properties | Number Of Properties | VarInt     | VarInt                 | Number of elements in the following array.            |
	 * | 0x64      | Play  | Client   | Property             | Key                  | Array      | Identifier             | See below.                                            |
	 * | 0x64      | Play  | Client   | Property             | Value                | Array      | Double                 | See below.                                            |
	 * | 0x64      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array.            |
	 * | 0x64      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x65
type PlayEntityEffectPkt = internal.PlayEntityEffect_757_2

// ID=0x66
type PlayDeclareRecipesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name  | Field Type | Field Type       | Notes                                                                   |
	 * |-----------|-------|----------|-------------|-------------|------------|------------------|-------------------------------------------------------------------------|
	 * | 0x66      | Play  | Client   | Num Recipes | Num Recipes | VarInt     | VarInt           | Number of elements in the following array.                              |
	 * | 0x66      | Play  | Client   | Recipe      | Type        | Array      | Identifier       | The recipe type, see below.                                             |
	 * | 0x66      | Play  | Client   | Recipe      | Recipe ID   | Array      | Identifier       |                                                                         |
	 * | 0x66      | Play  | Client   | Recipe      | Data        | Array      | Optional, varies | Additional data for the recipe.  For some types, there will be no data. |
	 * 
	 */
}

// ID=0x67
type PlayTagsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type | Field Type  | Notes                                                                                                                                        |
	 * |-----------|-------|----------|---------------------|---------------------|------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x67      | Play  | Client   | Length of the array | Length of the array | VarInt     | VarInt      |                                                                                                                                              |
	 * | 0x67      | Play  | Client   | Array of tags       | Tag type            | Array      | Identifier  | Tag identifier (Vanilla required tags are minecraft:block, minecraft:item, minecraft:fluid, minecraft:entity_type, and minecraft:game_event) |
	 * | 0x67      | Play  | Client   | Array of tags       | Array of Tag        | Array      | (See below) |                                                                                                                                              |
	 * 
	 */
}

// ======== END play ========
