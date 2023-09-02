
// Generated at 2023-09-01 20:57:33.569 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=13223
// Protocol: 335
// Protocol Name: 1.12

package packet_1_12

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
type PlayPrepareCraftingGridPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name    | Field Name    | Field Type | Field Type | Notes                                                                               |
	 * |-----------|-------|----------|---------------|---------------|------------|------------|-------------------------------------------------------------------------------------|
	 * | 0x01      | Play  | Server   | Window ID     | Window ID     | Byte       | Byte       | The window id.                                                                      |
	 * | 0x01      | Play  | Server   | Action number | Action number | Short      | Short      | The transaction number. Will be send to the client in a Confirm Transaction packet. |
	 * | 0x01      | Play  | Server   | Array size    | Array size    | Short      | Short      | Number of elements in the following array                                           |
	 * | 0x01      | Play  | Server   | Return Entry  | Item          | Array      | Slot       | The item stack that will be put in the inventory slot                               |
	 * | 0x01      | Play  | Server   | Return Entry  | Crafting Slot | Array      | Byte       | The crafting slot index in the active container                                     |
	 * | 0x01      | Play  | Server   | Return Entry  | Player Slot   | Array      | Byte       | The player slot index in the player inventory                                       |
	 * | 0x01      | Play  | Server   | Array Size    | Array Size    | Short      | Short      | Number of elements in the following array                                           |
	 * | 0x01      | Play  | Server   | Prepare Entry | Item          | Array      | Slot       | The item stack that will be put in the crafting slot                                |
	 * | 0x01      | Play  | Server   | Prepare Entry | Crafting Slot | Array      | Byte       | The crafting slot index in the active container                                     |
	 * | 0x01      | Play  | Server   | Prepare Entry | Player Slot   | Array      | Byte       | The player slot index in the player inventory                                       |
	 * 
	 */
}

// ID=0x2
type PlayTabCompleteServerPkt = internal.PlayTabComplete_335_6

// ID=0x3
type PlayChatMessageServerPkt = internal.PlayChatMessage_404_11

// ID=0x4
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x5
type PlayClientSettingsPkt = internal.PlayClientSettings_754_2

// ID=0x6
type PlayConfirmTransactionServerPkt = internal.PlayConfirmTransaction_404_0

// ID=0x7
type PlayEnchantItemPkt = internal.PlayEnchantItem_404_0

// ID=0x8
type PlayClickWindowPkt = internal.PlayClickWindow_754_0

// ID=0x9
type PlayCloseWindowServerPkt = internal.PlayCloseWindow_758_0

// ID=0xa
type PlayPluginMessageServerPkt = internal.PlayPluginMessage_340_1

// ID=0xb
type PlayUseEntityPkt = internal.PlayUseEntity_404_0

// ID=0xc
type PlayKeepAliveServerPkt = internal.PlayKeepAlive_338_1

// ID=0xd
type PlayPlayerPkt = internal.PlayPlayer_404_0

// ID=0xe
type PlayPlayerPositionPkt = internal.PlayPlayerPosition_758_0

// ID=0xf
type PlayPlayerPositionAndLookServerPkt = internal.PlayPlayerPositionAndLook_335_4

// ID=0x10
type PlayPlayerLookPkt = internal.PlayPlayerLook_404_0

// ID=0x11
type PlayVehicleMoveServerPkt = internal.PlayVehicleMove_758_0

// ID=0x12
type PlaySteerBoatPkt = internal.PlaySteerBoat_401_1

// ID=0x13
type PlayPlayerAbilitiesServerPkt = internal.PlayPlayerAbilities_404_12

// ID=0x14
type PlayPlayerDiggingPkt = internal.PlayPlayerDigging_758_0

