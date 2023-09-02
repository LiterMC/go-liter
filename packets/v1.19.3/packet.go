
// Generated at 2023-09-01 20:45:22.208 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=18067
// Protocol: 761

package packet_1_19_3

import (
	"io"
	. "github.com/kmcsr/go-liter"
	internal "github.com/kmcsr/go-liter/packets/internal"
)

// ======== BEGIN login ========
// ---- login: serverbound ----

// ID=0x0
type LoginLoginStartPkt = internal.LoginLoginStart_763_0

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
type LoginLoginSuccessPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type  | Field Type              | Notes                                      |
	 * |-----------|-------|----------|----------------------|----------------------|-------------|-------------------------|--------------------------------------------|
	 * | 0x02      | Login | Client   | UUID                 | UUID                 | UUID        | UUID                    |                                            |
	 * | 0x02      | Login | Client   | Username             | Username             | String (16) | String (16)             |                                            |
	 * | 0x02      | Login | Client   | Number Of Properties | Number Of Properties | VarInt      | VarInt                  | Number of elements in the following array. |
	 * | 0x02      | Login | Client   | Property             | Name                 | Array       | String (32767)          |                                            |
	 * | 0x02      | Login | Client   | Property             | Value                | Array       | String (32767)          |                                            |
	 * | 0x02      | Login | Client   | Property             | Is Signed            | Array       | Boolean                 |                                            |
	 * | 0x02      | Login | Client   | Property             | Signature            | Array       | Optional String (32767) | Only if Is Signed is true.                 |
	 * 
	 */
}

// ID=0x3
type LoginSetCompressionPkt = internal.LoginSetCompression_763_0

// ID=0x4
type LoginLoginPluginRequestPkt = internal.LoginLoginPluginRequest_763_0

// ======== END login ========

// ======== BEGIN play ========
// ---- play: serverbound ----

// ID=0x0
type PlayConfirmTeleportationPkt = internal.PlayConfirmTeleportation_763_0

// ID=0x1
type PlayQueryBlockEntityTagPkt = internal.PlayQueryBlockEntityTag_763_0

// ID=0x2
type PlayChangeDifficultyServerPkt = internal.PlayChangeDifficulty_761_3

// ID=0x3
type PlayMessageAcknowledgmentPkt = internal.PlayMessageAcknowledgment_763_0

// ID=0x4
type PlayChatCommandPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                   | Field Name    | Field Type   | Field Type   | Notes                                                               |
	 * |-----------|-------|----------|------------------------------|---------------|--------------|--------------|---------------------------------------------------------------------|
	 * | 0x04      | Play  | Server   | Command                      | Command       | String (256) | String (256) | The command typed by the client.                                    |
	 * | 0x04      | Play  | Server   | Timestamp                    | Timestamp     | Long         | Long         | The timestamp that the command was executed.                        |
	 * | 0x04      | Play  | Server   | Salt                         | Salt          | Long         | Long         | The salt for the following argument signatures.                     |
	 * | 0x04      | Play  | Server   | Array length                 | Array length  | VarInt       | VarInt       | Number of entries in the following array                            |
	 * | 0x04      | Play  | Server   | Array of argument signatures | Argument name | Array        | String       | The name of the argument that is signed by the following signature. |
	 * | 0x04      | Play  | Server   | Array of argument signatures | Signature     | Array        | Byte Array   | The signature that verifies the argument. Always 256 bytes.         |
	 * | 0x04      | Play  | Server   | Message Count                | Message Count | VarInt       | VarInt       |                                                                     |
	 * | 0x04      | Play  | Server   | Acknowledged                 | Acknowledged  | BitSet (20)  | BitSet (20)  |                                                                     |
	 * 
	 */
}

// ID=0x5
type PlayChatMessagePkt = internal.PlayChatMessage_763_0

// ID=0x6
type PlayClientCommandPkt = internal.PlayClientCommand_763_0

// ID=0x7
type PlayClientInformationPkt = internal.PlayClientInformation_763_0

// ID=0x8
type PlayCommandSuggestionsRequestPkt = internal.PlayCommandSuggestionsRequest_763_0

// ID=0x9
type PlayClickContainerButtonPkt = internal.PlayClickContainerButton_763_0

// ID=0xa
type PlayClickContainerPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type    | Field Type    | Notes                                                                                                    |
	 * |-----------|-------|----------|---------------------|---------------------|---------------|---------------|----------------------------------------------------------------------------------------------------------|
	 * | 0x0A      | Play  | Server   | Window ID           | Window ID           | Unsigned Byte | Unsigned Byte | The ID of the window which was clicked. 0 for player inventory.                                          |
	 * | 0x0A      | Play  | Server   | State ID            | State ID            | VarInt        | VarInt        | The last recieved State ID from either a Set Container Slot or a Set Container Content packet            |
	 * | 0x0A      | Play  | Server   | Slot                | Slot                | Short         | Short         | The clicked slot number, see below.                                                                      |
	 * | 0x0A      | Play  | Server   | Button              | Button              | Byte          | Byte          | The button used in the click, see below.                                                                 |
	 * | 0x0A      | Play  | Server   | Mode                | Mode                | VarInt Enum   | VarInt Enum   | Inventory operation mode, see below.                                                                     |
	 * | 0x0A      | Play  | Server   | Length of the array | Length of the array | VarInt        | VarInt        |                                                                                                          |
	 * | 0x0A      | Play  | Server   | Array of slots      | Slot number         | Array         | Short         |                                                                                                          |
	 * | 0x0A      | Play  | Server   | Array of slots      | Slot data           | Array         | Slot          | New data for this slot                                                                                   |
	 * | 0x0A      | Play  | Server   | Carried item        | Carried item        | Slot          | Slot          | Item carried by the cursor. Has to be empty (item ID = -1) for drop mode, otherwise nothing will happen. |
	 * 
	 */
}

// ID=0xb
type PlayCloseContainerServerPkt = internal.PlayCloseContainer_763_0

