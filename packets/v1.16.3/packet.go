
// Generated at 2023-09-01 20:57:33.569 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=16091
// Protocol: 753
// Protocol Name: 1.16.3

package packet_1_16_3

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

// ID=0xd
type PlayQueryEntityNBTPkt = internal.PlayQueryEntityNBT_758_0

// ID=0x2
type PlaySetDifficultyPkt = internal.PlaySetDifficulty_758_0

// ID=0x3
type PlayChatMessageServerPkt = internal.PlayChatMessage_758_3

// ID=0x4
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x5
type PlayClientSettingsPkt = internal.PlayClientSettings_754_2

// ID=0x6
type PlayTabCompleteServerPkt = internal.PlayTabComplete_758_0

// ID=0x7
type PlayWindowConfirmationServerPkt = internal.PlayWindowConfirmation_754_0

// ID=0x8
type PlayClickWindowButtonPkt = internal.PlayClickWindowButton_758_0

// ID=0x9
type PlayClickWindowPkt = internal.PlayClickWindow_754_0

// ID=0xa
type PlayCloseWindowServerPkt = internal.PlayCloseWindow_758_0

// ID=0xb
type PlayPluginMessageServerPkt = internal.PlayPluginMessage_763_0

// ID=0xc
type PlayEditBookPkt = internal.PlayEditBook_755_3

// ID=0xe
type PlayInteractEntityPkt = internal.PlayInteractEntity_758_0

// ID=0xf
type PlayGenerateStructurePkt = internal.PlayGenerateStructure_758_0

// ID=0x10
type PlayKeepAliveServerPkt = internal.PlayKeepAlive_763_0

// ID=0x11
type PlayLockDifficultyPkt = internal.PlayLockDifficulty_763_0

// ID=0x12
type PlayPlayerPositionPkt = internal.PlayPlayerPosition_758_0

// ID=0x13
type PlayPlayerPositionAndRotationPkt = internal.PlayPlayerPositionAndRotation_758_0

// ID=0x14
type PlayPlayerRotationPkt = internal.PlayPlayerRotation_758_0

// ID=0x15
type PlayPlayerMovementPkt = internal.PlayPlayerMovement_758_0

// ID=0x16
type PlayVehicleMoveServerPkt = internal.PlayVehicleMove_758_0

// ID=0x17
type PlaySteerBoatPkt = internal.PlaySteerBoat_758_0

// ID=0x18
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x19
type PlayCraftRecipeRequestPkt = internal.PlayCraftRecipeRequest_758_0

// ID=0x1a
type PlayPlayerAbilitiesServerPkt = internal.PlayPlayerAbilities_763_1

// ID=0x1b
type PlayPlayerDiggingPkt = internal.PlayPlayerDigging_758_0

// ID=0x1c
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x1d
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x1e
type PlaySetDisplayedRecipePkt = internal.PlaySetDisplayedRecipe_758_0

// ID=0x1f
type PlaySetRecipeBookStatePkt = internal.PlaySetRecipeBookState_758_0

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

// ID=0x29
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x28
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
type PlayEntityAnimationPkt = internal.PlayEntityAnimation_763_0

// ID=0x6
type PlayStatisticsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name   | Field Type | Field Type | Notes                                     |
	 * |-----------|-------|----------|------------|--------------|------------|------------|-------------------------------------------|
	 * | 0x06      | Play  | Client   | Count      | Count        | VarInt     | VarInt     | Number of elements in the following array |
	 * | 0x06      | Play  | Client   | Statistic  | Category ID  | Array      | VarInt     | See below                                 |
	 * | 0x06      | Play  | Client   | Statistic  | Statistic ID | Array      | VarInt     | See below                                 |
	 * | 0x06      | Play  | Client   | Statistic  | Value        | Array      | VarInt     | The amount to set it to                   |
	 * 
	 */
}

// ID=0x7
type PlayAcknowledgePlayerDiggingPkt = internal.PlayAcknowledgePlayerDigging_758_0

// ID=0x8
type PlayBlockBreakAnimationPkt = internal.PlayBlockBreakAnimation_758_0

// ID=0x9
type PlayBlockEntityDataPkt = internal.PlayBlockEntityData_756_1

