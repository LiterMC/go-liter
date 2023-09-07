
// Generated at 2023-09-05 22:06:22.506 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=18242
// Protocol: 762
// Protocol Name: 1.19.4

package packet_1_19_4

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
type LoginStartPkt = internal.LoginStart_763_0

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
type LoginSuccessPkt = internal.LoginSuccess_763_0

// ID=0x3
type LoginSetCompressionPkt = internal.LoginSetCompression_763_0

// ID=0x4
type LoginPluginRequestPkt = internal.LoginPluginRequest_763_0

// ======== END login ========

// ======== BEGIN play ========
// ---- play: serverbound ----

// ID=0x0
type PlayConfirmTeleportationPkt = internal.PlayConfirmTeleportation_763_0

// ID=0x1
type PlayQueryBlockEntityTagPkt = internal.PlayQueryBlockEntityTag_763_0

// ID=0x2
type PlayChangeDifficultyServerPkt = internal.PlayChangeDifficultyServer_763_0

// ID=0x3
type PlayMessageAcknowledgmentPkt = internal.PlayMessageAcknowledgment_763_0

// ID=0x4
type PlayChatCommandPkt = internal.PlayChatCommand_763_0

// ID=0x5
type PlayChatMessagePkt = internal.PlayChatMessage_763_0

// ID=0x6
type PlayerSessionPkt = internal.PlayerSession_763_0

// ID=0x7
type PlayClientCommandPkt = internal.PlayClientCommand_763_0

// ID=0x8
type PlayClientInformationPkt = internal.PlayClientInformation_763_0

// ID=0x9
type PlayCommandSuggestionsRequestPkt = internal.PlayCommandSuggestionsRequest_763_0

// ID=0xa
type PlayClickContainerButtonPkt = internal.PlayClickContainerButton_763_0

// ID=0xb
type PlayClickContainerPkt = internal.PlayClickContainer_763_0

// ID=0xc
type PlayCloseContainerServerPkt = internal.PlayCloseContainerServer_763_0

// ID=0xd
type PlayPluginMessageServerPkt = internal.PlayPluginMessageServer_763_0

// ID=0xe
type PlayEditBookPkt = internal.PlayEditBook_763_0

// ID=0xf
type PlayQueryEntityTagPkt = internal.PlayQueryEntityTag_763_0

// ID=0x10
type PlayInteractPkt = internal.PlayInteract_763_0

// ID=0x11
type PlayJigsawGeneratePkt = internal.PlayJigsawGenerate_763_0

// ID=0x12
type PlayKeepAliveServerPkt = internal.PlayKeepAliveServer_763_0

// ID=0x13
type PlayLockDifficultyPkt = internal.PlayLockDifficulty_763_0

// ID=0x14
type PlaySetPlayerPositionPkt = internal.PlaySetPlayerPosition_763_0

// ID=0x15
type PlaySetPlayerPositionAndRotationPkt = internal.PlaySetPlayerPositionAndRotation_763_0

// ID=0x16
type PlaySetPlayerRotationPkt = internal.PlaySetPlayerRotation_763_0

// ID=0x17
type PlaySetPlayerOnGroundPkt = internal.PlaySetPlayerOnGround_763_0

// ID=0x18
type PlayMoveVehicleServerPkt = internal.PlayMoveVehicleServer_763_0

// ID=0x19
type PlayPaddleBoatPkt = internal.PlayPaddleBoat_763_0

// ID=0x1a
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x1b
type PlayPlaceRecipePkt = internal.PlayPlaceRecipe_763_0

// ID=0x1c
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_763_0

// ID=0x1d
type PlayerActionPkt = internal.PlayerAction_763_0

// ID=0x1e
type PlayerCommandPkt = internal.PlayerCommand_763_0

// ID=0x1f
type PlayerInputPkt = internal.PlayerInput_763_0

// ID=0x20
type PlayPongPkt = internal.PlayPong_763_0

// ID=0x21
type PlayChangeRecipeBookSettingsPkt = internal.PlayChangeRecipeBookSettings_763_0

// ID=0x22
type PlaySetSeenRecipePkt = internal.PlaySetSeenRecipe_763_0

