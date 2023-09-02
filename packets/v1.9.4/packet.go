
// Generated at 2023-09-01 20:45:22.208 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=7959
// Protocol: 110

package packet_1_9_4

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
type LoginEncryptionResponsePkt = internal.LoginEncryptionResponse_758_2

// ---- login: clientbound ----

// ID=0x0
type LoginDisconnectPkt = internal.LoginDisconnect_763_0

// ID=0x1
type LoginEncryptionRequestPkt = internal.LoginEncryptionRequest_763_0

// ID=0x2
type LoginLoginSuccessPkt = internal.LoginLoginSuccess_578_1

// ID=0x3
type LoginSetCompressionPkt = internal.LoginSetCompression_763_0

// ======== END login ========

// ======== BEGIN play ========
// ---- play: serverbound ----

// ID=0x0
type PlayTeleportConfirmPkt = internal.PlayTeleportConfirm_758_0

// ID=0x1
type PlayTabCompleteServerPkt = internal.PlayTabComplete_338_6

// ID=0x2
type PlayChatMessageServerPkt = internal.PlayChatMessage_404_13

// ID=0x3
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x4
type PlayClientSettingsPkt = internal.PlayClientSettings_754_2

// ID=0x5
type PlayConfirmTransactionServerPkt = internal.PlayConfirmTransaction_404_0

// ID=0x6
type PlayEnchantItemPkt = internal.PlayEnchantItem_404_0

// ID=0x7
type PlayClickWindowPkt = internal.PlayClickWindow_754_0

// ID=0x8
type PlayCloseWindowServerPkt = internal.PlayCloseWindow_758_0

// ID=0x9
type PlayPluginMessageServerPkt = internal.PlayPluginMessage_340_1

// ID=0xa
type PlayUseEntityPkt = internal.PlayUseEntity_404_0

// ID=0xb
type PlayKeepAliveServerPkt = internal.PlayKeepAlive_338_1

// ID=0xc
type PlayPlayerPositionPkt = internal.PlayPlayerPosition_758_0

// ID=0xd
type PlayPlayerPositionAndLookServerPkt = internal.PlayPlayerPositionAndLook_338_4

// ID=0xe
type PlayPlayerLookPkt = internal.PlayPlayerLook_404_0

// ID=0xf
type PlayPlayerPkt = internal.PlayPlayer_404_0

// ID=0x10
type PlayVehicleMoveServerPkt = internal.PlayVehicleMove_758_0

// ID=0x11
type PlaySteerBoatPkt = internal.PlaySteerBoat_401_1

// ID=0x12
type PlayPlayerAbilitiesServerPkt = internal.PlayPlayerAbilities_404_12

// ID=0x13
type PlayPlayerDiggingPkt = internal.PlayPlayerDigging_758_0

// ID=0x14
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x15
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x16
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_110_1

// ID=0x17
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChange_404_7

// ID=0x18
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x19
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x1a
type PlayAnimationServerPkt = internal.PlayAnimation_338_4

// ID=0x1b
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x1c
type PlayPlayerBlockPlacementPkt = internal.PlayPlayerBlockPlacement_210_2

// ID=0x1d
type PlayUseItemPkt = internal.PlayUseItem_758_1

// ---- play: clientbound ----

// ID=0x0
type PlaySpawnObjectPkt = internal.PlaySpawnObject_404_1

// ID=0x1
type PlaySpawnExperienceOrbPkt = internal.PlaySpawnExperienceOrb_763_0

// ID=0x2
type PlaySpawnGlobalEntityPkt = internal.PlaySpawnGlobalEntity_498_0

// ID=0x3
type PlaySpawnMobPkt = internal.PlaySpawnMob_210_1

// ID=0x4
type PlaySpawnPaintingPkt = internal.PlaySpawnPainting_340_1

// ID=0x5
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_498_1

// ID=0x6
type PlayAnimationClientPkt = internal.PlayAnimation_338_3