// ID=0xa
type PlayBlockActionPkt = internal.PlayBlockAction_758_1

// ID=0xb
type PlayBlockChangePkt = internal.PlayBlockChange_758_0

// ID=0xc
type PlayBossBarPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name       | Field Name | Field Type    | Notes                                                                                                                                    |
	 * |-----------|-------|----------|------------------|------------|---------------|------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0C      | Play  | Client   | UUID             | UUID       | UUID          | Unique ID for this bar                                                                                                                   |
	 * | 0x0C      | Play  | Client   | Action           | Action     | VarInt Enum   | Determines the layout of the remaining packet                                                                                            |
	 * | 0x0C      | Play  | Client   | Action           | Field Name |               |                                                                                                                                          |
	 * | 0x0C      | Play  | Client   | 0: add           | Title      | Chat          |                                                                                                                                          |
	 * | 0x0C      | Play  | Client   | 0: add           | Health     | Float         | From 0 to 1. Values greater than 1 do not crash a Notchian client, and start rendering part of a second health bar at around 1.5.        |
	 * | 0x0C      | Play  | Client   | 0: add           | Color      | VarInt Enum   | Color ID (see below)                                                                                                                     |
	 * | 0x0C      | Play  | Client   | 0: add           | Division   | VarInt Enum   | Type of division (see below)                                                                                                             |
	 * | 0x0C      | Play  | Client   | 0: add           | Flags      | Unsigned Byte | Bit mask. 0x1: should darken sky, 0x2: is dragon bar (used to play end music), 0x04: create fog (previously was also controlled by 0x02) |
	 * | 0x0C      | Play  | Client   | 1: remove        | no fields  | no fields     | Removes this boss bar                                                                                                                    |
	 * | 0x0C      | Play  | Client   | 2: update health | Health     | Float         | as above                                                                                                                                 |
	 * | 0x0C      | Play  | Client   | 3: update title  | Title      | Chat          |                                                                                                                                          |
	 * | 0x0C      | Play  | Client   | 4: update style  | Color      | VarInt Enum   | Color ID (see below)                                                                                                                     |
	 * | 0x0C      | Play  | Client   | 4: update style  | Dividers   | VarInt Enum   | as above                                                                                                                                 |
	 * | 0x0C      | Play  | Client   | 5: update flags  | Flags      | Unsigned Byte | as above                                                                                                                                 |
	 * 
	 */
}

// ID=0xd
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_758_0

// ID=0xe
type PlayChatMessageClientPkt = internal.PlayChatMessage_758_2

// ID=0xf
type PlayTabCompleteClientPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name  | Field Type | Field Type     | Notes                                                                                                                                                                                                  |
	 * |-----------|-------|----------|------------|-------------|------------|----------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0xF       | Play  | Client   |            |             |            |                |                                                                                                                                                                                                        |
	 * | 0xF       | Play  | Client   | ID         | ID          | VarInt     | VarInt         | Transaction ID                                                                                                                                                                                         |
	 * | 0xF       | Play  | Client   | Start      | Start       | VarInt     | VarInt         | Start of the text to replace                                                                                                                                                                           |
	 * | 0xF       | Play  | Client   | Length     | Length      | VarInt     | VarInt         | Length of the text to replace                                                                                                                                                                          |
	 * | 0xF       | Play  | Client   | Count      | Count       | VarInt     | VarInt         | Number of elements in the following array                                                                                                                                                              |
	 * | 0xF       | Play  | Client   | Matches    | Match       | Array      | String (32767) | One eligible value to insert, note that each command is sent separately instead of in a single string, hence the need for Count.  Note that for instance this doesn't include a leading / on commands. |
	 * | 0xF       | Play  | Client   | Matches    | Has tooltip | Array      | Boolean        | True if the following is present                                                                                                                                                                       |
	 * | 0xF       | Play  | Client   | Matches    | Tooltip     | Array      | Optional Chat  | Tooltip to display; only present if previous boolean is true                                                                                                                                           |
	 * 
	 */
}

// ID=0x10
type PlayDeclareCommandsPkt = internal.PlayDeclareCommands_758_0