// ID=0x23
type PlayRenameItemPkt = internal.PlayRenameItem_763_0

// ID=0x24
type PlayResourcePackServerPkt = internal.PlayResourcePackServer_763_0

// ID=0x25
type PlaySeenAdvancementsPkt = internal.PlaySeenAdvancements_763_0

// ID=0x26
type PlaySelectTradePkt = internal.PlaySelectTrade_763_0

// ID=0x27
type PlaySetBeaconEffectPkt = internal.PlaySetBeaconEffect_763_0

// ID=0x28
type PlaySetHeldItemServerPkt = internal.PlaySetHeldItemServer_763_0

// ID=0x29
type PlayProgramCommandBlockPkt = internal.PlayProgramCommandBlock_763_0

// ID=0x2a
type PlayProgramCommandBlockMinecartPkt = internal.PlayProgramCommandBlockMinecart_763_0

// ID=0x2b
type PlaySetCreativeModeSlotPkt = internal.PlaySetCreativeModeSlot_763_0

// ID=0x2c
type PlayProgramJigsawBlockPkt = internal.PlayProgramJigsawBlock_763_0

// ID=0x2d
type PlayProgramStructureBlockPkt = internal.PlayProgramStructureBlock_763_0

// ID=0x2e
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x2f
type PlaySwingArmPkt = internal.PlaySwingArm_763_0

// ID=0x30
type PlayTeleportToEntityPkt = internal.PlayTeleportToEntity_763_0

// ID=0x31
type PlayUseItemOnPkt = internal.PlayUseItemOn_763_0

// ID=0x32
type PlayUseItemPkt = internal.PlayUseItem_763_0

// ---- play: clientbound ----

// ID=0x0
type PlayBundleDelimiterPkt = internal.PlayBundleDelimiter_763_0

// ID=0x1
type PlaySpawnEntityPkt = internal.PlaySpawnEntity_763_0

// ID=0x2
type PlaySpawnExperienceOrbPkt = internal.PlaySpawnExperienceOrb_763_0

// ID=0x3
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_763_0

// ID=0x4
type PlayEntityAnimationPkt = internal.PlayEntityAnimation_763_0

// ID=0x5
type PlayAwardStatisticsPkt = internal.PlayAwardStatistics_763_0

// ID=0x6
type PlayAcknowledgeBlockChangePkt = internal.PlayAcknowledgeBlockChange_763_0

// ID=0x7
type PlaySetBlockDestroyStagePkt = internal.PlaySetBlockDestroyStage_763_0

// ID=0x8
type PlayBlockEntityDataPkt = internal.PlayBlockEntityData_763_0

// ID=0x9
type PlayBlockActionPkt = internal.PlayBlockAction_763_0

// ID=0xa
type PlayBlockUpdatePkt = internal.PlayBlockUpdate_763_0

// ID=0xb
type PlayBossBarPkt = internal.PlayBossBar_763_0

// ID=0xc
type PlayChangeDifficultyClientPkt = internal.PlayChangeDifficulty_763_0

// ID=0xd
type PlayChunkBiomesPkt = internal.PlayChunkBiomes_763_0

// ID=0xe
type PlayClearTitlesPkt = internal.PlayClearTitles_763_0

// ID=0xf
type PlayCommandSuggestionsResponsePkt = internal.PlayCommandSuggestionsResponse_763_0

// ID=0x10
type PlayCommandsPkt = internal.PlayCommands_763_0

// ID=0x11
type PlayCloseContainerClientPkt = internal.PlayCloseContainer_763_0

// ID=0x12
type PlaySetContainerContentPkt = internal.PlaySetContainerContent_763_0

// ID=0x13
type PlaySetContainerPropertyPkt = internal.PlaySetContainerProperty_763_0

// ID=0x14
type PlaySetContainerSlotPkt = internal.PlaySetContainerSlot_763_0

// ID=0x15
type PlaySetCooldownPkt = internal.PlaySetCooldown_763_0

// ID=0x16
type PlayChatSuggestionsPkt = internal.PlayChatSuggestions_763_0

// ID=0x17
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_763_0