// ID=0xc
type PlayPluginMessageServerPkt = internal.PlayPluginMessage_763_0

// ID=0xd
type PlayEditBookPkt = internal.PlayEditBook_763_0

// ID=0xe
type PlayQueryEntityTagPkt = internal.PlayQueryEntityTag_763_0

// ID=0xf
type PlayInteractPkt = internal.PlayInteract_763_0

// ID=0x10
type PlayJigsawGeneratePkt = internal.PlayJigsawGenerate_763_0

// ID=0x11
type PlayKeepAliveServerPkt = internal.PlayKeepAlive_763_0

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
type PlayMoveVehicleServerPkt = internal.PlayMoveVehicle_763_0

// ID=0x18
type PlayPaddleBoatPkt = internal.PlayPaddleBoat_763_0

// ID=0x19
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x1a
type PlayPlaceRecipePkt = internal.PlayPlaceRecipe_763_0

// ID=0x1b
type PlayPlayerAbilitiesServerPkt = internal.PlayPlayerAbilities_761_3

// ID=0x1c
type PlayPlayerActionPkt = internal.PlayPlayerAction_763_0

// ID=0x1d
type PlayPlayerCommandPkt = internal.PlayPlayerCommand_763_0

// ID=0x1e
type PlayPlayerInputPkt = internal.PlayPlayerInput_763_0

// ID=0x1f
type PlayPongPkt = internal.PlayPong_763_0

// ID=0x20
type PlayPlayerSessionPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name           | Field Type | Notes                                                                                                                                                                            |
	 * |-----------|-------|----------|------------|----------------------|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x20      | Play  | Server   | Session Id | Session Id           | UUID       |                                                                                                                                                                                  |
	 * | 0x20      | Play  | Server   | Public Key | Expires At           | Long       | The time the play session key expires in epoch milliseconds.                                                                                                                     |
	 * | 0x20      | Play  | Server   | Public Key | Public Key Length    | VarInt     | Length of the proceeding public key.                                                                                                                                             |
	 * | 0x20      | Play  | Server   | Public Key | Public Key           | Byte Array | A byte array of an X.509-encoded public key.                                                                                                                                     |
	 * | 0x20      | Play  | Server   | Public Key | Key Signature Length | VarInt     | Length of the proceeding key signature array.                                                                                                                                    |
	 * | 0x20      | Play  | Server   | Public Key | Key Signature        | Byte Array | The signature consists of the player UUID, the key expiration timestamp, and the public key data. These values are hashed using SHA-1 and signed using Mojang's private RSA key. |
	 * 
	 */
}

// ID=0x21
type PlayChangeRecipeBookSettingsPkt = internal.PlayChangeRecipeBookSettings_763_0

// ID=0x22
type PlaySetSeenRecipePkt = internal.PlaySetSeenRecipe_763_0

// ID=0x23
type PlayRenameItemPkt = internal.PlayRenameItem_763_0

// ID=0x24
type PlayResourcePackServerPkt = internal.PlayResourcePack_761_3

// ID=0x25
type PlaySeenAdvancementsPkt = internal.PlaySeenAdvancements_763_0

// ID=0x26
type PlaySelectTradePkt = internal.PlaySelectTrade_763_0

// ID=0x27
type PlaySetBeaconEffectPkt = internal.PlaySetBeaconEffect_763_0

// ID=0x28
type PlaySetHeldItemServerPkt = internal.PlaySetHeldItem_761_3

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
type PlaySpawnEntityPkt = internal.PlaySpawnEntity_763_0

// ID=0x1
type PlaySpawnExperienceOrbPkt = internal.PlaySpawnExperienceOrb_763_0

// ID=0x2
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_763_0

// ID=0x3
type PlayEntityAnimationPkt = internal.PlayEntityAnimation_763_0

// ID=0x4
type PlayAwardStatisticsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name   | Field Type | Field Type | Notes                                      |
	 * |-----------|-------|----------|------------|--------------|------------|------------|--------------------------------------------|
	 * | 0x04      | Play  | Client   | Count      | Count        | VarInt     | VarInt     | Number of elements in the following array. |
	 * | 0x04      | Play  | Client   | Statistic  | Category ID  | Array      | VarInt     | See below.                                 |
	 * | 0x04      | Play  | Client   | Statistic  | Statistic ID | Array      | VarInt     | See below.                                 |
	 * | 0x04      | Play  | Client   | Statistic  | Value        | Array      | VarInt     | The amount to set it to.                   |
	 * 
	 */
}

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
type PlayBossBarPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name       | Field Name | Field Type    | Notes                                                                                                                                     |
	 * |-----------|-------|----------|------------------|------------|---------------|-------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0A      | Play  | Client   | UUID             | UUID       | UUID          | Unique ID for this bar.                                                                                                                   |
	 * | 0x0A      | Play  | Client   | Action           | Action     | VarInt Enum   | Determines the layout of the remaining packet.                                                                                            |
	 * | 0x0A      | Play  | Client   | Action           | Field Name |               |                                                                                                                                           |
	 * | 0x0A      | Play  | Client   | 0: add           | Title      | Chat          |                                                                                                                                           |
	 * | 0x0A      | Play  | Client   | 0: add           | Health     | Float         | From 0 to 1. Values greater than 1 do not crash a Notchian client, and start rendering part of a second health bar at around 1.5.         |
	 * | 0x0A      | Play  | Client   | 0: add           | Color      | VarInt Enum   | Color ID (see below).                                                                                                                     |
	 * | 0x0A      | Play  | Client   | 0: add           | Division   | VarInt Enum   | Type of division (see below).                                                                                                             |
	 * | 0x0A      | Play  | Client   | 0: add           | Flags      | Unsigned Byte | Bit mask. 0x1: should darken sky, 0x2: is dragon bar (used to play end music), 0x04: create fog (previously was also controlled by 0x02). |
	 * | 0x0A      | Play  | Client   | 1: remove        | no fields  | no fields     | Removes this boss bar.                                                                                                                    |
	 * | 0x0A      | Play  | Client   | 2: update health | Health     | Float         | as above                                                                                                                                  |
	 * | 0x0A      | Play  | Client   | 3: update title  | Title      | Chat          |                                                                                                                                           |
	 * | 0x0A      | Play  | Client   | 4: update style  | Color      | VarInt Enum   | Color ID (see below).                                                                                                                     |
	 * | 0x0A      | Play  | Client   | 4: update style  | Dividers   | VarInt Enum   | as above                                                                                                                                  |
	 * | 0x0A      | Play  | Client   | 5: update flags  | Flags      | Unsigned Byte | as above                                                                                                                                  |
	 * 
	 */
}

