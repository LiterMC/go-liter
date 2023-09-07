
// Generated at 2023-09-05 22:06:22.506 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=16866
// Protocol: 755
// Protocol Name: 1.17

package packet_1_17

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
type PlayClientSettingsPkt = internal.PlayClientSettings_756_1

// ID=0x6
type PlayTabCompleteServerPkt = internal.PlayTabCompleteServer_758_0

// ID=0x7
type PlayClickWindowButtonPkt = internal.PlayClickWindowButton_758_0

// ID=0x8
type PlayClickWindowPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name          | Field Type    | Field Type    | Notes                                                                                                          |
	 * |-----------|-------|----------|---------------------|---------------------|---------------|---------------|----------------------------------------------------------------------------------------------------------------|
	 * | 0x08      | Play  | Server   | Window ID           | Window ID           | Unsigned Byte | Unsigned Byte | The ID of the window which was clicked. 0 for player inventory.                                                |
	 * | 0x08      | Play  | Server   | Slot                | Slot                | Short         | Short         | The clicked slot number, see below.                                                                            |
	 * | 0x08      | Play  | Server   | Button              | Button              | Byte          | Byte          | The button used in the click, see below.                                                                       |
	 * | 0x08      | Play  | Server   | Mode                | Mode                | VarInt Enum   | VarInt Enum   | Inventory operation mode, see below.                                                                           |
	 * | 0x08      | Play  | Server   | Length of the array | Length of the array | VarInt        | VarInt        |                                                                                                                |
	 * | 0x08      | Play  | Server   | Array of slots      | Slot number         | Array         | Short         |                                                                                                                |
	 * | 0x08      | Play  | Server   | Array of slots      | Slot data           | Array         | Slot          | New data for this slot                                                                                         |
	 * | 0x08      | Play  | Server   | Clicked item        | Clicked item        | Slot          | Slot          | The clicked slot. Has to be empty (item ID = -1) for drop mode. Is always empty for mode 2 and mode 5 packets. |
	 * 
	 */
}

// ID=0x9
type PlayCloseWindowServerPkt = internal.PlayCloseWindowServer_758_0

// ID=0xa
type PlayPluginMessageServerPkt = internal.PlayPluginMessageServer_763_0

// ID=0xb
type PlayEditBookPkt = internal.PlayEditBook_755_4

// ID=0xc
type PlayQueryEntityNBTPkt = internal.PlayQueryEntityNBT_758_0

// ID=0xd
type PlayInteractEntityPkt = internal.PlayInteractEntity_758_0

// ID=0xe
type PlayGenerateStructurePkt = internal.PlayGenerateStructure_758_0

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
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_763_0

// ID=0x1a
type PlayerDiggingPkt = internal.PlayerDigging_758_0

// ID=0x1b
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x1c
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x1d
type PlayPongPkt = internal.PlayPong_763_0

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
type PlaySculkVibrationSignalPkt = internal.PlaySculkVibrationSignal_758_0

// ID=0x6
type PlayEntityAnimationPkt = internal.PlayEntityAnimation_763_0

// ID=0x7
type PlayStatisticsPkt = internal.PlayStatistics_758_0

// ID=0x8
type PlayAcknowledgePlayerDiggingPkt = internal.PlayAcknowledgePlayerDigging_758_0

// ID=0x9
type PlayBlockBreakAnimationPkt = internal.PlayBlockBreakAnimation_758_0

// ID=0xa
type PlayBlockEntityDataPkt = internal.PlayBlockEntityData_756_1

// ID=0xb
type PlayBlockActionPkt = internal.PlayBlockAction_758_1

// ID=0xc
type PlayBlockChangePkt = internal.PlayBlockChange_758_0

// ID=0xd
type PlayBossBarPkt = internal.PlayBossBar_758_2

// ID=0xe
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_758_0

// ID=0xf
type PlayChatMessageClientPkt = internal.PlayChatMessage_758_3

// ID=0x10
type PlayClearTitlesPkt = internal.PlayClearTitles_763_0

// ID=0x11
type PlayTabCompleteClientPkt = internal.PlayTabComplete_758_0

// ID=0x12
type PlayDeclareCommandsPkt = internal.PlayDeclareCommands_758_0

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
type PlayExplosionPkt = internal.PlayExplosion_760_1

// ID=0x1d
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x1e
type PlayChangeGameStatePkt = internal.PlayChangeGameState_758_0

// ID=0x1f
type PlayOpenHorseWindowPkt = internal.PlayOpenHorseWindow_756_1

// ID=0x20
type PlayInitializeWorldBorderPkt = internal.PlayInitializeWorldBorder_763_0

// ID=0x21
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x22
type PlayChunkDataPkt = internal.PlayChunkData_756_0

// ID=0x23
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x24
type PlayParticlePkt = internal.PlayParticle_758_1

// ID=0x25
type PlayUpdateLightPkt = internal.PlayUpdateLight_756_6

// ID=0x26
type PlayJoinGamePkt = internal.PlayJoinGame_756_1

// ID=0x27
type PlayMapDataPkt = internal.PlayMapData_756_5

