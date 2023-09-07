
// Generated at 2023-09-05 22:06:22.506 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=15346
// Protocol: 498
// Protocol Name: 1.14.4

package packet_1_14_4

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
type LoginEncryptionResponsePkt = internal.LoginEncryptionResponse_763_0

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
type PlaySetDifficultyPkt = internal.PlaySetDifficulty_758_0

// ID=0x3
type PlayChatMessageServerPkt = internal.PlayChatMessageServer_758_0

// ID=0x4
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x5
type PlayClientSettingsPkt = internal.PlayClientSettings_754_2

// ID=0x6
type PlayTabCompleteServerPkt = internal.PlayTabCompleteServer_758_0

// ID=0x7
type PlayWindowConfirmationServerPkt = internal.PlayWindowConfirmationServer_754_0

// ID=0x8
type PlayClickWindowButtonPkt = internal.PlayClickWindowButton_758_0

// ID=0x9
type PlayClickWindowPkt = internal.PlayClickWindow_754_4

// ID=0xa
type PlayCloseWindowServerPkt = internal.PlayCloseWindowServer_758_0

// ID=0xb
type PlayPluginMessageServerPkt = internal.PlayPluginMessageServer_763_0

// ID=0xc
type PlayEditBookPkt = internal.PlayEditBook_755_4

// ID=0xd
type PlayEntityNBTRequestPkt struct {
	/* An incremental ID so that the client can verify that the response matches. */
	TransactionID VarInt // VarInt
	/* The ID of the entity to query. */
	EntityID VarInt // VarInt
}

var _ Packet = (*PlayEntityNBTRequestPkt)(nil)

func (p PlayEntityNBTRequestPkt)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionID)
	b.VarInt(p.EntityID)
}

func (p *PlayEntityNBTRequestPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.TransactionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0xe
type PlayInteractEntityPkt = internal.PlayInteractEntity_578_1

// ID=0xf
type PlayKeepAliveServerPkt = internal.PlayKeepAliveServer_763_0

// ID=0x10
type PlayLockDifficultyPkt = internal.PlayLockDifficulty_763_0

// ID=0x11
type PlayerPositionPkt = internal.PlayerPosition_758_0

// ID=0x12
type PlayerPositionAndRotationPkt = internal.PlayerPositionAndRotation_758_0

// ID=0x13
type PlayerRotationPkt = internal.PlayerRotation_758_0

// ID=0x14
type PlayerMovementPkt = internal.PlayerMovement_758_0

// ID=0x15
type PlayVehicleMoveServerPkt = internal.PlayVehicleMoveServer_758_0

// ID=0x16
type PlaySteerBoatPkt = internal.PlaySteerBoat_758_0

// ID=0x17
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x18
type PlayCraftRecipeRequestPkt = internal.PlayCraftRecipeRequest_758_0

// ID=0x19
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_578_1

// ID=0x1a
type PlayerDiggingPkt = internal.PlayerDigging_758_0

// ID=0x1b
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x1c
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x1d
type PlayRecipeBookDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                    | Field Name                    | Field Type                                                                                                                           | Notes                                                                           |
	 * |-----------|-------|----------|-------------------------------|-------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------|
	 * | 0x1D      | Play  | Server   | Type                          | Type                          | VarInt                                                                                                                               | Determines the format of the rest of the packet                                 |
	 * | 0x1D      | Play  | Server   | Type                          | Field Name                    |                                                                                                                                      |                                                                                 |
	 * | 0x1D      | Play  | Server   | 0: Displayed Recipe           | Recipe ID                     | Identifier                                                                                                                           | A recipe ID                                                                     |
	 * | 0x1D      | Play  | Server   | 1: Recipe Book States         | Crafting Recipe Book Open     | Boolean                                                                                                                              | Whether the player has the crafting recipe book currently opened/active.        |
	 * | 0x1D      | Play  | Server   | 1: Recipe Book States         | Crafting Recipe Filter Active | Boolean                                                                                                                              | Whether the player has the crafting recipe book filter option currently active. |
	 * | 0x1D      | Play  | Server   | 1: Recipe Book States         | Smelting Recipe Book Open     | Boolean                                                                                                                              | Whether the player has the smelting recipe book currently opened/active.        |
	 * | 0x1D      | Play  | Server   | 1: Recipe Book States         | Smelting Recipe Filter Active | Boolean                                                                                                                              | Whether the player has the smelting recipe book filter option currently active. |
	 * | 0x1D      | Play  | Server   | Blasting Recipe Book Open     | Boolean                       | May be swapped with smoking recipe book.  Also, the notchian client appears to use the same value for both of those books currently. |                                                                                 |
	 * | 0x1D      | Play  | Server   | Blasting Recipe Filter Active | Boolean                       | May be swapped with smoking recipe book.  Also, the notchian client appears to use the same value for both of those books currently. |                                                                                 |
	 * | 0x1D      | Play  | Server   | Smoking Recipe Book Open      | Boolean                       |                                                                                                                                      |                                                                                 |
	 * | 0x1D      | Play  | Server   | Smoking Recipe Filter Active  | Boolean                       |                                                                                                                                      |                                                                                 |
	 * 
	 */
}

