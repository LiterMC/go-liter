
// Generated at 2023-09-05 22:06:22.506 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=14889
// Protocol: 404
// Protocol Name: 1.13.2

package packet_1_13_2

import (
	"io"
	. "github.com/kmcsr/go-liter"
	nbt "github.com/kmcsr/go-liter/nbt"
	data "github.com/kmcsr/go-liter/data"
	internal "github.com/kmcsr/go-liter/packets/internal"
)

func assert(cond bool, msg any){
	if !cond {
		panic(msg)
	}
}

// ======== BEGIN login ========
// ---- login: serverbound ----

// ID=0x0
type LoginStartPkt = internal.LoginStart_758_2

// ID=0x1
type LoginEncryptionResponsePkt = internal.LoginEncryptionResponse_758_2

// ID=0x2
type LoginPluginResponsePkt = internal.LoginPluginResponse_763_0

// ---- login: clientbound ----

// ID=0x0
type LoginDisconnectPkt = internal.LoginDisconnect_763_0

// ID=0x1
type LoginEncryptionRequestPkt = internal.LoginEncryptionRequest_763_0

// ID=0x2
type LoginSuccessPkt = internal.LoginSuccess_578_2

// ID=0x3
type LoginSetCompressionPkt = internal.LoginSetCompression_763_0

// ID=0x4
type LoginPluginRequestPkt = internal.LoginPluginRequest_763_0

// ======== END login ========

// ======== BEGIN play ========
// ---- play: serverbound ----

// ID=0x0
type PlayTeleportConfirmPkt = internal.PlayTeleportConfirm_758_0

// ID=0x1
type PlayQueryBlockNBTPkt = internal.PlayQueryBlockNBT_758_0

// ID=0x2
type PlayChatMessageServerPkt = internal.PlayChatMessageServer_758_0

// ID=0x3
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x4
type PlayClientSettingsPkt = internal.PlayClientSettings_754_2

// ID=0x5
type PlayTabCompleteServerPkt = internal.PlayTabCompleteServer_758_0

// ID=0x6
type PlayConfirmTransactionServerPkt = internal.PlayConfirmTransactionServer_404_0

// ID=0x7
type PlayEnchantItemPkt = internal.PlayEnchantItem_404_0

// ID=0x8
type PlayClickWindowPkt = internal.PlayClickWindow_754_4

// ID=0x9
type PlayCloseWindowServerPkt = internal.PlayCloseWindowServer_758_0

// ID=0xa
type PlayPluginMessageServerPkt = internal.PlayPluginMessageServer_763_0

// ID=0xb
type PlayEditBookPkt = internal.PlayEditBook_755_4

// ID=0xc
type PlayQueryEntityNBTPkt = internal.PlayQueryEntityNBT_758_0

// ID=0xd
type PlayUseEntityPkt = internal.PlayUseEntity_404_0

// ID=0xe
type PlayKeepAliveServerPkt = internal.PlayKeepAliveServer_763_0

// ID=0xf
type PlayerPkt = internal.Player_404_0

// ID=0x10
type PlayerPositionPkt = internal.PlayerPosition_758_0

// ID=0x11
type PlayerPositionAndLookServerPkt = internal.PlayerPositionAndLookServer_404_0

// ID=0x12
type PlayerLookPkt = internal.PlayerLook_404_0

// ID=0x13
type PlayVehicleMoveServerPkt = internal.PlayVehicleMoveServer_758_0

// ID=0x14
type PlaySteerBoatPkt = internal.PlaySteerBoat_758_0

// ID=0x15
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x16
type PlayCraftRecipeRequestPkt = internal.PlayCraftRecipeRequest_758_0

// ID=0x17
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_578_1

// ID=0x18
type PlayerDiggingPkt = internal.PlayerDigging_758_0