// ID=0x7
type PlayStatisticsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name | Field Type | Field Type | Notes                                                     |
	 * |-----------|-------|----------|------------|------------|------------|------------|-----------------------------------------------------------|
	 * | 0x07      | Play  | Client   | Count      | Count      | VarInt     | VarInt     | Number of elements in the following array                 |
	 * | 0x07      | Play  | Client   | Statistic  | Name       | Array      | String     | https://gist.github.com/thinkofdeath/a1842c21a0cf2e1fb5e0 |
	 * | 0x07      | Play  | Client   | Statistic  | Value      | Array      | VarInt     | The amount to set it to                                   |
	 * 
	 */
}

// ID=0x8
type PlayBlockBreakAnimationPkt = internal.PlayBlockBreakAnimation_758_0

// ID=0x9
type PlayUpdateBlockEntityPkt = internal.PlayUpdateBlockEntity_498_0

// ID=0xa
type PlayBlockActionPkt = internal.PlayBlockAction_110_2

// ID=0xb
type PlayBlockChangePkt = internal.PlayBlockChange_758_0

// ID=0xc
type PlayBossBarPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name       | Field Name | Field Type    | Notes                                                                                                                             |
	 * |-----------|-------|----------|------------------|------------|---------------|-----------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0C      | Play  | Client   | UUID             | UUID       | UUID          | Unique ID for this bar                                                                                                            |
	 * | 0x0C      | Play  | Client   | Action           | Action     | VarInt Enum   | Determines the layout of the remaining packet                                                                                     |
	 * | 0x0C      | Play  | Client   | Action           | Field Name |               |                                                                                                                                   |
	 * | 0x0C      | Play  | Client   | 0: add           | Title      | Chat          |                                                                                                                                   |
	 * | 0x0C      | Play  | Client   | 0: add           | Health     | Float         | From 0 to 1. Values greater than 1 do not crash a Notchian client, and start rendering part of a second health bar at around 1.5. |
	 * | 0x0C      | Play  | Client   | 0: add           | Color      | VarInt Enum   | Color ID (see below)                                                                                                              |
	 * | 0x0C      | Play  | Client   | 0: add           | Division   | VarInt Enum   | Type of division (see below)                                                                                                      |
	 * | 0x0C      | Play  | Client   | 0: add           | Flags      | Unsigned Byte | Bit mask. 0x1: should darken sky, 0x2: is dragon bar (used to play end music)                                                     |
	 * | 0x0C      | Play  | Client   | 1: remove        | no fields  | no fields     | Removes this boss bar                                                                                                             |
	 * | 0x0C      | Play  | Client   | 2: update health | Health     | Float         | as above                                                                                                                          |
	 * | 0x0C      | Play  | Client   | 3: update title  | Title      | Chat          |                                                                                                                                   |
	 * | 0x0C      | Play  | Client   | 4: update style  | Color      | VarInt Enum   | Color ID (see below)                                                                                                              |
	 * | 0x0C      | Play  | Client   | 4: update style  | Dividers   | VarInt Enum   | as above                                                                                                                          |
	 * | 0x0C      | Play  | Client   | 5: update flags  | Flags      | Unsigned Byte | as above                                                                                                                          |
	 * 
	 */
}

// ID=0xd
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_404_1

// ID=0xe
type PlayTabCompleteClientPkt = internal.PlayTabComplete_338_5

// ID=0xf
type PlayChatMessageClientPkt = internal.PlayChatMessage_404_12