// ID=0x1e
type PlayNameItemPkt = internal.PlayNameItem_758_0

// ID=0x1f
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_758_0

// ID=0x20
type PlayAdvancementTabPkt = internal.PlayAdvancementTab_758_0

// ID=0x21
type PlaySelectTradePkt = internal.PlaySelectTrade_763_0

// ID=0x22
type PlaySetBeaconEffectPkt = internal.PlaySetBeaconEffect_758_2

// ID=0x23
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChangeServer_758_0

// ID=0x24
type PlayUpdateCommandBlockPkt = internal.PlayUpdateCommandBlock_758_0

// ID=0x25
type PlayUpdateCommandBlockMinecartPkt = internal.PlayUpdateCommandBlockMinecart_758_0

// ID=0x26
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x27
type PlayUpdateJigsawBlockPkt = internal.PlayUpdateJigsawBlock_578_1

// ID=0x28
type PlayUpdateStructureBlockPkt = internal.PlayUpdateStructureBlock_758_0

// ID=0x29
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x2a
type PlayAnimationPkt = internal.PlayAnimation_758_0

// ID=0x2b
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x2c
type PlayerBlockPlacementPkt = internal.PlayerBlockPlacement_758_0

// ID=0x2d
type PlayUseItemPkt = internal.PlayUseItem_758_1

// ---- play: clientbound ----

// ID=0x0
type PlaySpawnObjectPkt struct {
	/* EID of the object */
	EntityID VarInt // VarInt
	ObjectUUID UUID // UUID
	/* The type of object (same as in Spawn Mob) */
	Type VarInt // VarInt
	X Double // Double
	Y Double // Double
	Z Double // Double
	Pitch Angle // Angle
	Yaw Angle // Angle
	/* Meaning dependent on the value of the Type field, see Object Data for details. */
	Data Int // Int
	/* Same units as Entity Velocity.  Always sent, but only used when Data is greater than 0 (except for some entities which always ignore it; see Object Data for details). */
	VelocityX Short // Short
	VelocityY Short // Short
	VelocityZ Short // Short
}

var _ Packet = (*PlaySpawnObjectPkt)(nil)

func (p PlaySpawnObjectPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.ObjectUUID)
	b.VarInt(p.Type)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	p.Pitch.Encode(b)
	p.Yaw.Encode(b)
	b.Int(p.Data)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
}

func (p *PlaySpawnObjectPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ObjectUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return err
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return err
	}
	if p.Data, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.VelocityX, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.VelocityY, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.VelocityZ, ok = r.Short(); !ok {
		return io.EOF
	}
	return nil
}

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
type PlayEntityAnimationPkt = internal.PlayEntityAnimation_763_0

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
type PlayBossBarPkt = internal.PlayBossBar_754_3

// ID=0xd
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_758_0

// ID=0xe
type PlayChatMessageClientPkt = internal.PlayChatMessage_578_4

// ID=0xf
type PlayMultiBlockChangePkt = internal.PlayMultiBlockChange_498_2

// ID=0x10
type PlayTabCompleteClientPkt = internal.PlayTabComplete_498_4

// ID=0x11
type PlayDeclareCommandsPkt = internal.PlayDeclareCommands_758_0

// ID=0x12
type PlayWindowConfirmationClientPkt = internal.PlayWindowConfirmation_754_0

// ID=0x13
type PlayCloseWindowClientPkt = internal.PlayCloseWindow_758_0

// ID=0x14
type PlayWindowItemsPkt = internal.PlayWindowItems_755_1

// ID=0x15
type PlayWindowPropertyPkt = internal.PlayWindowProperty_758_0