// ID=0x18
type PlayDamageEventPkt = internal.PlayDamageEvent_763_0

// ID=0x19
type PlayDeleteMessagePkt = internal.PlayDeleteMessage_763_0

// ID=0x1a
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x1b
type PlayDisguisedChatMessagePkt = internal.PlayDisguisedChatMessage_763_0

// ID=0x1c
type PlayEntityEventPkt = internal.PlayEntityEvent_763_0

// ID=0x1d
type PlayExplosionPkt = internal.PlayExplosion_763_0

// ID=0x1e
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x1f
type PlayGameEventPkt = internal.PlayGameEvent_763_0

// ID=0x20
type PlayOpenHorseScreenPkt = internal.PlayOpenHorseScreen_763_0

// ID=0x21
type PlayHurtAnimationPkt = internal.PlayHurtAnimation_763_0

// ID=0x22
type PlayInitializeWorldBorderPkt = internal.PlayInitializeWorldBorder_763_0

// ID=0x23
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x24
type PlayChunkDataAndUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name               | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
	 * |-----------|-------|----------|--------------------------|--------------------------|------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x24      | Play  | Client   | Chunk X                  | Chunk X                  | Int        | Int                 | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x24      | Play  | Client   | Chunk Z                  | Chunk Z                  | Int        | Int                 | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x24      | Play  | Client   | Heightmaps               | Heightmaps               | NBT        | NBT                 | Compound containing one long array named MOTION_BLOCKING, which is a heightmap for the highest solid block at each position in the chunk (as a compacted long array with 256 entries, with the number of bits per entry varying depending on the world's height, defined by the formula ceil(log2(height + 1))). The Notchian server also adds a WORLD_SURFACE long array, the purpose of which is unknown, but it's not required for the chunk to be accepted. |
	 * | 0x24      | Play  | Client   | Size                     | Size                     | VarInt     | VarInt              | Size of Data in bytes                                                                                                                                                                                                                                                                                                                                                                                                                                           |
	 * | 0x24      | Play  | Client   | Data                     | Data                     | Byte array | Byte array          | See data structure in Chunk Format                                                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x24      | Play  | Client   | Number of block entities | Number of block entities | VarInt     | VarInt              | Number of elements in the following array                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x24      | Play  | Client   | Block Entity             | Packed XZ                | Array      | Byte                | The packed section coordinates, calculated from ((blockX & 15) << 4) | (blockZ & 15)                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x24      | Play  | Client   | Block Entity             | Y                        | Array      | Short               | The height relative to the world                                                                                                                                                                                                                                                                                                                                                                                                                                |
	 * | 0x24      | Play  | Client   | Block Entity             | Type                     | Array      | VarInt              | The type of block entity                                                                                                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x24      | Play  | Client   | Block Entity             | Data                     | Array      | NBT                 | The block entity's data, without the X, Y, and Z values                                                                                                                                                                                                                                                                                                                                                                                                         |
	 * | 0x24      | Play  | Client   | Trust Edges              | Trust Edges              | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x24      | Play  | Client   | Sky Light Mask           | Sky Light Mask           | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Sky Light array below.  The least significant bit is for blocks 16 blocks to 1 block below the min world height (one section below the world), while the most significant bit covers blocks 1 to 16 blocks above the max world height (one section above the world).                                          |
	 * | 0x24      | Play  | Client   | Block Light Mask         | Block Light Mask         | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Block Light array below.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                 |
	 * | 0x24      | Play  | Client   | Empty Sky Light Mask     | Empty Sky Light Mask     | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Sky Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                    |
	 * | 0x24      | Play  | Client   | Empty Block Light Mask   | Empty Block Light Mask   | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Block Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                  |
	 * | 0x24      | Play  | Client   | Sky Light array count    | Sky Light array count    | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Sky Light Mask                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x24      | Play  | Client   | Sky Light arrays         | Length                   | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x24      | Play  | Client   | Sky Light arrays         | Sky Light array          | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value. Indexed ((y<<8) | (z<<4) | x) / 2  If there's a remainder, masked 0xF0 else 0x0F.                                                                                                                                                                                                                                                |
	 * | 0x24      | Play  | Client   | Block Light array count  | Block Light array count  | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Block Light Mask                                                                                                                                                                                                                                                                                                                                                               |
	 * | 0x24      | Play  | Client   | Block Light arrays       | Length                   | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x24      | Play  | Client   | Block Light arrays       | Block Light array        | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value. Indexed ((y<<8) | (z<<4) | x) / 2  If there's a remainder, masked 0xF0 else 0x0F.                                                                                                                                                                                                                                              |
	 * 
	 */
}

