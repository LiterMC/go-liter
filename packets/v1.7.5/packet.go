
// Generated at 2023-09-05 22:06:22.506 -06:00
// Origin: https://wiki.vg/index.php?title=Protocol&oldid=5486
// Protocol: 4
// Protocol Name: 1.7.5

package packet_1_7_5

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
type LoginStartPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes |
	 * |-----------|------------|------------|-------|
	 * | 0x00      | Name       | String     |       |
	 * 
	 */
}

// ID=0x1
type LoginEncryptionResponsePkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes                  |
	 * |-----------|---------------|------------|------------------------|
	 * | 0x01      | Length        | Short      | length of public key   |
	 * | 0x01      | Shared Secret | Byte array |                        |
	 * | 0x01      | Length        | Short      | length of verify token |
	 * | 0x01      | Verify Token  | Byte array |                        |
	 * 
	 */
}

// ---- login: clientbound ----

// ID=0x0
type LoginDisconnectPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes |
	 * |-----------|------------|------------|-------|
	 * | 0x00      | JSON Data  | String     |       |
	 * 
	 */
}

// ID=0x1
type LoginEncryptionRequestPkt struct {
	/*
	 * | Packet ID | Field Name   | Field Type | Notes                  |
	 * |-----------|--------------|------------|------------------------|
	 * | 0x01      | Server ID    | String     |                        |
	 * | 0x01      | Length       | Short      | length of public key   |
	 * | 0x01      | Public Key   | Byte array |                        |
	 * | 0x01      | Length       | Short      | length of verify token |
	 * | 0x01      | Verify Token | Byte array |                        |
	 * 
	 */
}

// ID=0x2
type LoginSuccessPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes |
	 * |-----------|------------|------------|-------|
	 * | 0x02      | UUID       | String     |       |
	 * | 0x02      | Username   | String     |       |
	 * 
	 */
}

// ======== END login ========

// ======== BEGIN play ========
// ---- play: serverbound ----

// ID=0x0
type PlayKeepAliveServerPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes |
	 * |-----------|---------------|------------|-------|
	 * | 0x00      | Keep Alive ID | Int        |       |
	 * 
	 */
}

// ID=0x1
type PlayChatMessageServerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes |
	 * |-----------|------------|------------|-------|
	 * | 0x01      | Message    | String     |       |
	 * 
	 */
}

// ID=0x2
type PlayUseEntityPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                           |
	 * |-----------|------------|------------|---------------------------------|
	 * | 0x02      | Target     | Int        |                                 |
	 * | 0x02      | Mouse      | Byte       | 0 = Left-click, 1 = Right-click |
	 * 
	 */
}

// ID=0x3
type PlayerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                |
	 * |-----------|------------|------------|------------------------------------------------------|
	 * | 0x03      | On Ground  | Bool       | True if the client is on the ground, False otherwise |
	 * 
	 */
}

// ID=0x4
type PlayerPositionPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                                                                                          |
	 * |-----------|------------|------------|--------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x04      | X          | Double     | Absolute position                                                                                                              |
	 * | 0x04      | FeetY      | Double     | Absolute feet position, normally HeadY - 1.62. Used to modify the players bounding box when going up stairs, crouching, etc… |
	 * | 0x04      | HeadY      | Double     | Absolute head position                                                                                                         |
	 * | 0x04      | Z          | Double     | Absolute position                                                                                                              |
	 * | 0x04      | On Ground  | Bool       | True if the client is on the ground, False otherwise                                                                           |
	 * 
	 */
}

// ID=0x5
type PlayerLookPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                |
	 * |-----------|------------|------------|------------------------------------------------------|
	 * | 0x05      | Yaw        | Float      | Absolute rotation on the X Axis, in degrees          |
	 * | 0x05      | Pitch      | Float      | Absolute rotation on the Y Axis, in degrees          |
	 * | 0x05      | On Ground  | Bool       | True if the client is on the ground, False otherwise |
	 * 
	 */
}

// ID=0x6
type PlayerPositionAndLookServerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                                                                                             |
	 * |-----------|------------|------------|-----------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x06      | X          | Double     | Absolute position                                                                                                                 |
	 * | 0x06      | FeetY      | Double     | Absolute feet position. Is normally HeadY - 1.62. Used to modify the players bounding box when going up stairs, crouching, etc… |
	 * | 0x06      | HeadY      | Double     | Absolute head position                                                                                                            |
	 * | 0x06      | Z          | Double     | Absolute position                                                                                                                 |
	 * | 0x06      | Yaw        | Float      | Absolute rotation on the X Axis, in degrees                                                                                       |
	 * | 0x06      | Pitch      | Float      | Absolute rotation on the Y Axis, in degrees                                                                                       |
	 * | 0x06      | On Ground  | Bool       | True if the client is on the ground, False otherwise                                                                              |
	 * 
	 */
}

// ID=0x7
type PlayerDiggingPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes                                                         |
	 * |-----------|------------|---------------|---------------------------------------------------------------|
	 * | 0x07      | Status     | Byte          | The action the player is taking against the block (see below) |
	 * | 0x07      | X          | Int           | Block position                                                |
	 * | 0x07      | Y          | Unsigned Byte | Block position                                                |
	 * | 0x07      | Z          | Int           | Block position                                                |
	 * | 0x07      | Face       | Byte          | The face being hit (see below)                                |
	 * 
	 */
}

// ID=0x8
type PlayerBlockPlacementPkt struct {
	/*
	 * | Packet ID | Field Name        | Field Type    | Notes                                                  |
	 * |-----------|-------------------|---------------|--------------------------------------------------------|
	 * | 0x08      | X                 | Int           | Block position                                         |
	 * | 0x08      | Y                 | Unsigned Byte | Block position                                         |
	 * | 0x08      | Z                 | Int           | Block position                                         |
	 * | 0x08      | Direction         | Byte          | The offset to use for block/item placement (see below) |
	 * | 0x08      | Held item         | Slot          |                                                        |
	 * | 0x08      | Cursor position X | Byte          | The position of the crosshair on the block             |
	 * | 0x08      | Cursor position Y | Byte          |                                                        |
	 * | 0x08      | Cursor position Z | Byte          |                                                        |
	 * 
	 */
}