// ID=0x16
type PlaySetSlotPkt = internal.PlaySetSlot_755_1

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
type PlayExplosionPkt = internal.PlayExplosion_498_3

// ID=0x1d
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x1e
type PlayChangeGameStatePkt = internal.PlayChangeGameState_758_0

// ID=0x1f
type PlayOpenHorseWindowPkt = internal.PlayOpenHorseWindow_756_1

// ID=0x20
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x21
type PlayChunkDataPkt struct {
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkX Int // Int
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkZ Int // Int
	/* See Chunk Format */
	FullChunk Bool // Boolean
	/* Bitmask with bits set to 1 for every 16×16×16 chunk section whose data is included in Data. The least significant bit represents the chunk section at the bottom of the chunk column (from y=0 to y=15). */
	PrimaryBitMask VarInt // VarInt
	/* Compound containing one long array named MOTION_BLOCKING, which is a heightmap for the highest solid block at each position in the chunk (as a compacted long array with 256 entries at 9 bits per entry totaling 36 longs). The Notchian server also adds a WORLD_SURFACE long array, the purpose of which is unknown, but it's not required for the chunk to be accepted. */
	Heightmaps nbt.NBT // NBT
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
	WriteNBT(b, p.Heightmaps)
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
	if p.Heightmaps, err = nbt.ReadNBT(r); err != nil {
		return err
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

// ID=0x22
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x23
type PlayParticlePkt = internal.PlayParticle_498_2

// ID=0x24
type PlayUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name             | Field Name             | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                           |
	 * |-----------|-------|----------|------------------------|------------------------|------------|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x24      | Play  | Client   | Chunk X                | Chunk X                | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                 |
	 * | 0x24      | Play  | Client   | Chunk Z                | Chunk Z                | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                 |
	 * | 0x24      | Play  | Client   | Sky Light Mask         | Sky Light Mask         | VarInt     | VarInt              | Mask containing 18 bits, with the lowest bit corresponding to chunk section -1 (in the void, y=-16 to y=-1) and the highest bit for chunk section 16 (above the world, y=256 to y=271)                                                                                                                                                                                          |
	 * | 0x24      | Play  | Client   | Block Light Mask       | Block Light Mask       | VarInt     | VarInt              | Mask containing 18 bits, with the same order as sky light                                                                                                                                                                                                                                                                                                                       |
	 * | 0x24      | Play  | Client   | Empty Sky Light Mask   | Empty Sky Light Mask   | VarInt     | VarInt              | Mask containing 18 bits, which indicates sections that have 0 for all their sky light values.  If a section is set in both this mask and the main sky light mask, it is ignored for this mask and it is included in the sky light arrays (the notchian server does not create such masks).  If it is only set in this mask, it is not included in the sky light arrays.         |
	 * | 0x24      | Play  | Client   | Empty Block Light Mask | Empty Block Light Mask | VarInt     | VarInt              | Mask containing 18 bits, which indicates sections that have 0 for all their block light values.  If a section is set in both this mask and the main block light mask, it is ignored for this mask and it is included in the block light arrays (the notchian server does not create such masks).  If it is only set in this mask, it is not included in the block light arrays. |
	 * | 0x24      | Play  | Client   | Sky Light arrays       | Length                 | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                            |
	 * | 0x24      | Play  | Client   | Sky Light arrays       | Sky Light array        | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                  |
	 * | 0x24      | Play  | Client   | Block Light arrays     | Length                 | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                            |
	 * | 0x24      | Play  | Client   | Block Light arrays     | Block Light array      | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                |
	 * 
	 */
}

// ID=0x25
type PlayJoinGamePkt struct {
	/* The player's Entity ID (EID) */
	EntityID Int // Int
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. Bit 3 (0x8) is the hardcore flag. */
	Gamemode UByte // Unsigned Byte
	/* -1: Nether, 0: Overworld, 1: End; also, note that this is not a VarInt but instead a regular int. */
	Dimension Int // Int Enum
	/* Was once used by the client to draw the player list, but now is ignored */
	MaxPlayers UByte // Unsigned Byte
	/* default, flat, largeBiomes, amplified, customized, buffet, default_1_1 */
	LevelType String // String Enum (16)
	/* Render distance (2-32) */
	ViewDistance VarInt // VarInt
	/* If true, a Notchian client shows reduced information on the debug screen.  For servers in development, this should almost always be false. */
	ReducedDebugInfo Bool // Boolean
}

var _ Packet = (*PlayJoinGamePkt)(nil)

func (p PlayJoinGamePkt)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.UByte(p.Gamemode)
	b.Int(p.Dimension)
	b.UByte(p.MaxPlayers)
	b.String(p.LevelType)
	b.VarInt(p.ViewDistance)
	b.Bool(p.ReducedDebugInfo)
}