// ID=0x11
type PlayWindowConfirmationClientPkt = internal.PlayWindowConfirmation_754_0

// ID=0x12
type PlayCloseWindowClientPkt = internal.PlayCloseWindow_758_0

// ID=0x13
type PlayWindowItemsPkt = internal.PlayWindowItems_755_1

// ID=0x14
type PlayWindowPropertyPkt = internal.PlayWindowProperty_758_0

// ID=0x15
type PlaySetSlotPkt = internal.PlaySetSlot_755_1

// ID=0x16
type PlaySetCooldownPkt = internal.PlaySetCooldown_763_0

// ID=0x17
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_763_0

// ID=0x18
type PlayNamedSoundEffectPkt = internal.PlayNamedSoundEffect_758_0

// ID=0x19
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x1a
type PlayEntityStatusPkt = internal.PlayEntityStatus_758_0

// ID=0x1b
type PlayExplosionPkt = internal.PlayExplosion_754_2

// ID=0x1c
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x1d
type PlayChangeGameStatePkt = internal.PlayChangeGameState_758_0

// ID=0x1e
type PlayOpenHorseWindowPkt = internal.PlayOpenHorseWindow_756_1

// ID=0x1f
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x20
type PlayChunkDataPkt = internal.PlayChunkData_754_1

// ID=0x21
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x22
type PlayParticlePkt = internal.PlayParticle_758_1

// ID=0x23
type PlayUpdateLightPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name             | Field Name             | Field Type | Field Type          | Notes                                                                                                                                                                                                                                                                                                                                                                           |
	 * |-----------|-------|----------|------------------------|------------------------|------------|---------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x23      | Play  | Client   | Chunk X                | Chunk X                | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                 |
	 * | 0x23      | Play  | Client   | Chunk Z                | Chunk Z                | VarInt     | VarInt              | Chunk coordinate (block coordinate divided by 16, rounded down)                                                                                                                                                                                                                                                                                                                 |
	 * | 0x23      | Play  | Client   | Trust Edges            | Trust Edges            | Boolean    | Boolean             | If edges should be trusted for light updates.                                                                                                                                                                                                                                                                                                                                   |
	 * | 0x23      | Play  | Client   | Sky Light Mask         | Sky Light Mask         | VarInt     | VarInt              | Mask containing 18 bits, with the lowest bit corresponding to chunk section -1 (in the void, y=-16 to y=-1) and the highest bit for chunk section 16 (above the world, y=256 to y=271)                                                                                                                                                                                          |
	 * | 0x23      | Play  | Client   | Block Light Mask       | Block Light Mask       | VarInt     | VarInt              | Mask containing 18 bits, with the same order as sky light                                                                                                                                                                                                                                                                                                                       |
	 * | 0x23      | Play  | Client   | Empty Sky Light Mask   | Empty Sky Light Mask   | VarInt     | VarInt              | Mask containing 18 bits, which indicates sections that have 0 for all their sky light values.  If a section is set in both this mask and the main sky light mask, it is ignored for this mask and it is included in the sky light arrays (the notchian server does not create such masks).  If it is only set in this mask, it is not included in the sky light arrays.         |
	 * | 0x23      | Play  | Client   | Empty Block Light Mask | Empty Block Light Mask | VarInt     | VarInt              | Mask containing 18 bits, which indicates sections that have 0 for all their block light values.  If a section is set in both this mask and the main block light mask, it is ignored for this mask and it is included in the block light arrays (the notchian server does not create such masks).  If it is only set in this mask, it is not included in the block light arrays. |
	 * | 0x23      | Play  | Client   | Sky Light arrays       | Length                 | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                            |
	 * | 0x23      | Play  | Client   | Sky Light arrays       | Sky Light array        | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the sky light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                  |
	 * | 0x23      | Play  | Client   | Block Light arrays     | Length                 | Array      | VarInt              | Length of the following array in bytes (always 2048)                                                                                                                                                                                                                                                                                                                            |
	 * | 0x23      | Play  | Client   | Block Light arrays     | Block Light array      | Array      | Array of 2048 bytes | There is 1 array for each bit set to true in the block light mask, starting with the lowest value.  Half a byte per light value.                                                                                                                                                                                                                                                |
	 * 
	 */
}