// ID=0xb
type PlayChangeDifficultyClientPkt = internal.PlayChangeDifficulty_761_2

// ID=0xc
type PlayClearTitlesPkt = internal.PlayClearTitles_763_0

// ID=0xd
type PlayCommandSuggestionsResponsePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name  | Field Type | Field Type     | Notes                                                                                                                                                                                                  |
	 * |-----------|-------|----------|------------|-------------|------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0D      | Play  | Client   |            |             |            |                |                                                                                                                                                                                                        |
	 * | 0x0D      | Play  | Client   | ID         | ID          | VarInt     | VarInt         | Transaction ID.                                                                                                                                                                                        |
	 * | 0x0D      | Play  | Client   | Start      | Start       | VarInt     | VarInt         | Start of the text to replace.                                                                                                                                                                          |
	 * | 0x0D      | Play  | Client   | Length     | Length      | VarInt     | VarInt         | Length of the text to replace.                                                                                                                                                                         |
	 * | 0x0D      | Play  | Client   | Count      | Count       | VarInt     | VarInt         | Number of elements in the following array.                                                                                                                                                             |
	 * | 0x0D      | Play  | Client   | Matches    | Match       | Array      | String (32767) | One eligible value to insert, note that each command is sent separately instead of in a single string, hence the need for Count.  Note that for instance this doesn't include a leading / on commands. |
	 * | 0x0D      | Play  | Client   | Matches    | Has tooltip | Array      | Boolean        | True if the following is present.                                                                                                                                                                      |
	 * | 0x0D      | Play  | Client   | Matches    | Tooltip     | Array      | Optional Chat  | Tooltip to display; only present if previous boolean is true.                                                                                                                                          |
	 * 
	 */
}

// ID=0xe
type PlayCommandsPkt = internal.PlayCommands_763_0

// ID=0xf
type PlayCloseContainerClientPkt = internal.PlayCloseContainer_763_0

// ID=0x10
type PlaySetContainerContentPkt = internal.PlaySetContainerContent_763_0

// ID=0x11
type PlaySetContainerPropertyPkt = internal.PlaySetContainerProperty_763_0

// ID=0x12
type PlaySetContainerSlotPkt = internal.PlaySetContainerSlot_763_0

// ID=0x13
type PlaySetCooldownPkt = internal.PlaySetCooldown_763_0

// ID=0x14
type PlayChatSuggestionsPkt = internal.PlayChatSuggestions_763_0

// ID=0x15
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_763_0

// ID=0x16
type PlayDeleteMessagePkt = internal.PlayDeleteMessage_763_0

// ID=0x17
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x18
type PlayDisguisedChatMessagePkt = internal.PlayDisguisedChatMessage_763_0

// ID=0x19
type PlayEntityEventPkt = internal.PlayEntityEvent_763_0

// ID=0x1a
type PlayExplosionPkt = internal.PlayExplosion_763_0

// ID=0x1b
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x1c
type PlayGameEventPkt = internal.PlayGameEvent_763_0

// ID=0x1d
type PlayOpenHorseScreenPkt = internal.PlayOpenHorseScreen_763_0

// ID=0x1e
type PlayInitializeWorldBorderPkt = internal.PlayInitializeWorldBorder_763_0

// ID=0x1f
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x20
type PlayChunkDataAndUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name               | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
	 * |-----------|-------|----------|--------------------------|--------------------------|------------|---------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x20      | Play  | Client   | Chunk X                  | Chunk X                  | Int        | Int                 | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x20      | Play  | Client   | Chunk Z                  | Chunk Z                  | Int        | Int                 | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x20      | Play  | Client   | Heightmaps               | Heightmaps               | NBT        | NBT                 | Compound containing one long array named MOTION_BLOCKING, which is a heightmap for the highest solid block at each position in the chunk (as a compacted long array with 256 entries, with the number of bits per entry varying depending on the world's height, defined by the formula ceil(log2(height + 1))). The Notchian server also adds a WORLD_SURFACE long array, the purpose of which is unknown, but it's not required for the chunk to be accepted. |
	 * | 0x20      | Play  | Client   | Size                     | Size                     | VarInt     | VarInt              | Size of Data in bytes                                                                                                                                                                                                                                                                                                                                                                                                                                           |
	 * | 0x20      | Play  | Client   | Data                     | Data                     | Byte array | Byte array          | See data structure in Chunk Format                                                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x20      | Play  | Client   | Number of block entities | Number of block entities | VarInt     | VarInt              | Number of elements in the following array                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x20      | Play  | Client   | Block Entity             | Packed XZ                | Array      | Byte                | The packed section coordinates, calculated from ((blockX & 15) << 4) | (blockZ & 15)                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x20      | Play  | Client   | Block Entity             | Y                        | Array      | Short               | The height relative to the world                                                                                                                                                                                                                                                                                                                                                                                                                                |
	 * | 0x20      | Play  | Client   | Block Entity             | Type                     | Array      | VarInt              | The type of block entity                                                                                                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x20      | Play  | Client   | Block Entity             | Data                     | Array      | NBT                 | The block entity's data, without the X, Y, and Z values                                                                                                                                                                                                                                                                                                                                                                                                         |
	 * | 0x20      | Play  | Client   | Trust Edges              | Trust Edges              | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x20      | Play  | Client   | Sky Light Mask           | Sky Light Mask           | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Sky Light array below.  The least significant bit is for blocks 16 blocks to 1 block below the min world height (one section below the world), while the most significant bit covers blocks 1 to 16 blocks above the max world height (one section above the world).                                          |
	 * | 0x20      | Play  | Client   | Block Light Mask         | Block Light Mask         | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Block Light array below.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                 |
	 * | 0x20      | Play  | Client   | Empty Sky Light Mask     | Empty Sky Light Mask     | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Sky Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                    |
	 * | 0x20      | Play  | Client   | Empty Block Light Mask   | Empty Block Light Mask   | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Block Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                                                                  |
	 * | 0x20      | Play  | Client   | Sky Light array count    | Sky Light array count    | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Sky Light Mask                                                                                                                                                                                                                                                                                                                                                                 |
	 * | 0x20      | Play  | Client   | Sky Light arrays         | Length                   | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x20      | Play  | Client   | Sky Light arrays         | Sky Light array          | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value. Indexed ((y<<8) | (z<<4) | x) / 2  If there's a remainder, masked 0xF0 else 0x0F.                                                                                                                                                                                                                                                |
	 * | 0x20      | Play  | Client   | Block Light array count  | Block Light array count  | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Block Light Mask                                                                                                                                                                                                                                                                                                                                                               |
	 * | 0x20      | Play  | Client   | Block Light arrays       | Length                   | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                                                            |
	 * | 0x20      | Play  | Client   | Block Light arrays       | Block Light array        | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value. Indexed ((y<<8) | (z<<4) | x) / 2  If there's a remainder, masked 0xF0 else 0x0F.                                                                                                                                                                                                                                              |
	 * 
	 */
}