// ID=0x25
type PlayWorldEventPkt = internal.PlayWorldEvent_763_0

// ID=0x26
type PlayParticlePkt = internal.PlayParticle_763_0

// ID=0x27
type PlayUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name              | Field Name              | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                  |
	 * |-----------|-------|----------|-------------------------|-------------------------|------------|---------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x27      | Play  | Client   | Chunk X                 | Chunk X                 | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x27      | Play  | Client   | Chunk Z                 | Chunk Z                 | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x27      | Play  | Client   | Trust Edges             | Trust Edges             | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                                                          |
	 * | 0x27      | Play  | Client   | Sky Light Mask          | Sky Light Mask          | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Sky Light array below.  The least significant bit is for blocks 16 blocks to 1 block below the min world height (one section below the world), while the most significant bit covers blocks 1 to 16 blocks above the max world height (one section above the world). |
	 * | 0x27      | Play  | Client   | Block Light Mask        | Block Light Mask        | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Block Light array below.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                        |
	 * | 0x27      | Play  | Client   | Empty Sky Light Mask    | Empty Sky Light Mask    | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Sky Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                           |
	 * | 0x27      | Play  | Client   | Empty Block Light Mask  | Empty Block Light Mask  | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Block Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                         |
	 * | 0x27      | Play  | Client   | Sky Light array count   | Sky Light array count   | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Sky Light Mask                                                                                                                                                                                                                                                                                                                        |
	 * | 0x27      | Play  | Client   | Sky Light arrays        | Length                  | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x27      | Play  | Client   | Sky Light arrays        | Sky Light array         | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                                                         |
	 * | 0x27      | Play  | Client   | Block Light array count | Block Light array count | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Block Light Mask                                                                                                                                                                                                                                                                                                                      |
	 * | 0x27      | Play  | Client   | Block Light arrays      | Length                  | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x27      | Play  | Client   | Block Light arrays      | Block Light array       | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                                                       |
	 * 
	 */
}

// ID=0x28
type PlayLoginPkt struct {
	/* The player's Entity ID (EID). */
	EntityID Int // Int
	IsHardcore Bool // Boolean
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. */
	GameMode UByte // Unsigned Byte
	/* -1: Undefined (null), 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. The previous game mode. Vanilla client uses this for the debug (F3 + N & F3 + F4) game mode switch. (More information needed) */
	PreviousGameMode Byte // Byte
	/* Size of the following array. */
	DimensionCount VarInt // VarInt
	/* Identifiers for all dimensions on the server. */
	DimensionNames []String // Array of Identifier
	/* Represents certain registries that are sent from the server and are applied on the client. */
	RegistryCodec *nbt.NBTCompound // NBT Tag Compound
	/* Name of the dimension type being spawned into. */
	DimensionType String // Identifier
	/* Name of the dimension being spawned into. */
	DimensionName String // Identifier
	/* First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise */
	HashedSeed Long // Long
	/* Was once used by the client to draw the player list, but now is ignored. */
	MaxPlayers VarInt // VarInt
	/* Render distance (2-32). */
	ViewDistance VarInt // VarInt
	/* The distance that the client will process specific things, such as entities. */
	SimulationDistance VarInt // VarInt
	/* If true, a Notchian client shows reduced information on the debug screen.  For servers in development, this should almost always be false. */
	ReducedDebugInfo Bool // Boolean
	/* Set to false when the doImmediateRespawn gamerule is true. */
	EnableRespawnScreen Bool // Boolean
	/* True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks. */
	IsDebug Bool // Boolean
	/* True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63. */
	IsFlat Bool // Boolean
	/* If true, then the next two fields are present. */
	HasDeathLocation Bool // Boolean
	/* Name of the dimension the player died in. */
	DeathDimensionName Optional[String] // Optional Identifier
	/* The location that the player died at. */
	DeathLocation Optional[Position] // Optional Position
}