// ID=0x10
type PlayMultiBlockChangePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name   | Field Name          | Field Type | Field Type    | Notes                                                                                                                                                                    |
	 * |-----------|-------|----------|--------------|---------------------|------------|---------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x10      | Play  | Client   | Chunk X      | Chunk X             | Int        | Int           | Chunk X coordinate                                                                                                                                                       |
	 * | 0x10      | Play  | Client   | Chunk Z      | Chunk Z             | Int        | Int           | Chunk Z coordinate                                                                                                                                                       |
	 * | 0x10      | Play  | Client   | Record Count | Record Count        | VarInt     | VarInt        | Number of elements in the following array, i.e. the number of blocks affected                                                                                            |
	 * | 0x10      | Play  | Client   | Record       | Horizontal Position | Array      | Unsigned Byte | The 4 most significant bits (0xF0) encode the X coordinate, relative to the chunk. The 4 least significant bits (0x0F) encode the Z coordinate, relative to the chunk.   |
	 * | 0x10      | Play  | Client   | Record       | Y Coordinate        | Array      | Unsigned Byte | Y coordinate of the block                                                                                                                                                |
	 * | 0x10      | Play  | Client   | Record       | Block ID            | Array      | VarInt        | The new block state ID for the block as given in the global palette (When reading data: type = id >> 4, meta = id & 15, when writing data: id = type << 4 | (meta & 15)) |
	 * 
	 */
}

// ID=0x11
type PlayConfirmTransactionClientPkt = internal.PlayConfirmTransaction_404_0

// ID=0x12
type PlayCloseWindowClientPkt = internal.PlayCloseWindow_758_0

// ID=0x13
type PlayOpenWindowPkt = internal.PlayOpenWindow_404_1

// ID=0x14
type PlayWindowItemsPkt = internal.PlayWindowItems_755_1

// ID=0x15
type PlayWindowPropertyPkt = internal.PlayWindowProperty_758_0

// ID=0x16
type PlaySetSlotPkt = internal.PlaySetSlot_755_1

// ID=0x17
type PlaySetCooldownPkt = internal.PlaySetCooldown_763_0

// ID=0x18
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_340_1

// ID=0x19
type PlayNamedSoundEffectPkt struct {
	/* All sound effect names as of 1.9.0 can be seen here. */
	SoundName String // String
	/* The category that this sound will be played from (current categories) */
	SoundCategory VarInt // VarInt Enum
	/* Effect X multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionX Int // Int
	/* Effect Y multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionY Int // Int
	/* Effect Z multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionZ Int // Int
	/* 1 is 100%, can be more */
	Volume Float // Float
	/* 63 is 100%, can be more */
	Pitch UByte // Unsigned Byte
}

var _ Packet = (*PlayNamedSoundEffectPkt)(nil)

func (p PlayNamedSoundEffectPkt)Encode(b *PacketBuilder){
	b.String(p.SoundName)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.UByte(p.Pitch)
}

func (p *PlayNamedSoundEffectPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SoundName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.SoundCategory, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectPositionX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EffectPositionY, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EffectPositionZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Volume, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// ID=0x1a
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x1b
type PlayEntityStatusPkt = internal.PlayEntityStatus_758_0

// ID=0x1c
type PlayExplosionPkt = internal.PlayExplosion_498_3

// ID=0x1d
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x1e
type PlayChangeGameStatePkt = internal.PlayChangeGameState_758_0

// ID=0x1f
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_338_1

// ID=0x20
type PlayChunkDataPkt = internal.PlayChunkData_316_6

// ID=0x21
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x22
type PlayParticlePkt = internal.PlayParticle_340_3

// ID=0x23
type PlayJoinGamePkt = internal.PlayJoinGame_404_4

// ID=0x24
type PlayMapPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name         | Field Type                      | Field Type                      | Notes                                                                                                      |
	 * |-----------|-------|----------|-------------------|--------------------|---------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x24      | Play  | Client   | Item Damage       | Item Damage        | VarInt                          | VarInt                          | The damage value (map ID) of the map being modified                                                        |
	 * | 0x24      | Play  | Client   | Scale             | Scale              | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x24      | Play  | Client   | Tracking Position | Tracking Position  | Boolean                         | Boolean                         | Specifies whether the icons are shown                                                                      |
	 * | 0x24      | Play  | Client   | Icon Count        | Icon Count         | VarInt                          | VarInt                          | Number of elements in the following array                                                                  |
	 * | 0x24      | Play  | Client   | Icon              | Direction And Type | Array                           | Byte                            | 0xF0 = Direction, 0x0F = Type                                                                              |
	 * | 0x24      | Play  | Client   | Icon              | X                  | Array                           | Byte                            |                                                                                                            |
	 * | 0x24      | Play  | Client   | Icon              | Z                  | Array                           | Byte                            |                                                                                                            |
	 * | 0x24      | Play  | Client   | Columns           | Columns            | Byte                            | Byte                            | Number of columns updated                                                                                  |
	 * | 0x24      | Play  | Client   | Rows              | Rows               | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; number of rows updated                                                     |
	 * | 0x24      | Play  | Client   | X                 | X                  | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column                                         |
	 * | 0x24      | Play  | Client   | Z                 | Z                  | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row                                           |
	 * | 0x24      | Play  | Client   | Length            | Length             | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array                                              |
	 * | 0x24      | Play  | Client   | Data              | Data               | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format                                                        |
	 * 
	 */
}