func (p *PlayJoinGamePkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Dimension, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.MaxPlayers, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.LevelType, ok = r.String(); !ok {
		return io.EOF
	}
	if p.ViewDistance, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ReducedDebugInfo, ok = r.Bool(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x26
type PlayMapDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type                      | Field Type                      | Notes                                                                                                      |
	 * |-----------|-------|----------|-------------------|-------------------|---------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x26      | Play  | Client   | Map ID            | Map ID            | VarInt                          | VarInt                          | Map ID of the map being modified                                                                           |
	 * | 0x26      | Play  | Client   | Scale             | Scale             | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x26      | Play  | Client   | Tracking Position | Tracking Position | Boolean                         | Boolean                         | Specifies whether player and item frame icons are shown                                                    |
	 * | 0x26      | Play  | Client   | Locked            | Locked            | Boolean                         | Boolean                         | True if the map has been locked in a cartography table                                                     |
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
type PlayTradeListPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name                   | Field Type | Field Type    | Notes                                                                                                                                                                              |
	 * |-----------|-------|----------|---------------------|------------------------------|------------|---------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x27      | Play  | Client   |                     |                              |            |               |                                                                                                                                                                                    |
	 * | 0x27      | Play  | Client   | Window ID           | Window ID                    | VarInt     | VarInt        | The ID of the window that is open; this is an int rather than a byte.                                                                                                              |
	 * | 0x27      | Play  | Client   | Size                | Size                         | Byte       | Byte          | The number of trades in the following array                                                                                                                                        |
	 * | 0x27      | Play  | Client   | Trades              | Input item 1                 | Array      | Slot          | The first item the villager is buying                                                                                                                                              |
	 * | 0x27      | Play  | Client   | Trades              | Output item                  | Array      | Slot          | The item the villager is selling                                                                                                                                                   |
	 * | 0x27      | Play  | Client   | Trades              | Has second item              | Array      | Boolean       | Whether there is a second item                                                                                                                                                     |
	 * | 0x27      | Play  | Client   | Trades              | Input item 2                 | Array      | Optional Slot | The second item the villager is buying; only present if they have a second item.                                                                                                   |
	 * | 0x27      | Play  | Client   | Trades              | Trade disabled               | Array      | Boolean       | True if the trade is disabled; false if the trade is enabled.                                                                                                                      |
	 * | 0x27      | Play  | Client   | Trades              | Number of trade uses         | Array      | Integer       | Number of times the trade has been used so far                                                                                                                                     |
	 * | 0x27      | Play  | Client   | Trades              | Maximum number of trade uses | Array      | Integer       | Number of times this trade can be used                                                                                                                                             |
	 * | 0x27      | Play  | Client   | Trades              | XP                           | Array      | Integer       |                                                                                                                                                                                    |
	 * | 0x27      | Play  | Client   | Trades              | Special Price                | Array      | Integer       |                                                                                                                                                                                    |
	 * | 0x27      | Play  | Client   | Trades              | Price Multiplier             | Array      | Float         |                                                                                                                                                                                    |
	 * | 0x27      | Play  | Client   | Trades              | Demand                       | Array      | Integer       |                                                                                                                                                                                    |
	 * | 0x27      | Play  | Client   | Villager level      | Villager level               | VarInt     | VarInt        | Appears on the trade GUI; meaning comes from the translation key merchant.level. + level.
	 * 1: Novice, 2: Apprentice, 3: Journeyman, 4: Expert, 5: Master                            |
	 * | 0x27      | Play  | Client   | Experience          | Experience                   | VarInt     | VarInt        | Total experience for this villager (always 0 for the wandering trader)                                                                                                             |
	 * | 0x27      | Play  | Client   | Is regular villager | Is regular villager          | Boolean    | Boolean       | True if this is a regular villager; false for the wandering trader.  When false, hides the villager level and some other GUI elements.                                             |
	 * | 0x27      | Play  | Client   | Can restock         | Can restock                  | Boolean    | Boolean       | True for regular villagers and false for the wandering trader.  If true, the "Villagers restock up to two times per day." message is displayed when hovering over disabled trades. |
	 * 
	 */
}

// ID=0x28
type PlayEntityPositionPkt = internal.PlayEntityPosition_758_0

// ID=0x29
type PlayEntityPositionAndRotationPkt = internal.PlayEntityPositionAndRotation_758_0

// ID=0x2a
type PlayEntityRotationPkt = internal.PlayEntityRotation_758_0

// ID=0x2b
type PlayEntityMovementPkt = internal.PlayEntityMovement_754_0

// ID=0x2c
type PlayVehicleMoveClientPkt = internal.PlayVehicleMove_758_0

// ID=0x2d
type PlayOpenBookPkt = internal.PlayOpenBook_763_0

// ID=0x2e
type PlayOpenWindowPkt = internal.PlayOpenWindow_758_0

// ID=0x2f
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x30
type PlayCraftRecipeResponsePkt = internal.PlayCraftRecipeResponse_758_0

// ID=0x31
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x32
type PlayCombatEventPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name      | Field Name | Field Type  | Notes                                         |
	 * |-----------|-------|----------|-----------------|------------|-------------|-----------------------------------------------|
	 * | 0x32      | Play  | Client   | Event           | Event      | VarInt Enum | Determines the layout of the remaining packet |
	 * | 0x32      | Play  | Client   | Event           | Field Name |             |                                               |
	 * | 0x32      | Play  | Client   | 0: enter combat | no fields  | no fields   |                                               |
	 * | 0x32      | Play  | Client   | 1: end combat   | Duration   | VarInt      |                                               |
	 * | 0x32      | Play  | Client   | 1: end combat   | Entity ID  | Int         |                                               |
	 * | 0x32      | Play  | Client   | 2: entity dead  | Player ID  | VarInt      |                                               |
	 * | 0x32      | Play  | Client   | 2: entity dead  | Entity ID  | Int         |                                               |
	 * | 0x32      | Play  | Client   | 2: entity dead  | Message    | Chat        |                                               |
	 * 
	 */
}