// ID=0x19
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x1a
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x1b
type PlayRecipeBookDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name            | Field Name                    | Field Type | Notes                                                                           |
	 * |-----------|-------|----------|-----------------------|-------------------------------|------------|---------------------------------------------------------------------------------|
	 * | 0x1B      | Play  | Server   | Type                  | Type                          | VarInt     | Determines the format of the rest of the packet                                 |
	 * | 0x1B      | Play  | Server   | Type                  | Field Name                    |            |                                                                                 |
	 * | 0x1B      | Play  | Server   | 0: Displayed Recipe   | Recipe ID                     | Identifier | A recipe ID                                                                     |
	 * | 0x1B      | Play  | Server   | 1: Recipe Book States | Crafting Recipe Book Open     | Boolean    | Whether the player has the crafting recipe book currently opened/active.        |
	 * | 0x1B      | Play  | Server   | 1: Recipe Book States | Crafting Recipe Filter Active | Boolean    | Whether the player has the crafting recipe book filter option currently active. |
	 * | 0x1B      | Play  | Server   | 1: Recipe Book States | Smelting Recipe Book Open     | Boolean    | Whether the player has the smelting recipe book currently opened/active.        |
	 * | 0x1B      | Play  | Server   | 1: Recipe Book States | Smelting Recipe Filter Active | Boolean    | Whether the player has the smelting recipe book filter option currently active. |
	 * 
	 */
}

// ID=0x1c
type PlayNameItemPkt = internal.PlayNameItem_758_0

// ID=0x1d
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_758_0

// ID=0x1e
type PlayAdvancementTabPkt = internal.PlayAdvancementTab_758_0

// ID=0x1f
type PlaySelectTradePkt = internal.PlaySelectTrade_763_0

// ID=0x20
type PlaySetBeaconEffectPkt = internal.PlaySetBeaconEffect_758_2

// ID=0x21
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChangeServer_758_0

// ID=0x22
type PlayUpdateCommandBlockPkt = internal.PlayUpdateCommandBlock_758_0

// ID=0x23
type PlayUpdateCommandBlockMinecartPkt = internal.PlayUpdateCommandBlockMinecart_758_0

// ID=0x24
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x25
type PlayUpdateStructureBlockPkt = internal.PlayUpdateStructureBlock_758_0

// ID=0x26
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x27
type PlayAnimationServerPkt = internal.PlayAnimationServer_404_0

// ID=0x28
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x29
type PlayerBlockPlacementPkt = internal.PlayerBlockPlacement_404_1

// ID=0x2a
type PlayUseItemPkt = internal.PlayUseItem_758_1

// ---- play: clientbound ----

// ID=0x0
type PlaySpawnObjectPkt = internal.PlaySpawnObject_404_1

// ID=0x1
type PlaySpawnExperienceOrbPkt = internal.PlaySpawnExperienceOrb_763_0

// ID=0x2
type PlaySpawnGlobalEntityPkt = internal.PlaySpawnGlobalEntity_498_0

// ID=0x3
type PlaySpawnMobPkt = internal.PlaySpawnMob_498_0

// ID=0x4
type PlaySpawnPaintingPkt = internal.PlaySpawnPainting_758_0

// ID=0x5
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_498_1

// ID=0x6
type PlayAnimationClientPkt = internal.PlayAnimation_404_1

// ID=0x7
type PlayStatisticsPkt = internal.PlayStatistics_498_3

// ID=0x8
type PlayBlockBreakAnimationPkt = internal.PlayBlockBreakAnimation_758_0

// ID=0x9
type PlayUpdateBlockEntityPkt = internal.PlayUpdateBlockEntity_498_0

// ID=0xa
type PlayBlockActionPkt = internal.PlayBlockAction_758_1

// ID=0xb
type PlayBlockChangePkt = internal.PlayBlockChange_758_0

// ID=0xc
type PlayBossBarPkt = internal.PlayBossBar_404_5

// ID=0xd
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_404_1

// ID=0xe
type PlayChatMessageClientPkt = internal.PlayChatMessage_578_4

// ID=0xf
type PlayMultiBlockChangePkt = internal.PlayMultiBlockChange_498_2

// ID=0x10
type PlayTabCompleteClientPkt = internal.PlayTabComplete_498_4

// ID=0x11
type PlayDeclareCommandsPkt = internal.PlayDeclareCommands_758_0

// ID=0x12
type PlayConfirmTransactionClientPkt = internal.PlayConfirmTransaction_404_0

// ID=0x13
type PlayCloseWindowClientPkt = internal.PlayCloseWindow_758_0

// ID=0x14
type PlayOpenWindowPkt = internal.PlayOpenWindow_404_1

