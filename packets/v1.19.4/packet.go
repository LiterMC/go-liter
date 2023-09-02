
// Generated at 2023-09-01 20:57:33.569 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=18242
// Protocol: 762
// Protocol Name: 1.19.4

package packet_1_19_4

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
type PlayChangeDifficultyServerPkt = internal.PlayChangeDifficulty_763_1

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
type PlayPlayerSessionPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name           | Field Type | Notes                                                                                                                                                                            |
	 * |-----------|-------|----------|------------|----------------------|------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x06      | Play  | Server   | Session Id | Session Id           | UUID       |                                                                                                                                                                                  |
	 * | 0x06      | Play  | Server   | Public Key | Expires At           | Long       | The time the play session key expires in epoch milliseconds.                                                                                                                     |
	 * | 0x06      | Play  | Server   | Public Key | Public Key Length    | VarInt     | Length of the proceeding public key.                                                                                                                                             |
	 * | 0x06      | Play  | Server   | Public Key | Public Key           | Byte Array | A byte array of an X.509-encoded public key.                                                                                                                                     |
	 * | 0x06      | Play  | Server   | Public Key | Key Signature Length | VarInt     | Length of the proceeding key signature array.                                                                                                                                    |
	 * | 0x06      | Play  | Server   | Public Key | Key Signature        | Byte Array | The signature consists of the player UUID, the key expiration timestamp, and the public key data. These values are hashed using SHA-1 and signed using Mojang's private RSA key. |
	 * 
	 */
}

// ID=0x7
type PlayClientCommandPkt = internal.PlayClientCommand_763_0

// ID=0x8
type PlayClientInformationPkt = internal.PlayClientInformation_763_0

// ID=0x9
type PlayCommandSuggestionsRequestPkt = internal.PlayCommandSuggestionsRequest_763_0

// ID=0xa
type PlayClickContainerButtonPkt = internal.PlayClickContainerButton_763_0

// ID=0xb
type PlayClickContainerPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type    | Field Type    | Notes                                                                                                    |
	 * |-----------|-------|----------|---------------------|---------------------|---------------|---------------|----------------------------------------------------------------------------------------------------------|
	 * | 0x0B      | Play  | Server   | Window ID           | Window ID           | Unsigned Byte | Unsigned Byte | The ID of the window which was clicked. 0 for player inventory.                                          |
	 * | 0x0B      | Play  | Server   | State ID            | State ID            | VarInt        | VarInt        | The last recieved State ID from either a Set Container Slot or a Set Container Content packet            |
	 * | 0x0B      | Play  | Server   | Slot                | Slot                | Short         | Short         | The clicked slot number, see below.                                                                      |
	 * | 0x0B      | Play  | Server   | Button              | Button              | Byte          | Byte          | The button used in the click, see below.                                                                 |
	 * | 0x0B      | Play  | Server   | Mode                | Mode                | VarInt Enum   | VarInt Enum   | Inventory operation mode, see below.                                                                     |
	 * | 0x0B      | Play  | Server   | Length of the array | Length of the array | VarInt        | VarInt        |                                                                                                          |
	 * | 0x0B      | Play  | Server   | Array of slots      | Slot number         | Array         | Short         |                                                                                                          |
	 * | 0x0B      | Play  | Server   | Array of slots      | Slot data           | Array         | Slot          | New data for this slot                                                                                   |
	 * | 0x0B      | Play  | Server   | Carried item        | Carried item        | Slot          | Slot          | Item carried by the cursor. Has to be empty (item ID = -1) for drop mode, otherwise nothing will happen. |
	 * 
	 */
}

// ID=0xc
type PlayCloseContainerServerPkt = internal.PlayCloseContainer_763_0

// ID=0xd
type PlayPluginMessageServerPkt = internal.PlayPluginMessage_763_0

// ID=0xe
type PlayEditBookPkt = internal.PlayEditBook_763_0

// ID=0xf
type PlayQueryEntityTagPkt = internal.PlayQueryEntityTag_763_0

// ID=0x10
type PlayInteractPkt = internal.PlayInteract_763_0

// ID=0x11
type PlayJigsawGeneratePkt = internal.PlayJigsawGenerate_763_0

// ID=0x12
type PlayKeepAliveServerPkt = internal.PlayKeepAlive_763_0

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
type PlayMoveVehicleServerPkt = internal.PlayMoveVehicle_763_0

// ID=0x19
type PlayPaddleBoatPkt = internal.PlayPaddleBoat_763_0

// ID=0x1a
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x1b
type PlayPlaceRecipePkt = internal.PlayPlaceRecipe_763_0

// ID=0x1c
type PlayPlayerAbilitiesServerPkt = internal.PlayPlayerAbilities_763_1