// ID=0x21
type PlayWorldEventPkt = internal.PlayWorldEvent_763_0

// ID=0x22
type PlayParticlePkt = internal.PlayParticle_763_0

// ID=0x23
type PlayUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name              | Field Name              | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                  |
	 * |-----------|-------|----------|-------------------------|-------------------------|------------|---------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x23      | Play  | Client   | Chunk X                 | Chunk X                 | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x23      | Play  | Client   | Chunk Z                 | Chunk Z                 | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                                                        |
	 * | 0x23      | Play  | Client   | Trust Edges             | Trust Edges             | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                                                          |
	 * | 0x23      | Play  | Client   | Sky Light Mask          | Sky Light Mask          | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Sky Light array below.  The least significant bit is for blocks 16 blocks to 1 block below the min world height (one section below the world), while the most significant bit covers blocks 1 to 16 blocks above the max world height (one section above the world). |
	 * | 0x23      | Play  | Client   | Block Light Mask        | Block Light Mask        | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has data in the Block Light array below.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                        |
	 * | 0x23      | Play  | Client   | Empty Sky Light Mask    | Empty Sky Light Mask    | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Sky Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                           |
	 * | 0x23      | Play  | Client   | Empty Block Light Mask  | Empty Block Light Mask  | BitSet     | BitSet              | BitSet containing bits for each section in the world + 2.  Each set bit indicates that the corresponding 16×16×16 chunk section has all zeros for its Block Light data.  The order of bits is the same as in Sky Light Mask.                                                                                                                                                                                         |
	 * | 0x23      | Play  | Client   | Sky Light array count   | Sky Light array count   | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Sky Light Mask                                                                                                                                                                                                                                                                                                                        |
	 * | 0x23      | Play  | Client   | Sky Light arrays        | Length                  | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x23      | Play  | Client   | Sky Light arrays        | Sky Light array         | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                                                         |
	 * | 0x23      | Play  | Client   | Block Light array count | Block Light array count | VarInt     | VarInt              | Number of entries in the following array; should match the number of bits set in Block Light Mask                                                                                                                                                                                                                                                                                                                      |
	 * | 0x23      | Play  | Client   | Block Light arrays      | Length                  | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x23      | Play  | Client   | Block Light arrays      | Block Light array       | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                                                       |
	 * 
	 */
}

// ID=0x24
type PlayLoginPkt = internal.PlayLogin_761_2

// ID=0x25
type PlayMapDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name       | Field Type                      | Field Type                      | Notes                                                                                                      |
	 * |-----------|-------|----------|------------|------------------|---------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x25      | Play  | Client   | Map ID     | Map ID           | VarInt                          | VarInt                          | Map ID of the map being modified                                                                           |
	 * | 0x25      | Play  | Client   | Scale      | Scale            | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x25      | Play  | Client   | Locked     | Locked           | Boolean                         | Boolean                         | True if the map has been locked in a cartography table                                                     |
	 * | 0x25      | Play  | Client   | Has Icons  | Has Icons        | Boolean                         | Boolean                         |                                                                                                            |
	 * | 0x25      | Play  | Client   | Icon Count | Icon Count       | Optional VarInt                 | Optional VarInt                 | Number of elements in the following array. Only present if previous Boolean is true.                       |
	 * | 0x25      | Play  | Client   | Icon       | Type             | Optional Array                  | VarInt Enum                     | See below                                                                                                  |
	 * | 0x25      | Play  | Client   | Icon       | X                | Optional Array                  | Byte                            | Map coordinates: -128 for furthest left, +127 for furthest right                                           |
	 * | 0x25      | Play  | Client   | Icon       | Z                | Optional Array                  | Byte                            | Map coordinates: -128 for highest, +127 for lowest                                                         |
	 * | 0x25      | Play  | Client   | Icon       | Direction        | Optional Array                  | Byte                            | 0-15                                                                                                       |
	 * | 0x25      | Play  | Client   | Icon       | Has Display Name | Optional Array                  | Boolean                         |                                                                                                            |
	 * | 0x25      | Play  | Client   | Icon       | Display Name     | Optional Array                  | Optional Chat                   | Only present if previous Boolean is true                                                                   |
	 * | 0x25      | Play  | Client   | Columns    | Columns          | Unsigned Byte                   | Unsigned Byte                   | Number of columns updated                                                                                  |
	 * | 0x25      | Play  | Client   | Rows       | Rows             | Optional Unsigned Byte          | Optional Unsigned Byte          | Only if Columns is more than 0; number of rows updated                                                     |
	 * | 0x25      | Play  | Client   | X          | X                | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column                                         |
	 * | 0x25      | Play  | Client   | Z          | Z                | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row                                           |
	 * | 0x25      | Play  | Client   | Length     | Length           | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array                                              |
	 * | 0x25      | Play  | Client   | Data       | Data             | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format                                                        |
	 * 
	 */
}