// ID=0x15
type PlayWindowItemsPkt = internal.PlayWindowItems_755_1

// ID=0x16
type PlayWindowPropertyPkt = internal.PlayWindowProperty_758_0

// ID=0x17
type PlaySetSlotPkt = internal.PlaySetSlot_755_1

// ID=0x18
type PlaySetCooldownPkt = internal.PlaySetCooldown_763_0

// ID=0x19
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_763_0

// ID=0x1a
type PlayNamedSoundEffectPkt = internal.PlayNamedSoundEffect_758_0

// ID=0x1b
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x1c
type PlayEntityStatusPkt = internal.PlayEntityStatus_758_0

// ID=0x1d
type PlayNBTQueryResponsePkt = internal.PlayNBTQueryResponse_758_0

// ID=0x1e
type PlayExplosionPkt = internal.PlayExplosion_498_3

// ID=0x1f
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x20
type PlayChangeGameStatePkt = internal.PlayChangeGameState_758_0

// ID=0x21
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x22
type PlayChunkDataPkt struct {
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkX Int // Int
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkZ Int // Int
	/* See Chunk Format */
	FullChunk Bool // Boolean
	/* Bitmask with bits set to 1 for every 16×16×16 chunk section whose data is included in Data. The least significant bit represents the chunk section at the bottom of the chunk column (from y=0 to y=15). */
	PrimaryBitMask VarInt // VarInt
	/* Size of Data in bytes */
	Size VarInt // VarInt
	/* See data structure in Chunk Format */
	Data ByteArray // Byte array
	/* Number of elements in the following array */
	NumberOfBlockEntities VarInt // VarInt
	/* All block entities in the chunk.  Use the x, y, and z tags in the NBT to determine their positions. */
	BlockEntities []nbt.NBT // Array of NBT Tag
}

var _ Packet = (*PlayChunkDataPkt)(nil)

func (p PlayChunkDataPkt)Encode(b *PacketBuilder){
	b.Int(p.ChunkX)
	b.Int(p.ChunkZ)
	b.Bool(p.FullChunk)
	b.VarInt(p.PrimaryBitMask)
	p.Size = (VarInt)(len(p.Data))
	b.VarInt(p.Size)
	b.ByteArray(p.Data)
	p.NumberOfBlockEntities = (VarInt)(len(p.BlockEntities))
	b.VarInt(p.NumberOfBlockEntities)
	for _, v := range p.BlockEntities {
		WriteNBT(b, v)
	}
}

func (p *PlayChunkDataPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.ChunkX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.FullChunk, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.PrimaryBitMask, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Size, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Data = make(ByteArray, p.Size)
	if ok = r.ByteArray(p.Data); !ok {
		return io.EOF
	}
	if p.NumberOfBlockEntities, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.BlockEntities = make([]nbt.NBT, p.NumberOfBlockEntities)
	for i, _ := range p.BlockEntities {
		if p.BlockEntities[i], err = nbt.ReadNBT(r); err != nil {
			return err
		}
	}
	return nil
}

// ID=0x23
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x24
type PlayParticlePkt = internal.PlayParticle_498_2

// ID=0x25
type PlayJoinGamePkt = internal.PlayJoinGame_404_4