// ID=0x15
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x16
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x17
type PlayCraftingBookDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name              | Field Name         | Field Type | Notes                                                               |
	 * |-----------|-------|----------|-------------------------|--------------------|------------|---------------------------------------------------------------------|
	 * | 0x17      | Play  | Server   | Type                    | Type               | VarInt     | Determines the format of the rest of the packet                     |
	 * | 0x17      | Play  | Server   | Type                    | Field Name         |            |                                                                     |
	 * | 0x17      | Play  | Server   | 0: Displayed Recipe     | Recipe ID          | Int        | The internal id of the displayed recipe.                            |
	 * | 0x17      | Play  | Server   | 1: Crafting Book Status | Crafting Book Open | Boolean    | Whether the player has the crafting book currently openened/active. |
	 * | 0x17      | Play  | Server   | 1: Crafting Book Status | Crafting Filter    | Boolean    | Whether the player has the crafting filter option currently active. |
	 * 
	 */
}

// ID=0x18
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_758_0

// ID=0x19
type PlayAdvancementTabPkt = internal.PlayAdvancementTab_758_0

// ID=0x1a
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChange_404_5

// ID=0x1b
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x1c
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x1d
type PlayAnimationServerPkt = internal.PlayAnimation_335_4

// ID=0x1e
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x1f
type PlayPlayerBlockPlacementPkt = internal.PlayPlayerBlockPlacement_404_1

// ID=0x20
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
type PlaySpawnPaintingPkt = internal.PlaySpawnPainting_340_1

// ID=0x5
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_498_1

// ID=0x6
type PlayAnimationClientPkt = internal.PlayAnimation_335_3

// ID=0x7
type PlayStatisticsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name | Field Type | Field Type     | Notes                                                             |
	 * |-----------|-------|----------|------------|------------|------------|----------------|-------------------------------------------------------------------|
	 * | 0x07      | Play  | Client   | Count      | Count      | VarInt     | VarInt         | Number of elements in the following array                         |
	 * | 0x07      | Play  | Client   | Statistic  | Name       | Array      | String (32767) | https://gist.github.com/Alvin-LB/8d0d13db00b3c00fd0e822a562025eff |
	 * | 0x07      | Play  | Client   | Statistic  | Value      | Array      | VarInt         | The amount to set it to                                           |
	 * 
	 */
}

// ID=0x8
type PlayBlockBreakAnimationPkt = internal.PlayBlockBreakAnimation_758_0

// ID=0x9
type PlayUpdateBlockEntityPkt = internal.PlayUpdateBlockEntity_498_0

// ID=0xa
type PlayBlockActionPkt = internal.PlayBlockAction_758_1

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
type PlayTabCompleteClientPkt = internal.PlayTabComplete_335_5

// ID=0xf
type PlayChatMessageClientPkt = internal.PlayChatMessage_404_10

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
type PlayNamedSoundEffectPkt = internal.PlayNamedSoundEffect_340_1

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
type PlayChunkDataPkt = internal.PlayChunkData_401_5

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
type PlayEntityPkt = internal.PlayEntity_404_0

// ID=0x26
type PlayEntityRelativeMovePkt = internal.PlayEntityRelativeMove_404_0

// ID=0x27
type PlayEntityLookAndRelativeMovePkt = internal.PlayEntityLookAndRelativeMove_404_0

// ID=0x28
type PlayEntityLookPkt = internal.PlayEntityLook_404_0

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
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type              | Notes                                                   |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-------------------------|---------------------------------------------------------|
	 * | 0x2D      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt                  | Determines the rest of the Player format after the UUID |
	 * | 0x2D      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x2D      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID                    |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                         |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String (16)   | String (16)             |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String (32767)          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String (32767)          |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean                 |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String (32767) | Only if Is Signed is true                               |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only if Has Display Name is true                        |
	 * | 0x2D      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x2D      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true                   |
	 * | 0x2D      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields               |                                                         |
	 * 
	 */
}

// ID=0x2e
type PlayPlayerPositionAndLookClientPkt = internal.PlayPlayerPositionAndLook_335_3

// ID=0x2f
type PlayUseBedPkt = internal.PlayUseBed_404_0

// ID=0x30
type PlayUnlockRecipesPkt = internal.PlayUnlockRecipes_340_3

// ID=0x31
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x32
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x33
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x34
type PlayRespawnPkt = internal.PlayRespawn_404_7

// ID=0x35
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x36
type PlaySelectAdvancementTabPkt = internal.PlaySelectAdvancementTab_758_0