// ID=0x1d
type PlayPlayerActionPkt = internal.PlayPlayerAction_763_0

// ID=0x1e
type PlayPlayerCommandPkt = internal.PlayPlayerCommand_763_0

// ID=0x1f
type PlayPlayerInputPkt = internal.PlayPlayerInput_763_0

// ID=0x20
type PlayPongPkt = internal.PlayPong_763_0

// ID=0x21
type PlayChangeRecipeBookSettingsPkt = internal.PlayChangeRecipeBookSettings_763_0

// ID=0x22
type PlaySetSeenRecipePkt = internal.PlaySetSeenRecipe_763_0

// ID=0x23
type PlayRenameItemPkt = internal.PlayRenameItem_763_0

// ID=0x24
type PlayResourcePackServerPkt = internal.PlayResourcePack_763_1

// ID=0x25
type PlaySeenAdvancementsPkt = internal.PlaySeenAdvancements_763_0

// ID=0x26
type PlaySelectTradePkt = internal.PlaySelectTrade_763_0

// ID=0x27
type PlaySetBeaconEffectPkt = internal.PlaySetBeaconEffect_763_0

// ID=0x28
type PlaySetHeldItemServerPkt = internal.PlaySetHeldItem_763_1

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
type PlayAwardStatisticsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name   | Field Type | Field Type | Notes                                      |
	 * |-----------|-------|----------|------------|--------------|------------|------------|--------------------------------------------|
	 * | 0x05      | Play  | Client   | Count      | Count        | VarInt     | VarInt     | Number of elements in the following array. |
	 * | 0x05      | Play  | Client   | Statistic  | Category ID  | Array      | VarInt     | See below.                                 |
	 * | 0x05      | Play  | Client   | Statistic  | Statistic ID | Array      | VarInt     | See below.                                 |
	 * | 0x05      | Play  | Client   | Statistic  | Value        | Array      | VarInt     | The amount to set it to.                   |
	 * 
	 */
}

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
type PlayBossBarPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name       | Field Name | Field Type    | Notes                                                                                                                                     |
	 * |-----------|-------|----------|------------------|------------|---------------|-------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0B      | Play  | Client   | UUID             | UUID       | UUID          | Unique ID for this bar.                                                                                                                   |
	 * | 0x0B      | Play  | Client   | Action           | Action     | VarInt Enum   | Determines the layout of the remaining packet.                                                                                            |
	 * | 0x0B      | Play  | Client   | Action           | Field Name |               |                                                                                                                                           |
	 * | 0x0B      | Play  | Client   | 0: add           | Title      | Chat          |                                                                                                                                           |
	 * | 0x0B      | Play  | Client   | 0: add           | Health     | Float         | From 0 to 1. Values greater than 1 do not crash a Notchian client, and start rendering part of a second health bar at around 1.5.         |
	 * | 0x0B      | Play  | Client   | 0: add           | Color      | VarInt Enum   | Color ID (see below).                                                                                                                     |
	 * | 0x0B      | Play  | Client   | 0: add           | Division   | VarInt Enum   | Type of division (see below).                                                                                                             |
	 * | 0x0B      | Play  | Client   | 0: add           | Flags      | Unsigned Byte | Bit mask. 0x1: should darken sky, 0x2: is dragon bar (used to play end music), 0x04: create fog (previously was also controlled by 0x02). |
	 * | 0x0B      | Play  | Client   | 1: remove        | no fields  | no fields     | Removes this boss bar.                                                                                                                    |
	 * | 0x0B      | Play  | Client   | 2: update health | Health     | Float         | as above                                                                                                                                  |
	 * | 0x0B      | Play  | Client   | 3: update title  | Title      | Chat          |                                                                                                                                           |
	 * | 0x0B      | Play  | Client   | 4: update style  | Color      | VarInt Enum   | Color ID (see below).                                                                                                                     |
	 * | 0x0B      | Play  | Client   | 4: update style  | Dividers   | VarInt Enum   | as above                                                                                                                                  |
	 * | 0x0B      | Play  | Client   | 5: update flags  | Flags      | Unsigned Byte | as above                                                                                                                                  |
	 * 
	 */
}

// ID=0xc
type PlayChangeDifficultyClientPkt = internal.PlayChangeDifficulty_763_0