// ID=0x24
type PlayJoinGamePkt = internal.PlayJoinGame_756_1

// ID=0x25
type PlayMapDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type                      | Field Type                      | Notes                                                                                                      |
	 * |-----------|-------|----------|-------------------|-------------------|---------------------------------|---------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x25      | Play  | Client   | Map ID            | Map ID            | VarInt                          | VarInt                          | Map ID of the map being modified                                                                           |
	 * | 0x25      | Play  | Client   | Scale             | Scale             | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x25      | Play  | Client   | Tracking Position | Tracking Position | Boolean                         | Boolean                         | Specifies whether player and item frame icons are shown                                                    |
	 * | 0x25      | Play  | Client   | Locked            | Locked            | Boolean                         | Boolean                         | True if the map has been locked in a cartography table                                                     |
	 * | 0x25      | Play  | Client   | Icon Count        | Icon Count        | VarInt                          | VarInt                          | Number of elements in the following array                                                                  |
	 * | 0x25      | Play  | Client   | Icon              | Type              | Array                           | VarInt enum                     | See below                                                                                                  |
	 * | 0x25      | Play  | Client   | Icon              | X                 | Array                           | Byte                            | Map coordinates: -128 for furthest left, +127 for furthest right                                           |
	 * | 0x25      | Play  | Client   | Icon              | Z                 | Array                           | Byte                            | Map coordinates: -128 for highest, +127 for lowest                                                         |
	 * | 0x25      | Play  | Client   | Icon              | Direction         | Array                           | Byte                            | 0-15                                                                                                       |
	 * | 0x25      | Play  | Client   | Icon              | Has Display Name  | Array                           | Boolean                         |                                                                                                            |
	 * | 0x25      | Play  | Client   | Icon              | Display Name      | Array                           | Optional Chat                   | Only present if previous Boolean is true                                                                   |
	 * | 0x25      | Play  | Client   | Columns           | Columns           | Unsigned Byte                   | Unsigned Byte                   | Number of columns updated                                                                                  |
	 * | 0x25      | Play  | Client   | Rows              | Rows              | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; number of rows updated                                                     |
	 * | 0x25      | Play  | Client   | X                 | X                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column                                         |
	 * | 0x25      | Play  | Client   | Z                 | Z                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row                                           |
	 * | 0x25      | Play  | Client   | Length            | Length            | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array                                              |
	 * | 0x25      | Play  | Client   | Data              | Data              | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format                                                        |
	 * 
	 */
}