// ID=0x37
type PlayWorldBorderPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name            | Field Name               | Field Type  | Notes                                                                                                                                                                                                                                        |
	 * |-----------|-------|----------|-----------------------|--------------------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x37      | Play  | Client   | Action                | Action                   | VarInt Enum | Determines the format of the rest of the packet                                                                                                                                                                                              |
	 * | 0x37      | Play  | Client   | Action                | Field Name               |             |                                                                                                                                                                                                                                              |
	 * | 0x37      | Play  | Client   | 0: set size           | Diameter                 | Double      | Length of a single side of the world border, in meters                                                                                                                                                                                       |
	 * | 0x37      | Play  | Client   | 1: lerp size          | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x37      | Play  | Client   | 1: lerp size          | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x37      | Play  | Client   | 1: lerp size          | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x37      | Play  | Client   | 2: set center         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x37      | Play  | Client   | 2: set center         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x37      | Play  | Client   | 3: initialize         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x37      | Play  | Client   | 3: initialize         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x37      | Play  | Client   | 3: initialize         | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x37      | Play  | Client   | 3: initialize         | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x37      | Play  | Client   | 3: initialize         | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x37      | Play  | Client   | 3: initialize         | Portal Teleport Boundary | VarInt      | Resulting coordinates from a portal teleport are limited to ±value. Usually 29999984.                                                                                                                                                       |
	 * | 0x37      | Play  | Client   | 3: initialize         | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x37      | Play  | Client   | 3: initialize         | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * | 0x37      | Play  | Client   | 4: set warning time   | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x37      | Play  | Client   | 5: set warning blocks | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * 
	 */
}

// ID=0x38
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x39
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_404_4

// ID=0x3a
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x3b
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x3c
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x3d
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x3e
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_578_0

// ID=0x3f
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x40
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x41
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_340_1

// ID=0x42
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x43
type PlayTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                  | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|-----------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x43      | Play  | Client   | Team Name                   | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x43      | Play  | Client   | Mode                        | Mode                | Byte                 | Determines the layout of the remaining packet                                                                          |
	 * | 0x43      | Play  | Client   | 0: create team              | Team Display Name   | String (32)          |                                                                                                                        |
	 * | 0x43      | Play  | Client   | 0: create team              | Team Prefix         | String (16)          | Displayed before the names of players that are part of this team                                                       |
	 * | 0x43      | Play  | Client   | 0: create team              | Team Suffix         | String (16)          | Displayed after the names of players that are part of this team                                                        |
	 * | 0x43      | Play  | Client   | 0: create team              | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team                                      |
	 * | 0x43      | Play  | Client   | 0: create team              | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x43      | Play  | Client   | 0: create team              | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x43      | Play  | Client   | 0: create team              | Color               | Byte                 | For colors, the same Chat colors (0-15).  -1 indicates RESET/no color.                                                 |
	 * | 0x43      | Play  | Client   | 0: create team              | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x43      | Play  | Client   | 0: create team              | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x43      | Play  | Client   | 1: remove team              | no fields           | no fields            |                                                                                                                        |
	 * | 0x43      | Play  | Client   | 2: update team info         | Team Display Name   | String (32)          |                                                                                                                        |
	 * | 0x43      | Play  | Client   | 2: update team info         | Team Prefix         | String (16)          | Displayed before the names of entities that are part of this team                                                      |
	 * | 0x43      | Play  | Client   | 2: update team info         | Team Suffix         | String (16)          | Displayed after the names of entities that are part of this team                                                       |
	 * | 0x43      | Play  | Client   | 2: update team info         | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team                                     |
	 * | 0x43      | Play  | Client   | 2: update team info         | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x43      | Play  | Client   | 2: update team info         | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x43      | Play  | Client   | 2: update team info         | Color               | Byte                 | For colors, the same Chat colors (0-15).  -1 indicates RESET/no color.                                                 |
	 * | 0x43      | Play  | Client   | 3: add players to team      | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x43      | Play  | Client   | 3: add players to team      | Entities            | Array of String (40) | Identifiers for the entities added.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x43      | Play  | Client   | 4: remove players from team | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x43      | Play  | Client   | 4: remove players from team | Entities            | Array of String (40) | Identifiers for the entities removed.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