// ID=0xd
type PlayChunkBiomesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name       | Field Name       | Field Type | Field Type | Notes                                                                |
	 * |-----------|-------|----------|------------------|------------------|------------|------------|----------------------------------------------------------------------|
	 * | 0x0D      | Play  | Client   |                  |                  |            |            |                                                                      |
	 * | 0x0D      | Play  | Client   | Number of chunks | Number of chunks | VarInt     | VarInt     | Number of elements in the following array                            |
	 * | 0x0D      | Play  | Client   | Chunk biome data | Chunk X          | Array      | Int        | Chunk coordinate (block coordinate divided by 16, rounded down)      |
	 * | 0x0D      | Play  | Client   | Chunk biome data | Chunk Z          | Array      | Int        | Chunk coordinate (block coordinate divided by 16, rounded down)      |
	 * | 0x0D      | Play  | Client   | Chunk biome data | Size             | Array      | VarInt     | Size of Data in bytes                                                |
	 * | 0x0D      | Play  | Client   | Chunk biome data | Data             | Array      | Byte array | Chunk data structure, with sections containing only the Biomes field |
	 * 
	 */
}

// ID=0xe
type PlayClearTitlesPkt = internal.PlayClearTitles_763_0

// ID=0xf
type PlayCommandSuggestionsResponsePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name  | Field Type | Field Type     | Notes                                                                                                                                                                                                  |
	 * |-----------|-------|----------|------------|-------------|------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0F      | Play  | Client   |            |             |            |                |                                                                                                                                                                                                        |
	 * | 0x0F      | Play  | Client   | ID         | ID          | VarInt     | VarInt         | Transaction ID.                                                                                                                                                                                        |
	 * | 0x0F      | Play  | Client   | Start      | Start       | VarInt     | VarInt         | Start of the text to replace.                                                                                                                                                                          |
	 * | 0x0F      | Play  | Client   | Length     | Length      | VarInt     | VarInt         | Length of the text to replace.                                                                                                                                                                         |
	 * | 0x0F      | Play  | Client   | Count      | Count       | VarInt     | VarInt         | Number of elements in the following array.                                                                                                                                                             |
	 * | 0x0F      | Play  | Client   | Matches    | Match       | Array      | String (32767) | One eligible value to insert, note that each command is sent separately instead of in a single string, hence the need for Count.  Note that for instance this doesn't include a leading / on commands. |
	 * | 0x0F      | Play  | Client   | Matches    | Has tooltip | Array      | Boolean        | True if the following is present.                                                                                                                                                                      |
	 * | 0x0F      | Play  | Client   | Matches    | Tooltip     | Array      | Optional Chat  | Tooltip to display; only present if previous boolean is true.                                                                                                                                          |
	 * 
	 */
}

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
type PlayDamageEventPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type      | Field Type      | Notes                                                                                                                                                                                                                                                                                                                                                                          |
	 * |-----------|-------|----------|---------------------|---------------------|-----------------|-----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x18      | Play  | Client   |                     |                     |                 |                 |                                                                                                                                                                                                                                                                                                                                                                                |
	 * | 0x18      | Play  | Client   | Entity ID           | Entity ID           | VarInt          | VarInt          | The ID of the entity taking damage                                                                                                                                                                                                                                                                                                                                             |
	 * | 0x18      | Play  | Client   | Source Type ID      | Source Type ID      | VarInt          | VarInt          | The ID of the type of damage taken                                                                                                                                                                                                                                                                                                                                             |
	 * | 0x18      | Play  | Client   | Source Cause ID     | Source Cause ID     | VarInt          | VarInt          | The ID + 1 of the entity responsible for the damage, if present. If not present, the value is 0                                                                                                                                                                                                                                                                                |
	 * | 0x18      | Play  | Client   | Source Direct ID    | Source Direct ID    | VarInt          | VarInt          | The ID + 1 of the entity that directly dealt the damage, if present. If not present, the value is 0. If this field is present:
	 * and damage was dealt indirectly, such as by the use of a projectile, this field will contain the ID of such projectile;
	 * and damage was dealt dirctly, such as by manually attacking, this field will contain the same value as Source Cause ID. |
	 * | 0x18      | Play  | Client   | Has Source Position | Has Source Position | Boolean         | Boolean         | Indicates the presence of the three following fields.
	 * The Notchian server sends the Source Position when the damage was dealt by the /damage command and a position was specified                                                                                                                                                                                              |
	 * | 0x18      | Play  | Client   | Source Position X   | Source Position X   | Optional Double | Optional Double | Only present if Has Source Position is true                                                                                                                                                                                                                                                                                                                                    |
	 * | 0x18      | Play  | Client   | Source Position Y   | Source Position Y   | Optional Double | Optional Double | Only present if Has Source Position is true                                                                                                                                                                                                                                                                                                                                    |
	 * | 0x18      | Play  | Client   | Source Position Z   | Source Position Z   | Optional Double | Optional Double | Only present if Has Source Position is true                                                                                                                                                                                                                                                                                                                                    |
	 * 
	 */
}

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
type PlayHurtAnimationPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name | Field Type | Field Type | Notes                                                             |
	 * |-----------|-------|----------|------------|------------|------------|------------|-------------------------------------------------------------------|
	 * | 0x21      | Play  | Client   |            |            |            |            |                                                                   |
	 * | 0x21      | Play  | Client   | Entity ID  | Entity ID  | VarInt     | VarInt     | The ID of the entity taking damage                                |
	 * | 0x21      | Play  | Client   | Yaw        | Yaw        | Float      | Float      | The direction the damage is coming from in relation to the entity |
	 * 
	 */
}

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
	DimensionNames []Identifier // Array of Identifier
	/* Represents certain registries that are sent from the server and are applied on the client. */
	RegistryCodec *NBTCompound // NBT Tag Compound
	/* Name of the dimension type being spawned into. */
	DimensionType Identifier // Identifier
	/* Name of the dimension being spawned into. */
	DimensionName Identifier // Identifier
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
	DeathDimensionName Identifier // Optional Identifier
	/* The location that the player died at. */
	DeathLocation Position // Optional Position
}