// ID=0x9
type PlayHeldItemChangeServerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                        |
	 * |-----------|------------|------------|----------------------------------------------|
	 * | 0x09      | Slot       | Short      | The slot which the player has selected (0-8) |
	 * 
	 */
}

// ID=0xa
type PlayAnimationServerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes        |
	 * |-----------|------------|------------|--------------|
	 * | 0x0A      | Entity ID  | Int        | Player ID    |
	 * | 0x0A      | Animation  | Byte       | Animation ID |
	 * 
	 */
}

// ID=0xb
type PlayEntityActionPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                   |
	 * |-----------|------------|------------|-----------------------------------------|
	 * | 0x0B      | Entity ID  | Int        | Player ID                               |
	 * | 0x0B      | Action ID  | Byte       | The ID of the action, see below.        |
	 * | 0x0B      | Jump Boost | Int        | Horse jump boost. Ranged from 0 -> 100. |
	 * 
	 */
}

// ID=0xc
type PlaySteerVehiclePkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                              |
	 * |-----------|------------|------------|------------------------------------|
	 * | 0x0C      | Sideways   | Float      | Positive to the left of the player |
	 * | 0x0C      | Forward    | Float      | Positive forward                   |
	 * | 0x0C      | Jump       | Bool       |                                    |
	 * | 0x0C      | Unmount    | Bool       | True when leaving the vehicle      |
	 * 
	 */
}

// ID=0xd
type PlayCloseWindowServerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                          |
	 * |-----------|------------|------------|----------------------------------------------------------------|
	 * | 0x0D      | Window id  | byte       | This is the id of the window that was closed. 0 for inventory. |
	 * 
	 */
}

// ID=0xe
type PlayClickWindowPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes                                                                                       |
	 * |-----------|---------------|------------|---------------------------------------------------------------------------------------------|
	 * | 0x0E      | Window ID     | Byte       | The id of the window which was clicked. 0 for player inventory.                             |
	 * | 0x0E      | Slot          | Short      | The clicked slot. See below.                                                                |
	 * | 0x0E      | Button        | Byte       | The button used in the click. See below.                                                    |
	 * | 0x0E      | Action number | Short      | A unique number for the action, used for transaction handling (See the Transaction packet). |
	 * | 0x0E      | Mode          | Byte       | Inventory operation mode. See below.                                                        |
	 * | 0x0E      | Clicked item  | Slot       |                                                                                             |
	 * 
	 */
}

// ID=0xf
type PlayConfirmTransactionServerPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes                                                                                           |
	 * |-----------|---------------|------------|-------------------------------------------------------------------------------------------------|
	 * | 0x0F      | Window ID     | Byte       | The id of the window that the action occurred in.                                               |
	 * | 0x0F      | Action number | Short      | Every action that is to be accepted has a unique number. This field corresponds to that number. |
	 * | 0x0F      | Accepted      | Bool       | Whether the action was accepted.                                                                |
	 * 
	 */
}

// ID=0x10
type PlayCreativeInventoryActionPkt struct {
	/*
	 * | Packet ID | Field Name   | Field Type | Notes          |
	 * |-----------|--------------|------------|----------------|
	 * | 0x10      | Slot         | Short      | Inventory slot |
	 * | 0x10      | Clicked item | Slot       |                |
	 * 
	 */
}

// ID=0x11
type PlayEnchantItemPkt struct {
	/*
	 * | Packet ID | Field Name  | Field Type | Notes                                                                                                |
	 * |-----------|-------------|------------|------------------------------------------------------------------------------------------------------|
	 * | 0x11      | Window ID   | Byte       | The ID sent by Open Window                                                                           |
	 * | 0x11      | Enchantment | Byte       | The position of the enchantment on the enchantment table window, starting with 0 as the topmost one. |
	 * 
	 */
}

// ID=0x12
type PlayUpdateSignServerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                           |
	 * |-----------|------------|------------|---------------------------------|
	 * | 0x12      | X          | Int        | Block X Coordinate              |
	 * | 0x12      | Y          | Short      | Block Y Coordinate              |
	 * | 0x12      | Z          | Int        | Block Z Coordinate              |
	 * | 0x12      | Line 1     | String     | First line of text in the sign  |
	 * | 0x12      | Line 2     | String     | Second line of text in the sign |
	 * | 0x12      | Line 3     | String     | Third line of text in the sign  |
	 * | 0x12      | Line 4     | String     | Fourth line of text in the sign |
	 * 
	 */
}

// ID=0x13
type PlayerAbilitiesServerPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes                                 |
	 * |-----------|---------------|------------|---------------------------------------|
	 * | 0x13      | Flags         | Byte       |                                       |
	 * | 0x13      | Flying speed  | Float      | previous integer value divided by 250 |
	 * | 0x13      | Walking speed | Float      | previous integer value divided by 250 |
	 * 
	 */
}

// ID=0x14
type PlayTabCompleteServerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes |
	 * |-----------|------------|------------|-------|
	 * | 0x14      | Text       | String     |       |
	 * 
	 */
}

// ID=0x15
type PlayClientSettingsPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes                                   |
	 * |-----------|---------------|------------|-----------------------------------------|
	 * | 0x15      | Locale        | String     | en_GB                                   |
	 * | 0x15      | View distance | Byte       | Client-side render distance(chunks)     |
	 * | 0x15      | Chat flags    | Byte       | Chat settings. See notes below.         |
	 * | 0x15      | Chat colours  | Bool       | "Colours" multiplayer setting           |
	 * | 0x15      | Difficulty    | Byte       | Client-side difficulty from options.txt |
	 * | 0x15      | Show Cape     | Bool       | Client-side "show cape" option          |
	 * 
	 */
}

// ID=0x16
type PlayClientStatusPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes     |
	 * |-----------|------------|------------|-----------|
	 * | 0x16      | Action ID  | Byte       | See below |
	 * 
	 */
}

// ID=0x17
type PlayPluginMessageServerPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                        |
	 * |-----------|------------|------------|----------------------------------------------|
	 * | 0x17      | Channel    | String     | Name of the "channel" used to send the data. |
	 * | 0x17      | Length     | Short      | Length of the following byte array           |
	 * | 0x17      | Data       | Byte Array | Any data.                                    |
	 * 
	 */
}

// ---- play: clientbound ----