// ID=0x26
type PlayTradeListPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name                   | Field Type | Field Type    | Notes                                                                                                                                                                              |
	 * |-----------|-------|----------|---------------------|------------------------------|------------|---------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x26      | Play  | Client   |                     |                              |            |               |                                                                                                                                                                                    |
	 * | 0x26      | Play  | Client   | Window ID           | Window ID                    | VarInt     | VarInt        | The ID of the window that is open; this is an int rather than a byte.                                                                                                              |
	 * | 0x26      | Play  | Client   | Size                | Size                         | Byte       | Byte          | The number of trades in the following array                                                                                                                                        |
	 * | 0x26      | Play  | Client   | Trades              | Input item 1                 | Array      | Slot          | The first item the villager is buying                                                                                                                                              |
	 * | 0x26      | Play  | Client   | Trades              | Output item                  | Array      | Slot          | The item the villager is selling                                                                                                                                                   |
	 * | 0x26      | Play  | Client   | Trades              | Has second item              | Array      | Boolean       | Whether there is a second item                                                                                                                                                     |
	 * | 0x26      | Play  | Client   | Trades              | Input item 2                 | Array      | Optional Slot | The second item the villager is buying; only present if they have a second item.                                                                                                   |
	 * | 0x26      | Play  | Client   | Trades              | Trade disabled               | Array      | Boolean       | True if the trade is disabled; false if the trade is enabled.                                                                                                                      |
	 * | 0x26      | Play  | Client   | Trades              | Number of trade uses         | Array      | Integer       | Number of times the trade has been used so far                                                                                                                                     |
	 * | 0x26      | Play  | Client   | Trades              | Maximum number of trade uses | Array      | Integer       | Number of times this trade can be used                                                                                                                                             |
	 * | 0x26      | Play  | Client   | Trades              | XP                           | Array      | Integer       |                                                                                                                                                                                    |
	 * | 0x26      | Play  | Client   | Trades              | Special Price                | Array      | Integer       |                                                                                                                                                                                    |
	 * | 0x26      | Play  | Client   | Trades              | Price Multiplier             | Array      | Float         |                                                                                                                                                                                    |
	 * | 0x26      | Play  | Client   | Trades              | Demand                       | Array      | Integer       |                                                                                                                                                                                    |
	 * | 0x26      | Play  | Client   | Villager level      | Villager level               | VarInt     | VarInt        | Appears on the trade GUI; meaning comes from the translation key merchant.level. + level.
	 * 1: Novice, 2: Apprentice, 3: Journeyman, 4: Expert, 5: Master                            |
	 * | 0x26      | Play  | Client   | Experience          | Experience                   | VarInt     | VarInt        | Total experience for this villager (always 0 for the wandering trader)                                                                                                             |
	 * | 0x26      | Play  | Client   | Is regular villager | Is regular villager          | Boolean    | Boolean       | True if this is a regular villager; false for the wandering trader.  When false, hides the villager level and some other GUI elements.                                             |
	 * | 0x26      | Play  | Client   | Can restock         | Can restock                  | Boolean    | Boolean       | True for regular villagers and false for the wandering trader.  If true, the "Villagers restock up to two times per day." message is displayed when hovering over disabled trades. |
	 * 
	 */
}

// ID=0x27
type PlayEntityPositionPkt = internal.PlayEntityPosition_758_0

// ID=0x28
type PlayEntityPositionAndRotationPkt = internal.PlayEntityPositionAndRotation_758_0

// ID=0x29
type PlayEntityRotationPkt = internal.PlayEntityRotation_758_0

// ID=0x2a
type PlayEntityMovementPkt = internal.PlayEntityMovement_754_0

// ID=0x2b
type PlayVehicleMoveClientPkt = internal.PlayVehicleMove_758_0

// ID=0x2c
type PlayOpenBookPkt = internal.PlayOpenBook_763_0

// ID=0x2d
type PlayOpenWindowPkt = internal.PlayOpenWindow_758_0

// ID=0x2e
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x2f
type PlayCraftRecipeResponsePkt = internal.PlayCraftRecipeResponse_758_0

// ID=0x30
type PlayPlayerAbilitiesClientPkt = internal.PlayPlayerAbilities_763_0

// ID=0x31
type PlayCombatEventPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name      | Field Name | Field Type  | Notes                                                                                          |
	 * |-----------|-------|----------|-----------------|------------|-------------|------------------------------------------------------------------------------------------------|
	 * | 0x31      | Play  | Client   | Event           | Event      | VarInt Enum | Determines the layout of the remaining packet                                                  |
	 * | 0x31      | Play  | Client   | Event           | Field Name |             |                                                                                                |
	 * | 0x31      | Play  | Client   | 0: enter combat | no fields  | no fields   |                                                                                                |
	 * | 0x31      | Play  | Client   | 1: end combat   | Duration   | VarInt      | Length of the combat in ticks.                                                                 |
	 * | 0x31      | Play  | Client   | 1: end combat   | Entity ID  | Int         | ID of the primary opponent of the ended combat, or -1 if there is no obvious primary opponent. |
	 * | 0x31      | Play  | Client   | 2: entity dead  | Player ID  | VarInt      | Entity ID of the player that died (should match the client's entity ID).                       |
	 * | 0x31      | Play  | Client   | 2: entity dead  | Entity ID  | Int         | The killing entity's ID, or -1 if there is no obvious killer.                                  |
	 * | 0x31      | Play  | Client   | 2: entity dead  | Message    | Chat        | The death message                                                                              |
	 * 
	 */
}