var _ Packet = (*PlayLoginPkt)(nil)

func (p PlayLoginPkt)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.Bool(p.IsHardcore)
	b.UByte(p.GameMode)
	b.Byte(p.PreviousGameMode)
	p.DimensionCount = len(p.DimensionNames)
	b.VarInt(p.DimensionCount)
	for _, v := range p.DimensionNames {
		b.Encode(v)
	}
	b.Encode(p.RegistryCodec)
	b.Encode(p.DimensionType)
	b.Encode(p.DimensionName)
	b.Long(p.HashedSeed)
	b.VarInt(p.MaxPlayers)
	b.VarInt(p.ViewDistance)
	b.VarInt(p.SimulationDistance)
	b.Bool(p.ReducedDebugInfo)
	b.Bool(p.EnableRespawnScreen)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
	b.Bool(p.HasDeathLocation)
	b.Encode(p.DeathDimensionName)
	b.Encode(p.DeathLocation)
}

func (p *PlayLoginPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	p.DimensionNames = make([]Identifier, p.DimensionCount)
	for i, _ := range p.DimensionNames {
		if err = p.DimensionNames[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if err = p.RegistryCodec.DecodeFrom(r); err != nil {
		return
	}
	if err = p.DimensionType.DecodeFrom(r); err != nil {
		return
	}
	if err = p.DimensionName.DecodeFrom(r); err != nil {
		return
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
	if err = p.DeathDimensionName.DecodeFrom(r); err != nil {
		return
	}
	if err = p.DeathLocation.DecodeFrom(r); err != nil {
		return
	}
}

// ID=0x29
type PlayMapDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name       | Field Type                      | Field Type                      | Notes                                                                                                      |
	 * |-----------|-------|----------|------------|------------------|---------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x29      | Play  | Client   | Map ID     | Map ID           | VarInt                          | VarInt                          | Map ID of the map being modified                                                                           |
	 * | 0x29      | Play  | Client   | Scale      | Scale            | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x29      | Play  | Client   | Locked     | Locked           | Boolean                         | Boolean                         | True if the map has been locked in a cartography table                                                     |
	 * | 0x29      | Play  | Client   | Has Icons  | Has Icons        | Boolean                         | Boolean                         |                                                                                                            |
	 * | 0x29      | Play  | Client   | Icon Count | Icon Count       | Optional VarInt                 | Optional VarInt                 | Number of elements in the following array. Only present if previous Boolean is true.                       |
	 * | 0x29      | Play  | Client   | Icon       | Type             | Optional Array                  | VarInt Enum                     | See below                                                                                                  |
	 * | 0x29      | Play  | Client   | Icon       | X                | Optional Array                  | Byte                            | Map coordinates: -128 for furthest left, +127 for furthest right                                           |
	 * | 0x29      | Play  | Client   | Icon       | Z                | Optional Array                  | Byte                            | Map coordinates: -128 for highest, +127 for lowest                                                         |
	 * | 0x29      | Play  | Client   | Icon       | Direction        | Optional Array                  | Byte                            | 0-15                                                                                                       |
	 * | 0x29      | Play  | Client   | Icon       | Has Display Name | Optional Array                  | Boolean                         |                                                                                                            |
	 * | 0x29      | Play  | Client   | Icon       | Display Name     | Optional Array                  | Optional Chat                   | Only present if previous Boolean is true                                                                   |
	 * | 0x29      | Play  | Client   | Columns    | Columns          | Unsigned Byte                   | Unsigned Byte                   | Number of columns updated                                                                                  |
	 * | 0x29      | Play  | Client   | Rows       | Rows             | Optional Unsigned Byte          | Optional Unsigned Byte          | Only if Columns is more than 0; number of rows updated                                                     |
	 * | 0x29      | Play  | Client   | X          | X                | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column                                         |
	 * | 0x29      | Play  | Client   | Z          | Z                | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row                                           |
	 * | 0x29      | Play  | Client   | Length     | Length           | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array                                              |
	 * | 0x29      | Play  | Client   | Data       | Data             | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format                                                        |
	 * 
	 */
}

// ID=0x2a
type PlayMerchantOffersPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name                   | Field Type | Field Type | Notes                                                                                                                                                                             |
	 * |-----------|-------|----------|---------------------|------------------------------|------------|------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x2A      | Play  | Client   | Window ID           | Window ID                    | VarInt     | VarInt     | The ID of the window that is open; this is an int rather than a byte.                                                                                                             |
	 * | 0x2A      | Play  | Client   | Size                | Size                         | VarInt     | VarInt     | The number of trades in the following array.                                                                                                                                      |
	 * | 0x2A      | Play  | Client   | Trades              | Input item 1                 | Array      | Slot       | The first item the player has to supply for this villager trade. The count of the item stack is the default "price" of this trade.                                                |
	 * | 0x2A      | Play  | Client   | Trades              | Output item                  | Array      | Slot       | The item the player will receive from this villager trade.                                                                                                                        |
	 * | 0x2A      | Play  | Client   | Trades              | Input item 2                 | Array      | Slot       | The second item the player has to supply for this villager trade. May be an empty slot.                                                                                           |
	 * | 0x2A      | Play  | Client   | Trades              | Trade disabled               | Array      | Boolean    | True if the trade is disabled; false if the trade is enabled.                                                                                                                     |
	 * | 0x2A      | Play  | Client   | Trades              | Number of trade uses         | Array      | Int        | Number of times the trade has been used so far. If equal to the maximum number of trades, the client will display a red X.                                                        |
	 * | 0x2A      | Play  | Client   | Trades              | Maximum number of trade uses | Array      | Int        | Number of times this trade can be used before it's exhausted.                                                                                                                     |
	 * | 0x2A      | Play  | Client   | Trades              | XP                           | Array      | Int        | Amount of XP the villager will earn each time the trade is used.                                                                                                                  |
	 * | 0x2A      | Play  | Client   | Trades              | Special Price                | Array      | Int        | Can be zero or negative. The number is added to the price when an item is discounted due to player reputation or other effects.                                                   |
	 * | 0x2A      | Play  | Client   | Trades              | Price Multiplier             | Array      | Float      | Can be low (0.05) or high (0.2). Determines how much demand, player reputation, and temporary effects will adjust the price.                                                      |
	 * | 0x2A      | Play  | Client   | Trades              | Demand                       | Array      | Int        | If positive, causes the price to increase. Negative values seem to be treated the same as zero.                                                                                   |
	 * | 0x2A      | Play  | Client   | Villager level      | Villager level               | VarInt     | VarInt     | Appears on the trade GUI; meaning comes from the translation key merchant.level. + level.
	 * 1: Novice, 2: Apprentice, 3: Journeyman, 4: Expert, 5: Master.                          |
	 * | 0x2A      | Play  | Client   | Experience          | Experience                   | VarInt     | VarInt     | Total experience for this villager (always 0 for the wandering trader).                                                                                                           |
	 * | 0x2A      | Play  | Client   | Is regular villager | Is regular villager          | Boolean    | Boolean    | True if this is a regular villager; false for the wandering trader.  When false, hides the villager level and some other GUI elements.                                            |
	 * | 0x2A      | Play  | Client   | Can restock         | Can restock                  | Boolean    | Boolean    | True for regular villagers and false for the wandering trader. If true, the "Villagers restock up to two times per day." message is displayed when hovering over disabled trades. |
	 * 
	 */
}

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
type PlayPlayerAbilitiesClientPkt = internal.PlayPlayerAbilities_763_0

// ID=0x35
type PlayPlayerChatMessagePkt struct {
	/*
	 * | Packet ID | State | Bound To | Sector            | Field Name                  | Field Name                  | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
	 * |-----------|-------|----------|-------------------|-----------------------------|-----------------------------|---------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x35      | Play  | Client   | Header            | Sender                      | Sender                      | UUID                | Used by the Notchian client for the disableChat launch option. Setting both longs to 0 will always display the message regardless of the setting.                                                                                                                                                                                                                                                                                                                                                                                                     |
	 * | 0x35      | Play  | Client   | Header            | Index                       | Index                       | VarInt              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | Header            | Message Signature Present   | Message Signature Present   | Boolean             | States if a message signature is present                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | Header            | Message Signature bytes     | Message Signature bytes     | Optional Byte Array | Only present if Message Signature Present is true. Cryptography, the signature consists of the Sender UUID, Session UUID from the Player Session packet, Index, Salt, Timestamp in epoch seconds, the length of the original chat content, the original content itself, the length of Previous Messages, and all of the Previous message signatures. These values are hashed with SHA-256 and signed using the RSA cryptosystem. Modifying any of these values in the packet will cause this signature to fail. This buffer is always 256 bytes long. |
	 * | 0x35      | Play  | Client   | Body              | Message                     | Message                     | String              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | Body              | Timestamp                   | Timestamp                   | Long                | Represents the time the message was signed as milliseconds since the epoch, used to check if the message was received within 2 minutes of it being sent.                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | Body              | Salt                        | Salt                        | Long                | Cryptography, used for validating the message signature.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
	 * | 0x35      | Play  | Client   | Previous Messages | Total Previous Messages     | Total Previous Messages     | VarInt              |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | Previous Messages | Array                       | Message ID                  | VarInt              | The message Id, used for validating message signature.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
	 * | 0x35      | Play  | Client   | Previous Messages | Array                       | Signature                   | Optional Byte Array | The previous message's signature. Contains the same type of data as Message Signature bytes above.                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
	 * | 0x35      | Play  | Client   | Other             | Unsigned Content Present    | Unsigned Content Present    | Boolean             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | Other             | Unsigned Content            | Unsigned Content            | Optional Chat       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | Other             | Filter Type                 | Filter Type                 | Enum VarInt         | If the message has been filtered                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
	 * | 0x35      | Play  | Client   | Other             | Filter Type Bits            | Filter Type Bits            | Optional BitSet     |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | Network target    | Chat Type                   | Chat Type                   | VarInt              | The chat type from the Login (play) packet used for this message                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
	 * | 0x35      | Play  | Client   | Network target    | Network name                | Network name                | Chat                |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | Network target    | Network target name present | Network target name present | Boolean             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * | 0x35      | Play  | Client   | Network target    | Network target name         | Network target name         | Optional Chat       |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
	 * 
	 */
}

// ID=0x36
type PlayEndCombatPkt = internal.PlayEndCombat_762_1

// ID=0x37
type PlayEnterCombatPkt = internal.PlayEnterCombat_763_0

// ID=0x38
type PlayCombatDeathPkt = internal.PlayCombatDeath_762_1

// ID=0x39
type PlayPlayerInfoRemovePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type    | Notes                                      |
	 * |-----------|-------|----------|-------------------|-------------------|---------------|--------------------------------------------|
	 * | 0x39      | Play  | Client   | Number of Players | Number of Players | VarInt        | Number of elements in the following array. |
	 * | 0x39      | Play  | Client   | Player            | Player Id         | Array of UUID | UUIDs of players to remove.                |
	 * 
	 */
}