// ID=0x25
type PlayEntityRelativeMovePkt = internal.PlayEntityRelativeMove_404_0

// ID=0x26
type PlayEntityLookAndRelativeMovePkt = internal.PlayEntityLookAndRelativeMove_404_0

// ID=0x27
type PlayEntityLookPkt = internal.PlayEntityLook_404_0

// ID=0x28
type PlayEntityPkt = internal.PlayEntity_404_0

// ID=0x29
type PlayVehicleMoveClientPkt = internal.PlayVehicleMove_758_0

// ID=0x2a
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x2b
type PlayPlayerAbilitiesClientPkt = internal.PlayPlayerAbilities_404_11

// ID=0x2c
type PlayCombatEventPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name      | Field Name | Field Type  | Notes                                         |
	 * |-----------|-------|----------|-----------------|------------|-------------|-----------------------------------------------|
	 * | 0x2C      | Play  | Client   | Event           | Event      | VarInt Enum | Determines the layout of the remaining packet |
	 * | 0x2C      | Play  | Client   | Event           | Field Name |             |                                               |
	 * | 0x2C      | Play  | Client   | 0: enter combat | no fields  | no fields   |                                               |
	 * | 0x2C      | Play  | Client   | 1: end combat   | Duration   | VarInt      |                                               |
	 * | 0x2C      | Play  | Client   | 1: end combat   | Entity ID  | Int         |                                               |
	 * | 0x2C      | Play  | Client   | 2: entity dead  | Player ID  | VarInt      |                                               |
	 * | 0x2C      | Play  | Client   | 2: entity dead  | Entity ID  | Int         |                                               |
	 * | 0x2C      | Play  | Client   | 2: entity dead  | Message    | Chat        |                                               |
	 * 
	 */
}

// ID=0x2d
type PlayPlayerListItemPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type      | Notes                                                   |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-----------------|---------------------------------------------------------|
	 * | 0x2D      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt          | Determines the rest of the Player format after the UUID |
	 * | 0x2D      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt          | Number of elements in the following array               |
	 * | 0x2D      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID            |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                 |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String        | String          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt          | Number of elements in the following array               |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean         |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String | Only if Is Signed is true                               |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean         |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat   | Only if Has Display Name is true                        |
	 * | 0x2D      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean         |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat   | Only send if Has Display Name is true                   |
	 * | 0x2D      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields       |                                                         |
	 * 
	 */
}

// ID=0x2e
type PlayPlayerPositionAndLookClientPkt = internal.PlayPlayerPositionAndLook_338_3

// ID=0x2f
type PlayUseBedPkt = internal.PlayUseBed_404_0

// ID=0x30
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x31
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x32
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x33
type PlayRespawnPkt = internal.PlayRespawn_404_7

