
// Generated at 2023-09-05 22:06:22.506 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=7368
// Protocol: 47
// Protocol Name: 1.8.9

package packet_1_8_9

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
type PlayKeepAliveServerPkt = internal.PlayKeepAliveServer_338_1

// ID=0x1
type PlayChatMessageServerPkt = internal.PlayChatMessageServer_758_0

// ID=0x2
type PlayUseEntityPkt struct {
	Target VarInt // VarInt
	/* 0: interact, 1: attack, 2: interact at */
	Type VarInt // VarInt
	/* Only if Type is interact at */
	TargetX Optional[Float] // Optional Float
	/* Only if Type is interact at */
	TargetY Optional[Float] // Optional Float
	/* Only if Type is interact at */
	TargetZ Optional[Float] // Optional Float
}

var _ Packet = (*PlayUseEntityPkt)(nil)

func (p PlayUseEntityPkt)Encode(b *PacketBuilder){
	b.VarInt(p.Target)
	b.VarInt(p.Type)
	if p.TargetX.Ok = TODO; p.TargetX.Ok {
		b.Float(p.TargetX.V)
	}
	if p.TargetY.Ok = TODO; p.TargetY.Ok {
		b.Float(p.TargetY.V)
	}
	if p.TargetZ.Ok = TODO; p.TargetZ.Ok {
		b.Float(p.TargetZ.V)
	}
}

func (p *PlayUseEntityPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.Target, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.TargetX.Ok = TODO; p.TargetX.Ok {
		if p.TargetX.V, ok = r.Float(); !ok {
			return io.EOF
		}
	}
	if p.TargetY.Ok = TODO; p.TargetY.Ok {
		if p.TargetY.V, ok = r.Float(); !ok {
			return io.EOF
		}
	}
	if p.TargetZ.Ok = TODO; p.TargetZ.Ok {
		if p.TargetZ.V, ok = r.Float(); !ok {
			return io.EOF
		}
	}
	return nil
}

// ID=0x3
type PlayerPkt = internal.Player_404_0

// ID=0x4
type PlayerPositionPkt = internal.PlayerPosition_758_0

// ID=0x5
type PlayerLookPkt = internal.PlayerLook_404_0

// ID=0x6
type PlayerPositionAndLookServerPkt = internal.PlayerPositionAndLookServer_404_0

// ID=0x7
type PlayerDiggingPkt struct {
	/* The action the player is taking against the block (see below) */
	Status Byte // Byte
	/* Block position */
	Location Position // Position
	/* The face being hit (see below) */
	Face Byte // Byte
}

var _ Packet = (*PlayerDiggingPkt)(nil)

func (p PlayerDiggingPkt)Encode(b *PacketBuilder){
	b.Byte(p.Status)
	p.Location.Encode(b)
	b.Byte(p.Face)
}