// ID=0x3a
type PlayPlayerInfoUpdatePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name                         | Field Name                | Field Name                | Field Type | Field Type    | Field Type              | Notes                                                                                                          |
	 * |-----------|-------|----------|-------------------|------------------------------------|---------------------------|---------------------------|------------|---------------|-------------------------|----------------------------------------------------------------------------------------------------------------|
	 * | 0x3A      | Play  | Client   |                   |                                    |                           |                           |            |               |                         |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Actions           | Actions                            | Actions                   | Actions                   | Byte       | Byte          | Byte                    | Bit Mask. The actions to process. This must have a bit set for every action below, whether it's true or false. |
	 * | 0x3A      | Play  | Client   | Number Of Actions | Number Of Actions                  | Number Of Actions         | Number Of Actions         | VarInt     | VarInt        | VarInt                  | Number of elements in the following array.                                                                     |
	 * | 0x3A      | Play  | Client   | Action            | UUID                               | UUID                      | UUID                      | Array      | UUID          | UUID                    | The player UUID                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Action                             | Field Name                | Field Name                | Array      |               |                         |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 0: add player          | Name                      | Name                      | Array      | String (16)   | String (16)             |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 0: add player          | Number Of Properties      | Number Of Properties      | Array      | VarInt        | VarInt                  | Number of elements in the following array.                                                                     |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 0: add player          | Property                  | Name                      | Array      | Array         | String (32767)          |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 0: add player          | Property                  | Value                     | Array      | Array         | String (32767)          |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 0: add player          | Property                  | Is Signed                 | Array      | Array         | Boolean                 |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 0: add player          | Property                  | Signature                 | Array      | Array         | Optional String (32767) | Only if Is Signed is true.                                                                                     |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Has Signature Data        | Has Signature Data        | Array      | Boolean       | Boolean                 |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Chat session ID           | Chat session ID           | Array      | UUID          | UUID                    | Only send if Has Signature Data is true.                                                                       |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Public key expiry time    | Public key expiry time    | Array      | Long          | Long                    | Key expiry time, as a UNIX timestamp in milliseconds. Only send if Has Signature Data is true.                 |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Encoded public key size   | Encoded public key size   | Array      | VarInt        | VarInt                  | Size of the following array. Only send if Has Signature Data is true.                                          |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Encoded public key        | Encoded public key        | Array      | Byte Array    | Byte Array              | The player's public key, in bytes. Only send if Has Signature Data is true.                                    |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Public key signature size | Public key signature size | Array      | VarInt        | VarInt                  | Size of the following array. Only send if Has Signature Data is true.                                          |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 1: initialize chat     | Public key signature      | Public key signature      | Array      | Byte Array    | Byte Array              | The public key's digital signature. Only send if Has Signature Data is true.                                   |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 2: update game mode    | Game mode                 | Game mode                 | Array      | VarInt        | VarInt                  |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 3: update listed       | Listed                    | Listed                    | Array      | Boolean       | Boolean                 | Whether the player should be listed on the player list.                                                        |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 4: update latency      | Ping                      | Ping                      | Array      | VarInt        | VarInt                  | Measured in milliseconds.                                                                                      |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 5: update display name | Has Display Name          | Has Display Name          | Array      | Boolean       | Boolean                 |                                                                                                                |
	 * | 0x3A      | Play  | Client   | Action            | Actions bit 5: update display name | Display Name              | Display Name              | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true.                                                                         |
	 * 
	 */
}

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
	DimensionType Identifier // Identifier
	/* Name of the dimension being spawned into. */
	DimensionName Identifier // Identifier
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
	DeathDimensionName Identifier // Optional Identifier
	/* The location that the player died at. */
	DeathLocation Position // Optional Position
}

