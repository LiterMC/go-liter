
// Generated at 2023-09-02 21:26:37.473 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=7959
// Protocol: 110
// Protocol Name: 1.9.4

package packet_1_9_4

import (
	"io"
	. "github.com/kmcsr/go-liter"
	internal "github.com/kmcsr/go-liter/packets/internal"
)

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
type PlayTabCompleteServerPkt = internal.PlayTabCompleteServer_340_1

// ID=0x2
type PlayChatMessageServerPkt = internal.PlayChatMessageServer_758_0

// ID=0x3
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x4
type PlayClientSettingsPkt = internal.PlayClientSettings_754_2

// ID=0x5
type PlayConfirmTransactionServerPkt = internal.PlayConfirmTransactionServer_404_0

// ID=0x6
type PlayEnchantItemPkt = internal.PlayEnchantItem_404_0

// ID=0x7
type PlayClickWindowPkt = internal.PlayClickWindow_754_4

// ID=0x8
type PlayCloseWindowServerPkt = internal.PlayCloseWindowServer_758_0

// ID=0x9
type PlayPluginMessageServerPkt = internal.PlayPluginMessageServer_763_0

// ID=0xa
type PlayUseEntityPkt = internal.PlayUseEntity_404_0

// ID=0xb
type PlayKeepAliveServerPkt = internal.PlayKeepAliveServer_338_1

// ID=0xc
type PlayerPositionPkt = internal.PlayerPosition_758_0

// ID=0xd
type PlayerPositionAndLookServerPkt = internal.PlayerPositionAndLookServer_404_0

// ID=0xe
type PlayerLookPkt = internal.PlayerLook_404_0

// ID=0xf
type PlayerPkt = internal.Player_404_0

// ID=0x10
type PlayVehicleMoveServerPkt = internal.PlayVehicleMoveServer_758_0

// ID=0x11
type PlaySteerBoatPkt = internal.PlaySteerBoat_340_1

// ID=0x12
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_578_1

// ID=0x13
type PlayerDiggingPkt = internal.PlayerDigging_758_0

// ID=0x14
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x15
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x16
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_110_1

// ID=0x17
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChangeServer_758_0

// ID=0x18
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x19
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x1a
type PlayAnimationServerPkt = internal.PlayAnimationServer_404_0

// ID=0x1b
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x1c
type PlayerBlockPlacementPkt = internal.PlayerBlockPlacement_210_2

// ID=0x1d
type PlayUseItemPkt = internal.PlayUseItem_758_1

// ---- play: clientbound ----

// ID=0x0
type PlaySpawnObjectPkt = internal.PlaySpawnObject_404_1

// ID=0x1
type PlaySpawnExperienceOrbPkt = internal.PlaySpawnExperienceOrb_763_0

// ID=0x2
type PlaySpawnGlobalEntityPkt = internal.PlaySpawnGlobalEntity_498_0

// ID=0x3
type PlaySpawnMobPkt = internal.PlaySpawnMob_210_1

// ID=0x4
type PlaySpawnPaintingPkt = internal.PlaySpawnPainting_340_1

// ID=0x5
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_498_1

// ID=0x6
type PlayAnimationClientPkt = internal.PlayAnimation_404_1

// ID=0x7
type PlayStatisticsPkt = internal.PlayStatistics_315_4

// ID=0x8
type PlayBlockBreakAnimationPkt = internal.PlayBlockBreakAnimation_758_0

// ID=0x9
type PlayUpdateBlockEntityPkt = internal.PlayUpdateBlockEntity_498_0

// ID=0xa
type PlayBlockActionPkt = internal.PlayBlockAction_110_2

// ID=0xb
type PlayBlockChangePkt = internal.PlayBlockChange_758_0

// ID=0xc
type PlayBossBarPkt = internal.PlayBossBar_754_3

// ID=0xd
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_404_1

// ID=0xe
type PlayTabCompleteClientPkt = internal.PlayTabComplete_340_4

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
type PlayNamedSoundEffectPkt struct {
	/* All sound effect names as of 1.9.0 can be seen here. */
	SoundName String // String
	/* The category that this sound will be played from (current categories) */
	SoundCategory VarInt // VarInt Enum
	/* Effect X multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionX Int // Int
	/* Effect Y multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionY Int // Int
	/* Effect Z multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionZ Int // Int
	/* 1 is 100%, can be more */
	Volume Float // Float
	/* 63 is 100%, can be more */
	Pitch UByte // Unsigned Byte
}