// ID=0x32
type PlayPlayerInfoPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type              | Notes                                                   |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-------------------------|---------------------------------------------------------|
	 * | 0x32      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt                  | Determines the rest of the Player format after the UUID |
	 * | 0x32      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x32      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID                    |                                                         |
	 * | 0x32      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                         |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String (16)   | String (16)             |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String (32767)          |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String (32767)          |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean                 |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String (32767) | Only if Is Signed is true                               |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds                                |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only if Has Display Name is true                        |
	 * | 0x32      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds                                |
	 * | 0x32      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x32      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true                   |
	 * | 0x32      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields               |                                                         |
	 * 
	 */
}

// ID=0x33
type PlayFacePlayerPkt = internal.PlayFacePlayer_757_0

// ID=0x34
type PlayPlayerPositionAndLookPkt = internal.PlayPlayerPositionAndLook_754_1

// ID=0x35
type PlayUnlockRecipesPkt struct {
	/* 0: init, 1: add, 2: remove */
	Action VarInt // VarInt
	/* If true, then the crafting recipe book will be open when the player opens its inventory. */
	CraftingRecipeBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	CraftingRecipeBookFilterActive Bool // Boolean
	/* If true, then the smelting recipe book will be open when the player opens its inventory. */
	SmeltingRecipeBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	SmeltingRecipeBookFilterActive Bool // Boolean
	/* If true, then the blast furnace recipe book will be open when the player opens its inventory. */
	BlastFuranceRecipeBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	BlastFuranceRecipeBookFilterActive Bool // Boolean
	/* If true, then the smoker recipe book will be open when the player opens its inventory. */
	SmokerRecipeBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	SmokerRecipeBookFilterActive Bool // Boolean
	/* Number of elements in the following array */
	ArraySize1 VarInt // VarInt
	RecipeIDs []Identifier // Array of Identifier
	/* Number of elements in the following array, only present if mode is 0 (init) */
	ArraySize2 VarInt // Optional VarInt
	RecipeIDs []Identifier // Optional Array of Identifier, only present if mode is 0 (init)
}

var _ Packet = (*PlayUnlockRecipesPkt)(nil)

func (p PlayUnlockRecipesPkt)Encode(b *PacketBuilder){
	b.VarInt(p.Action)
	b.Bool(p.CraftingRecipeBookOpen)
	b.Bool(p.CraftingRecipeBookFilterActive)
	b.Bool(p.SmeltingRecipeBookOpen)
	b.Bool(p.SmeltingRecipeBookFilterActive)
	b.Bool(p.BlastFuranceRecipeBookOpen)
	b.Bool(p.BlastFuranceRecipeBookFilterActive)
	b.Bool(p.SmokerRecipeBookOpen)
	b.Bool(p.SmokerRecipeBookFilterActive)
	b.VarInt(p.ArraySize1)
	TODO_Encode_Array(p.RecipeIDs)
	b.VarInt(p.ArraySize2)
	TODO_Encode_Array(p.RecipeIDs)
}

func (p *PlayUnlockRecipesPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Action, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CraftingRecipeBookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.CraftingRecipeBookFilterActive, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.SmeltingRecipeBookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.SmeltingRecipeBookFilterActive, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.BlastFuranceRecipeBookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.BlastFuranceRecipeBookFilterActive, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.SmokerRecipeBookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.SmokerRecipeBookFilterActive, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.ArraySize1, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.RecipeIDs)
	if p.ArraySize2, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.RecipeIDs)
}

// ID=0x36
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x37
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x38
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x39
type PlayRespawnPkt = internal.PlayRespawn_756_4

// ID=0x3a
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x3b
type PlayMultiBlockChangePkt = internal.PlayMultiBlockChange_758_0

// ID=0x3c
type PlaySelectAdvancementTabPkt = internal.PlaySelectAdvancementTab_758_0