// ID=0x28
type PlayTradeListPkt = internal.PlayTradeList_758_0

// ID=0x29
type PlayEntityPositionPkt = internal.PlayEntityPosition_758_0

// ID=0x2a
type PlayEntityPositionAndRotationPkt = internal.PlayEntityPositionAndRotation_758_0

// ID=0x2b
type PlayEntityRotationPkt = internal.PlayEntityRotation_758_0

// ID=0x2c
type PlayVehicleMoveClientPkt = internal.PlayVehicleMove_758_0

// ID=0x2d
type PlayOpenBookPkt = internal.PlayOpenBook_763_0

// ID=0x2e
type PlayOpenWindowPkt = internal.PlayOpenWindow_758_0

// ID=0x2f
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x30
type PlayPingPkt = internal.PlayPing_763_0

// ID=0x31
type PlayCraftRecipeResponsePkt = internal.PlayCraftRecipeResponse_758_0

// ID=0x32
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x33
type PlayEndCombatEventPkt = internal.PlayEndCombatEvent_758_0

// ID=0x34
type PlayEnterCombatEventPkt = internal.PlayEnterCombatEvent_758_0

// ID=0x35
type PlayDeathCombatEventPkt = internal.PlayDeathCombatEvent_758_0

// ID=0x36
type PlayerInfoPkt = internal.PlayerInfo_758_2

// ID=0x37
type PlayFacePlayerPkt = internal.PlayFacePlayer_757_0

// ID=0x38
type PlayerPositionAndLookPkt = internal.PlayerPositionAndLook_758_0

// ID=0x39
type PlayUnlockRecipesPkt = internal.PlayUnlockRecipes_758_0

// ID=0x3a
type PlayDestroyEntityPkt struct {
	/* The ID of the entity to be destroyed */
	EntityID VarInt // VarInt
}

var _ Packet = (*PlayDestroyEntityPkt)(nil)

func (p PlayDestroyEntityPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
}

func (p *PlayDestroyEntityPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x3b
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x3c
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_758_0

// ID=0x3d
type PlayRespawnPkt = internal.PlayRespawn_756_5

// ID=0x3e
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x3f
type PlayMultiBlockChangePkt = internal.PlayMultiBlockChange_758_0

// ID=0x40
type PlaySelectAdvancementTabPkt = internal.PlaySelectAdvancementTab_758_0

// ID=0x41
type PlayActionBarPkt = internal.PlayActionBar_758_0

// ID=0x42
type PlayWorldBorderCenterPkt = internal.PlayWorldBorderCenter_758_0

// ID=0x43
type PlayWorldBorderLerpSizePkt = internal.PlayWorldBorderLerpSize_758_0

// ID=0x44
type PlayWorldBorderSizePkt = internal.PlayWorldBorderSize_758_0

// ID=0x45
type PlayWorldBorderWarningDelayPkt = internal.PlayWorldBorderWarningDelay_758_0

// ID=0x46
type PlayWorldBorderWarningReachPkt = internal.PlayWorldBorderWarningReach_758_0

// ID=0x47
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x48
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_758_0

// ID=0x49
type PlayUpdateViewPositionPkt = internal.PlayUpdateViewPosition_758_0

// ID=0x4a
type PlayUpdateViewDistancePkt = internal.PlayUpdateViewDistance_758_0

// ID=0x4b
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_758_0

// ID=0x4c
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x4d
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x4e
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x4f
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x50
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_758_0

// ID=0x51
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x52
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x53
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_758_0

// ID=0x54
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x55
type PlayTeamsPkt = internal.PlayTeams_758_0

// ID=0x56
type PlayUpdateScorePkt = internal.PlayUpdateScore_757_1

// ID=0x57
type PlaySetTitleSubTitlePkt = internal.PlaySetTitleSubTitle_758_0

// ID=0x58
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x59
type PlaySetTitleTextPkt = internal.PlaySetTitleText_763_0

// ID=0x5a
type PlaySetTitleTimesPkt = internal.PlaySetTitleTimes_758_0

// ID=0x5b
type PlayEntitySoundEffectPkt = internal.PlayEntitySoundEffect_760_1

// ID=0x5c
type PlaySoundEffectPkt = internal.PlaySoundEffect_758_2

// ID=0x5d
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x5e
type PlayerListHeaderAndFooterPkt = internal.PlayerListHeaderAndFooter_758_0

// ID=0x5f
type PlayNBTQueryResponsePkt = internal.PlayNBTQueryResponse_758_0

// ID=0x60
type PlayCollectItemPkt = internal.PlayCollectItem_758_0

// ID=0x61
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x62
type PlayAdvancementsPkt = internal.PlayAdvancements_756_1

// ID=0x63
type PlayEntityPropertiesPkt = internal.PlayEntityProperties_756_1

// ID=0x64
type PlayEntityEffectPkt = internal.PlayEntityEffect_757_2

// ID=0x65
type PlayDeclareRecipesPkt = internal.PlayDeclareRecipes_756_1

// ID=0x66
type PlayTagsPkt = internal.PlayTags_756_1

// ======== END play ========