// ID=0x44
type PlayUpdateScorePkt = internal.PlayUpdateScore_757_1

// ID=0x45
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_754_1

// ID=0x46
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x47
type PlayTitlePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name      | Field Type  | Notes                                                                                                                                                                            |
	 * |-----------|-------|----------|--------------------------|-----------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x47      | Play  | Client   | Action                   | Action          | VarInt Enum |                                                                                                                                                                                  |
	 * | 0x47      | Play  | Client   | Action                   | Field Name      |             |                                                                                                                                                                                  |
	 * | 0x47      | Play  | Client   | 0: set title             | Title Text      | Chat        |                                                                                                                                                                                  |
	 * | 0x47      | Play  | Client   | 1: set subtitle          | Subtitle Text   | Chat        |                                                                                                                                                                                  |
	 * | 0x47      | Play  | Client   | 2: set action bar        | Action bar text | Chat        | Displays a message above the hotbar (the same as position 2 in Chat Message (clientbound), except that it correctly renders formatted chat. See MC-119145 for more information.) |
	 * | 0x47      | Play  | Client   | 3: set times and display | Fade In         | Int         | Ticks to spend fading in                                                                                                                                                         |
	 * | 0x47      | Play  | Client   | 3: set times and display | Stay            | Int         | Ticks to keep the title displayed                                                                                                                                                |
	 * | 0x47      | Play  | Client   | 3: set times and display | Fade Out        | Int         | Ticks to spend out, not when to start fading out                                                                                                                                 |
	 * | 0x47      | Play  | Client   | 4: hide                  | no fields       | no fields   |                                                                                                                                                                                  |
	 * | 0x47      | Play  | Client   | 5: reset                 | no fields       | no fields   |                                                                                                                                                                                  |
	 * 
	 */
}

// ID=0x48
type PlaySoundEffectPkt = internal.PlaySoundEffect_758_2

// ID=0x49
type PlayPlayerListHeaderAndFooterPkt = internal.PlayPlayerListHeaderAndFooter_758_0

// ID=0x4a
type PlayCollectItemPkt = internal.PlayCollectItem_758_0

// ID=0x4b
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x4c
type PlayAdvancementsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name    | Field Type          | Field Type           | Notes                                                      |
	 * |-----------|-------|----------|---------------------|---------------|---------------------|----------------------|------------------------------------------------------------|
	 * | 0x4C      | Play  | Client   | Reset/Clear         | Reset/Clear   | Boolean             | Boolean              | Whether to reset/clear the current advancements            |
	 * | 0x4C      | Play  | Client   | Mapping size        | Mapping size  | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x4C      | Play  | Client   | Advancement mapping | Key           | Array               | Identifier           | The identifier of the advancement                          |
	 * | 0x4C      | Play  | Client   | Advancement mapping | Value         | Array               | Advancement          | See below                                                  |
	 * | 0x4C      | Play  | Client   | List size           | List size     | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x4C      | Play  | Client   | Identifiers         | Identifiers   | Array of Identifier | Array of Identifier  | The identifiers of the advancements that should be removed |
	 * | 0x4C      | Play  | Client   | Progress size       | Progress size | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x4C      | Play  | Client   | Progress mapping    | Key           | Array               | Identifier           | The identifier of the advancement                          |
	 * | 0x4C      | Play  | Client   | Progress mapping    | Value         | Array               | Advancement progress | See below                                                  |
	 * 
	 */
}

// ID=0x4d
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x4D      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x4D      | Play  | Client   | Number Of Properties | Number Of Properties | Int        | Int                    | Number of elements in the following array             |
	 * | 0x4D      | Play  | Client   | Property             | Key                  | Array      | String (64)            | See below                                             |
	 * | 0x4D      | Play  | Client   | Property             | Value                | Array      | Double                 | See below                                             |
	 * | 0x4D      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array             |
	 * | 0x4D      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x4e
type PlayEntityEffectPkt = internal.PlayEntityEffect_757_2

// ======== END play ========