// ID=0x0
type PlayKeepAliveClientPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes |
	 * |-----------|---------------|------------|-------|
	 * | 0x00      | Keep Alive ID | Int        |       |
	 * 
	 */
}

// ID=0x1
type PlayJoinGamePkt struct {
	/*
	 * | Packet ID | Field Name  | Field Type    | Notes                                                                    |
	 * |-----------|-------------|---------------|--------------------------------------------------------------------------|
	 * | 0x01      | Entity ID   | Int           | The player's Entity ID                                                   |
	 * | 0x01      | Gamemode    | Unsigned Byte | 0: survival, 1: creative, 2: adventure. Bit 3 (0x8) is the hardcore flag |
	 * | 0x01      | Dimension   | Byte          | -1: nether, 0: overworld, 1: end                                         |
	 * | 0x01      | Difficulty  | Unsigned Byte | 0 thru 3 for Peaceful, Easy, Normal, Hard                                |
	 * | 0x01      | Max Players | Unsigned Byte | Used by the client to draw the player list                               |
	 * | 0x01      | Level Type  | String        | default, flat, largeBiomes, amplified, default_1_1                       |
	 * 
	 */
}

// ID=0x2
type PlayChatMessageClientPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                                              |
	 * |-----------|------------|------------|------------------------------------------------------------------------------------|
	 * | 0x02      | JSON Data  | String     | https://gist.github.com/thinkofdeath/e882ce057ed83bac0a1c , Limited to 32767 bytes |
	 * 
	 */
}

// ID=0x3
type PlayTimeUpdatePkt struct {
	/*
	 * | Packet ID | Field Name       | Field Type | Notes                                                                                                  |
	 * |-----------|------------------|------------|--------------------------------------------------------------------------------------------------------|
	 * | 0x03      | Age of the world | Long       | In ticks; not changed by server commands                                                               |
	 * | 0x03      | Time of day      | Long       | The world (or region) time, in ticks. If negative the sun will stop moving at the Math.abs of the time |
	 * 
	 */
}

// ID=0x4
type PlayEntityEquipmentPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                                                        |
	 * |-----------|------------|------------|----------------------------------------------------------------------------------------------|
	 * | 0x04      | EntityID   | Int        | Entity's ID                                                                                  |
	 * | 0x04      | Slot       | Short      | Equipment slot: 0=held, 1-4=armor slot (1 - boots, 2 - leggings, 3 - chestplate, 4 - helmet) |
	 * | 0x04      | Item       | Slot       | Item in slot format                                                                          |
	 * 
	 */
}

// ID=0x5
type PlaySpawnPositionPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                        |
	 * |-----------|------------|------------|------------------------------|
	 * | 0x05      | X          | Int        | Spawn X in block coordinates |
	 * | 0x05      | Y          | Int        | Spawn Y in block coordinates |
	 * | 0x05      | Z          | Int        | Spawn Z in block coordinates |
	 * 
	 */
}

// ID=0x6
type PlayUpdateHealthPkt struct {
	/*
	 * | Packet ID | Field Name      | Field Type | Notes                                               |
	 * |-----------|-----------------|------------|-----------------------------------------------------|
	 * | 0x06      | Health          | Float      | 0 or less = dead, 20 = full HP                      |
	 * | 0x06      | Food            | Short      | 0 - 20                                              |
	 * | 0x06      | Food Saturation | Float      | Seems to vary from 0.0 to 5.0 in integer increments |
	 * 
	 */
}

// ID=0x7
type PlayRespawnPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes                                                                     |
	 * |-----------|------------|---------------|---------------------------------------------------------------------------|
	 * | 0x07      | Dimension  | Int           | -1: The Nether, 0: The Overworld, 1: The End                              |
	 * | 0x07      | Difficulty | Unsigned Byte | 0 thru 3 for Peaceful, Easy, Normal, Hard.                                |
	 * | 0x07      | Gamemode   | Unsigned Byte | 0: survival, 1: creative, 2: adventure. The hardcore flag is not included |
	 * | 0x07      | Level Type | String        | Same as Join Game                                                         |
	 * 
	 */
}

// ID=0x8
type PlayerPositionAndLookClientPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                |
	 * |-----------|------------|------------|------------------------------------------------------|
	 * | 0x08      | X          | Double     | Absolute position                                    |
	 * | 0x08      | Y          | Double     | Absolute position (eyes pos)                         |
	 * | 0x08      | Z          | Double     | Absolute position                                    |
	 * | 0x08      | Yaw        | Float      | Absolute rotation on the X Axis, in degrees          |
	 * | 0x08      | Pitch      | Float      | Absolute rotation on the Y Axis, in degrees          |
	 * | 0x08      | On Ground  | Bool       | True if the client is on the ground, False otherwise |
	 * 
	 */
}

// ID=0x9
type PlayHeldItemChangeClientPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                        |
	 * |-----------|------------|------------|----------------------------------------------|
	 * | 0x09      | Slot       | Byte       | The slot which the player has selected (0-8) |
	 * 
	 */
}

// ID=0xa
type PlayUseBedPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes                               |
	 * |-----------|------------|---------------|-------------------------------------|
	 * | 0x0A      | Entity ID  | Int           | Player ID                           |
	 * | 0x0A      | X          | Int           | Bed headboard X as block coordinate |
	 * | 0x0A      | Y          | Unsigned Byte | Bed headboard Y as block coordinate |
	 * | 0x0A      | Z          | Int           | Bed headboard Z as block coordinate |
	 * 
	 */
}

// ID=0xb
type PlayAnimationClientPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes        |
	 * |-----------|------------|---------------|--------------|
	 * | 0x0B      | Entity ID  | VarInt        | Player ID    |
	 * | 0x0B      | Animation  | Unsigned Byte | Animation ID |
	 * 
	 */
}

