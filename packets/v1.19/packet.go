
// Generated at 2023-09-02 21:26:37.473 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=17753
// Protocol: 759
// Protocol Name: 1.19

package packet_1_19

import (
	"io"
	. "github.com/kmcsr/go-liter"
	internal "github.com/kmcsr/go-liter/packets/internal"
)

// ======== BEGIN login ========
// ---- login: serverbound ----

// ID=0x0
type LoginStartPkt = internal.LoginStart_760_1

// ID=0x1
type LoginEncryptionResponsePkt = internal.LoginEncryptionResponse_760_1

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
type PlayChatCommandPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                   | Field Name       | Field Type | Field Type | Notes                                                               |
	 * |-----------|-------|----------|------------------------------|------------------|------------|------------|---------------------------------------------------------------------|
	 * | 0x03      | Play  | Server   |                              |                  |            |            |                                                                     |
	 * | 0x03      | Play  | Server   | Command                      | Command          | String     | String     | The command typed by the client.                                    |
	 * | 0x03      | Play  | Server   | Timestamp                    | Timestamp        | Long       | Long       | The timestamp that the command was executed.                        |
	 * | 0x03      | Play  | Server   | Salt                         | Salt             | Long       | Long       | The salt for the following argument signatures.                     |
	 * | 0x03      | Play  | Server   | Array length                 | Array length     | VarInt     | VarInt     | The length of the following array                                   |
	 * | 0x03      | Play  | Server   | Array of argument signatures | Argument name    | Array      | String     | The name of the argument that is signed by the following signature. |
	 * | 0x03      | Play  | Server   | Array of argument signatures | Signature length | Array      | VarInt     | The length of the follow byte array                                 |
	 * | 0x03      | Play  | Server   | Array of argument signatures | Signature        | Array      | Byte Array | The signature that verifies the argument.                           |
	 * | 0x03      | Play  | Server   | Signed Preview               | Signed Preview   | Boolean    | Boolean    |                                                                     |
	 * 
	 */
}

// ID=0x4
type PlayChatMessagePkt struct {
	Message String // String (256 chars)
	Timestamp Long // Instant
	/* The salt used to verify the signature hash. */
	Salt Long // Long
	/* The length of the following array. */
	SignatureLength VarInt // VarInt
	/* The signature used to verify the chat message's authentication. */
	Signature ByteArray // Byte Array
	SignedPreview Bool // Boolean
}

var _ Packet = (*PlayChatMessagePkt)(nil)

func (p PlayChatMessagePkt)Encode(b *PacketBuilder){
	b.String(p.Message)
	b.Long(p.Timestamp)
	b.Long(p.Salt)
	p.SignatureLength = (VarInt)(len(p.Signature))
	b.VarInt(p.SignatureLength)
	b.ByteArray(p.Signature)
	b.Bool(p.SignedPreview)
}