var _ Packet = (*PlayNamedSoundEffectPkt)(nil)

func (p PlayNamedSoundEffectPkt)Encode(b *PacketBuilder){
	b.String(p.SoundName)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.UByte(p.Pitch)
}

func (p *PlayNamedSoundEffectPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.SoundName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.SoundCategory, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectPositionX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EffectPositionY, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EffectPositionZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Volume, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.UByte(); !ok {
		return io.EOF
	}
	return nil
}

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
type PlayChunkDataPkt = internal.PlayChunkData_316_6

// ID=0x21
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x22
type PlayParticlePkt = internal.PlayParticle_340_3

// ID=0x23
type PlayJoinGamePkt = internal.PlayJoinGame_404_4

// ID=0x24
type PlayMapPkt = internal.PlayMap_340_0

// ID=0x25
type PlayEntityRelativeMovePkt = internal.PlayEntityRelativeMove_404_0

// ID=0x26
type PlayEntityLookAndRelativeMovePkt = internal.PlayEntityLookAndRelativeMove_404_0

// ID=0x27
type PlayEntityLookPkt = internal.PlayEntityLook_404_0

// ID=0x28
type PlayEntityPkt = internal.PlayEntity_404_0

// ID=0x29
type PlayVehicleMoveClientPkt = internal.PlayVehicleMove_758_0

// ID=0x2a
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x2b
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x2c
type PlayCombatEventPkt = internal.PlayCombatEvent_335_5

// ID=0x2d
type PlayerListItemPkt = internal.PlayerListItem_315_2

// ID=0x2e
type PlayerPositionAndLookClientPkt = internal.PlayerPositionAndLook_754_1

// ID=0x2f
type PlayUseBedPkt = internal.PlayUseBed_404_0

// ID=0x30
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x31
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x32
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x33
type PlayRespawnPkt = internal.PlayRespawn_404_8

// ID=0x34
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x35
type PlayWorldBorderPkt = internal.PlayWorldBorder_316_7

// ID=0x36
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x37
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_758_0

// ID=0x38
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x39
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x3a
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x3b
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x3c
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_578_2

// ID=0x3d
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x3e
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x3f
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_340_1

// ID=0x40
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x41
type PlayTeamsPkt = internal.PlayTeams_315_8

// ID=0x42
type PlayUpdateScorePkt = internal.PlayUpdateScore_315_2

// ID=0x43
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_754_1

// ID=0x44
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x45
type PlayTitlePkt = internal.PlayTitle_210_7

// ID=0x46
type PlaySoundEffectPkt struct {
	/* ID of hardcoded sound event (events as of 1.9.0) */
	SoundID VarInt // VarInt
	/* The category that this sound will be played from (current categories) */
	SoundCategory VarInt // VarInt Enum
	/* Effect X multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionX Int // Int
	/* Effect Y multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionY Int // Int
	/* Effect Z multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part) */
	EffectPositionZ Int // Int
	/* 1.0 is 100%, capped between 0.0 and 1.0 by Notchian clients */
	Volume Float // Float
	/* Float multiplied by 63.5, capped between 0.5 and 2.0 by Notchian clients */
	Pitch UByte // Unsigned Byte
}

var _ Packet = (*PlaySoundEffectPkt)(nil)

func (p PlaySoundEffectPkt)Encode(b *PacketBuilder){
	b.VarInt(p.SoundID)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.UByte(p.Pitch)
}

func (p *PlaySoundEffectPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.SoundID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.SoundCategory, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectPositionX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EffectPositionY, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EffectPositionZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Volume, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.UByte(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x47
type PlayerListHeaderAndFooterPkt = internal.PlayerListHeaderAndFooter_758_0

// ID=0x48
type PlayCollectItemPkt = internal.PlayCollectItem_210_1

// ID=0x49
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x4a
type PlayEntityPropertiesPkt = internal.PlayEntityProperties_315_9

// ID=0x4b
type PlayEntityEffectPkt = internal.PlayEntityEffect_110_3

// ======== END play ========