// ID=0x26
type PlayMerchantOffersPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name                   | Field Type | Field Type | Notes                                                                                                                                                                             |
	 * |-----------|-------|----------|---------------------|------------------------------|------------|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x26      | Play  | Client   | Window ID           | Window ID                    | VarInt     | VarInt     | The ID of the window that is open; this is an int rather than a byte.                                                                                                             |
	 * | 0x26      | Play  | Client   | Size                | Size                         | VarInt     | VarInt     | The number of trades in the following array.                                                                                                                                      |
	 * | 0x26      | Play  | Client   | Trades              | Input item 1                 | Array      | Slot       | The first item the player has to supply for this villager trade. The count of the item stack is the default "price" of this trade.                                                |
	 * | 0x26      | Play  | Client   | Trades              | Output item                  | Array      | Slot       | The item the player will receive from this villager trade.                                                                                                                        |
	 * | 0x26      | Play  | Client   | Trades              | Input item 2                 | Array      | Slot       | The second item the player has to supply for this villager trade. May be an empty slot.                                                                                           |
	 * | 0x26      | Play  | Client   | Trades              | Trade disabled               | Array      | Boolean    | True if the trade is disabled; false if the trade is enabled.                                                                                                                     |
	 * | 0x26      | Play  | Client   | Trades              | Number of trade uses         | Array      | Int        | Number of times the trade has been used so far. If equal to the maximum number of trades, the client will display a red X.                                                        |
	 * | 0x26      | Play  | Client   | Trades              | Maximum number of trade uses | Array      | Int        | Number of times this trade can be used before it's exhausted.                                                                                                                     |
	 * | 0x26      | Play  | Client   | Trades              | XP                           | Array      | Int        | Amount of XP the villager will earn each time the trade is used.                                                                                                                  |
	 * | 0x26      | Play  | Client   | Trades              | Special Price                | Array      | Int        | Can be zero or negative. The number is added to the price when an item is discounted due to player reputation or other effects.                                                   |
	 * | 0x26      | Play  | Client   | Trades              | Price Multiplier             | Array      | Float      | Can be low (0.05) or high (0.2). Determines how much demand, player reputation, and temporary effects will adjust the price.                                                      |
	 * | 0x26      | Play  | Client   | Trades              | Demand                       | Array      | Int        | If positive, causes the price to increase. Negative values seem to be treated the same as zero.                                                                                   |
	 * | 0x26      | Play  | Client   | Villager level      | Villager level               | VarInt     | VarInt     | Appears on the trade GUI; meaning comes from the translation key merchant.level. + level.
	 * 1: Novice, 2: Apprentice, 3: Journeyman, 4: Expert, 5: Master.                          |
	 * | 0x26      | Play  | Client   | Experience          | Experience                   | VarInt     | VarInt     | Total experience for this villager (always 0 for the wandering trader).                                                                                                           |
	 * | 0x26      | Play  | Client   | Is regular villager | Is regular villager          | Boolean    | Boolean    | True if this is a regular villager; false for the wandering trader.  When false, hides the villager level and some other GUI elements.                                            |
	 * | 0x26      | Play  | Client   | Can restock         | Can restock                  | Boolean    | Boolean    | True for regular villagers and false for the wandering trader. If true, the "Villagers restock up to two times per day." message is displayed when hovering over disabled trades. |
	 * 
	 */
}

// ID=0x27
type PlayUpdateEntityPositionPkt = internal.PlayUpdateEntityPosition_763_0

// ID=0x28
type PlayUpdateEntityPositionAndRotationPkt = internal.PlayUpdateEntityPositionAndRotation_763_0

// ID=0x29
type PlayUpdateEntityRotationPkt = internal.PlayUpdateEntityRotation_763_0

// ID=0x2a
type PlayMoveVehicleClientPkt = internal.PlayMoveVehicle_763_0

// ID=0x2b
type PlayOpenBookPkt = internal.PlayOpenBook_763_0

// ID=0x2c
type PlayOpenScreenPkt = internal.PlayOpenScreen_763_0

// ID=0x2d
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x2e
type PlayPingPkt = internal.PlayPing_763_0

// ID=0x2f
type PlayPlaceGhostRecipePkt = internal.PlayPlaceGhostRecipe_763_0

// ID=0x30
type PlayPlayerAbilitiesClientPkt = internal.PlayPlayerAbilities_761_2