// ID=0xc
type PlaySpawnPlayerPkt struct {
	/*
	 * | Packet ID | Field Name   | Field Type | Notes                                                                                                                                                  |
	 * |-----------|--------------|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x0C      | Entity ID    | VarInt     | Player's Entity ID                                                                                                                                     |
	 * | 0x0C      | Player UUID  | String     | Player's UUID                                                                                                                                          |
	 * | 0x0C      | Player Name  | String     | Player's Name                                                                                                                                          |
	 * | 0x0C      | X            | Int        | Player X as a Fixed-Point number                                                                                                                       |
	 * | 0x0C      | Y            | Int        | Player X as a Fixed-Point number                                                                                                                       |
	 * | 0x0C      | Z            | Int        | Player X as a Fixed-Point number                                                                                                                       |
	 * | 0x0C      | Yaw          | Byte       | Player rotation as a packed byte                                                                                                                       |
	 * | 0x0C      | Pitch        | Byte       | Player rotation as a packet byte                                                                                                                       |
	 * | 0x0C      | Current Item | Short      | The item the player is currently holding. Note that this should be 0 for "no item", unlike -1 used in other packets. A negative value crashes clients. |
	 * | 0x0C      | Metadata     | Metadata   | The client will crash if no metadata is sent                                                                                                           |
	 * 
	 */
}

// ID=0xd
type PlayCollectItemPkt struct {
	/*
	 * | Packet ID | Field Name          | Field Type | Notes |
	 * |-----------|---------------------|------------|-------|
	 * | 0x0D      | Collected Entity ID | Int        |       |
	 * | 0x0D      | Collector Entity ID | Int        |       |
	 * 
	 */
}

// ID=0xe
type PlaySpawnObjectPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type  | Notes                              |
	 * |-----------|------------|-------------|------------------------------------|
	 * | 0x0E      | Entity ID  | VarInt      | Entity ID of the object            |
	 * | 0x0E      | Type       | Byte        | The type of object (See Objects)   |
	 * | 0x0E      | X          | Int         | X position as a Fixed-Point number |
	 * | 0x0E      | Y          | Int         | Y position as a Fixed-Point number |
	 * | 0x0E      | Z          | Int         | Z position as a Fixed-Point number |
	 * | 0x0E      | Pitch      | Byte        | The pitch in steps of 2p/256       |
	 * | 0x0E      | Yaw        | Byte        | The yaw in steps of 2p/256         |
	 * | 0x0E      | Data       | Object Data |                                    |
	 * 
	 */
}

// ID=0xf
type PlaySpawnMobPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes                              |
	 * |-----------|------------|---------------|------------------------------------|
	 * | 0x0F      | Entity ID  | VarInt        | Entity's ID                        |
	 * | 0x0F      | Type       | Unsigned Byte | The type of mob. See Mobs          |
	 * | 0x0F      | X          | Int           | X position as a Fixed-Point number |
	 * | 0x0F      | Y          | Int           | Y position as a Fixed-Point number |
	 * | 0x0F      | Z          | Int           | Z position as a Fixed-Point number |
	 * | 0x0F      | Pitch      | Byte          | The pitch in steps of 2p/256       |
	 * | 0x0F      | Head Pitch | Byte          | The pitch in steps of 2p/256       |
	 * | 0x0F      | Yaw        | Byte          | The yaw in steps of 2p/256         |
	 * | 0x0F      | Velocity X | Short         |                                    |
	 * | 0x0F      | Velocity Y | Short         |                                    |
	 * | 0x0F      | Velocity Z | Short         |                                    |
	 * | 0x0F      | Metadata   | Metadata      |                                    |
	 * 
	 */
}

// ID=0x10
type PlaySpawnPaintingPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                 |
	 * |-----------|------------|------------|-------------------------------------------------------|
	 * | 0x10      | Entity ID  | VarInt     | Entity's ID                                           |
	 * | 0x10      | Title      | String     | Name of the painting. Max length 13                   |
	 * | 0x10      | X          | Int        | Center X coordinate                                   |
	 * | 0x10      | Y          | Int        | Center Y coordinate                                   |
	 * | 0x10      | Z          | Int        | Center Z coordinate                                   |
	 * | 0x10      | Direction  | Int        | Direction the painting faces (0 -z, 1 -x, 2 +z, 3 +x) |
	 * 
	 */
}

// ID=0x11
type PlaySpawnExperienceOrbPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                        |
	 * |-----------|------------|------------|--------------------------------------------------------------|
	 * | 0x11      | Entity ID  | VarInt     | Entity's ID                                                  |
	 * | 0x11      | X          | Int        | X position as a Fixed-Point number                           |
	 * | 0x11      | Y          | Int        | Y position as a Fixed-Point number                           |
	 * | 0x11      | Z          | Int        | Z position as a Fixed-Point number                           |
	 * | 0x11      | Count      | Short      | The amount of experience this orb will reward once collected |
	 * 
	 */
}

// ID=0x12
type PlayEntityVelocityPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                  |
	 * |-----------|------------|------------|------------------------|
	 * | 0x12      | Entity ID  | Int        | Entity's ID            |
	 * | 0x12      | Velocity X | Short      | Velocity on the X axis |
	 * | 0x12      | Velocity Y | Short      | Velocity on the Y axis |
	 * | 0x12      | Velocity Z | Short      | Velocity on the Z axis |
	 * 
	 */
}

// ID=0x13
type PlayDestroyEntitiesPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type   | Notes                           |
	 * |-----------|------------|--------------|---------------------------------|
	 * | 0x13      | Count      | Byte         | Length of following array       |
	 * | 0x13      | Entity IDs | Array of Int | The list of entities of destroy |
	 * 
	 */
}

// ID=0x14
type PlayEntityPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes       |
	 * |-----------|------------|------------|-------------|
	 * | 0x14      | Entity ID  | Int        | Entity's ID |
	 * 
	 */
}

// ID=0x15
type PlayEntityRelativeMovePkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                        |
	 * |-----------|------------|------------|----------------------------------------------|
	 * | 0x15      | Entity ID  | Int        | Entity's ID                                  |
	 * | 0x15      | DX         | Byte       | Change in X position as a Fixed-Point number |
	 * | 0x15      | DY         | Byte       | Change in Y position as a Fixed-Point number |
	 * | 0x15      | DZ         | Byte       | Change in Z position as a Fixed-Point number |
	 * 
	 */
}

// ID=0x16
type PlayEntityLookPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                    |
	 * |-----------|------------|------------|------------------------------------------|
	 * | 0x16      | Entity ID  | Int        | Entity's ID                              |
	 * | 0x16      | Yaw        | Byte       | The X Axis rotation as a fraction of 360 |
	 * | 0x16      | Pitch      | Byte       | The Y Axis rotation as a fraction of 360 |
	 * 
	 */
}