// ID=0x33
type PlayerInfoPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type              | Notes                                                   |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-------------------------|---------------------------------------------------------|
	 * | 0x33      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt                  | Determines the rest of the Player format after the UUID |
	 * | 0x33      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x33      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID                    |                                                         |
	 * | 0x33      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                         |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String (16)   | String (16)             |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String (32767)          |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String (32767)          |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean                 |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String (32767) | Only if Is Signed is true                               |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds                                |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only if Has Display Name is true                        |
	 * | 0x33      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds                                |
	 * | 0x33      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x33      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true                   |
	 * | 0x33      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields               |                                                         |
	 * 
	 */
}

// ID=0x34
type PlayFacePlayerPkt = internal.PlayFacePlayer_757_0

// ID=0x35
type PlayerPositionAndLookPkt = internal.PlayerPositionAndLook_754_1

// ID=0x36
type PlayUnlockRecipesPkt = internal.PlayUnlockRecipes_578_2

// ID=0x37
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x38
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x39
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x3a
type PlayRespawnPkt struct {
	/* -1: The Nether, 0: The Overworld, 1: The End */
	Dimension Int // Int Enum
	/* 0: survival, 1: creative, 2: adventure, 3: spectator. The hardcore flag is not included */
	Gamemode UByte // Unsigned Byte
	/* Same as Join Game */
	LevelType String // String (16)
}

var _ Packet = (*PlayRespawnPkt)(nil)

func (p PlayRespawnPkt)Encode(b *PacketBuilder){
	b.Int(p.Dimension)
	b.UByte(p.Gamemode)
	b.String(p.LevelType)
}