// ID=0x31
type PlayPlayerChatMessagePkt struct {
	/*
	 * | Packet ID | State | Bound To | Sector            | Field Name                  | Field Name                  | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * |-----------|-------|----------|-------------------|-----------------------------|-----------------------------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x31      | Play  | Client   | Header            | Sender                      | Sender                      | UUID                | Used by the Notchian client for the disableChat launch option. Setting both longs to 0 will always display the message regardless of the setting.                                                                                                                                                                                                                                                                                                                                                                                                     |
	 * | 0x31      | Play  | Client   | Header            | Index                       | Index                       | VarInt              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x31      | Play  | Client   | Header            | Message Signature Present   | Message Signature Present   | Boolean             | States if a message signature is present                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x31      | Play  | Client   | Header            | Message Signature bytes     | Message Signature bytes     | Optional Byte Array | Only present if Message Signature Present is true. Cryptography, the signature consists of the Sender UUID, Session UUID from the Player Session packet, Index, Salt, Timestamp in epoch seconds, the length of the original chat content, the original content itself, the length of Previous Messages, and all of the Previous message signatures. These values are hashed with SHA-256 and signed using the RSA cryptosystem. Modifying any of these values in the packet will cause this signature to fail. This buffer is always 256 bytes long. |
	 * | 0x31      | Play  | Client   | Body              | Message                     | Message                     | String              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x31      | Play  | Client   | Body              | Timestamp                   | Timestamp                   | Long                | Represents the time the message was signed as milliseconds since the epoch, used to check if the message was received within 2 minutes of it being sent.                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x31      | Play  | Client   | Body              | Salt                        | Salt                        | Long                | Cryptography, used for validating the message signature.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x31      | Play  | Client   | Previous Messages | Total Previous Messages     | Total Previous Messages     | VarInt              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x31      | Play  | Client   | Previous Messages | Array                       | Message ID                  | VarInt              | The message Id, used for validating message signature.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
	 * | 0x31      | Play  | Client   | Previous Messages | Array                       | Signature                   | Optional Byte Array | The previous message's signature. Contains the same type of data as Message Signature bytes above.                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
	 * | 0x31      | Play  | Client   | Other             | Unsigned Content Present    | Unsigned Content Present    | Boolean             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x31      | Play  | Client   | Other             | Unsigned Content            | Unsigned Content            | Optional Chat       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x31      | Play  | Client   | Other             | Filter Type                 | Filter Type                 | Enum VarInt         | If the message has been filtered                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
	 * | 0x31      | Play  | Client   | Other             | Filter Type Bits            | Filter Type Bits            | Optional BitSet     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x31      | Play  | Client   | Network target    | Chat Type                   | Chat Type                   | VarInt              | The chat type from the Login (play) packet used for this message                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
	 * | 0x31      | Play  | Client   | Network target    | Network name                | Network name                | Chat                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x31      | Play  | Client   | Network target    | Network target name present | Network target name present | Boolean             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x31      | Play  | Client   | Network target    | Network target name         | Network target name         | Optional Chat       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * 
	 */
}

// ID=0x32
type PlayEndCombatPkt = internal.PlayEndCombat_762_1

// ID=0x33
type PlayEnterCombatPkt = internal.PlayEnterCombat_763_0

// ID=0x34
type PlayCombatDeathPkt = internal.PlayCombatDeath_762_1

// ID=0x35
type PlayPlayerInfoRemovePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type    | Notes                                      |
	 * |-----------|-------|----------|-------------------|-------------------|---------------|--------------------------------------------|
	 * | 0x35      | Play  | Client   | Number of Players | Number of Players | VarInt        | Number of elements in the following array. |
	 * | 0x35      | Play  | Client   | Player            | Player Id         | Array of UUID | UUIDs of players to remove.                |
	 * 
	 */
}

// ID=0x36
type PlayPlayerInfoUpdatePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name                         | Field Name                | Field Name                | Field Type | Field Type    | Field Type              | Notes                                                                                                          |
	 * |-----------|-------|----------|-------------------|------------------------------------|---------------------------|---------------------------|------------|---------------|-------------------------|----------------------------------------------------------------------------------------------------------------|
	 * | 0x36      | Play  | Client   |                   |                                    |                           |                           |            |               |                         |                                                                                                                |
	 * | 0x36      | Play  | Client   | Actions           | Actions                            | Actions                   | Actions                   | Byte       | Byte          | Byte                    | Bit Mask. The actions to process. This must have a bit set for every action below, whether it's true or false. |
	 * | 0x36      | Play  | Client   | Number Of Actions | Number Of Actions                  | Number Of Actions         | Number Of Actions         | VarInt     | VarInt        | VarInt                  | Number of elements in the following array.                                                                     |
	 * | 0x36      | Play  | Client   | Action            | UUID                               | UUID                      | UUID                      | Array      | UUID          | UUID                    | The player UUID                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Action                             | Field Name                | Field Name                | Array      |               |                         |                                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 0: add player          | Name                      | Name                      | Array      | String (16)   | String (16)             |                                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 0: add player          | Number Of Properties      | Number Of Properties      | Array      | VarInt        | VarInt                  | Number of elements in the following array.                                                                     |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 0: add player          | Property                  | Name                      | Array      | Array         | String (32767)          |                                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 0: add player          | Property                  | Value                     | Array      | Array         | String (32767)          |                                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 0: add player          | Property                  | Is Signed                 | Array      | Array         | Boolean                 |                                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 0: add player          | Property                  | Signature                 | Array      | Array         | Optional String (32767) | Only if Is Signed is true.                                                                                     |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Has Signature Data        | Has Signature Data        | Array      | Boolean       | Boolean                 |                                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Chat session ID           | Chat session ID           | Array      | UUID          | UUID                    | Only send if Has Signature Data is true.                                                                       |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Public key expiry time    | Public key expiry time    | Array      | Long          | Long                    | Key expiry time, as a UNIX timestamp in milliseconds. Only send if Has Signature Data is true.                 |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Encoded public key size   | Encoded public key size   | Array      | VarInt        | VarInt                  | Size of the following array. Only send if Has Signature Data is true.                                          |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Encoded public key        | Encoded public key        | Array      | Byte Array    | Byte Array              | The player's public key, in bytes. Only send if Has Signature Data is true.                                    |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Public key signature size | Public key signature size | Array      | VarInt        | VarInt                  | Size of the following array. Only send if Has Signature Data is true.                                          |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Public key signature      | Public key signature      | Array      | Byte Array    | Byte Array              | The public key's digital signature. Only send if Has Signature Data is true.                                   |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 2: update gamemode     | Gamemode                  | Gamemode                  | Array      | VarInt        | VarInt                  |                                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 3: update listed       | Listed                    | Listed                    | Array      | Boolean       | Boolean                 | Whether the player should be listed on the player list.                                                        |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 4: update latency      | Ping                      | Ping                      | Array      | VarInt        | VarInt                  | Measured in milliseconds.                                                                                      |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 5: update display name | Has Display Name          | Has Display Name          | Array      | Boolean       | Boolean                 |                                                                                                                |
	 * | 0x36      | Play  | Client   | Action            | Actions bit 5: update display name | Display Name              | Display Name              | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true.                                                                         |
	 * 
	 */
}