// ID=0x17
type PlayEntityLookAndRelativeMovePkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                        |
	 * |-----------|------------|------------|----------------------------------------------|
	 * | 0x17      | Entity ID  | Int        | Entity's ID                                  |
	 * | 0x17      | DX         | Byte       | Change in X position as a Fixed-Point number |
	 * | 0x17      | DY         | Byte       | Change in Y position as a Fixed-Point number |
	 * | 0x17      | DZ         | Byte       | Change in Z position as a Fixed-Point number |
	 * | 0x17      | Yaw        | Byte       | The X Axis rotation as a fraction of 360     |
	 * | 0x17      | Pitch      | Byte       | The Y Axis rotation as a fraction of 360     |
	 * 
	 */
}

// ID=0x18
type PlayEntityTeleportPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                    |
	 * |-----------|------------|------------|------------------------------------------|
	 * | 0x18      | Entity ID  | Int        | Entity's ID                              |
	 * | 0x18      | X          | Int        | X position as a Fixed-Point number       |
	 * | 0x18      | Y          | Int        | Y position as a Fixed-Point number       |
	 * | 0x18      | Z          | Int        | Z position as a Fixed-Point number       |
	 * | 0x18      | Yaw        | Byte       | The X Axis rotation as a fraction of 360 |
	 * | 0x18      | Pitch      | Byte       | The Y Axis rotation as a fraction of 360 |
	 * 
	 */
}

// ID=0x19
type PlayEntityHeadLookPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                       |
	 * |-----------|------------|------------|-----------------------------|
	 * | 0x19      | Entity ID  | Int        | Entity's ID                 |
	 * | 0x19      | Head Yaw   | Byte       | Head yaw in steps of 2p/256 |
	 * 
	 */
}

// ID=0x1a
type PlayEntityStatusPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes       |
	 * |-----------|---------------|------------|-------------|
	 * | 0x1A      | Entity ID     | Int        | Entity's ID |
	 * | 0x1A      | Entity Status | Byte       | See below   |
	 * 
	 */
}

// ID=0x1b
type PlayAttachEntityPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                     |
	 * |-----------|------------|------------|-------------------------------------------|
	 * | 0x1B      | Entity ID  | Int        | Entity's ID                               |
	 * | 0x1B      | Vehicle ID | Int        | Vechicle's Entity ID                      |
	 * | 0x1B      | Leash      | Bool       | If true leashes the entity to the vehicle |
	 * 
	 */
}

// ID=0x1c
type PlayEntityMetadataPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes       |
	 * |-----------|------------|------------|-------------|
	 * | 0x1C      | Entity ID  | Int        | Entity's ID |
	 * | 0x1C      | Metadata   | Metadata   |             |
	 * 
	 */
}

// ID=0x1d
type PlayEntityEffectPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes       |
	 * |-----------|------------|------------|-------------|
	 * | 0x1D      | Entity ID  | Int        | Entity's ID |
	 * | 0x1D      | Effect ID  | Byte       | See [[1]]   |
	 * | 0x1D      | Amplifier  | Byte       |             |
	 * | 0x1D      | Duration   | Short      |             |
	 * 
	 */
}

// ID=0x1e
type PlayRemoveEntityEffectPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes       |
	 * |-----------|------------|------------|-------------|
	 * | 0x1E      | Entity ID  | Int        | Entity's ID |
	 * | 0x1E      | Effect ID  | Byte       |             |
	 * 
	 */
}

// ID=0x1f
type PlaySetExperiencePkt struct {
	/*
	 * | Packet ID | Field Name       | Field Type | Notes           |
	 * |-----------|------------------|------------|-----------------|
	 * | 0x1F      | Experience bar   | Float      | Between 0 and 1 |
	 * | 0x1F      | Level            | Short      |                 |
	 * | 0x1F      | Total Experience | Short      |                 |
	 * 
	 */
}

// ID=0x20
type PlayEntityPropertiesPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type             | Notes                     |
	 * |-----------|------------|------------------------|---------------------------|
	 * | 0x20      | Entity ID  | Int                    | Entity's ID               |
	 * | 0x20      | Count      | Int                    | Length of following array |
	 * | 0x20      | Properties | Array of Property Data |                           |
	 * 
	 */
}

// ID=0x21
type PlayChunkDataPkt struct {
	/*
	 * | Packet ID | Field Name           | Field Type     | Notes                                                                                                                                                                  |
	 * |-----------|----------------------|----------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x21      | Chunk X              | Int            | Chunk X coordinate                                                                                                                                                     |
	 * | 0x21      | Chunk Z              | Int            | Chunk Z coordinate                                                                                                                                                     |
	 * | 0x21      | Ground-Up continuous | Boolean        | This is True if the packet represents all sections in this vertical column, where the primary bit map specifies exactly which sections are included, and which are air |
	 * | 0x21      | Primary bit map      | Unsigned Short | Bitmask with 1 for every 16x16x16 section which data follows in the compressed data.                                                                                   |
	 * | 0x21      | Add bit map          | Unsigned Short | Same as above, but this is used exclusively for the 'add' portion of the payload                                                                                       |
	 * | 0x21      | Compressed size      | Int            | Size of compressed chunk data                                                                                                                                          |
	 * | 0x21      | Compressed data      | Byte array     | The chunk data is compressed using Zlib Deflate                                                                                                                        |
	 * 
	 */
}

// ID=0x22
type PlayMultiBlockChangePkt struct {
	/*
	 * | Packet ID | Field Name   | Field Type       | Notes                                                                 |
	 * |-----------|--------------|------------------|-----------------------------------------------------------------------|
	 * | 0x22      | Chunk X      | Int              | Chunk X coordinate                                                    |
	 * | 0x22      | Chunk Z      | Int              | Chunk Z Coordinate                                                    |
	 * | 0x22      | Record count | Short            | The number of blocks affected                                         |
	 * | 0x22      | Data size    | Int              | The total size of the data, in bytes. Should always be 4*record count |
	 * | 0x22      | Records      | Array of Records |                                                                       |
	 * 
	 */
}