// ID=0x26
type PlayMapDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type                      | Field Type                      | Notes                                                                                                      |
	 * |-----------|-------|----------|-------------------|-------------------|---------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x26      | Play  | Client   | Map ID            | Map ID            | VarInt                          | VarInt                          | Map ID of the map being modified                                                                           |
	 * | 0x26      | Play  | Client   | Scale             | Scale             | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x26      | Play  | Client   | Tracking Position | Tracking Position | Boolean                         | Boolean                         | Specifies whether player and item frame icons are shown                                                    |
	 * | 0x26      | Play  | Client   | Icon Count        | Icon Count        | VarInt                          | VarInt                          | Number of elements in the following array                                                                  |
	 * | 0x26      | Play  | Client   | Icon              | Type              | Array                           | VarInt enum                     | See below                                                                                                  |
	 * | 0x26      | Play  | Client   | Icon              | X                 | Array                           | Byte                            | Map coordinates: -128 for furthest left, +127 for furthest right                                           |
	 * | 0x26      | Play  | Client   | Icon              | Z                 | Array                           | Byte                            | Map coordinates: -128 for highest, +127 for lowest                                                         |
	 * | 0x26      | Play  | Client   | Icon              | Direction         | Array                           | Byte                            | 0-15                                                                                                       |
	 * | 0x26      | Play  | Client   | Icon              | Has Display Name  | Array                           | Boolean                         |                                                                                                            |
	 * | 0x26      | Play  | Client   | Icon              | Display Name      | Array                           | Optional Chat                   | Only present if previous Boolean is true                                                                   |
	 * | 0x26      | Play  | Client   | Columns           | Columns           | Unsigned Byte                   | Unsigned Byte                   | Number of columns updated                                                                                  |
	 * | 0x26      | Play  | Client   | Rows              | Rows              | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; number of rows updated                                                     |
	 * | 0x26      | Play  | Client   | X                 | X                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column                                         |
	 * | 0x26      | Play  | Client   | Z                 | Z                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row                                           |
	 * | 0x26      | Play  | Client   | Length            | Length            | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array                                              |
	 * | 0x26      | Play  | Client   | Data              | Data              | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format                                                        |
	 * 
	 */
}

// ID=0x27
type PlayEntityPkt = internal.PlayEntity_404_0

// ID=0x28
type PlayEntityRelativeMovePkt = internal.PlayEntityRelativeMove_404_0

// ID=0x29
type PlayEntityLookAndRelativeMovePkt = internal.PlayEntityLookAndRelativeMove_404_0

// ID=0x2a
type PlayEntityLookPkt = internal.PlayEntityLook_404_0

// ID=0x2b
type PlayVehicleMoveClientPkt = internal.PlayVehicleMove_758_0

// ID=0x2c
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x2d
type PlayCraftRecipeResponsePkt = internal.PlayCraftRecipeResponse_758_0

// ID=0x2e
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x2f
type PlayCombatEventPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name      | Field Name | Field Type  | Notes                                         |
	 * |-----------|-------|----------|-----------------|------------|-------------|-----------------------------------------------|
	 * | 0x2F      | Play  | Client   | Event           | Event      | VarInt Enum | Determines the layout of the remaining packet |
	 * | 0x2F      | Play  | Client   | Event           | Field Name |             |                                               |
	 * | 0x2F      | Play  | Client   | 0: enter combat | no fields  | no fields   |                                               |
	 * | 0x2F      | Play  | Client   | 1: end combat   | Duration   | VarInt      |                                               |
	 * | 0x2F      | Play  | Client   | 1: end combat   | Entity ID  | Int         |                                               |
	 * | 0x2F      | Play  | Client   | 2: entity dead  | Player ID  | VarInt      |                                               |
	 * | 0x2F      | Play  | Client   | 2: entity dead  | Entity ID  | Int         |                                               |
	 * | 0x2F      | Play  | Client   | 2: entity dead  | Message    | Chat        |                                               |
	 * 
	 */
}

// ID=0x30
type PlayerInfoPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type              | Notes                                                   |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-------------------------|---------------------------------------------------------|
	 * | 0x30      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt                  | Determines the rest of the Player format after the UUID |
	 * | 0x30      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x30      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID                    |                                                         |
	 * | 0x30      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                         |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String (16)   | String (16)             |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String (32767)          |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String (32767)          |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean                 |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String (32767) | Only if Is Signed is true                               |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds                                |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only if Has Display Name is true                        |
	 * | 0x30      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds                                |
	 * | 0x30      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true                   |
	 * | 0x30      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields               |                                                         |
	 * 
	 */
}

// ID=0x31
type PlayFacePlayerPkt = internal.PlayFacePlayer_757_0

// ID=0x32
type PlayerPositionAndLookClientPkt = internal.PlayerPositionAndLook_754_1

// ID=0x33
type PlayUseBedPkt = internal.PlayUseBed_404_0

// ID=0x34
type PlayUnlockRecipesPkt = internal.PlayUnlockRecipes_578_2

// ID=0x35
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x36
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x37
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x38
type PlayRespawnPkt = internal.PlayRespawn_404_8

// ID=0x39
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x3a
type PlaySelectAdvancementTabPkt = internal.PlaySelectAdvancementTab_758_0