// ID=0x37
type PlayLookAtPkt = internal.PlayLookAt_763_0

// ID=0x38
type PlaySynchronizePlayerPositionPkt = internal.PlaySynchronizePlayerPosition_761_1

// ID=0x39
type PlayUpdateRecipeBookPkt = internal.PlayUpdateRecipeBook_763_0

// ID=0x3a
type PlayRemoveEntitiesPkt = internal.PlayRemoveEntities_763_0

// ID=0x3b
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_763_0

// ID=0x3c
type PlayResourcePackClientPkt = internal.PlayResourcePack_761_2

// ID=0x3d
type PlayRespawnPkt = internal.PlayRespawn_761_2

// ID=0x3e
type PlaySetHeadRotationPkt = internal.PlaySetHeadRotation_763_0

// ID=0x3f
type PlayUpdateSectionBlocksPkt = internal.PlayUpdateSectionBlocks_762_1

// ID=0x40
type PlaySelectAdvancementsTabPkt = internal.PlaySelectAdvancementsTab_763_0

// ID=0x41
type PlayServerDataPkt struct {
	HasMOTD Bool // Boolean
	MOTD Object // Optional Chat
	HasIcon Bool // Boolean
	/* Icon PNG base64 String */
	Icon String // Optional String (32767)
	EnforcesSecureChat Bool // Boolean
}

var _ Packet = (*PlayServerDataPkt)(nil)

func (p PlayServerDataPkt)Encode(b *PacketBuilder){
	b.Bool(p.HasMOTD)
	b.JSON(p.MOTD)
	b.Bool(p.HasIcon)
	b.String(p.Icon)
	b.Bool(p.EnforcesSecureChat)
}

func (p *PlayServerDataPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.HasMOTD, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.MOTD); err != nil {
		return
	}
	if p.HasIcon, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Icon, ok = r.String(); !ok {
		return io.EOF
	}
	if p.EnforcesSecureChat, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// ID=0x42
type PlaySetActionBarTextPkt = internal.PlaySetActionBarText_763_0

// ID=0x43
type PlaySetBorderCenterPkt = internal.PlaySetBorderCenter_763_0

// ID=0x44
type PlaySetBorderLerpSizePkt = internal.PlaySetBorderLerpSize_763_0

// ID=0x45
type PlaySetBorderSizePkt = internal.PlaySetBorderSize_763_0

// ID=0x46
type PlaySetBorderWarningDelayPkt = internal.PlaySetBorderWarningDelay_763_0

// ID=0x47
type PlaySetBorderWarningDistancePkt = internal.PlaySetBorderWarningDistance_763_0

// ID=0x48
type PlaySetCameraPkt = internal.PlaySetCamera_763_0

// ID=0x49
type PlaySetHeldItemClientPkt = internal.PlaySetHeldItem_761_2

// ID=0x4a
type PlaySetCenterChunkPkt = internal.PlaySetCenterChunk_763_0

// ID=0x4b
type PlaySetRenderDistancePkt = internal.PlaySetRenderDistance_763_0

// ID=0x4c
type PlaySetDefaultSpawnPositionPkt = internal.PlaySetDefaultSpawnPosition_763_0

// ID=0x4d
type PlayDisplayObjectivePkt = internal.PlayDisplayObjective_763_0

// ID=0x4e
type PlaySetEntityMetadataPkt = internal.PlaySetEntityMetadata_763_0

// ID=0x4f
type PlayLinkEntitiesPkt = internal.PlayLinkEntities_763_0

// ID=0x50
type PlaySetEntityVelocityPkt = internal.PlaySetEntityVelocity_763_0

// ID=0x51
type PlaySetEquipmentPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name | Field Type | Field Type | Notes                                                                                                                                                                                                                          |
	 * |-----------|-------|----------|------------|------------|------------|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x51      | Play  | Client   | Entity ID  | Entity ID  | VarInt     | VarInt     | Entity's ID.                                                                                                                                                                                                                   |
	 * | 0x51      | Play  | Client   | Equipment  | Slot       | Array      | Byte Enum  | Equipment slot. 0: main hand, 1: off hand, 2–5: armor slot (2: boots, 3: leggings, 4: chestplate, 5: helmet).  Also has the top bit set if another entry follows, and otherwise unset if this is the last item in the array. |
	 * | 0x51      | Play  | Client   | Equipment  | Item       | Array      | Slot       |                                                                                                                                                                                                                                |
	 * 
	 */
}

// ID=0x52
type PlaySetExperiencePkt = internal.PlaySetExperience_763_0

// ID=0x53
type PlaySetHealthPkt = internal.PlaySetHealth_763_0

// ID=0x54
type PlayUpdateObjectivesPkt = internal.PlayUpdateObjectives_763_0