// ID=0x23
type PlayBlockChangePkt struct {
	/*
	 * | Packet ID | Field Name     | Field Type    | Notes                          |
	 * |-----------|----------------|---------------|--------------------------------|
	 * | 0x23      | X              | Int           | Block X Coordinate             |
	 * | 0x23      | Y              | Unsigned Byte | Block Y Coordinate             |
	 * | 0x23      | Z              | Int           | Block Z Coordinate             |
	 * | 0x23      | Block ID       | VarInt        | The new block ID for the block |
	 * | 0x23      | Block Metadata | Unsigned Byte | The new metadata for the block |
	 * 
	 */
}

// ID=0x24
type PlayBlockActionPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes                                         |
	 * |-----------|------------|---------------|-----------------------------------------------|
	 * | 0x24      | X          | Int           | Block X Coordinate                            |
	 * | 0x24      | Y          | Short         | Block Y Coordinate                            |
	 * | 0x24      | Z          | Int           | Block Z Coordinate                            |
	 * | 0x24      | Byte 1     | Unsigned Byte | Varies depending on block - see Block_Actions |
	 * | 0x24      | Byte 2     | Unsigned Byte | Varies depending on block - see Block_Actions |
	 * | 0x24      | Block Type | VarInt        | The block type for the block                  |
	 * 
	 */
}

// ID=0x25
type PlayBlockBreakAnimationPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes          |
	 * |-----------|---------------|------------|----------------|
	 * | 0x25      | Entity ID     | VarInt     | Entity's ID    |
	 * | 0x25      | X             | Int        | Block Position |
	 * | 0x25      | Y             | Int        |                |
	 * | 0x25      | Z             | Int        |                |
	 * | 0x25      | Destroy Stage | Byte       | Block Position |
	 * 
	 */
}

// ID=0x26
type PlayMapChunkBulkPkt struct {
	/*
	 * | Packet ID | Field Name         | Field Type | Notes                                                                                                                  |
	 * |-----------|--------------------|------------|------------------------------------------------------------------------------------------------------------------------|
	 * | 0x26      | Chunk column count | Short      | The number of chunk in this packet                                                                                     |
	 * | 0x26      | Data length        | Int        | The size of the data field                                                                                             |
	 * | 0x26      | Sky light sent     | Bool       | Whether or not the chunk data contains a light nibble array. This is true in the main world, false in the end + nether |
	 * | 0x26      | Data               | Byte Array | Compressed chunk data                                                                                                  |
	 * | 0x26      | Meta information   | Meta       | See below                                                                                                              |
	 * 
	 */
}

// ID=0x27
type PlayExplosionPkt struct {
	/*
	 * | Packet ID | Field Name      | Field Type                  | Notes                                                                                                 |
	 * |-----------|-----------------|-----------------------------|-------------------------------------------------------------------------------------------------------|
	 * | 0x27      | X               | Float                       |                                                                                                       |
	 * | 0x27      | Y               | Float                       |                                                                                                       |
	 * | 0x27      | Z               | Float                       |                                                                                                       |
	 * | 0x27      | Radius          | Float                       | Currently unused in the client                                                                        |
	 * | 0x27      | Record count    | Int                         | This is the count, not the size. The size is 3 times this value.                                      |
	 * | 0x27      | Records         | (Byte, Byte, Byte) × count | Each record is 3 signed bytes long, each bytes are the XYZ (respectively) offsets of affected blocks. |
	 * | 0x27      | Player Motion X | Float                       | X velocity of the player being pushed by the explosion                                                |
	 * | 0x27      | Player Motion Y | Float                       | Y velocity of the player being pushed by the explosion                                                |
	 * | 0x27      | Player Motion Z | Float                       | Z velocity of the player being pushed by the explosion                                                |
	 * 
	 */
}

// ID=0x28
type PlayEffectPkt struct {
	/*
	 * | Packet ID | Field Name              | Field Type | Notes                                      |
	 * |-----------|-------------------------|------------|--------------------------------------------|
	 * | 0x28      | Effect ID               | Int        | The ID of the effect, see below.           |
	 * | 0x28      | X                       | Int        | The X location of the effect               |
	 * | 0x28      | Y                       | Byte       | The Y location of the effect               |
	 * | 0x28      | Z                       | Int        | The Z location of the effect               |
	 * | 0x28      | Data                    | Int        | Extra data for certain effects, see below. |
	 * | 0x28      | Disable relative volume | Bool       | See above                                  |
	 * 
	 */
}

// ID=0x29
type PlaySoundEffectPkt struct {
	/*
	 * | Packet ID | Field Name        | Field Type    | Notes                    |
	 * |-----------|-------------------|---------------|--------------------------|
	 * | 0x29      | Sound name        | String        |                          |
	 * | 0x29      | Effect position X | Int           | Effect X multiplied by 8 |
	 * | 0x29      | Effect position Y | Int           | Effect Y multiplied by 8 |
	 * | 0x29      | Effect position Z | Int           | Effect Z multiplied by 8 |
	 * | 0x29      | Volume            | Float         | 1 is 100%, can be more   |
	 * | 0x29      | Pitch             | Unsigned Byte | 63 is 100%, can be more  |
	 * 
	 */
}

// ID=0x2a
type PlayParticlePkt struct {
	/*
	 * | Packet ID | Field Name          | Field Type | Notes                                                                           |
	 * |-----------|---------------------|------------|---------------------------------------------------------------------------------|
	 * | 0x2A      | Particle name       | String     | The name of the particle to create. A list can be found here                    |
	 * | 0x2A      | X                   | Float      | X position of the particle                                                      |
	 * | 0x2A      | Y                   | Float      | Y position of the particle                                                      |
	 * | 0x2A      | Z                   | Float      | Z position of the particle                                                      |
	 * | 0x2A      | Offset X            | Float      | This is added to the X position after being multiplied by random.nextGaussian() |
	 * | 0x2A      | Offset Y            | Float      | This is added to the Y position after being multiplied by random.nextGaussian() |
	 * | 0x2A      | Offset Z            | Float      | This is added to the Z position after being multiplied by random.nextGaussian() |
	 * | 0x2A      | Particle data       | Float      | The data of each particle                                                       |
	 * | 0x2A      | Number of particles | Int        | The number of particles to create                                               |
	 * 
	 */
}

