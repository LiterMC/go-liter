
// Generated at 2023-09-01 20:57:33.569 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=14301
// Protocol: 401
// Protocol Name: 1.13.1

package packet_1_13_1

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

// ID=0x2
type LoginLoginPluginResponsePkt = internal.LoginLoginPluginResponse_763_0

// ---- login: clientbound ----

// ID=0x0
type LoginDisconnectPkt = internal.LoginDisconnect_763_0

// ID=0x1
type LoginEncryptionRequestPkt = internal.LoginEncryptionRequest_763_0

// ID=0x2
type LoginLoginSuccessPkt = internal.LoginLoginSuccess_578_1

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

// ID=0x2
type PlayChatMessageServerPkt = internal.PlayChatMessage_404_11

// ID=0x3
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x4
type PlayClientSettingsPkt = internal.PlayClientSettings_754_2

// ID=0x5
type PlayTabCompleteServerPkt struct {
	/* The id received in the tab completion request packet, must match or the client will ignore this packet.  Client generates this and increments it each time it sends another tab completion that doesn't get a response. */
	TransactionId VarInt // VarInt
	/* All text behind the cursor without the / (e.g. to the left of the cursor in left-to-right languages like English) */
	Text String // String (32500)
}

var _ Packet = (*PlayTabCompleteServerPkt)(nil)

func (p PlayTabCompleteServerPkt)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionId)
	b.String(p.Text)
}

func (p *PlayTabCompleteServerPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionId, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Text, ok = r.String(); !ok {
		return io.EOF
	}
}

// ID=0x6
type PlayConfirmTransactionServerPkt = internal.PlayConfirmTransaction_404_0

// ID=0x7
type PlayEnchantItemPkt = internal.PlayEnchantItem_404_0

// ID=0x8
type PlayClickWindowPkt = internal.PlayClickWindow_754_0

// ID=0x9
type PlayCloseWindowServerPkt = internal.PlayCloseWindow_758_0

// ID=0xa
type PlayPluginMessageServerPkt = internal.PlayPluginMessage_763_0

// ID=0xb
type PlayEditBookPkt = internal.PlayEditBook_755_3

// ID=0xc
type PlayQueryEntityNBTPkt = internal.PlayQueryEntityNBT_758_0

// ID=0xd
type PlayUseEntityPkt = internal.PlayUseEntity_404_0

// ID=0xe
type PlayKeepAliveServerPkt = internal.PlayKeepAlive_763_0

// ID=0xf
type PlayPlayerPkt = internal.PlayPlayer_404_0

// ID=0x10
type PlayPlayerPositionPkt = internal.PlayPlayerPosition_758_0

// ID=0x11
type PlayPlayerPositionAndLookServerPkt = internal.PlayPlayerPositionAndLook_404_2

// ID=0x12
type PlayPlayerLookPkt = internal.PlayPlayerLook_404_0

// ID=0x13
type PlayVehicleMoveServerPkt = internal.PlayVehicleMove_758_0

// ID=0x14
type PlaySteerBoatPkt = internal.PlaySteerBoat_401_1

// ID=0x15
type PlayPickItemPkt = internal.PlayPickItem_763_0

// ID=0x16
type PlayCraftRecipeRequestPkt = internal.PlayCraftRecipeRequest_758_0

// ID=0x17
type PlayPlayerAbilitiesServerPkt = internal.PlayPlayerAbilities_404_12

// ID=0x18
type PlayPlayerDiggingPkt = internal.PlayPlayerDigging_758_0

// ID=0x19
type PlayEntityActionPkt = internal.PlayEntityAction_758_0