func (p *PlayerDiggingPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.Status, ok = r.Byte(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return err
	}
	if p.Face, ok = r.Byte(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x8
type PlayerBlockPlacementPkt struct {
	/* Block position */
	Location Position // Position
	/* The face on which the block is placed (see above) */
	Face Byte // Byte
	HeldItem *data.Slot // Slot
	/* The position of the crosshair on the block */
	CursorPositionX Byte // Byte
	CursorPositionY Byte // Byte
	CursorPositionZ Byte // Byte
}

var _ Packet = (*PlayerBlockPlacementPkt)(nil)

func (p PlayerBlockPlacementPkt)Encode(b *PacketBuilder){
	p.Location.Encode(b)
	b.Byte(p.Face)
	p.HeldItem.Encode(b)
	b.Byte(p.CursorPositionX)
	b.Byte(p.CursorPositionY)
	b.Byte(p.CursorPositionZ)
}

func (p *PlayerBlockPlacementPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if err = p.Location.DecodeFrom(r); err != nil {
		return err
	}
	if p.Face, ok = r.Byte(); !ok {
		return io.EOF
	}
	p.HeldItem = new(data.Slot)
	if err = p.HeldItem.DecodeFrom(r); err != nil {
		return err
	}
	if p.CursorPositionX, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.CursorPositionY, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.CursorPositionZ, ok = r.Byte(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x9
type PlayHeldItemChangeServerPkt = internal.PlayHeldItemChangeServer_758_0

// ID=0xa
type PlayAnimationServerPkt struct {
}

var _ Packet = (*PlayAnimationServerPkt)(nil)

func (p PlayAnimationServerPkt)Encode(b *PacketBuilder){
}

func (p *PlayAnimationServerPkt)DecodeFrom(r *PacketReader)(error){ return nil }

// ID=0xb
type PlayEntityActionPkt struct {
	/* Player ID */
	EntityID VarInt // VarInt
	/* The ID of the action, see below */
	ActionID VarInt // VarInt
	/* Only used by Horse Jump Boost, in which case it ranges from 0 to 100.  In all other cases it is 0. */
	ActionParameter VarInt // VarInt
}

var _ Packet = (*PlayEntityActionPkt)(nil)

func (p PlayEntityActionPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.ActionID)
	b.VarInt(p.ActionParameter)
}

func (p *PlayEntityActionPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ActionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ActionParameter, ok = r.VarInt(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0xc
type PlaySteerVehiclePkt = internal.PlaySteerVehicle_758_0

// ID=0xd
type PlayCloseWindowServerPkt = internal.PlayCloseWindowServer_758_0

// ID=0xe
type PlayClickWindowPkt struct {
	/* The ID of the window which was clicked. 0 for player inventory. */
	WindowID UByte // Unsigned Byte
	/* The clicked slot number, see below */
	Slot Short // Short
	/* The button used in the click, see below */
	Button Byte // Byte
	/* A unique number for the action, implemented by Notchian as a counter, starting at 1. Used by the server to send back a Confirm Transaction. */
	ActionNumber Short // Short
	/* Inventory operation mode, see below */
	Mode Byte // Byte Enum
	/* The clicked slot. Has to be empty (item ID = -1) for drop mode. */
	ClickedItem *data.Slot // Slot
}

var _ Packet = (*PlayClickWindowPkt)(nil)

func (p PlayClickWindowPkt)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.Short(p.Slot)
	b.Byte(p.Button)
	b.Short(p.ActionNumber)
	b.Byte(p.Mode)
	p.ClickedItem.Encode(b)
}

func (p *PlayClickWindowPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.Button, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ActionNumber, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.Byte(); !ok {
		return io.EOF
	}
	p.ClickedItem = new(data.Slot)
	if err = p.ClickedItem.DecodeFrom(r); err != nil {
		return err
	}
	return nil
}

// ID=0xf
type PlayConfirmTransactionServerPkt = internal.PlayConfirmTransactionServer_404_0

// ID=0x10
type PlayCreativeInventoryActionPkt = internal.PlayCreativeInventoryAction_758_0

// ID=0x11
type PlayEnchantItemPkt = internal.PlayEnchantItem_404_0

// ID=0x12
type PlayUpdateSignServerPkt struct {
	/* Block Coordinates */
	Location Position // Position
	/* First line of text in the sign */
	Line1 Object // Chat
	/* Second line of text in the sign */
	Line2 Object // Chat
	/* Third line of text in the sign */
	Line3 Object // Chat
	/* Fourth line of text in the sign */
	Line4 Object // Chat
}

var _ Packet = (*PlayUpdateSignServerPkt)(nil)

func (p PlayUpdateSignServerPkt)Encode(b *PacketBuilder){
	p.Location.Encode(b)
	b.JSON(p.Line1)
	b.JSON(p.Line2)
	b.JSON(p.Line3)
	b.JSON(p.Line4)
}

func (p *PlayUpdateSignServerPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if err = p.Location.DecodeFrom(r); err != nil {
		return err
	}
	if err = r.JSON(&p.Line1); err != nil {
		return err
	}
	if err = r.JSON(&p.Line2); err != nil {
		return err
	}
	if err = r.JSON(&p.Line3); err != nil {
		return err
	}
	if err = r.JSON(&p.Line4); err != nil {
		return err
	}
	return nil
}

// ID=0x13
type PlayerAbilitiesServerPkt = internal.PlayerAbilitiesServer_578_1

// ID=0x14
type PlayTabCompleteServerPkt struct {
	/* All text behind the cursor */
	Text String // String
	HasPosition Bool // Boolean
	/* The position of the block being looked at. Only sent if Has Position is true. */
	LookedAtBlock Optional[Position] // Optional Position
}

var _ Packet = (*PlayTabCompleteServerPkt)(nil)

func (p PlayTabCompleteServerPkt)Encode(b *PacketBuilder){
	b.String(p.Text)
	b.Bool(p.HasPosition)
	if p.LookedAtBlock.Ok = TODO; p.LookedAtBlock.Ok {
		p.LookedAtBlock.V.Encode(b)
	}
}

func (p *PlayTabCompleteServerPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.Text, ok = r.String(); !ok {
		return io.EOF
	}
	if p.HasPosition, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.LookedAtBlock.Ok = TODO; p.LookedAtBlock.Ok {
		if err = p.LookedAtBlock.V.DecodeFrom(r); err != nil {
			return err
		}
	}
	return nil
}

// ID=0x15
type PlayClientSettingsPkt struct {
	/* e.g. en_GB */
	Locale String // String
	/* Client-side render distance, in chunks */
	ViewDistance Byte // Byte
	/* 0: enabled, 1: commands only, 2: hidden */
	ChatMode Byte // Byte
	/* “Colors” multiplayer setting */
	ChatColors Bool // Boolean
	/* Skin parts, see note below */
	DisplayedSkinParts UByte // Unsigned Byte
}

var _ Packet = (*PlayClientSettingsPkt)(nil)

func (p PlayClientSettingsPkt)Encode(b *PacketBuilder){
	b.String(p.Locale)
	b.Byte(p.ViewDistance)
	b.Byte(p.ChatMode)
	b.Bool(p.ChatColors)
	b.UByte(p.DisplayedSkinParts)
}

func (p *PlayClientSettingsPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.Locale, ok = r.String(); !ok {
		return io.EOF
	}
	if p.ViewDistance, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ChatMode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ChatColors, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.DisplayedSkinParts, ok = r.UByte(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x16
type PlayClientStatusPkt = internal.PlayClientStatus_758_0

// ID=0x17
type PlayPluginMessageServerPkt = internal.PlayPluginMessageServer_763_0

// ID=0x18
type PlaySpectatePkt = internal.PlaySpectate_758_0

// ID=0x19
type PlayResourcePackStatusPkt = internal.PlayResourcePackStatus_110_1

// ---- play: clientbound ----

// ID=0x0
type PlayKeepAliveClientPkt = internal.PlayKeepAlive_338_1

// ID=0x1
type PlayJoinGamePkt struct {
	/* The player's Entity ID (EID) */
	EntityID Int // Int
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. Bit 3 (0x8) is the hardcore flag. */
	Gamemode UByte // Unsigned Byte
	/* -1: Nether, 0: Overworld, 1: End */
	Dimension Byte // Byte
	/* 0: peaceful, 1: easy, 2: normal, 3: hard */
	Difficulty UByte // Unsigned Byte
	/* Used by the client to draw the player list */
	MaxPlayers UByte // Unsigned Byte
	/* default, flat, largeBiomes, amplified, default_1_1 */
	LevelType String // String
	/* If true, a Notchian client shows reduced information on the debug screen. */
	ReducedDebugInfo Bool // Boolean
}

var _ Packet = (*PlayJoinGamePkt)(nil)

func (p PlayJoinGamePkt)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.UByte(p.Gamemode)
	b.Byte(p.Dimension)
	b.UByte(p.Difficulty)
	b.UByte(p.MaxPlayers)
	b.String(p.LevelType)
	b.Bool(p.ReducedDebugInfo)
}

func (p *PlayJoinGamePkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Dimension, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Difficulty, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.MaxPlayers, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.LevelType, ok = r.String(); !ok {
		return io.EOF
	}
	if p.ReducedDebugInfo, ok = r.Bool(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x2
type PlayChatMessageClientPkt = internal.PlayChatMessage_578_4

// ID=0x3
type PlayTimeUpdatePkt = internal.PlayTimeUpdate_758_0

// ID=0x4
type PlayEntityEquipmentPkt struct {
	/* Entity's EID */
	EntityID VarInt // VarInt
	/* Equipment slot. 0: held, 1–4: armor slot (1: boots, 2: leggings, 3: chestplate, 4: helmet) */
	Slot Short // Short
	/* Item in slot format */
	Item *data.Slot // Slot
}

var _ Packet = (*PlayEntityEquipmentPkt)(nil)

func (p PlayEntityEquipmentPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.Slot)
	p.Item.Encode(b)
}

func (p *PlayEntityEquipmentPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
	p.Item = new(data.Slot)
	if err = p.Item.DecodeFrom(r); err != nil {
		return err
	}
	return nil
}

// ID=0x5
type PlaySpawnPositionPkt = internal.PlaySpawnPosition_754_1

// ID=0x6
type PlayUpdateHealthPkt = internal.PlayUpdateHealth_758_0

// ID=0x7
type PlayRespawnPkt = internal.PlayRespawn_404_8

// ID=0x8
type PlayerPositionAndLookClientPkt struct {
	/* Absolute or relative position, depending on Flags */
	X Double // Double
	/* Absolute or relative position, depending on Flags */
	Y Double // Double
	/* Absolute or relative position, depending on Flags */
	Z Double // Double
	/* Absolute or relative rotation on the X Axis, in degrees */
	Yaw Float // Float
	/* Absolute or relative rotation on the Y Axis, in degrees */
	Pitch Float // Float
	/* Bit field, see below */
	Flags Byte // Byte
}

var _ Packet = (*PlayerPositionAndLookClientPkt)(nil)

func (p PlayerPositionAndLookClientPkt)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Byte(p.Flags)
}

func (p *PlayerPositionAndLookClientPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Yaw, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x9
type PlayHeldItemChangeClientPkt = internal.PlayHeldItemChange_758_0

// ID=0xa
type PlayUseBedPkt = internal.PlayUseBed_404_0

// ID=0xb
type PlayAnimationClientPkt = internal.PlayAnimation_404_1

// ID=0xc
type PlaySpawnPlayerPkt struct {
	/* Player's EID */
	EntityID VarInt // VarInt
	PlayerUUID UUID // UUID
	/* Player X as a Fixed-Point number */
	X Int // Int
	/* Player Y as a Fixed-Point number */
	Y Int // Int
	/* Player Z as a Fixed-Point number */
	Z Int // Int
	Yaw Angle // Angle
	Pitch Angle // Angle
	/* The item the player is currently holding. Note that this should be 0 for “no item”, unlike -1 used in other packets. */
	CurrentItem Short // Short
	Metadata *data.EntityMetadata // Metadata
}

var _ Packet = (*PlaySpawnPlayerPkt)(nil)

func (p PlaySpawnPlayerPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.PlayerUUID)
	b.Int(p.X)
	b.Int(p.Y)
	b.Int(p.Z)
	p.Yaw.Encode(b)
	p.Pitch.Encode(b)
	b.Short(p.CurrentItem)
	p.Metadata.Encode(b)
}

func (p *PlaySpawnPlayerPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.PlayerUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return err
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return err
	}
	if p.CurrentItem, ok = r.Short(); !ok {
		return io.EOF
	}
	p.Metadata = new(data.EntityMetadata)
	if err = p.Metadata.DecodeFrom(r); err != nil {
		return err
	}
	return nil
}

// ID=0xd
type PlayCollectItemPkt = internal.PlayCollectItem_210_1

// ID=0xe
type PlaySpawnObjectPkt struct {
	/* EID of the object */
	EntityID VarInt // VarInt
	/* The type of object (see Entities#Objects) */
	Type Byte // Byte
	/* X position as a Fixed-Point number */
	X Int // Int
	/* Y position as a Fixed-Point number */
	Y Int // Int
	/* Z position as a Fixed-Point number */
	Z Int // Int
	Pitch Angle // Angle
	Yaw Angle // Angle
	/* Meaning dependent on the value of the Type field, see Object Data for details. */
	Data Int // Int
	/* Only sent if the Data field is nonzero. Same units as Entity Velocity. */
	VelocityX Optional[Short] // Optional Short
	/* Only sent if the Data field is nonzero. Same units as Entity Velocity. */
	VelocityY Optional[Short] // Optional Short
	/* Only sent if the Data field is nonzero. Same units as Entity Velocity. */
	VelocityZ Optional[Short] // Optional Short
}

var _ Packet = (*PlaySpawnObjectPkt)(nil)

func (p PlaySpawnObjectPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Byte(p.Type)
	b.Int(p.X)
	b.Int(p.Y)
	b.Int(p.Z)
	p.Pitch.Encode(b)
	p.Yaw.Encode(b)
	b.Int(p.Data)
	if p.VelocityX.Ok = TODO; p.VelocityX.Ok {
		b.Short(p.VelocityX.V)
	}
	if p.VelocityY.Ok = TODO; p.VelocityY.Ok {
		b.Short(p.VelocityY.V)
	}
	if p.VelocityZ.Ok = TODO; p.VelocityZ.Ok {
		b.Short(p.VelocityZ.V)
	}
}

func (p *PlaySpawnObjectPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return err
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return err
	}
	if p.Data, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.VelocityX.Ok = TODO; p.VelocityX.Ok {
		if p.VelocityX.V, ok = r.Short(); !ok {
			return io.EOF
		}
	}
	if p.VelocityY.Ok = TODO; p.VelocityY.Ok {
		if p.VelocityY.V, ok = r.Short(); !ok {
			return io.EOF
		}
	}
	if p.VelocityZ.Ok = TODO; p.VelocityZ.Ok {
		if p.VelocityZ.V, ok = r.Short(); !ok {
			return io.EOF
		}
	}
	return nil
}

// ID=0xf
type PlaySpawnMobPkt struct {
	EntityID VarInt // VarInt
	/* The type of mob. See Mobs */
	Type UByte // Unsigned Byte
	/* X position as a Fixed-Point number */
	X Int // Int
	/* Y position as a Fixed-Point number */
	Y Int // Int
	/* Z position as a Fixed-Point number */
	Z Int // Int
	Yaw Angle // Angle
	Pitch Angle // Angle
	HeadPitch Angle // Angle
	/* Same units as Entity Velocity */
	VelocityX Short // Short
	/* Same units as Entity Velocity */
	VelocityY Short // Short
	/* Same units as Entity Velocity */
	VelocityZ Short // Short
	Metadata *data.EntityMetadata // Metadata
}

var _ Packet = (*PlaySpawnMobPkt)(nil)

func (p PlaySpawnMobPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UByte(p.Type)
	b.Int(p.X)
	b.Int(p.Y)
	b.Int(p.Z)
	p.Yaw.Encode(b)
	p.Pitch.Encode(b)
	p.HeadPitch.Encode(b)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
	p.Metadata.Encode(b)
}

func (p *PlaySpawnMobPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return err
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return err
	}
	if err = p.HeadPitch.DecodeFrom(r); err != nil {
		return err
	}
	if p.VelocityX, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.VelocityY, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.VelocityZ, ok = r.Short(); !ok {
		return io.EOF
	}
	p.Metadata = new(data.EntityMetadata)
	if err = p.Metadata.DecodeFrom(r); err != nil {
		return err
	}
	return nil
}

// ID=0x10
type PlaySpawnPaintingPkt struct {
	EntityID VarInt // VarInt
	/* Name of the painting. Max length 13 */
	Title String // String
	/* Center coordinates */
	Location Position // Position
	/* Direction the painting faces. 0: north (-z), 1: west (-x), 2: south (+z), 3: east (+x) */
	Direction UByte // Unsigned Byte
}

var _ Packet = (*PlaySpawnPaintingPkt)(nil)

func (p PlaySpawnPaintingPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.String(p.Title)
	p.Location.Encode(b)
	b.UByte(p.Direction)
}

func (p *PlaySpawnPaintingPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Title, ok = r.String(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return err
	}
	if p.Direction, ok = r.UByte(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x11
type PlaySpawnExperienceOrbPkt struct {
	/* Entity's ID */
	EntityID VarInt // VarInt
	/* X position as a Fixed-Point number */
	X Int // Int
	/* Y position as a Fixed-Point number */
	Y Int // Int
	/* Z position as a Fixed-Point number */
	Z Int // Int
	/* The amount of experience this orb will reward once collected */
	Count Short // Short
}

var _ Packet = (*PlaySpawnExperienceOrbPkt)(nil)

func (p PlaySpawnExperienceOrbPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Int(p.X)
	b.Int(p.Y)
	b.Int(p.Z)
	b.Short(p.Count)
}

func (p *PlaySpawnExperienceOrbPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Count, ok = r.Short(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x12
type PlayEntityVelocityPkt = internal.PlayEntityVelocity_758_0

// ID=0x13
type PlayDestroyEntitiesPkt = internal.PlayDestroyEntities_758_0

// ID=0x14
type PlayEntityPkt = internal.PlayEntity_404_0

// ID=0x15
type PlayEntityRelativeMovePkt struct {
	EntityID VarInt // VarInt
	/* Change in X position as a Fixed-Point number */
	DeltaX Byte // Byte
	/* Change in Y position as a Fixed-Point number */
	DeltaY Byte // Byte
	/* Change in Z position as a Fixed-Point number */
	DeltaZ Byte // Byte
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityRelativeMovePkt)(nil)

func (p PlayEntityRelativeMovePkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Byte(p.DeltaX)
	b.Byte(p.DeltaY)
	b.Byte(p.DeltaZ)
	b.Bool(p.OnGround)
}

func (p *PlayEntityRelativeMovePkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DeltaX, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.DeltaY, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.DeltaZ, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x16
type PlayEntityLookPkt = internal.PlayEntityLook_404_0

// ID=0x17
type PlayEntityLookAndRelativeMovePkt struct {
	EntityID VarInt // VarInt
	/* Change in X position as a Fixed-Point number */
	DeltaX Byte // Byte
	/* Change in Y position as a Fixed-Point number */
	DeltaY Byte // Byte
	/* Change in Z position as a Fixed-Point number */
	DeltaZ Byte // Byte
	/* New angle, not a delta */
	Yaw Angle // Angle
	/* New angle, not a delta */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityLookAndRelativeMovePkt)(nil)

func (p PlayEntityLookAndRelativeMovePkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Byte(p.DeltaX)
	b.Byte(p.DeltaY)
	b.Byte(p.DeltaZ)
	p.Yaw.Encode(b)
	p.Pitch.Encode(b)
	b.Bool(p.OnGround)
}

func (p *PlayEntityLookAndRelativeMovePkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DeltaX, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.DeltaY, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.DeltaZ, ok = r.Byte(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return err
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return err
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x18
type PlayEntityTeleportPkt struct {
	EntityID VarInt // VarInt
	/* X position as a Fixed-Point number */
	X Int // Int
	/* Y position as a Fixed-Point number */
	Y Int // Int
	/* Z position as a Fixed-Point number */
	Z Int // Int
	/* New angle, not a delta */
	Yaw Angle // Angle
	/* New angle, not a delta */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityTeleportPkt)(nil)

func (p PlayEntityTeleportPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Int(p.X)
	b.Int(p.Y)
	b.Int(p.Z)
	p.Yaw.Encode(b)
	p.Pitch.Encode(b)
	b.Bool(p.OnGround)
}

func (p *PlayEntityTeleportPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return err
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return err
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x19
type PlayEntityHeadLookPkt = internal.PlayEntityHeadLook_758_0

// ID=0x1a
type PlayEntityStatusPkt = internal.PlayEntityStatus_758_0

// ID=0x1b
type PlayAttachEntityPkt struct {
	/* Attached entity's EID */
	EntityID Int // Int
	/* Vechicle's Entity ID. Set to -1 to detach */
	VehicleID Int // Int
	/* If true leashes the entity to the vehicle */
	Leash Bool // Boolean
}

var _ Packet = (*PlayAttachEntityPkt)(nil)

func (p PlayAttachEntityPkt)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.Int(p.VehicleID)
	b.Bool(p.Leash)
}

func (p *PlayAttachEntityPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.VehicleID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Leash, ok = r.Bool(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x1c
type PlayEntityMetadataPkt = internal.PlayEntityMetadata_758_0

// ID=0x1d
type PlayEntityEffectPkt = internal.PlayEntityEffect_110_3

// ID=0x1e
type PlayRemoveEntityEffectPkt = internal.PlayRemoveEntityEffect_757_1

// ID=0x1f
type PlaySetExperiencePkt = internal.PlaySetExperience_760_1

// ID=0x20
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name           | Field Name           | Field Type | Field Type             | Notes                                     |
	 * |-----------|-------|----------|----------------------|----------------------|------------|------------------------|-------------------------------------------|
	 * | 0x20      | Play  | Client   | Entity ID            | Entity ID            | VarInt     | VarInt                 |                                           |
	 * | 0x20      | Play  | Client   | Number Of Properties | Number Of Properties | Int        | Int                    | Number of elements in the following array |
	 * | 0x20      | Play  | Client   | Property             | Key                  | Array      | String                 | See below                                 |
	 * | 0x20      | Play  | Client   | Property             | Value                | Array      | Double                 | See below                                 |
	 * | 0x20      | Play  | Client   | Property             | Number Of Modifiers  | Array      | VarInt                 | Number of elements in the following array |
	 * | 0x20      | Play  | Client   | Property             | Modifiers            | Array      | Array of Modifier Data | See Attribute#Modifiers                   |
	 * 
	 */
}

// ID=0x21
type PlayChunkDataPkt struct {
	/* Chunk X coordinate */
	ChunkX Int // Int
	/* Chunk Z coordinate */
	ChunkZ Int // Int
	/* This is true if the packet represents all sections in this vertical column, where the Primary Bit Mask specifies exactly which sections are included, and which are air */
	GroundUpContinuous Bool // Boolean
	/* Bitmask with 1 for every 16x16x16 section whose data follows in the compressed data */
	PrimaryBitMask UShort // Unsigned Short
	/* Size of Data */
	Size VarInt // VarInt
	Data *ChunkSection // Chunk
}

var _ Packet = (*PlayChunkDataPkt)(nil)

func (p PlayChunkDataPkt)Encode(b *PacketBuilder){
	b.Int(p.ChunkX)
	b.Int(p.ChunkZ)
	b.Bool(p.GroundUpContinuous)
	p.PrimaryBitMask.Encode(b)
	b.VarInt(p.Size)
	p.Data.Encode(b)
}

func (p *PlayChunkDataPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.ChunkX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.GroundUpContinuous, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = p.PrimaryBitMask.DecodeFrom(r); err != nil {
		return err
	}
	if p.Size, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Data = new(ChunkSection)
	if err = p.Data.DecodeFrom(r); err != nil {
		return err
	}
	return nil
}

// ID=0x22
type PlayMultiBlockChangePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name   | Field Name          | Field Type | Field Type    | Notes                                                                                                                                                                    |
	 * |-----------|-------|----------|--------------|---------------------|------------|---------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x22      | Play  | Client   | Chunk X      | Chunk X             | Int        | Int           | Chunk X coordinate                                                                                                                                                       |
	 * | 0x22      | Play  | Client   | Chunk Z      | Chunk Z             | Int        | Int           | Chunk Z coordinate                                                                                                                                                       |
	 * | 0x22      | Play  | Client   | Record Count | Record Count        | VarInt     | VarInt        | Number of elements in the following array, i.e. the number of blocks affected                                                                                            |
	 * | 0x22      | Play  | Client   | Record       | Horizontal Position | Array      | Unsigned Byte | The 4 most significant bits (0xF0) encode the X coordinate, relative to the chunk. The 4 least significant bits (0x0F) encode the Z coordinate, relative to the chunk.   |
	 * | 0x22      | Play  | Client   | Record       | Y Coordinate        | Array      | Unsigned Byte |                                                                                                                                                                          |
	 * | 0x22      | Play  | Client   | Record       | Block ID            | Array      | VarInt        | The new block state ID for the block as given in the global palette (When reading data: type = id >> 4, meta = id & 15, when writing data: id = type << 4 | (meta & 15)) |
	 * 
	 */
}

// ID=0x23
type PlayBlockChangePkt = internal.PlayBlockChange_758_0

// ID=0x24
type PlayBlockActionPkt = internal.PlayBlockAction_110_2

// ID=0x25
type PlayBlockBreakAnimationPkt = internal.PlayBlockBreakAnimation_758_0

// ID=0x26
type PlayMapChunkBulkPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name         | Field Name         | Field Type     | Field Type     | Notes                                                                                                            |
	 * |-----------|-------|----------|--------------------|--------------------|----------------|----------------|------------------------------------------------------------------------------------------------------------------|
	 * | 0x26      | Play  | Client   | Sky Light Sent     | Sky Light Sent     | Boolean        | Boolean        | Whether or not Chunk Data contains light nibble arrays. This is true in the Overworld, false in the End + Nether |
	 * | 0x26      | Play  | Client   | Chunk Column Count | Chunk Column Count | VarInt         | VarInt         | Number of elements in each of the following arrays                                                               |
	 * | 0x26      | Play  | Client   | Chunk Meta         | Chunk X            | Array          | Int            | The X coordinate of the chunk                                                                                    |
	 * | 0x26      | Play  | Client   | Chunk Meta         | Chunk Z            | Array          | Int            | The Z coordinate of the chunk                                                                                    |
	 * | 0x26      | Play  | Client   | Chunk Meta         | Primary Bit Mask   | Array          | Unsigned Short | A bit mask which specifies which sections are not empty in this chunk                                            |
	 * | 0x26      | Play  | Client   | Chunk Data         | Chunk Data         | Array of Chunk | Array of Chunk | Each chunk in this array corresponds to the data at the same position in Chunk Meta                              |
	 * 
	 */
}

// ID=0x27
type PlayExplosionPkt = internal.PlayExplosion_498_3

// ID=0x28
type PlayEffectPkt = internal.PlayEffect_758_0

// ID=0x29
type PlaySoundEffectPkt struct {
	/* All known sound effect names can be seen here */
	SoundName String // String
	/* Effect X multiplied by 8 */
	EffectPositionX Int // Int
	/* Effect Y multiplied by 8 */
	EffectPositionY Int // Int
	/* Effect Z multiplied by 8 */
	EffectPositionZ Int // Int
	/* 1 is 100%, can be more */
	Volume Float // Float
	/* 63 is 100%, can be more */
	Pitch UByte // Unsigned Byte
}

var _ Packet = (*PlaySoundEffectPkt)(nil)

func (p PlaySoundEffectPkt)Encode(b *PacketBuilder){
	b.String(p.SoundName)
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
	if p.SoundName, ok = r.String(); !ok {
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

// ID=0x2a
type PlayParticlePkt = internal.PlayParticle_340_3

// ID=0x2b
type PlayChangeGameStatePkt = internal.PlayChangeGameState_758_0

// ID=0x2c
type PlaySpawnGlobalEntityPkt struct {
	/* The EID of the thunderbolt */
	EntityID VarInt // VarInt
	/* The global entity type, currently always 1 for thunderbolt */
	Type Byte // Byte
	/* Thunderbolt X, a fixed-point number */
	X Int // Int
	/* Thunderbolt Y, a fixed-point number */
	Y Int // Int
	/* Thunderbolt Z, a fixed-point number */
	Z Int // Int
}

var _ Packet = (*PlaySpawnGlobalEntityPkt)(nil)

func (p PlaySpawnGlobalEntityPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Byte(p.Type)
	b.Int(p.X)
	b.Int(p.Y)
	b.Int(p.Z)
}

func (p *PlaySpawnGlobalEntityPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Int(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x2d
type PlayOpenWindowPkt = internal.PlayOpenWindow_404_1

// ID=0x2e
type PlayCloseWindowClientPkt = internal.PlayCloseWindow_758_0

// ID=0x2f
type PlaySetSlotPkt = internal.PlaySetSlot_755_1

// ID=0x30
type PlayWindowItemsPkt = internal.PlayWindowItems_755_1

// ID=0x31
type PlayWindowPropertyPkt = internal.PlayWindowProperty_758_0

// ID=0x32
type PlayConfirmTransactionClientPkt = internal.PlayConfirmTransaction_404_0

// ID=0x33
type PlayUpdateSignClientPkt struct {
	Location Position // Position
	/* First line of text in the sign */
	Line1 Object // Chat
	/* Second line of text in the sign */
	Line2 Object // Chat
	/* Third line of text in the sign */
	Line3 Object // Chat
	/* Fourth line of text in the sign */
	Line4 Object // Chat
}

var _ Packet = (*PlayUpdateSignClientPkt)(nil)

func (p PlayUpdateSignClientPkt)Encode(b *PacketBuilder){
	p.Location.Encode(b)
	b.JSON(p.Line1)
	b.JSON(p.Line2)
	b.JSON(p.Line3)
	b.JSON(p.Line4)
}

func (p *PlayUpdateSignClientPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if err = p.Location.DecodeFrom(r); err != nil {
		return err
	}
	if err = r.JSON(&p.Line1); err != nil {
		return err
	}
	if err = r.JSON(&p.Line2); err != nil {
		return err
	}
	if err = r.JSON(&p.Line3); err != nil {
		return err
	}
	if err = r.JSON(&p.Line4); err != nil {
		return err
	}
	return nil
}

// ID=0x34
type PlayMapPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name  | Field Name         | Field Type                      | Field Type                      | Notes                                                              |
	 * |-----------|-------|----------|-------------|--------------------|---------------------------------|---------------------------------|--------------------------------------------------------------------|
	 * | 0x34      | Play  | Client   | Item Damage | Item Damage        | VarInt                          | VarInt                          | The damage value (map ID) of the map being modified                |
	 * | 0x34      | Play  | Client   | Scale       | Scale              | Byte                            | Byte                            |                                                                    |
	 * | 0x34      | Play  | Client   | Icon Count  | Icon Count         | VarInt                          | VarInt                          | Number of elements in the following array                          |
	 * | 0x34      | Play  | Client   | Icon        | Direction And Type | Array                           | Byte                            | 0xF0 = Direction, 0x0F = Type                                      |
	 * | 0x34      | Play  | Client   | Icon        | X                  | Array                           | Byte                            |                                                                    |
	 * | 0x34      | Play  | Client   | Icon        | Z                  | Array                           | Byte                            |                                                                    |
	 * | 0x34      | Play  | Client   | Columns     | Columns            | Byte                            | Byte                            | Number of columns updated                                          |
	 * | 0x34      | Play  | Client   | Rows        | Rows               | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; number of rows updated             |
	 * | 0x34      | Play  | Client   | X           | X                  | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; x offset of the westernmost column |
	 * | 0x34      | Play  | Client   | Z           | Z                  | Optional Byte                   | Optional Byte                   | Only if Columns is more than 0; z offset of the northernmost row   |
	 * | 0x34      | Play  | Client   | Length      | Length             | Optional VarInt                 | Optional VarInt                 | Only if Columns is more than 0; length of the following array      |
	 * | 0x34      | Play  | Client   | Data        | Data               | Optional Array of Unsigned Byte | Optional Array of Unsigned Byte | Only if Columns is more than 0; see Map item format                |
	 * 
	 */
}

// ID=0x35
type PlayUpdateBlockEntityPkt struct {
	Location Position // Position
	/* The type of update to perform, see below */
	Action UByte // Unsigned Byte
	/* If not present then it's a TAG_END (0) */
	NBTData Optional[nbt.NBT] // Optional NBT Tag
}

var _ Packet = (*PlayUpdateBlockEntityPkt)(nil)

func (p PlayUpdateBlockEntityPkt)Encode(b *PacketBuilder){
	p.Location.Encode(b)
	b.UByte(p.Action)
	if p.NBTData.Ok = TODO; p.NBTData.Ok {
		WriteNBT(b, p.NBTData.V)
	}
}

func (p *PlayUpdateBlockEntityPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if err = p.Location.DecodeFrom(r); err != nil {
		return err
	}
	if p.Action, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.NBTData.Ok = TODO; p.NBTData.Ok {
		if p.NBTData.V, err = nbt.ReadNBT(r); err != nil {
			return err
		}
	}
	return nil
}

// ID=0x36
type PlayOpenSignEditorPkt = internal.PlayOpenSignEditor_762_1

// ID=0x37
type PlayStatisticsPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name | Field Name | Field Type | Field Type | Notes                                                     |
	 * |-----------|-------|----------|------------|------------|------------|------------|-----------------------------------------------------------|
	 * | 0x37      | Play  | Client   | Count      | Count      | VarInt     | VarInt     | Number of elements in the following array                 |
	 * | 0x37      | Play  | Client   | Statistic  | Name       | Array      | String     | https://gist.github.com/thinkofdeath/a1842c21a0cf2e1fb5e0 |
	 * | 0x37      | Play  | Client   | Statistic  | Value      | Array      | VarInt     | The amount to set it to                                   |
	 * 
	 */
}

// ID=0x38
type PlayerListItemPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name        | Field Name             | Field Name           | Field Name           | Field Type | Field Type    | Field Type      | Notes                                                   |
	 * |-----------|-------|----------|-------------------|------------------------|----------------------|----------------------|------------|---------------|-----------------|---------------------------------------------------------|
	 * | 0x38      | Play  | Client   | Action            | Action                 | Action               | Action               | VarInt     | VarInt        | VarInt          | Determines the rest of the Player format after the UUID |
	 * | 0x38      | Play  | Client   | Number Of Players | Number Of Players      | Number Of Players    | Number Of Players    | VarInt     | VarInt        | VarInt          | Number of elements in the following array               |
	 * | 0x38      | Play  | Client   | Player            | UUID                   | UUID                 | UUID                 | Array      | UUID          | UUID            |                                                         |
	 * | 0x38      | Play  | Client   | Player            | Action                 | Field Name           | Field Name           | Array      |               |                 |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Name                 | Name                 | Array      | String        | String          |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Number Of Properties | Number Of Properties | Array      | VarInt        | VarInt          | Number of elements in the following array               |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Property             | Name                 | Array      | Array         | String          |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Property             | Value                | Array      | Array         | String          |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Property             | Is Signed            | Array      | Array         | Boolean         |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Property             | Signature            | Array      | Array         | Optional String | Only if Is Signed is true                               |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Gamemode             | Gamemode             | Array      | VarInt        | VarInt          |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Ping                 | Ping                 | Array      | VarInt        | VarInt          |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean         |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 0: add player          | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat   | Only if Has Display Name is true                        |
	 * | 0x38      | Play  | Client   | Player            | 1: update gamemode     | Gamemode             | Gamemode             | Array      | VarInt        | VarInt          |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 2: update latency      | Ping                 | Ping                 | Array      | VarInt        | VarInt          |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 3: update display name | Has Display Name     | Has Display Name     | Array      | Boolean       | Boolean         |                                                         |
	 * | 0x38      | Play  | Client   | Player            | 3: update display name | Display Name         | Display Name         | Array      | Optional Chat | Optional Chat   | Only send if Has Display Name is true                   |
	 * | 0x38      | Play  | Client   | Player            | 4: remove player       | no fields            | no fields            | Array      | no fields     | no fields       |                                                         |
	 * 
	 */
}

// ID=0x39
type PlayerAbilitiesClientPkt = internal.PlayerAbilities_763_0

// ID=0x3a
type PlayTabCompleteClientPkt = internal.PlayTabComplete_340_5

// ID=0x3b
type PlayScoreboardObjectivePkt = internal.PlayScoreboardObjective_340_1

// ID=0x3c
type PlayUpdateScorePkt = internal.PlayUpdateScore_315_2

// ID=0x3d
type PlayDisplayScoreboardPkt = internal.PlayDisplayScoreboard_758_0

// ID=0x3e
type PlayTeamsPkt struct {
	/* A unique name for the team. (Shared with scoreboard). */
	TeamName String // String
	/* If 0 then the team is created.
	 * If 1 then the team is removed. 
	 * If 2 the team team information is updated. 
	 * If 3 then new players are added to the team. 
	 * If 4 then players are removed from the team. */
	Mode Byte // Byte
	/* Only if Mode = 0 or 2. */
	TeamDisplayName Optional[String] // Optional String
	/* Only if Mode = 0 or 2. Displayed before the players' name that are part of this team */
	TeamPrefix Optional[String] // Optional String
	/* Only if Mode = 0 or 2. Displayed after the players' name that are part of this team */
	TeamSuffix Optional[String] // Optional String
	/* Only if Mode = 0 or 2. 0 for off, 1 for on, 3 for seeing friendly invisibles */
	FriendlyFire Optional[Byte] // Optional Byte
	/* Only if Mode = 0 or 2. always, hideForOtherTeams, hideForOwnTeam, never */
	NameTagVisibility Optional[String] // Optional String
	/* Only if Mode = 0 or 2. Same as Chat colors */
	Color Optional[Byte] // Optional Byte
	/* Only if Mode = 0 or 3 or 4. Number of players in the array */
	PlayerCount Optional[VarInt] // Optional VarInt
	/* Only if Mode = 0 or 3 or 4. Players to be added/remove from the team. Max 40 characters so may be uuid's later */
	Players Optional[[]String] // Optional Array of String
}

var _ Packet = (*PlayTeamsPkt)(nil)

func (p PlayTeamsPkt)Encode(b *PacketBuilder){
	b.String(p.TeamName)
	b.Byte(p.Mode)
	if p.TeamDisplayName.Ok = TODO; p.TeamDisplayName.Ok {
		b.String(p.TeamDisplayName.V)
	}
	if p.TeamPrefix.Ok = TODO; p.TeamPrefix.Ok {
		b.String(p.TeamPrefix.V)
	}
	if p.TeamSuffix.Ok = TODO; p.TeamSuffix.Ok {
		b.String(p.TeamSuffix.V)
	}
	if p.FriendlyFire.Ok = TODO; p.FriendlyFire.Ok {
		b.Byte(p.FriendlyFire.V)
	}
	if p.NameTagVisibility.Ok = TODO; p.NameTagVisibility.Ok {
		b.String(p.NameTagVisibility.V)
	}
	if p.Color.Ok = TODO; p.Color.Ok {
		b.Byte(p.Color.V)
	}
	if p.PlayerCount.Ok = TODO; p.PlayerCount.Ok {
		b.VarInt(p.PlayerCount.V)
	}
	if p.Players.Ok = TODO; p.Players.Ok {
		TODO_Encode_Array(p.Players.V)
	}
}

func (p *PlayTeamsPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.TeamName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.TeamDisplayName.Ok = TODO; p.TeamDisplayName.Ok {
		if p.TeamDisplayName.V, ok = r.String(); !ok {
			return io.EOF
		}
	}
	if p.TeamPrefix.Ok = TODO; p.TeamPrefix.Ok {
		if p.TeamPrefix.V, ok = r.String(); !ok {
			return io.EOF
		}
	}
	if p.TeamSuffix.Ok = TODO; p.TeamSuffix.Ok {
		if p.TeamSuffix.V, ok = r.String(); !ok {
			return io.EOF
		}
	}
	if p.FriendlyFire.Ok = TODO; p.FriendlyFire.Ok {
		if p.FriendlyFire.V, ok = r.Byte(); !ok {
			return io.EOF
		}
	}
	if p.NameTagVisibility.Ok = TODO; p.NameTagVisibility.Ok {
		if p.NameTagVisibility.V, ok = r.String(); !ok {
			return io.EOF
		}
	}
	if p.Color.Ok = TODO; p.Color.Ok {
		if p.Color.V, ok = r.Byte(); !ok {
			return io.EOF
		}
	}
	if p.PlayerCount.Ok = TODO; p.PlayerCount.Ok {
		if p.PlayerCount.V, ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
	if p.Players.Ok = TODO; p.Players.Ok {
		TODO_Decode_Array(p.Players.V)
	}
	return nil
}

// ID=0x3f
type PlayPluginMessageClientPkt = internal.PlayPluginMessage_763_0

// ID=0x40
type PlayDisconnectPkt = internal.PlayDisconnect_763_0

// ID=0x41
type PlayServerDifficultyPkt = internal.PlayServerDifficulty_404_1

// ID=0x42
type PlayCombatEventPkt struct {
	/* 0: enter combat, 1: end combat, 2: entity dead */
	Event VarInt // VarInt
	/* Only for end combat */
	Duration Optional[VarInt] // Optional VarInt
	/* Only for entity dead */
	PlayerID Optional[VarInt] // Optional VarInt
	/* Only for end combat and entity dead */
	EntityID Optional[Int] // Optional Int
	/* Only for entity dead */
	Message String // String
}

var _ Packet = (*PlayCombatEventPkt)(nil)

func (p PlayCombatEventPkt)Encode(b *PacketBuilder){
	b.VarInt(p.Event)
	if p.Duration.Ok = TODO; p.Duration.Ok {
		b.VarInt(p.Duration.V)
	}
	if p.PlayerID.Ok = TODO; p.PlayerID.Ok {
		b.VarInt(p.PlayerID.V)
	}
	if p.EntityID.Ok = TODO; p.EntityID.Ok {
		b.Int(p.EntityID.V)
	}
	b.String(p.Message)
}

func (p *PlayCombatEventPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.Event, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Duration.Ok = TODO; p.Duration.Ok {
		if p.Duration.V, ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
	if p.PlayerID.Ok = TODO; p.PlayerID.Ok {
		if p.PlayerID.V, ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
	if p.EntityID.Ok = TODO; p.EntityID.Ok {
		if p.EntityID.V, ok = r.Int(); !ok {
			return io.EOF
		}
	}
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x43
type PlayCameraPkt = internal.PlayCamera_758_0

// ID=0x44
type PlayWorldBorderPkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name            | Field Name               | Field Type | Notes                                                                                                                                                                                                       |
	 * |-----------|-------|----------|-----------------------|--------------------------|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x44      | Play  | Client   | Action                | Action                   | VarInt     | Determines the format of the rest of the packet                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | Action                | Field Name               |            |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 0: set size           | Radius                   | Double     | meters                                                                                                                                                                                                      |
	 * | 0x44      | Play  | Client   | 1: lerp size          | Old Radius               | Double     | meters                                                                                                                                                                                                      |
	 * | 0x44      | Play  | Client   | 1: lerp size          | New Radius               | Double     | meters                                                                                                                                                                                                      |
	 * | 0x44      | Play  | Client   | 1: lerp size          | Speed                    | VarLong    | number of real-time ticks/seconds (?) until New Radius is reached. From experiments, it appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag |
	 * | 0x44      | Play  | Client   | 2: set center         | X                        | Double     |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 2: set center         | Z                        | Double     |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 3: initialize         | X                        | Double     |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 3: initialize         | Z                        | Double     |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 3: initialize         | Old Radius               | Double     |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 3: initialize         | New Radius               | Double     |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 3: initialize         | Speed                    | VarLong    |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 3: initialize         | Portal Teleport Boundary | VarInt     | Resulting coordinates from a portal teleport are limited to +-value. Usually 29999984.                                                                                                                      |
	 * | 0x44      | Play  | Client   | 3: initialize         | Warning Time             | VarInt     |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 3: initialize         | Warning Blocks           | VarInt     |                                                                                                                                                                                                             |
	 * | 0x44      | Play  | Client   | 4: set warning time   | Warning Time             | VarInt     | unit?                                                                                                                                                                                                       |
	 * | 0x44      | Play  | Client   | 5: set warning blocks | Warning Blocks           | VarInt     |                                                                                                                                                                                                             |
	 * 
	 */
}

// ID=0x45
type PlayTitlePkt struct {
	/*
	 * | Packet ID | State | Bound To | Field Name               | Field Name    | Field Type | Notes |
	 * |-----------|-------|----------|--------------------------|---------------|------------|-------|
	 * | 0x45      | Play  | Client   | Action                   | Action        | VarInt     |       |
	 * | 0x45      | Play  | Client   | Action                   | Field Name    |            |       |
	 * | 0x45      | Play  | Client   | 0: set title             | Title Text    | Chat       |       |
	 * | 0x45      | Play  | Client   | 1: set subtitle          | Subtitle Text | Chat       |       |
	 * | 0x45      | Play  | Client   | 2: set times and display | Fade In       | Int        | ticks |
	 * | 0x45      | Play  | Client   | 2: set times and display | Stay          | Int        | ticks |
	 * | 0x45      | Play  | Client   | 2: set times and display | Fade Out      | Int        | ticks |
	 * | 0x45      | Play  | Client   | 3: hide                  | no fields     | no fields  |       |
	 * | 0x45      | Play  | Client   | 4: reset                 | no fields     | no fields  |       |
	 * 
	 */
}

// ID=0x46
type PlaySetCompressionPkt struct {
	/* Packets of this size or higher may be compressed */
	Threshold VarInt // VarInt
}

var _ Packet = (*PlaySetCompressionPkt)(nil)

func (p PlaySetCompressionPkt)Encode(b *PacketBuilder){
	b.VarInt(p.Threshold)
}

func (p *PlaySetCompressionPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.Threshold, ok = r.VarInt(); !ok {
		return io.EOF
	}
	return nil
}

// ID=0x47
type PlayerListHeaderAndFooterPkt = internal.PlayerListHeaderAndFooter_758_0

// ID=0x48
type PlayResourcePackSendPkt = internal.PlayResourcePackSend_754_1

// ID=0x49
type PlayUpdateEntityNBTPkt struct {
	EntityID VarInt // VarInt
	Tag nbt.NBT // NBT Tag
}

var _ Packet = (*PlayUpdateEntityNBTPkt)(nil)

func (p PlayUpdateEntityNBTPkt)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	WriteNBT(b, p.Tag)
}

func (p *PlayUpdateEntityNBTPkt)DecodeFrom(r *PacketReader)(error){
	var ok bool
	_ = ok
	var err error
	_ = err
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Tag, err = nbt.ReadNBT(r); err != nil {
		return err
	}
	return nil
}

// ======== END play ========