// ID=0x3b
type PlayWorldBorderPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name            | Field Name               | Field Type  | Notes                                                                                                                                                                                                                                        |
	 * |-----------|-------|----------|-----------------------|--------------------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x3B      | Play  | Client   | Action                | Action                   | VarInt Enum | Determines the format of the rest of the packet                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | Action                | Field Name               |             |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 0: set size           | Diameter                 | Double      | Length of a single side of the world border, in meters                                                                                                                                                                                       |
	 * | 0x3B      | Play  | Client   | 1: lerp size          | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x3B      | Play  | Client   | 1: lerp size          | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x3B      | Play  | Client   | 1: lerp size          | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x3B      | Play  | Client   | 2: set center         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 2: set center         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 3: initialize         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x3B      | Play  | Client   | 3: initialize         | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Portal Teleport Boundary | VarInt      | Resulting coordinates from a portal teleport are limited to ±value. Usually 29999984.                                                                                                                                                       |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * | 0x3B      | Play  | Client   | 4: set warning time   | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x3B      | Play  | Client   | 5: set warning blocks | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * 
	 */
}

// ID=0x3c
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x3d
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_758_0

// ID=0x3e
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x3f
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x40
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x41
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x42
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_578_2

// ID=0x43
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x44
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x45
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_758_0

// ID=0x46
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x47
type PlayTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                  | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|-----------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x47      | Play  | Client   | Team Name                   | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x47      | Play  | Client   | Mode                        | Mode                | Byte                 | Determines the layout of the remaining packet                                                                          |
	 * | 0x47      | Play  | Client   | 0: create team              | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x47      | Play  | Client   | 0: create team              | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team                                      |
	 * | 0x47      | Play  | Client   | 0: create team              | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x47      | Play  | Client   | 0: create team              | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x47      | Play  | Client   | 0: create team              | Team Color          | VarInt enum          | Used to color the name of players on the team; see below                                                               |
	 * | 0x47      | Play  | Client   | 0: create team              | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team                                                       |
	 * | 0x47      | Play  | Client   | 0: create team              | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team                                                        |
	 * | 0x47      | Play  | Client   | 0: create team              | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x47      | Play  | Client   | 0: create team              | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x47      | Play  | Client   | 1: remove team              | no fields           | no fields            |                                                                                                                        |
	 * | 0x47      | Play  | Client   | 2: update team info         | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x47      | Play  | Client   | 2: update team info         | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team                                     |
	 * | 0x47      | Play  | Client   | 2: update team info         | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x47      | Play  | Client   | 2: update team info         | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x47      | Play  | Client   | 2: update team info         | Team Color          | VarInt enum          | Used to color the name of players on the team; see below                                                               |
	 * | 0x47      | Play  | Client   | 2: update team info         | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team                                                       |
	 * | 0x47      | Play  | Client   | 2: update team info         | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team                                                        |
	 * | 0x47      | Play  | Client   | 3: add players to team      | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x47      | Play  | Client   | 3: add players to team      | Entities            | Array of String (40) | Identifiers for the entities added.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x47      | Play  | Client   | 4: remove players from team | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x47      | Play  | Client   | 4: remove players from team | Entities            | Array of String (40) | Identifiers for the entities removed.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

// ID=0x48
type PlayUpdateScorePkt = internal.PlayUpdateScore_757_1

// ID=0x49
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_754_1

// ID=0x4a
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x4b
type PlayTitlePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name      | Field Type  | Notes                                                                                                                                                                            |
	 * |-----------|-------|----------|--------------------------|-----------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x4B      | Play  | Client   | Action                   | Action          | VarInt Enum |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | Action                   | Field Name      |             |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | 0: set title             | Title Text      | Chat        |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | 1: set subtitle          | Subtitle Text   | Chat        |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | 2: set action bar        | Action bar text | Chat        | Displays a message above the hotbar (the same as position 2 in Chat Message (clientbound), except that it correctly renders formatted chat. See MC-119145 for more information.) |
	 * | 0x4B      | Play  | Client   | 3: set times and display | Fade In         | Int         | Ticks to spend fading in                                                                                                                                                         |
	 * | 0x4B      | Play  | Client   | 3: set times and display | Stay            | Int         | Ticks to keep the title displayed                                                                                                                                                |
	 * | 0x4B      | Play  | Client   | 3: set times and display | Fade Out        | Int         | Ticks to spend out, not when to start fading out                                                                                                                                 |
	 * | 0x4B      | Play  | Client   | 4: hide                  | no fields       | no fields   |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | 5: reset                 | no fields       | no fields   |                                                                                                                                                                                  |
	 * 
	 */
}