// ID=0x1a
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0x1b
type PlayRecipeBookDataPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name            | Field Name                    | Field Type | Notes                                                                           |
	 * |-----------|-------|----------|-----------------------|-------------------------------|------------|---------------------------------------------------------------------------------|
	 * | 0x1B      | Play  | Server   | Type                  | Type                          | VarInt     | Determines the format of the rest of the packet                                 |
	 * | 0x1B      | Play  | Server   | Type                  | Field Name                    |            |                                                                                 |
	 * | 0x1B      | Play  | Server   | 0: Displayed Recipe   | Recipe ID                     | Identifier | A recipe ID                                                                     |
	 * | 0x1B      | Play  | Server   | 1: Recipe Book States | Crafting Recipe Book Open     | Boolean    | Whether the player has the crafting recipe book currently opened/active.        |
	 * | 0x1B      | Play  | Server   | 1: Recipe Book States | Crafting Recipe Filter Active | Boolean    | Whether the player has the crafting recipe book filter option currently active. |
	 * | 0x1B      | Play  | Server   | 1: Recipe Book States | Smelting Recipe Book Open     | Boolean    | Whether the player has the smelting recipe book currently opened/active.        |
	 * | 0x1B      | Play  | Server   | 1: Recipe Book States | Smelting Recipe Filter Active | Boolean    | Whether the player has the smelting recipe book filter option currently active. |
	 * 
	 */
}

// ID=0x1c
type PlayNameItemPkt = internal.PlayNameItem_758_0

// ID=0x1d
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_758_0

// ID=0x1e
type PlayAdvancementTabPkt = internal.PlayAdvancementTab_758_0

// ID=0x1f
type PlaySelectTradePkt = internal.PlaySelectTrade_763_0

// ID=0x20
type PlaySetBeaconEffectPkt = internal.PlaySetBeaconEffect_758_2

// ID=0x21
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChange_404_5

// ID=0x22
type PlayUpdateCommandBlockPkt = internal.PlayUpdateCommandBlock_758_0

// ID=0x23
type PlayUpdateCommandBlockMinecartPkt = internal.PlayUpdateCommandBlockMinecart_758_0

// ID=0x24
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x25
type PlayUpdateStructureBlockPkt = internal.PlayUpdateStructureBlock_758_0

// ID=0x26
type PlayUpdateSignPkt = internal.PlayUpdateSign_762_1

// ID=0x27
type PlayAnimationServerPkt = internal.PlayAnimation_404_2

// ID=0x28
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x29
type PlayPlayerBlockPlacementPkt = internal.PlayPlayerBlockPlacement_404_1

// ID=0x2a
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
type PlaySpawnPaintingPkt = internal.PlaySpawnPainting_758_0

// ID=0x5
type PlaySpawnPlayerPkt = internal.PlaySpawnPlayer_498_1

// ID=0x6
type PlayAnimationClientPkt = internal.PlayAnimation_404_1

// ID=0x7
type PlayStatisticsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name   | Field Type              | Field Type | Notes                                     |
	 * |-----------|-------|----------|------------|--------------|-------------------------|------------|-------------------------------------------|
	 * | 0x07      | Play  | Client   | Count      | Count        | VarInt                  | VarInt     | Number of elements in the following array |
	 * | 0x07      | Play  | Client   | Statistic  | Category ID  | Array                   | VarInt     | See below                                 |
	 * | 0x07      | Play  | Client   | Statistic  | Statistic ID | Array                   | VarInt     | See below                                 |
	 * | 0x07      | Play  | Client   | Value      | VarInt       | The amount to set it to |            |                                           |
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
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_404_1

// ID=0xe
type PlayChatMessageClientPkt = internal.PlayChatMessage_404_10