// ID=0x55
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x56
type PlayUpdateTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                   | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|------------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x56      | Play  | Client   | Team Name                    | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x56      | Play  | Client   | Mode                         | Mode                | Byte                 | Determines the layout of the remaining packet.                                                                         |
	 * | 0x56      | Play  | Client   | 0: create team               | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x56      | Play  | Client   | 0: create team               | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team.                                     |
	 * | 0x56      | Play  | Client   | 0: create team               | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never.                                                                      |
	 * | 0x56      | Play  | Client   | 0: create team               | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never.                                                                            |
	 * | 0x56      | Play  | Client   | 0: create team               | Team Color          | VarInt Enum          | Used to color the name of players on the team; see below.                                                              |
	 * | 0x56      | Play  | Client   | 0: create team               | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team.                                                      |
	 * | 0x56      | Play  | Client   | 0: create team               | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team.                                                       |
	 * | 0x56      | Play  | Client   | 0: create team               | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x56      | Play  | Client   | 0: create team               | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x56      | Play  | Client   | 1: remove team               | no fields           | no fields            |                                                                                                                        |
	 * | 0x56      | Play  | Client   | 2: update team info          | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x56      | Play  | Client   | 2: update team info          | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team.                                    |
	 * | 0x56      | Play  | Client   | 2: update team info          | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x56      | Play  | Client   | 2: update team info          | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x56      | Play  | Client   | 2: update team info          | Team Color          | VarInt Enum          | Used to color the name of players on the team; see below.                                                              |
	 * | 0x56      | Play  | Client   | 2: update team info          | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team.                                                      |
	 * | 0x56      | Play  | Client   | 2: update team info          | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team.                                                       |
	 * | 0x56      | Play  | Client   | 3: add entities to team      | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x56      | Play  | Client   | 3: add entities to team      | Entities            | Array of String (40) | Identifiers for the added entities.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x56      | Play  | Client   | 4: remove entities from team | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x56      | Play  | Client   | 4: remove entities from team | Entities            | Array of String (40) | Identifiers for the removed entities.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

// ID=0x57
type PlayUpdateScorePkt = internal.PlayUpdateScore_763_0

// ID=0x58
type PlaySetSimulationDistancePkt = internal.PlaySetSimulationDistance_763_0

// ID=0x59
type PlaySetSubtitleTextPkt = internal.PlaySetSubtitleText_763_0

// ID=0x5a
type PlayUpdateTimePkt = internal.PlayUpdateTime_763_0

// ID=0x5b
type PlaySetTitleTextPkt = internal.PlaySetTitleText_763_0

// ID=0x5c
type PlaySetTitleAnimationTimesPkt = internal.PlaySetTitleAnimationTimes_763_0

// ID=0x5d
type PlayEntitySoundEffectPkt = internal.PlayEntitySoundEffect_763_0

// ID=0x5e
type PlaySoundEffectPkt = internal.PlaySoundEffect_763_0

// ID=0x5f
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x60
type PlaySystemChatMessagePkt = internal.PlaySystemChatMessage_763_0

// ID=0x61
type PlaySetTabListHeaderAndFooterPkt = internal.PlaySetTabListHeaderAndFooter_763_0

// ID=0x62
type PlayTagQueryResponsePkt = internal.PlayTagQueryResponse_763_0

// ID=0x63
type PlayPickupItemPkt = internal.PlayPickupItem_763_0

// ID=0x64
type PlayTeleportEntityPkt = internal.PlayTeleportEntity_763_0

// ID=0x65
type PlayUpdateAdvancementsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name    | Field Type          | Field Type           | Notes                                                       |
	 * |-----------|-------|----------|---------------------|---------------|---------------------|----------------------|-------------------------------------------------------------|
	 * | 0x65      | Play  | Client   | Reset/Clear         | Reset/Clear   | Boolean             | Boolean              | Whether to reset/clear the current advancements.            |
	 * | 0x65      | Play  | Client   | Mapping size        | Mapping size  | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x65      | Play  | Client   | Advancement mapping | Key           | Array               | Identifier           | The identifier of the advancement.                          |
	 * | 0x65      | Play  | Client   | Advancement mapping | Value         | Array               | Advancement          | See below                                                   |
	 * | 0x65      | Play  | Client   | List size           | List size     | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x65      | Play  | Client   | Identifiers         | Identifiers   | Array of Identifier | Array of Identifier  | The identifiers of the advancements that should be removed. |
	 * | 0x65      | Play  | Client   | Progress size       | Progress size | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x65      | Play  | Client   | Progress mapping    | Key           | Array               | Identifier           | The identifier of the advancement.                          |
	 * | 0x65      | Play  | Client   | Progress mapping    | Value         | Array               | Advancement progress | See below.                                                  |
	 * 
	 */
}

// ID=0x66
type PlayUpdateAttributesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x66      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x66      | Play  | Client   | Number Of Properties | Number Of Properties | VarInt     | VarInt                 | Number of elements in the following array.            |
	 * | 0x66      | Play  | Client   | Property             | Key                  | Array      | Identifier             | See below.                                            |
	 * | 0x66      | Play  | Client   | Property             | Value                | Array      | Double                 | See below.                                            |
	 * | 0x66      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array.            |
	 * | 0x66      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x67
type PlayFeatureFlagsPkt = internal.PlayFeatureFlags_763_0

// ID=0x68
type PlayEntityEffectPkt = internal.PlayEntityEffect_763_0

// ID=0x69
type PlayUpdateRecipesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name  | Field Type | Field Type | Notes                                      |
	 * |-----------|-------|----------|-------------|-------------|------------|------------|--------------------------------------------|
	 * | 0x69      | Play  | Client   | Num Recipes | Num Recipes | VarInt     | VarInt     | Number of elements in the following array. |
	 * | 0x69      | Play  | Client   | Recipe      | Type        | Array      | Identifier | The recipe type, see below.                |
	 * | 0x69      | Play  | Client   | Recipe      | Recipe ID   | Array      | Identifier |                                            |
	 * | 0x69      | Play  | Client   | Recipe      | Data        | Array      | Varies     | Additional data for the recipe.            |
	 * 
	 */
}

// ID=0x6a
type PlayUpdateTagsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type | Field Type  | Notes                                                                                                                                        |
	 * |-----------|-------|----------|---------------------|---------------------|------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x6A      | Play  | Client   | Length of the array | Length of the array | VarInt     | VarInt      |                                                                                                                                              |
	 * | 0x6A      | Play  | Client   | Array of tags       | Tag type            | Array      | Identifier  | Tag identifier (Vanilla required tags are minecraft:block, minecraft:item, minecraft:fluid, minecraft:entity_type, and minecraft:game_event) |
	 * | 0x6A      | Play  | Client   | Array of tags       | Array of Tag        | Array      | (See below) |                                                                                                                                              |
	 * 
	 */
}

// ======== END play ========