// ID=0x3d
type PlayWorldBorderPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name            | Field Name               | Field Type  | Notes                                                                                                                                                                                                                                        |
	 * |-----------|-------|----------|-----------------------|--------------------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x3D      | Play  | Client   | Action                | Action                   | VarInt Enum | Determines the format of the rest of the packet                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | Action                | Field Name               |             |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 0: set size           | Diameter                 | Double      | Length of a single side of the world border, in meters                                                                                                                                                                                       |
	 * | 0x3D      | Play  | Client   | 1: lerp size          | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x3D      | Play  | Client   | 1: lerp size          | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x3D      | Play  | Client   | 1: lerp size          | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x3D      | Play  | Client   | 2: set center         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 2: set center         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 3: initialize         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x3D      | Play  | Client   | 3: initialize         | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Portal Teleport Boundary | VarInt      | Resulting coordinates from a portal teleport are limited to ±value. Usually 29999984.                                                                                                                                                       |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * | 0x3D      | Play  | Client   | 4: set warning time   | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x3D      | Play  | Client   | 5: set warning blocks | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * 
	 */
}

// ID=0x3e
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x3f
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_758_0

// ID=0x40
type PlayUpdateViewPositionPkt = internal.PlayUpdateViewPosition_758_0

// ID=0x41
type PlayUpdateViewDistancePkt = internal.PlayUpdateViewDistance_758_0

// ID=0x42
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_754_1

// ID=0x43
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x44
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x45
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x46
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x47
type PlayEntityEquipmentPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name | Field Type | Field Type | Notes                                                                                                                                                                                                                          |
	 * |-----------|-------|----------|------------|------------|------------|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x47      | Play  | Client   | Entity ID  | Entity ID  | VarInt     | VarInt     | Entity's EID                                                                                                                                                                                                                   |
	 * | 0x47      | Play  | Client   | Equipment  | Slot       | Array      | Byte Enum  | Equipment slot. 0: main hand, 1: off hand, 2–5: armor slot (2: boots, 3: leggings, 4: chestplate, 5: helmet).  Also has the top bit set if another entry follows, and otherwise unset if this is the last item in the array. |
	 * | 0x47      | Play  | Client   | Equipment  | Item       | Array      | Slot       |                                                                                                                                                                                                                                |
	 * 
	 */
}

// ID=0x48
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x49
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x4a
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_758_0

// ID=0x4b
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x4c
type PlayTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                  | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|-----------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x4C      | Play  | Client   | Team Name                   | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x4C      | Play  | Client   | Mode                        | Mode                | Byte                 | Determines the layout of the remaining packet                                                                          |
	 * | 0x4C      | Play  | Client   | 0: create team              | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x4C      | Play  | Client   | 0: create team              | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team                                      |
	 * | 0x4C      | Play  | Client   | 0: create team              | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x4C      | Play  | Client   | 0: create team              | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x4C      | Play  | Client   | 0: create team              | Team Color          | VarInt enum          | Used to color the name of players on the team; see below                                                               |
	 * | 0x4C      | Play  | Client   | 0: create team              | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team                                                       |
	 * | 0x4C      | Play  | Client   | 0: create team              | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team                                                        |
	 * | 0x4C      | Play  | Client   | 0: create team              | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x4C      | Play  | Client   | 0: create team              | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x4C      | Play  | Client   | 1: remove team              | no fields           | no fields            |                                                                                                                        |
	 * | 0x4C      | Play  | Client   | 2: update team info         | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x4C      | Play  | Client   | 2: update team info         | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team                                     |
	 * | 0x4C      | Play  | Client   | 2: update team info         | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x4C      | Play  | Client   | 2: update team info         | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x4C      | Play  | Client   | 2: update team info         | Team Color          | VarInt enum          | Used to color the name of players on the team; see below                                                               |
	 * | 0x4C      | Play  | Client   | 2: update team info         | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team                                                       |
	 * | 0x4C      | Play  | Client   | 2: update team info         | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team                                                        |
	 * | 0x4C      | Play  | Client   | 3: add players to team      | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x4C      | Play  | Client   | 3: add players to team      | Entities            | Array of String (40) | Identifiers for the entities added.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x4C      | Play  | Client   | 4: remove players from team | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x4C      | Play  | Client   | 4: remove players from team | Entities            | Array of String (40) | Identifiers for the entities removed.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

// ID=0x4d
type PlayUpdateScorePkt = internal.PlayUpdateScore_757_1