// ID=0xf
type PlayMultiBlockChangePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name   | Field Name          | Field Type | Field Type    | Notes                                                                                                                                                                  |
	 * |-----------|-------|----------|--------------|---------------------|------------|---------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0F      | Play  | Client   | Chunk X      | Chunk X             | Int        | Int           | Chunk X coordinate                                                                                                                                                     |
	 * | 0x0F      | Play  | Client   | Chunk Z      | Chunk Z             | Int        | Int           | Chunk Z coordinate                                                                                                                                                     |
	 * | 0x0F      | Play  | Client   | Record Count | Record Count        | VarInt     | VarInt        | Number of elements in the following array, i.e. the number of blocks affected                                                                                          |
	 * | 0x0F      | Play  | Client   | Record       | Horizontal Position | Array      | Unsigned Byte | The 4 most significant bits (0xF0) encode the X coordinate, relative to the chunk. The 4 least significant bits (0x0F) encode the Z coordinate, relative to the chunk. |
	 * | 0x0F      | Play  | Client   | Record       | Y Coordinate        | Array      | Unsigned Byte | Y coordinate of the block                                                                                                                                              |
	 * | 0x0F      | Play  | Client   | Record       | Block ID            | Array      | VarInt        | The new block state ID for the block as given in the global palette. See that section for more information.                                                            |
	 * 
	 */
}

// ID=0x10
type PlayTabCompleteClientPkt struct {
	/* VarInt */
	Start Unknown_start // Start
	/* VarInt */
	Length Unknown_length // Length
	/* VarInt */
	Count Unknown_count // Count
	/* Array */
	Matches Unknown_match // Match
	/* Array */
	Matches Unknown_has_tooltip // Has tooltip
	/* Array */
	Matches Unknown_tooltip // Tooltip
}

var _ Packet = (*PlayTabCompleteClientPkt)(nil)

func (p PlayTabCompleteClientPkt)Encode(b *PacketBuilder){
	b.Encode(p.Start)
	b.Encode(p.Length)
	b.Encode(p.Count)
	b.Encode(p.Matches)
	b.Encode(p.Matches)
	b.Encode(p.Matches)
}

func (p *PlayTabCompleteClientPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Start.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Length.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Count.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Matches.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Matches.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Matches.DecodeFrom(r); err != nil {
		return
	}
}

// ID=0x11
type PlayDeclareCommandsPkt = internal.PlayDeclareCommands_758_0

// ID=0x12
type PlayConfirmTransactionClientPkt = internal.PlayConfirmTransaction_404_0

// ID=0x13
type PlayCloseWindowClientPkt = internal.PlayCloseWindow_758_0

// ID=0x14
type PlayOpenWindowPkt = internal.PlayOpenWindow_404_1

// ID=0x15
type PlayWindowItemsPkt = internal.PlayWindowItems_755_1

// ID=0x16
type PlayWindowPropertyPkt = internal.PlayWindowProperty_758_0

// ID=0x17
type PlaySetSlotPkt = internal.PlaySetSlot_755_1

// ID=0x18
type PlaySetCooldownPkt = internal.PlaySetCooldown_763_0

// ID=0x19
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_763_0

// ID=0x1a
type PlayNamedSoundEffectPkt = internal.PlayNamedSoundEffect_758_0

// ID=0x1b
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x1c
type PlayEntityStatusPkt = internal.PlayEntityStatus_758_0

// ID=0x1d
type PlayNBTQueryResponsePkt = internal.PlayNBTQueryResponse_758_0

// ID=0x1e
type PlayExplosionPkt = internal.PlayExplosion_498_3

// ID=0x1f
type PlayUnloadChunkPkt = internal.PlayUnloadChunk_763_0

// ID=0x20
type PlayChangeGameStatePkt = internal.PlayChangeGameState_758_0

// ID=0x21
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_763_0

// ID=0x22
type PlayChunkDataPkt = internal.PlayChunkData_401_5

// ID=0x23
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x24
type PlayParticlePkt = internal.PlayParticle_498_2

// ID=0x25
type PlayJoinGamePkt = internal.PlayJoinGame_404_4