func (p *PlayChatMessagePkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Timestamp, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.Salt, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.SignatureLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Signature = make(ByteArray, p.SignatureLength)
	if ok = r.ByteArray(p.Signature); !ok {
		return io.EOF
	}
	if p.SignedPreview, ok = r.Bool(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x5
type PlayChatPreviewServerPkt = internal.PlayChatPreviewServer_760_0

// ID=0x6
type PlayClientCommandPkt = internal.PlayClientCommand_763_0

// ID=0x7
type PlayClientInformationPkt = internal.PlayClientInformation_763_0

// ID=0x8
type PlayCommandSuggestionsRequestPkt = internal.PlayCommandSuggestionsRequest_763_0

// ID=0x9
type PlayClickContainerButtonPkt = internal.PlayClickContainerButton_763_0

// ID=0xa
type PlayClickContainerPkt = internal.PlayClickContainer_761_1

// ID=0xb
type PlayCloseContainerServerPkt = internal.PlayCloseContainerServer_763_0

// ID=0xc
type PlayPluginMessageServerPkt = internal.PlayPluginMessageServer_763_0

// ID=0xd
type PlayEditBookPkt = internal.PlayEditBook_763_0

// ID=0xe
type PlayQueryEntityTagPkt = internal.PlayQueryEntityTag_763_0

// ID=0xf
type PlayInteractPkt = internal.PlayInteract_763_0

// ID=0x10
type PlayJigsawGeneratePkt = internal.PlayJigsawGenerate_763_0

// ID=0x11
type PlayKeepAliveServerPkt = internal.PlayKeepAliveServer_763_0

// ID=0x12
type PlayLockDifficultyPkt = internal.PlayLockDifficulty_763_0

// ID=0x13
type PlaySetPlayerPositionPkt = internal.PlaySetPlayerPosition_763_0

// ID=0x14
type PlaySetPlayerPositionAndRotationPkt = internal.PlaySetPlayerPositionAndRotation_763_0

// ID=0x15
type PlaySetPlayerRotationPkt = internal.PlaySetPlayerRotation_763_0

// ID=0x16
type PlaySetPlayerOnGroundPkt = internal.PlaySetPlayerOnGround_763_0

// ID=0x17
type PlayMoveVehicleServerPkt = internal.PlayMoveVehicleServer_763_0

// ID=0x18
type PlayPaddleBoatPkt = internal.PlayPaddleBoat_763_0

// ID=0x19
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x1a
type PlayPlaceRecipePkt = internal.PlayPlaceRecipe_763_0

// ID=0x1b
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_763_0

// ID=0x1c
type PlayerActionPkt = internal.PlayerAction_763_0

// ID=0x1d
type PlayerCommandPkt = internal.PlayerCommand_763_0

// ID=0x1e
type PlayerInputPkt = internal.PlayerInput_763_0

// ID=0x1f
type PlayPongPkt = internal.PlayPong_763_0

// ID=0x20
type PlayChangeRecipeBookSettingsPkt = internal.PlayChangeRecipeBookSettings_763_0

// ID=0x21
type PlaySetSeenRecipePkt = internal.PlaySetSeenRecipe_763_0

// ID=0x22
type PlayRenameItemPkt = internal.PlayRenameItem_763_0

// ID=0x23
type PlayResourcePackServerPkt = internal.PlayResourcePackServer_763_0

// ID=0x24
type PlaySeenAdvancementsPkt = internal.PlaySeenAdvancements_763_0

// ID=0x25
type PlaySelectTradePkt = internal.PlaySelectTrade_763_0

// ID=0x26
type PlaySetBeaconEffectPkt struct {
	PrimaryEffectPresent Bool // Boolean
	/* A Potion ID. (Was a full Integer for the plugin message). */
	PrimaryEffect VarInt // VarInt
	SecondaryEffectPresent Bool // Boolean
	/* A Potion ID. (Was a full Integer for the plugin message). */
	SecondaryEffect VarInt // VarInt
}

var _ Packet = (*PlaySetBeaconEffectPkt)(nil)

func (p PlaySetBeaconEffectPkt)Encode(b *PacketBuilder){
	b.Bool(p.PrimaryEffectPresent)
	b.VarInt(p.PrimaryEffect)
	b.Bool(p.SecondaryEffectPresent)
	b.VarInt(p.SecondaryEffect)
}

func (p *PlaySetBeaconEffectPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.PrimaryEffectPresent, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.PrimaryEffect, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.SecondaryEffectPresent, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.SecondaryEffect, ok = r.VarInt(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x27
type PlaySetHeldItemServerPkt = internal.PlaySetHeldItemServer_763_0

// ID=0x28
type PlayProgramCommandBlockPkt = internal.PlayProgramCommandBlock_763_0

// ID=0x29
type PlayProgramCommandBlockMinecartPkt = internal.PlayProgramCommandBlockMinecart_763_0

// ID=0x2a
type PlaySetCreativeModeSlotPkt = internal.PlaySetCreativeModeSlot_763_0

// ID=0x2b
type PlayProgramJigsawBlockPkt = internal.PlayProgramJigsawBlock_763_0

// ID=0x2c
type PlayProgramStructureBlockPkt = internal.PlayProgramStructureBlock_763_0

// ID=0x2d
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x2e
type PlaySwingArmPkt = internal.PlaySwingArm_763_0

// ID=0x2f
type PlayTeleportToEntityPkt = internal.PlayTeleportToEntity_763_0

// ID=0x30
type PlayUseItemOnPkt = internal.PlayUseItemOn_763_0

// ID=0x31
type PlayUseItemPkt = internal.PlayUseItem_763_0

// ---- play: clientbound ----

// ID=0x0
type PlaySpawnEntityPkt = internal.PlaySpawnEntity_763_0

// ID=0x1
type PlaySpawnExperienceOrbPkt = internal.PlaySpawnExperienceOrb_763_0

// ID=0x2
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_763_0

// ID=0x3
type PlayEntityAnimationPkt = internal.PlayEntityAnimation_763_0

// ID=0x4
type PlayAwardStatisticsPkt = internal.PlayAwardStatistics_761_1

// ID=0x5
type PlayAcknowledgeBlockChangePkt = internal.PlayAcknowledgeBlockChange_763_0

// ID=0x6
type PlaySetBlockDestroyStagePkt = internal.PlaySetBlockDestroyStage_763_0

// ID=0x7
type PlayBlockEntityDataPkt = internal.PlayBlockEntityData_763_0

// ID=0x8
type PlayBlockActionPkt = internal.PlayBlockAction_763_0

// ID=0x9
type PlayBlockUpdatePkt = internal.PlayBlockUpdate_763_0

// ID=0xa
type PlayBossBarPkt = internal.PlayBossBar_761_1

// ID=0xb
type PlayChangeDifficultyClientPkt = internal.PlayChangeDifficulty_763_0

// ID=0xc
type PlayChatPreviewClientPkt = internal.PlayChatPreview_760_0

// ID=0xd
type PlayClearTitlesPkt = internal.PlayClearTitles_763_0

// ID=0xe
type PlayCommandSuggestionsResponsePkt = internal.PlayCommandSuggestionsResponse_760_2

// ID=0xf
type PlayCommandsPkt = internal.PlayCommands_763_0

// ID=0x10
type PlayCloseContainerClientPkt = internal.PlayCloseContainer_763_0

// ID=0x11
type PlaySetContainerContentPkt = internal.PlaySetContainerContent_763_0

// ID=0x12
type PlaySetContainerPropertyPkt = internal.PlaySetContainerProperty_763_0

// ID=0x13
type PlaySetContainerSlotPkt = internal.PlaySetContainerSlot_763_0

// ID=0x14
type PlaySetCooldownPkt = internal.PlaySetCooldown_763_0

// ID=0x15
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_763_0

// ID=0x16
type PlayCustomSoundEffectPkt = internal.PlayCustomSoundEffect_760_0

// ID=0x17
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x18
type PlayEntityEventPkt = internal.PlayEntityEvent_763_0

// ID=0x19
type PlayExplosionPkt = internal.PlayExplosion_763_0

// ID=0x1a
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x1b
type PlayGameEventPkt = internal.PlayGameEvent_763_0

// ID=0x1c
type PlayOpenHorseScreenPkt = internal.PlayOpenHorseScreen_763_0

// ID=0x1d
type PlayInitializeWorldBorderPkt = internal.PlayInitializeWorldBorder_763_0

// ID=0x1e
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x1f
type PlayChunkDataAndUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name               | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
	 * |-----------|-------|----------|--------------------------|--------------------------|------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x1F      | Play  | Client   | Chunk X                  | Chunk X                  | Int        | Int                 | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x1F      | Play  | Client   | Chunk Z                  | Chunk Z                  | Int        | Int                 | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x1F      | Play  | Client   | Heightmaps               | Heightmaps               | NBT        | NBT                 | Compound containing one long array named MOTION_BLOCKING, which is a heightmap for the highest solid block at each position in the chunk (as a compacted long array with 256 entries, with the number of bits per entry varying depending on the world's height, defined by the formula ceil(log2(height + 1))). The Notchian server also adds a WORLD_SURFACE long array, the purpose of which is unknown, but it's not required for the chunk to be accepted. |
	 * | 0x1F      | Play  | Client   | Size                     | Size                     | VarInt     | VarInt              | Size of Data in bytes                                                                                                                                                                                                                                                                                                                                                                                                                                           |
	 * | 0x1F      | Play  | Client   | Data                     | Data                     | Byte array | Byte array          | See data structure in Chunk Format                                                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x1F      | Play  | Client   | Number of block entities | Number of block entities | VarInt     | VarInt              | Number of elements in the following array                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x1F      | Play  | Client   | Block Entity             | Packed XZ                | Array      | Byte                | The packed section coordinates, calculated from ((blockX & 15) << 4) | (blockZ & 15)                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x1F      | Play  | Client   | Block Entity             | Y                        | Array      | Short               | The height relative to the world                                                                                                                                                                                                                                                                                                                                                                                                                                |
	 * | 0x1F      | Play  | Client   | Block Entity             | Type                     | Array      | VarInt              | The type of block entity                                                                                                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x1F      | Play  | Client   | Block Entity             | Data                     | Array      | NBT                 | The block entity's data, without the X, Y, and Z values                                                                                                                                                                                                                                                                                                                                                                                                         |
	 * | 0x1F      | Play  | Client   | Trust Edges              | Trust Edges              | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x1F      | Play  | Client   | Sky Light Mask           | Sky Light Mask           | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Sky Light array below.  The least significant bit is for blocks 16 blocks to 1 block below the min world height (one section below the world), while the most significant bit covers blocks 1 to 16 blocks above the max world height (one section above the world).                                          |
	 * | 0x1F      | Play  | Client   | Block Light Mask         | Block Light Mask         | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Block Light array below.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                 |
	 * | 0x1F      | Play  | Client   | Empty Sky Light Mask     | Empty Sky Light Mask     | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Sky Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                    |
	 * | 0x1F      | Play  | Client   | Empty Block Light Mask   | Empty Block Light Mask   | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Block Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                  |
	 * | 0x1F      | Play  | Client   | Sky Light array count    | Sky Light array count    | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Sky Light Mask                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x1F      | Play  | Client   | Sky Light arrays         | Length                   | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x1F      | Play  | Client   | Sky Light arrays         | Sky Light array          | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value. Indexed ((y<<8) | (z<<4) | x) / 2  If there's a remainder, masked 0xF0 else 0x0F.                                                                                                                                                                                                                                                |
	 * | 0x1F      | Play  | Client   | Block Light array count  | Block Light array count  | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Block Light Mask                                                                                                                                                                                                                                                                                                                                                               |
	 * | 0x1F      | Play  | Client   | Block Light arrays       | Length                   | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x1F      | Play  | Client   | Block Light arrays       | Block Light array        | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value. Indexed ((y<<8) | (z<<4) | x) / 2  If there's a remainder, masked 0xF0 else 0x0F.                                                                                                                                                                                                                                              |
	 * 
	 */
}

// ID=0x20
type PlayWorldEventPkt = internal.PlayWorldEvent_763_0

// ID=0x21
type PlayParticlePkt = internal.PlayParticle_763_0

// ID=0x22
type PlayUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name              | Field Name              | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                  |
	 * |-----------|-------|----------|-------------------------|-------------------------|------------|---------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x22      | Play  | Client   | Chunk X                 | Chunk X                 | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x22      | Play  | Client   | Chunk Z                 | Chunk Z                 | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x22      | Play  | Client   | Trust Edges             | Trust Edges             | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                                                          |
	 * | 0x22      | Play  | Client   | Sky Light Mask          | Sky Light Mask          | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Sky Light array below.  The least significant bit is for blocks 16 blocks to 1 block below the min world height (one section below the world), while the most significant bit covers blocks 1 to 16 blocks above the max world height (one section above the world). |
	 * | 0x22      | Play  | Client   | Block Light Mask        | Block Light Mask        | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Block Light array below.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                        |
	 * | 0x22      | Play  | Client   | Empty Sky Light Mask    | Empty Sky Light Mask    | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Sky Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                           |
	 * | 0x22      | Play  | Client   | Empty Block Light Mask  | Empty Block Light Mask  | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Block Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                         |
	 * | 0x22      | Play  | Client   | Sky Light array count   | Sky Light array count   | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Sky Light Mask                                                                                                                                                                                                                                                                                                                        |
	 * | 0x22      | Play  | Client   | Sky Light arrays        | Length                  | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x22      | Play  | Client   | Sky Light arrays        | Sky Light array         | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                                                         |
	 * | 0x22      | Play  | Client   | Block Light array count | Block Light array count | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Block Light Mask                                                                                                                                                                                                                                                                                                                      |
	 * | 0x22      | Play  | Client   | Block Light arrays      | Length                  | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x22      | Play  | Client   | Block Light arrays      | Block Light array       | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                                                       |
	 * 
	 */
}

// ID=0x23
type PlayLoginPkt = internal.PlayLogin_760_3

// ID=0x24
type PlayMapDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type                      | Field Type                      | Notes                                                                                                      |
	 * |-----------|-------|----------|-------------------|-------------------|---------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x24      | Play  | Client   | Map ID            | Map ID            | VarInt                          | VarInt                          | Map ID of the map being modified                                                                           |
	 * | 0x24      | Play  | Client   | Scale             | Scale             | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x24      | Play  | Client   | Locked            | Locked            | Boolean                         | Boolean                         | True if the map has been locked in a cartography table                                                     |
	 * | 0x24      | Play  | Client   | Tracking Position | Tracking Position | Boolean                         | Boolean                         | Specifies whether player and item frame icons are shown                                                    |
	 * | 0x24      | Play  | Client   | Icon Count        | Icon Count        | VarInt                          | VarInt                          | Number of elements in the following array. Only present if "Tracking Position" is true.                    |
	 * | 0x24      | Play  | Client   | Icon              | Type              | Array                           | VarInt enum                     | See below                                                                                                  |
	 * | 0x24      | Play  | Client   | Icon              | X                 | Array                           | Byte                            | Map coordinates: -128 for furthest left, +127 for furthest right                                           |
	 * | 0x24      | Play  | Client   | Icon              | Z                 | Array                           | Byte                            | Map coordinates: -128 for highest, +127 for lowest                                                         |
	 * | 0x24      | Play  | Client   | Icon              | Direction         | Array                           | Byte                            | 0-15                                                                                                       |
	 * | 0x24      | Play  | Client   | Icon              | Has Display Name  | Array                           | Boolean                         |                                                                                                            |
	 * | 0x24      | Play  | Client   | Icon              | Display Name      | Array                           | Optional Chat                   | Only present if previous Boolean is true                                                                   |
	 * | 0x24      | Play  | Client   | Columns           | Columns           | Unsigned Byte                   | Unsigned Byte                   | Number of columns updated                                                                                  |
	 * | 0x24      | Play  | Client   | Rows              | Rows              | Optional Unsigned Byte          | Optional Unsigned Byte          | Only if Columns is more than 0; number of rows updated                                                     |
	 * | 0x24      | Play  | Client   | X                 | X                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column                                         |
	 * | 0x24      | Play  | Client   | Z                 | Z                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row                                           |
	 * | 0x24      | Play  | Client   | Length            | Length            | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array                                              |
	 * | 0x24      | Play  | Client   | Data              | Data              | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format                                                        |
	 * 
	 */
}

// ID=0x25
type PlayMerchantOffersPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name                   | Field Type | Field Type    | Notes                                                                                                                                                                             |
	 * |-----------|-------|----------|---------------------|------------------------------|------------|---------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x25      | Play  | Client   |                     |                              |            |               |                                                                                                                                                                                   |
	 * | 0x25      | Play  | Client   | Window ID           | Window ID                    | VarInt     | VarInt        | The ID of the window that is open; this is an int rather than a byte.                                                                                                             |
	 * | 0x25      | Play  | Client   | Size                | Size                         | VarInt     | VarInt        | The number of trades in the following array.                                                                                                                                      |
	 * | 0x25      | Play  | Client   | Trades              | Input item 1                 | Array      | Slot          | The first item the player has to supply for this villager trade. The count of the item stack is the default "price" of this trade.                                                |
	 * | 0x25      | Play  | Client   | Trades              | Output item                  | Array      | Slot          | The item the player will receive from this villager trade.                                                                                                                        |
	 * | 0x25      | Play  | Client   | Trades              | Has second item              | Array      | Boolean       | Whether there is a second item.                                                                                                                                                   |
	 * | 0x25      | Play  | Client   | Trades              | Input item 2                 | Array      | Optional Slot | The second item the player has to supply for this villager trade; only present if has a second item is true.                                                                      |
	 * | 0x25      | Play  | Client   | Trades              | Trade disabled               | Array      | Boolean       | True if the trade is disabled; false if the trade is enabled.                                                                                                                     |
	 * | 0x25      | Play  | Client   | Trades              | Number of trade uses         | Array      | Integer       | Number of times the trade has been used so far. If equal to the maximum number of trades, the client will display a red X.                                                        |
	 * | 0x25      | Play  | Client   | Trades              | Maximum number of trade uses | Array      | Integer       | Number of times this trade can be used before it's exhausted.                                                                                                                     |
	 * | 0x25      | Play  | Client   | Trades              | XP                           | Array      | Integer       | Amount of XP the villager will earn each time the trade is used.                                                                                                                  |
	 * | 0x25      | Play  | Client   | Trades              | Special Price                | Array      | Integer       | Can be zero or negative. The number is added to the price when an item is discounted due to player reputation or other effects.                                                   |
	 * | 0x25      | Play  | Client   | Trades              | Price Multiplier             | Array      | Float         | Can be low (0.05) or high (0.2). Determines how much demand, player reputation, and temporary effects will adjust the price.                                                      |
	 * | 0x25      | Play  | Client   | Trades              | Demand                       | Array      | Integer       | If positive, causes the price to increase. Negative values seem to be treated the same as zero.                                                                                   |
	 * | 0x25      | Play  | Client   | Villager level      | Villager level               | VarInt     | VarInt        | Appears on the trade GUI; meaning comes from the translation key merchant.level. + level.
	 * 1: Novice, 2: Apprentice, 3: Journeyman, 4: Expert, 5: Master.                          |
	 * | 0x25      | Play  | Client   | Experience          | Experience                   | VarInt     | VarInt        | Total experience for this villager (always 0 for the wandering trader).                                                                                                           |
	 * | 0x25      | Play  | Client   | Is regular villager | Is regular villager          | Boolean    | Boolean       | True if this is a regular villager; false for the wandering trader.  When false, hides the villager level and some other GUI elements.                                            |
	 * | 0x25      | Play  | Client   | Can restock         | Can restock                  | Boolean    | Boolean       | True for regular villagers and false for the wandering trader. If true, the "Villagers restock up to two times per day." message is displayed when hovering over disabled trades. |
	 * 
	 */
}

// ID=0x26
type PlayUpdateEntityPositionPkt = internal.PlayUpdateEntityPosition_763_0

// ID=0x27
type PlayUpdateEntityPositionAndRotationPkt = internal.PlayUpdateEntityPositionAndRotation_763_0

// ID=0x28
type PlayUpdateEntityRotationPkt = internal.PlayUpdateEntityRotation_763_0

// ID=0x29
type PlayMoveVehicleClientPkt = internal.PlayMoveVehicle_763_0

// ID=0x2a
type PlayOpenBookPkt = internal.PlayOpenBook_763_0

// ID=0x2b
type PlayOpenScreenPkt = internal.PlayOpenScreen_763_0

// ID=0x2c
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x2d
type PlayPingPkt = internal.PlayPing_763_0

// ID=0x2e
type PlayPlaceGhostRecipePkt = internal.PlayPlaceGhostRecipe_763_0

// ID=0x2f
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x30
type PlayerChatMessagePkt struct {
	/* Player message, this could either be the player typed message or the server previewed message if it was signed by the client. Not modifiable by the server without invalidating the sigature. */
	SignedChatContent Object // Chat
	HasUnsignedChatContent Bool // Boolean
	/* Server modifiable message.
	 * The client will prefer to choose this UNLESS show only secure chat is enabled on the client or has unsigned chat content is false. If these conditions are met, it will instead use the signed chat content. */
	UnsignedChatContent Object // Chat
	/* Registered on the server, sent to the client.
	 * Default: 0: chat (chat box), 1: system message (chat box), 2: game info (above hotbar), 3: say command, 4: msg command, 5: team msg command, 6: emote command, 7: tellraw command */
	Type VarInt // VarInt
	/* Used by the Notchian client for the disableChat launch option. Setting both longs to 0 will always display the message regardless of the setting. */
	SenderUUID UUID // UUID
	/* This can be modified without disrupting the message signature. */
	SenderDisplayName Object // Chat
	HasSenderTeamName Bool // Boolean
	/* Used in the team_name decorator parameter. */
	SenderTeamName Object // Chat
	/* Represents the time the message was signed, used to check if the message was received within 2 minutes of it being sent. */
	Timestamp Long // Long
	/* Cryptography, used for validating the message signature. */
	Salt Long // Long
	/* Array Length */
	SignatureLength VarInt // VarInt
	/* Cryptography, the signature consists of the Sender UUID, Timestamp and Original Chat Content. Modifying any of these values in the packet will cause this signature to fail. */
	MessageSignature []Byte // Bytes
}

var _ Packet = (*PlayerChatMessagePkt)(nil)

func (p PlayerChatMessagePkt)Encode(b *PacketBuilder){
	b.JSON(p.SignedChatContent)
	b.Bool(p.HasUnsignedChatContent)
	b.JSON(p.UnsignedChatContent)
	b.VarInt(p.Type)
	b.UUID(p.SenderUUID)
	b.JSON(p.SenderDisplayName)
	b.Bool(p.HasSenderTeamName)
	b.JSON(p.SenderTeamName)
	b.Long(p.Timestamp)
	b.Long(p.Salt)
	p.SignatureLength = (VarInt)(len(p.MessageSignature))
	b.VarInt(p.SignatureLength)
	for _, v := range p.MessageSignature {
		b.Byte(v)
	}
}

func (p *PlayerChatMessagePkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if err = r.JSON(&p.SignedChatContent); err != nil {
		return err
	}
	if p.HasUnsignedChatContent, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.UnsignedChatContent); err != nil {
		return err
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.SenderUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.SenderDisplayName); err != nil {
		return err
	}
	if p.HasSenderTeamName, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.SenderTeamName); err != nil {
		return err
	}
	if p.Timestamp, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.Salt, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.SignatureLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.MessageSignature = make([]Byte, p.SignatureLength)
	for i, _ := range p.MessageSignature {
		if p.MessageSignature[i], ok = r.Byte(); !ok {
			return io.EOF
		}
	}
	return nil
}

