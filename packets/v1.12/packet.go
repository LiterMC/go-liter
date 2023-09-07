
// Generated at 2023-09-05 22:06:22.506 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=13223
// Protocol: 335
// Protocol Name: 1.12

package packet_1_12

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

// ---- login: clientbound ----

// ID=0x0
type LoginDisconnectPkt = internal.LoginDisconnect_763_0

// ID=0x1
type LoginEncryptionRequestPkt = internal.LoginEncryptionRequest_763_0

// ID=0x2
type LoginSuccessPkt = internal.LoginSuccess_578_2

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
type PlayTabCompleteServerPkt = internal.PlayTabCompleteServer_340_1

// ID=0x3
type PlayChatMessageServerPkt = internal.PlayChatMessageServer_758_0

// ID=0x4
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x5
type PlayClientSettingsPkt = internal.PlayClientSettings_754_2

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
type PlayUseEntityPkt = internal.PlayUseEntity_404_0

// ID=0xc
type PlayKeepAliveServerPkt = internal.PlayKeepAliveServer_338_1

// ID=0xd
type PlayerPkt = internal.Player_404_0

// ID=0xe
type PlayerPositionPkt = internal.PlayerPosition_758_0

// ID=0xf
type PlayerPositionAndLookServerPkt = internal.PlayerPositionAndLookServer_404_0

// ID=0x10
type PlayerLookPkt = internal.PlayerLook_404_0

// ID=0x11
type PlayVehicleMoveServerPkt = internal.PlayVehicleMoveServer_758_0

// ID=0x12
type PlaySteerBoatPkt = internal.PlaySteerBoat_340_1

// ID=0x13
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_578_1

// ID=0x14
type PlayerDiggingPkt = internal.PlayerDigging_758_0

// ID=0x15
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x16
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x17
type PlayCraftingBookDataPkt = internal.PlayCraftingBookData_340_0

// ID=0x18
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_758_0

// ID=0x19
type PlayAdvancementTabPkt = internal.PlayAdvancementTab_758_0

// ID=0x1a
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChangeServer_758_0

// ID=0x1b
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x1c
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x1d
type PlayAnimationServerPkt = internal.PlayAnimationServer_404_0

// ID=0x1e
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x1f
type PlayerBlockPlacementPkt = internal.PlayerBlockPlacement_404_1

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
type PlayAnimationClientPkt = internal.PlayAnimation_404_1

// ID=0x7
type PlayStatisticsPkt = internal.PlayStatistics_340_4

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
type PlayTabCompleteClientPkt = internal.PlayTabComplete_340_5

// ID=0xf
type PlayChatMessageClientPkt = internal.PlayChatMessage_578_4

// ID=0x10
type PlayMultiBlockChangePkt = internal.PlayMultiBlockChange_340_3

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
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_338_1

// ID=0x20
type PlayChunkDataPkt = internal.PlayChunkData_340_5

// ID=0x21
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x22
type PlayParticlePkt = internal.PlayParticle_340_3

// ID=0x23
type PlayJoinGamePkt = internal.PlayJoinGame_404_4

// ID=0x24
type PlayMapPkt = internal.PlayMap_340_0

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
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x2c
type PlayCombatEventPkt = internal.PlayCombatEvent_335_5

// ID=0x2d
type PlayerListItemPkt = internal.PlayerListItem_335_1

// ID=0x2e
type PlayerPositionAndLookClientPkt = internal.PlayerPositionAndLook_754_1

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
type PlayRespawnPkt = internal.PlayRespawn_404_8

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
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_758_0

// ID=0x3a
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x3b
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x3c
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x3d
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x3e
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_578_2

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
type PlayerListHeaderAndFooterPkt = internal.PlayerListHeaderAndFooter_758_0

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