// ID=0x26
type PlayMapPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name        | Field Type                      | Field Type                               | Notes                                                                                                      |
	 * |-----------|-------|----------|-------------------|-------------------|---------------------------------|------------------------------------------|------------------------------------------------------------------------------------------------------------|
	 * | 0x26      | Play  | Client   | Map ID            | Map ID            | VarInt                          | VarInt                                   | Map ID of the map being modified                                                                           |
	 * | 0x26      | Play  | Client   | Scale             | Scale             | Byte                            | Byte                                     | From 0 for a fully zoomed-in map (1 block per pixel) to 4 for a fully zoomed-out map (16 blocks per pixel) |
	 * | 0x26      | Play  | Client   | Tracking Position | Tracking Position | Boolean                         | Boolean                                  | Specifies whether the icons are shown                                                                      |
	 * | 0x26      | Play  | Client   | Icon Count        | Icon Count        | VarInt                          | VarInt                                   | Number of elements in the following array                                                                  |
	 * | 0x26      | Play  | Client   | Icon              | Type              | Array                           | VarInt enum                              | See below                                                                                                  |
	 * | 0x26      | Play  | Client   | Icon              | X                 | Array                           | Byte                                     |                                                                                                            |
	 * | 0x26      | Play  | Client   | Icon              | Z                 | Array                           | Byte                                     |                                                                                                            |
	 * | 0x26      | Play  | Client   | Icon              | Direction         | Array                           | Byte                                     | 0-15                                                                                                       |
	 * | 0x26      | Play  | Client   | Icon              | Has Display Name  | Array                           | Boolean                                  |                                                                                                            |
	 * | 0x26      | Play  | Client   | Display Name      | Optional Chat     | Array                           | Only present if previous Boolean is true |                                                                                                            |
	 * | 0x26      | Play  | Client   | Columns           | Columns           | Byte                            | Byte                                     | Number of columns updated                                                                                  |
	 * | 0x26      | Play  | Client   | Rows              | Rows              | Optional Byte                   | Optional Byte                            | Only if Columns is more than 0; number of rows updated                                                     |
	 * | 0x26      | Play  | Client   | X                 | X                 | Optional Byte                   | Optional Byte                            | Only if Columns is more than 0; x offset of the westernmost column                                         |
	 * | 0x26      | Play  | Client   | Z                 | Z                 | Optional Byte                   | Optional Byte                            | Only if Columns is more than 0; z offset of the northernmost row                                           |
	 * | 0x26      | Play  | Client   | Length            | Length            | Optional VarInt                 | Optional VarInt                          | Only if Columns is more than 0; length of the following array                                              |
	 * | 0x26      | Play  | Client   | Data              | Data              | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte          | Only if Columns is more than 0; see Map item format                                                        |
	 * 
	 */
}

// ID=0x27
type PlayEntityPkt = internal.PlayEntity_404_0

// ID=0x28
type PlayEntityRelativeMovePkt = internal.PlayEntityRelativeMove_404_0

// ID=0x29
type PlayEntityLookAndRelativeMovePkt = internal.PlayEntityLookAndRelativeMove_404_0

// ID=0x2a
type PlayEntityLookPkt = internal.PlayEntityLook_404_0

// ID=0x2b
type PlayVehicleMoveClientPkt = internal.PlayVehicleMove_758_0

// ID=0x2c
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x2d
type PlayCraftRecipeResponsePkt = internal.PlayCraftRecipeResponse_758_0

// ID=0x2e
type PlayPlayerAbilitiesClientPkt = internal.PlayPlayerAbilities_404_11

// ID=0x2f
type PlayCombatEventPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name      | Field Name | Field Type  | Notes                                         |
	 * |-----------|-------|----------|-----------------|------------|-------------|-----------------------------------------------|
	 * | 0x2F      | Play  | Client   | Event           | Event      | VarInt Enum | Determines the layout of the remaining packet |
	 * | 0x2F      | Play  | Client   | Event           | Field Name |             |                                               |
	 * | 0x2F      | Play  | Client   | 0: enter combat | no fields  | no fields   |                                               |
	 * | 0x2F      | Play  | Client   | 1: end combat   | Duration   | VarInt      |                                               |
	 * | 0x2F      | Play  | Client   | 1: end combat   | Entity ID  | Int         |                                               |
	 * | 0x2F      | Play  | Client   | 2: entity dead  | Player ID  | VarInt      |                                               |
	 * | 0x2F      | Play  | Client   | 2: entity dead  | Entity ID  | Int         |                                               |
	 * | 0x2F      | Play  | Client   | 2: entity dead  | Message    | Chat        |                                               |
	 * 
	 */
}