// ID=0x2b
type PlayChangeGameStatePkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes             |
	 * |-----------|------------|---------------|-------------------|
	 * | 0x2B      | Reason     | Unsigned Byte |                   |
	 * | 0x2B      | Value      | Float         | Depends on reason |
	 * 
	 */
}

// ID=0x2c
type PlaySpawnGlobalEntityPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                       |
	 * |-----------|------------|------------|-------------------------------------------------------------|
	 * | 0x2C      | Entity ID  | VarInt     | The entity ID of the thunderbolt                            |
	 * | 0x2C      | Type       | Byte       | The global entity type, currently always 1 for thunderbolt. |
	 * | 0x2C      | X          | Int        | Thunderbolt X a fixed-point number                          |
	 * | 0x2C      | Y          | Int        | Thunderbolt Y a fixed-point number                          |
	 * | 0x2C      | Z          | Int        | Thunderbolt Z a fixed-point number                          |
	 * 
	 */
}

// ID=0x2d
type PlayOpenWindowPkt struct {
	/*
	 * | Packet ID | Field Name                | Field Type    | Notes                                                                                                                 |
	 * |-----------|---------------------------|---------------|-----------------------------------------------------------------------------------------------------------------------|
	 * | 0x2D      | Window id                 | Unsigned Byte | A unique id number for the window to be displayed.  Notchian server implementation is a counter, starting at 1.       |
	 * | 0x2D      | Inventory Type            | Unsigned Byte | The window type to use for display.  Check below                                                                      |
	 * | 0x2D      | Window title              | String        | The title of the window.                                                                                              |
	 * | 0x2D      | Number of Slots           | Unsigned Byte | Number of slots in the window (excluding the number of slots in the player inventory).                                |
	 * | 0x2D      | Use provided window title | Bool          | If false, the client will look up a string like "window.minecart". If true, the client uses what the server provides. |
	 * | 0x2D      | Entity ID                 | Int           | EntityHorse's entityId. Only sent when window type is equal to 11 (AnimalChest).                                      |
	 * 
	 */
}

// ID=0x2e
type PlayCloseWindowClientPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes                                                          |
	 * |-----------|------------|---------------|----------------------------------------------------------------|
	 * | 0x2E      | Window ID  | Unsigned Byte | This is the id of the window that was closed. 0 for inventory. |
	 * 
	 */
}

// ID=0x2f
type PlaySetSlotPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                                                                                                                                                                                                                                                                                                                                                |
	 * |-----------|------------|------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x2F      | Window ID  | Byte       | The window which is being updated. 0 for player inventory. Note that all known window types include the player inventory. This packet will only be sent for the currently opened window while the player is performing actions, even if it affects the player inventory. After the window is closed, a number of these packets are sent to update the player's inventory window (0). |
	 * | 0x2F      | Slot       | Short      | The slot that should be updated                                                                                                                                                                                                                                                                                                                                                      |
	 * | 0x2F      | Slot data  | Slot       |                                                                                                                                                                                                                                                                                                                                                                                      |
	 * 
	 */
}

// ID=0x30
type PlayWindowItemsPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type     | Notes                                                                    |
	 * |-----------|------------|----------------|--------------------------------------------------------------------------|
	 * | 0x30      | Window ID  | Unsigned Byte  | The id of window which items are being sent for. 0 for player inventory. |
	 * | 0x30      | Count      | Short          | The number of slots (see below)                                          |
	 * | 0x30      | Slot data  | Array of Slots |                                                                          |
	 * 
	 */
}

// ID=0x31
type PlayWindowPropertyPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type    | Notes                             |
	 * |-----------|------------|---------------|-----------------------------------|
	 * | 0x31      | Window ID  | Unsigned Byte | The id of the window.             |
	 * | 0x31      | Property   | Short         | Which property should be updated. |
	 * | 0x31      | Value      | Short         | The new value for the property.   |
	 * 
	 */
}

// ID=0x32
type PlayConfirmTransactionClientPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type    | Notes                                                                                           |
	 * |-----------|---------------|---------------|-------------------------------------------------------------------------------------------------|
	 * | 0x32      | Window ID     | Unsigned Byte | The id of the window that the action occurred in.                                               |
	 * | 0x32      | Action number | Short         | Every action that is to be accepted has a unique number. This field corresponds to that number. |
	 * | 0x32      | Accepted      | Bool          | Whether the action was accepted.                                                                |
	 * 
	 */
}

// ID=0x33
type PlayUpdateSignClientPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                           |
	 * |-----------|------------|------------|---------------------------------|
	 * | 0x33      | X          | Int        | Block X Coordinate              |
	 * | 0x33      | Y          | Short      | Block Y Coordinate              |
	 * | 0x33      | Z          | Int        | Block Z Coordinate              |
	 * | 0x33      | Line 1     | String     | First line of text in the sign  |
	 * | 0x33      | Line 2     | String     | Second line of text in the sign |
	 * | 0x33      | Line 3     | String     | Third line of text in the sign  |
	 * | 0x33      | Line 4     | String     | Fourth line of text in the sign |
	 * 
	 */
}

// ID=0x34
type PlayMapsPkt struct {
	/*
	 * | Packet ID | Field Name  | Field Type | Notes                                      |
	 * |-----------|-------------|------------|--------------------------------------------|
	 * | 0x34      | Item Damage | VarInt     | The damage value of the map being modified |
	 * | 0x34      | Length      | Short      | Length of following byte array             |
	 * | 0x34      | Data        | Byte Array | Array data                                 |
	 * 
	 */
}

// ID=0x35
type PlayUpdateBlockEntityPkt struct {
	/*
	 * | Packet ID | Field Name  | Field Type    | Notes                                                    |
	 * |-----------|-------------|---------------|----------------------------------------------------------|
	 * | 0x35      | X           | Int           |                                                          |
	 * | 0x35      | Y           | Short         |                                                          |
	 * | 0x35      | Z           | Int           |                                                          |
	 * | 0x35      | Action      | Unsigned Byte | The type of update to perform                            |
	 * | 0x35      | Data length | Short         | Varies                                                   |
	 * | 0x35      | NBT Data    | Byte Array    | Present if data length > 0. Compressed with gzip. Varies |
	 * 
	 */
}