// ID=0x4e
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x4f
type PlayTitlePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name      | Field Type  | Notes                                                                                                                                                                            |
	 * |-----------|-------|----------|--------------------------|-----------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x4F      | Play  | Client   | Action                   | Action          | VarInt Enum |                                                                                                                                                                                  |
	 * | 0x4F      | Play  | Client   | Action                   | Field Name      |             |                                                                                                                                                                                  |
	 * | 0x4F      | Play  | Client   | 0: set title             | Title Text      | Chat        |                                                                                                                                                                                  |
	 * | 0x4F      | Play  | Client   | 1: set subtitle          | Subtitle Text   | Chat        |                                                                                                                                                                                  |
	 * | 0x4F      | Play  | Client   | 2: set action bar        | Action bar text | Chat        | Displays a message above the hotbar (the same as position 2 in Chat Message (clientbound), except that it correctly renders formatted chat. See MC-119145 for more information.) |
	 * | 0x4F      | Play  | Client   | 3: set times and display | Fade In         | Int         | Ticks to spend fading in                                                                                                                                                         |
	 * | 0x4F      | Play  | Client   | 3: set times and display | Stay            | Int         | Ticks to keep the title displayed                                                                                                                                                |
	 * | 0x4F      | Play  | Client   | 3: set times and display | Fade Out        | Int         | Ticks to spend out, not when to start fading out                                                                                                                                 |
	 * | 0x4F      | Play  | Client   | 4: hide                  | no fields       | no fields   |                                                                                                                                                                                  |
	 * | 0x4F      | Play  | Client   | 5: reset                 | no fields       | no fields   |                                                                                                                                                                                  |
	 * 
	 */
}

// ID=0x50
type PlayEntitySoundEffectPkt = internal.PlayEntitySoundEffect_760_1

// ID=0x51
type PlaySoundEffectPkt = internal.PlaySoundEffect_758_2

// ID=0x52
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x53
type PlayPlayerListHeaderAndFooterPkt = internal.PlayPlayerListHeaderAndFooter_758_0

// ID=0x54
type PlayNBTQueryResponsePkt = internal.PlayNBTQueryResponse_758_0

// ID=0x55
type PlayCollectItemPkt = internal.PlayCollectItem_758_0

// ID=0x56
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x57
type PlayAdvancementsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name    | Field Type          | Field Type           | Notes                                                      |
	 * |-----------|-------|----------|---------------------|---------------|---------------------|----------------------|------------------------------------------------------------|
	 * | 0x57      | Play  | Client   | Reset/Clear         | Reset/Clear   | Boolean             | Boolean              | Whether to reset/clear the current advancements            |
	 * | 0x57      | Play  | Client   | Mapping size        | Mapping size  | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x57      | Play  | Client   | Advancement mapping | Key           | Array               | Identifier           | The identifier of the advancement                          |
	 * | 0x57      | Play  | Client   | Advancement mapping | Value         | Array               | Advancement          | See below                                                  |
	 * | 0x57      | Play  | Client   | List size           | List size     | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x57      | Play  | Client   | Identifiers         | Identifiers   | Array of Identifier | Array of Identifier  | The identifiers of the advancements that should be removed |
	 * | 0x57      | Play  | Client   | Progress size       | Progress size | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x57      | Play  | Client   | Progress mapping    | Key           | Array               | Identifier           | The identifier of the advancement                          |
	 * | 0x57      | Play  | Client   | Progress mapping    | Value         | Array               | Advancement progress | See below                                                  |
	 * 
	 */
}

// ID=0x58
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x58      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x58      | Play  | Client   | Number Of Properties | Number Of Properties | Int        | Int                    | Number of elements in the following array             |
	 * | 0x58      | Play  | Client   | Property             | Key                  | Array      | Identifier             | See below                                             |
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
	 * | 0x5A      | Play  | Client   | Recipe      | Type        | Array      | Identifier       | The recipe type, see below                                              |
	 * | 0x5A      | Play  | Client   | Recipe      | Recipe ID   | Array      | String           |                                                                         |
	 * | 0x5A      | Play  | Client   | Recipe      | Data        | Array      | Optional, varies | Additional data for the recipe.  For some types, there will be no data. |
	 * 
	 */
}

// ID=0x5b
type PlayTagsPkt = internal.PlayTags_754_0

// ======== END play ========