// ID=0x30
type PlayPlayerListItemPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type              | Notes                                                   |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-------------------------|---------------------------------------------------------|
	 * | 0x30      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt                  | Determines the rest of the Player format after the UUID |
	 * | 0x30      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x30      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID                    |                                                         |
	 * | 0x30      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                         |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String (16)   | String (16)             |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt                  | Number of elements in the following array               |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String (32767)          |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String (32767)          |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean                 |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String (32767) | Only if Is Signed is true                               |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds                                |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only if Has Display Name is true                        |
	 * | 0x30      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt                  |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt                  | Measured in milliseconds                                |
	 * | 0x30      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean                 |                                                         |
	 * | 0x30      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat           | Only send if Has Display Name is true                   |
	 * | 0x30      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields               |                                                         |
	 * 
	 */
}

// ID=0x31
type PlayFacePlayerPkt = internal.PlayFacePlayer_757_0

// ID=0x32
type PlayPlayerPositionAndLookClientPkt = internal.PlayPlayerPositionAndLook_754_1

// ID=0x33
type PlayUseBedPkt = internal.PlayUseBed_404_0

// ID=0x34
type PlayUnlockRecipesPkt = internal.PlayUnlockRecipes_578_2

// ID=0x35
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x36
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x37
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x38
type PlayRespawnPkt = internal.PlayRespawn_404_7

// ID=0x39
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x3a
type PlaySelectAdvancementTabPkt = internal.PlaySelectAdvancementTab_758_0

// ID=0x3b
type PlayWorldBorderPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name            | Field Name               | Field Type  | Notes                                                                                                                                                                                                                                        |
	 * |-----------|-------|----------|-----------------------|--------------------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x3B      | Play  | Client   | Action                | Action                   | VarInt Enum | Determines the format of the rest of the packet                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | Action                | Field Name               |             |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 0: set size           | Diameter                 | Double      | Length of a single side of the world border, in meters                                                                                                                                                                                       |
	 * | 0x3B      | Play  | Client   | 1: lerp size          | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x3B      | Play  | Client   | 1: lerp size          | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x3B      | Play  | Client   | 1: lerp size          | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x3B      | Play  | Client   | 2: set center         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 2: set center         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 3: initialize         | X                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Z                        | Double      |                                                                                                                                                                                                                                              |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Old Diameter             | Double      | Current length of a single side of the world border, in meters                                                                                                                                                                               |
	 * | 0x3B      | Play  | Client   | 3: initialize         | New Diameter             | Double      | Target length of a single side of the world border, in meters                                                                                                                                                                                |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Speed                    | VarLong     | Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Portal Teleport Boundary | VarInt      | Resulting coordinates from a portal teleport are limited to ±value. Usually 29999984.                                                                                                                                                       |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x3B      | Play  | Client   | 3: initialize         | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * | 0x3B      | Play  | Client   | 4: set warning time   | Warning Time             | VarInt      | In seconds as set by /worldborder warning time                                                                                                                                                                                               |
	 * | 0x3B      | Play  | Client   | 5: set warning blocks | Warning Blocks           | VarInt      | In meters                                                                                                                                                                                                                                    |
	 * 
	 */
}

// ID=0x3c
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x3d
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_404_4

// ID=0x3e
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x3f
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x40
type PlayAttachEntityPkt = internal.PlayAttachEntity_758_0

// ID=0x41
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x42
type PlayEntityEquipmentPkt = internal.PlayEntityEquipment_578_0

// ID=0x43
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x44
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x45
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_758_0

// ID=0x46
type PlaySetPassengersPkt = internal.PlaySetPassengers_763_0