var _ Packet = (*PlayRespawnPkt)(nil)

func (p PlayRespawnPkt)Encode(b *PacketBuilder){
	b.Encode(p.DimensionType)
	b.Encode(p.DimensionName)
	b.Long(p.HashedSeed)
	b.UByte(p.GameMode)
	b.Byte(p.PreviousGameMode)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
	b.Byte(p.DataKept)
	b.Bool(p.HasDeathLocation)
	b.Encode(p.DeathDimensionName)
	b.Encode(p.DeathLocation)
}

func (p *PlayRespawnPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.DimensionType.DecodeFrom(r); err != nil {
		return
	}
	if err = p.DimensionName.DecodeFrom(r); err != nil {
		return
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
	if err = p.DeathDimensionName.DecodeFrom(r); err != nil {
		return
	}
	if err = p.DeathLocation.DecodeFrom(r); err != nil {
		return
	}
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
type PlaySetEquipmentPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name | Field Type | Field Type | Notes                                                                                                                                                                                                                          |
	 * |-----------|-------|----------|------------|------------|------------|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x55      | Play  | Client   | Entity ID  | Entity ID  | VarInt     | VarInt     | Entity's ID.                                                                                                                                                                                                                   |
	 * | 0x55      | Play  | Client   | Equipment  | Slot       | Array      | Byte Enum  | Equipment slot. 0: main hand, 1: off hand, 2–5: armor slot (2: boots, 3: leggings, 4: chestplate, 5: helmet).  Also has the top bit set if another entry follows, and otherwise unset if this is the last item in the array. |
	 * | 0x55      | Play  | Client   | Equipment  | Item       | Array      | Slot       |                                                                                                                                                                                                                                |
	 * 
	 */
}

// ID=0x56
type PlaySetExperiencePkt = internal.PlaySetExperience_763_0

// ID=0x57
type PlaySetHealthPkt = internal.PlaySetHealth_763_0

// ID=0x58
type PlayUpdateObjectivesPkt = internal.PlayUpdateObjectives_763_0

// ID=0x59
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x5a
type PlayUpdateTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                   | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|------------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x5A      | Play  | Client   | Team Name                    | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x5A      | Play  | Client   | Mode                         | Mode                | Byte                 | Determines the layout of the remaining packet.                                                                         |
	 * | 0x5A      | Play  | Client   | 0: create team               | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x5A      | Play  | Client   | 0: create team               | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team.                                     |
	 * | 0x5A      | Play  | Client   | 0: create team               | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never.                                                                      |
	 * | 0x5A      | Play  | Client   | 0: create team               | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never.                                                                            |
	 * | 0x5A      | Play  | Client   | 0: create team               | Team Color          | VarInt Enum          | Used to color the name of players on the team; see below.                                                              |
	 * | 0x5A      | Play  | Client   | 0: create team               | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team.                                                      |
	 * | 0x5A      | Play  | Client   | 0: create team               | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team.                                                       |
	 * | 0x5A      | Play  | Client   | 0: create team               | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x5A      | Play  | Client   | 0: create team               | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x5A      | Play  | Client   | 1: remove team               | no fields           | no fields            |                                                                                                                        |
	 * | 0x5A      | Play  | Client   | 2: update team info          | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x5A      | Play  | Client   | 2: update team info          | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team.                                    |
	 * | 0x5A      | Play  | Client   | 2: update team info          | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x5A      | Play  | Client   | 2: update team info          | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x5A      | Play  | Client   | 2: update team info          | Team Color          | VarInt Enum          | Used to color the name of players on the team; see below.                                                              |
	 * | 0x5A      | Play  | Client   | 2: update team info          | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team.                                                      |
	 * | 0x5A      | Play  | Client   | 2: update team info          | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team.                                                       |
	 * | 0x5A      | Play  | Client   | 3: add entities to team      | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x5A      | Play  | Client   | 3: add entities to team      | Entities            | Array of String (40) | Identifiers for the added entities.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x5A      | Play  | Client   | 4: remove entities from team | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x5A      | Play  | Client   | 4: remove entities from team | Entities            | Array of String (40) | Identifiers for the removed entities.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

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
type PlayUpdateAdvancementsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name    | Field Type          | Field Type           | Notes                                                       |
	 * |-----------|-------|----------|---------------------|---------------|---------------------|----------------------|-------------------------------------------------------------|
	 * | 0x69      | Play  | Client   | Reset/Clear         | Reset/Clear   | Boolean             | Boolean              | Whether to reset/clear the current advancements.            |
	 * | 0x69      | Play  | Client   | Mapping size        | Mapping size  | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x69      | Play  | Client   | Advancement mapping | Key           | Array               | Identifier           | The identifier of the advancement.                          |
	 * | 0x69      | Play  | Client   | Advancement mapping | Value         | Array               | Advancement          | See below                                                   |
	 * | 0x69      | Play  | Client   | List size           | List size     | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x69      | Play  | Client   | Identifiers         | Identifiers   | Array of Identifier | Array of Identifier  | The identifiers of the advancements that should be removed. |
	 * | 0x69      | Play  | Client   | Progress size       | Progress size | VarInt              | VarInt               | Size of the following array.                                |
	 * | 0x69      | Play  | Client   | Progress mapping    | Key           | Array               | Identifier           | The identifier of the advancement.                          |
	 * | 0x69      | Play  | Client   | Progress mapping    | Value         | Array               | Advancement progress | See below.                                                  |
	 * 
	 */
}