// ID=0x4c
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x4d
type PlaySoundEffectPkt = internal.PlaySoundEffect_758_2

// ID=0x4e
type PlayerListHeaderAndFooterPkt = internal.PlayerListHeaderAndFooter_758_0

// ID=0x4f
type PlayCollectItemPkt = internal.PlayCollectItem_758_0

// ID=0x50
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x51
type PlayAdvancementsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name    | Field Type          | Field Type           | Notes                                                      |
	 * |-----------|-------|----------|---------------------|---------------|---------------------|----------------------|------------------------------------------------------------|
	 * | 0x51      | Play  | Client   | Reset/Clear         | Reset/Clear   | Boolean             | Boolean              | Whether to reset/clear the current advancements            |
	 * | 0x51      | Play  | Client   | Mapping size        | Mapping size  | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x51      | Play  | Client   | Advancement mapping | Key           | Array               | Identifier           | The identifier of the advancement                          |
	 * | 0x51      | Play  | Client   | Advancement mapping | Value         | Array               | Advancement          | See below                                                  |
	 * | 0x51      | Play  | Client   | List size           | List size     | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x51      | Play  | Client   | Identifiers         | Identifiers   | Array of Identifier | Array of Identifier  | The identifiers of the advancements that should be removed |
	 * | 0x51      | Play  | Client   | Progress size       | Progress size | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x51      | Play  | Client   | Progress mapping    | Key           | Array               | Identifier           | The identifier of the advancement                          |
	 * | 0x51      | Play  | Client   | Progress mapping    | Value         | Array               | Advancement progress | See below                                                  |
	 * 
	 */
}

// ID=0x52
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x52      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x52      | Play  | Client   | Number Of Properties | Number Of Properties | Int        | Int                    | Number of elements in the following array             |
	 * | 0x52      | Play  | Client   | Property             | Key                  | Array      | String (64)            | See below                                             |
	 * | 0x52      | Play  | Client   | Property             | Value                | Array      | Double                 | See below                                             |
	 * | 0x52      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array             |
	 * | 0x52      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x53
type PlayEntityEffectPkt = internal.PlayEntityEffect_757_2

// ID=0x54
type PlayDeclareRecipesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name  | Field Type | Field Type       | Notes                                                                   |
	 * |-----------|-------|----------|-------------|-------------|------------|------------------|-------------------------------------------------------------------------|
	 * | 0x54      | Play  | Client   | Num Recipes | Num Recipes | VarInt     | VarInt           | Number of elements in the following array                               |
	 * | 0x54      | Play  | Client   | Recipe      | Recipe ID   | Array      | Identifier       |                                                                         |
	 * | 0x54      | Play  | Client   | Recipe      | Type        | Array      | String           | The recipe type, see below                                              |
	 * | 0x54      | Play  | Client   | Recipe      | Data        | Array      | Optional, varies | Additional data for the recipe.  For some types, there will be no data. |
	 * 
	 */
}

// ID=0x55
type PlayTagsPkt struct {
	/* IDs are block IDs */
	BlockTags []Tag // (See below)
	/* IDs are item IDs */
	ItemTags []Tag // (See below)
	/* IDs are fluid IDs */
	FluidTags []Tag // (See below)
}

var _ Packet = (*PlayTagsPkt)(nil)

func (p PlayTagsPkt)Encode(b *PacketBuilder){
	TODO_Encode_Array(p.BlockTags)
	TODO_Encode_Array(p.ItemTags)
	TODO_Encode_Array(p.FluidTags)
}

func (p *PlayTagsPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	TODO_Decode_Array(p.BlockTags)
	TODO_Decode_Array(p.ItemTags)
	TODO_Decode_Array(p.FluidTags)
	return nil
}

// ======== END play ========