// ID=0x47
type PlayTeamsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name                  | Field Name          | Field Type           | Notes                                                                                                                  |
	 * |-----------|-------|----------|-----------------------------|---------------------|----------------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x47      | Play  | Client   | Team Name                   | Team Name           | String (16)          | A unique name for the team. (Shared with scoreboard).                                                                  |
	 * | 0x47      | Play  | Client   | Mode                        | Mode                | Byte                 | Determines the layout of the remaining packet                                                                          |
	 * | 0x47      | Play  | Client   | 0: create team              | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x47      | Play  | Client   | 0: create team              | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible players on same team                                      |
	 * | 0x47      | Play  | Client   | 0: create team              | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x47      | Play  | Client   | 0: create team              | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x47      | Play  | Client   | 0: create team              | Formatting          | VarInt enum          | See below                                                                                                              |
	 * | 0x47      | Play  | Client   | 0: create team              | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team                                                       |
	 * | 0x47      | Play  | Client   | 0: create team              | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team                                                        |
	 * | 0x47      | Play  | Client   | 0: create team              | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x47      | Play  | Client   | 0: create team              | Entities            | Array of String (40) | Identifiers for the entities in this team.  For players, this is their username; for other entities, it is their UUID. |
	 * | 0x47      | Play  | Client   | 1: remove team              | no fields           | no fields            |                                                                                                                        |
	 * | 0x47      | Play  | Client   | 2: update team info         | Team Display Name   | Chat                 |                                                                                                                        |
	 * | 0x47      | Play  | Client   | 2: update team info         | Friendly Flags      | Byte                 | Bit mask. 0x01: Allow friendly fire, 0x02: can see invisible entities on same team                                     |
	 * | 0x47      | Play  | Client   | 2: update team info         | Name Tag Visibility | String Enum (32)     | always, hideForOtherTeams, hideForOwnTeam, never                                                                       |
	 * | 0x47      | Play  | Client   | 2: update team info         | Collision Rule      | String Enum (32)     | always, pushOtherTeams, pushOwnTeam, never                                                                             |
	 * | 0x47      | Play  | Client   | 2: update team info         | Formatting          | VarInt enum          | See below                                                                                                              |
	 * | 0x47      | Play  | Client   | 2: update team info         | Team Prefix         | Chat                 | Displayed before the names of players that are part of this team                                                       |
	 * | 0x47      | Play  | Client   | 2: update team info         | Team Suffix         | Chat                 | Displayed after the names of players that are part of this team                                                        |
	 * | 0x47      | Play  | Client   | 3: add players to team      | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x47      | Play  | Client   | 3: add players to team      | Entities            | Array of String (40) | Identifiers for the entities added.  For players, this is their username; for other entities, it is their UUID.        |
	 * | 0x47      | Play  | Client   | 4: remove players from team | Entity Count        | VarInt               | Number of elements in the following array                                                                              |
	 * | 0x47      | Play  | Client   | 4: remove players from team | Entities            | Array of String (40) | Identifiers for the entities removed.  For players, this is their username; for other entities, it is their UUID.      |
	 * 
	 */
}

// ID=0x48
type PlayUpdateScorePkt = internal.PlayUpdateScore_757_1

// ID=0x49
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_754_1

// ID=0x4a
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x4b
type PlayTitlePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name      | Field Type  | Notes                                                                                                                                                                            |
	 * |-----------|-------|----------|--------------------------|-----------------|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x4B      | Play  | Client   | Action                   | Action          | VarInt Enum |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | Action                   | Field Name      |             |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | 0: set title             | Title Text      | Chat        |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | 1: set subtitle          | Subtitle Text   | Chat        |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | 2: set action bar        | Action bar text | Chat        | Displays a message above the hotbar (the same as position 2 in Chat Message (clientbound), except that it correctly renders formatted chat. See MC-119145 for more information.) |
	 * | 0x4B      | Play  | Client   | 3: set times and display | Fade In         | Int         | Ticks to spend fading in                                                                                                                                                         |
	 * | 0x4B      | Play  | Client   | 3: set times and display | Stay            | Int         | Ticks to keep the title displayed                                                                                                                                                |
	 * | 0x4B      | Play  | Client   | 3: set times and display | Fade Out        | Int         | Ticks to spend out, not when to start fading out                                                                                                                                 |
	 * | 0x4B      | Play  | Client   | 4: hide                  | no fields       | no fields   |                                                                                                                                                                                  |
	 * | 0x4B      | Play  | Client   | 5: reset                 | no fields       | no fields   |                                                                                                                                                                                  |
	 * 
	 */
}