// ID=0x31
type PlayEndCombatPkt = internal.PlayEndCombat_762_1

// ID=0x32
type PlayEnterCombatPkt = internal.PlayEnterCombat_763_0

// ID=0x33
type PlayCombatDeathPkt = internal.PlayCombatDeath_762_1

// ID=0x34
type PlayerInfoPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type              | Notes                                                                  |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-------------------------|------------------------------------------------------------------------|
	 * | 0x34      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt                  | Determines the rest of the Player format after the UUID.               |
	 * | 0x34      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt                  | Number of elements in the following array.                             |
	 * | 0x34      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID                    |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                         |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String (16)   | String (16)             |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt                  | Number of elements in the following array.                             |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String (32767)          |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String (32767)          |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean                 |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String (32767) | Only if Is Signed is true.                                             |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds.                                              |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only if Has Display Name is true.                                      |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Has Sig Data         | Has Sig Data         | Array      | Boolean       | Boolean                 | Whether or not the next 5 fields should be sent.                       |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Timestamp            | Timestamp            | Array      | Long          | Long                    | When the key data will expire.                                         |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Public Key Length    | Public Key Length    | Array      | VarInt        | VarInt                  | Length of Public Key. Optional.                                        |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Public Key           | Public Key           | Array      | Byte Array    | Byte Array              | The encoded bytes of the public key the client received from Mojang.   |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Signature Length     | Signature Length     | Array      | VarInt        | VarInt                  | Length of Signature.                                                   |
	 * | 0x34      | Play  | Client   | Player            | 0: add player          | Signature            | Signature            | Array      | Byte Array    | Byte Array              | The bytes of the public key signature the client received from Mojang. |
	 * | 0x34      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds.                                              |
	 * | 0x34      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                                        |
	 * | 0x34      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true.                                 |
	 * | 0x34      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields               |                                                                        |
	 * 
	 */
}