// ID=0x34
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x35
type PlayWorldBorderPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name            | Field Name               | Field Type  | Notes                                                                                                                                                                                                                                        |
	 * |-----------|-------|----------|-----------------------|--------------------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x35      | Play  | Client   | Action                | Action                   | VarInt Enum | Determines the format of the rest of the packet                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | Action                | Field Name               |             |                                                                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | 0: set size           | Diameter                 | Double      | Length of a single side of the world border, in meters                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | 1: lerp size          | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x35      | Play  | Client   | 1: lerp size          | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x35      | Play  | Client   | 1: lerp size          | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x35      | Play  | Client   | 2: set center         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | 2: set center         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | 3: initialize         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | 3: initialize         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | 3: initialize         | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x35      | Play  | Client   | 3: initialize         | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x35      | Play  | Client   | 3: initialize         | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x35      | Play  | Client   | 3: initialize         | Portal Teleport Boundary | VarInt      | Resulting coordinates from a portal teleport are limited to ±value. Usually 29999984.                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | 3: initialize         | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x35      | Play  | Client   | 3: initialize         | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * | 0x35      | Play  | Client   | 4: set warning time   | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x35      | Play  | Client   | 5: set warning blocks | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * 
	 */
}

// ID=0x36
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x37
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_404_6

// ID=0x38
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x39
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x3a
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x3b
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x3c
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_578_0

// ID=0x3d
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x3e
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x3f
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_340_1

// ID=0x40
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x41
type PlayTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                  | Field Name          | Field Type      | Notes                                                                             |
	 * |-----------|-------|----------|-----------------------------|---------------------|-----------------|-----------------------------------------------------------------------------------|
	 * | 0x41      | Play  | Client   | Team Name                   | Team Name           | String          | A unique name for the team. (Shared with scoreboard).                             |
	 * | 0x41      | Play  | Client   | Mode                        | Mode                | Byte            | Determines the layout of the remaining packet                                     |
	 * | 0x41      | Play  | Client   | 0: create team              | Team Display Name   | String          |                                                                                   |
	 * | 0x41      | Play  | Client   | 0: create team              | Team Prefix         | String          | Displayed before the names of players that are part of this team                  |
	 * | 0x41      | Play  | Client   | 0: create team              | Team Suffix         | String          | Displayed after the names of players that are part of this team                   |
	 * | 0x41      | Play  | Client   | 0: create team              | Friendly Flags      | Byte            | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team |
	 * | 0x41      | Play  | Client   | 0: create team              | Name Tag Visibility | String Enum     | always, hideForOtherTeams, hideForOwnTeam, never                                  |
	 * | 0x41      | Play  | Client   | 0: create team              | Collision Rule      | String Enum     | always, pushOtherTeams, pushOwnTeam, never                                        |
	 * | 0x41      | Play  | Client   | 0: create team              | Color               | Byte            | Same as Chat colors                                                               |
	 * | 0x41      | Play  | Client   | 0: create team              | Player Count        | VarInt          | Number of elements in the following array                                         |
	 * | 0x41      | Play  | Client   | 0: create team              | Players             | Array of String | Usernames of the players in this team. Maximum 40 characters per entry.           |
	 * | 0x41      | Play  | Client   | 1: remove team              | no fields           | no fields       |                                                                                   |
	 * | 0x41      | Play  | Client   | 2: update team info         | Team Display Name   | String          |                                                                                   |
	 * | 0x41      | Play  | Client   | 2: update team info         | Team Prefix         | String          | Displayed before the names of players that are part of this team                  |
	 * | 0x41      | Play  | Client   | 2: update team info         | Team Suffix         | String          | Displayed after the names of players that are part of this team                   |
	 * | 0x41      | Play  | Client   | 2: update team info         | Friendly Flags      | Byte            | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team |
	 * | 0x41      | Play  | Client   | 2: update team info         | Name Tag Visibility | String Enum     | always, hideForOtherTeams, hideForOwnTeam, never                                  |
	 * | 0x41      | Play  | Client   | 2: update team info         | Collision Rule      | String Enum     | always, pushOtherTeams, pushOwnTeam, never                                        |
	 * | 0x41      | Play  | Client   | 2: update team info         | Color               | Byte            | Same as Chat colors                                                               |
	 * | 0x41      | Play  | Client   | 3: add players to team      | Player Count        | VarInt          | Number of elements in the following array                                         |
	 * | 0x41      | Play  | Client   | 3: add players to team      | Players             | Array of String | Usernames of the added players. Maximum 40 characters per entry.                  |
	 * | 0x41      | Play  | Client   | 4: remove players from team | Player Count        | VarInt          | Number of elements in the following array                                         |
	 * | 0x41      | Play  | Client   | 4: remove players from team | Players             | Array of String | Usernames of the removed players. Maximum 40 characters per entry.                |
	 * 
	 */
}