// ID=0x4c
type PlayStopSoundPkt = internal.PlayStopSound_763_0

// ID=0x4d
type PlaySoundEffectPkt = internal.PlaySoundEffect_758_2

// ID=0x4e
type PlayPlayerListHeaderAndFooterPkt = internal.PlayPlayerListHeaderAndFooter_758_0

// ID=0x4f
type PlayCollectItemPkt = internal.PlayCollectItem_758_0

// ID=0x50
type PlayEntityTeleportPkt = internal.PlayEntityTeleport_758_0

// ID=0x51
type PlayAdvancementsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name          | Field Name    | Field Type          | Field Type           | Notes                                                      |
	 * |-----------|-------|----------|---------------------|---------------|---------------------|----------------------|------------------------------------------------------------|
	 * | 0x51      | Play  | Client   | Reset/Clear         | Reset/Clear   | Boolean             | Boolean              | Whether to reset/clear the current advancements            |
	 * | 0x51      | Play  | Client   | Mapping size        | Mapping size  | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x51      | Play  | Client   | Advancement mapping | Key           | Array               | Identifier           | The identifier of the advancement                          |
	 * | 0x51      | Play  | Client   | Advancement mapping | Value         | Array               | Advancement          | See below                                                  |
	 * | 0x51      | Play  | Client   | List size           | List size     | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x51      | Play  | Client   | Identifiers         | Identifiers   | Array of Identifier | Array of Identifier  | The identifiers of the advancements that should be removed |
	 * | 0x51      | Play  | Client   | Progress size       | Progress size | VarInt              | VarInt               | Size of the following array                                |
	 * | 0x51      | Play  | Client   | Progress mapping    | Key           | Array               | Identifier           | The identifier of the advancement                          |
	 * | 0x51      | Play  | Client   | Progress mapping    | Value         | Array               | Advancement progress | See below                                                  |
	 * 
	 */
}

// ID=0x52
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                                 |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------------------|
	 * | 0x52      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                                       |
	 * | 0x52      | Play  | Client   | Number Of Properties | Number Of Properties | Int        | Int                    | Number of elements in the following array             |
	 * | 0x52      | Play  | Client   | Property             | Key                  | Array      | String (64)            | See below                                             |
	 * | 0x52      | Play  | Client   | Property             | Value                | Array      | Double                 | See below                                             |
	 * | 0x52      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array             |
	 * | 0x52      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers. Modifier Data defined below. |
	 * 
	 */
}

// ID=0x53
type PlayEntityEffectPkt = internal.PlayEntityEffect_757_2

// ID=0x54
type PlayDeclareRecipesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name  | Field Type | Field Type       | Notes                                                                   |
	 * |-----------|-------|----------|-------------|-------------|------------|------------------|-------------------------------------------------------------------------|
	 * | 0x54      | Play  | Client   | Num Recipes | Num Recipes | VarInt     | VarInt           | Number of elements in the following array                               |
	 * | 0x54      | Play  | Client   | Recipe      | Recipe ID   | Array      | Identifier       |                                                                         |
	 * | 0x54      | Play  | Client   | Recipe      | Type        | Array      | String           | The recipe type, see below                                              |
	 * | 0x54      | Play  | Client   | Recipe      | Data        | Array      | Optional, varies | Additional data for the recipe.  For some types, there will be no data. |
	 * 
	 */
}

// ID=0x55
type PlayTagsPkt = internal.PlayTags_404_1

// ======== END play ========