var _ Packet = (*PlayLoginPkt)(nil)

func (p PlayLoginPkt)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.Bool(p.IsHardcore)
	b.UByte(p.GameMode)
	b.Byte(p.PreviousGameMode)
	p.DimensionCount = (VarInt)(len(p.DimensionNames))
	b.VarInt(p.DimensionCount)
	for _, v := range p.DimensionNames {
		b.String(v)
	}
	p.RegistryCodec.Encode(b)
	b.String(p.DimensionType)
	b.String(p.DimensionName)
	b.Long(p.HashedSeed)
	b.VarInt(p.MaxPlayers)
	b.VarInt(p.ViewDistance)
	b.VarInt(p.SimulationDistance)
	b.Bool(p.ReducedDebugInfo)
	b.Bool(p.EnableRespawnScreen)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
	b.Bool(p.HasDeathLocation)
	if p.DeathDimensionName.Ok = TODO; p.DeathDimensionName.Ok {
		b.String(p.DeathDimensionName.V)
	}
	if p.DeathLocation.Ok = TODO; p.DeathLocation.Ok {
		p.DeathLocation.V.Encode(b)
	}
}

func (p *PlayLoginPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.IsHardcore, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.GameMode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.PreviousGameMode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.DimensionCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.DimensionNames = make([]String, p.DimensionCount)
	for i, _ := range p.DimensionNames {
		if p.DimensionNames[i], ok = r.String(); !ok {
			return io.EOF
		}
	}
	p.RegistryCodec = new(nbt.NBTCompound)
	if err = p.RegistryCodec.DecodeFrom(r); err != nil {
		return err
	}
	if p.DimensionType, ok = r.String(); !ok {
		return io.EOF
	}
	if p.DimensionName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.HashedSeed, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.MaxPlayers, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ViewDistance, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.SimulationDistance, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ReducedDebugInfo, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.EnableRespawnScreen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.IsDebug, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.IsFlat, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.HasDeathLocation, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.DeathDimensionName.Ok = TODO; p.DeathDimensionName.Ok {
		if p.DeathDimensionName.V, ok = r.String(); !ok {
			return io.EOF
		}
	}
	if p.DeathLocation.Ok = TODO; p.DeathLocation.Ok {
		if err = p.DeathLocation.V.DecodeFrom(r); err != nil {
			return err
		}
	}
	return nil
}

// ID=0x29
type PlayMapDataPkt = internal.PlayMapData_763_0

// ID=0x2a
type PlayMerchantOffersPkt = internal.PlayMerchantOffers_763_0

// ID=0x2b
type PlayUpdateEntityPositionPkt = internal.PlayUpdateEntityPosition_763_0

// ID=0x2c
type PlayUpdateEntityPositionAndRotationPkt = internal.PlayUpdateEntityPositionAndRotation_763_0

// ID=0x2d
type PlayUpdateEntityRotationPkt = internal.PlayUpdateEntityRotation_763_0

// ID=0x2e
type PlayMoveVehicleClientPkt = internal.PlayMoveVehicle_763_0

// ID=0x2f
type PlayOpenBookPkt = internal.PlayOpenBook_763_0

// ID=0x30
type PlayOpenScreenPkt = internal.PlayOpenScreen_763_0

// ID=0x31
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x32
type PlayPingPkt = internal.PlayPing_763_0

// ID=0x33
type PlayPlaceGhostRecipePkt = internal.PlayPlaceGhostRecipe_763_0

// ID=0x34
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x35
type PlayerChatMessagePkt = internal.PlayerChatMessage_763_0

// ID=0x36
type PlayEndCombatPkt = internal.PlayEndCombat_762_1

// ID=0x37
type PlayEnterCombatPkt = internal.PlayEnterCombat_763_0

// ID=0x38
type PlayCombatDeathPkt = internal.PlayCombatDeath_762_1

