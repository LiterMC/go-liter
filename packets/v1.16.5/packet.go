
// Generated at 2023-09-05 22:06:22.506 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=16681
// Protocol: 754
// Protocol Name: 1.16.5

package packet_1_16_5

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
type LoginSuccessPkt = internal.LoginSuccess_758_1

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
type PlayQueryEntityNBTPkt = internal.PlayQueryEntityNBT_758_0

// ID=0xe
type PlayInteractEntityPkt = internal.PlayInteractEntity_758_0

// ID=0xf
type PlayGenerateStructurePkt = internal.PlayGenerateStructure_758_0

// ID=0x10
type PlayKeepAliveServerPkt = internal.PlayKeepAliveServer_763_0

// ID=0x11
type PlayLockDifficultyPkt = internal.PlayLockDifficulty_763_0

// ID=0x12
type PlayerPositionPkt = internal.PlayerPosition_758_0

// ID=0x13
type PlayerPositionAndRotationPkt = internal.PlayerPositionAndRotation_758_0

// ID=0x14
type PlayerRotationPkt = internal.PlayerRotation_758_0

// ID=0x15
type PlayerMovementPkt = internal.PlayerMovement_758_0

// ID=0x16
type PlayVehicleMoveServerPkt = internal.PlayVehicleMoveServer_758_0

// ID=0x17
type PlaySteerBoatPkt = internal.PlaySteerBoat_758_0

// ID=0x18
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x19
type PlayCraftRecipeRequestPkt = internal.PlayCraftRecipeRequest_758_0

// ID=0x1a
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_763_0

// ID=0x1b
type PlayerDiggingPkt = internal.PlayerDigging_758_0

// ID=0x1c
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x1d
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x1e
type PlaySetRecipeBookStatePkt = internal.PlaySetRecipeBookState_758_0

// ID=0x1f
type PlaySetDisplayedRecipePkt = internal.PlaySetDisplayedRecipe_758_0

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
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChangeServer_758_0

// ID=0x26
type PlayUpdateCommandBlockPkt = internal.PlayUpdateCommandBlock_758_0

// ID=0x27
type PlayUpdateCommandBlockMinecartPkt = internal.PlayUpdateCommandBlockMinecart_758_0

// ID=0x28
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x29
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
type PlayerBlockPlacementPkt = internal.PlayerBlockPlacement_758_0

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
type PlayStatisticsPkt = internal.PlayStatistics_754_1

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
type PlayBossBarPkt = internal.PlayBossBar_754_3

// ID=0xd
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_758_0

// ID=0xe
type PlayChatMessageClientPkt = internal.PlayChatMessage_758_3

// ID=0xf
type PlayTabCompleteClientPkt struct {
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
type PlayUpdateLightPkt = internal.PlayUpdateLight_754_7

// ID=0x24
type PlayJoinGamePkt = internal.PlayJoinGame_756_1

// ID=0x25
type PlayMapDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type                      | Field Type                      | Notes                                                                                                             |
	 * |-----------|-------|----------|-------------------|-------------------|---------------------------------|---------------------------------|-------------------------------------------------------------------------------------------------------------------|
	 * | 0x25      | Play  | Client   | Map ID            | Map ID            | VarInt                          | VarInt                          | Map ID of the map being modified.                                                                                 |
	 * | 0x25      | Play  | Client   | Scale             | Scale             | Byte                            | Byte                            | From 0 for a fully zoomed-in map (1x1 blocks per pixel) to 4 for a fully zoomed-out map (16x16 blocks per pixel). |
	 * | 0x25      | Play  | Client   | Tracking Position | Tracking Position | Boolean                         | Boolean                         | Specifies whether player and item frame icons are shown.                                                          |
	 * | 0x25      | Play  | Client   | Locked            | Locked            | Boolean                         | Boolean                         | True if the map has been locked in a cartography table.                                                           |
	 * | 0x25      | Play  | Client   | Icon Count        | Icon Count        | VarInt                          | VarInt                          | Number of elements in the following array.                                                                        |
	 * | 0x25      | Play  | Client   | Icon              | Type              | Array                           | VarInt enum                     | See below.                                                                                                        |
	 * | 0x25      | Play  | Client   | Icon              | X                 | Array                           | Byte                            | Map coordinates: -128 for furthest left, +127 for furthest right.                                                 |
	 * | 0x25      | Play  | Client   | Icon              | Z                 | Array                           | Byte                            | Map coordinates: -128 for highest, +127 for lowest.                                                               |
	 * | 0x25      | Play  | Client   | Icon              | Direction         | Array                           | Byte                            | 0-15.                                                                                                             |
	 * | 0x25      | Play  | Client   | Icon              | Has Display Name  | Array                           | Boolean                         |                                                                                                                   |
	 * | 0x25      | Play  | Client   | Icon              | Display Name      | Array                           | Optional Chat                   | Only present if previous Boolean is true.                                                                         |
	 * | 0x25      | Play  | Client   | Columns           | Columns           | Unsigned Byte                   | Unsigned Byte                   | Number of columns updated.                                                                                        |
	 * | 0x25      | Play  | Client   | Rows              | Rows              | Optional Unsigned Byte          | Optional Unsigned Byte          | Only if Columns is more than 0; number of rows updated.                                                           |
	 * | 0x25      | Play  | Client   | X                 | X                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column.                                               |
	 * | 0x25      | Play  | Client   | Z                 | Z                 | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row.                                                 |
	 * | 0x25      | Play  | Client   | Length            | Length            | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array.                                                    |
	 * | 0x25      | Play  | Client   | Data              | Data              | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format.                                                              |
	 * 
	 */
}