// ID=0x35
type PlayLookAtPkt = internal.PlayLookAt_763_0

// ID=0x36
type PlaySynchronizePlayerPositionPkt = internal.PlaySynchronizePlayerPosition_761_1

// ID=0x37
type PlayUpdateRecipeBookPkt = internal.PlayUpdateRecipeBook_763_0

// ID=0x38
type PlayRemoveEntitiesPkt = internal.PlayRemoveEntities_763_0

// ID=0x39
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_763_0

// ID=0x3a
type PlayResourcePackClientPkt = internal.PlayResourcePack_763_0

// ID=0x3b
type PlayRespawnPkt = internal.PlayRespawn_760_3

// ID=0x3c
type PlaySetHeadRotationPkt = internal.PlaySetHeadRotation_763_0

// ID=0x3d
type PlayUpdateSectionBlocksPkt = internal.PlayUpdateSectionBlocks_762_1

// ID=0x3e
type PlaySelectAdvancementsTabPkt = internal.PlaySelectAdvancementsTab_763_0

// ID=0x3f
type PlayServerDataPkt = internal.PlayServerData_760_2

// ID=0x40
type PlaySetActionBarTextPkt = internal.PlaySetActionBarText_763_0

// ID=0x41
type PlaySetBorderCenterPkt = internal.PlaySetBorderCenter_763_0