// ID=0x39
type PlayerInfoRemovePkt = internal.PlayerInfoRemove_763_0

// ID=0x3a
type PlayerInfoUpdatePkt = internal.PlayerInfoUpdate_763_0

// ID=0x3b
type PlayLookAtPkt = internal.PlayLookAt_763_0

// ID=0x3c
type PlaySynchronizePlayerPositionPkt = internal.PlaySynchronizePlayerPosition_763_0

// ID=0x3d
type PlayUpdateRecipeBookPkt = internal.PlayUpdateRecipeBook_763_0

// ID=0x3e
type PlayRemoveEntitiesPkt = internal.PlayRemoveEntities_763_0

// ID=0x3f
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_763_0

// ID=0x40
type PlayResourcePackClientPkt = internal.PlayResourcePack_763_0

// ID=0x41
type PlayRespawnPkt struct {
	/* Valid dimensions are defined per dimension registry sent in Login (play) */
	DimensionType String // Identifier
	/* Name of the dimension being spawned into. */
	DimensionName String // Identifier
	/* First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise */
	HashedSeed Long // Long
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. */
	GameMode UByte // Unsigned Byte
	/* -1: Undefined (null), 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. The previous game mode. Vanilla client uses this for the debug (F3 + N & F3 + F4) game mode switch. (More information needed) */
	PreviousGameMode Byte // Byte
	/* True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks. */
	IsDebug Bool // Boolean
	/* True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63. */
	IsFlat Bool // Boolean
	/* Bit mask. 0x01: Keep attributes, 0x02: Keep metadata. Tells which data should be kept on the client side once the player has respawned.
	 * In the Notchian implementation, this is context dependent:
	 * 
	 * normal respawns (after death) keep no data;
	 * exiting the end poem/credits keeps the attributes;
	 * other dimension changes (portals or teleports) keep all data. */
	DataKept Byte // Byte
	/* If true, then the next two fields are present. */
	HasDeathLocation Bool // Boolean
	/* Name of the dimension the player died in. */
	DeathDimensionName Optional[String] // Optional Identifier
	/* The location that the player died at. */
	DeathLocation Optional[Position] // Optional Position
}

var _ Packet = (*PlayRespawnPkt)(nil)

func (p PlayRespawnPkt)Encode(b *PacketBuilder){
	b.String(p.DimensionType)
	b.String(p.DimensionName)
	b.Long(p.HashedSeed)
	b.UByte(p.GameMode)
	b.Byte(p.PreviousGameMode)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
	b.Byte(p.DataKept)
	b.Bool(p.HasDeathLocation)
	if p.DeathDimensionName.Ok = TODO; p.DeathDimensionName.Ok {
		b.String(p.DeathDimensionName.V)
	}
	if p.DeathLocation.Ok = TODO; p.DeathLocation.Ok {
		p.DeathLocation.V.Encode(b)
	}
}

func (p *PlayRespawnPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.DimensionType, ok = r.String(); !ok {
		return io.EOF
	}
	if p.DimensionName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.HashedSeed, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.GameMode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.PreviousGameMode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.IsDebug, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.IsFlat, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.DataKept, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.HasDeathLocation, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.DeathDimensionName.Ok = TODO; p.DeathDimensionName.Ok {
		if p.DeathDimensionName.V, ok = r.String(); !ok {
			return io.EOF
		}
	}
	if p.DeathLocation.Ok = TODO; p.DeathLocation.Ok {
		if err = p.DeathLocation.V.DecodeFrom(r); err != nil {
			return err
		}
	}
	return nil
}

// ID=0x42
type PlaySetHeadRotationPkt = internal.PlaySetHeadRotation_763_0

// ID=0x43
type PlayUpdateSectionBlocksPkt = internal.PlayUpdateSectionBlocks_762_1

// ID=0x44
type PlaySelectAdvancementsTabPkt = internal.PlaySelectAdvancementsTab_763_0

// ID=0x45
type PlayServerDataPkt = internal.PlayServerData_763_0

// ID=0x46
type PlaySetActionBarTextPkt = internal.PlaySetActionBarText_763_0

// ID=0x47
type PlaySetBorderCenterPkt = internal.PlaySetBorderCenter_763_0