// ID=0x36
type PlaySignEditorOpenPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                  |
	 * |-----------|------------|------------|------------------------|
	 * | 0x36      | X          | Int        | X in block coordinates |
	 * | 0x36      | Y          | Int        | Y in block coordinates |
	 * | 0x36      | Z          | Int        | Z in block coordinates |
	 * 
	 */
}

// ID=0x37
type PlayStatisticsPkt struct {
	/*
	 * | Packet ID | Field Name | Field Name       | Field Type | Notes                                                     |
	 * |-----------|------------|------------------|------------|-----------------------------------------------------------|
	 * | 0x37      | Count      | Count            | VarInt     | Number of entrys                                          |
	 * | 0x37      | Entry      | Statistic's name | String     | https://gist.github.com/thinkofdeath/a1842c21a0cf2e1fb5e0 |
	 * | 0x37      | Entry      | Value            | VarInt     | The amount to set it to                                   |
	 * 
	 */
}

// ID=0x38
type PlayerListItemPkt struct {
	/*
	 * | Packet ID | Field Name  | Field Type | Notes                                                   |
	 * |-----------|-------------|------------|---------------------------------------------------------|
	 * | 0x38      | Player name | String     | Supports chat colouring, limited to 16 characters.      |
	 * | 0x38      | Online      | Bool       | The client will remove the user from the list if false. |
	 * | 0x38      | Ping        | Short      | Ping, presumably in ms.                                 |
	 * 
	 */
}

// ID=0x39
type PlayerAbilitiesClientPkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes                                 |
	 * |-----------|---------------|------------|---------------------------------------|
	 * | 0x39      | Flags         | Byte       |                                       |
	 * | 0x39      | Flying speed  | Float      | previous integer value divided by 250 |
	 * | 0x39      | Walking speed | Float      | previous integer value divided by 250 |
	 * 
	 */
}

// ID=0x3a
type PlayTabCompleteClientPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                                                                                   |
	 * |-----------|------------|------------|-------------------------------------------------------------------------------------------------------------------------|
	 * | 0x3A      | Count      | VarInt     | Number of following strings                                                                                             |
	 * | 0x3A      | Match      | String     | One eligible command, note that each command is sent separately instead of in a single string, hence the need for Count |
	 * 
	 */
}

// ID=0x3b
type PlayScoreboardObjectivePkt struct {
	/*
	 * | Packet ID | Field Name      | Field Type | Notes                                                                                 |
	 * |-----------|-----------------|------------|---------------------------------------------------------------------------------------|
	 * | 0x3B      | Objective name  | String     | An unique name for the objective                                                      |
	 * | 0x3B      | Objective value | String     | The text to be displayed for the score.                                               |
	 * | 0x3B      | Create/Remove   | Byte       | 0 to create the scoreboard. 1 to remove the scoreboard. 2 to update the display text. |
	 * 
	 */
}

// ID=0x3c
type PlayUpdateScorePkt struct {
	/*
	 * | Packet ID | Field Name    | Field Type | Notes                                                                                            |
	 * |-----------|---------------|------------|--------------------------------------------------------------------------------------------------|
	 * | 0x3C      | Item Name     | String     | An unique name to be displayed in the list.                                                      |
	 * | 0x3C      | Update/Remove | Byte       | 0 to create/update an item. 1 to remove an item.                                                 |
	 * | 0x3C      | Score Name    | String     | The unique name for the scoreboard to be updated. Only sent when Update/Remove does not equal 1. |
	 * | 0x3C      | Value         | Int        | The score to be displayed next to the entry. Only sent when Update/Remove does not equal 1.      |
	 * 
	 */
}

// ID=0x3d
type PlayDisplayScoreboardPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                                 |
	 * |-----------|------------|------------|-----------------------------------------------------------------------|
	 * | 0x3D      | Position   | Byte       | The position of the scoreboard. 0 = list, 1 = sidebar, 2 = belowName. |
	 * | 0x3D      | Score Name | String     | The unique name for the scoreboard to be displayed.                   |
	 * 
	 */
}

// ID=0x3e
type PlayTeamsPkt struct {
	/*
	 * | Packet ID | Field Name        | Field Type       | Notes                                                                                                                                                                                                 |
	 * |-----------|-------------------|------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
	 * | 0x3E      | Team Name         | String           | A unique name for the team. (Shared with scoreboard).                                                                                                                                                 |
	 * | 0x3E      | Mode              | Byte             | If 0 then the team is created.
	 * If 1 then the team is removed. 
	 * If 2 the team team information is updated. 
	 * If 3 then new players are added to the team. 
	 * If 4 then players are removed from the team. |
	 * | 0x3E      | Team Display Name | String           | Only if Mode = 0 or 2.                                                                                                                                                                                |
	 * | 0x3E      | Team Prefix       | String           | Only if Mode = 0 or 2. Displayed before the players' name that are part of this team.                                                                                                                 |
	 * | 0x3E      | Team Suffix       | String           | Only if Mode = 0 or 2. Displayed after the players' name that are part of this team.                                                                                                                  |
	 * | 0x3E      | Friendly fire     | Byte             | Only if Mode = 0 or 2; 0 for off, 1 for on, 3 for seeing friendly invisibles                                                                                                                          |
	 * | 0x3E      | Player count      | Short            | Only if Mode = 0 or 3 or 4. Number of players in the array                                                                                                                                            |
	 * | 0x3E      | Players           | Array of strings | Only if Mode = 0 or 3 or 4. Players to be added/remove from the team.                                                                                                                                 |
	 * 
	 */
}

// ID=0x3f
type PlayPluginMessageClientPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                        |
	 * |-----------|------------|------------|----------------------------------------------|
	 * | 0x3F      | Channel    | String     | Name of the "channel" used to send the data. |
	 * | 0x3F      | Length     | Short      | Length of the following byte array           |
	 * | 0x3F      | Data       | Byte Array | Any data.                                    |
	 * 
	 */
}

// ID=0x40
type PlayDisconnectPkt struct {
	/*
	 * | Packet ID | Field Name | Field Type | Notes                                                                       |
	 * |-----------|------------|------------|-----------------------------------------------------------------------------|
	 * | 0x40      | Reason     | String     | Displayed to the client when the connection terminates. Must be valid JSON. |
	 * 
	 */
}

// ======== END play ========