// ID=0x6a
type PlayUpdateAttributesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x6A      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x6A      | Play  | Client   | Number Of Properties | Number Of Properties | VarInt     | VarInt                 | Number of elements in the following array.            |
	 * | 0x6A      | Play  | Client   | Property             | Key                  | Array      | Identifier             | See below.                                            |
	 * | 0x6A      | Play  | Client   | Property             | Value                | Array      | Double                 | See below.                                            |
	 * | 0x6A      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array.            |
	 * | 0x6A      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x6b
type PlayFeatureFlagsPkt = internal.PlayFeatureFlags_763_0

// ID=0x6c
type PlayEntityEffectPkt = internal.PlayEntityEffect_763_0

// ID=0x6d
type PlayUpdateRecipesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name  | Field Type | Field Type | Notes                                      |
	 * |-----------|-------|----------|-------------|-------------|------------|------------|--------------------------------------------|
	 * | 0x6D      | Play  | Client   | Num Recipes | Num Recipes | VarInt     | VarInt     | Number of elements in the following array. |
	 * | 0x6D      | Play  | Client   | Recipe      | Type        | Array      | Identifier | The recipe type, see below.                |
	 * | 0x6D      | Play  | Client   | Recipe      | Recipe ID   | Array      | Identifier |                                            |
	 * | 0x6D      | Play  | Client   | Recipe      | Data        | Array      | Varies     | Additional data for the recipe.            |
	 * 
	 */
}

// ID=0x6e
type PlayUpdateTagsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type | Field Type  | Notes                                                                                                                                        |
	 * |-----------|-------|----------|---------------------|---------------------|------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x6E      | Play  | Client   | Length of the array | Length of the array | VarInt     | VarInt      |                                                                                                                                              |
	 * | 0x6E      | Play  | Client   | Array of tags       | Tag type            | Array      | Identifier  | Tag identifier (Vanilla required tags are minecraft:block, minecraft:item, minecraft:fluid, minecraft:entity_type, and minecraft:game_event) |
	 * | 0x6E      | Play  | Client   | Array of tags       | Array of Tag        | Array      | (See below) |                                                                                                                                              |
	 * 
	 */
}

// ======== END play ========