// ID=0x48
type PlaySetBorderLerpSizePkt = internal.PlaySetBorderLerpSize_763_0

// ID=0x49
type PlaySetBorderSizePkt = internal.PlaySetBorderSize_763_0

// ID=0x4a
type PlaySetBorderWarningDelayPkt = internal.PlaySetBorderWarningDelay_763_0

// ID=0x4b
type PlaySetBorderWarningDistancePkt = internal.PlaySetBorderWarningDistance_763_0

// ID=0x4c
type PlaySetCameraPkt = internal.PlaySetCamera_763_0

// ID=0x4d
type PlaySetHeldItemClientPkt = internal.PlaySetHeldItem_763_0

// ID=0x4e
type PlaySetCenterChunkPkt = internal.PlaySetCenterChunk_763_0

// ID=0x4f
type PlaySetRenderDistancePkt = internal.PlaySetRenderDistance_763_0

// ID=0x50
type PlaySetDefaultSpawnPositionPkt = internal.PlaySetDefaultSpawnPosition_763_0

// ID=0x51
type PlayDisplayObjectivePkt = internal.PlayDisplayObjective_763_0

// ID=0x52
type PlaySetEntityMetadataPkt = internal.PlaySetEntityMetadata_763_0

// ID=0x53
type PlayLinkEntitiesPkt = internal.PlayLinkEntities_763_0

// ID=0x54
type PlaySetEntityVelocityPkt = internal.PlaySetEntityVelocity_763_0

// ID=0x55
type PlaySetEquipmentPkt = internal.PlaySetEquipment_763_0

// ID=0x56
type PlaySetExperiencePkt = internal.PlaySetExperience_763_0

// ID=0x57
type PlaySetHealthPkt = internal.PlaySetHealth_763_0

// ID=0x58
type PlayUpdateObjectivesPkt = internal.PlayUpdateObjectives_763_0

// ID=0x59
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x5a
type PlayUpdateTeamsPkt = internal.PlayUpdateTeams_763_0

// ID=0x5b
type PlayUpdateScorePkt = internal.PlayUpdateScore_763_0

// ID=0x5c
type PlaySetSimulationDistancePkt = internal.PlaySetSimulationDistance_763_0

// ID=0x5d
type PlaySetSubtitleTextPkt = internal.PlaySetSubtitleText_763_0

// ID=0x5e
type PlayUpdateTimePkt = internal.PlayUpdateTime_763_0

// ID=0x5f
type PlaySetTitleTextPkt = internal.PlaySetTitleText_763_0

// ID=0x60
type PlaySetTitleAnimationTimesPkt = internal.PlaySetTitleAnimationTimes_763_0

// ID=0x61
type PlayEntitySoundEffectPkt = internal.PlayEntitySoundEffect_763_0

// ID=0x62
type PlaySoundEffectPkt = internal.PlaySoundEffect_763_0

// ID=0x63
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x64
type PlaySystemChatMessagePkt = internal.PlaySystemChatMessage_763_0

// ID=0x65
type PlaySetTabListHeaderAndFooterPkt = internal.PlaySetTabListHeaderAndFooter_763_0

// ID=0x66
type PlayTagQueryResponsePkt = internal.PlayTagQueryResponse_763_0

// ID=0x67
type PlayPickupItemPkt = internal.PlayPickupItem_763_0

// ID=0x68
type PlayTeleportEntityPkt = internal.PlayTeleportEntity_763_0

// ID=0x69
type PlayUpdateAdvancementsPkt = internal.PlayUpdateAdvancements_763_0

// ID=0x6a
type PlayUpdateAttributesPkt = internal.PlayUpdateAttributes_763_0

// ID=0x6b
type PlayFeatureFlagsPkt = internal.PlayFeatureFlags_763_0

// ID=0x6c
type PlayEntityEffectPkt = internal.PlayEntityEffect_763_0

// ID=0x6d
type PlayUpdateRecipesPkt = internal.PlayUpdateRecipes_763_0

// ID=0x6e
type PlayUpdateTagsPkt = internal.PlayUpdateTags_763_0

// ======== END play ========