// ID=0x42
type PlayUpdateScorePkt = internal.PlayUpdateScore_315_2

// ID=0x43
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_754_1

// ID=0x44
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x45
type PlayTitlePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name    | Field Type  | Notes                                            |
	 * |-----------|-------|----------|--------------------------|---------------|-------------|--------------------------------------------------|
	 * | 0x45      | Play  | Client   | Action                   | Action        | VarInt Enum |                                                  |
	 * | 0x45      | Play  | Client   | Action                   | Field Name    |             |                                                  |
	 * | 0x45      | Play  | Client   | 0: set title             | Title Text    | Chat        |                                                  |
	 * | 0x45      | Play  | Client   | 1: set subtitle          | Subtitle Text | Chat        |                                                  |
	 * | 0x45      | Play  | Client   | 2: set times and display | Fade In       | Int         | Ticks to spend fading in                         |
	 * | 0x45      | Play  | Client   | 2: set times and display | Stay          | Int         | Ticks to keep the title displayed                |
	 * | 0x45      | Play  | Client   | 2: set times and display | Fade Out      | Int         | Ticks to spend out, not when to start fading out |
	 * | 0x45      | Play  | Client   | 3: hide                  | no fields     | no fields   |                                                  |
	 * | 0x45      | Play  | Client   | 4: reset                 | no fields     | no fields   |                                                  |
	 * 
	 */
}

// ID=0x46
type PlaySoundEffectPkt struct {
	/* ID of hardcoded sound event (events as of 1.9.0) */
	SoundID VarInt // VarInt
	/* The category that this sound will be played from (current categories) */
	SoundCategory VarInt // VarInt Enum
	/* Effect X multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionX Int // Int
	/* Effect Y multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionY Int // Int
	/* Effect Z multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionZ Int // Int
	/* 1.0 is 100%, capped between 0.0 and 1.0 by Notchian clients */
	Volume Float // Float
	/* Float multiplied by 63.5, capped between 0.5 and 2.0 by Notchian clients */
	Pitch UByte // Unsigned Byte
}

var _ Packet = (*PlaySoundEffectPkt)(nil)

func (p PlaySoundEffectPkt)Encode(b *PacketBuilder){
	b.VarInt(p.SoundID)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.UByte(p.Pitch)
}

func (p *PlaySoundEffectPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SoundID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.SoundCategory, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectPositionX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EffectPositionY, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EffectPositionZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Volume, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// ID=0x47
type PlayPlayerListHeaderAndFooterPkt = internal.PlayPlayerListHeaderAndFooter_758_0

// ID=0x48
type PlayCollectItemPkt = internal.PlayCollectItem_210_1

// ID=0x49
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x4a
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x4A      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x4A      | Play  | Client   | Number Of Properties | Number Of Properties | Int        | Int                    | Number of elements in the following array             |
	 * | 0x4A      | Play  | Client   | Property             | Key                  | Array      | String                 | See below                                             |
	 * | 0x4A      | Play  | Client   | Property             | Value                | Array      | Double                 | See below                                             |
	 * | 0x4A      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array             |
	 * | 0x4A      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x4b
type PlayEntityEffectPkt = internal.PlayEntityEffect_110_3

// ======== END play ========