func (p *PlayRespawnPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.Dimension, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.LevelType, ok = r.String(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x3b
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x3c
type PlaySelectAdvancementTabPkt = internal.PlaySelectAdvancementTab_758_0

// ID=0x3d
type PlayWorldBorderPkt = internal.PlayWorldBorder_753_1

// ID=0x3e
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x3f
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_758_0

// ID=0x40
type PlayUpdateViewPositionPkt = internal.PlayUpdateViewPosition_758_0

// ID=0x41
type PlayUpdateViewDistancePkt = internal.PlayUpdateViewDistance_758_0

// ID=0x42
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x43
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x44
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x45
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x46
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_578_2

// ID=0x47
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x48
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x49
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_758_0

// ID=0x4a
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x4b
type PlayTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                  | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|-----------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x4B      | Play  | Client   | Team Name                   | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x4B      | Play  | Client   | Mode                        | Mode                | Byte                 | Determines the layout of the remaining packet                                                                          |
	 * | 0x4B      | Play  | Client   | 0: create team              | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x4B      | Play  | Client   | 0: create team              | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team                                      |
	 * | 0x4B      | Play  | Client   | 0: create team              | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x4B      | Play  | Client   | 0: create team              | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x4B      | Play  | Client   | 0: create team              | Team Color          | VarInt enum          | Used to color the name of players on the team; see below                                                               |
	 * | 0x4B      | Play  | Client   | 0: create team              | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team                                                       |
	 * | 0x4B      | Play  | Client   | 0: create team              | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team                                                        |
	 * | 0x4B      | Play  | Client   | 0: create team              | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x4B      | Play  | Client   | 0: create team              | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x4B      | Play  | Client   | 1: remove team              | no fields           | no fields            |                                                                                                                        |
	 * | 0x4B      | Play  | Client   | 2: update team info         | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x4B      | Play  | Client   | 2: update team info         | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team                                     |
	 * | 0x4B      | Play  | Client   | 2: update team info         | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x4B      | Play  | Client   | 2: update team info         | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x4B      | Play  | Client   | 2: update team info         | Team Color          | VarInt enum          | Used to color the name of players on the team; see below                                                               |
	 * | 0x4B      | Play  | Client   | 2: update team info         | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team                                                       |
	 * | 0x4B      | Play  | Client   | 2: update team info         | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team                                                        |
	 * | 0x4B      | Play  | Client   | 3: add players to team      | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x4B      | Play  | Client   | 3: add players to team      | Entities            | Array of String (40) | Identifiers for the entities added.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x4B      | Play  | Client   | 4: remove players from team | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x4B      | Play  | Client   | 4: remove players from team | Entities            | Array of String (40) | Identifiers for the entities removed.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

// ID=0x4c
type PlayUpdateScorePkt = internal.PlayUpdateScore_757_1

// ID=0x4d
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_754_1

// ID=0x4e
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x4f
type PlayTitlePkt = internal.PlayTitle_754_0

// ID=0x50
type PlayEntitySoundEffectPkt = internal.PlayEntitySoundEffect_760_1

// ID=0x51
type PlaySoundEffectPkt = internal.PlaySoundEffect_758_2

// ID=0x52
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x53
type PlayerListHeaderAndFooterPkt = internal.PlayerListHeaderAndFooter_758_0

// ID=0x54
type PlayNBTQueryResponsePkt = internal.PlayNBTQueryResponse_758_0

// ID=0x55
type PlayCollectItemPkt = internal.PlayCollectItem_758_0

// ID=0x56
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x57
type PlayAdvancementsPkt = internal.PlayAdvancements_754_2

// ID=0x58
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x58      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x58      | Play  | Client   | Number Of Properties | Number Of Properties | Int        | Int                    | Number of elements in the following array             |
	 * | 0x58      | Play  | Client   | Property             | Key                  | Array      | String (64)            | See below                                             |
	 * | 0x58      | Play  | Client   | Property             | Value                | Array      | Double                 | See below                                             |
	 * | 0x58      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array             |
	 * | 0x58      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x59
type PlayEntityEffectPkt = internal.PlayEntityEffect_757_2

// ID=0x5a
type PlayDeclareRecipesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name  | Field Type | Field Type       | Notes                                                                   |
	 * |-----------|-------|----------|-------------|-------------|------------|------------------|-------------------------------------------------------------------------|
	 * | 0x5A      | Play  | Client   | Num Recipes | Num Recipes | VarInt     | VarInt           | Number of elements in the following array                               |
	 * | 0x5A      | Play  | Client   | Recipe      | Recipe ID   | Array      | Identifier       |                                                                         |
	 * | 0x5A      | Play  | Client   | Recipe      | Type        | Array      | String           | The recipe type, see below                                              |
	 * | 0x5A      | Play  | Client   | Recipe      | Data        | Array      | Optional, varies | Additional data for the recipe.  For some types, there will be no data. |
	 * 
	 */
}

// ID=0x5b
type PlayTagsPkt = internal.PlayTags_754_2

// ID=0x5c
type PlayAcknowledgePlayerDiggingPkt = internal.PlayAcknowledgePlayerDigging_758_0

// ======== END play ========