// ID=0x42
type PlaySetBorderLerpSizePkt = internal.PlaySetBorderLerpSize_763_0

// ID=0x43
type PlaySetBorderSizePkt = internal.PlaySetBorderSize_763_0

// ID=0x44
type PlaySetBorderWarningDelayPkt = internal.PlaySetBorderWarningDelay_763_0

// ID=0x45
type PlaySetBorderWarningDistancePkt = internal.PlaySetBorderWarningDistance_763_0

// ID=0x46
type PlaySetCameraPkt = internal.PlaySetCamera_763_0

// ID=0x47
type PlaySetHeldItemClientPkt = internal.PlaySetHeldItem_763_0

// ID=0x48
type PlaySetCenterChunkPkt = internal.PlaySetCenterChunk_763_0

// ID=0x49
type PlaySetRenderDistancePkt = internal.PlaySetRenderDistance_763_0

// ID=0x4a
type PlaySetDefaultSpawnPositionPkt = internal.PlaySetDefaultSpawnPosition_763_0

// ID=0x4b
type PlaySetDisplayChatPreviewPkt = internal.PlaySetDisplayChatPreview_760_0

// ID=0x4c
type PlayDisplayObjectivePkt = internal.PlayDisplayObjective_763_0

// ID=0x4d
type PlaySetEntityMetadataPkt = internal.PlaySetEntityMetadata_763_0