// ID=0x26
type PlayTradeListPkt = internal.PlayTradeList_754_1

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
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x31
type PlayCombatEventPkt = internal.PlayCombatEvent_754_0

// ID=0x32
type PlayerInfoPkt = internal.PlayerInfo_754_3

// ID=0x33
type PlayFacePlayerPkt = internal.PlayFacePlayer_757_0

// ID=0x34
type PlayerPositionAndLookPkt = internal.PlayerPositionAndLook_754_1

// ID=0x35
type PlayUnlockRecipesPkt = internal.PlayUnlockRecipes_758_0

// ID=0x36
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x37
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x38
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x39
type PlayRespawnPkt = internal.PlayRespawn_756_5

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
	 * | 0x3D      | Play  | Client   | Action                | Action                   | VarInt Enum | Determines the format of the rest of the packet.                                                                                                                                                                                             |
	 * | 0x3D      | Play  | Client   | Action                | Field Name               |             |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 0: set size           | Diameter                 | Double      | Length of a single side of the world border, in meters.                                                                                                                                                                                      |
	 * | 0x3D      | Play  | Client   | 1: lerp size          | Old Diameter             | Double      | Current length of a single side of the world border, in meters.                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 1: lerp size          | New Diameter             | Double      | Target length of a single side of the world border, in meters.                                                                                                                                                                               |
	 * | 0x3D      | Play  | Client   | 1: lerp size          | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x3D      | Play  | Client   | 2: set center         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 2: set center         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 3: initialize         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Old Diameter             | Double      | Current length of a single side of the world border, in meters.                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 3: initialize         | New Diameter             | Double      | Target length of a single side of the world border, in meters.                                                                                                                                                                               |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Portal Teleport Boundary | VarInt      | Resulting coordinates from a portal teleport are limited to ±value. Usually 29999984.                                                                                                                                                       |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Warning Blocks           | VarInt      | In meters.                                                                                                                                                                                                                                   |
	 * | 0x3D      | Play  | Client   | 3: initialize         | Warning Time             | VarInt      | In seconds as set by /worldborder warning time.                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 4: set warning time   | Warning Time             | VarInt      | In seconds as set by /worldborder warning time.                                                                                                                                                                                              |
	 * | 0x3D      | Play  | Client   | 5: set warning blocks | Warning Blocks           | VarInt      | In meters.                                                                                                                                                                                                                                   |
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
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_754_1

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
	 * | Packet ID | State | Bound To | Field Name                   | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|------------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x4C      | Play  | Client   | Team Name                    | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x4C      | Play  | Client   | Mode                         | Mode                | Byte                 | Determines the layout of the remaining packet.                                                                         |
	 * | 0x4C      | Play  | Client   | 0: create team               | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x4C      | Play  | Client   | 0: create team               | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team.                                     |
	 * | 0x4C      | Play  | Client   | 0: create team               | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never.                                                                      |
	 * | 0x4C      | Play  | Client   | 0: create team               | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never.                                                                            |
	 * | 0x4C      | Play  | Client   | 0: create team               | Team Color          | VarInt enum          | Used to color the name of players on the team; see below.                                                              |
	 * | 0x4C      | Play  | Client   | 0: create team               | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team.                                                      |
	 * | 0x4C      | Play  | Client   | 0: create team               | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team.                                                       |
	 * | 0x4C      | Play  | Client   | 0: create team               | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x4C      | Play  | Client   | 0: create team               | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x4C      | Play  | Client   | 1: remove team               | no fields           | no fields            |                                                                                                                        |
	 * | 0x4C      | Play  | Client   | 2: update team info          | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x4C      | Play  | Client   | 2: update team info          | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team.                                    |
	 * | 0x4C      | Play  | Client   | 2: update team info          | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x4C      | Play  | Client   | 2: update team info          | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x4C      | Play  | Client   | 2: update team info          | Team Color          | VarInt enum          | Used to color the name of players on the team; see below.                                                              |
	 * | 0x4C      | Play  | Client   | 2: update team info          | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team.                                                      |
	 * | 0x4C      | Play  | Client   | 2: update team info          | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team.                                                       |
	 * | 0x4C      | Play  | Client   | 3: add entities to team      | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x4C      | Play  | Client   | 3: add entities to team      | Entities            | Array of String (40) | Identifiers for the added entities.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x4C      | Play  | Client   | 4: remove entities from team | Entity Count        | VarInt               | Number of elements in the following array.                                                                             |
	 * | 0x4C      | Play  | Client   | 4: remove entities from team | Entities            | Array of String (40) | Identifiers for the removed entities.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

// ID=0x4d
type PlayUpdateScorePkt = internal.PlayUpdateScore_757_1

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
type PlayEntityPropertiesPkt = internal.PlayEntityProperties_754_2

// ID=0x59
type PlayEntityEffectPkt = internal.PlayEntityEffect_757_2

// ID=0x5a
type PlayDeclareRecipesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name  | Field Type | Field Type       | Notes                                                                   |
	 * |-----------|-------|----------|-------------|-------------|------------|------------------|-------------------------------------------------------------------------|
	 * | 0x5A      | Play  | Client   | Num Recipes | Num Recipes | VarInt     | VarInt           | Number of elements in the following array.                              |
	 * | 0x5A      | Play  | Client   | Recipe      | Type        | Array      | Identifier       | The recipe type, see below.                                             |
	 * | 0x5A      | Play  | Client   | Recipe      | Recipe ID   | Array      | Identifier       |                                                                         |
	 * | 0x5A      | Play  | Client   | Recipe      | Data        | Array      | Optional, varies | Additional data for the recipe.  For some types, there will be no data. |
	 * 
	 */
}

// ID=0x5b
type PlayTagsPkt = internal.PlayTags_754_2

// ======== END play ========