// ID=0x4e
type PlayLinkEntitiesPkt = internal.PlayLinkEntities_763_0

// ID=0x4f
type PlaySetEntityVelocityPkt = internal.PlaySetEntityVelocity_763_0

// ID=0x50
type PlaySetEquipmentPkt struct {
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
type PlaySetHealthPkt = internal.PlaySetHealth_763_0

// ID=0x53
type PlayUpdateObjectivesPkt = internal.PlayUpdateObjectives_763_0

// ID=0x54
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x55
type PlayUpdateTeamsPkt struct {
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
type PlayUpdateScorePkt = internal.PlayUpdateScore_763_0

// ID=0x57
type PlaySetSimulationDistancePkt = internal.PlaySetSimulationDistance_763_0

// ID=0x58
type PlaySetSubtitleTextPkt = internal.PlaySetSubtitleText_763_0

// ID=0x59
type PlayUpdateTimePkt = internal.PlayUpdateTime_763_0

// ID=0x5a
type PlaySetTitleTextPkt = internal.PlaySetTitleText_763_0

// ID=0x5b
type PlaySetTitleAnimationTimesPkt = internal.PlaySetTitleAnimationTimes_763_0

// ID=0x5c
type PlayEntitySoundEffectPkt = internal.PlayEntitySoundEffect_760_1

// ID=0x5d
type PlaySoundEffectPkt = internal.PlaySoundEffect_760_1

// ID=0x5e
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x5f
type PlaySystemChatMessagePkt = internal.PlaySystemChatMessage_760_1

// ID=0x60
type PlaySetTabListHeaderAndFooterPkt = internal.PlaySetTabListHeaderAndFooter_763_0

// ID=0x61
type PlayTagQueryResponsePkt = internal.PlayTagQueryResponse_763_0

// ID=0x62
type PlayPickupItemPkt = internal.PlayPickupItem_763_0

// ID=0x63
type PlayTeleportEntityPkt = internal.PlayTeleportEntity_763_0

// ID=0x64
type PlayUpdateAdvancementsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name    | Field Type          | Field Type           | Notes                                                       |
	 * |-----------|-------|----------|---------------------|---------------|---------------------|----------------------|-------------------------------------------------------------|
	 * | 0x64      | Play  | Client   | Reset/Clear         | Reset/Clear   | Boolean             | Boolean              | Whether to reset/clear the current advancements.            |
	 * | 0x64      | Play  | Client   | Mapping size        | Mapping size  | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x64      | Play  | Client   | Advancement mapping | Key           | Array               | Identifier           | The identifier of the advancement.                          |
	 * | 0x64      | Play  | Client   | Advancement mapping | Value         | Array               | Advancement          | See below                                                   |
	 * | 0x64      | Play  | Client   | List size           | List size     | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x64      | Play  | Client   | Identifiers         | Identifiers   | Array of Identifier | Array of Identifier  | The identifiers of the advancements that should be removed. |
	 * | 0x64      | Play  | Client   | Progress size       | Progress size | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x64      | Play  | Client   | Progress mapping    | Key           | Array               | Identifier           | The identifier of the advancement.                          |
	 * | 0x64      | Play  | Client   | Progress mapping    | Value         | Array               | Advancement progress | See below.                                                  |
	 * 
	 */
}

// ID=0x65
type PlayUpdateAttributesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x65      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x65      | Play  | Client   | Number Of Properties | Number Of Properties | VarInt     | VarInt                 | Number of elements in the following array.            |
	 * | 0x65      | Play  | Client   | Property             | Key                  | Array      | Identifier             | See below.                                            |
	 * | 0x65      | Play  | Client   | Property             | Value                | Array      | Double                 | See below.                                            |
	 * | 0x65      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array.            |
	 * | 0x65      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x66
type PlayEntityEffectPkt = internal.PlayEntityEffect_763_0

// ID=0x67
type PlayUpdateRecipesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name  | Field Type | Field Type       | Notes                                                                   |
	 * |-----------|-------|----------|-------------|-------------|------------|------------------|-------------------------------------------------------------------------|
	 * | 0x67      | Play  | Client   | Num Recipes | Num Recipes | VarInt     | VarInt           | Number of elements in the following array.                              |
	 * | 0x67      | Play  | Client   | Recipe      | Type        | Array      | Identifier       | The recipe type, see below.                                             |
	 * | 0x67      | Play  | Client   | Recipe      | Recipe ID   | Array      | Identifier       |                                                                         |
	 * | 0x67      | Play  | Client   | Recipe      | Data        | Array      | Optional, varies | Additional data for the recipe.  For some types, there will be no data. |
	 * 
	 */
}

// ID=0x68
type PlayUpdateTagsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type | Field Type  | Notes                                                                                                                                        |
	 * |-----------|-------|----------|---------------------|---------------------|------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x68      | Play  | Client   | Length of the array | Length of the array | VarInt     | VarInt      |                                                                                                                                              |
	 * | 0x68      | Play  | Client   | Array of tags       | Tag type            | Array      | Identifier  | Tag identifier (Vanilla required tags are minecraft:block, minecraft:item, minecraft:fluid, minecraft:entity_type, and minecraft:game_event) |
	 * | 0x68      | Play  | Client   | Array of tags       | Array of Tag        | Array      | (See below) |                                                                                                                                              |
	 * 
	 */
}

// ======== END play ========
