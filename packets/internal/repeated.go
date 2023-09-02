
// Generated at 2023-09-01 20:57:33.569 -06:00

package internal

import (
	"io"
	. "github.com/kmcsr/go-liter"
)

// Protocol=763; ProtocolName=1.20; State=login; Bound=client; ID=0x0
// Protocol=762; ProtocolName=1.19.4; State=login; Bound=client; ID=0x0
// Protocol=761; ProtocolName=1.19.3; State=login; Bound=client; ID=0x0
// Protocol=760; ProtocolName=1.19.2; State=login; Bound=client; ID=0x0
// Protocol=759; ProtocolName=1.19; State=login; Bound=client; ID=0x0
// Protocol=758; ProtocolName=1.18.2; State=login; Bound=client; ID=0x0
// Protocol=757; ProtocolName=1.18.1; State=login; Bound=client; ID=0x0
// Protocol=756; ProtocolName=1.17.1; State=login; Bound=client; ID=0x0
// Protocol=755; ProtocolName=1.17; State=login; Bound=client; ID=0x0
// Protocol=754; ProtocolName=1.16.5; State=login; Bound=client; ID=0x0
// Protocol=753; ProtocolName=1.16.3; State=login; Bound=client; ID=0x0
// Protocol=578; ProtocolName=1.15.2; State=login; Bound=client; ID=0x0
// Protocol=498; ProtocolName=1.14.4; State=login; Bound=client; ID=0x0
// Protocol=404; ProtocolName=1.13.2; State=login; Bound=client; ID=0x0
// Protocol=401; ProtocolName=1.13.1; State=login; Bound=client; ID=0x0
// Protocol=340; ProtocolName=1.12.2; State=login; Bound=client; ID=0x0
// Protocol=338; ProtocolName=1.12.1; State=login; Bound=client; ID=0x0
// Protocol=335; ProtocolName=1.12; State=login; Bound=client; ID=0x0
// Protocol=316; ProtocolName=1.11.2; State=login; Bound=client; ID=0x0
// Protocol=315; ProtocolName=1.11; State=login; Bound=client; ID=0x0
// Protocol=210; ProtocolName=1.10.2; State=login; Bound=client; ID=0x0
// Protocol=110; ProtocolName=1.9.4; State=login; Bound=client; ID=0x0
// Protocol=47; ProtocolName=1.8.9; State=login; Bound=client; ID=0x0
type LoginDisconnect_763_0 struct {
	Reason Object // Chat
}

var _ Packet = (*LoginDisconnect_763_0)(nil)

func (p LoginDisconnect_763_0)Encode(b *PacketBuilder){
	b.JSON(p.Reason)
}

func (p *LoginDisconnect_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.Reason); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=login; Bound=client; ID=0x1
// Protocol=762; ProtocolName=1.19.4; State=login; Bound=client; ID=0x1
// Protocol=761; ProtocolName=1.19.3; State=login; Bound=client; ID=0x1
// Protocol=760; ProtocolName=1.19.2; State=login; Bound=client; ID=0x1
// Protocol=759; ProtocolName=1.19; State=login; Bound=client; ID=0x1
// Protocol=758; ProtocolName=1.18.2; State=login; Bound=client; ID=0x1
// Protocol=757; ProtocolName=1.18.1; State=login; Bound=client; ID=0x1
// Protocol=756; ProtocolName=1.17.1; State=login; Bound=client; ID=0x1
// Protocol=755; ProtocolName=1.17; State=login; Bound=client; ID=0x1
// Protocol=754; ProtocolName=1.16.5; State=login; Bound=client; ID=0x1
// Protocol=753; ProtocolName=1.16.3; State=login; Bound=client; ID=0x1
// Protocol=578; ProtocolName=1.15.2; State=login; Bound=client; ID=0x1
// Protocol=498; ProtocolName=1.14.4; State=login; Bound=client; ID=0x1
// Protocol=404; ProtocolName=1.13.2; State=login; Bound=client; ID=0x1
// Protocol=401; ProtocolName=1.13.1; State=login; Bound=client; ID=0x1
// Protocol=340; ProtocolName=1.12.2; State=login; Bound=client; ID=0x1
// Protocol=338; ProtocolName=1.12.1; State=login; Bound=client; ID=0x1
// Protocol=335; ProtocolName=1.12; State=login; Bound=client; ID=0x1
// Protocol=316; ProtocolName=1.11.2; State=login; Bound=client; ID=0x1
// Protocol=315; ProtocolName=1.11; State=login; Bound=client; ID=0x1
// Protocol=210; ProtocolName=1.10.2; State=login; Bound=client; ID=0x1
// Protocol=110; ProtocolName=1.9.4; State=login; Bound=client; ID=0x1
// Protocol=47; ProtocolName=1.8.9; State=login; Bound=client; ID=0x1
type LoginEncryptionRequest_763_0 struct {
	/* Appears to be empty */
	ServerID String // String
	/* Length of Public Key */
	PublicKeyLength VarInt // VarInt
	PublicKey ByteArray // Byte Array
	/* Length of Verify Token */
	VerifyTokenLength VarInt // VarInt
	VerifyToken ByteArray // Byte Array
}

var _ Packet = (*LoginEncryptionRequest_763_0)(nil)

func (p LoginEncryptionRequest_763_0)Encode(b *PacketBuilder){
	b.String(p.ServerID)
	p.PublicKeyLength = len(p.PublicKey)
	b.VarInt(p.PublicKeyLength)
	b.ByteArray(p.PublicKey)
	p.VerifyTokenLength = len(p.VerifyToken)
	b.VarInt(p.VerifyTokenLength)
	b.ByteArray(p.VerifyToken)
}

func (p *LoginEncryptionRequest_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ServerID, ok = r.String(); !ok {
		return io.EOF
	}
	if p.PublicKeyLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.PublicKey = make(ByteArray, p.PublicKeyLength)
	if ok = r.ByteArray(p.PublicKey); !ok {
		return io.EOF
	}
	if p.VerifyTokenLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.VerifyToken = make(ByteArray, p.VerifyTokenLength)
	if ok = r.ByteArray(p.VerifyToken); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=login; Bound=server; ID=0x1
// Protocol=762; ProtocolName=1.19.4; State=login; Bound=server; ID=0x1
// Protocol=761; ProtocolName=1.19.3; State=login; Bound=server; ID=0x1
// Protocol=757; ProtocolName=1.18.1; State=login; Bound=server; ID=0x1
// Protocol=753; ProtocolName=1.16.3; State=login; Bound=server; ID=0x1
// Protocol=578; ProtocolName=1.15.2; State=login; Bound=server; ID=0x1
// Protocol=340; ProtocolName=1.12.2; State=login; Bound=server; ID=0x1
// Protocol=338; ProtocolName=1.12.1; State=login; Bound=server; ID=0x1
// Protocol=316; ProtocolName=1.11.2; State=login; Bound=server; ID=0x1
// Protocol=315; ProtocolName=1.11; State=login; Bound=server; ID=0x1
// Protocol=47; ProtocolName=1.8.9; State=login; Bound=server; ID=0x1
type LoginEncryptionResponse_763_0 struct {
	/* Length of Shared Secret */
	SharedSecretLength VarInt // VarInt
	SharedSecret ByteArray // Byte Array
	/* Length of Verify Token */
	VerifyTokenLength VarInt // VarInt
	VerifyToken ByteArray // Byte Array
}

var _ Packet = (*LoginEncryptionResponse_763_0)(nil)

func (p LoginEncryptionResponse_763_0)Encode(b *PacketBuilder){
	p.SharedSecretLength = len(p.SharedSecret)
	b.VarInt(p.SharedSecretLength)
	b.ByteArray(p.SharedSecret)
	p.VerifyTokenLength = len(p.VerifyToken)
	b.VarInt(p.VerifyTokenLength)
	b.ByteArray(p.VerifyToken)
}

func (p *LoginEncryptionResponse_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SharedSecretLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.SharedSecret = make(ByteArray, p.SharedSecretLength)
	if ok = r.ByteArray(p.SharedSecret); !ok {
		return io.EOF
	}
	if p.VerifyTokenLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.VerifyToken = make(ByteArray, p.VerifyTokenLength)
	if ok = r.ByteArray(p.VerifyToken); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=login; Bound=server; ID=0x1
// Protocol=759; ProtocolName=1.19; State=login; Bound=server; ID=0x1
type LoginEncryptionResponse_760_1 struct {
	/* Length of Shared Secret. */
	SharedSecretLength VarInt // VarInt
	/* Shared Secret value, encrypted with the server's public key. */
	SharedSecret ByteArray // Byte Array
	/* Whether or not the Verify Token should be sent. If not, then the salt and signature will be sent. */
	HasVerifyToken Bool // Boolean
	/* Length of Verify Token. Optional and only sent if Has Verify Token is true. */
	OptionalVerifyTokenLength VarInt // VarInt
	/* Verify Token value, encrypted with the same public key as the shared secret. Optional and only sent if Has Verify Token is true. */
	OptionalVerifyToken ByteArray // Byte Array
	/* Cryptography, used for validating the message signature. Optional and only sent if Has Verify Token is false. */
	OptionalSalt Long // Long
	/* Array Length. Optional and only sent if Has Verify Token is false. */
	OptionalMessageSignatureLength VarInt // VarInt
	/* The bytes of the public key signature the client received from Mojang. Optional and only sent if Has Verify Token is false. */
	OptionalMessageSignature ByteArray // Byte Array
}

var _ Packet = (*LoginEncryptionResponse_760_1)(nil)

func (p LoginEncryptionResponse_760_1)Encode(b *PacketBuilder){
	p.SharedSecretLength = len(p.SharedSecret)
	b.VarInt(p.SharedSecretLength)
	b.ByteArray(p.SharedSecret)
	b.Bool(p.HasVerifyToken)
	p.OptionalVerifyTokenLength = len(p.OptionalVerifyToken)
	b.VarInt(p.OptionalVerifyTokenLength)
	b.ByteArray(p.OptionalVerifyToken)
	b.Long(p.OptionalSalt)
	p.OptionalMessageSignatureLength = len(p.OptionalMessageSignature)
	b.VarInt(p.OptionalMessageSignatureLength)
	b.ByteArray(p.OptionalMessageSignature)
}

func (p *LoginEncryptionResponse_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SharedSecretLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.SharedSecret = make(ByteArray, p.SharedSecretLength)
	if ok = r.ByteArray(p.SharedSecret); !ok {
		return io.EOF
	}
	if p.HasVerifyToken, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.OptionalVerifyTokenLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.OptionalVerifyToken = make(ByteArray, p.OptionalVerifyTokenLength)
	if ok = r.ByteArray(p.OptionalVerifyToken); !ok {
		return io.EOF
	}
	if p.OptionalSalt, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.OptionalMessageSignatureLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.OptionalMessageSignature = make(ByteArray, p.OptionalMessageSignatureLength)
	if ok = r.ByteArray(p.OptionalMessageSignature); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=login; Bound=server; ID=0x1
// Protocol=756; ProtocolName=1.17.1; State=login; Bound=server; ID=0x1
// Protocol=755; ProtocolName=1.17; State=login; Bound=server; ID=0x1
// Protocol=754; ProtocolName=1.16.5; State=login; Bound=server; ID=0x1
// Protocol=498; ProtocolName=1.14.4; State=login; Bound=server; ID=0x1
// Protocol=404; ProtocolName=1.13.2; State=login; Bound=server; ID=0x1
// Protocol=401; ProtocolName=1.13.1; State=login; Bound=server; ID=0x1
// Protocol=335; ProtocolName=1.12; State=login; Bound=server; ID=0x1
// Protocol=210; ProtocolName=1.10.2; State=login; Bound=server; ID=0x1
// Protocol=110; ProtocolName=1.9.4; State=login; Bound=server; ID=0x1
type LoginEncryptionResponse_758_2 struct {
	/* Length of Shared Secret */
	SharedSecretLength VarInt // VarInt
	SharedSecret ByteArray // Byte Array
	/* Length of Verify Token */
	VerifyTokenLength VarInt // VarInt
	VerifyToken ByteArray // Byte Array
}

var _ Packet = (*LoginEncryptionResponse_758_2)(nil)

func (p LoginEncryptionResponse_758_2)Encode(b *PacketBuilder){
	p.SharedSecretLength = len(p.SharedSecret)
	b.VarInt(p.SharedSecretLength)
	b.ByteArray(p.SharedSecret)
	p.VerifyTokenLength = len(p.VerifyToken)
	b.VarInt(p.VerifyTokenLength)
	b.ByteArray(p.VerifyToken)
}

func (p *LoginEncryptionResponse_758_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SharedSecretLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.SharedSecret = make(ByteArray, p.SharedSecretLength)
	if ok = r.ByteArray(p.SharedSecret); !ok {
		return io.EOF
	}
	if p.VerifyTokenLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.VerifyToken = make(ByteArray, p.VerifyTokenLength)
	if ok = r.ByteArray(p.VerifyToken); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=login; Bound=client; ID=0x4
// Protocol=762; ProtocolName=1.19.4; State=login; Bound=client; ID=0x4
// Protocol=761; ProtocolName=1.19.3; State=login; Bound=client; ID=0x4
// Protocol=760; ProtocolName=1.19.2; State=login; Bound=client; ID=0x4
// Protocol=759; ProtocolName=1.19; State=login; Bound=client; ID=0x4
// Protocol=758; ProtocolName=1.18.2; State=login; Bound=client; ID=0x4
// Protocol=757; ProtocolName=1.18.1; State=login; Bound=client; ID=0x4
// Protocol=756; ProtocolName=1.17.1; State=login; Bound=client; ID=0x4
// Protocol=755; ProtocolName=1.17; State=login; Bound=client; ID=0x4
// Protocol=754; ProtocolName=1.16.5; State=login; Bound=client; ID=0x4
// Protocol=753; ProtocolName=1.16.3; State=login; Bound=client; ID=0x4
// Protocol=578; ProtocolName=1.15.2; State=login; Bound=client; ID=0x4
// Protocol=498; ProtocolName=1.14.4; State=login; Bound=client; ID=0x4
// Protocol=404; ProtocolName=1.13.2; State=login; Bound=client; ID=0x4
// Protocol=401; ProtocolName=1.13.1; State=login; Bound=client; ID=0x4
type LoginLoginPluginRequest_763_0 struct {
	/* Generated by the server - should be unique to the connection. */
	MessageID VarInt // VarInt
	/* Name of the plugin channel used to send the data */
	Channel Identifier // Identifier
	/* Any data, depending on the channel. The length of this array must be inferred from the packet length. */
	Data ByteArray // Byte Array
}

var _ Packet = (*LoginLoginPluginRequest_763_0)(nil)

func (p LoginLoginPluginRequest_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.MessageID)
	b.Encode(p.Channel)
	b.ReadAll(p.Data)
}

func (p *LoginLoginPluginRequest_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.MessageID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Channel.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=login; Bound=server; ID=0x2
// Protocol=762; ProtocolName=1.19.4; State=login; Bound=server; ID=0x2
// Protocol=761; ProtocolName=1.19.3; State=login; Bound=server; ID=0x2
// Protocol=760; ProtocolName=1.19.2; State=login; Bound=server; ID=0x2
// Protocol=759; ProtocolName=1.19; State=login; Bound=server; ID=0x2
// Protocol=758; ProtocolName=1.18.2; State=login; Bound=server; ID=0x2
// Protocol=757; ProtocolName=1.18.1; State=login; Bound=server; ID=0x2
// Protocol=756; ProtocolName=1.17.1; State=login; Bound=server; ID=0x2
// Protocol=755; ProtocolName=1.17; State=login; Bound=server; ID=0x2
// Protocol=754; ProtocolName=1.16.5; State=login; Bound=server; ID=0x2
// Protocol=753; ProtocolName=1.16.3; State=login; Bound=server; ID=0x2
// Protocol=578; ProtocolName=1.15.2; State=login; Bound=server; ID=0x2
// Protocol=498; ProtocolName=1.14.4; State=login; Bound=server; ID=0x2
// Protocol=404; ProtocolName=1.13.2; State=login; Bound=server; ID=0x2
// Protocol=401; ProtocolName=1.13.1; State=login; Bound=server; ID=0x2
type LoginLoginPluginResponse_763_0 struct {
	/* Should match ID from server. */
	MessageID VarInt // VarInt
	/* true if the client understands the request, false otherwise. When false, no payload follows. */
	Successful Bool // Boolean
	/* Any data, depending on the channel. The length of this array must be inferred from the packet length. */
	Data ByteArray // Optional Byte Array
}

var _ Packet = (*LoginLoginPluginResponse_763_0)(nil)

func (p LoginLoginPluginResponse_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.MessageID)
	b.Bool(p.Successful)
	b.ReadAll(p.Data)
}

func (p *LoginLoginPluginResponse_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.MessageID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Successful, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=login; Bound=server; ID=0x0
// Protocol=762; ProtocolName=1.19.4; State=login; Bound=server; ID=0x0
// Protocol=761; ProtocolName=1.19.3; State=login; Bound=server; ID=0x0
type LoginLoginStart_763_0 struct {
	/* Player's Username. */
	Name String // String (16)
	/* Whether or not the next field should be sent. */
	HasPlayerUUID Bool // Boolean
	/* The UUID of the player logging in. Only sent if Has Player UUID is true. */
	PlayerUUID UUID // Optional UUID
}

var _ Packet = (*LoginLoginStart_763_0)(nil)

func (p LoginLoginStart_763_0)Encode(b *PacketBuilder){
	b.String(p.Name)
	b.Bool(p.HasPlayerUUID)
	b.UUID(p.PlayerUUID)
}

func (p *LoginLoginStart_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Name, ok = r.String(); !ok {
		return io.EOF
	}
	if p.HasPlayerUUID, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.PlayerUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=login; Bound=server; ID=0x0
// Protocol=759; ProtocolName=1.19; State=login; Bound=server; ID=0x0
type LoginLoginStart_760_1 struct {
	/* Player's Username. */
	Name String // String (16)
	/* Whether or not the next 5 fields should be sent. */
	HasSigData Bool // Boolean
	/* When the key data will expire. Optional. Only sent if Has Sig Data is true. */
	Timestamp Long // Long
	/* Length of Public Key. Optional. Only sent if Has Sig Data is true. */
	PublicKeyLength VarInt // VarInt
	/* The encoded bytes of the public key the client received from Mojang. Optional. Only sent if Has Sig Data is true. */
	PublicKey ByteArray // Byte Array
	/* Length of Signature. Optional. Only sent if Has Sig Data is true. */
	SignatureLength VarInt // VarInt
	/* The bytes of the public key signature the client received from Mojang. Optional. Only sent if Has Sig Data is true. */
	Signature ByteArray // Byte Array
	/* Whether or not the next field should be sent. */
	HasPlayerUUID Bool // Boolean
	/* The UUID of the player logging in. Optional. Only sent if Has Player UUID is true. */
	PlayerUUID UUID // UUID
}

var _ Packet = (*LoginLoginStart_760_1)(nil)

func (p LoginLoginStart_760_1)Encode(b *PacketBuilder){
	b.String(p.Name)
	b.Bool(p.HasSigData)
	b.Long(p.Timestamp)
	p.PublicKeyLength = len(p.PublicKey)
	b.VarInt(p.PublicKeyLength)
	b.ByteArray(p.PublicKey)
	p.SignatureLength = len(p.Signature)
	b.VarInt(p.SignatureLength)
	b.ByteArray(p.Signature)
	b.Bool(p.HasPlayerUUID)
	b.UUID(p.PlayerUUID)
}

func (p *LoginLoginStart_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Name, ok = r.String(); !ok {
		return io.EOF
	}
	if p.HasSigData, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Timestamp, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.PublicKeyLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.PublicKey = make(ByteArray, p.PublicKeyLength)
	if ok = r.ByteArray(p.PublicKey); !ok {
		return io.EOF
	}
	if p.SignatureLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Signature = make(ByteArray, p.SignatureLength)
	if ok = r.ByteArray(p.Signature); !ok {
		return io.EOF
	}
	if p.HasPlayerUUID, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.PlayerUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=login; Bound=server; ID=0x0
// Protocol=757; ProtocolName=1.18.1; State=login; Bound=server; ID=0x0
// Protocol=756; ProtocolName=1.17.1; State=login; Bound=server; ID=0x0
// Protocol=755; ProtocolName=1.17; State=login; Bound=server; ID=0x0
// Protocol=754; ProtocolName=1.16.5; State=login; Bound=server; ID=0x0
// Protocol=753; ProtocolName=1.16.3; State=login; Bound=server; ID=0x0
// Protocol=578; ProtocolName=1.15.2; State=login; Bound=server; ID=0x0
// Protocol=498; ProtocolName=1.14.4; State=login; Bound=server; ID=0x0
// Protocol=404; ProtocolName=1.13.2; State=login; Bound=server; ID=0x0
// Protocol=401; ProtocolName=1.13.1; State=login; Bound=server; ID=0x0
// Protocol=340; ProtocolName=1.12.2; State=login; Bound=server; ID=0x0
// Protocol=338; ProtocolName=1.12.1; State=login; Bound=server; ID=0x0
// Protocol=335; ProtocolName=1.12; State=login; Bound=server; ID=0x0
// Protocol=316; ProtocolName=1.11.2; State=login; Bound=server; ID=0x0
// Protocol=315; ProtocolName=1.11; State=login; Bound=server; ID=0x0
// Protocol=210; ProtocolName=1.10.2; State=login; Bound=server; ID=0x0
// Protocol=110; ProtocolName=1.9.4; State=login; Bound=server; ID=0x0
// Protocol=47; ProtocolName=1.8.9; State=login; Bound=server; ID=0x0
type LoginLoginStart_758_2 struct {
	Name String // String
}

var _ Packet = (*LoginLoginStart_758_2)(nil)

func (p LoginLoginStart_758_2)Encode(b *PacketBuilder){
	b.String(p.Name)
}

func (p *LoginLoginStart_758_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Name, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=login; Bound=client; ID=0x2
// Protocol=757; ProtocolName=1.18.1; State=login; Bound=client; ID=0x2
// Protocol=756; ProtocolName=1.17.1; State=login; Bound=client; ID=0x2
// Protocol=755; ProtocolName=1.17; State=login; Bound=client; ID=0x2
// Protocol=754; ProtocolName=1.16.5; State=login; Bound=client; ID=0x2
// Protocol=753; ProtocolName=1.16.3; State=login; Bound=client; ID=0x2
type LoginLoginSuccess_758_0 struct {
	UUID UUID // UUID
	Username String // String (16)
}

var _ Packet = (*LoginLoginSuccess_758_0)(nil)

func (p LoginLoginSuccess_758_0)Encode(b *PacketBuilder){
	b.UUID(p.UUID)
	b.String(p.Username)
}

func (p *LoginLoginSuccess_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.UUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Username, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=578; ProtocolName=1.15.2; State=login; Bound=client; ID=0x2
// Protocol=498; ProtocolName=1.14.4; State=login; Bound=client; ID=0x2
// Protocol=404; ProtocolName=1.13.2; State=login; Bound=client; ID=0x2
// Protocol=401; ProtocolName=1.13.1; State=login; Bound=client; ID=0x2
// Protocol=340; ProtocolName=1.12.2; State=login; Bound=client; ID=0x2
// Protocol=338; ProtocolName=1.12.1; State=login; Bound=client; ID=0x2
// Protocol=335; ProtocolName=1.12; State=login; Bound=client; ID=0x2
// Protocol=316; ProtocolName=1.11.2; State=login; Bound=client; ID=0x2
// Protocol=315; ProtocolName=1.11; State=login; Bound=client; ID=0x2
// Protocol=210; ProtocolName=1.10.2; State=login; Bound=client; ID=0x2
// Protocol=110; ProtocolName=1.9.4; State=login; Bound=client; ID=0x2
// Protocol=47; ProtocolName=1.8.9; State=login; Bound=client; ID=0x2
type LoginLoginSuccess_578_1 struct {
	/* Unlike in other packets, this field contains the UUID as a string with hyphens. */
	UUID String // String
	Username String // String
}

var _ Packet = (*LoginLoginSuccess_578_1)(nil)

func (p LoginLoginSuccess_578_1)Encode(b *PacketBuilder){
	b.String(p.UUID)
	b.String(p.Username)
}

func (p *LoginLoginSuccess_578_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.UUID, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Username, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=login; Bound=client; ID=0x3
// Protocol=762; ProtocolName=1.19.4; State=login; Bound=client; ID=0x3
// Protocol=761; ProtocolName=1.19.3; State=login; Bound=client; ID=0x3
// Protocol=760; ProtocolName=1.19.2; State=login; Bound=client; ID=0x3
// Protocol=759; ProtocolName=1.19; State=login; Bound=client; ID=0x3
// Protocol=758; ProtocolName=1.18.2; State=login; Bound=client; ID=0x3
// Protocol=757; ProtocolName=1.18.1; State=login; Bound=client; ID=0x3
// Protocol=756; ProtocolName=1.17.1; State=login; Bound=client; ID=0x3
// Protocol=755; ProtocolName=1.17; State=login; Bound=client; ID=0x3
// Protocol=754; ProtocolName=1.16.5; State=login; Bound=client; ID=0x3
// Protocol=753; ProtocolName=1.16.3; State=login; Bound=client; ID=0x3
// Protocol=578; ProtocolName=1.15.2; State=login; Bound=client; ID=0x3
// Protocol=498; ProtocolName=1.14.4; State=login; Bound=client; ID=0x3
// Protocol=404; ProtocolName=1.13.2; State=login; Bound=client; ID=0x3
// Protocol=401; ProtocolName=1.13.1; State=login; Bound=client; ID=0x3
// Protocol=340; ProtocolName=1.12.2; State=login; Bound=client; ID=0x3
// Protocol=338; ProtocolName=1.12.1; State=login; Bound=client; ID=0x3
// Protocol=335; ProtocolName=1.12; State=login; Bound=client; ID=0x3
// Protocol=316; ProtocolName=1.11.2; State=login; Bound=client; ID=0x3
// Protocol=315; ProtocolName=1.11; State=login; Bound=client; ID=0x3
// Protocol=210; ProtocolName=1.10.2; State=login; Bound=client; ID=0x3
// Protocol=110; ProtocolName=1.9.4; State=login; Bound=client; ID=0x3
// Protocol=47; ProtocolName=1.8.9; State=login; Bound=client; ID=0x3
type LoginSetCompression_763_0 struct {
	/* Maximum size of a packet before its compressed */
	Threshold VarInt // VarInt
}

var _ Packet = (*LoginSetCompression_763_0)(nil)

func (p LoginSetCompression_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Threshold)
}

func (p *LoginSetCompression_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Threshold, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x6
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x6
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x5
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x5
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x5
type PlayAcknowledgeBlockChange_763_0 struct {
	/* Represents the sequence to acknowledge, this is used for properly syncing block changes to the client after interactions. */
	SequenceID VarInt // VarInt
}

var _ Packet = (*PlayAcknowledgeBlockChange_763_0)(nil)

func (p PlayAcknowledgeBlockChange_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.SequenceID)
}

func (p *PlayAcknowledgeBlockChange_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SequenceID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x8
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x8
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x8
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x8
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x7
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x7
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x8
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x5c
type PlayAcknowledgePlayerDigging_758_0 struct {
	/* Position where the digging was happening */
	Location Position // Position
	/* Block state ID of the block that should be at that position now. */
	Block VarInt // VarInt
	/* Same as Player Digging.  Only Started digging (0), Cancelled digging (1), and Finished digging (2) are used. */
	Status VarInt // VarInt enum
	/* True if the digging succeeded; false if the client should undo any changes it made locally.  (How does this work?) */
	Successful Bool // Boolean
}

var _ Packet = (*PlayAcknowledgePlayerDigging_758_0)(nil)

func (p PlayAcknowledgePlayerDigging_758_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.Block)
	b.VarInt(p.Status)
	b.Bool(p.Successful)
}

func (p *PlayAcknowledgePlayerDigging_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Block, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Status, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Successful, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x41
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x41
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x41
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x41
type PlayActionBar_758_0 struct {
	/* Displays a message above the hotbar (the same as position 2 in Chat Message (clientbound). */
	ActionBarText Object // Chat
}

var _ Packet = (*PlayActionBar_758_0)(nil)

func (p PlayActionBar_758_0)Encode(b *PacketBuilder){
	b.JSON(p.ActionBarText)
}

func (p *PlayActionBar_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.ActionBarText); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x22
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x22
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x22
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x22
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x22
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x22
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x20
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x20
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x1e
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x1e
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x19
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x19
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x19
type PlayAdvancementTab_758_0 struct {
	/* 0: Opened tab, 1: Closed screen */
	Action VarInt // VarInt enum
	/* Only present if action is Opened tab */
	TabID Identifier // Optional identifier
}

var _ Packet = (*PlayAdvancementTab_758_0)(nil)

func (p PlayAdvancementTab_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.Action)
	b.Encode(p.TabID)
}

func (p *PlayAdvancementTab_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Action, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.TabID.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x2c
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x2c
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x2c
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x2c
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x2c
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x2c
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x2a
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x2a
type PlayAnimation_758_0 struct {
	/* Hand used for the animation. 0: main hand, 1: off hand. */
	Hand VarInt // VarInt Enum
}

var _ Packet = (*PlayAnimation_758_0)(nil)

func (p PlayAnimation_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
}

func (p *PlayAnimation_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x6
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x6
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x6
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x6
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x6
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0xb
type PlayAnimation_404_1 struct {
	/* Player ID */
	EntityID VarInt // VarInt
	/* Animation ID (see below) */
	Animation UByte // Unsigned Byte
}

var _ Packet = (*PlayAnimation_404_1)(nil)

func (p PlayAnimation_404_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UByte(p.Animation)
}

func (p *PlayAnimation_404_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Animation, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x27
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x27
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x1d
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x1d
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x1a
type PlayAnimation_404_2 struct {
	/* Hand used for the animation. 0: main hand, 1: off hand. */
	Hand VarInt // VarInt Enum
}

var _ Packet = (*PlayAnimation_404_2)(nil)

func (p PlayAnimation_404_2)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
}

func (p *PlayAnimation_404_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x6
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x6
type PlayAnimation_335_3 struct {
	/* Player ID */
	EntityID VarInt // VarInt
	/* Animation ID (see below) */
	Animation UByte // Unsigned Byte
}

var _ Packet = (*PlayAnimation_335_3)(nil)

func (p PlayAnimation_335_3)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UByte(p.Animation)
}

func (p *PlayAnimation_335_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Animation, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x1d
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x1a
type PlayAnimation_335_4 struct {
	/* Hand used for the animation. 0: main hand, 1: off hand. */
	Hand VarInt // VarInt Enum
}

var _ Packet = (*PlayAnimation_335_4)(nil)

func (p PlayAnimation_335_4)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
}

func (p *PlayAnimation_335_4)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x6
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x6
type PlayAnimation_210_5 struct {
	/* Player ID */
	EntityID VarInt // VarInt
	/* Animation ID (see below) */
	Animation UByte // Unsigned Byte
}

var _ Packet = (*PlayAnimation_210_5)(nil)

func (p PlayAnimation_210_5)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UByte(p.Animation)
}

func (p *PlayAnimation_210_5)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Animation, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x1a
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x1a
type PlayAnimation_210_6 struct {
	/* Hand used for the animation. 0: main hand, 1: off hand. */
	Hand VarInt // VarInt Enum
}

var _ Packet = (*PlayAnimation_210_6)(nil)

func (p PlayAnimation_210_6)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
}

func (p *PlayAnimation_210_6)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x4e
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x4e
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x4e
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x4e
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x45
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x45
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x45
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x44
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x40
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x40
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x3d
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x3d
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x3c
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x3a
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x3a
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x3a
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x3a
type PlayAttachEntity_758_0 struct {
	/* Attached entity's EID */
	AttachedEntityID Int // Int
	/* ID of the entity holding the lead. Set to -1 to detach. */
	HoldingEntityID Int // Int
}

var _ Packet = (*PlayAttachEntity_758_0)(nil)

func (p PlayAttachEntity_758_0)Encode(b *PacketBuilder){
	b.Int(p.AttachedEntityID)
	b.Int(p.HoldingEntityID)
}

func (p *PlayAttachEntity_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.AttachedEntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.HoldingEntityID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x9
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x9
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x8
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x8
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x8
type PlayBlockAction_763_0 struct {
	/* Block coordinates. */
	Location Position // Position
	/* Varies depending on block — see Block Actions. */
	ActionID UByte // Unsigned Byte
	/* Varies depending on block — see Block Actions. */
	ActionParameter UByte // Unsigned Byte
	/* The block type ID for the block.  This must match the block at the given coordinates. */
	BlockType VarInt // VarInt
}

var _ Packet = (*PlayBlockAction_763_0)(nil)

func (p PlayBlockAction_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.UByte(p.ActionID)
	b.UByte(p.ActionParameter)
	b.VarInt(p.BlockType)
}

func (p *PlayBlockAction_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.ActionID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.ActionParameter, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.BlockType, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0xb
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0xb
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0xb
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0xb
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0xa
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0xa
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0xb
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0xa
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0xa
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0xa
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0xa
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0xa
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0xa
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0xa
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0xa
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0xa
type PlayBlockAction_758_1 struct {
	/* Block coordinates */
	Location Position // Position
	/* Varies depending on block — see Block Actions */
	ActionID UByte // Unsigned Byte
	/* Varies depending on block — see Block Actions */
	ActionParam UByte // Unsigned Byte
	/* The block type ID for the block, not including metadata/damage value.  This must match the block at the given coordinates. */
	BlockType VarInt // VarInt
}

var _ Packet = (*PlayBlockAction_758_1)(nil)

func (p PlayBlockAction_758_1)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.UByte(p.ActionID)
	b.UByte(p.ActionParam)
	b.VarInt(p.BlockType)
}

func (p *PlayBlockAction_758_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.ActionID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.ActionParam, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.BlockType, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0xa
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x24
type PlayBlockAction_110_2 struct {
	/* Block coordinates */
	Location Position // Position
	/* Varies depending on block — see Block Actions */
	Byte1 UByte // Unsigned Byte
	/* Varies depending on block — see Block Actions */
	Byte2 UByte // Unsigned Byte
	/* The block type ID for the block, not including metadata/damage value */
	BlockType VarInt // VarInt
}

var _ Packet = (*PlayBlockAction_110_2)(nil)

func (p PlayBlockAction_110_2)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.UByte(p.Byte1)
	b.UByte(p.Byte2)
	b.VarInt(p.BlockType)
}

func (p *PlayBlockAction_110_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Byte1, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Byte2, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.BlockType, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x9
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x9
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x9
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x9
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x8
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x8
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x9
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x8
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x8
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x8
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x8
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x8
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x8
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x8
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x8
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x8
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x8
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x25
type PlayBlockBreakAnimation_758_0 struct {
	/* EID for the animation */
	EntityID VarInt // VarInt
	/* Block Position */
	Location Position // Position
	/* 0–9 to set it, any other value to remove it */
	DestroyStage Byte // Byte
}

var _ Packet = (*PlayBlockBreakAnimation_758_0)(nil)

func (p PlayBlockBreakAnimation_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.Location)
	b.Byte(p.DestroyStage)
}

func (p *PlayBlockBreakAnimation_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.DestroyStage, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0xc
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0xc
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0xc
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0xc
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0xb
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0xb
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0xc
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0xb
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0xb
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0xb
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0xb
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0xb
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0xb
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0xb
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0xb
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0xb
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0xb
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x23
type PlayBlockChange_758_0 struct {
	/* Block Coordinates */
	Location Position // Position
	/* The new block state ID for the block as given in the global palette (When reading data: type = id >> 4, meta = id & 15, when writing data: id = type << 4 | (meta & 15)) */
	BlockID VarInt // VarInt
}

var _ Packet = (*PlayBlockChange_758_0)(nil)

func (p PlayBlockChange_758_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.BlockID)
}

func (p *PlayBlockChange_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.BlockID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x8
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x8
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x7
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x7
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x7
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0xa
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0xa
type PlayBlockEntityData_763_0 struct {
	Location Position // Position
	/* The type of the block entity */
	Type VarInt // VarInt
	/* Data to set.  May be a TAG_END (0), in which case the block entity at the given location is removed (though this is not required since the client will remove the block entity automatically on chunk unload or block removal). */
	NBTData NBT // NBT Tag
}

var _ Packet = (*PlayBlockEntityData_763_0)(nil)

func (p PlayBlockEntityData_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.Type)
	b.Encode(p.NBTData)
}

func (p *PlayBlockEntityData_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.NBTData.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0xa
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0xa
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x9
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x9
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0xa
type PlayBlockEntityData_756_1 struct {
	Location Position // Position
	/* The type of update to perform, see below */
	Action UByte // Unsigned Byte
	/* Data to set.  May be a TAG_END (0), in which case the block entity at the given location is removed (though this is not required since the client will remove the block entity automatically on chunk unload or block removal) */
	NBTData NBT // NBT Tag
}

var _ Packet = (*PlayBlockEntityData_756_1)(nil)

func (p PlayBlockEntityData_756_1)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.UByte(p.Action)
	b.Encode(p.NBTData)
}

func (p *PlayBlockEntityData_756_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Action, ok = r.UByte(); !ok {
		return io.EOF
	}
	if err = p.NBTData.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0xa
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0xa
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x9
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x9
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x9
type PlayBlockUpdate_763_0 struct {
	/* Block Coordinates. */
	Location Position // Position
	/* The new block state ID for the block as given in the global palette. See that section for more information. */
	BlockID VarInt // VarInt
}

var _ Packet = (*PlayBlockUpdate_763_0)(nil)

func (p PlayBlockUpdate_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.BlockID)
}

func (p *PlayBlockUpdate_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.BlockID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x0
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x0
type PlayBundleDelimiter_763_0 struct {
}

var _ Packet = (*PlayBundleDelimiter_763_0)(nil)

func (p PlayBundleDelimiter_763_0)Encode(b *PacketBuilder){
}

func (p *PlayBundleDelimiter_763_0)DecodeFrom(r *PacketReader)(err error){ return }

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x47
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x47
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x47
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x47
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x3e
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x3e
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x3f
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x3e
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x3c
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x3c
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x39
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x39
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x38
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x36
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x36
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x36
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x36
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x43
type PlayCamera_758_0 struct {
	/* ID of the entity to set the client's camera to */
	CameraID VarInt // VarInt
}

var _ Packet = (*PlayCamera_758_0)(nil)

func (p PlayCamera_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.CameraID)
}

func (p *PlayCamera_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.CameraID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0xc
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0xc
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0xb
type PlayChangeDifficulty_763_0 struct {
	/* 0: peaceful, 1: easy, 2: normal, 3: hard. */
	Difficulty UByte // Unsigned Byte
	DifficultyLocked Bool // Boolean
}

var _ Packet = (*PlayChangeDifficulty_763_0)(nil)

func (p PlayChangeDifficulty_763_0)Encode(b *PacketBuilder){
	b.UByte(p.Difficulty)
	b.Bool(p.DifficultyLocked)
}

func (p *PlayChangeDifficulty_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Difficulty, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.DifficultyLocked, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x2
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x2
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x2
type PlayChangeDifficulty_763_1 struct {
	/* 0: peaceful, 1: easy, 2: normal, 3: hard . */
	NewDifficulty Byte // Byte
}

var _ Packet = (*PlayChangeDifficulty_763_1)(nil)

func (p PlayChangeDifficulty_763_1)Encode(b *PacketBuilder){
	b.Byte(p.NewDifficulty)
}

func (p *PlayChangeDifficulty_763_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.NewDifficulty, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0xb
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0xb
type PlayChangeDifficulty_760_2 struct {
	/* 0: peaceful, 1: easy, 2: normal, 3: hard. */
	Difficulty UByte // Unsigned Byte
	DifficultyLocked Bool // Boolean
}

var _ Packet = (*PlayChangeDifficulty_760_2)(nil)

func (p PlayChangeDifficulty_760_2)Encode(b *PacketBuilder){
	b.UByte(p.Difficulty)
	b.Bool(p.DifficultyLocked)
}

func (p *PlayChangeDifficulty_760_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Difficulty, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.DifficultyLocked, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x2
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x2
type PlayChangeDifficulty_760_3 struct {
	/* 0: peaceful, 1: easy, 2: normal, 3: hard . */
	NewDifficulty Byte // Byte
}

var _ Packet = (*PlayChangeDifficulty_760_3)(nil)

func (p PlayChangeDifficulty_760_3)Encode(b *PacketBuilder){
	b.Byte(p.NewDifficulty)
}

func (p *PlayChangeDifficulty_760_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.NewDifficulty, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x1e
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x1e
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x1e
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x1e
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x1d
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x1d
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x1f
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x1e
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x20
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x20
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x1e
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x1e
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x1e
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x1e
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x1e
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x1e
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x1e
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x2b
type PlayChangeGameState_758_0 struct {
	/* See below */
	Reason UByte // Unsigned Byte
	/* Depends on Reason */
	Value Float // Float
}

var _ Packet = (*PlayChangeGameState_758_0)(nil)

func (p PlayChangeGameState_758_0)Encode(b *PacketBuilder){
	b.UByte(p.Reason)
	b.Float(p.Value)
}

func (p *PlayChangeGameState_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Reason, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Value, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x21
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x21
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x21
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x21
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x20
type PlayChangeRecipeBookSettings_763_0 struct {
	/* 0: crafting, 1: furnace, 2: blast furnace, 3: smoker. */
	BookID VarInt // VarInt Enum
	BookOpen Bool // Boolean
	FilterActive Bool // Boolean
}

var _ Packet = (*PlayChangeRecipeBookSettings_763_0)(nil)

func (p PlayChangeRecipeBookSettings_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.BookID)
	b.Bool(p.BookOpen)
	b.Bool(p.FilterActive)
}

func (p *PlayChangeRecipeBookSettings_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.BookID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.BookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.FilterActive, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x5
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x5
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x5
type PlayChatMessage_763_0 struct {
	Message String // String (256 chars)
	Timestamp Long // Long
	/* The salt used to verify the signature hash. */
	Salt Long // Long
	/* Whether the next field is present. */
	HasSignature Bool // Boolean
	/* The signature used to verify the chat message's authentication. When present, always 256 bytes. */
	Signature ByteArray // Optional Byte Array
	MessageCount VarInt // VarInt
	Acknowledged BitSet // BitSet (20)
}

var _ Packet = (*PlayChatMessage_763_0)(nil)

func (p PlayChatMessage_763_0)Encode(b *PacketBuilder){
	b.String(p.Message)
	b.Long(p.Timestamp)
	b.Long(p.Salt)
	b.Bool(p.HasSignature)
	b.ByteArray(p.Signature)
	b.VarInt(p.MessageCount)
	b.Encode(p.Acknowledged)
}

func (p *PlayChatMessage_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Timestamp, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.Salt, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.HasSignature, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = p.Signature.DecodeFrom(r); err != nil {
		return
	}
	if p.MessageCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Acknowledged.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0xf
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0xf
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0xe
type PlayChatMessage_758_2 struct {
	/* Limited to 32767 bytes */
	JSONData Object // Chat
	/* 0: chat (chat box), 1: system message (chat box), 2: game info (above hotbar). */
	Position Byte // Byte
	/* Used by the Notchian client for the disableChat launch option. Setting both longs to 0 will always display the message regardless of the setting. */
	Sender UUID // UUID
}

var _ Packet = (*PlayChatMessage_758_2)(nil)

func (p PlayChatMessage_758_2)Encode(b *PacketBuilder){
	b.JSON(p.JSONData)
	b.Byte(p.Position)
	b.UUID(p.Sender)
}

func (p *PlayChatMessage_758_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.JSONData); err != nil {
		return
	}
	if p.Position, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Sender, ok = r.UUID(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x3
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x3
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x3
type PlayChatMessage_758_3 struct {
	/* The client sends the raw input, not a Chat component */
	Message String // String (256)
}

var _ Packet = (*PlayChatMessage_758_3)(nil)

func (p PlayChatMessage_758_3)Encode(b *PacketBuilder){
	b.String(p.Message)
}

func (p *PlayChatMessage_758_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0xf
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0xf
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0xe
type PlayChatMessage_756_4 struct {
	/* Limited to 262144 bytes. */
	JSONData Object // Chat
	/* 0: chat (chat box), 1: system message (chat box), 2: game info (above hotbar). */
	Position Byte // Byte
	/* Used by the Notchian client for the disableChat launch option. Setting both longs to 0 will always display the message regardless of the setting. */
	Sender UUID // UUID
}

var _ Packet = (*PlayChatMessage_756_4)(nil)

func (p PlayChatMessage_756_4)Encode(b *PacketBuilder){
	b.JSON(p.JSONData)
	b.Byte(p.Position)
	b.UUID(p.Sender)
}

func (p *PlayChatMessage_756_4)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.JSONData); err != nil {
		return
	}
	if p.Position, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Sender, ok = r.UUID(); !ok {
		return io.EOF
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x3
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x3
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x3
type PlayChatMessage_756_5 struct {
	/* The client sends the raw input, not a Chat component. */
	Message String // String (256)
}

var _ Packet = (*PlayChatMessage_756_5)(nil)

func (p PlayChatMessage_756_5)Encode(b *PacketBuilder){
	b.String(p.Message)
}

func (p *PlayChatMessage_756_5)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0xf
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0xf
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x2
type PlayChatMessage_578_6 struct {
	/* Limited to 32767 bytes */
	JSONData Object // Chat
	/* 0: chat (chat box), 1: system message (chat box), 2: above hotbar */
	Position Byte // Byte
}

var _ Packet = (*PlayChatMessage_578_6)(nil)

func (p PlayChatMessage_578_6)Encode(b *PacketBuilder){
	b.JSON(p.JSONData)
	b.Byte(p.Position)
}

func (p *PlayChatMessage_578_6)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.JSONData); err != nil {
		return
	}
	if p.Position, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x3
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x2
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x1
type PlayChatMessage_578_7 struct {
	/* The client sends the raw input, not Chat */
	Message String // String
}

var _ Packet = (*PlayChatMessage_578_7)(nil)

func (p PlayChatMessage_578_7)Encode(b *PacketBuilder){
	b.String(p.Message)
}

func (p *PlayChatMessage_578_7)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0xe
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0xf
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0xf
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0xf
type PlayChatMessage_498_8 struct {
	/* Limited to 32767 bytes */
	JSONData Object // Chat
	/* 0: chat (chat box), 1: system message (chat box), 2: game info (above hotbar). */
	Position Byte // Byte
}

var _ Packet = (*PlayChatMessage_498_8)(nil)

func (p PlayChatMessage_498_8)Encode(b *PacketBuilder){
	b.JSON(p.JSONData)
	b.Byte(p.Position)
}

func (p *PlayChatMessage_498_8)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.JSONData); err != nil {
		return
	}
	if p.Position, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x3
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x2
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x2
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x2
type PlayChatMessage_498_9 struct {
	/* The client sends the raw input, not a Chat component */
	Message String // String (256)
}

var _ Packet = (*PlayChatMessage_498_9)(nil)

func (p PlayChatMessage_498_9)Encode(b *PacketBuilder){
	b.String(p.Message)
}

func (p *PlayChatMessage_498_9)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0xe
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0xe
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0xf
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0xf
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0xf
type PlayChatMessage_404_10 struct {
	/* Limited to 32767 bytes */
	JSONData Object // Chat
	/* 0: chat (chat box), 1: system message (chat box), 2: above hotbar */
	Position Byte // Byte
}

var _ Packet = (*PlayChatMessage_404_10)(nil)

func (p PlayChatMessage_404_10)Encode(b *PacketBuilder){
	b.JSON(p.JSONData)
	b.Byte(p.Position)
}

func (p *PlayChatMessage_404_10)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.JSONData); err != nil {
		return
	}
	if p.Position, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x2
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x2
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x3
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x2
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x2
type PlayChatMessage_404_11 struct {
	/* The client sends the raw input, not a Chat component */
	Message String // String
}

var _ Packet = (*PlayChatMessage_404_11)(nil)

func (p PlayChatMessage_404_11)Encode(b *PacketBuilder){
	b.String(p.Message)
}

func (p *PlayChatMessage_404_11)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0xc
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0xc
type PlayChatPreview_760_0 struct {
	QueryID Int // Integer
	ComponentIsPresent Bool // Boolean
	MessageToPreview Object // Component
}

var _ Packet = (*PlayChatPreview_760_0)(nil)

func (p PlayChatPreview_760_0)Encode(b *PacketBuilder){
	b.Int(p.QueryID)
	b.Bool(p.ComponentIsPresent)
	b.JSON(p.MessageToPreview)
}

func (p *PlayChatPreview_760_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.QueryID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.ComponentIsPresent, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.MessageToPreview); err != nil {
		return
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x6
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x5
type PlayChatPreview_760_1 struct {
	Query Int // Integer
	Message String // String (256)
}

var _ Packet = (*PlayChatPreview_760_1)(nil)

func (p PlayChatPreview_760_1)Encode(b *PacketBuilder){
	b.Int(p.Query)
	b.String(p.Message)
}

func (p *PlayChatPreview_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Query, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Message, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x16
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x16
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x14
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x15
type PlayChatSuggestions_763_0 struct {
	/* 0: Add, 1: Remove, 2: Set */
	Action VarInt // VarInt Enum
	/* Number of elements in the following array. */
	Count VarInt // VarInt
	Entries []String // Array of String
}

var _ Packet = (*PlayChatSuggestions_763_0)(nil)

func (p PlayChatSuggestions_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Action)
	p.Count = len(p.Entries)
	b.VarInt(p.Count)
	for _, v := range p.Entries {
		b.String(v)
	}
}

func (p *PlayChatSuggestions_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Action, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Entries = make([]String, p.Count)
	for i, _ := range p.Entries {
		if p.Entries[i], ok = r.String(); !ok {
			return io.EOF
		}
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x22
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x22
type PlayChunkData_756_0 struct {
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkX Int // Int
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkZ Int // Int
	/* Length of the following array */
	BitMaskLength VarInt // VarInt
	/* BitSet with bits (world height in blocks / 16) set to 1 for every 16×16×16 chunk section whose data is included in Data. The least significant bit represents the chunk section at the bottom of the chunk column (from the lowest y to 15 blocks above). */
	PrimaryBitMask []Long // Array of Long
	/* Compound containing one long array named MOTION_BLOCKING, which is a heightmap for the highest solid block at each position in the chunk (as a compacted long array with 256 entries at 9 bits per entry totaling 36 longs). The Notchian server also adds a WORLD_SURFACE long array, the purpose of which is unknown, but it's not required for the chunk to be accepted. */
	Heightmaps NBT // NBT
	/* Size of the following array; should always be 1024. */
	BiomesLength VarInt // VarInt
	/* 1024 biome IDs, ordered by x then z then y, in 4×4×4 blocks.  See Chunk Format § Biomes. */
	Biomes []VarInt // Array of VarInt
	/* Size of Data in bytes */
	Size VarInt // VarInt
	/* See data structure in Chunk Format */
	Data ByteArray // Byte array
	/* Number of elements in the following array */
	NumberOfBlockEntities VarInt // VarInt
	/* All block entities in the chunk.  Use the x, y, and z tags in the NBT to determine their positions. */
	BlockEntities []NBT // Array of NBT Tag
}

var _ Packet = (*PlayChunkData_756_0)(nil)

func (p PlayChunkData_756_0)Encode(b *PacketBuilder){
	b.Int(p.ChunkX)
	b.Int(p.ChunkZ)
	b.VarInt(p.BitMaskLength)
	TODO_Encode_Array(p.PrimaryBitMask)
	b.Encode(p.Heightmaps)
	p.BiomesLength = len(p.Biomes)
	b.VarInt(p.BiomesLength)
	for _, v := range p.Biomes {
		b.VarInt(v)
	}
	b.VarInt(p.Size)
	b.ByteArray(p.Data)
	p.NumberOfBlockEntities = len(p.BlockEntities)
	b.VarInt(p.NumberOfBlockEntities)
	for _, v := range p.BlockEntities {
		b.Encode(v)
	}
}

func (p *PlayChunkData_756_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.BitMaskLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.PrimaryBitMask = make([]Long, p.BitMaskLength)
	for i, _ := range p.PrimaryBitMask {
		if p.PrimaryBitMask[i], ok = r.Long(); !ok {
			return io.EOF
		}
	}
	if err = p.Heightmaps.DecodeFrom(r); err != nil {
		return
	}
	if p.BiomesLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Biomes = make([]VarInt, p.BiomesLength)
	for i, _ := range p.Biomes {
		if p.Biomes[i], ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
	if p.Size, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
	if p.NumberOfBlockEntities, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.BlockEntities = make([]NBT, p.NumberOfBlockEntities)
	for i, _ := range p.BlockEntities {
		if err = p.BlockEntities[i].DecodeFrom(r); err != nil {
			return
		}
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x20
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x20
type PlayChunkData_754_1 struct {
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkX Int // Int
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkZ Int // Int
	/* See Chunk Format */
	FullChunk Bool // Boolean
	/* Bitmask with bits set to 1 for every 16×16×16 chunk section whose data is included in Data. The least significant bit represents the chunk section at the bottom of the chunk column (from y=0 to y=15). */
	PrimaryBitMask VarInt // VarInt
	/* Compound containing one long array named MOTION_BLOCKING, which is a heightmap for the highest solid block at each position in the chunk (as a compacted long array with 256 entries at 9 bits per entry totaling 36 longs). The Notchian server also adds a WORLD_SURFACE long array, the purpose of which is unknown, but it's not required for the chunk to be accepted. */
	Heightmaps NBT // NBT
	/* Size of the following array.  Not present if full chunk is false. */
	BiomesLength VarInt // Optional VarInt
	/* 1024 biome IDs, ordered by x then z then y, in 4×4×4 blocks.Possibly something else.  Not present if full chunk is false. */
	Biomes []VarInt // Optional array of VarInt
	/* Size of Data in bytes */
	Size VarInt // VarInt
	/* See data structure in Chunk Format */
	Data ByteArray // Byte array
	/* Number of elements in the following array */
	NumberOfBlockEntities VarInt // VarInt
	/* All block entities in the chunk.  Use the x, y, and z tags in the NBT to determine their positions. */
	BlockEntities []NBT // Array of NBT Tag
}

var _ Packet = (*PlayChunkData_754_1)(nil)

func (p PlayChunkData_754_1)Encode(b *PacketBuilder){
	b.Int(p.ChunkX)
	b.Int(p.ChunkZ)
	b.Bool(p.FullChunk)
	b.VarInt(p.PrimaryBitMask)
	b.Encode(p.Heightmaps)
	p.BiomesLength = len(p.Biomes)
	b.VarInt(p.BiomesLength)
	for _, v := range p.Biomes {
		b.VarInt(v)
	}
	b.VarInt(p.Size)
	b.ByteArray(p.Data)
	p.NumberOfBlockEntities = len(p.BlockEntities)
	b.VarInt(p.NumberOfBlockEntities)
	for _, v := range p.BlockEntities {
		b.Encode(v)
	}
}

func (p *PlayChunkData_754_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.FullChunk, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.PrimaryBitMask, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Heightmaps.DecodeFrom(r); err != nil {
		return
	}
	if p.BiomesLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Biomes = make([]VarInt, p.BiomesLength)
	for i, _ := range p.Biomes {
		if p.Biomes[i], ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
	if p.Size, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
	if p.NumberOfBlockEntities, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.BlockEntities = make([]NBT, p.NumberOfBlockEntities)
	for i, _ := range p.BlockEntities {
		if err = p.BlockEntities[i].DecodeFrom(r); err != nil {
			return
		}
	}
}

// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x22
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x20
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x20
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x20
type PlayChunkData_401_5 struct {
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkX Int // Int
	/* Chunk coordinate (block coordinate divided by 16, rounded down) */
	ChunkZ Int // Int
	/* See Chunk Format */
	GroundUpContinuous Bool // Boolean
	/* Bitmask with bits set to 1 for every 16×16×16 chunk section whose data is included in Data. The least significant bit represents the chunk section at the bottom of the chunk column (from y=0 to y=15). */
	PrimaryBitMask VarInt // VarInt
	/* Size of Data in bytes */
	Size VarInt // VarInt
	/* See data structure in Chunk Format */
	Data ByteArray // Byte array
	/* Length of the following array */
	NumberOfBlockEntities VarInt // VarInt
	/* All block entities in the chunk.  Use the x, y, and z tags in the NBT to determine their positions. */
	BlockEntities []NBT // Array of NBT Tag
}

var _ Packet = (*PlayChunkData_401_5)(nil)

func (p PlayChunkData_401_5)Encode(b *PacketBuilder){
	b.Int(p.ChunkX)
	b.Int(p.ChunkZ)
	b.Bool(p.GroundUpContinuous)
	b.VarInt(p.PrimaryBitMask)
	b.VarInt(p.Size)
	b.ByteArray(p.Data)
	p.NumberOfBlockEntities = len(p.BlockEntities)
	b.VarInt(p.NumberOfBlockEntities)
	for _, v := range p.BlockEntities {
		b.Encode(v)
	}
}

func (p *PlayChunkData_401_5)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.GroundUpContinuous, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.PrimaryBitMask, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Size, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
	if p.NumberOfBlockEntities, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.BlockEntities = make([]NBT, p.NumberOfBlockEntities)
	for i, _ := range p.BlockEntities {
		if err = p.BlockEntities[i].DecodeFrom(r); err != nil {
			return
		}
	}
}

// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x20
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x20
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x20
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x20
type PlayChunkData_316_6 struct {
	/* Block coordinate divided by 16, rounded down */
	ChunkX Int // Int
	/* Block coordinate divided by 16, rounded down */
	ChunkZ Int // Int
	/* This is true if the packet represents all chunk sections in this vertical chunk column, where the Primary Bit Mask specifies exactly which chunk sections are included, and which are air. */
	GroundUpContinuous Bool // Boolean
	/* Bitmask with bits set to 1 for every 16×16×16 chunk section whose data is included in Data. The least significant bit represents the chunk section at the bottom of the chunk column (from y=0 to y=15). */
	PrimaryBitMask VarInt // VarInt
	/* Size of Data in bytes, plus size of Biomes in bytes if present */
	Size VarInt // VarInt
	/* The length of the array is equal to the number of bits set in Primary Bit Mask. Chunks are sent bottom-to-top, i.e. the first chunk, if sent, extends from Y=0 to Y=15. */
	Data []Unknown_chunk_section // Array of Chunk Section
	/* Only if Ground-Up Continuous: biome array, byte per XZ coordinate, 256 bytes total. */
	Biomes ByteArray // Optional Byte Array
	/* Length of the following array */
	NumberOfBlockEntities VarInt // VarInt
	/* All block entities in the chunk.  Use the x, y, and z tags in the NBT to determine their positions. */
	BlockEntities []NBT // Array of NBT Tag
}

var _ Packet = (*PlayChunkData_316_6)(nil)

func (p PlayChunkData_316_6)Encode(b *PacketBuilder){
	b.Int(p.ChunkX)
	b.Int(p.ChunkZ)
	b.Bool(p.GroundUpContinuous)
	b.VarInt(p.PrimaryBitMask)
	b.VarInt(p.Size)
	TODO_Encode_Array(p.Data)
	b.ByteArray(p.Biomes)
	p.NumberOfBlockEntities = len(p.BlockEntities)
	b.VarInt(p.NumberOfBlockEntities)
	for _, v := range p.BlockEntities {
		b.Encode(v)
	}
}

func (p *PlayChunkData_316_6)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.GroundUpContinuous, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.PrimaryBitMask, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Size, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.Data)
	if err = p.Biomes.DecodeFrom(r); err != nil {
		return
	}
	if p.NumberOfBlockEntities, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.BlockEntities = make([]NBT, p.NumberOfBlockEntities)
	for i, _ := range p.BlockEntities {
		if err = p.BlockEntities[i].DecodeFrom(r); err != nil {
			return
		}
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0xe
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0xe
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0xc
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0xd
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0xd
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x10
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x10
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x10
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x10
type PlayClearTitles_763_0 struct {
	Reset Bool // Boolean
}

var _ Packet = (*PlayClearTitles_763_0)(nil)

func (p PlayClearTitles_763_0)Encode(b *PacketBuilder){
	b.Bool(p.Reset)
}

func (p *PlayClearTitles_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Reset, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0xa
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0xa
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x9
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0xa
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x9
type PlayClickContainerButton_763_0 struct {
	/* The ID of the window sent by Open Screen. */
	WindowID Byte // Byte
	/* Meaning depends on window type; see below. */
	ButtonID Byte // Byte
}

var _ Packet = (*PlayClickContainerButton_763_0)(nil)

func (p PlayClickContainerButton_763_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Byte(p.ButtonID)
}

func (p *PlayClickContainerButton_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ButtonID, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x9
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x9
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x9
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x9
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x8
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x8
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x7
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x7
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x8
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x7
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x7
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x7
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x7
type PlayClickWindow_754_0 struct {
	/* The ID of the window which was clicked. 0 for player inventory. */
	WindowID UByte // Unsigned Byte
	/* The clicked slot number, see below */
	Slot Short // Short
	/* The button used in the click, see below */
	Button Byte // Byte
	/* A unique number for the action, implemented by Notchian as a counter, starting at 1. Used by the server to send back a Confirm Transaction (clientbound). */
	ActionNumber Short // Short
	/* Inventory operation mode, see below */
	Mode VarInt // VarInt Enum
	/* The clicked slot. Has to be empty (item ID = -1) for drop mode. */
	ClickedItem *Slot // Slot
}

var _ Packet = (*PlayClickWindow_754_0)(nil)

func (p PlayClickWindow_754_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.Short(p.Slot)
	b.Byte(p.Button)
	b.Short(p.ActionNumber)
	b.VarInt(p.Mode)
	b.Encode(p.ClickedItem)
}

func (p *PlayClickWindow_754_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.Mode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.ClickedItem.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x7
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x7
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x7
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x7
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x8
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x8
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x8
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x8
type PlayClickWindowButton_758_0 struct {
	/* The ID of the window sent by Open Window */
	WindowID Byte // Byte
	/* Meaning depends on window type; see below */
	ButtonID Byte // Byte
}

var _ Packet = (*PlayClickWindowButton_758_0)(nil)

func (p PlayClickWindowButton_758_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Byte(p.ButtonID)
}

func (p *PlayClickWindowButton_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ButtonID, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x7
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x7
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x6
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x7
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x6
type PlayClientCommand_763_0 struct {
	/* See below */
	ActionID VarInt // VarInt Enum
}

var _ Packet = (*PlayClientCommand_763_0)(nil)

func (p PlayClientCommand_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.ActionID)
}

func (p *PlayClientCommand_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ActionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x8
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x8
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x7
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x8
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x7
type PlayClientInformation_763_0 struct {
	/* e.g. en_GB. */
	Locale String // String (16)
	/* Client-side render distance, in chunks. */
	ViewDistance Byte // Byte
	/* 0: enabled, 1: commands only, 2: hidden.  See processing chat for more information. */
	ChatMode VarInt // VarInt Enum
	/* “Colors” multiplayer setting. Can the chat be colored? */
	ChatColors Bool // Boolean
	/* Bit mask, see below. */
	DisplayedSkinParts UByte // Unsigned Byte
	/* 0: Left, 1: Right. */
	MainHand VarInt // VarInt Enum
	/* Enables filtering of text on signs and written book titles. Currently always false (i.e. the filtering is disabled) */
	EnableTextFiltering Bool // Boolean
	/* Servers usually list online players, this option should let you not show up in that list. */
	AllowServerListings Bool // Boolean
}

var _ Packet = (*PlayClientInformation_763_0)(nil)

func (p PlayClientInformation_763_0)Encode(b *PacketBuilder){
	b.String(p.Locale)
	b.Byte(p.ViewDistance)
	b.VarInt(p.ChatMode)
	b.Bool(p.ChatColors)
	b.UByte(p.DisplayedSkinParts)
	b.VarInt(p.MainHand)
	b.Bool(p.EnableTextFiltering)
	b.Bool(p.AllowServerListings)
}

func (p *PlayClientInformation_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Locale, ok = r.String(); !ok {
		return io.EOF
	}
	if p.ViewDistance, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ChatMode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ChatColors, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.DisplayedSkinParts, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.MainHand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EnableTextFiltering, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.AllowServerListings, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x5
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x5
type PlayClientSettings_758_0 struct {
	/* e.g. en_GB. */
	Locale String // String (16)
	/* Client-side render distance, in chunks. */
	ViewDistance Byte // Byte
	/* 0: enabled, 1: commands only, 2: hidden.  See processing chat for more information. */
	ChatMode VarInt // VarInt Enum
	/* “Colors” multiplayer setting. Can the chat be colored? */
	ChatColors Bool // Boolean
	/* Bit mask, see below. */
	DisplayedSkinParts UByte // Unsigned Byte
	/* 0: Left, 1: Right. */
	MainHand VarInt // VarInt Enum
	/* Enables filtering of text on signs and written book titles. Currently always false (i.e. the filtering is disabled) */
	EnableTextFiltering Bool // Boolean
	/* Servers usually list online players, this option should let you not show up in that list. */
	AllowServerListings Bool // Boolean
}

var _ Packet = (*PlayClientSettings_758_0)(nil)

func (p PlayClientSettings_758_0)Encode(b *PacketBuilder){
	b.String(p.Locale)
	b.Byte(p.ViewDistance)
	b.VarInt(p.ChatMode)
	b.Bool(p.ChatColors)
	b.UByte(p.DisplayedSkinParts)
	b.VarInt(p.MainHand)
	b.Bool(p.EnableTextFiltering)
	b.Bool(p.AllowServerListings)
}

func (p *PlayClientSettings_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Locale, ok = r.String(); !ok {
		return io.EOF
	}
	if p.ViewDistance, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ChatMode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ChatColors, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.DisplayedSkinParts, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.MainHand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EnableTextFiltering, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.AllowServerListings, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x5
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x5
type PlayClientSettings_756_1 struct {
	/* e.g. en_GB. */
	Locale String // String (16)
	/* Client-side render distance, in chunks. */
	ViewDistance Byte // Byte
	/* 0: enabled, 1: commands only, 2: hidden.  See processing chat for more information. */
	ChatMode VarInt // VarInt Enum
	/* “Colors” multiplayer setting. */
	ChatColors Bool // Boolean
	/* Bit mask, see below. */
	DisplayedSkinParts UByte // Unsigned Byte
	/* 0: Left, 1: Right. */
	MainHand VarInt // VarInt Enum
	/* Disables filtering of text on signs and written book titles. Currently always true (i.e. the filtering is disabled) */
	DisableTextFiltering Bool // Boolean
}

var _ Packet = (*PlayClientSettings_756_1)(nil)

func (p PlayClientSettings_756_1)Encode(b *PacketBuilder){
	b.String(p.Locale)
	b.Byte(p.ViewDistance)
	b.VarInt(p.ChatMode)
	b.Bool(p.ChatColors)
	b.UByte(p.DisplayedSkinParts)
	b.VarInt(p.MainHand)
	b.Bool(p.DisableTextFiltering)
}

func (p *PlayClientSettings_756_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Locale, ok = r.String(); !ok {
		return io.EOF
	}
	if p.ViewDistance, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ChatMode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ChatColors, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.DisplayedSkinParts, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.MainHand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DisableTextFiltering, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x5
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x5
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x5
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x5
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x4
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x4
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x4
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x4
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x5
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x4
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x4
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x4
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x4
type PlayClientSettings_754_2 struct {
	/* e.g. en_GB */
	Locale String // String (7)
	/* Client-side render distance, in chunks */
	ViewDistance Byte // Byte
	/* 0: enabled, 1: commands only, 2: hidden.  See processing chat for more information. */
	ChatMode VarInt // VarInt Enum
	/* “Colors” multiplayer setting */
	ChatColors Bool // Boolean
	/* Bit mask, see below */
	DisplayedSkinParts UByte // Unsigned Byte
	/* 0: Left, 1: Right */
	MainHand VarInt // VarInt Enum
}

var _ Packet = (*PlayClientSettings_754_2)(nil)

func (p PlayClientSettings_754_2)Encode(b *PacketBuilder){
	b.String(p.Locale)
	b.Byte(p.ViewDistance)
	b.VarInt(p.ChatMode)
	b.Bool(p.ChatColors)
	b.UByte(p.DisplayedSkinParts)
	b.VarInt(p.MainHand)
}

func (p *PlayClientSettings_754_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Locale, ok = r.String(); !ok {
		return io.EOF
	}
	if p.ViewDistance, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ChatMode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ChatColors, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.DisplayedSkinParts, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.MainHand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x4
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x4
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x4
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x4
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x4
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x4
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x4
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x4
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x3
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x3
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x3
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x3
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x4
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x3
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x3
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x3
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x3
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x16
type PlayClientStatus_758_0 struct {
	/* See below */
	ActionID VarInt // VarInt
}

var _ Packet = (*PlayClientStatus_758_0)(nil)

func (p PlayClientStatus_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.ActionID)
}

func (p *PlayClientStatus_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ActionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x11
// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0xc
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x11
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0xc
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0xf
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0xb
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x10
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0xc
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x10
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0xb
type PlayCloseContainer_763_0 struct {
	/* This is the ID of the window that was closed. 0 for inventory. */
	WindowID UByte // Unsigned Byte
}

var _ Packet = (*PlayCloseContainer_763_0)(nil)

func (p PlayCloseContainer_763_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
}

func (p *PlayCloseContainer_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x13
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x9
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x13
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x9
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x13
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x9
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x13
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x9
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x12
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0xa
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x12
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0xa
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x14
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0xa
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x13
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0xa
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x13
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x9
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x13
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x9
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x12
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x8
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x12
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x8
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x12
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x9
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x12
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x8
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x12
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x8
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x12
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x8
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x12
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x8
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x2e
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0xd
type PlayCloseWindow_758_0 struct {
	/* This is the ID of the window that was closed. 0 for inventory. */
	WindowID UByte // Unsigned Byte
}

var _ Packet = (*PlayCloseWindow_758_0)(nil)

func (p PlayCloseWindow_758_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
}

func (p *PlayCloseWindow_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x61
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x61
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x60
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x60
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x55
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x55
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x56
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x55
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x4f
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x4f
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x4b
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x4b
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x4a
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x48
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x48
type PlayCollectItem_758_0 struct {
	CollectedEntityID VarInt // VarInt
	CollectorEntityID VarInt // VarInt
	/* Seems to be 1 for XP orbs, otherwise the number of items in the stack. */
	PickupItemCount VarInt // VarInt
}

var _ Packet = (*PlayCollectItem_758_0)(nil)

func (p PlayCollectItem_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.CollectedEntityID)
	b.VarInt(p.CollectorEntityID)
	b.VarInt(p.PickupItemCount)
}

func (p *PlayCollectItem_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.CollectedEntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CollectorEntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.PickupItemCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x48
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x48
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0xd
type PlayCollectItem_210_1 struct {
	CollectedEntityID VarInt // VarInt
	CollectorEntityID VarInt // VarInt
}

var _ Packet = (*PlayCollectItem_210_1)(nil)

func (p PlayCollectItem_210_1)Encode(b *PacketBuilder){
	b.VarInt(p.CollectedEntityID)
	b.VarInt(p.CollectorEntityID)
}

func (p *PlayCollectItem_210_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.CollectedEntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CollectorEntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x38
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x34
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x36
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x33
type PlayCombatDeath_762_1 struct {
	/* Entity ID of the player that died (should match the client's entity ID). */
	PlayerID VarInt // VarInt
	/* The killer entity's ID, or -1 if there is no obvious killer. */
	EntityID Int // Int
	/* The death message. */
	Message Object // Chat
}

var _ Packet = (*PlayCombatDeath_762_1)(nil)

func (p PlayCombatDeath_762_1)Encode(b *PacketBuilder){
	b.VarInt(p.PlayerID)
	b.Int(p.EntityID)
	b.JSON(p.Message)
}

func (p *PlayCombatDeath_762_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.PlayerID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.Message); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x9
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x9
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x8
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x9
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x8
type PlayCommandSuggestionsRequest_763_0 struct {
	/* The id of the transaction that the server will send back to the client in the response of this packet. Client generates this and increments it each time it sends another tab completion that doesn't get a response. */
	TransactionId VarInt // VarInt
	/* All text behind the cursor without the / (e.g. to the left of the cursor in left-to-right languages like English). */
	Text String // String (32500)
}

var _ Packet = (*PlayCommandSuggestionsRequest_763_0)(nil)

func (p PlayCommandSuggestionsRequest_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionId)
	b.String(p.Text)
}

func (p *PlayCommandSuggestionsRequest_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionId, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Text, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x10
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x10
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0xe
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0xf
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0xf
type PlayCommands_763_0 struct {
	/* Number of elements in the following array. */
	Count VarInt // VarInt
	/* An array of nodes. */
	Nodes []*CommandNode // Array of Node
	/* Index of the root node in the previous array. */
	RootIndex VarInt // VarInt
}

var _ Packet = (*PlayCommands_763_0)(nil)

func (p PlayCommands_763_0)Encode(b *PacketBuilder){
	p.Count = len(p.Nodes)
	b.VarInt(p.Count)
	for _, v := range p.Nodes {
		b.Encode(v)
	}
	b.VarInt(p.RootIndex)
}

func (p *PlayCommands_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Nodes = make([]*CommandNode, p.Count)
	for i, _ := range p.Nodes {
		if err = p.Nodes[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if p.RootIndex, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x0
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x0
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x0
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x0
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x0
type PlayConfirmTeleportation_763_0 struct {
	/* The ID given by the Synchronize Player Position packet. */
	TeleportID VarInt // VarInt
}

var _ Packet = (*PlayConfirmTeleportation_763_0)(nil)

func (p PlayConfirmTeleportation_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.TeleportID)
}

func (p *PlayConfirmTeleportation_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TeleportID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x12
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x6
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x12
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x6
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x11
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x5
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x11
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x5
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x11
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x6
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x11
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x5
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x11
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x5
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x11
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x5
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x11
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x5
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x32
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0xf
type PlayConfirmTransaction_404_0 struct {
	/* The ID of the window that the action occurred in */
	WindowID Byte // Byte
	/* Every action that is to be accepted has a unique number. This field corresponds to that number. */
	ActionNumber Short // Short
	/* Whether the action was accepted */
	Accepted Bool // Boolean
}

var _ Packet = (*PlayConfirmTransaction_404_0)(nil)

func (p PlayConfirmTransaction_404_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Short(p.ActionNumber)
	b.Bool(p.Accepted)
}

func (p *PlayConfirmTransaction_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ActionNumber, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.Accepted, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x18
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x18
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x18
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x18
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x19
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x19
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x18
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x18
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x16
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x16
type PlayCraftRecipeRequest_758_0 struct {
	WindowID Byte // Byte
	/* A recipe ID */
	Recipe Identifier // Identifier
	/* Affects the amount of items processed; true if shift is down when clicked */
	MakeAll Bool // Boolean
}

var _ Packet = (*PlayCraftRecipeRequest_758_0)(nil)

func (p PlayCraftRecipeRequest_758_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Encode(p.Recipe)
	b.Bool(p.MakeAll)
}

func (p *PlayCraftRecipeRequest_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if err = p.Recipe.DecodeFrom(r); err != nil {
		return
	}
	if p.MakeAll, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x12
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x12
type PlayCraftRecipeRequest_340_1 struct {
	WindowID Byte // Byte
	/* A recipe ID */
	Recipe VarInt // VarInt
	/* Affects the amount of items processed; true if shift is down when clicked */
	MakeAll Bool // Boolean
}

var _ Packet = (*PlayCraftRecipeRequest_340_1)(nil)

func (p PlayCraftRecipeRequest_340_1)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.VarInt(p.Recipe)
	b.Bool(p.MakeAll)
}

func (p *PlayCraftRecipeRequest_340_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Recipe, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.MakeAll, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x31
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x31
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x31
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x31
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x2f
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x2f
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x31
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x30
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x2d
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x2d
type PlayCraftRecipeResponse_758_0 struct {
	WindowID Byte // Byte
	/* A recipe ID */
	Recipe Identifier // Identifier
}

var _ Packet = (*PlayCraftRecipeResponse_758_0)(nil)

func (p PlayCraftRecipeResponse_758_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Encode(p.Recipe)
}

func (p *PlayCraftRecipeResponse_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if err = p.Recipe.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x2b
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x2b
type PlayCraftRecipeResponse_340_1 struct {
	WindowID Byte // Byte
	/* A recipe ID */
	Recipe VarInt // VarInt
}

var _ Packet = (*PlayCraftRecipeResponse_340_1)(nil)

func (p PlayCraftRecipeResponse_340_1)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.VarInt(p.Recipe)
}

func (p *PlayCraftRecipeResponse_340_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Recipe, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x28
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x28
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x28
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x28
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x28
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x29
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x26
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x26
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x24
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x24
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x1b
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x1b
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x1b
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x18
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x18
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x18
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x18
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x10
type PlayCreativeInventoryAction_758_0 struct {
	/* Inventory slot */
	Slot Short // Short
	ClickedItem *Slot // Slot
}

var _ Packet = (*PlayCreativeInventoryAction_758_0)(nil)

func (p PlayCreativeInventoryAction_758_0)Encode(b *PacketBuilder){
	b.Short(p.Slot)
	b.Encode(p.ClickedItem)
}

func (p *PlayCreativeInventoryAction_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
	if err = p.ClickedItem.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x17
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x16
type PlayCustomSoundEffect_760_0 struct {
	/* All sound effect names can be seen here. */
	SoundName Identifier // Identifier
	/* The category that this sound will be played from (current categories). */
	SoundCategory VarInt // VarInt Enum
	/* Effect X multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionX Int // Int
	/* Effect Y multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionY Int // Int
	/* Effect Z multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionZ Int // Int
	/* 1 is 100%, can be more. */
	Volume Float // Float
	/* Float between 0.5 and 2.0 by Notchian clients. */
	Pitch Float // Float
	/* Seed used to pick sound variant. */
	Seed Long // long
}

var _ Packet = (*PlayCustomSoundEffect_760_0)(nil)

func (p PlayCustomSoundEffect_760_0)Encode(b *PacketBuilder){
	b.Encode(p.SoundName)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.Float(p.Pitch)
	b.Long(p.Seed)
}

func (p *PlayCustomSoundEffect_760_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.SoundName.DecodeFrom(r); err != nil {
		return
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
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Seed, ok = r.Long(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x35
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x35
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x35
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x35
type PlayDeathCombatEvent_758_0 struct {
	/* Entity ID of the player that died (should match the client's entity ID). */
	PlayerID VarInt // VarInt
	/* The killer entity's ID, or -1 if there is no obvious killer. */
	EntityID Int // Int
	/* The death message. */
	Message Object // Chat
}

var _ Packet = (*PlayDeathCombatEvent_758_0)(nil)

func (p PlayDeathCombatEvent_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.PlayerID)
	b.Int(p.EntityID)
	b.JSON(p.Message)
}

func (p *PlayDeathCombatEvent_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.PlayerID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.Message); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x12
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x12
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x12
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x12
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x10
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x10
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x12
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x11
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x11
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x11
type PlayDeclareCommands_758_0 struct {
	/* Number of elements in the following array */
	Count VarInt // VarInt
	/* An array of nodes */
	Nodes []*CommandNode // Array of Node
	/* Index of the root node in the previous array */
	RootIndex VarInt // VarInt
}

var _ Packet = (*PlayDeclareCommands_758_0)(nil)

func (p PlayDeclareCommands_758_0)Encode(b *PacketBuilder){
	p.Count = len(p.Nodes)
	b.VarInt(p.Count)
	for _, v := range p.Nodes {
		b.Encode(v)
	}
	b.VarInt(p.RootIndex)
}

func (p *PlayDeclareCommands_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Nodes = make([]*CommandNode, p.Count)
	for i, _ := range p.Nodes {
		if err = p.Nodes[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if p.RootIndex, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x19
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x19
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x16
type PlayDeleteMessage_763_0 struct {
	/* Length of Signature. */
	SignatureLength VarInt // VarInt
	/* Bytes of the signature. */
	Signature ByteArray // Byte Array
}

var _ Packet = (*PlayDeleteMessage_763_0)(nil)

func (p PlayDeleteMessage_763_0)Encode(b *PacketBuilder){
	p.SignatureLength = len(p.Signature)
	b.VarInt(p.SignatureLength)
	b.ByteArray(p.Signature)
}

func (p *PlayDeleteMessage_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SignatureLength, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Signature = make(ByteArray, p.SignatureLength)
	if ok = r.ByteArray(p.Signature); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x3a
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x3a
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x3a
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x36
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x36
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x38
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x37
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x35
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x35
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x32
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x32
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x31
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x30
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x30
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x30
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x30
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x13
type PlayDestroyEntities_758_0 struct {
	/* Number of elements in the following array */
	Count VarInt // VarInt
	/* The list of entities of destroy */
	EntityIDs []VarInt // Array of VarInt
}

var _ Packet = (*PlayDestroyEntities_758_0)(nil)

func (p PlayDestroyEntities_758_0)Encode(b *PacketBuilder){
	p.Count = len(p.EntityIDs)
	b.VarInt(p.Count)
	for _, v := range p.EntityIDs {
		b.VarInt(v)
	}
}

func (p *PlayDestroyEntities_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.EntityIDs = make([]VarInt, p.Count)
	for i, _ := range p.EntityIDs {
		if p.EntityIDs[i], ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x1a
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x1a
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x17
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x19
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x17
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x1a
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x1a
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x1a
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x1a
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x19
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x19
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x1b
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x1a
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x1b
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x1b
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x1a
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x1a
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x1a
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x1a
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x1a
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x1a
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x1a
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x40
type PlayDisconnect_763_0 struct {
	/* Displayed to the client when the connection terminates. */
	Reason Object // Chat
}

var _ Packet = (*PlayDisconnect_763_0)(nil)

func (p PlayDisconnect_763_0)Encode(b *PacketBuilder){
	b.JSON(p.Reason)
}

func (p *PlayDisconnect_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.Reason); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x1b
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x1b
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x18
type PlayDisguisedChatMessage_763_0 struct {
	Message Object // Chat
	/* The chat message type. */
	ChatType VarInt // VarInt
	/* The name associated with the chat type. Usually the message sender's display name. */
	ChatTypeName Object // Chat
	/* True if target name is present. */
	HasTargetName Bool // Boolean
	/* The target name associated with the chat type. Usually the message target's display name. Only present if previous boolean is true. */
	TargetName Object // Chat
}

var _ Packet = (*PlayDisguisedChatMessage_763_0)(nil)

func (p PlayDisguisedChatMessage_763_0)Encode(b *PacketBuilder){
	b.JSON(p.Message)
	b.VarInt(p.ChatType)
	b.JSON(p.ChatTypeName)
	b.Bool(p.HasTargetName)
	b.JSON(p.TargetName)
}

func (p *PlayDisguisedChatMessage_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.Message); err != nil {
		return
	}
	if p.ChatType, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.ChatTypeName); err != nil {
		return
	}
	if p.HasTargetName, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.TargetName); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x51
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x51
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x4d
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x4f
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x4c
type PlayDisplayObjective_763_0 struct {
	/* The position of the scoreboard. 0: list, 1: sidebar, 2: below name, 3 - 18: team specific sidebar, indexed as 3 + team color. */
	Position Byte // Byte
	/* The unique name for the scoreboard to be displayed. */
	ScoreName String // String (16)
}

var _ Packet = (*PlayDisplayObjective_763_0)(nil)

func (p PlayDisplayObjective_763_0)Encode(b *PacketBuilder){
	b.Byte(p.Position)
	b.String(p.ScoreName)
}

func (p *PlayDisplayObjective_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Position, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ScoreName, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x4c
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x4c
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x4c
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x4c
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x43
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x43
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x43
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x42
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x3e
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x3e
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x3b
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x3b
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x3a
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x38
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x38
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x38
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x38
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x3d
type PlayDisplayScoreboard_758_0 struct {
	/* The position of the scoreboard. 0: list, 1: sidebar, 2: below name. */
	Position Byte // Byte
	/* The unique name for the scoreboard to be displayed. */
	ScoreName String // String
}

var _ Packet = (*PlayDisplayScoreboard_758_0)(nil)

func (p PlayDisplayScoreboard_758_0)Encode(b *PacketBuilder){
	b.Byte(p.Position)
	b.String(p.ScoreName)
}

func (p *PlayDisplayScoreboard_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Position, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ScoreName, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0xe
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0xe
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0xd
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0xe
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0xd
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0xb
type PlayEditBook_763_0 struct {
	/* The hotbar slot where the written book is located */
	Slot VarInt // VarInt
	/* Number of elements in the following array */
	Count VarInt // VarInt
	/* Text from each page. */
	Entries [][]String // Array of Strings (8192 chars)
	/* If true, the next field is present. */
	HasTitle Bool // Boolean
	/* Title of book. */
	Title String // Optional String (128 chars)
}

var _ Packet = (*PlayEditBook_763_0)(nil)

func (p PlayEditBook_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Slot)
	p.Count = len(p.Entries)
	b.VarInt(p.Count)
	for _, v := range p.Entries {
		TODO_Encode_Array(v)
	}
	b.Bool(p.HasTitle)
	b.String(p.Title)
}

func (p *PlayEditBook_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Entries = make([][]String, p.Count)
	for i, _ := range p.Entries {
		TODO_Decode_Array(Entries[i])
	}
	if p.HasTitle, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Title, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0xb
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0xc
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0xc
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0xc
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0xc
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0xb
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0xb
type PlayEditBook_755_3 struct {
	NewBook *Slot // Slot
	/* True if the player is signing the book; false if the player is saving a draft. */
	IsSigning Bool // Boolean
	/* 0: Main hand, 1: Off hand */
	Hand VarInt // VarInt enum
}

var _ Packet = (*PlayEditBook_755_3)(nil)

func (p PlayEditBook_755_3)Encode(b *PacketBuilder){
	b.Encode(p.NewBook)
	b.Bool(p.IsSigning)
	b.VarInt(p.Hand)
}

func (p *PlayEditBook_755_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.NewBook.DecodeFrom(r); err != nil {
		return
	}
	if p.IsSigning, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x23
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x23
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x23
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x23
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x21
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x21
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x23
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x22
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x23
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x23
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x21
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x21
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x21
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x21
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x21
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x21
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x21
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x28
type PlayEffect_758_0 struct {
	/* The ID of the effect, see below */
	EffectID Int // Int
	/* The location of the effect */
	Location Position // Position
	/* Extra data for certain effects, see below */
	Data Int // Int
	/* See above */
	DisableRelativeVolume Bool // Boolean
}

var _ Packet = (*PlayEffect_758_0)(nil)

func (p PlayEffect_758_0)Encode(b *PacketBuilder){
	b.Int(p.EffectID)
	b.Encode(p.Location)
	b.Int(p.Data)
	b.Bool(p.DisableRelativeVolume)
}

func (p *PlayEffect_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EffectID, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Data, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.DisableRelativeVolume, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x7
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x7
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x6
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x6
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x7
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x6
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x6
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x6
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x6
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x11
type PlayEnchantItem_404_0 struct {
	/* The ID of the enchantment table window sent by Open Window */
	WindowID Byte // Byte
	/* The position of the enchantment on the enchantment table window, starting with 0 as the topmost one */
	Enchantment Byte // Byte
}

var _ Packet = (*PlayEnchantItem_404_0)(nil)

func (p PlayEnchantItem_404_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Byte(p.Enchantment)
}

func (p *PlayEnchantItem_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Enchantment, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x36
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x32
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x34
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x31
type PlayEndCombat_762_1 struct {
	/* Length of the combat in ticks. */
	Duration VarInt // VarInt
	/* ID of the primary opponent of the ended combat, or -1 if there is no obvious primary opponent. */
	EntityID Int // Int
}

var _ Packet = (*PlayEndCombat_762_1)(nil)

func (p PlayEndCombat_762_1)Encode(b *PacketBuilder){
	b.VarInt(p.Duration)
	b.Int(p.EntityID)
}

func (p *PlayEndCombat_762_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Duration, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x33
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x33
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x33
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x33
type PlayEndCombatEvent_758_0 struct {
	/* Length of the combat in ticks. */
	Duration VarInt // VarInt
	/* ID of the primary opponent of the ended combat, or -1 if there is no obvious primary opponent. */
	EntityID Int // Int
}

var _ Packet = (*PlayEndCombatEvent_758_0)(nil)

func (p PlayEndCombatEvent_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.Duration)
	b.Int(p.EntityID)
}

func (p *PlayEndCombatEvent_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Duration, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x37
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x37
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x33
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x35
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x32
type PlayEnterCombat_763_0 struct {
}

var _ Packet = (*PlayEnterCombat_763_0)(nil)

func (p PlayEnterCombat_763_0)Encode(b *PacketBuilder){
}

func (p *PlayEnterCombat_763_0)DecodeFrom(r *PacketReader)(err error){ return }

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x34
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x34
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x34
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x34
type PlayEnterCombatEvent_758_0 struct {
}

var _ Packet = (*PlayEnterCombatEvent_758_0)(nil)

func (p PlayEnterCombatEvent_758_0)Encode(b *PacketBuilder){
}

func (p *PlayEnterCombatEvent_758_0)DecodeFrom(r *PacketReader)(err error){ return }

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x27
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x27
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x25
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x25
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x25
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x28
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x28
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x28
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x28
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x14
type PlayEntity_404_0 struct {
	EntityID VarInt // VarInt
}

var _ Packet = (*PlayEntity_404_0)(nil)

func (p PlayEntity_404_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
}

func (p *PlayEntity_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x1b
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x1b
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x1b
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x1b
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x1c
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x1c
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x1b
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x1b
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x19
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x19
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x15
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x15
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x15
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x14
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x14
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x14
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x14
type PlayEntityAction_758_0 struct {
	/* Player ID */
	EntityID VarInt // VarInt
	/* The ID of the action, see below */
	ActionID VarInt // VarInt Enum
	/* Only used by the “start jump with horse” action, in which case it ranges from 0 to 100. In all other cases it is 0. */
	JumpBoost VarInt // VarInt
}

var _ Packet = (*PlayEntityAction_758_0)(nil)

func (p PlayEntityAction_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.ActionID)
	b.VarInt(p.JumpBoost)
}

func (p *PlayEntityAction_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ActionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.JumpBoost, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x4
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x4
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x3
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x3
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x3
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x6
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x6
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x6
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x6
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x5
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x5
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x6
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x6
type PlayEntityAnimation_763_0 struct {
	/* Player ID */
	EntityID VarInt // VarInt
	/* Animation ID (see below) */
	Animation UByte // Unsigned Byte
}

var _ Packet = (*PlayEntityAnimation_763_0)(nil)

func (p PlayEntityAnimation_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UByte(p.Animation)
}

func (p *PlayEntityAnimation_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Animation, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x6c
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x6c
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x68
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x69
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x66
type PlayEntityEffect_763_0 struct {
	EntityID VarInt // VarInt
	/* See this table. */
	EffectID VarInt // VarInt
	/* Notchian client displays effect level as Amplifier + 1. */
	Amplifier Byte // Byte
	/* Duration in ticks. */
	Duration VarInt // VarInt
	/* Bit field, see below. */
	Flags Byte // Byte
	/* Used in DARKNESS effect */
	HasFactorData Bool // Boolean
	/* See below */
	FactorCodec NBT // NBT Tag
}

var _ Packet = (*PlayEntityEffect_763_0)(nil)

func (p PlayEntityEffect_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.EffectID)
	b.Byte(p.Amplifier)
	b.VarInt(p.Duration)
	b.Byte(p.Flags)
	b.Bool(p.HasFactorData)
	b.Encode(p.FactorCodec)
}

func (p *PlayEntityEffect_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Amplifier, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Duration, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.HasFactorData, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = p.FactorCodec.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x65
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x64
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x64
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x59
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x59
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x5a
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x59
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x53
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x53
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x4f
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x4f
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x4e
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x4b
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x4b
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x4b
type PlayEntityEffect_757_2 struct {
	EntityID VarInt // VarInt
	/* See this table */
	EffectID Byte // Byte
	/* Notchian client displays effect level as Amplifier + 1 */
	Amplifier Byte // Byte
	/* Seconds */
	Duration VarInt // VarInt
	/* Bit field, see below. */
	Flags Byte // Byte
}

var _ Packet = (*PlayEntityEffect_757_2)(nil)

func (p PlayEntityEffect_757_2)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Byte(p.EffectID)
	b.Byte(p.Amplifier)
	b.VarInt(p.Duration)
	b.Byte(p.Flags)
}

func (p *PlayEntityEffect_757_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Amplifier, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Duration, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x4b
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x1d
type PlayEntityEffect_110_3 struct {
	EntityID VarInt // VarInt
	/* See this table */
	EffectID Byte // Byte
	/* Notchian client displays effect level as Amplifier + 1 */
	Amplifier Byte // Byte
	/* Seconds */
	Duration VarInt // VarInt
	HideParticles Bool // Boolean
}

var _ Packet = (*PlayEntityEffect_110_3)(nil)

func (p PlayEntityEffect_110_3)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Byte(p.EffectID)
	b.Byte(p.Amplifier)
	b.VarInt(p.Duration)
	b.Bool(p.HideParticles)
}

func (p *PlayEntityEffect_110_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Amplifier, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Duration, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.HideParticles, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x47
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x46
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x42
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x42
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x3f
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x3f
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x3e
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x3c
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x3c
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x3c
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x3c
type PlayEntityEquipment_578_0 struct {
	/* Entity's EID */
	EntityID VarInt // VarInt
	/* Equipment slot. 0: main hand, 1: off hand, 2–5: armor slot (2: boots, 3: leggings, 4: chestplate, 5: helmet) */
	Slot VarInt // VarInt Enum
	Item *Slot // Slot
}

var _ Packet = (*PlayEntityEquipment_578_0)(nil)

func (p PlayEntityEquipment_578_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.Slot)
	b.Encode(p.Item)
}

func (p *PlayEntityEquipment_578_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Slot, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Item.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x1c
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x1c
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x19
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x1a
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x18
type PlayEntityEvent_763_0 struct {
	EntityID Int // Int
	/* See Entity statuses for a list of which statuses are valid for each type of entity. */
	EntityStatus Byte // Byte Enum
}

var _ Packet = (*PlayEntityEvent_763_0)(nil)

func (p PlayEntityEvent_763_0)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.Byte(p.EntityStatus)
}

func (p *PlayEntityEvent_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EntityStatus, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x3e
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x3e
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x3e
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x3e
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x3a
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x3a
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x3c
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x3b
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x39
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x39
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x36
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x36
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x35
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x34
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x34
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x34
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x34
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x19
type PlayEntityHeadLook_758_0 struct {
	EntityID VarInt // VarInt
	/* New angle, not a delta */
	HeadYaw Angle // Angle
}

var _ Packet = (*PlayEntityHeadLook_758_0)(nil)

func (p PlayEntityHeadLook_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.HeadYaw)
}

func (p *PlayEntityHeadLook_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.HeadYaw.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x2a
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x2a
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x28
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x28
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x28
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x27
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x27
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x27
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x27
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x16
type PlayEntityLook_404_0 struct {
	EntityID VarInt // VarInt
	/* New angle, not a delta */
	Yaw Angle // Angle
	/* New angle, not a delta */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityLook_404_0)(nil)

func (p PlayEntityLook_404_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayEntityLook_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x29
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x29
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x27
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x27
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x27
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x26
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x26
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x26
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x26
type PlayEntityLookAndRelativeMove_404_0 struct {
	EntityID VarInt // VarInt
	/* Change in X position as (currentX * 32 - prevX * 32) * 128 */
	DeltaX Short // Short
	/* Change in Y position as (currentY * 32 - prevY * 32) * 128 */
	DeltaY Short // Short
	/* Change in Z position as (currentZ * 32 - prevZ * 32) * 128 */
	DeltaZ Short // Short
	/* New angle, not a delta */
	Yaw Angle // Angle
	/* New angle, not a delta */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityLookAndRelativeMove_404_0)(nil)

func (p PlayEntityLookAndRelativeMove_404_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.DeltaX)
	b.Short(p.DeltaY)
	b.Short(p.DeltaZ)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayEntityLookAndRelativeMove_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DeltaX, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaY, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaZ, ok = r.Short(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x4d
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x4d
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x4d
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x4d
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x44
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x44
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x44
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x43
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x3f
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x3f
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x3c
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x3c
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x3b
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x39
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x39
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x39
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x39
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x1c
type PlayEntityMetadata_758_0 struct {
	EntityID VarInt // VarInt
	Metadata *EntityMetadata // Metadata
}

var _ Packet = (*PlayEntityMetadata_758_0)(nil)

func (p PlayEntityMetadata_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.Metadata)
}

func (p *PlayEntityMetadata_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Metadata.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x2a
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x2a
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x2c
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x2b
type PlayEntityMovement_754_0 struct {
	EntityID VarInt // VarInt
}

var _ Packet = (*PlayEntityMovement_754_0)(nil)

func (p PlayEntityMovement_754_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
}

func (p *PlayEntityMovement_754_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x29
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x29
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x29
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x29
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x27
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x27
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x29
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x28
type PlayEntityPosition_758_0 struct {
	EntityID VarInt // VarInt
	/* Change in X position as (currentX * 32 - prevX * 32) * 128 */
	DeltaX Short // Short
	/* Change in Y position as (currentY * 32 - prevY * 32) * 128 */
	DeltaY Short // Short
	/* Change in Z position as (currentZ * 32 - prevZ * 32) * 128 */
	DeltaZ Short // Short
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityPosition_758_0)(nil)

func (p PlayEntityPosition_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.DeltaX)
	b.Short(p.DeltaY)
	b.Short(p.DeltaZ)
	b.Bool(p.OnGround)
}

func (p *PlayEntityPosition_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DeltaX, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaY, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaZ, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x2a
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x2a
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x2a
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x2a
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x28
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x28
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x2a
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x29
type PlayEntityPositionAndRotation_758_0 struct {
	EntityID VarInt // VarInt
	/* Change in X position as (currentX * 32 - prevX * 32) * 128 */
	DeltaX Short // Short
	/* Change in Y position as (currentY * 32 - prevY * 32) * 128 */
	DeltaY Short // Short
	/* Change in Z position as (currentZ * 32 - prevZ * 32) * 128 */
	DeltaZ Short // Short
	/* New angle, not a delta */
	Yaw Angle // Angle
	/* New angle, not a delta */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityPositionAndRotation_758_0)(nil)

func (p PlayEntityPositionAndRotation_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.DeltaX)
	b.Short(p.DeltaY)
	b.Short(p.DeltaZ)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayEntityPositionAndRotation_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DeltaX, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaY, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaZ, ok = r.Short(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x28
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x28
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x26
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x26
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x26
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x25
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x25
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x25
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x25
type PlayEntityRelativeMove_404_0 struct {
	EntityID VarInt // VarInt
	/* Change in X position as (currentX * 32 - prevX * 32) * 128 */
	DeltaX Short // Short
	/* Change in Y position as (currentY * 32 - prevY * 32) * 128 */
	DeltaY Short // Short
	/* Change in Z position as (currentZ * 32 - prevZ * 32) * 128 */
	DeltaZ Short // Short
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityRelativeMove_404_0)(nil)

func (p PlayEntityRelativeMove_404_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.DeltaX)
	b.Short(p.DeltaY)
	b.Short(p.DeltaZ)
	b.Bool(p.OnGround)
}

func (p *PlayEntityRelativeMove_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DeltaX, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaY, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaZ, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x2b
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x2b
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x2b
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x2b
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x29
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x29
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x2b
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x2a
type PlayEntityRotation_758_0 struct {
	EntityID VarInt // VarInt
	/* New angle, not a delta */
	Yaw Angle // Angle
	/* New angle, not a delta */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityRotation_758_0)(nil)

func (p PlayEntityRotation_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayEntityRotation_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x61
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x61
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x5d
type PlayEntitySoundEffect_763_0 struct {
	/* Represents the Sound ID + 1. If the value is 0, the packet contains a sound specified by Identifier. */
	SoundID VarInt // VarInt
	/* Only present if Sound ID is 0 */
	SoundName Identifier // Optional Identifier
	/* Only present if Sound ID is 0. */
	HasFixedRange Bool // Optional Boolean
	/* The fixed range of the sound. Only present if previous boolean is true and Sound ID is 0. */
	Range Float // Optional Float
	/* The category that this sound will be played from (current categories). */
	SoundCategory VarInt // VarInt Enum
	EntityID VarInt // VarInt
	/* 1.0 is 100%, capped between 0.0 and 1.0 by Notchian clients. */
	Volume Float // Float
	/* Float between 0.5 and 2.0 by Notchian clients. */
	Pitch Float // Float
	/* Seed used to pick sound variant. */
	Seed Long // Long
}

var _ Packet = (*PlayEntitySoundEffect_763_0)(nil)

func (p PlayEntitySoundEffect_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.SoundID)
	b.Encode(p.SoundName)
	b.Bool(p.HasFixedRange)
	b.Float(p.Range)
	b.VarInt(p.SoundCategory)
	b.VarInt(p.EntityID)
	b.Float(p.Volume)
	b.Float(p.Pitch)
	b.Long(p.Seed)
}

func (p *PlayEntitySoundEffect_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SoundID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.SoundName.DecodeFrom(r); err != nil {
		return
	}
	if p.HasFixedRange, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Range, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.SoundCategory, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Volume, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Seed, ok = r.Long(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x5f
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x5c
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x5c
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x5c
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x5b
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x5b
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x50
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x50
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x51
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x50
type PlayEntitySoundEffect_760_1 struct {
	/* ID of hardcoded sound event (events as of 1.15.2) */
	SoundID VarInt // VarInt
	/* The category that this sound will be played from (current categories) */
	SoundCategory VarInt // VarInt Enum
	EntityID VarInt // VarInt
	/* 1.0 is 100%, capped between 0.0 and 1.0 by Notchian clients */
	Volume Float // Float
	/* Float between 0.5 and 2.0 by Notchian clients */
	Pitch Float // Float
}

var _ Packet = (*PlayEntitySoundEffect_760_1)(nil)

func (p PlayEntitySoundEffect_760_1)Encode(b *PacketBuilder){
	b.VarInt(p.SoundID)
	b.VarInt(p.SoundCategory)
	b.VarInt(p.EntityID)
	b.Float(p.Volume)
	b.Float(p.Pitch)
}

func (p *PlayEntitySoundEffect_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SoundID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.SoundCategory, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Volume, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x1b
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x1b
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x1b
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x1b
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x1a
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x1a
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x1c
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x1b
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x1c
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x1c
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x1b
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x1b
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x1b
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x1b
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x1b
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x1b
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x1b
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x1a
type PlayEntityStatus_758_0 struct {
	EntityID Int // Int
	/* See below */
	EntityStatus Byte // Byte
}

var _ Packet = (*PlayEntityStatus_758_0)(nil)

func (p PlayEntityStatus_758_0)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.Byte(p.EntityStatus)
}

func (p *PlayEntityStatus_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.EntityStatus, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x62
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x62
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x61
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x61
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x56
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x56
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x57
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x56
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x50
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x50
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x4c
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x4c
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x4b
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x49
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x49
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x49
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x49
type PlayEntityTeleport_758_0 struct {
	EntityID VarInt // VarInt
	X Double // Double
	Y Double // Double
	Z Double // Double
	/* New angle, not a delta */
	Yaw Angle // Angle
	/* New angle, not a delta */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayEntityTeleport_758_0)(nil)

func (p PlayEntityTeleport_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayEntityTeleport_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x4f
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x4f
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x4f
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x4f
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x46
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x46
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x46
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x45
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x41
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x41
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x3e
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x3e
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x3d
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x3b
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x3b
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x3b
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x3b
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x12
type PlayEntityVelocity_758_0 struct {
	EntityID VarInt // VarInt
	/* Velocity on the X axis */
	VelocityX Short // Short
	/* Velocity on the Y axis */
	VelocityY Short // Short
	/* Velocity on the Z axis */
	VelocityZ Short // Short
}

var _ Packet = (*PlayEntityVelocity_758_0)(nil)

func (p PlayEntityVelocity_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
}

func (p *PlayEntityVelocity_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
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
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x1d
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x1d
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x1a
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x19
type PlayExplosion_763_0 struct {
	X Double // Double
	Y Double // Double
	Z Double // Double
	/* A strength greater than or equal to 2.0 spawns a minecraft:explosion_emitter particle, while a lesser strength spawns a minecraft:explosion particle. */
	Strength Float // Float
	/* Number of elements in the following array. */
	RecordCount VarInt // VarInt
	/* Each record is 3 signed bytes long; the 3 bytes are the XYZ (respectively) signed offsets of affected blocks. */
	Records [][3]byte // Array of (Byte, Byte, Byte)
	/* X velocity of the player being pushed by the explosion. */
	PlayerMotionX Float // Float
	/* Y velocity of the player being pushed by the explosion. */
	PlayerMotionY Float // Float
	/* Z velocity of the player being pushed by the explosion. */
	PlayerMotionZ Float // Float
}

var _ Packet = (*PlayExplosion_763_0)(nil)

func (p PlayExplosion_763_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Strength)
	p.RecordCount = len(p.Records)
	b.VarInt(p.RecordCount)
	for _, v := range p.Records {
		b.Encode(v)
	}
	b.Float(p.PlayerMotionX)
	b.Float(p.PlayerMotionY)
	b.Float(p.PlayerMotionZ)
}

func (p *PlayExplosion_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Strength, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.RecordCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Records = make([][3]byte, p.RecordCount)
	for i, _ := range p.Records {
		if err = p.Records[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if p.PlayerMotionX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.PlayerMotionY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.PlayerMotionZ, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x1b
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x1c
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x1c
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x1c
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x1c
type PlayExplosion_760_1 struct {
	X Float // Float
	Y Float // Float
	Z Float // Float
	/* A strength greater than or equal to 2.0 spawns a minecraft:explosion_emitter particle, while a lesser strength spawns a minecraft:explosion particle. */
	Strength Float // Float
	/* Number of elements in the following array. */
	RecordCount VarInt // VarInt
	/* Each record is 3 signed bytes long; the 3 bytes are the XYZ (respectively) signed offsets of affected blocks. */
	Records [][3]byte // Array of (Byte, Byte, Byte)
	/* X velocity of the player being pushed by the explosion. */
	PlayerMotionX Float // Float
	/* Y velocity of the player being pushed by the explosion. */
	PlayerMotionY Float // Float
	/* Z velocity of the player being pushed by the explosion. */
	PlayerMotionZ Float // Float
}

var _ Packet = (*PlayExplosion_760_1)(nil)

func (p PlayExplosion_760_1)Encode(b *PacketBuilder){
	b.Float(p.X)
	b.Float(p.Y)
	b.Float(p.Z)
	b.Float(p.Strength)
	p.RecordCount = len(p.Records)
	b.VarInt(p.RecordCount)
	for _, v := range p.Records {
		b.Encode(v)
	}
	b.Float(p.PlayerMotionX)
	b.Float(p.PlayerMotionY)
	b.Float(p.PlayerMotionZ)
}

func (p *PlayExplosion_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Strength, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.RecordCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Records = make([][3]byte, p.RecordCount)
	for i, _ := range p.Records {
		if err = p.Records[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if p.PlayerMotionX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.PlayerMotionY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.PlayerMotionZ, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x1b
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x1b
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x1d
type PlayExplosion_754_2 struct {
	X Float // Float
	Y Float // Float
	Z Float // Float
	/* A strength greater than or equal to 2.0 spawns a minecraft:explosion_emitter particle, while a lesser strength spawns a minecraft:explosion particle. */
	Strength Float // Float
	/* Number of elements in the following array */
	RecordCount Int // Int
	/* Each record is 3 signed bytes long; the 3 bytes are the XYZ (respectively) signed offsets of affected blocks. */
	Records [][3]byte // Array of (Byte, Byte, Byte)
	/* X velocity of the player being pushed by the explosion */
	PlayerMotionX Float // Float
	/* Y velocity of the player being pushed by the explosion */
	PlayerMotionY Float // Float
	/* Z velocity of the player being pushed by the explosion */
	PlayerMotionZ Float // Float
}

var _ Packet = (*PlayExplosion_754_2)(nil)

func (p PlayExplosion_754_2)Encode(b *PacketBuilder){
	b.Float(p.X)
	b.Float(p.Y)
	b.Float(p.Z)
	b.Float(p.Strength)
	p.RecordCount = len(p.Records)
	b.Int(p.RecordCount)
	for _, v := range p.Records {
		b.Encode(v)
	}
	b.Float(p.PlayerMotionX)
	b.Float(p.PlayerMotionY)
	b.Float(p.PlayerMotionZ)
}

func (p *PlayExplosion_754_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Strength, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.RecordCount, ok = r.Int(); !ok {
		return io.EOF
	}
	p.Records = make([][3]byte, p.RecordCount)
	for i, _ := range p.Records {
		if err = p.Records[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if p.PlayerMotionX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.PlayerMotionY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.PlayerMotionZ, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x1c
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x1e
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x1e
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x1c
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x1c
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x1c
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x1c
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x1c
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x1c
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x1c
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x27
type PlayExplosion_498_3 struct {
	X Float // Float
	Y Float // Float
	Z Float // Float
	/* Currently unused in the client */
	Radius Float // Float
	/* Number of elements in the following array */
	RecordCount Int // Int
	/* Each record is 3 signed bytes long, each bytes are the XYZ (respectively) offsets of affected blocks. */
	Records [][3]byte // Array of (Byte, Byte, Byte)
	/* X velocity of the player being pushed by the explosion */
	PlayerMotionX Float // Float
	/* Y velocity of the player being pushed by the explosion */
	PlayerMotionY Float // Float
	/* Z velocity of the player being pushed by the explosion */
	PlayerMotionZ Float // Float
}

var _ Packet = (*PlayExplosion_498_3)(nil)

func (p PlayExplosion_498_3)Encode(b *PacketBuilder){
	b.Float(p.X)
	b.Float(p.Y)
	b.Float(p.Z)
	b.Float(p.Radius)
	p.RecordCount = len(p.Records)
	b.Int(p.RecordCount)
	for _, v := range p.Records {
		b.Encode(v)
	}
	b.Float(p.PlayerMotionX)
	b.Float(p.PlayerMotionY)
	b.Float(p.PlayerMotionZ)
}

func (p *PlayExplosion_498_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Radius, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.RecordCount, ok = r.Int(); !ok {
		return io.EOF
	}
	p.Records = make([][3]byte, p.RecordCount)
	for i, _ := range p.Records {
		if err = p.Records[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if p.PlayerMotionX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.PlayerMotionY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.PlayerMotionZ, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x37
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x37
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x37
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x33
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x33
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x35
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x34
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x31
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x31
type PlayFacePlayer_757_0 struct {
	/* Values are feet=0, eyes=1.  If set to eyes, aims using the head position; otherwise aims using the feet position. */
	FeetOrEyes VarInt // VarInt enum
	/* x coordinate of the point to face towards */
	TargetX Double // Double
	/* y coordinate of the point to face towards */
	TargetY Double // Double
	/* z coordinate of the point to face towards */
	TargetZ Double // Double
	/* If true, additional information about an entity is provided. */
	IsEntity Bool // Boolean
	/* Only if is entity is true — the entity to face towards */
	EntityID VarInt // Optional VarInt
	/* Whether to look at the entity's eyes or feet.  Same values and meanings as before, just for the entity's head/feet. */
	EntityFeetOrEyes VarInt // Optional VarInt enum
}

var _ Packet = (*PlayFacePlayer_757_0)(nil)

func (p PlayFacePlayer_757_0)Encode(b *PacketBuilder){
	b.VarInt(p.FeetOrEyes)
	b.Double(p.TargetX)
	b.Double(p.TargetY)
	b.Double(p.TargetZ)
	b.Bool(p.IsEntity)
	b.VarInt(p.EntityID)
	b.VarInt(p.EntityFeetOrEyes)
}

func (p *PlayFacePlayer_757_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.FeetOrEyes, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.TargetX, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.TargetY, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.TargetZ, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.IsEntity, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityFeetOrEyes, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x6b
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x6b
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x67
type PlayFeatureFlags_763_0 struct {
	/* Number of features that appear in the array below. */
	TotalFeatures VarInt // VarInt
	FeatureFlags []Identifier // Identifier Array
}

var _ Packet = (*PlayFeatureFlags_763_0)(nil)

func (p PlayFeatureFlags_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.TotalFeatures)
	TODO_Encode_Array(p.FeatureFlags)
}

func (p *PlayFeatureFlags_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TotalFeatures, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.FeatureFlags)
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x1f
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x1f
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x1c
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x1d
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x1b
type PlayGameEvent_763_0 struct {
	/* See below. */
	Event UByte // Unsigned Byte
	/* Depends on Event. */
	Value Float // Float
}

var _ Packet = (*PlayGameEvent_763_0)(nil)

func (p PlayGameEvent_763_0)Encode(b *PacketBuilder){
	b.UByte(p.Event)
	b.Float(p.Value)
}

func (p *PlayGameEvent_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Event, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Value, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0xe
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0xe
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0xe
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0xe
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0xf
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0xf
type PlayGenerateStructure_758_0 struct {
	/* Block entity location */
	Location Position // Position
	/* Value of the levels slider/max depth to generate */
	Levels VarInt // VarInt
	KeepJigsaws Bool // Boolean
}

var _ Packet = (*PlayGenerateStructure_758_0)(nil)

func (p PlayGenerateStructure_758_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.Levels)
	b.Bool(p.KeepJigsaws)
}

func (p *PlayGenerateStructure_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Levels, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.KeepJigsaws, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x48
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x48
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x3f
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x40
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x37
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x9
type PlayHeldItemChange_758_0 struct {
	/* The slot which the player has selected (0–8) */
	Slot Byte // Byte
}

var _ Packet = (*PlayHeldItemChange_758_0)(nil)

func (p PlayHeldItemChange_758_0)Encode(b *PacketBuilder){
	b.Byte(p.Slot)
}

func (p *PlayHeldItemChange_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x25
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x25
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x25
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x23
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x17
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x9
type PlayHeldItemChange_758_1 struct {
	/* The slot which the player has selected (0–8) */
	Slot Short // Short
}

var _ Packet = (*PlayHeldItemChange_758_1)(nil)

func (p PlayHeldItemChange_758_1)Encode(b *PacketBuilder){
	b.Short(p.Slot)
}

func (p *PlayHeldItemChange_758_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x48
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x48
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x3f
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x3f
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x3a
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x3a
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x37
type PlayHeldItemChange_756_2 struct {
	/* The slot which the player has selected (0–8) */
	Slot Byte // Byte
}

var _ Packet = (*PlayHeldItemChange_756_2)(nil)

func (p PlayHeldItemChange_756_2)Encode(b *PacketBuilder){
	b.Byte(p.Slot)
}

func (p *PlayHeldItemChange_756_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x25
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x25
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x25
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x23
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x1a
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x1a
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x17
type PlayHeldItemChange_756_3 struct {
	/* The slot which the player has selected (0–8) */
	Slot Short // Short
}

var _ Packet = (*PlayHeldItemChange_756_3)(nil)

func (p PlayHeldItemChange_756_3)Encode(b *PacketBuilder){
	b.Short(p.Slot)
}

func (p *PlayHeldItemChange_756_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x3d
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x3d
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x39
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x37
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x37
type PlayHeldItemChange_404_4 struct {
	/* The slot which the player has selected (0–8) */
	Slot Byte // Byte
}

var _ Packet = (*PlayHeldItemChange_404_4)(nil)

func (p PlayHeldItemChange_404_4)Encode(b *PacketBuilder){
	b.Byte(p.Slot)
}

func (p *PlayHeldItemChange_404_4)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x21
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x21
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x1a
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x17
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x17
type PlayHeldItemChange_404_5 struct {
	/* The slot which the player has selected (0–8) */
	Slot Short // Short
}

var _ Packet = (*PlayHeldItemChange_404_5)(nil)

func (p PlayHeldItemChange_404_5)Encode(b *PacketBuilder){
	b.Short(p.Slot)
}

func (p *PlayHeldItemChange_404_5)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x22
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x22
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x1e
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x1f
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x1d
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x20
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x20
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x20
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x20
type PlayInitializeWorldBorder_763_0 struct {
	X Double // Double
	Z Double // Double
	/* Current length of a single side of the world border, in meters. */
	OldDiameter Double // Double
	/* Target length of a single side of the world border, in meters. */
	NewDiameter Double // Double
	/* Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. */
	Speed VarLong // VarLong
	/* Resulting coordinates from a portal teleport are limited to ±value. Usually 29999984. */
	PortalTeleportBoundary VarInt // VarInt
	/* In meters. */
	WarningBlocks VarInt // VarInt
	/* In seconds as set by /worldborder warning time. */
	WarningTime VarInt // VarInt
}

var _ Packet = (*PlayInitializeWorldBorder_763_0)(nil)

func (p PlayInitializeWorldBorder_763_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Z)
	b.Double(p.OldDiameter)
	b.Double(p.NewDiameter)
	b.VarLong(p.Speed)
	b.VarInt(p.PortalTeleportBoundary)
	b.VarInt(p.WarningBlocks)
	b.VarInt(p.WarningTime)
}

func (p *PlayInitializeWorldBorder_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.OldDiameter, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.NewDiameter, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Speed, ok = r.VarLong(); !ok {
		return io.EOF
	}
	if p.PortalTeleportBoundary, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.WarningBlocks, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.WarningTime, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x10
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x10
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0xf
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x10
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0xf
type PlayInteract_763_0 struct {
	/* The ID of the entity to interact. */
	EntityID VarInt // VarInt
	/* 0: interact, 1: attack, 2: interact at. */
	Type VarInt // VarInt Enum
	/* Only if Type is interact at. */
	TargetX Float // Optional Float
	/* Only if Type is interact at. */
	TargetY Float // Optional Float
	/* Only if Type is interact at. */
	TargetZ Float // Optional Float
	/* Only if Type is interact or interact at; 0: main hand, 1: off hand. */
	Hand VarInt // Optional VarInt Enum
	/* If the client is sneaking. */
	Sneaking Bool // Boolean
}

var _ Packet = (*PlayInteract_763_0)(nil)

func (p PlayInteract_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.Type)
	b.Float(p.TargetX)
	b.Float(p.TargetY)
	b.Float(p.TargetZ)
	b.VarInt(p.Hand)
	b.Bool(p.Sneaking)
}

func (p *PlayInteract_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.TargetX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TargetY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TargetZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Sneaking, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0xd
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0xd
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0xd
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0xd
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0xe
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0xe
type PlayInteractEntity_758_0 struct {
	/* The ID of the entity to interact */
	EntityID VarInt // VarInt
	/* 0: interact, 1: attack, 2: interact at */
	Type VarInt // VarInt Enum
	/* Only if Type is interact at */
	TargetX Float // Optional Float
	/* Only if Type is interact at */
	TargetY Float // Optional Float
	/* Only if Type is interact at */
	TargetZ Float // Optional Float
	/* Only if Type is interact or interact at; 0: main hand, 1: off hand */
	Hand VarInt // Optional VarInt Enum
	/* If the client is sneaking. */
	Sneaking Bool // Boolean
}

var _ Packet = (*PlayInteractEntity_758_0)(nil)

func (p PlayInteractEntity_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.Type)
	b.Float(p.TargetX)
	b.Float(p.TargetY)
	b.Float(p.TargetZ)
	b.VarInt(p.Hand)
	b.Bool(p.Sneaking)
}

func (p *PlayInteractEntity_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.TargetX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TargetY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TargetZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Sneaking, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0xe
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0xe
type PlayInteractEntity_578_1 struct {
	/* The ID of the entity to interact */
	EntityID VarInt // VarInt
	/* 0: interact, 1: attack, 2: interact at */
	Type VarInt // VarInt Enum
	/* Only if Type is interact at */
	TargetX Float // Optional Float
	/* Only if Type is interact at */
	TargetY Float // Optional Float
	/* Only if Type is interact at */
	TargetZ Float // Optional Float
	/* Only if Type is interact or interact at; 0: main hand, 1: off hand */
	Hand VarInt // Optional VarInt Enum
}

var _ Packet = (*PlayInteractEntity_578_1)(nil)

func (p PlayInteractEntity_578_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.Type)
	b.Float(p.TargetX)
	b.Float(p.TargetY)
	b.Float(p.TargetZ)
	b.VarInt(p.Hand)
}

func (p *PlayInteractEntity_578_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.TargetX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TargetY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TargetZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x11
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x11
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x10
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x11
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x10
type PlayJigsawGenerate_763_0 struct {
	/* Block entity location. */
	Location Position // Position
	/* Value of the levels slider/max depth to generate. */
	Levels VarInt // VarInt
	KeepJigsaws Bool // Boolean
}

var _ Packet = (*PlayJigsawGenerate_763_0)(nil)

func (p PlayJigsawGenerate_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.Levels)
	b.Bool(p.KeepJigsaws)
}

func (p *PlayJigsawGenerate_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Levels, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.KeepJigsaws, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x26
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x26
type PlayJoinGame_758_0 struct {
	/* The player's Entity ID (EID). */
	EntityID Int // Int
	IsHardcore Bool // Boolean
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. */
	Gamemode UByte // Unsigned Byte
	/* 0: survival, 1: creative, 2: adventure, 3: spectator. The hardcore flag is not included. The previous gamemode. Defaults to -1 if there is no previous gamemode. (More information needed) */
	PreviousGamemode Byte // Byte
	/* Size of the following array. */
	WorldCount VarInt // VarInt
	/* Identifiers for all dimensions on the server. */
	DimensionNames []Identifier // Array of Identifier
	/* The full extent of these is still unknown, but the tag represents a dimension and biome registry. See below for the vanilla default. */
	DimensionCodec *NBTCompound // NBT Tag Compound
	/* Valid dimensions are defined per dimension registry sent before this. The structure of this tag is a dimension type (see below). */
	Dimension *NBTCompound // NBT Tag Compound
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
}

var _ Packet = (*PlayJoinGame_758_0)(nil)

func (p PlayJoinGame_758_0)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.Bool(p.IsHardcore)
	b.UByte(p.Gamemode)
	b.Byte(p.PreviousGamemode)
	b.VarInt(p.WorldCount)
	TODO_Encode_Array(p.DimensionNames)
	b.Encode(p.DimensionCodec)
	b.Encode(p.Dimension)
	b.Encode(p.DimensionName)
	b.Long(p.HashedSeed)
	b.VarInt(p.MaxPlayers)
	b.VarInt(p.ViewDistance)
	b.VarInt(p.SimulationDistance)
	b.Bool(p.ReducedDebugInfo)
	b.Bool(p.EnableRespawnScreen)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
}

func (p *PlayJoinGame_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.IsHardcore, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.PreviousGamemode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.WorldCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.DimensionNames)
	if err = p.DimensionCodec.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Dimension.DecodeFrom(r); err != nil {
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
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x26
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x26
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x24
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x24
type PlayJoinGame_756_1 struct {
	/* The player's Entity ID (EID) */
	EntityID Int // Int
	IsHardcore Bool // Boolean
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. */
	Gamemode UByte // Unsigned Byte
	/* 0: survival, 1: creative, 2: adventure, 3: spectator. The hardcore flag is not included. The previous gamemode. Defaults to -1 if there is no previous gamemode. (More information needed) */
	PreviousGamemode Byte // Byte
	/* Size of the following array */
	WorldCount VarInt // VarInt
	/* Identifiers for all worlds on the server */
	WorldNames []Identifier // Array of Identifier
	/* The full extent of these is still unknown, but the tag represents a dimension and biome registry. See below for the vanilla default. */
	DimensionCodec *NBTCompound // NBT Tag Compound
	/* Valid dimensions are defined per dimension registry sent before this */
	Dimension *NBTCompound // NBT Tag Compound
	/* Name of the world being spawned into */
	WorldName Identifier // Identifier
	/* First 8 bytes of the SHA-256 hash of the world's seed. */
	HashedSeed Long // Long
	/* Was once used by the client to draw the player list, but now is ignored */
	MaxPlayers VarInt // VarInt
	/* Render distance (2-32) */
	ViewDistance VarInt // VarInt
	/* If true, a Notchian client shows reduced information on the debug screen.  For servers in development, this should almost always be false. */
	ReducedDebugInfo Bool // Boolean
	/* Set to false when the doImmediateRespawn gamerule is true */
	EnableRespawnScreen Bool // Boolean
	/* True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks */
	IsDebug Bool // Boolean
	/* True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63 */
	IsFlat Bool // Boolean
}

var _ Packet = (*PlayJoinGame_756_1)(nil)

func (p PlayJoinGame_756_1)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.Bool(p.IsHardcore)
	b.UByte(p.Gamemode)
	b.Byte(p.PreviousGamemode)
	p.WorldCount = len(p.WorldNames)
	b.VarInt(p.WorldCount)
	for _, v := range p.WorldNames {
		b.Encode(v)
	}
	b.Encode(p.DimensionCodec)
	b.Encode(p.Dimension)
	b.Encode(p.WorldName)
	b.Long(p.HashedSeed)
	b.VarInt(p.MaxPlayers)
	b.VarInt(p.ViewDistance)
	b.Bool(p.ReducedDebugInfo)
	b.Bool(p.EnableRespawnScreen)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
}

func (p *PlayJoinGame_756_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.IsHardcore, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.PreviousGamemode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.WorldCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.WorldNames = make([]Identifier, p.WorldCount)
	for i, _ := range p.WorldNames {
		if err = p.WorldNames[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if err = p.DimensionCodec.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Dimension.DecodeFrom(r); err != nil {
		return
	}
	if err = p.WorldName.DecodeFrom(r); err != nil {
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
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x25
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x25
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x23
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x23
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x23
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x23
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x23
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x23
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x23
type PlayJoinGame_404_4 struct {
	/* The player's Entity ID (EID) */
	EntityID Int // Int
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. Bit 3 (0x8) is the hardcore flag. */
	Gamemode UByte // Unsigned Byte
	/* -1: Nether, 0: Overworld, 1: End; also, note that this is not a VarInt but instead a regular int. */
	Dimension Int // Int Enum
	/* 0: peaceful, 1: easy, 2: normal, 3: hard */
	Difficulty UByte // Unsigned Byte
	/* Was once used by the client to draw the player list, but now is ignored */
	MaxPlayers UByte // Unsigned Byte
	/* default, flat, largeBiomes, amplified, default_1_1 */
	LevelType String // String Enum (16)
	/* If true, a Notchian client shows reduced information on the debug screen.  For servers in development, this should almost always be false. */
	ReducedDebugInfo Bool // Boolean
}

var _ Packet = (*PlayJoinGame_404_4)(nil)

func (p PlayJoinGame_404_4)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.UByte(p.Gamemode)
	b.Int(p.Dimension)
	b.UByte(p.Difficulty)
	b.UByte(p.MaxPlayers)
	b.String(p.LevelType)
	b.Bool(p.ReducedDebugInfo)
}

func (p *PlayJoinGame_404_4)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Dimension, ok = r.Int(); !ok {
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
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x23
// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x12
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x23
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x12
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x1f
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x11
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x20
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x12
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x1e
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x11
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x21
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0xf
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x21
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0xf
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x21
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0xf
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x21
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0xf
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x1f
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x10
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x1f
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x10
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x21
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0xf
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x20
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0xf
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x21
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0xe
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x21
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0xe
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x1f
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0xb
type PlayKeepAlive_763_0 struct {
	KeepAliveID Long // Long
}

var _ Packet = (*PlayKeepAlive_763_0)(nil)

func (p PlayKeepAlive_763_0)Encode(b *PacketBuilder){
	b.Long(p.KeepAliveID)
}

func (p *PlayKeepAlive_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.KeepAliveID, ok = r.Long(); !ok {
		return io.EOF
	}
}

// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x1f
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0xb
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x1f
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0xc
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x1f
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0xb
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x1f
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0xb
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x1f
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0xb
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x1f
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0xb
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x0
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x0
type PlayKeepAlive_338_1 struct {
	KeepAliveID VarInt // VarInt
}

var _ Packet = (*PlayKeepAlive_338_1)(nil)

func (p PlayKeepAlive_338_1)Encode(b *PacketBuilder){
	b.VarInt(p.KeepAliveID)
}

func (p *PlayKeepAlive_338_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.KeepAliveID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x53
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x53
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x4f
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x51
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x4e
type PlayLinkEntities_763_0 struct {
	/* Attached entity's EID. */
	AttachedEntityID Int // Int
	/* ID of the entity holding the lead. Set to -1 to detach. */
	HoldingEntityID Int // Int
}

var _ Packet = (*PlayLinkEntities_763_0)(nil)

func (p PlayLinkEntities_763_0)Encode(b *PacketBuilder){
	b.Int(p.AttachedEntityID)
	b.Int(p.HoldingEntityID)
}

func (p *PlayLinkEntities_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.AttachedEntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.HoldingEntityID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x13
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x13
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x12
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x13
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x12
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x10
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x10
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x10
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x10
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x11
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x11
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x10
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x10
type PlayLockDifficulty_763_0 struct {
	Locked Bool // Boolean
}

var _ Packet = (*PlayLockDifficulty_763_0)(nil)

func (p PlayLockDifficulty_763_0)Encode(b *PacketBuilder){
	b.Bool(p.Locked)
}

func (p *PlayLockDifficulty_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Locked, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x24
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x25
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x23
type PlayLogin_761_2 struct {
	/* The player's Entity ID (EID). */
	EntityID Int // Int
	IsHardcore Bool // Boolean
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. */
	Gamemode UByte // Unsigned Byte
	/* -1: Undefined (null), 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. The previous gamemode. Vanilla client uses this for the debug (F3 + N & F3 + F4) gamemode switch. (More information needed) */
	PreviousGamemode Byte // Byte
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

var _ Packet = (*PlayLogin_761_2)(nil)

func (p PlayLogin_761_2)Encode(b *PacketBuilder){
	b.Int(p.EntityID)
	b.Bool(p.IsHardcore)
	b.UByte(p.Gamemode)
	b.Byte(p.PreviousGamemode)
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

func (p *PlayLogin_761_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.IsHardcore, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.PreviousGamemode, ok = r.Byte(); !ok {
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

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x3b
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x3b
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x37
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x38
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x35
type PlayLookAt_763_0 struct {
	/* Values are feet=0, eyes=1.  If set to eyes, aims using the head position; otherwise aims using the feet position. */
	FeetOrEyes VarInt // VarInt Enum
	/* x coordinate of the point to face towards. */
	TargetX Double // Double
	/* y coordinate of the point to face towards. */
	TargetY Double // Double
	/* z coordinate of the point to face towards. */
	TargetZ Double // Double
	/* If true, additional information about an entity is provided. */
	IsEntity Bool // Boolean
	/* Only if is entity is true — the entity to face towards. */
	EntityID VarInt // Optional VarInt
	/* Whether to look at the entity's eyes or feet.  Same values and meanings as before, just for the entity's head/feet. */
	EntityFeetOrEyes VarInt // Optional VarInt Enum
}

var _ Packet = (*PlayLookAt_763_0)(nil)

func (p PlayLookAt_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.FeetOrEyes)
	b.Double(p.TargetX)
	b.Double(p.TargetY)
	b.Double(p.TargetZ)
	b.Bool(p.IsEntity)
	b.VarInt(p.EntityID)
	b.VarInt(p.EntityFeetOrEyes)
}

func (p *PlayLookAt_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.FeetOrEyes, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.TargetX, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.TargetY, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.TargetZ, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.IsEntity, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityFeetOrEyes, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x3
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x3
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x3
type PlayMessageAcknowledgment_763_0 struct {
	MessageCount VarInt // VarInt
}

var _ Packet = (*PlayMessageAcknowledgment_763_0)(nil)

func (p PlayMessageAcknowledgment_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.MessageCount)
}

func (p *PlayMessageAcknowledgment_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.MessageCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x2e
// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x18
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x2e
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x18
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x2a
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x17
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x2b
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x18
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x29
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x17
type PlayMoveVehicle_763_0 struct {
	/* Absolute position (X coordinate). */
	X Double // Double
	/* Absolute position (Y coordinate). */
	Y Double // Double
	/* Absolute position (Z coordinate). */
	Z Double // Double
	/* Absolute rotation on the vertical axis, in degrees. */
	Yaw Float // Float
	/* Absolute rotation on the horizontal axis, in degrees. */
	Pitch Float // Float
}

var _ Packet = (*PlayMoveVehicle_763_0)(nil)

func (p PlayMoveVehicle_763_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
}

func (p *PlayMoveVehicle_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x3f
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x3f
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x3f
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x3f
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x3b
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x3b
type PlayMultiBlockChange_758_0 struct {
	/* Chunk section coordinate (encoded chunk x and z with each 22 bits, and section y with 20 bits, from left to right) */
	ChunkSectionPosition Long // Long
	/* Number of elements in the following array */
	BlocksArraySize VarInt // VarInt
	/* Each entry is composed of the block id, shifted right by 12, and the relative block position in the chunk section (4 bits for x, z, and y, from left to right). */
	Blocks []VarLong // Array of VarLong
}

var _ Packet = (*PlayMultiBlockChange_758_0)(nil)

func (p PlayMultiBlockChange_758_0)Encode(b *PacketBuilder){
	b.Long(p.ChunkSectionPosition)
	b.VarInt(p.BlocksArraySize)
	TODO_Encode_Array(p.Blocks)
}

func (p *PlayMultiBlockChange_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkSectionPosition, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.BlocksArraySize, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.Blocks)
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x60
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x60
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x5f
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x5f
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x54
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x54
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x55
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x54
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x1d
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x1d
type PlayNBTQueryResponse_758_0 struct {
	/* Can be compared to the one sent in the original query packet. */
	TransactionID VarInt // VarInt
	/* The NBT of the block or entity.  May be a TAG_END (0) in which case no NBT is present. */
	NBT NBT // NBT Tag
}

var _ Packet = (*PlayNBTQueryResponse_758_0)(nil)

func (p PlayNBTQueryResponse_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionID)
	b.Encode(p.NBT)
}

func (p *PlayNBTQueryResponse_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.NBT.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x20
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x20
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x20
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x20
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x20
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x20
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x1e
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x1e
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x1c
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x1c
type PlayNameItem_758_0 struct {
	/* The new name of the item */
	ItemName String // String (32767)
}

var _ Packet = (*PlayNameItem_758_0)(nil)

func (p PlayNameItem_758_0)Encode(b *PacketBuilder){
	b.String(p.ItemName)
}

func (p *PlayNameItem_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ItemName, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x19
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x19
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x19
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x19
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x18
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x18
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x1a
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x19
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x1a
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x1a
type PlayNamedSoundEffect_758_0 struct {
	/* All sound effect names as of 1.15.2 can be seen here. */
	SoundName Identifier // Identifier
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
	/* Float between 0.5 and 2.0 by Notchian clients */
	Pitch Float // Float
}

var _ Packet = (*PlayNamedSoundEffect_758_0)(nil)

func (p PlayNamedSoundEffect_758_0)Encode(b *PacketBuilder){
	b.Encode(p.SoundName)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.Float(p.Pitch)
}

func (p *PlayNamedSoundEffect_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.SoundName.DecodeFrom(r); err != nil {
		return
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
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x19
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x19
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x19
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x19
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x19
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x19
type PlayNamedSoundEffect_340_1 struct {
	/* All sound effect names as of 1.11.0 can be seen here. */
	SoundName String // String (256)
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
	/* Float between 0.5 and 2.0 by Notchian clients */
	Pitch Float // Float
}

var _ Packet = (*PlayNamedSoundEffect_340_1)(nil)

func (p PlayNamedSoundEffect_340_1)Encode(b *PacketBuilder){
	b.String(p.SoundName)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.Float(p.Pitch)
}

func (p *PlayNamedSoundEffect_340_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x2f
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x2f
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x2b
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x2c
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x2a
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x2d
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x2d
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x2d
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x2d
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x2c
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x2c
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x2e
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x2d
type PlayOpenBook_763_0 struct {
	/* 0: Main hand, 1: Off hand */
	Hand VarInt // VarInt enum
}

var _ Packet = (*PlayOpenBook_763_0)(nil)

func (p PlayOpenBook_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
}

func (p *PlayOpenBook_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x20
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x20
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x1d
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x1e
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x1c
type PlayOpenHorseScreen_763_0 struct {
	WindowID UByte // Unsigned Byte
	SlotCount VarInt // VarInt
	EntityID Int // Int
}

var _ Packet = (*PlayOpenHorseScreen_763_0)(nil)

func (p PlayOpenHorseScreen_763_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.VarInt(p.SlotCount)
	b.Int(p.EntityID)
}

func (p *PlayOpenHorseScreen_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.SlotCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x1f
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x1f
type PlayOpenHorseWindow_758_0 struct {
	WindowID UByte // Unsigned Byte
	SlotCount VarInt // VarInt
	EntityID Int // Integer
}

var _ Packet = (*PlayOpenHorseWindow_758_0)(nil)

func (p PlayOpenHorseWindow_758_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.VarInt(p.SlotCount)
	b.Int(p.EntityID)
}

func (p *PlayOpenHorseWindow_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.SlotCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x1f
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x1f
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x1e
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x1e
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x20
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x1f
type PlayOpenHorseWindow_756_1 struct {
	WindowID Byte // Byte
	NumberOfSlots VarInt // VarInt
	EntityID Int // Integer
}

var _ Packet = (*PlayOpenHorseWindow_756_1)(nil)

func (p PlayOpenHorseWindow_756_1)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.VarInt(p.NumberOfSlots)
	b.Int(p.EntityID)
}

func (p *PlayOpenHorseWindow_756_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.NumberOfSlots, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x30
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x30
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x2c
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x2d
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x2b
type PlayOpenScreen_763_0 struct {
	/* A unique id number for the window to be displayed. Notchian server implementation is a counter, starting at 1. */
	WindowID VarInt // VarInt
	/* The window type to use for display. Contained in the minecraft:menu registry; see Inventory for the different values. */
	WindowType VarInt // VarInt
	/* The title of the window. */
	WindowTitle Object // Chat
}

var _ Packet = (*PlayOpenScreen_763_0)(nil)

func (p PlayOpenScreen_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.WindowID)
	b.VarInt(p.WindowType)
	b.JSON(p.WindowTitle)
}

func (p *PlayOpenScreen_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.WindowType, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.WindowTitle); err != nil {
		return
	}
}

// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x31
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x2d
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x2e
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x2c
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x2f
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x2f
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x2f
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x2f
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x2e
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x2e
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x30
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x2f
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x2c
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x2c
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x2a
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x2a
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x2a
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x2a
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x2a
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x2a
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x2a
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x36
type PlayOpenSignEditor_762_1 struct {
	Location Position // Position
}

var _ Packet = (*PlayOpenSignEditor_762_1)(nil)

func (p PlayOpenSignEditor_762_1)Encode(b *PacketBuilder){
	b.Encode(p.Location)
}

func (p *PlayOpenSignEditor_762_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x2e
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x2e
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x2e
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x2e
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x2d
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x2d
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x2f
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x2e
type PlayOpenWindow_758_0 struct {
	/* A unique id number for the window to be displayed. Notchian server implementation is a counter, starting at 1. */
	WindowID VarInt // VarInt
	/* The window type to use for display. Contained in the minecraft:menu regisry; see Inventory for the different values. */
	WindowType VarInt // VarInt
	/* The title of the window */
	WindowTitle Object // Chat
}

var _ Packet = (*PlayOpenWindow_758_0)(nil)

func (p PlayOpenWindow_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.WindowID)
	b.VarInt(p.WindowType)
	b.JSON(p.WindowTitle)
}

func (p *PlayOpenWindow_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.WindowType, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.WindowTitle); err != nil {
		return
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x14
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x14
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x13
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x13
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x13
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x13
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x13
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x13
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x13
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x2d
type PlayOpenWindow_404_1 struct {
	/* A unique id number for the window to be displayed. Notchian server implementation is a counter, starting at 1. */
	WindowID UByte // Unsigned Byte
	/* The window type to use for display. See Inventory for a list. */
	WindowType String // String
	/* The title of the window */
	WindowTitle Object // Chat
	/* Number of slots in the window (excluding the number of slots in the player inventory) */
	NumberOfSlots UByte // Unsigned Byte
	/* EntityHorse's EID. Only sent when Window Type is “EntityHorse” */
	EntityID Int // Optional Int
}

var _ Packet = (*PlayOpenWindow_404_1)(nil)

func (p PlayOpenWindow_404_1)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.String(p.WindowType)
	b.JSON(p.WindowTitle)
	b.UByte(p.NumberOfSlots)
	b.Int(p.EntityID)
}

func (p *PlayOpenWindow_404_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.WindowType, ok = r.String(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.WindowTitle); err != nil {
		return
	}
	if p.NumberOfSlots, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x19
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x19
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x18
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x19
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x18
type PlayPaddleBoat_763_0 struct {
	LeftPaddleTurning Bool // Boolean
	RightPaddleTurning Bool // Boolean
}

var _ Packet = (*PlayPaddleBoat_763_0)(nil)

func (p PlayPaddleBoat_763_0)Encode(b *PacketBuilder){
	b.Bool(p.LeftPaddleTurning)
	b.Bool(p.RightPaddleTurning)
}

func (p *PlayPaddleBoat_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.LeftPaddleTurning, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.RightPaddleTurning, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x26
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x26
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x22
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x23
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x21
type PlayParticle_763_0 struct {
	/* The particle ID listed in the particle data type. */
	ParticleID VarInt // VarInt
	/* If true, particle distance increases from 256 to 65536. */
	LongDistance Bool // Boolean
	/* X position of the particle. */
	X Double // Double
	/* Y position of the particle. */
	Y Double // Double
	/* Z position of the particle. */
	Z Double // Double
	/* This is added to the X position after being multiplied by random.nextGaussian(). */
	OffsetX Float // Float
	/* This is added to the Y position after being multiplied by random.nextGaussian(). */
	OffsetY Float // Float
	/* This is added to the Z position after being multiplied by random.nextGaussian(). */
	OffsetZ Float // Float
	MaxSpeed Float // Float
	/* The number of particles to create. */
	ParticleCount Int // Int
	/* The variable data listed in the particle data type. */
	Data any // Varies
}

var _ Packet = (*PlayParticle_763_0)(nil)

func (p PlayParticle_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.ParticleID)
	b.Bool(p.LongDistance)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.OffsetX)
	b.Float(p.OffsetY)
	b.Float(p.OffsetZ)
	b.Float(p.MaxSpeed)
	b.Int(p.ParticleCount)
	b.Encode(p.Data)
}

func (p *PlayParticle_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ParticleID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.LongDistance, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.OffsetX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.MaxSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.ParticleCount, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x24
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x24
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x24
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x24
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x22
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x22
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x24
type PlayParticle_758_1 struct {
	/* The particle ID listed in the particle data type. */
	ParticleID Int // Int
	/* If true, particle distance increases from 256 to 65536 */
	LongDistance Bool // Boolean
	/* X position of the particle */
	X Double // Double
	/* Y position of the particle */
	Y Double // Double
	/* Z position of the particle */
	Z Double // Double
	/* This is added to the X position after being multiplied by random.nextGaussian() */
	OffsetX Float // Float
	/* This is added to the Y position after being multiplied by random.nextGaussian() */
	OffsetY Float // Float
	/* This is added to the Z position after being multiplied by random.nextGaussian() */
	OffsetZ Float // Float
	/* The data of each particle */
	ParticleData Float // Float
	/* The number of particles to create */
	ParticleCount Int // Int
	/* The variable data listed in the particle data type. */
	Data any // Varies
}

var _ Packet = (*PlayParticle_758_1)(nil)

func (p PlayParticle_758_1)Encode(b *PacketBuilder){
	b.Int(p.ParticleID)
	b.Bool(p.LongDistance)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.OffsetX)
	b.Float(p.OffsetY)
	b.Float(p.OffsetZ)
	b.Float(p.ParticleData)
	b.Int(p.ParticleCount)
	b.Encode(p.Data)
}

func (p *PlayParticle_758_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ParticleID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.LongDistance, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.OffsetX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.ParticleData, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.ParticleCount, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x23
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x24
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x24
type PlayParticle_498_2 struct {
	/* The particle ID listed in the particle data type. */
	ParticleID Int // Int
	/* If true, particle distance increases from 256 to 65536 */
	LongDistance Bool // Boolean
	/* X position of the particle */
	X Float // Float
	/* Y position of the particle */
	Y Float // Float
	/* Z position of the particle */
	Z Float // Float
	/* This is added to the X position after being multiplied by random.nextGaussian() */
	OffsetX Float // Float
	/* This is added to the Y position after being multiplied by random.nextGaussian() */
	OffsetY Float // Float
	/* This is added to the Z position after being multiplied by random.nextGaussian() */
	OffsetZ Float // Float
	/* The data of each particle */
	ParticleData Float // Float
	/* The number of particles to create */
	ParticleCount Int // Int
	/* The variable data listed in the particle data type. */
	Data any // Varies
}

var _ Packet = (*PlayParticle_498_2)(nil)

func (p PlayParticle_498_2)Encode(b *PacketBuilder){
	b.Int(p.ParticleID)
	b.Bool(p.LongDistance)
	b.Float(p.X)
	b.Float(p.Y)
	b.Float(p.Z)
	b.Float(p.OffsetX)
	b.Float(p.OffsetY)
	b.Float(p.OffsetZ)
	b.Float(p.ParticleData)
	b.Int(p.ParticleCount)
	b.Encode(p.Data)
}

func (p *PlayParticle_498_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ParticleID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.LongDistance, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.ParticleData, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.ParticleCount, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x22
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x22
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x22
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x22
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x22
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x22
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x22
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x2a
type PlayParticle_340_3 struct {
	/* See below */
	ParticleID Int // Int
	/* If true, particle distance increases from 256 to 65536 */
	LongDistance Bool // Boolean
	/* X position of the particle */
	X Float // Float
	/* Y position of the particle */
	Y Float // Float
	/* Z position of the particle */
	Z Float // Float
	/* This is added to the X position after being multiplied by random.nextGaussian() */
	OffsetX Float // Float
	/* This is added to the Y position after being multiplied by random.nextGaussian() */
	OffsetY Float // Float
	/* This is added to the Z position after being multiplied by random.nextGaussian() */
	OffsetZ Float // Float
	/* The data of each particle */
	ParticleData Float // Float
	/* The number of particles to create */
	ParticleCount Int // Int
	/* Length depends on particle. "iconcrack" has length of 2, "blockcrack", and "blockdust" have lengths of 1, the rest have 0. */
	Data []VarInt // Array of VarInt
}

var _ Packet = (*PlayParticle_340_3)(nil)

func (p PlayParticle_340_3)Encode(b *PacketBuilder){
	b.Int(p.ParticleID)
	b.Bool(p.LongDistance)
	b.Float(p.X)
	b.Float(p.Y)
	b.Float(p.Z)
	b.Float(p.OffsetX)
	b.Float(p.OffsetY)
	b.Float(p.OffsetZ)
	b.Float(p.ParticleData)
	b.Int(p.ParticleCount)
	TODO_Encode_Array(p.Data)
}

func (p *PlayParticle_340_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ParticleID, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.LongDistance, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OffsetZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.ParticleData, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.ParticleCount, ok = r.Int(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.Data)
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x1a
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x1a
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x19
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x1a
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x19
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x17
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x17
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x17
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x17
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x18
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x18
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x17
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x17
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x15
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x15
type PlayPickItem_763_0 struct {
	/* See Inventory */
	SlotToUse VarInt // VarInt
}

var _ Packet = (*PlayPickItem_763_0)(nil)

func (p PlayPickItem_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.SlotToUse)
}

func (p *PlayPickItem_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SlotToUse, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x67
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x67
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x63
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x65
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x62
type PlayPickupItem_763_0 struct {
	CollectedEntityID VarInt // VarInt
	CollectorEntityID VarInt // VarInt
	/* Seems to be 1 for XP orbs, otherwise the number of items in the stack. */
	PickupItemCount VarInt // VarInt
}

var _ Packet = (*PlayPickupItem_763_0)(nil)

func (p PlayPickupItem_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.CollectedEntityID)
	b.VarInt(p.CollectorEntityID)
	b.VarInt(p.PickupItemCount)
}

func (p *PlayPickupItem_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.CollectedEntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CollectorEntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.PickupItemCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x32
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x32
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x2e
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x2f
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x2d
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x30
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x30
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x30
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x30
type PlayPing_763_0 struct {
	ID Int // Int
}

var _ Packet = (*PlayPing_763_0)(nil)

func (p PlayPing_763_0)Encode(b *PacketBuilder){
	b.Int(p.ID)
}

func (p *PlayPing_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x33
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x33
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x2f
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x30
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x2e
type PlayPlaceGhostRecipe_763_0 struct {
	WindowID Byte // Byte
	/* A recipe ID. */
	Recipe Identifier // Identifier
}

var _ Packet = (*PlayPlaceGhostRecipe_763_0)(nil)

func (p PlayPlaceGhostRecipe_763_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Encode(p.Recipe)
}

func (p *PlayPlaceGhostRecipe_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if err = p.Recipe.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x1b
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x1b
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x1a
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x1b
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x1a
type PlayPlaceRecipe_763_0 struct {
	WindowID Byte // Byte
	/* A recipe ID. */
	Recipe Identifier // Identifier
	/* Affects the amount of items processed; true if shift is down when clicked. */
	MakeAll Bool // Boolean
}

var _ Packet = (*PlayPlaceRecipe_763_0)(nil)

func (p PlayPlaceRecipe_763_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Encode(p.Recipe)
	b.Bool(p.MakeAll)
}

func (p *PlayPlaceRecipe_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if err = p.Recipe.DecodeFrom(r); err != nil {
		return
	}
	if p.MakeAll, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0xf
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0xf
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0xc
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0xc
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0xd
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0xf
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0xf
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0xf
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0xf
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x3
type PlayPlayer_404_0 struct {
	/* True if the client is on the ground, false otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayer_404_0)(nil)

func (p PlayPlayer_404_0)Encode(b *PacketBuilder){
	b.Bool(p.OnGround)
}

func (p *PlayPlayer_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x34
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x34
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x30
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x30
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x32
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x2b
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x39
type PlayPlayerAbilities_763_0 struct {
	/* Bit field, see below */
	Flags Byte // Byte
	FlyingSpeed Float // Float
	/* Modifies the field of view, like a speed potion. A Notchian server will use the same value as the movement speed (send in the Entity Properties packet). */
	FieldOfViewModifier Float // Float
}

var _ Packet = (*PlayPlayerAbilities_763_0)(nil)

func (p PlayPlayerAbilities_763_0)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.Float(p.FlyingSpeed)
	b.Float(p.FieldOfViewModifier)
}

func (p *PlayPlayerAbilities_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.FlyingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.FieldOfViewModifier, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x1c
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x1c
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x1b
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x1a
type PlayPlayerAbilities_763_1 struct {
	/* Bit mask. 0x02: is flying */
	Flags Byte // Byte
}

var _ Packet = (*PlayPlayerAbilities_763_1)(nil)

func (p PlayPlayerAbilities_763_1)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
}

func (p *PlayPlayerAbilities_763_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x31
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x2f
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x32
type PlayPlayerAbilities_760_2 struct {
	/* Bit field, see below. */
	Flags Byte // Byte
	/* 0.05 by default. */
	FlyingSpeed Float // Float
	/* Modifies the field of view, like a speed potion. A Notchian server will use the same value as the movement speed sent in the Entity Properties packet, which defaults to 0.1 for players. */
	FieldOfViewModifier Float // Float
}

var _ Packet = (*PlayPlayerAbilities_760_2)(nil)

func (p PlayPlayerAbilities_760_2)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.Float(p.FlyingSpeed)
	b.Float(p.FieldOfViewModifier)
}

func (p *PlayPlayerAbilities_760_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.FlyingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.FieldOfViewModifier, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x1c
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x1b
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x19
type PlayPlayerAbilities_760_3 struct {
	/* Bit mask. 0x02: is flying. */
	Flags Byte // Byte
}

var _ Packet = (*PlayPlayerAbilities_760_3)(nil)

func (p PlayPlayerAbilities_760_3)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
}

func (p *PlayPlayerAbilities_760_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x32
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x32
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x30
type PlayPlayerAbilities_758_4 struct {
	/* Bit field, see below. */
	Flags Byte // Byte
	/* 0.05 by default. */
	FlyingSpeed Float // Float
	/* Modifies the field of view, like a speed potion. A Notchian server will use the same value as the movement speed sent in the Entity Properties packet, which defaults to 0.1 for players. */
	FieldOfViewModifier Float // Float
}

var _ Packet = (*PlayPlayerAbilities_758_4)(nil)

func (p PlayPlayerAbilities_758_4)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.Float(p.FlyingSpeed)
	b.Float(p.FieldOfViewModifier)
}

func (p *PlayPlayerAbilities_758_4)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.FlyingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.FieldOfViewModifier, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x19
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x19
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x1a
type PlayPlayerAbilities_756_7 struct {
	/* Bit mask. 0x02: is flying. */
	Flags Byte // Byte
}

var _ Packet = (*PlayPlayerAbilities_756_7)(nil)

func (p PlayPlayerAbilities_756_7)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
}

func (p *PlayPlayerAbilities_756_7)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x19
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x12
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x13
type PlayPlayerAbilities_578_8 struct {
	/* Bit mask. 0x08: damage disabled (god mode), 0x04: can fly, 0x02: is flying, 0x01: is Creative */
	Flags Byte // Byte
	FlyingSpeed Float // Float
	WalkingSpeed Float // Float
}

var _ Packet = (*PlayPlayerAbilities_578_8)(nil)

func (p PlayPlayerAbilities_578_8)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.Float(p.FlyingSpeed)
	b.Float(p.WalkingSpeed)
}

func (p *PlayPlayerAbilities_578_8)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.FlyingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.WalkingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x31
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x2c
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x2c
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x2b
type PlayPlayerAbilities_498_9 struct {
	/* Bit field, see below */
	Flags Byte // Byte
	FlyingSpeed Float // Float
	/* Modifies the field of view, like a speed potion. A Notchian server will use the same value as the movement speed (send in the Entity Properties packet). */
	FieldOfViewModifier Float // Float
}

var _ Packet = (*PlayPlayerAbilities_498_9)(nil)

func (p PlayPlayerAbilities_498_9)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.Float(p.FlyingSpeed)
	b.Float(p.FieldOfViewModifier)
}

func (p *PlayPlayerAbilities_498_9)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.FlyingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.FieldOfViewModifier, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x19
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x13
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x13
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x12
type PlayPlayerAbilities_498_10 struct {
	/* Bit mask. 0x08: damage disabled (god mode), 0x04: can fly, 0x02: is flying, 0x01: is Creative */
	Flags Byte // Byte
	FlyingSpeed Float // Float
	WalkingSpeed Float // Float
}

var _ Packet = (*PlayPlayerAbilities_498_10)(nil)

func (p PlayPlayerAbilities_498_10)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.Float(p.FlyingSpeed)
	b.Float(p.WalkingSpeed)
}

func (p *PlayPlayerAbilities_498_10)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.FlyingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.WalkingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x2e
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x2e
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x2b
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x2b
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x2b
type PlayPlayerAbilities_404_11 struct {
	/* Bit field, see below */
	Flags Byte // Byte
	FlyingSpeed Float // Float
	/* Modifies the field of view, like a speed potion. A Notchian server will use the same value as the movement speed (send in the Entity Properties packet). */
	FieldOfViewModifier Float // Float
}

var _ Packet = (*PlayPlayerAbilities_404_11)(nil)

func (p PlayPlayerAbilities_404_11)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.Float(p.FlyingSpeed)
	b.Float(p.FieldOfViewModifier)
}

func (p *PlayPlayerAbilities_404_11)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.FlyingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.FieldOfViewModifier, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x17
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x17
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x13
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x12
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x12
type PlayPlayerAbilities_404_12 struct {
	/* Bit mask. 0x08: damage disabled (god mode), 0x04: can fly, 0x02: is flying, 0x01: is Creative */
	Flags Byte // Byte
	FlyingSpeed Float // Float
	WalkingSpeed Float // Float
}

var _ Packet = (*PlayPlayerAbilities_404_12)(nil)

func (p PlayPlayerAbilities_404_12)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.Float(p.FlyingSpeed)
	b.Float(p.WalkingSpeed)
}

func (p *PlayPlayerAbilities_404_12)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.FlyingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.WalkingSpeed, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x1d
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x1d
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x1c
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x1d
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x1c
type PlayPlayerAction_763_0 struct {
	/* The action the player is taking against the block (see below). */
	Status VarInt // VarInt Enum
	/* Block position. */
	Location Position // Position
	/* The face being hit (see below). */
	Face Byte // Byte Enum
	Sequence VarInt // VarInt
}

var _ Packet = (*PlayPlayerAction_763_0)(nil)

func (p PlayPlayerAction_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Status)
	b.Encode(p.Location)
	b.Byte(p.Face)
	b.VarInt(p.Sequence)
}

func (p *PlayPlayerAction_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Status, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Face, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Sequence, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x2e
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x2e
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x2e
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x2e
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x2e
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x2e
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x2c
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x2c
type PlayPlayerBlockPlacement_758_0 struct {
	/* The hand from which the block is placed; 0: main hand, 1: off hand */
	Hand VarInt // VarInt Enum
	/* Block position */
	Location Position // Position
	/* The face on which the block is placed (as documented at Player Digging) */
	Face VarInt // VarInt Enum
	/* The position of the crosshair on the block, from 0 to 1 increasing from west to east */
	CursorPositionX Float // Float
	/* The position of the crosshair on the block, from 0 to 1 increasing from bottom to top */
	CursorPositionY Float // Float
	/* The position of the crosshair on the block, from 0 to 1 increasing from north to south */
	CursorPositionZ Float // Float
	/* True when the player's head is inside of a block. */
	InsideBlock Bool // Boolean
}

var _ Packet = (*PlayPlayerBlockPlacement_758_0)(nil)

func (p PlayPlayerBlockPlacement_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
	b.Encode(p.Location)
	b.VarInt(p.Face)
	b.Float(p.CursorPositionX)
	b.Float(p.CursorPositionY)
	b.Float(p.CursorPositionZ)
	b.Bool(p.InsideBlock)
}

func (p *PlayPlayerBlockPlacement_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Face, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CursorPositionX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.CursorPositionY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.CursorPositionZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.InsideBlock, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x29
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x29
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x1f
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x1f
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x1f
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x1c
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x1c
type PlayPlayerBlockPlacement_404_1 struct {
	/* Block position */
	Location Position // Position
	/* The face on which the block is placed (as documented at Player Digging) */
	Face VarInt // VarInt Enum
	/* The hand from which the block is placed; 0: main hand, 1: off hand */
	Hand VarInt // VarInt Enum
	/* The position of the crosshair on the block, from 0 to 1 increasing from west to east */
	CursorPositionX Float // Float
	/* The position of the crosshair on the block, from 0 to 1 increasing from bottom to top */
	CursorPositionY Float // Float
	/* The position of the crosshair on the block, from 0 to 1 increasing from north to south */
	CursorPositionZ Float // Float
}

var _ Packet = (*PlayPlayerBlockPlacement_404_1)(nil)

func (p PlayPlayerBlockPlacement_404_1)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.Face)
	b.VarInt(p.Hand)
	b.Float(p.CursorPositionX)
	b.Float(p.CursorPositionY)
	b.Float(p.CursorPositionZ)
}

func (p *PlayPlayerBlockPlacement_404_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Face, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CursorPositionX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.CursorPositionY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.CursorPositionZ, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x1c
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x1c
type PlayPlayerBlockPlacement_210_2 struct {
	/* Block position */
	Location Position // Position
	/* The face on which the block is placed (as documented at Player Digging) */
	Face VarInt // VarInt Enum
	/* The hand from which the block is placed; 0: main hand, 1: off hand */
	Hand VarInt // VarInt Enum
	/* The position of the crosshair on the block, from 0 to 15 increasing from west to east */
	CursorPositionX UByte // Unsigned Byte
	/* The position of the crosshair on the block, from 0 to 15 increasing from bottom to top */
	CursorPositionY UByte // Unsigned Byte
	/* The position of the crosshair on the block, from 0 to 15 increasing from north to south */
	CursorPositionZ UByte // Unsigned Byte
}

var _ Packet = (*PlayPlayerBlockPlacement_210_2)(nil)

func (p PlayPlayerBlockPlacement_210_2)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.Face)
	b.VarInt(p.Hand)
	b.UByte(p.CursorPositionX)
	b.UByte(p.CursorPositionY)
	b.UByte(p.CursorPositionZ)
}

func (p *PlayPlayerBlockPlacement_210_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Face, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CursorPositionX, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.CursorPositionY, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.CursorPositionZ, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x1e
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x1e
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x1d
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x1e
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x1d
type PlayPlayerCommand_763_0 struct {
	/* Player ID */
	EntityID VarInt // VarInt
	/* The ID of the action, see below. */
	ActionID VarInt // VarInt Enum
	/* Only used by the “start jump with horse” action, in which case it ranges from 0 to 100. In all other cases it is 0. */
	JumpBoost VarInt // VarInt
}

var _ Packet = (*PlayPlayerCommand_763_0)(nil)

func (p PlayPlayerCommand_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.ActionID)
	b.VarInt(p.JumpBoost)
}

func (p *PlayPlayerCommand_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ActionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.JumpBoost, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x1a
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x1a
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x1a
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x1a
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x1b
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x1b
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x1a
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x1a
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x18
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x18
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x14
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x14
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x14
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x13
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x13
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x13
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x13
type PlayPlayerDigging_758_0 struct {
	/* The action the player is taking against the block (see below) */
	Status VarInt // VarInt Enum
	/* Block position */
	Location Position // Position
	/* The face being hit (see below) */
	Face Byte // Byte Enum
}

var _ Packet = (*PlayPlayerDigging_758_0)(nil)

func (p PlayPlayerDigging_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.Status)
	b.Encode(p.Location)
	b.Byte(p.Face)
}

func (p *PlayPlayerDigging_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Status, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Face, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x1f
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x1f
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x1e
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x1f
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x1e
type PlayPlayerInput_763_0 struct {
	/* Positive to the left of the player. */
	Sideways Float // Float
	/* Positive forward. */
	Forward Float // Float
	/* Bit mask. 0x1: jump, 0x2: unmount. */
	Flags UByte // Unsigned Byte
}

var _ Packet = (*PlayPlayerInput_763_0)(nil)

func (p PlayPlayerInput_763_0)Encode(b *PacketBuilder){
	b.Float(p.Sideways)
	b.Float(p.Forward)
	b.UByte(p.Flags)
}

func (p *PlayPlayerInput_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Sideways, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Forward, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x5f
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x5f
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x5e
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x5e
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x53
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x53
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x54
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x53
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x4e
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x4e
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x4a
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x4a
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x49
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x47
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x47
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x47
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x47
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x47
type PlayPlayerListHeaderAndFooter_758_0 struct {
	Header Object // Chat
	Footer Object // Chat
}

var _ Packet = (*PlayPlayerListHeaderAndFooter_758_0)(nil)

func (p PlayPlayerListHeaderAndFooter_758_0)Encode(b *PacketBuilder){
	b.JSON(p.Header)
	b.JSON(p.Footer)
}

func (p *PlayPlayerListHeaderAndFooter_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.Header); err != nil {
		return
	}
	if err = r.JSON(&p.Footer); err != nil {
		return
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x12
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x12
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0xf
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0xf
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x10
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0xe
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0xe
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0xe
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0xe
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x5
type PlayPlayerLook_404_0 struct {
	/* Absolute rotation on the X Axis, in degrees */
	Yaw Float // Float
	/* Absolute rotation on the Y Axis, in degrees */
	Pitch Float // Float
	/* True if the client is on the ground, False otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayerLook_404_0)(nil)

func (p PlayPlayerLook_404_0)Encode(b *PacketBuilder){
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayPlayerLook_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Yaw, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x14
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x14
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x14
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x14
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x15
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x15
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x14
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x14
type PlayPlayerMovement_758_0 struct {
	/* True if the client is on the ground, false otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayerMovement_758_0)(nil)

func (p PlayPlayerMovement_758_0)Encode(b *PacketBuilder){
	b.Bool(p.OnGround)
}

func (p *PlayPlayerMovement_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x11
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x11
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x11
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x11
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x12
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x12
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x11
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x11
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x10
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x10
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0xd
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0xd
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0xe
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0xc
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0xc
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0xc
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0xc
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x4
type PlayPlayerPosition_758_0 struct {
	/* Absolute position */
	X Double // Double
	/* Absolute position, normally Head Y - 1.62 */
	FeetY Double // Double
	/* Absolute position */
	Z Double // Double
	/* True if the client is on the ground, false otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayerPosition_758_0)(nil)

func (p PlayPlayerPosition_758_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.FeetY)
	b.Double(p.Z)
	b.Bool(p.OnGround)
}

func (p *PlayPlayerPosition_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.FeetY, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x38
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x38
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x38
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x38
type PlayPlayerPositionAndLook_758_0 struct {
	/* Absolute or relative position, depending on Flags. */
	X Double // Double
	/* Absolute or relative position, depending on Flags. */
	Y Double // Double
	/* Absolute or relative position, depending on Flags. */
	Z Double // Double
	/* Absolute or relative rotation on the X axis, in degrees. */
	Yaw Float // Float
	/* Absolute or relative rotation on the Y axis, in degrees. */
	Pitch Float // Float
	/* Bit field, see below. */
	Flags Byte // Byte
	/* Client should confirm this packet with Teleport Confirm containing the same Teleport ID. */
	TeleportID VarInt // VarInt
	/* True if the player should dismount their vehicle. */
	DismountVehicle Bool // Boolean
}

var _ Packet = (*PlayPlayerPositionAndLook_758_0)(nil)

func (p PlayPlayerPositionAndLook_758_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Byte(p.Flags)
	b.VarInt(p.TeleportID)
	b.Bool(p.DismountVehicle)
}

func (p *PlayPlayerPositionAndLook_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.TeleportID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DismountVehicle, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x34
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x34
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x36
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x35
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x32
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x32
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x2f
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x2f
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x2e
type PlayPlayerPositionAndLook_754_1 struct {
	/* Absolute or relative position, depending on Flags */
	X Double // Double
	/* Absolute or relative position, depending on Flags */
	Y Double // Double
	/* Absolute or relative position, depending on Flags */
	Z Double // Double
	/* Absolute or relative rotation on the X axis, in degrees */
	Yaw Float // Float
	/* Absolute or relative rotation on the Y axis, in degrees */
	Pitch Float // Float
	/* Bit field, see below */
	Flags Byte // Byte
	/* Client should confirm this packet with Teleport Confirm containing the same Teleport ID */
	TeleportID VarInt // VarInt
}

var _ Packet = (*PlayPlayerPositionAndLook_754_1)(nil)

func (p PlayPlayerPositionAndLook_754_1)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Byte(p.Flags)
	b.VarInt(p.TeleportID)
}

func (p *PlayPlayerPositionAndLook_754_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.TeleportID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x11
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x11
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0xe
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0xe
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0xd
type PlayPlayerPositionAndLook_404_2 struct {
	/* Absolute position */
	X Double // Double
	/* Absolute feet position, normally Head Y - 1.62 */
	FeetY Double // Double
	/* Absolute position */
	Z Double // Double
	/* Absolute rotation on the X Axis, in degrees */
	Yaw Float // Float
	/* Absolute rotation on the Y Axis, in degrees */
	Pitch Float // Float
	/* True if the client is on the ground, false otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayerPositionAndLook_404_2)(nil)

func (p PlayPlayerPositionAndLook_404_2)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.FeetY)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayPlayerPositionAndLook_404_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.FeetY, ok = r.Double(); !ok {
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
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x2e
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x2e
type PlayPlayerPositionAndLook_335_3 struct {
	/* Absolute or relative position, depending on Flags */
	X Double // Double
	/* Absolute or relative position, depending on Flags */
	Y Double // Double
	/* Absolute or relative position, depending on Flags */
	Z Double // Double
	/* Absolute or relative rotation on the X axis, in degrees */
	Yaw Float // Float
	/* Absolute or relative rotation on the Y axis, in degrees */
	Pitch Float // Float
	/* Bit field, see below */
	Flags Byte // Byte
	/* Client should confirm this packet with Teleport Confirm containing the same Teleport ID */
	TeleportID VarInt // VarInt
}

var _ Packet = (*PlayPlayerPositionAndLook_335_3)(nil)

func (p PlayPlayerPositionAndLook_335_3)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Byte(p.Flags)
	b.VarInt(p.TeleportID)
}

func (p *PlayPlayerPositionAndLook_335_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.TeleportID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0xf
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0xd
type PlayPlayerPositionAndLook_335_4 struct {
	/* Absolute position */
	X Double // Double
	/* Absolute feet position, normally Head Y - 1.62 */
	FeetY Double // Double
	/* Absolute position */
	Z Double // Double
	/* Absolute rotation on the X Axis, in degrees */
	Yaw Float // Float
	/* Absolute rotation on the Y Axis, in degrees */
	Pitch Float // Float
	/* True if the client is on the ground, false otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayerPositionAndLook_335_4)(nil)

func (p PlayPlayerPositionAndLook_335_4)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.FeetY)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayPlayerPositionAndLook_335_4)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.FeetY, ok = r.Double(); !ok {
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
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x2e
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x2e
type PlayPlayerPositionAndLook_210_5 struct {
	/* Absolute or relative position, depending on Flags */
	X Double // Double
	/* Absolute or relative position, depending on Flags */
	Y Double // Double
	/* Absolute or relative position, depending on Flags */
	Z Double // Double
	/* Absolute or relative rotation on the X axis, in degrees */
	Yaw Float // Float
	/* Absolute or relative rotation on the Y axis, in degrees */
	Pitch Float // Float
	/* Bit field, see below */
	Flags Byte // Byte
	/* Client should confirm this packet with Teleport Confirm containing the same Teleport ID */
	TeleportID VarInt // VarInt
}

var _ Packet = (*PlayPlayerPositionAndLook_210_5)(nil)

func (p PlayPlayerPositionAndLook_210_5)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Byte(p.Flags)
	b.VarInt(p.TeleportID)
}

func (p *PlayPlayerPositionAndLook_210_5)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.TeleportID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0xd
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0xd
type PlayPlayerPositionAndLook_210_6 struct {
	/* Absolute position */
	X Double // Double
	/* Absolute feet position, normally Head Y - 1.62 */
	FeetY Double // Double
	/* Absolute position */
	Z Double // Double
	/* Absolute rotation on the X Axis, in degrees */
	Yaw Float // Float
	/* Absolute rotation on the Y Axis, in degrees */
	Pitch Float // Float
	/* True if the client is on the ground, false otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayerPositionAndLook_210_6)(nil)

func (p PlayPlayerPositionAndLook_210_6)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.FeetY)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayPlayerPositionAndLook_210_6)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.FeetY, ok = r.Double(); !ok {
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
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x12
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x12
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x12
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x12
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x13
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x13
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x12
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x12
type PlayPlayerPositionAndRotation_758_0 struct {
	/* Absolute position */
	X Double // Double
	/* Absolute feet position, normally Head Y - 1.62 */
	FeetY Double // Double
	/* Absolute position */
	Z Double // Double
	/* Absolute rotation on the X Axis, in degrees */
	Yaw Float // Float
	/* Absolute rotation on the Y Axis, in degrees */
	Pitch Float // Float
	/* True if the client is on the ground, false otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayerPositionAndRotation_758_0)(nil)

func (p PlayPlayerPositionAndRotation_758_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.FeetY)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayPlayerPositionAndRotation_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.FeetY, ok = r.Double(); !ok {
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
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x13
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x13
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x13
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x13
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x14
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x14
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x13
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x13
type PlayPlayerRotation_758_0 struct {
	/* Absolute rotation on the X Axis, in degrees */
	Yaw Float // Float
	/* Absolute rotation on the Y Axis, in degrees */
	Pitch Float // Float
	/* True if the client is on the ground, False otherwise */
	OnGround Bool // Boolean
}

var _ Packet = (*PlayPlayerRotation_758_0)(nil)

func (p PlayPlayerRotation_758_0)Encode(b *PacketBuilder){
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayPlayerRotation_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Yaw, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x17
// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0xd
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x17
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0xd
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x15
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0xc
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x16
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0xd
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x15
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0xc
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x18
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0xa
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x18
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0xa
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x18
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0xa
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x18
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0xa
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x17
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0xb
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x17
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0xb
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x19
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0xb
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x18
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0xb
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x19
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0xa
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x19
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0xa
type PlayPluginMessage_763_0 struct {
	/* Name of the plugin channel used to send the data */
	Channel Identifier // Identifier
	/* Any data, depending on the channel. minecraft: channels are documented here.  The length of this array must be inferred from the packet length. */
	Data ByteArray // Byte Array
}

var _ Packet = (*PlayPluginMessage_763_0)(nil)

func (p PlayPluginMessage_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Channel)
	b.ReadAll(p.Data)
}

func (p *PlayPluginMessage_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Channel.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x18
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x9
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x18
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x9
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x18
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0xa
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x18
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x9
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x18
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x9
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x18
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x9
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x18
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x9
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x3f
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x17
type PlayPluginMessage_340_1 struct {
	/* Name of the plugin channel used to send the data */
	Channel String // String
	/* Any data, depending on the channel. MC| channels are documented here. */
	Data ByteArray // Byte Array
}

var _ Packet = (*PlayPluginMessage_340_1)(nil)

func (p PlayPluginMessage_340_1)Encode(b *PacketBuilder){
	b.String(p.Channel)
	b.ReadAll(p.Data)
}

func (p *PlayPluginMessage_340_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Channel, ok = r.String(); !ok {
		return io.EOF
	}
	if err = p.Data.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x20
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x20
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x1f
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x20
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x1f
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x1d
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x1d
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x1d
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x1d
type PlayPong_763_0 struct {
	/* id is the same as the ping packet */
	ID Int // Int
}

var _ Packet = (*PlayPong_763_0)(nil)

func (p PlayPong_763_0)Encode(b *PacketBuilder){
	b.Int(p.ID)
}

func (p *PlayPong_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ID, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x29
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x29
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x29
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x29
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x28
type PlayProgramCommandBlock_763_0 struct {
	Location Position // Position
	Command String // String (32767)
	/* One of SEQUENCE (0), AUTO (1), or REDSTONE (2). */
	Mode VarInt // VarInt Enum
	/* 0x01: Track Output (if false, the output of the previous command will not be stored within the command block); 0x02: Is conditional; 0x04: Automatic. */
	Flags Byte // Byte
}

var _ Packet = (*PlayProgramCommandBlock_763_0)(nil)

func (p PlayProgramCommandBlock_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.String(p.Command)
	b.VarInt(p.Mode)
	b.Byte(p.Flags)
}

func (p *PlayProgramCommandBlock_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Command, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x2a
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x2a
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x2a
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x2a
type PlayProgramCommandBlockMinecart_763_0 struct {
	EntityID VarInt // VarInt
	Command String // String (32767)
	/* If false, the output of the previous command will not be stored within the command block. */
	TrackOutput Bool // Boolean
}

var _ Packet = (*PlayProgramCommandBlockMinecart_763_0)(nil)

func (p PlayProgramCommandBlockMinecart_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.String(p.Command)
	b.Bool(p.TrackOutput)
}

func (p *PlayProgramCommandBlockMinecart_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Command, ok = r.String(); !ok {
		return io.EOF
	}
	if p.TrackOutput, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x2c
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x2c
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x2c
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x2c
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x2b
type PlayProgramJigsawBlock_763_0 struct {
	/* Block entity location */
	Location Position // Position
	Name Identifier // Identifier
	Target Identifier // Identifier
	Pool Identifier // Identifier
	/* "Turns into" on the GUI, final_state in NBT. */
	FinalState String // String (32767)
	/* rollable if the attached piece can be rotated, else aligned. */
	JointType String // String
}

var _ Packet = (*PlayProgramJigsawBlock_763_0)(nil)

func (p PlayProgramJigsawBlock_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.Encode(p.Name)
	b.Encode(p.Target)
	b.Encode(p.Pool)
	b.String(p.FinalState)
	b.String(p.JointType)
}

func (p *PlayProgramJigsawBlock_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Name.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Target.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pool.DecodeFrom(r); err != nil {
		return
	}
	if p.FinalState, ok = r.String(); !ok {
		return io.EOF
	}
	if p.JointType, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x2d
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x2d
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x2d
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x2d
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x2c
type PlayProgramStructureBlock_763_0 struct {
	/* Block entity location. */
	Location Position // Position
	/* An additional action to perform beyond simply saving the given data; see below. */
	Action VarInt // VarInt Enum
	/* One of SAVE (0), LOAD (1), CORNER (2), DATA (3). */
	Mode VarInt // VarInt Enum
	Name String // String (32767)
	/* Between -32 and 32. */
	OffsetX Byte // Byte
	/* Between -32 and 32. */
	OffsetY Byte // Byte
	/* Between -32 and 32. */
	OffsetZ Byte // Byte
	/* Between 0 and 32. */
	SizeX Byte // Byte
	/* Between 0 and 32. */
	SizeY Byte // Byte
	/* Between 0 and 32. */
	SizeZ Byte // Byte
	/* One of NONE (0), LEFT_RIGHT (1), FRONT_BACK (2). */
	Mirror VarInt // VarInt Enum
	/* One of NONE (0), CLOCKWISE_90 (1), CLOCKWISE_180 (2), COUNTERCLOCKWISE_90 (3). */
	Rotation VarInt // VarInt Enum
	Metadata String // String (128)
	/* Between 0 and 1. */
	Integrity Float // Float
	Seed VarLong // VarLong
	/* 0x01: Ignore entities; 0x02: Show air; 0x04: Show bounding box. */
	Flags Byte // Byte
}

var _ Packet = (*PlayProgramStructureBlock_763_0)(nil)

func (p PlayProgramStructureBlock_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.Action)
	b.VarInt(p.Mode)
	b.String(p.Name)
	b.Byte(p.OffsetX)
	b.Byte(p.OffsetY)
	b.Byte(p.OffsetZ)
	b.Byte(p.SizeX)
	b.Byte(p.SizeY)
	b.Byte(p.SizeZ)
	b.VarInt(p.Mirror)
	b.VarInt(p.Rotation)
	b.String(p.Metadata)
	b.Float(p.Integrity)
	b.VarLong(p.Seed)
	b.Byte(p.Flags)
}

func (p *PlayProgramStructureBlock_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Action, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Name, ok = r.String(); !ok {
		return io.EOF
	}
	if p.OffsetX, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.OffsetY, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.OffsetZ, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.SizeX, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.SizeY, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.SizeZ, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Mirror, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Rotation, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Metadata, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Integrity, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Seed, ok = r.VarLong(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x1
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x1
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x1
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x1
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x1
type PlayQueryBlockEntityTag_763_0 struct {
	/* An incremental ID so that the client can verify that the response matches. */
	TransactionID VarInt // VarInt
	/* The location of the block to check. */
	Location Position // Position
}

var _ Packet = (*PlayQueryBlockEntityTag_763_0)(nil)

func (p PlayQueryBlockEntityTag_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionID)
	b.Encode(p.Location)
}

func (p *PlayQueryBlockEntityTag_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x1
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x1
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x1
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x1
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x1
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x1
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x1
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x1
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x1
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x1
type PlayQueryBlockNBT_758_0 struct {
	/* An incremental ID so that the client can verify that the response matches. */
	TransactionID VarInt // VarInt
	/* The location of the block to check. */
	Location Position // Position
}

var _ Packet = (*PlayQueryBlockNBT_758_0)(nil)

func (p PlayQueryBlockNBT_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionID)
	b.Encode(p.Location)
}

func (p *PlayQueryBlockNBT_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0xc
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0xc
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0xc
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0xc
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0xd
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0xd
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0xd
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0xc
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0xc
type PlayQueryEntityNBT_758_0 struct {
	/* An incremental ID so that the client can verify that the response matches. */
	TransactionID VarInt // VarInt
	/* The ID of the entity to query. */
	EntityID VarInt // VarInt
}

var _ Packet = (*PlayQueryEntityNBT_758_0)(nil)

func (p PlayQueryEntityNBT_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionID)
	b.VarInt(p.EntityID)
}

func (p *PlayQueryEntityNBT_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0xf
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0xf
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0xe
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0xf
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0xe
type PlayQueryEntityTag_763_0 struct {
	/* An incremental ID so that the client can verify that the response matches. */
	TransactionID VarInt // VarInt
	/* The ID of the entity to query. */
	EntityID VarInt // VarInt
}

var _ Packet = (*PlayQueryEntityTag_763_0)(nil)

func (p PlayQueryEntityTag_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionID)
	b.VarInt(p.EntityID)
}

func (p *PlayQueryEntityTag_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x3e
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x3e
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x3a
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x3b
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x38
type PlayRemoveEntities_763_0 struct {
	/* Number of elements in the following array. */
	Count VarInt // VarInt
	/* The list of entities to destroy. */
	EntityIDs []VarInt // Array of VarInt
}

var _ Packet = (*PlayRemoveEntities_763_0)(nil)

func (p PlayRemoveEntities_763_0)Encode(b *PacketBuilder){
	p.Count = len(p.EntityIDs)
	b.VarInt(p.Count)
	for _, v := range p.EntityIDs {
		b.VarInt(v)
	}
}

func (p *PlayRemoveEntities_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.EntityIDs = make([]VarInt, p.Count)
	for i, _ := range p.EntityIDs {
		if p.EntityIDs[i], ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x3f
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x3f
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x3b
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x3c
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x39
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x3b
type PlayRemoveEntityEffect_763_0 struct {
	EntityID VarInt // VarInt
	/* See this table. */
	EffectID VarInt // VarInt
}

var _ Packet = (*PlayRemoveEntityEffect_763_0)(nil)

func (p PlayRemoveEntityEffect_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.VarInt(p.EffectID)
}

func (p *PlayRemoveEntityEffect_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x3b
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x3b
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x3b
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x37
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x37
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x39
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x38
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x36
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x36
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x33
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x33
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x32
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x31
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x31
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x31
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x31
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x1e
type PlayRemoveEntityEffect_757_1 struct {
	EntityID VarInt // VarInt
	/* See this table */
	EffectID Byte // Byte
}

var _ Packet = (*PlayRemoveEntityEffect_757_1)(nil)

func (p PlayRemoveEntityEffect_757_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Byte(p.EffectID)
}

func (p *PlayRemoveEntityEffect_757_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EffectID, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x23
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x23
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x23
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x23
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x22
type PlayRenameItem_763_0 struct {
	/* The new name of the item. */
	ItemName String // String (32767)
}

var _ Packet = (*PlayRenameItem_763_0)(nil)

func (p PlayRenameItem_763_0)Encode(b *PacketBuilder){
	b.String(p.ItemName)
}

func (p *PlayRenameItem_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ItemName, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x40
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x40
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x3c
type PlayResourcePack_763_0 struct {
	/* The URL to the resource pack. */
	URL String // String (32767)
	/* A 40 character hexadecimal and lowercase SHA-1 hash of the resource pack file.If it's not a 40 character hexadecimal string, the client will not use it for hash verification and likely waste bandwidth — but it will still treat it as a unique id */
	Hash String // String (40)
	/* The notchian client will be forced to use the resource pack from the server. If they decline they will be kicked from the server. */
	Forced Bool // Boolean
	/* true If the next field will be sent false otherwise. When false, this is the end of the packet */
	HasPromptMessage Bool // Boolean
	/* This is shown in the prompt making the client accept or decline the resource pack. */
	PromptMessage Object // Optional Chat
}

var _ Packet = (*PlayResourcePack_763_0)(nil)

func (p PlayResourcePack_763_0)Encode(b *PacketBuilder){
	b.String(p.URL)
	b.String(p.Hash)
	b.Bool(p.Forced)
	b.Bool(p.HasPromptMessage)
	b.JSON(p.PromptMessage)
}

func (p *PlayResourcePack_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.URL, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Hash, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Forced, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.HasPromptMessage, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.PromptMessage); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x24
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x24
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x24
type PlayResourcePack_763_1 struct {
	/* 0: successfully loaded, 1: declined, 2: failed download, 3: accepted. */
	Result VarInt // VarInt Enum
}

var _ Packet = (*PlayResourcePack_763_1)(nil)

func (p PlayResourcePack_763_1)Encode(b *PacketBuilder){
	b.VarInt(p.Result)
}

func (p *PlayResourcePack_763_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Result, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x3d
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x3a
type PlayResourcePack_760_2 struct {
	/* The URL to the resource pack. */
	URL String // String (32767)
	/* A 40 character hexadecimal and lowercase SHA-1 hash of the resource pack file. (must be lower case in order to work)If it's not a 40 character hexadecimal string, the client will not use it for hash verification and likely waste bandwidth — but it will still treat it as a unique id */
	Hash String // String (40)
	/* The notchian client will be forced to use the resource pack from the server. If they decline they will be kicked from the server. */
	Forced Bool // Boolean
	/* true If the next field will be sent false otherwise. When false, this is the end of the packet */
	HasPromptMessage Bool // Boolean
	/* This is shown in the prompt making the client accept or decline the resource pack. */
	PromptMessage Object // Optional Chat
}

var _ Packet = (*PlayResourcePack_760_2)(nil)

func (p PlayResourcePack_760_2)Encode(b *PacketBuilder){
	b.String(p.URL)
	b.String(p.Hash)
	b.Bool(p.Forced)
	b.Bool(p.HasPromptMessage)
	b.JSON(p.PromptMessage)
}

func (p *PlayResourcePack_760_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.URL, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Hash, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Forced, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.HasPromptMessage, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.PromptMessage); err != nil {
		return
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x24
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x23
type PlayResourcePack_760_3 struct {
	/* 0: successfully loaded, 1: declined, 2: failed download, 3: accepted. */
	Result VarInt // VarInt Enum
}

var _ Packet = (*PlayResourcePack_760_3)(nil)

func (p PlayResourcePack_760_3)Encode(b *PacketBuilder){
	b.VarInt(p.Result)
}

func (p *PlayResourcePack_760_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Result, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x3c
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x3c
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x3c
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x3c
type PlayResourcePackSend_758_0 struct {
	/* The URL to the resource pack. */
	URL String // String (32767)
	/* A 40 character hexadecimal and lowercase SHA-1 hash of the resource pack file. (must be lower case in order to work)If it's not a 40 character hexadecimal string, the client will not use it for hash verification and likely waste bandwidth — but it will still treat it as a unique id */
	Hash String // String (40)
	/* The notchian client will be forced to use the resource pack from the server. If they decline they will be kicked from the server. */
	Forced Bool // Boolean
	/* true If the next field will be sent false otherwise. When false, this is the end of the packet */
	HasPromptMessage Bool // Boolean
	/* This is shown in the prompt making the client accept or decline the resource pack. */
	PromptMessage Object // Optional Chat
}

var _ Packet = (*PlayResourcePackSend_758_0)(nil)

func (p PlayResourcePackSend_758_0)Encode(b *PacketBuilder){
	b.String(p.URL)
	b.String(p.Hash)
	b.Bool(p.Forced)
	b.Bool(p.HasPromptMessage)
	b.JSON(p.PromptMessage)
}

func (p *PlayResourcePackSend_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.URL, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Hash, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Forced, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.HasPromptMessage, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.PromptMessage); err != nil {
		return
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x38
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x38
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x3a
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x39
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x37
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x37
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x34
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x34
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x33
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x32
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x32
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x32
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x32
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x48
type PlayResourcePackSend_754_1 struct {
	/* The URL to the resource pack. */
	URL String // String
	/* A 40 character hexadecimal and lowercase SHA-1 hash of the resource pack file. (must be lower case in order to work)If it's not a 40 character hexadecimal string, the client will not use it for hash verification and likely waste bandwidth — but it will still treat it as a unique id */
	Hash String // String
}

var _ Packet = (*PlayResourcePackSend_754_1)(nil)

func (p PlayResourcePackSend_754_1)Encode(b *PacketBuilder){
	b.String(p.URL)
	b.String(p.Hash)
}

func (p *PlayResourcePackSend_754_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.URL, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Hash, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x21
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x21
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x21
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x21
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x21
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x21
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x1f
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x1f
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x1d
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x1d
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x18
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x18
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x18
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x16
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x16
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x16
type PlayResourcePackStatus_758_0 struct {
	/* 0: successfully loaded, 1: declined, 2: failed download, 3: accepted */
	Result VarInt // VarInt Enum
}

var _ Packet = (*PlayResourcePackStatus_758_0)(nil)

func (p PlayResourcePackStatus_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.Result)
}

func (p *PlayResourcePackStatus_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Result, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x16
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x19
type PlayResourcePackStatus_110_1 struct {
	/* The hash sent in the Resource Pack Send packet. */
	Hash String // String
	/* 0: successfully loaded, 1: declined, 2: failed download, 3: accepted */
	Result VarInt // VarInt Enum
}

var _ Packet = (*PlayResourcePackStatus_110_1)(nil)

func (p PlayResourcePackStatus_110_1)Encode(b *PacketBuilder){
	b.String(p.Hash)
	b.VarInt(p.Result)
}

func (p *PlayResourcePackStatus_110_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hash, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Result, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x3d
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x3e
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x3b
type PlayRespawn_761_2 struct {
	/* Valid dimensions are defined per dimension registry sent in Login (play) */
	DimensionType Identifier // Identifier
	/* Name of the dimension being spawned into. */
	DimensionName Identifier // Identifier
	/* First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise */
	HashedSeed Long // Long
	/* 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. */
	Gamemode UByte // Unsigned Byte
	/* -1: Undefined (null), 0: Survival, 1: Creative, 2: Adventure, 3: Spectator. The previous gamemode. Vanilla client uses this for the debug (F3 + N & F3 + F4) gamemode switch. (More information needed) */
	PreviousGamemode Byte // Byte
	/* True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks. */
	IsDebug Bool // Boolean
	/* True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63. */
	IsFlat Bool // Boolean
	/* If false, metadata is reset on the respawned player entity.  Set to true for dimension changes (including the dimension change triggered by sending client status perform respawn to exit the end poem/credits), and false for normal respawns. */
	CopyMetadata Bool // Boolean
	/* If true, then the next two fields are present. */
	HasDeathLocation Bool // Boolean
	/* Name of the dimension the player died in. */
	DeathDimensionName Identifier // Optional Identifier
	/* The location that the player died at. */
	DeathLocation Position // Optional Position
}

var _ Packet = (*PlayRespawn_761_2)(nil)

func (p PlayRespawn_761_2)Encode(b *PacketBuilder){
	b.Encode(p.DimensionType)
	b.Encode(p.DimensionName)
	b.Long(p.HashedSeed)
	b.UByte(p.Gamemode)
	b.Byte(p.PreviousGamemode)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
	b.Bool(p.CopyMetadata)
	b.Bool(p.HasDeathLocation)
	b.Encode(p.DeathDimensionName)
	b.Encode(p.DeathLocation)
}

func (p *PlayRespawn_761_2)DecodeFrom(r *PacketReader)(err error){
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
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.PreviousGamemode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.IsDebug, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.IsFlat, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.CopyMetadata, ok = r.Bool(); !ok {
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

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x3d
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x3d
type PlayRespawn_758_3 struct {
	/* Valid dimensions are defined per dimension registry sent in Join Game */
	Dimension *NBTCompound // NBT Tag Compound
	/* Name of the dimension being spawned into. */
	DimensionName Identifier // Identifier
	/* First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise */
	HashedSeed Long // Long
	/* 0: survival, 1: creative, 2: adventure, 3: spectator. The hardcore flag is not included */
	Gamemode UByte // Unsigned Byte
	/* -1: null 0: survival, 1: creative, 2: adventure, 3: spectator. The hardcore flag is not included. The previous gamemode. (More information needed) */
	PreviousGamemode UByte // Unsigned Byte
	/* True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks. */
	IsDebug Bool // Boolean
	/* True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63. */
	IsFlat Bool // Boolean
	/* If false, metadata is reset on the respawned player entity.  Set to true for dimension changes (including the dimension change triggered by sending client status perform respawn to exit the end poem/credits), and false for normal respawns. */
	CopyMetadata Bool // Boolean
}

var _ Packet = (*PlayRespawn_758_3)(nil)

func (p PlayRespawn_758_3)Encode(b *PacketBuilder){
	b.Encode(p.Dimension)
	b.Encode(p.DimensionName)
	b.Long(p.HashedSeed)
	b.UByte(p.Gamemode)
	b.UByte(p.PreviousGamemode)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
	b.Bool(p.CopyMetadata)
}

func (p *PlayRespawn_758_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Dimension.DecodeFrom(r); err != nil {
		return
	}
	if err = p.DimensionName.DecodeFrom(r); err != nil {
		return
	}
	if p.HashedSeed, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.PreviousGamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.IsDebug, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.IsFlat, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.CopyMetadata, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x3d
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x3d
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x39
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x39
type PlayRespawn_756_4 struct {
	/* Valid dimensions are defined per dimension registry sent in Join Game */
	Dimension *NBTCompound // NBT Tag Compound
	/* Name of the world being spawned into */
	WorldName Identifier // Identifier
	/* First 8 bytes of the SHA-256 hash of the world's seed. */
	HashedSeed Long // Long
	/* 0: survival, 1: creative, 2: adventure, 3: spectator. The hardcore flag is not included */
	Gamemode UByte // Unsigned Byte
	/* 0: survival, 1: creative, 2: adventure, 3: spectator. The hardcore flag is not included. The previous gamemode. (More information needed) */
	PreviousGamemode UByte // Unsigned Byte
	/* True if the world is a debug mode world; debug mode worlds cannot be modified and have predefined blocks */
	IsDebug Bool // Boolean
	/* True if the world is a superflat world; flat worlds have different void fog and a horizon at y=0 instead of y=63 */
	IsFlat Bool // Boolean
	/* If false, metadata is reset on the respawned player entity.  Set to true for dimension changes (including the dimension change triggered by sending client status perform respawn to exit the end poem/credits), and false for normal respawns. */
	CopyMetadata Bool // Boolean
}

var _ Packet = (*PlayRespawn_756_4)(nil)

func (p PlayRespawn_756_4)Encode(b *PacketBuilder){
	b.Encode(p.Dimension)
	b.Encode(p.WorldName)
	b.Long(p.HashedSeed)
	b.UByte(p.Gamemode)
	b.UByte(p.PreviousGamemode)
	b.Bool(p.IsDebug)
	b.Bool(p.IsFlat)
	b.Bool(p.CopyMetadata)
}

func (p *PlayRespawn_756_4)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Dimension.DecodeFrom(r); err != nil {
		return
	}
	if err = p.WorldName.DecodeFrom(r); err != nil {
		return
	}
	if p.HashedSeed, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.PreviousGamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.IsDebug, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.IsFlat, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.CopyMetadata, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x38
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x38
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x35
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x35
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x34
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x33
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x33
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x33
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x33
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x7
type PlayRespawn_404_7 struct {
	/* -1: The Nether, 0: The Overworld, 1: The End */
	Dimension Int // Int
	/* 0: Peaceful, 1: Easy, 2: Normal, 3: Hard */
	Difficulty UByte // Unsigned Byte
	/* 0: survival, 1: creative, 2: adventure. The hardcore flag is not included */
	Gamemode UByte // Unsigned Byte
	/* Same as Join Game */
	LevelType String // String
}

var _ Packet = (*PlayRespawn_404_7)(nil)

func (p PlayRespawn_404_7)Encode(b *PacketBuilder){
	b.Int(p.Dimension)
	b.UByte(p.Difficulty)
	b.UByte(p.Gamemode)
	b.String(p.LevelType)
}

func (p *PlayRespawn_404_7)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Dimension, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Difficulty, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Gamemode, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.LevelType, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x53
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x53
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x53
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x53
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x4a
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x4a
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x4a
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x49
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x45
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x45
type PlayScoreboardObjective_758_0 struct {
	/* An unique name for the objective */
	ObjectiveName String // String (16)
	/* 0 to create the scoreboard. 1 to remove the scoreboard. 2 to update the display text. */
	Mode Byte // Byte
	/* Only if mode is 0 or 2. The text to be displayed for the score */
	ObjectiveValue Object // Optional Chat
	/* Only if mode is 0 or 2. 0 = "integer", 1 = "hearts". */
	Type VarInt // Optional VarInt enum
}

var _ Packet = (*PlayScoreboardObjective_758_0)(nil)

func (p PlayScoreboardObjective_758_0)Encode(b *PacketBuilder){
	b.String(p.ObjectiveName)
	b.Byte(p.Mode)
	b.JSON(p.ObjectiveValue)
	b.VarInt(p.Type)
}

func (p *PlayScoreboardObjective_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ObjectiveName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.ObjectiveValue); err != nil {
		return
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x42
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x42
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x41
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x3f
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x3f
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x3f
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x3f
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x3b
type PlayScoreboardObjective_340_1 struct {
	/* An unique name for the objective */
	ObjectiveName String // String
	/* 0 to create the scoreboard. 1 to remove the scoreboard. 2 to update the display text. */
	Mode Byte // Byte
	/* Only if mode is 0 or 2. The text to be displayed for the score */
	ObjectiveValue String // Optional String
	/* Only if mode is 0 or 2. “integer” or “hearts” */
	Type String // Optional String
}

var _ Packet = (*PlayScoreboardObjective_340_1)(nil)

func (p PlayScoreboardObjective_340_1)Encode(b *PacketBuilder){
	b.String(p.ObjectiveName)
	b.Byte(p.Mode)
	b.String(p.ObjectiveValue)
	b.String(p.Type)
}

func (p *PlayScoreboardObjective_340_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ObjectiveName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ObjectiveValue, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x5
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x5
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x5
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x5
type PlaySculkVibrationSignal_758_0 struct {
	/* Source position for the vibration. */
	SourcePosition Position // Position
	/* Identifier of the destination codec type. */
	DestinationIdentifier Identifier // Identifier
	/* Vanilla default destinations are a block position encoded as a Position for "block" or an entity id encoded as a VarInt for "entity". */
	Destination any // Varies
	/* Ticks for the signal to arrive at the destination. */
	ArrivalTicks VarInt // VarInt
}

var _ Packet = (*PlaySculkVibrationSignal_758_0)(nil)

func (p PlaySculkVibrationSignal_758_0)Encode(b *PacketBuilder){
	b.Encode(p.SourcePosition)
	b.Encode(p.DestinationIdentifier)
	b.Encode(p.Destination)
	b.VarInt(p.ArrivalTicks)
}

func (p *PlaySculkVibrationSignal_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.SourcePosition.DecodeFrom(r); err != nil {
		return
	}
	if err = p.DestinationIdentifier.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Destination.DecodeFrom(r); err != nil {
		return
	}
	if p.ArrivalTicks, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x25
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x25
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x25
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x25
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x24
type PlaySeenAdvancements_763_0 struct {
	/* 0: Opened tab, 1: Closed screen. */
	Action VarInt // VarInt Enum
	/* Only present if action is Opened tab. */
	TabID Identifier // Optional identifier
}

var _ Packet = (*PlaySeenAdvancements_763_0)(nil)

func (p PlaySeenAdvancements_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Action)
	b.Encode(p.TabID)
}

func (p *PlaySeenAdvancements_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Action, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.TabID.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x40
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x40
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x40
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x40
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x3c
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x3c
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x3d
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x3c
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x3a
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x3a
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x37
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x37
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x36
type PlaySelectAdvancementTab_758_0 struct {
	/* Indicates if the next field is present */
	HasId Bool // Boolean
	/* See below */
	OptionalIdentifier String // String (32767)
}

var _ Packet = (*PlaySelectAdvancementTab_758_0)(nil)

func (p PlaySelectAdvancementTab_758_0)Encode(b *PacketBuilder){
	b.Bool(p.HasId)
	b.String(p.OptionalIdentifier)
}

func (p *PlaySelectAdvancementTab_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.HasId, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.OptionalIdentifier, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x44
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x44
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x40
type PlaySelectAdvancementsTab_763_0 struct {
	/* Indicates if the next field is present. */
	HasId Bool // Boolean
	/* See below. */
	OptionalIdentifier Identifier // Identifier
}

var _ Packet = (*PlaySelectAdvancementsTab_763_0)(nil)

func (p PlaySelectAdvancementsTab_763_0)Encode(b *PacketBuilder){
	b.Bool(p.HasId)
	b.Encode(p.OptionalIdentifier)
}

func (p *PlaySelectAdvancementsTab_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.HasId, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = p.OptionalIdentifier.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x41
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x3e
type PlaySelectAdvancementsTab_760_1 struct {
	/* Indicates if the next field is present. */
	HasId Bool // Boolean
	/* See below. */
	OptionalIdentifier String // String (32767)
}

var _ Packet = (*PlaySelectAdvancementsTab_760_1)(nil)

func (p PlaySelectAdvancementsTab_760_1)Encode(b *PacketBuilder){
	b.Bool(p.HasId)
	b.String(p.OptionalIdentifier)
}

func (p *PlaySelectAdvancementsTab_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.HasId, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.OptionalIdentifier, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x26
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x26
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x26
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x26
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x25
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x23
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x23
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x23
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x23
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x23
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x23
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x21
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x21
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x1f
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x1f
type PlaySelectTrade_763_0 struct {
	/* The selected slot int the players current (trading) inventory. (Was a full Integer for the plugin message) */
	SelectedSlot VarInt // VarInt
}

var _ Packet = (*PlaySelectTrade_763_0)(nil)

func (p PlaySelectTrade_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.SelectedSlot)
}

func (p *PlaySelectTrade_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SelectedSlot, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x45
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x45
type PlayServerData_763_0 struct {
	MOTD Object // Chat
	HasIcon Bool // Boolean
	/* Icon PNG base64 String */
	Icon String // Optional String (32767)
	EnforcesSecureChat Bool // Boolean
}

var _ Packet = (*PlayServerData_763_0)(nil)

func (p PlayServerData_763_0)Encode(b *PacketBuilder){
	b.JSON(p.MOTD)
	b.Bool(p.HasIcon)
	b.String(p.Icon)
	b.Bool(p.EnforcesSecureChat)
}

func (p *PlayServerData_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.MOTD); err != nil {
		return
	}
	if p.HasIcon, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Icon, ok = r.String(); !ok {
		return io.EOF
	}
	if p.EnforcesSecureChat, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x42
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x3f
type PlayServerData_760_2 struct {
	HasMOTD Bool // Boolean
	MOTD Object // Optional Chat
	HasIcon Bool // Boolean
	/* Icon PNG base64 String */
	Icon String // Optional String (32767)
	PreviewsChat Bool // Boolean
	EnforcesSecureChat Bool // Boolean
}

var _ Packet = (*PlayServerData_760_2)(nil)

func (p PlayServerData_760_2)Encode(b *PacketBuilder){
	b.Bool(p.HasMOTD)
	b.JSON(p.MOTD)
	b.Bool(p.HasIcon)
	b.String(p.Icon)
	b.Bool(p.PreviewsChat)
	b.Bool(p.EnforcesSecureChat)
}

func (p *PlayServerData_760_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.HasMOTD, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.MOTD); err != nil {
		return
	}
	if p.HasIcon, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Icon, ok = r.String(); !ok {
		return io.EOF
	}
	if p.PreviewsChat, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.EnforcesSecureChat, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0xe
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0xe
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0xe
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0xe
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0xd
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0xd
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0xe
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0xd
type PlayServerDifficulty_758_0 struct {
	/* 0: peaceful, 1: easy, 2: normal, 3: hard */
	Difficulty UByte // Unsigned Byte
	DifficultyLocked Bool // Boolean
}

var _ Packet = (*PlayServerDifficulty_758_0)(nil)

func (p PlayServerDifficulty_758_0)Encode(b *PacketBuilder){
	b.UByte(p.Difficulty)
	b.Bool(p.DifficultyLocked)
}

func (p *PlayServerDifficulty_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Difficulty, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.DifficultyLocked, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0xd
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0xd
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0xd
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0xd
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0xd
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0xd
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0xd
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0xd
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0xd
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x41
type PlayServerDifficulty_404_1 struct {
	/* 0: peaceful, 1: easy, 2: normal, 3: hard */
	Difficulty UByte // Unsigned Byte
}

var _ Packet = (*PlayServerDifficulty_404_1)(nil)

func (p PlayServerDifficulty_404_1)Encode(b *PacketBuilder){
	b.UByte(p.Difficulty)
}

func (p *PlayServerDifficulty_404_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Difficulty, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x46
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x46
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x42
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x43
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x40
type PlaySetActionBarText_763_0 struct {
	/* Displays a message above the hotbar (the same as position 2 in Player Chat Message. */
	ActionBarText Object // Chat
}

var _ Packet = (*PlaySetActionBarText_763_0)(nil)

func (p PlaySetActionBarText_763_0)Encode(b *PacketBuilder){
	b.JSON(p.ActionBarText)
}

func (p *PlaySetActionBarText_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.ActionBarText); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x27
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x27
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x27
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x27
type PlaySetBeaconEffect_763_0 struct {
	HasPrimaryEffect Bool // Boolean
	/* A Potion ID. (Was a full Integer for the plugin message). */
	PrimaryEffect VarInt // VarInt
	HasSecondaryEffect Bool // Boolean
	/* A Potion ID. (Was a full Integer for the plugin message). */
	SecondaryEffect VarInt // VarInt
}

var _ Packet = (*PlaySetBeaconEffect_763_0)(nil)

func (p PlaySetBeaconEffect_763_0)Encode(b *PacketBuilder){
	b.Bool(p.HasPrimaryEffect)
	b.VarInt(p.PrimaryEffect)
	b.Bool(p.HasSecondaryEffect)
	b.VarInt(p.SecondaryEffect)
}

func (p *PlaySetBeaconEffect_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.HasPrimaryEffect, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.PrimaryEffect, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.HasSecondaryEffect, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.SecondaryEffect, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x24
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x24
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x24
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x24
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x24
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x24
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x22
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x22
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x20
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x20
type PlaySetBeaconEffect_758_2 struct {
	/* A Potion ID. (Was a full Integer for the plugin message) */
	PrimaryEffect VarInt // VarInt
	/* A Potion ID. (Was a full Integer for the plugin message) */
	SecondaryEffect VarInt // VarInt
}

var _ Packet = (*PlaySetBeaconEffect_758_2)(nil)

func (p PlaySetBeaconEffect_758_2)Encode(b *PacketBuilder){
	b.VarInt(p.PrimaryEffect)
	b.VarInt(p.SecondaryEffect)
}

func (p *PlaySetBeaconEffect_758_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.PrimaryEffect, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.SecondaryEffect, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x7
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x7
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x6
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x6
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x6
type PlaySetBlockDestroyStage_763_0 struct {
	/* The ID of the entity breaking the block. */
	EntityID VarInt // VarInt
	/* Block Position. */
	Location Position // Position
	/* 0–9 to set it, any other value to remove it. */
	DestroyStage Byte // Byte
}

var _ Packet = (*PlaySetBlockDestroyStage_763_0)(nil)

func (p PlaySetBlockDestroyStage_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.Location)
	b.Byte(p.DestroyStage)
}

func (p *PlaySetBlockDestroyStage_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.DestroyStage, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x47
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x47
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x43
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x44
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x41
type PlaySetBorderCenter_763_0 struct {
	X Double // Double
	Z Double // Double
}

var _ Packet = (*PlaySetBorderCenter_763_0)(nil)

func (p PlaySetBorderCenter_763_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Z)
}

func (p *PlaySetBorderCenter_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x48
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x48
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x44
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x45
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x42
type PlaySetBorderLerpSize_763_0 struct {
	/* Current length of a single side of the world border, in meters. */
	OldDiameter Double // Double
	/* Target length of a single side of the world border, in meters. */
	NewDiameter Double // Double
	/* Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. */
	Speed VarLong // VarLong
}

var _ Packet = (*PlaySetBorderLerpSize_763_0)(nil)

func (p PlaySetBorderLerpSize_763_0)Encode(b *PacketBuilder){
	b.Double(p.OldDiameter)
	b.Double(p.NewDiameter)
	b.VarLong(p.Speed)
}

func (p *PlaySetBorderLerpSize_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.OldDiameter, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.NewDiameter, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Speed, ok = r.VarLong(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x49
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x49
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x45
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x46
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x43
type PlaySetBorderSize_763_0 struct {
	/* Length of a single side of the world border, in meters. */
	Diameter Double // Double
}

var _ Packet = (*PlaySetBorderSize_763_0)(nil)

func (p PlaySetBorderSize_763_0)Encode(b *PacketBuilder){
	b.Double(p.Diameter)
}

func (p *PlaySetBorderSize_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Diameter, ok = r.Double(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x4a
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x4a
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x46
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x47
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x44
type PlaySetBorderWarningDelay_763_0 struct {
	/* In seconds as set by /worldborder warning time. */
	WarningTime VarInt // VarInt
}

var _ Packet = (*PlaySetBorderWarningDelay_763_0)(nil)

func (p PlaySetBorderWarningDelay_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.WarningTime)
}

func (p *PlaySetBorderWarningDelay_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WarningTime, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x4b
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x4b
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x47
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x48
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x45
type PlaySetBorderWarningDistance_763_0 struct {
	/* In meters. */
	WarningBlocks VarInt // VarInt
}

var _ Packet = (*PlaySetBorderWarningDistance_763_0)(nil)

func (p PlaySetBorderWarningDistance_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.WarningBlocks)
}

func (p *PlaySetBorderWarningDistance_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WarningBlocks, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x4c
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x4c
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x48
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x49
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x46
type PlaySetCamera_763_0 struct {
	/* ID of the entity to set the client's camera to. */
	CameraID VarInt // VarInt
}

var _ Packet = (*PlaySetCamera_763_0)(nil)

func (p PlaySetCamera_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.CameraID)
}

func (p *PlaySetCamera_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.CameraID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x4e
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x4e
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x4a
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x4b
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x48
type PlaySetCenterChunk_763_0 struct {
	/* Chunk X coordinate of the player's position. */
	ChunkX VarInt // VarInt
	/* Chunk Z coordinate of the player's position. */
	ChunkZ VarInt // VarInt
}

var _ Packet = (*PlaySetCenterChunk_763_0)(nil)

func (p PlaySetCenterChunk_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.ChunkX)
	b.VarInt(p.ChunkZ)
}

func (p *PlaySetCenterChunk_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkX, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x12
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x12
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x10
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x11
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x11
type PlaySetContainerContent_763_0 struct {
	/* The ID of window which items are being sent for. 0 for player inventory. */
	WindowID UByte // Unsigned Byte
	/* The last received State ID from either a Set Container Slot or a Set Container Content packet */
	StateID VarInt // VarInt
	/* Number of elements in the following array. */
	Count VarInt // VarInt
	SlotData []*Slot // Array of Slot
	/* Item held by player. */
	CarriedItem *Slot // Slot
}

var _ Packet = (*PlaySetContainerContent_763_0)(nil)

func (p PlaySetContainerContent_763_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.VarInt(p.StateID)
	p.Count = len(p.SlotData)
	b.VarInt(p.Count)
	for _, v := range p.SlotData {
		b.Encode(v)
	}
	b.Encode(p.CarriedItem)
}

func (p *PlaySetContainerContent_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.StateID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.SlotData = make([]*Slot, p.Count)
	for i, _ := range p.SlotData {
		if err = p.SlotData[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if err = p.CarriedItem.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x13
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x13
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x11
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x12
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x12
type PlaySetContainerProperty_763_0 struct {
	WindowID UByte // Unsigned Byte
	/* The property to be updated, see below. */
	Property Short // Short
	/* The new value for the property, see below. */
	Value Short // Short
}

var _ Packet = (*PlaySetContainerProperty_763_0)(nil)

func (p PlaySetContainerProperty_763_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.Short(p.Property)
	b.Short(p.Value)
}

func (p *PlaySetContainerProperty_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Property, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.Value, ok = r.Short(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x14
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x14
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x12
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x13
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x13
type PlaySetContainerSlot_763_0 struct {
	/* The window which is being updated. 0 for player inventory. Note that all known window types include the player inventory. This packet will only be sent for the currently opened window while the player is performing actions, even if it affects the player inventory. After the window is closed, a number of these packets are sent to update the player's inventory window (0). */
	WindowID Byte // Byte
	/* The last received State ID from either a Set Container Slot or a Set Container Content packet */
	StateID VarInt // VarInt
	/* The slot that should be updated. */
	Slot Short // Short
	SlotData *Slot // Slot
}

var _ Packet = (*PlaySetContainerSlot_763_0)(nil)

func (p PlaySetContainerSlot_763_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.VarInt(p.StateID)
	b.Short(p.Slot)
	b.Encode(p.SlotData)
}

func (p *PlaySetContainerSlot_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.StateID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
	if err = p.SlotData.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x15
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x15
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x13
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x14
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x14
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x17
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x17
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x17
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x17
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x16
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x16
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x18
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x17
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x18
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x18
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x17
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x17
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x17
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x17
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x17
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x17
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x17
type PlaySetCooldown_763_0 struct {
	/* Numeric ID of the item to apply a cooldown to. */
	ItemID VarInt // VarInt
	/* Number of ticks to apply a cooldown for, or 0 to clear the cooldown. */
	CooldownTicks VarInt // VarInt
}

var _ Packet = (*PlaySetCooldown_763_0)(nil)

func (p PlaySetCooldown_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.ItemID)
	b.VarInt(p.CooldownTicks)
}

func (p *PlaySetCooldown_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ItemID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CooldownTicks, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x2b
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x2b
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x2b
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x2b
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x2a
type PlaySetCreativeModeSlot_763_0 struct {
	/* Inventory slot. */
	Slot Short // Short
	ClickedItem *Slot // Slot
}

var _ Packet = (*PlaySetCreativeModeSlot_763_0)(nil)

func (p PlaySetCreativeModeSlot_763_0)Encode(b *PacketBuilder){
	b.Short(p.Slot)
	b.Encode(p.ClickedItem)
}

func (p *PlaySetCreativeModeSlot_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
	if err = p.ClickedItem.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x50
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x50
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x4c
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x4d
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x4a
type PlaySetDefaultSpawnPosition_763_0 struct {
	/* Spawn location. */
	Location Position // Position
	/* The angle at which to respawn at. */
	Angle Float // Float
}

var _ Packet = (*PlaySetDefaultSpawnPosition_763_0)(nil)

func (p PlaySetDefaultSpawnPosition_763_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.Float(p.Angle)
}

func (p *PlaySetDefaultSpawnPosition_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Angle, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x2
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x2
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x2
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x2
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x2
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x2
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x2
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x2
type PlaySetDifficulty_758_0 struct {
	/* 0: peaceful, 1: easy, 2: normal, 3: hard */
	NewDifficulty Byte // Byte
}

var _ Packet = (*PlaySetDifficulty_758_0)(nil)

func (p PlaySetDifficulty_758_0)Encode(b *PacketBuilder){
	b.Byte(p.NewDifficulty)
}

func (p *PlaySetDifficulty_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.NewDifficulty, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x4e
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x4b
type PlaySetDisplayChatPreview_760_0 struct {
	ChatPreviewSetting Bool // Boolean
}

var _ Packet = (*PlaySetDisplayChatPreview_760_0)(nil)

func (p PlaySetDisplayChatPreview_760_0)Encode(b *PacketBuilder){
	b.Bool(p.ChatPreviewSetting)
}

func (p *PlaySetDisplayChatPreview_760_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChatPreviewSetting, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x1f
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x1f
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x1f
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x1f
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x1f
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x1e
type PlaySetDisplayedRecipe_758_0 struct {
	RecipeID Identifier // Identifier
}

var _ Packet = (*PlaySetDisplayedRecipe_758_0)(nil)

func (p PlaySetDisplayedRecipe_758_0)Encode(b *PacketBuilder){
	b.Encode(p.RecipeID)
}

func (p *PlaySetDisplayedRecipe_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.RecipeID.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x52
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x52
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x4e
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x50
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x4d
type PlaySetEntityMetadata_763_0 struct {
	EntityID VarInt // VarInt
	Metadata *EntityMetadata // Entity Metadata
}

var _ Packet = (*PlaySetEntityMetadata_763_0)(nil)

func (p PlaySetEntityMetadata_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.Metadata)
}

func (p *PlaySetEntityMetadata_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Metadata.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x54
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x54
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x50
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x52
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x4f
type PlaySetEntityVelocity_763_0 struct {
	EntityID VarInt // VarInt
	/* Velocity on the X axis. */
	VelocityX Short // Short
	/* Velocity on the Y axis. */
	VelocityY Short // Short
	/* Velocity on the Z axis. */
	VelocityZ Short // Short
}

var _ Packet = (*PlaySetEntityVelocity_763_0)(nil)

func (p PlaySetEntityVelocity_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
}

func (p *PlaySetEntityVelocity_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
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
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x56
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x56
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x52
type PlaySetExperience_763_0 struct {
	/* Between 0 and 1. */
	ExperienceBar Float // Float
	/* See Experience#Leveling up on the Minecraft Wiki for Total Experience to Level conversion. */
	TotalExperience VarInt // VarInt
	Level VarInt // VarInt
}

var _ Packet = (*PlaySetExperience_763_0)(nil)

func (p PlaySetExperience_763_0)Encode(b *PacketBuilder){
	b.Float(p.ExperienceBar)
	b.VarInt(p.TotalExperience)
	b.VarInt(p.Level)
}

func (p *PlaySetExperience_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ExperienceBar, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TotalExperience, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Level, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x54
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x51
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x51
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x51
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x51
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x51
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x48
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x48
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x48
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x47
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x43
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x43
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x40
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x40
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x3f
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x3d
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x3d
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x3d
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x3d
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x1f
type PlaySetExperience_760_1 struct {
	/* Between 0 and 1 */
	ExperienceBar Float // Float
	Level VarInt // VarInt
	/* See Experience#Leveling up on the Minecraft Wiki for Total Experience to Level conversion */
	TotalExperience VarInt // VarInt
}

var _ Packet = (*PlaySetExperience_760_1)(nil)

func (p PlaySetExperience_760_1)Encode(b *PacketBuilder){
	b.Float(p.ExperienceBar)
	b.VarInt(p.Level)
	b.VarInt(p.TotalExperience)
}

func (p *PlaySetExperience_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ExperienceBar, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Level, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.TotalExperience, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x42
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x42
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x3e
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x3f
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x3c
type PlaySetHeadRotation_763_0 struct {
	EntityID VarInt // VarInt
	/* New angle, not a delta. */
	HeadYaw Angle // Angle
}

var _ Packet = (*PlaySetHeadRotation_763_0)(nil)

func (p PlaySetHeadRotation_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.HeadYaw)
}

func (p *PlaySetHeadRotation_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.HeadYaw.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x57
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x57
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x53
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x55
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x52
type PlaySetHealth_763_0 struct {
	/* 0 or less = dead, 20 = full HP. */
	Health Float // Float
	/* 0–20. */
	Food VarInt // VarInt
	/* Seems to vary from 0.0 to 5.0 in integer increments. */
	FoodSaturation Float // Float
}

var _ Packet = (*PlaySetHealth_763_0)(nil)

func (p PlaySetHealth_763_0)Encode(b *PacketBuilder){
	b.Float(p.Health)
	b.VarInt(p.Food)
	b.Float(p.FoodSaturation)
}

func (p *PlaySetHealth_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Health, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Food, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.FoodSaturation, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x4d
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x4d
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x49
type PlaySetHeldItem_763_0 struct {
	/* The slot which the player has selected (0–8). */
	Slot Byte // Byte
}

var _ Packet = (*PlaySetHeldItem_763_0)(nil)

func (p PlaySetHeldItem_763_0)Encode(b *PacketBuilder){
	b.Byte(p.Slot)
}

func (p *PlaySetHeldItem_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x28
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x28
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x28
type PlaySetHeldItem_763_1 struct {
	/* The slot which the player has selected (0–8). */
	Slot Short // Short
}

var _ Packet = (*PlaySetHeldItem_763_1)(nil)

func (p PlaySetHeldItem_763_1)Encode(b *PacketBuilder){
	b.Short(p.Slot)
}

func (p *PlaySetHeldItem_763_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x4a
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x47
type PlaySetHeldItem_760_2 struct {
	/* The slot which the player has selected (0–8). */
	Slot Byte // Byte
}

var _ Packet = (*PlaySetHeldItem_760_2)(nil)

func (p PlaySetHeldItem_760_2)Encode(b *PacketBuilder){
	b.Byte(p.Slot)
}

func (p *PlaySetHeldItem_760_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x28
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x27
type PlaySetHeldItem_760_3 struct {
	/* The slot which the player has selected (0–8). */
	Slot Short // Short
}

var _ Packet = (*PlaySetHeldItem_760_3)(nil)

func (p PlaySetHeldItem_760_3)Encode(b *PacketBuilder){
	b.Short(p.Slot)
}

func (p *PlaySetHeldItem_760_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x59
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x59
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x55
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x57
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x54
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x54
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x54
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x54
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x54
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x4b
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x4b
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x4b
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x4a
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x46
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x46
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x43
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x43
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x42
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x40
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x40
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x40
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x40
type PlaySetPassengers_763_0 struct {
	/* Vehicle's EID */
	EntityID VarInt // VarInt
	/* Number of elements in the following array */
	PassengerCount VarInt // VarInt
	/* EIDs of entity's passengers */
	Passengers []VarInt // Array of VarInt
}

var _ Packet = (*PlaySetPassengers_763_0)(nil)

func (p PlaySetPassengers_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	p.PassengerCount = len(p.Passengers)
	b.VarInt(p.PassengerCount)
	for _, v := range p.Passengers {
		b.VarInt(v)
	}
}

func (p *PlaySetPassengers_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.PassengerCount, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Passengers = make([]VarInt, p.PassengerCount)
	for i, _ := range p.Passengers {
		if p.Passengers[i], ok = r.VarInt(); !ok {
			return io.EOF
		}
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x17
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x17
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x16
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x17
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x16
type PlaySetPlayerOnGround_763_0 struct {
	/* True if the client is on the ground, false otherwise. */
	OnGround Bool // Boolean
}

var _ Packet = (*PlaySetPlayerOnGround_763_0)(nil)

func (p PlaySetPlayerOnGround_763_0)Encode(b *PacketBuilder){
	b.Bool(p.OnGround)
}

func (p *PlaySetPlayerOnGround_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x14
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x14
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x13
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x14
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x13
type PlaySetPlayerPosition_763_0 struct {
	/* Absolute position. */
	X Double // Double
	/* Absolute feet position, normally Head Y - 1.62. */
	FeetY Double // Double
	/* Absolute position. */
	Z Double // Double
	/* True if the client is on the ground, false otherwise. */
	OnGround Bool // Boolean
}

var _ Packet = (*PlaySetPlayerPosition_763_0)(nil)

func (p PlaySetPlayerPosition_763_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.FeetY)
	b.Double(p.Z)
	b.Bool(p.OnGround)
}

func (p *PlaySetPlayerPosition_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.FeetY, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x15
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x15
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x14
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x15
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x14
type PlaySetPlayerPositionAndRotation_763_0 struct {
	/* Absolute position. */
	X Double // Double
	/* Absolute feet position, normally Head Y - 1.62. */
	FeetY Double // Double
	/* Absolute position. */
	Z Double // Double
	/* Absolute rotation on the X Axis, in degrees. */
	Yaw Float // Float
	/* Absolute rotation on the Y Axis, in degrees. */
	Pitch Float // Float
	/* True if the client is on the ground, false otherwise. */
	OnGround Bool // Boolean
}

var _ Packet = (*PlaySetPlayerPositionAndRotation_763_0)(nil)

func (p PlaySetPlayerPositionAndRotation_763_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.FeetY)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlaySetPlayerPositionAndRotation_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.FeetY, ok = r.Double(); !ok {
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
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x16
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x16
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x15
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x16
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x15
type PlaySetPlayerRotation_763_0 struct {
	/* Absolute rotation on the X Axis, in degrees. */
	Yaw Float // Float
	/* Absolute rotation on the Y Axis, in degrees. */
	Pitch Float // Float
	/* True if the client is on the ground, false otherwise. */
	OnGround Bool // Boolean
}

var _ Packet = (*PlaySetPlayerRotation_763_0)(nil)

func (p PlaySetPlayerRotation_763_0)Encode(b *PacketBuilder){
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlaySetPlayerRotation_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Yaw, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x1e
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x1e
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x1e
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x1e
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x1e
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x1f
type PlaySetRecipeBookState_758_0 struct {
	/* 0: crafting, 1: furnace, 2: blast furnace, 3: smoker */
	BookID VarInt // VarInt enum
	BookOpen Bool // Boolean
	FilterActive Bool // Boolean
}

var _ Packet = (*PlaySetRecipeBookState_758_0)(nil)

func (p PlaySetRecipeBookState_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.BookID)
	b.Bool(p.BookOpen)
	b.Bool(p.FilterActive)
}

func (p *PlaySetRecipeBookState_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.BookID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.BookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.FilterActive, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x4f
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x4f
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x4b
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x4c
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x49
type PlaySetRenderDistance_763_0 struct {
	/* Render distance (2-32). */
	ViewDistance VarInt // VarInt
}

var _ Packet = (*PlaySetRenderDistance_763_0)(nil)

func (p PlaySetRenderDistance_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.ViewDistance)
}

func (p *PlaySetRenderDistance_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ViewDistance, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x22
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x22
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x22
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x22
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x21
type PlaySetSeenRecipe_763_0 struct {
	RecipeID Identifier // Identifier
}

var _ Packet = (*PlaySetSeenRecipe_763_0)(nil)

func (p PlaySetSeenRecipe_763_0)Encode(b *PacketBuilder){
	b.Encode(p.RecipeID)
}

func (p *PlaySetSeenRecipe_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.RecipeID.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x5c
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x5c
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x58
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x5a
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x57
type PlaySetSimulationDistance_763_0 struct {
	/* The distance that the client will process specific things, such as entities. */
	SimulationDistance VarInt // VarInt
}

var _ Packet = (*PlaySetSimulationDistance_763_0)(nil)

func (p PlaySetSimulationDistance_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.SimulationDistance)
}

func (p *PlaySetSimulationDistance_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SimulationDistance, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x16
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x16
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x16
type PlaySetSlot_758_0 struct {
	/* The window which is being updated. 0 for player inventory. Note that all known window types include the player inventory. This packet will only be sent for the currently opened window while the player is performing actions, even if it affects the player inventory. After the window is closed, a number of these packets are sent to update the player's inventory window (0). */
	WindowID Byte // Byte
	/* See State ID */
	StateID VarInt // VarInt
	/* The slot that should be updated. */
	Slot Short // Short
	SlotData *Slot // Slot
}

var _ Packet = (*PlaySetSlot_758_0)(nil)

func (p PlaySetSlot_758_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.VarInt(p.StateID)
	b.Short(p.Slot)
	b.Encode(p.SlotData)
}

func (p *PlaySetSlot_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.StateID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
	if err = p.SlotData.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x16
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x15
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x15
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x17
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x16
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x17
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x17
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x16
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x16
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x16
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x16
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x16
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x16
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x16
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x2f
type PlaySetSlot_755_1 struct {
	/* The window which is being updated. 0 for player inventory. Note that all known window types include the player inventory. This packet will only be sent for the currently opened window while the player is performing actions, even if it affects the player inventory. After the window is closed, a number of these packets are sent to update the player's inventory window (0). */
	WindowID Byte // Byte
	/* The slot that should be updated */
	Slot Short // Short
	SlotData *Slot // Slot
}

var _ Packet = (*PlaySetSlot_755_1)(nil)

func (p PlaySetSlot_755_1)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Short(p.Slot)
	b.Encode(p.SlotData)
}

func (p *PlaySetSlot_755_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Slot, ok = r.Short(); !ok {
		return io.EOF
	}
	if err = p.SlotData.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x5d
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x5d
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x59
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x5b
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x58
type PlaySetSubtitleText_763_0 struct {
	SubtitleText Object // Chat
}

var _ Packet = (*PlaySetSubtitleText_763_0)(nil)

func (p PlaySetSubtitleText_763_0)Encode(b *PacketBuilder){
	b.JSON(p.SubtitleText)
}

func (p *PlaySetSubtitleText_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.SubtitleText); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x65
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x65
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x61
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x63
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x60
type PlaySetTabListHeaderAndFooter_763_0 struct {
	/* To remove the header, send a empty text component: {"text":""}. */
	Header Object // Chat
	/* To remove the footer, send a empty text component: {"text":""}. */
	Footer Object // Chat
}

var _ Packet = (*PlaySetTabListHeaderAndFooter_763_0)(nil)

func (p PlaySetTabListHeaderAndFooter_763_0)Encode(b *PacketBuilder){
	b.JSON(p.Header)
	b.JSON(p.Footer)
}

func (p *PlaySetTabListHeaderAndFooter_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.Header); err != nil {
		return
	}
	if err = r.JSON(&p.Footer); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x60
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x60
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x5c
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x5e
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x5b
type PlaySetTitleAnimationTimes_763_0 struct {
	/* Ticks to spend fading in. */
	FadeIn Int // Int
	/* Ticks to keep the title displayed. */
	Stay Int // Int
	/* Ticks to spend fading out, not when to start fading out. */
	FadeOut Int // Int
}

var _ Packet = (*PlaySetTitleAnimationTimes_763_0)(nil)

func (p PlaySetTitleAnimationTimes_763_0)Encode(b *PacketBuilder){
	b.Int(p.FadeIn)
	b.Int(p.Stay)
	b.Int(p.FadeOut)
}

func (p *PlaySetTitleAnimationTimes_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.FadeIn, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Stay, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.FadeOut, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x58
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x58
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x57
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x57
type PlaySetTitleSubTitle_758_0 struct {
	SubtitleText Object // Chat
}

var _ Packet = (*PlaySetTitleSubTitle_758_0)(nil)

func (p PlaySetTitleSubTitle_758_0)Encode(b *PacketBuilder){
	b.JSON(p.SubtitleText)
}

func (p *PlaySetTitleSubTitle_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.SubtitleText); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x5f
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x5f
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x5b
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x5d
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x5a
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x5a
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x5a
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x59
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x59
type PlaySetTitleText_763_0 struct {
	TitleText Object // Chat
}

var _ Packet = (*PlaySetTitleText_763_0)(nil)

func (p PlaySetTitleText_763_0)Encode(b *PacketBuilder){
	b.JSON(p.TitleText)
}

func (p *PlaySetTitleText_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.TitleText); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x5b
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x5b
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x5a
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x5a
type PlaySetTitleTimes_758_0 struct {
	/* Ticks to spend fading in. */
	FadeIn Int // Int
	/* Ticks to keep the title displayed. */
	Stay Int // Int
	/* Ticks to spend out, not when to start fading out. */
	FadeOut Int // Int
}

var _ Packet = (*PlaySetTitleTimes_758_0)(nil)

func (p PlaySetTitleTimes_758_0)Encode(b *PacketBuilder){
	b.Int(p.FadeIn)
	b.Int(p.Stay)
	b.Int(p.FadeOut)
}

func (p *PlaySetTitleTimes_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.FadeIn, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.Stay, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.FadeOut, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x62
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x62
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x5e
type PlaySoundEffect_763_0 struct {
	/* Represents the Sound ID + 1. If the value is 0, the packet contains a sound specified by Identifier. */
	SoundID VarInt // VarInt
	/* Only present if Sound ID is 0 */
	SoundName Identifier // Optional Identifier
	/* Only present if Sound ID is 0. */
	HasFixedRange Bool // Optional Boolean
	/* The fixed range of the sound. Only present if previous boolean is true and Sound ID is 0. */
	Range Float // Optional Float
	/* The category that this sound will be played from (current categories). */
	SoundCategory VarInt // VarInt Enum
	/* Effect X multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionX Int // Int
	/* Effect Y multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionY Int // Int
	/* Effect Z multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionZ Int // Int
	/* 1.0 is 100%, capped between 0.0 and 1.0 by Notchian clients. */
	Volume Float // Float
	/* Float between 0.5 and 2.0 by Notchian clients. */
	Pitch Float // Float
	/* Seed used to pick sound variant. */
	Seed Long // Long
}

var _ Packet = (*PlaySoundEffect_763_0)(nil)

func (p PlaySoundEffect_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.SoundID)
	b.Encode(p.SoundName)
	b.Bool(p.HasFixedRange)
	b.Float(p.Range)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.Float(p.Pitch)
	b.Long(p.Seed)
}

func (p *PlaySoundEffect_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SoundID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.SoundName.DecodeFrom(r); err != nil {
		return
	}
	if p.HasFixedRange, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Range, ok = r.Float(); !ok {
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
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Seed, ok = r.Long(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x60
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x5d
type PlaySoundEffect_760_1 struct {
	/* ID of hardcoded sound event (events). */
	SoundID VarInt // VarInt
	/* The category that this sound will be played from (current categories). */
	SoundCategory VarInt // VarInt Enum
	/* Effect X multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionX Int // Int
	/* Effect Y multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionY Int // Int
	/* Effect Z multiplied by 8 (fixed-point number with only 3 bits dedicated to the fractional part). */
	EffectPositionZ Int // Int
	/* 1.0 is 100%, capped between 0.0 and 1.0 by Notchian clients. */
	Volume Float // Float
	/* Float between 0.5 and 2.0 by Notchian clients. */
	Pitch Float // Float
	/* Seed used to pick sound variant. */
	Seed Long // long
}

var _ Packet = (*PlaySoundEffect_760_1)(nil)

func (p PlaySoundEffect_760_1)Encode(b *PacketBuilder){
	b.VarInt(p.SoundID)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.Float(p.Pitch)
	b.Long(p.Seed)
}

func (p *PlaySoundEffect_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Seed, ok = r.Long(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x5d
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x5d
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x5c
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x5c
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x51
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x51
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x52
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x51
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x4d
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x4d
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x49
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x49
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x48
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x46
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x46
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x46
type PlaySoundEffect_758_2 struct {
	/* ID of hardcoded sound event (events as of 1.11.0) */
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
	/* Float between 0.5 and 2.0 by Notchian clients */
	Pitch Float // Float
}

var _ Packet = (*PlaySoundEffect_758_2)(nil)

func (p PlaySoundEffect_758_2)Encode(b *PacketBuilder){
	b.VarInt(p.SoundID)
	b.VarInt(p.SoundCategory)
	b.Int(p.EffectPositionX)
	b.Int(p.EffectPositionY)
	b.Int(p.EffectPositionZ)
	b.Float(p.Volume)
	b.Float(p.Pitch)
}

func (p *PlaySoundEffect_758_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.Pitch, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x1
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x1
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x0
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x0
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x0
type PlaySpawnEntity_763_0 struct {
	/* A unique integer ID mostly used in the protocol to identify the entity. */
	EntityID VarInt // VarInt
	/* A unique identifier that is mostly used in persistence and places where the uniqueness matters more. */
	EntityUUID UUID // UUID
	/* The type of the entity (see "type" field of the list of Mob types). */
	Type VarInt // VarInt
	X Double // Double
	Y Double // Double
	Z Double // Double
	/* To get the real pitch, you must divide this by (256.0F / 360.0F) */
	Pitch Angle // Angle
	/* To get the real yaw, you must divide this by (256.0F / 360.0F) */
	Yaw Angle // Angle
	/* Only used by living entities, where the head of the entity may differ from the general body rotation. */
	HeadYaw Angle // Angle
	/* Meaning dependent on the value of the Type field, see Object Data for details. */
	Data VarInt // VarInt
	/* Same units as Set Entity Velocity. */
	VelocityX Short // Short
	VelocityY Short // Short
	VelocityZ Short // Short
}

var _ Packet = (*PlaySpawnEntity_763_0)(nil)

func (p PlaySpawnEntity_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.EntityUUID)
	b.VarInt(p.Type)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Pitch)
	b.Encode(p.Yaw)
	b.Encode(p.HeadYaw)
	b.VarInt(p.Data)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
}

func (p *PlaySpawnEntity_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.HeadYaw.DecodeFrom(r); err != nil {
		return
	}
	if p.Data, ok = r.VarInt(); !ok {
		return io.EOF
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
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x0
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x0
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x0
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x0
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x0
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x0
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x0
type PlaySpawnEntity_758_1 struct {
	/* EID of the entity */
	EntityID VarInt // VarInt
	ObjectUUID UUID // UUID
	/* The type of entity (same as in Spawn Living Entity) */
	Type VarInt // VarInt
	X Double // Double
	Y Double // Double
	Z Double // Double
	Pitch Angle // Angle
	Yaw Angle // Angle
	/* Meaning dependent on the value of the Type field, see Object Data for details. */
	Data Int // Int
	/* Same units as Entity Velocity.  Always sent, but only used when Data is greater than 0 (except for some entities which always ignore it; see Object Data for details). */
	VelocityX Short // Short
	VelocityY Short // Short
	VelocityZ Short // Short
}

var _ Packet = (*PlaySpawnEntity_758_1)(nil)

func (p PlaySpawnEntity_758_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.ObjectUUID)
	b.VarInt(p.Type)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Pitch)
	b.Encode(p.Yaw)
	b.Int(p.Data)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
}

func (p *PlaySpawnEntity_758_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ObjectUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if p.Data, ok = r.Int(); !ok {
		return io.EOF
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
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x2
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x2
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x1
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x1
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x1
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x1
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x1
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x1
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x1
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x1
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x1
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x1
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x1
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x1
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x1
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x1
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x1
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x1
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x1
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x1
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x1
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x1
type PlaySpawnExperienceOrb_763_0 struct {
	EntityID VarInt // VarInt
	X Double // Double
	Y Double // Double
	Z Double // Double
	/* The amount of experience this orb will reward once collected */
	Count Short // Short
}

var _ Packet = (*PlaySpawnExperienceOrb_763_0)(nil)

func (p PlaySpawnExperienceOrb_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Short(p.Count)
}

func (p *PlaySpawnExperienceOrb_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Count, ok = r.Short(); !ok {
		return io.EOF
	}
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x2
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x2
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x2
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x2
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x2
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x2
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x2
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x2
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x2
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x2
type PlaySpawnGlobalEntity_498_0 struct {
	/* The EID of the thunderbolt */
	EntityID VarInt // VarInt
	/* The global entity type, currently always 1 for thunderbolt */
	Type Byte // Byte Enum
	X Double // Double
	Y Double // Double
	Z Double // Double
}

var _ Packet = (*PlaySpawnGlobalEntity_498_0)(nil)

func (p PlaySpawnGlobalEntity_498_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Byte(p.Type)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
}

func (p *PlaySpawnGlobalEntity_498_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
}

// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x2
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x2
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x2
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x2
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x2
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x3
type PlaySpawnLivingEntity_757_1 struct {
	EntityID VarInt // VarInt
	EntityUUID UUID // UUID
	/* The type of mob. See Entities#Mobs */
	Type VarInt // VarInt
	X Double // Double
	Y Double // Double
	Z Double // Double
	Yaw Angle // Angle
	Pitch Angle // Angle
	HeadPitch Angle // Angle
	/* Same units as Entity Velocity */
	VelocityX Short // Short
	/* Same units as Entity Velocity */
	VelocityY Short // Short
	/* Same units as Entity Velocity */
	VelocityZ Short // Short
}

var _ Packet = (*PlaySpawnLivingEntity_757_1)(nil)

func (p PlaySpawnLivingEntity_757_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.EntityUUID)
	b.VarInt(p.Type)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Encode(p.HeadPitch)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
}

func (p *PlaySpawnLivingEntity_757_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if err = p.HeadPitch.DecodeFrom(r); err != nil {
		return
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
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x3
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x3
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x3
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x3
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x3
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x3
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x3
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x3
type PlaySpawnMob_498_0 struct {
	EntityID VarInt // VarInt
	EntityUUID UUID // UUID
	/* The type of mob. See Entities#Mobs */
	Type VarInt // VarInt
	X Double // Double
	Y Double // Double
	Z Double // Double
	Yaw Angle // Angle
	Pitch Angle // Angle
	HeadPitch Angle // Angle
	/* Same units as Entity Velocity */
	VelocityX Short // Short
	/* Same units as Entity Velocity */
	VelocityY Short // Short
	/* Same units as Entity Velocity */
	VelocityZ Short // Short
	Metadata *EntityMetadata // Entity Metadata
}

var _ Packet = (*PlaySpawnMob_498_0)(nil)

func (p PlaySpawnMob_498_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.EntityUUID)
	b.VarInt(p.Type)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Encode(p.HeadPitch)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
	b.Encode(p.Metadata)
}

func (p *PlaySpawnMob_498_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if err = p.HeadPitch.DecodeFrom(r); err != nil {
		return
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
	if err = p.Metadata.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x3
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x3
type PlaySpawnMob_210_1 struct {
	EntityID VarInt // VarInt
	EntityUUID UUID // UUID
	/* The type of mob. See Entities#Mobs */
	Type UByte // Unsigned Byte
	X Double // Double
	Y Double // Double
	Z Double // Double
	Yaw Angle // Angle
	Pitch Angle // Angle
	HeadPitch Angle // Angle
	/* Same units as Entity Velocity */
	VelocityX Short // Short
	/* Same units as Entity Velocity */
	VelocityY Short // Short
	/* Same units as Entity Velocity */
	VelocityZ Short // Short
	Metadata *EntityMetadata // Entity Metadata
}

var _ Packet = (*PlaySpawnMob_210_1)(nil)

func (p PlaySpawnMob_210_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.EntityUUID)
	b.UByte(p.Type)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Encode(p.HeadPitch)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
	b.Encode(p.Metadata)
}

func (p *PlaySpawnMob_210_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if err = p.HeadPitch.DecodeFrom(r); err != nil {
		return
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
	if err = p.Metadata.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x0
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x0
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x0
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x0
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x0
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x0
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x0
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x0
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x0
type PlaySpawnObject_404_1 struct {
	/* EID of the object */
	EntityID VarInt // VarInt
	ObjectUUID UUID // UUID
	/* The type of object (see Entities#Objects) */
	Type Byte // Byte
	X Double // Double
	Y Double // Double
	Z Double // Double
	Pitch Angle // Angle
	Yaw Angle // Angle
	/* Meaning dependent on the value of the Type field, see Object Data for details. */
	Data Int // Int
	/* Same units as Entity Velocity.  Always sent, but only used when Data is greater than 0 (except for some entities which always ignore it; see Object Data for details). */
	VelocityX Short // Short
	VelocityY Short // Short
	VelocityZ Short // Short
}

var _ Packet = (*PlaySpawnObject_404_1)(nil)

func (p PlaySpawnObject_404_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.ObjectUUID)
	b.Byte(p.Type)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Pitch)
	b.Encode(p.Yaw)
	b.Int(p.Data)
	b.Short(p.VelocityX)
	b.Short(p.VelocityY)
	b.Short(p.VelocityZ)
}

func (p *PlaySpawnObject_404_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ObjectUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if p.Data, ok = r.Int(); !ok {
		return io.EOF
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
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x3
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x3
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x3
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x3
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x3
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x3
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x4
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x4
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x4
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x4
type PlaySpawnPainting_758_0 struct {
	EntityID VarInt // VarInt
	EntityUUID UUID // UUID
	/* Panting's ID, see below */
	Motive VarInt // VarInt
	/* Center coordinates (see below) */
	Location Position // Position
	/* Direction the painting faces (North = 2, South = 0, West = 1, East = 3) */
	Direction Byte // Byte Enum
}

var _ Packet = (*PlaySpawnPainting_758_0)(nil)

func (p PlaySpawnPainting_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.EntityUUID)
	b.VarInt(p.Motive)
	b.Encode(p.Location)
	b.Byte(p.Direction)
}

func (p *PlaySpawnPainting_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Motive, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Direction, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x4
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x4
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x4
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x4
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x4
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x4
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x4
type PlaySpawnPainting_340_1 struct {
	EntityID VarInt // VarInt
	EntityUUID UUID // UUID
	/* Name of the painting. Max length 13 */
	Title String // String (13)
	/* Center coordinates (see below) */
	Location Position // Position
	/* Direction the painting faces (North = 2, South = 0, West = 1, East = 3) */
	Direction Byte // Byte Enum
}

var _ Packet = (*PlaySpawnPainting_340_1)(nil)

func (p PlaySpawnPainting_340_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.EntityUUID)
	b.String(p.Title)
	b.Encode(p.Location)
	b.Byte(p.Direction)
}

func (p *PlaySpawnPainting_340_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.EntityUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Title, ok = r.String(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Direction, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x3
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x3
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x2
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x2
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x2
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x4
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x4
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x4
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x4
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x4
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x4
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x5
type PlaySpawnPlayer_763_0 struct {
	/* Player's EID */
	EntityID VarInt // VarInt
	/* See below for notes on offline mode and NPCs */
	PlayerUUID UUID // UUID
	X Double // Double
	Y Double // Double
	Z Double // Double
	Yaw Angle // Angle
	Pitch Angle // Angle
}

var _ Packet = (*PlaySpawnPlayer_763_0)(nil)

func (p PlaySpawnPlayer_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.PlayerUUID)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
}

func (p *PlaySpawnPlayer_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.PlayerUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x5
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x5
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x5
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x5
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x5
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x5
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x5
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x5
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x5
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x5
type PlaySpawnPlayer_498_1 struct {
	/* Player's EID */
	EntityID VarInt // VarInt
	/* See below for notes on offline mode and NPCs */
	PlayerUUID UUID // UUID
	X Double // Double
	Y Double // Double
	Z Double // Double
	Yaw Angle // Angle
	Pitch Angle // Angle
	Metadata *EntityMetadata // Entity Metadata
}

var _ Packet = (*PlaySpawnPlayer_498_1)(nil)

func (p PlaySpawnPlayer_498_1)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.UUID(p.PlayerUUID)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Encode(p.Metadata)
}

func (p *PlaySpawnPlayer_498_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.PlayerUUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Metadata.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x4b
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x4b
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x4b
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x4b
type PlaySpawnPosition_758_0 struct {
	/* Spawn location. */
	Location Position // Position
	/* The angle at which to respawn at. */
	Angle Float // Float
}

var _ Packet = (*PlaySpawnPosition_758_0)(nil)

func (p PlaySpawnPosition_758_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.Float(p.Angle)
}

func (p *PlaySpawnPosition_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Angle, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x42
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x42
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x4e
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x4d
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x49
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x49
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x46
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x46
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x45
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x43
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x43
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x43
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x43
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x5
type PlaySpawnPosition_754_1 struct {
	/* Spawn location */
	Location Position // Position
}

var _ Packet = (*PlaySpawnPosition_754_1)(nil)

func (p PlaySpawnPosition_754_1)Encode(b *PacketBuilder){
	b.Encode(p.Location)
}

func (p *PlaySpawnPosition_754_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x2d
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x2d
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x2d
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x2d
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x2d
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x2d
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x2b
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x2b
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x28
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x28
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x1e
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x1e
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x1e
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x1b
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x1b
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x1b
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x1b
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x18
type PlaySpectate_758_0 struct {
	/* UUID of the player to teleport to (can also be an entity UUID) */
	TargetPlayer UUID // UUID
}

var _ Packet = (*PlaySpectate_758_0)(nil)

func (p PlaySpectate_758_0)Encode(b *PacketBuilder){
	b.UUID(p.TargetPlayer)
}

func (p *PlaySpectate_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TargetPlayer, ok = r.UUID(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x16
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x16
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x16
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x16
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x17
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x17
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x16
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x16
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x14
type PlaySteerBoat_758_0 struct {
	LeftPaddleTurning Bool // Boolean
	RightPaddleTurning Bool // Boolean
}

var _ Packet = (*PlaySteerBoat_758_0)(nil)

func (p PlaySteerBoat_758_0)Encode(b *PacketBuilder){
	b.Bool(p.LeftPaddleTurning)
	b.Bool(p.RightPaddleTurning)
}

func (p *PlaySteerBoat_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.LeftPaddleTurning, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.RightPaddleTurning, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x14
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x11
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x11
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x12
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x11
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x11
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x11
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x11
type PlaySteerBoat_401_1 struct {
	RightPaddleTurning Bool // Boolean
	LeftPaddleTurning Bool // Boolean
}

var _ Packet = (*PlaySteerBoat_401_1)(nil)

func (p PlaySteerBoat_401_1)Encode(b *PacketBuilder){
	b.Bool(p.RightPaddleTurning)
	b.Bool(p.LeftPaddleTurning)
}

func (p *PlaySteerBoat_401_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.RightPaddleTurning, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.LeftPaddleTurning, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x1c
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x1c
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x1c
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x1c
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x1d
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x1d
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x1c
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x1c
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x1a
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x1a
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x16
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x16
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x16
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x15
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x15
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x15
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x15
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0xc
type PlaySteerVehicle_758_0 struct {
	/* Positive to the left of the player */
	Sideways Float // Float
	/* Positive forward */
	Forward Float // Float
	/* Bit mask. 0x1: jump, 0x2: unmount */
	Flags UByte // Unsigned Byte
}

var _ Packet = (*PlaySteerVehicle_758_0)(nil)

func (p PlaySteerVehicle_758_0)Encode(b *PacketBuilder){
	b.Float(p.Sideways)
	b.Float(p.Forward)
	b.UByte(p.Flags)
}

func (p *PlaySteerVehicle_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Sideways, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Forward, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.UByte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x63
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x63
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x5f
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x61
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x5e
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x5e
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x5e
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x5d
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x5d
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x52
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x52
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x53
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x52
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x4c
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x4c
type PlayStopSound_763_0 struct {
	/* Controls which fields are present. */
	Flags Byte // Byte
	/* Only if flags is 3 or 1 (bit mask 0x1).  See below.  If not present, then sounds from all sources are cleared. */
	Source VarInt // Optional VarInt enum
	/* Only if flags is 2 or 3 (bit mask 0x2).  A sound effect name, see Named Sound Effect.  If not present, then all sounds are cleared. */
	Sound Identifier // Optional Identifier
}

var _ Packet = (*PlayStopSound_763_0)(nil)

func (p PlayStopSound_763_0)Encode(b *PacketBuilder){
	b.Byte(p.Flags)
	b.VarInt(p.Source)
	b.Encode(p.Sound)
}

func (p *PlayStopSound_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Source, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Sound.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x2f
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x2f
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x2f
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x2f
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x2e
type PlaySwingArm_763_0 struct {
	/* Hand used for the animation. 0: main hand, 1: off hand. */
	Hand VarInt // VarInt Enum
}

var _ Packet = (*PlaySwingArm_763_0)(nil)

func (p PlaySwingArm_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
}

func (p *PlaySwingArm_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x3c
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x3c
type PlaySynchronizePlayerPosition_763_0 struct {
	/* Absolute or relative position, depending on Flags. */
	X Double // Double
	/* Absolute or relative position, depending on Flags. */
	Y Double // Double
	/* Absolute or relative position, depending on Flags. */
	Z Double // Double
	/* Absolute or relative rotation on the X axis, in degrees. */
	Yaw Float // Float
	/* Absolute or relative rotation on the Y axis, in degrees. */
	Pitch Float // Float
	/* Bit field, see below. */
	Flags Byte // Byte
	/* Client should confirm this packet with Confirm Teleportation containing the same Teleport ID. */
	TeleportID VarInt // VarInt
}

var _ Packet = (*PlaySynchronizePlayerPosition_763_0)(nil)

func (p PlaySynchronizePlayerPosition_763_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Byte(p.Flags)
	b.VarInt(p.TeleportID)
}

func (p *PlaySynchronizePlayerPosition_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.TeleportID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x38
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x39
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x36
type PlaySynchronizePlayerPosition_761_1 struct {
	/* Absolute or relative position, depending on Flags. */
	X Double // Double
	/* Absolute or relative position, depending on Flags. */
	Y Double // Double
	/* Absolute or relative position, depending on Flags. */
	Z Double // Double
	/* Absolute or relative rotation on the X axis, in degrees. */
	Yaw Float // Float
	/* Absolute or relative rotation on the Y axis, in degrees. */
	Pitch Float // Float
	/* Bit field, see below. */
	Flags Byte // Byte
	/* Client should confirm this packet with Confirm Teleportation containing the same Teleport ID. */
	TeleportID VarInt // VarInt
	/* True if the player should dismount their vehicle. */
	DismountVehicle Bool // Boolean
}

var _ Packet = (*PlaySynchronizePlayerPosition_761_1)(nil)

func (p PlaySynchronizePlayerPosition_761_1)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
	b.Byte(p.Flags)
	b.VarInt(p.TeleportID)
	b.Bool(p.DismountVehicle)
}

func (p *PlaySynchronizePlayerPosition_761_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
	if p.TeleportID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DismountVehicle, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x64
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x64
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x60
type PlaySystemChatMessage_763_0 struct {
	/* Limited to 262144 bytes. */
	Content Object // Chat
	/* Whether the message is an actionbar or chat message. */
	Overlay Bool // Boolean
}

var _ Packet = (*PlaySystemChatMessage_763_0)(nil)

func (p PlaySystemChatMessage_763_0)Encode(b *PacketBuilder){
	b.JSON(p.Content)
	b.Bool(p.Overlay)
}

func (p *PlaySystemChatMessage_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.Content); err != nil {
		return
	}
	if p.Overlay, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x62
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x5f
type PlaySystemChatMessage_760_1 struct {
	/* Limited to 262144 bytes. */
	JSONData Object // Chat
	/* Whether the message is an actionbar or chat message. */
	Overlay Bool // Boolean
}

var _ Packet = (*PlaySystemChatMessage_760_1)(nil)

func (p PlaySystemChatMessage_760_1)Encode(b *PacketBuilder){
	b.JSON(p.JSONData)
	b.Bool(p.Overlay)
}

func (p *PlaySystemChatMessage_760_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = r.JSON(&p.JSONData); err != nil {
		return
	}
	if p.Overlay, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x6
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x6
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x6
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x6
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x6
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x6
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x6
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x6
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x5
type PlayTabComplete_758_0 struct {
	/* The id received in the tab completion request packet, must match or the client will ignore this packet.  Client generates this and increments it each time it sends another tab completion that doesn't get a response. */
	TransactionId VarInt // VarInt
	/* All text behind the cursor without the / (e.g. to the left of the cursor in left-to-right languages like English) */
	Text String // String (32500)
}

var _ Packet = (*PlayTabComplete_758_0)(nil)

func (p PlayTabComplete_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionId)
	b.String(p.Text)
}

func (p *PlayTabComplete_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionId, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Text, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0xe
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0xe
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0xe
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x3a
type PlayTabComplete_340_3 struct {
	/* Number of elements in the following array */
	Count VarInt // VarInt
	/* One eligible command, note that each command is sent separately instead of in a single string, hence the need for Count */
	Matches []String // Array of String
}

var _ Packet = (*PlayTabComplete_340_3)(nil)

func (p PlayTabComplete_340_3)Encode(b *PacketBuilder){
	p.Count = len(p.Matches)
	b.VarInt(p.Count)
	for _, v := range p.Matches {
		b.String(v)
	}
}

func (p *PlayTabComplete_340_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Matches = make([]String, p.Count)
	for i, _ := range p.Matches {
		if p.Matches[i], ok = r.String(); !ok {
			return io.EOF
		}
	}
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x1
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x1
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x1
type PlayTabComplete_340_4 struct {
	/* All text behind the cursor (e.g. to the left of the cursor in left-to-right languages like English) */
	Text String // String (32767)
	/* If true, the server will parse Text as a command even if it doesn't start with a /. Used in the command block GUI. */
	AssumeCommand Bool // Boolean
	HasPosition Bool // Boolean
	/* The position of the block being looked at. Only sent if Has Position is true. */
	LookedAtBlock Position // Optional Position
}

var _ Packet = (*PlayTabComplete_340_4)(nil)

func (p PlayTabComplete_340_4)Encode(b *PacketBuilder){
	b.String(p.Text)
	b.Bool(p.AssumeCommand)
	b.Bool(p.HasPosition)
	b.Encode(p.LookedAtBlock)
}

func (p *PlayTabComplete_340_4)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Text, ok = r.String(); !ok {
		return io.EOF
	}
	if p.AssumeCommand, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.HasPosition, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = p.LookedAtBlock.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0xe
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0xe
type PlayTabComplete_335_5 struct {
	/* Number of elements in the following array */
	Count VarInt // VarInt
	/* One eligible command, note that each command is sent separately instead of in a single string, hence the need for Count */
	Matches []String // Array of String
}

var _ Packet = (*PlayTabComplete_335_5)(nil)

func (p PlayTabComplete_335_5)Encode(b *PacketBuilder){
	p.Count = len(p.Matches)
	b.VarInt(p.Count)
	for _, v := range p.Matches {
		b.String(v)
	}
}

func (p *PlayTabComplete_335_5)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Matches = make([]String, p.Count)
	for i, _ := range p.Matches {
		if p.Matches[i], ok = r.String(); !ok {
			return io.EOF
		}
	}
}

// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x2
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x1
type PlayTabComplete_335_6 struct {
	/* All text behind the cursor (e.g. to the left of the cursor in left-to-right languages like English) */
	Text String // String
	/* If true, the server will parse Text as a command even if it doesn't start with a /. Used in the command block GUI. */
	AssumeCommand Bool // Boolean
	HasPosition Bool // Boolean
	/* The position of the block being looked at. Only sent if Has Position is true. */
	LookedAtBlock Position // Optional Position
}

var _ Packet = (*PlayTabComplete_335_6)(nil)

func (p PlayTabComplete_335_6)Encode(b *PacketBuilder){
	b.String(p.Text)
	b.Bool(p.AssumeCommand)
	b.Bool(p.HasPosition)
	b.Encode(p.LookedAtBlock)
}

func (p *PlayTabComplete_335_6)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Text, ok = r.String(); !ok {
		return io.EOF
	}
	if p.AssumeCommand, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.HasPosition, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = p.LookedAtBlock.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0xe
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0xe
type PlayTabComplete_210_7 struct {
	/* Number of elements in the following array */
	Count VarInt // VarInt
	/* One eligible command, note that each command is sent separately instead of in a single string, hence the need for Count */
	Matches []String // Array of String
}

var _ Packet = (*PlayTabComplete_210_7)(nil)

func (p PlayTabComplete_210_7)Encode(b *PacketBuilder){
	p.Count = len(p.Matches)
	b.VarInt(p.Count)
	for _, v := range p.Matches {
		b.String(v)
	}
}

func (p *PlayTabComplete_210_7)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Matches = make([]String, p.Count)
	for i, _ := range p.Matches {
		if p.Matches[i], ok = r.String(); !ok {
			return io.EOF
		}
	}
}

// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x1
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x1
type PlayTabComplete_210_8 struct {
	/* All text behind the cursor (e.g. to the left of the cursor in left-to-right languages like English) */
	Text String // String
	/* If true, the server will parse Text as a command even if it doesn't start with a /. Used in the command block GUI. */
	AssumeCommand Bool // Boolean
	HasPosition Bool // Boolean
	/* The position of the block being looked at. Only sent if Has Position is true. */
	LookedAtBlock Position // Optional Position
}

var _ Packet = (*PlayTabComplete_210_8)(nil)

func (p PlayTabComplete_210_8)Encode(b *PacketBuilder){
	b.String(p.Text)
	b.Bool(p.AssumeCommand)
	b.Bool(p.HasPosition)
	b.Encode(p.LookedAtBlock)
}

func (p *PlayTabComplete_210_8)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Text, ok = r.String(); !ok {
		return io.EOF
	}
	if p.AssumeCommand, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.HasPosition, ok = r.Bool(); !ok {
		return io.EOF
	}
	if err = p.LookedAtBlock.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x66
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x66
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x62
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x64
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x61
type PlayTagQueryResponse_763_0 struct {
	/* Can be compared to the one sent in the original query packet. */
	TransactionID VarInt // VarInt
	/* The NBT of the block or entity.  May be a TAG_END (0) in which case no NBT is present. */
	NBT NBT // NBT Tag
}

var _ Packet = (*PlayTagQueryResponse_763_0)(nil)

func (p PlayTagQueryResponse_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.TransactionID)
	b.Encode(p.NBT)
}

func (p *PlayTagQueryResponse_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TransactionID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.NBT.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x5b
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x5b
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x5c
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x5b
type PlayTags_754_0 struct {
	/* IDs are block IDs */
	BlockTags Unknown_ // (See below)
	/* IDs are item IDs */
	ItemTags Unknown_ // (See below)
	/* IDs are fluid IDs */
	FluidTags Unknown_ // (See below)
	/* IDs are entity IDs */
	EntityTags Unknown_ // (See below)
}

var _ Packet = (*PlayTags_754_0)(nil)

func (p PlayTags_754_0)Encode(b *PacketBuilder){
	b.Encode(p.BlockTags)
	b.Encode(p.ItemTags)
	b.Encode(p.FluidTags)
	b.Encode(p.EntityTags)
}

func (p *PlayTags_754_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.BlockTags.DecodeFrom(r); err != nil {
		return
	}
	if err = p.ItemTags.DecodeFrom(r); err != nil {
		return
	}
	if err = p.FluidTags.DecodeFrom(r); err != nil {
		return
	}
	if err = p.EntityTags.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x55
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x55
type PlayTags_404_1 struct {
	/* IDs are block IDs */
	BlockTags Unknown_ // (See below)
	/* IDs are item IDs */
	ItemTags Unknown_ // (See below)
	/* IDs are fluid IDs */
	FluidTags Unknown_ // (See below)
}

var _ Packet = (*PlayTags_404_1)(nil)

func (p PlayTags_404_1)Encode(b *PacketBuilder){
	b.Encode(p.BlockTags)
	b.Encode(p.ItemTags)
	b.Encode(p.FluidTags)
}

func (p *PlayTags_404_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.BlockTags.DecodeFrom(r); err != nil {
		return
	}
	if err = p.ItemTags.DecodeFrom(r); err != nil {
		return
	}
	if err = p.FluidTags.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x0
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x0
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x0
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x0
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x0
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x0
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x0
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x0
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x0
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x0
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x0
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x0
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x0
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x0
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x0
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x0
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x0
type PlayTeleportConfirm_758_0 struct {
	/* The ID given by the Player Position And Look packet */
	TeleportID VarInt // VarInt
}

var _ Packet = (*PlayTeleportConfirm_758_0)(nil)

func (p PlayTeleportConfirm_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.TeleportID)
}

func (p *PlayTeleportConfirm_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TeleportID, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x68
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x68
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x64
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x66
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x63
type PlayTeleportEntity_763_0 struct {
	EntityID VarInt // VarInt
	X Double // Double
	Y Double // Double
	Z Double // Double
	/* (Y Rot)New angle, not a delta. */
	Yaw Angle // Angle
	/* (X Rot)New angle, not a delta. */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayTeleportEntity_763_0)(nil)

func (p PlayTeleportEntity_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayTeleportEntity_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Y, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x30
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x30
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x30
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x30
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x2f
type PlayTeleportToEntity_763_0 struct {
	/* UUID of the player to teleport to (can also be an entity UUID). */
	TargetPlayer UUID // UUID
}

var _ Packet = (*PlayTeleportToEntity_763_0)(nil)

func (p PlayTeleportToEntity_763_0)Encode(b *PacketBuilder){
	b.UUID(p.TargetPlayer)
}

func (p *PlayTeleportToEntity_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.TargetPlayer, ok = r.UUID(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x59
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x59
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x58
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x58
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x4e
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x4e
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x4f
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x4e
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x4a
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x4a
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x47
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x47
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x46
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x44
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x44
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x44
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x44
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x3
type PlayTimeUpdate_758_0 struct {
	/* In ticks; not changed by server commands */
	WorldAge Long // Long
	/* The world (or region) time, in ticks. If negative the sun will stop moving at the Math.abs of the time */
	TimeOfDay Long // Long
}

var _ Packet = (*PlayTimeUpdate_758_0)(nil)

func (p PlayTimeUpdate_758_0)Encode(b *PacketBuilder){
	b.Long(p.WorldAge)
	b.Long(p.TimeOfDay)
}

func (p *PlayTimeUpdate_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WorldAge, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.TimeOfDay, ok = r.Long(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x1e
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x1e
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x1b
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x1c
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x1a
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x1d
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x1d
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x1d
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x1d
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x1c
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x1c
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x1e
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x1d
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x1f
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x1f
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x1d
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x1d
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x1d
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x1d
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x1d
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x1d
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x1d
type PlayUnloadChunk_763_0 struct {
	/* Block coordinate divided by 16, rounded down */
	ChunkX Int // Int
	/* Block coordinate divided by 16, rounded down */
	ChunkZ Int // Int
}

var _ Packet = (*PlayUnloadChunk_763_0)(nil)

func (p PlayUnloadChunk_763_0)Encode(b *PacketBuilder){
	b.Int(p.ChunkX)
	b.Int(p.ChunkZ)
}

func (p *PlayUnloadChunk_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkX, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.Int(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x39
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x39
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x39
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x39
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x35
type PlayUnlockRecipes_758_0 struct {
	/* 0: init, 1: add, 2: remove. */
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
	BlastFurnaceRecipeBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	BlastFurnaceRecipeBookFilterActive Bool // Boolean
	/* If true, then the smoker recipe book will be open when the player opens its inventory. */
	SmokerRecipeBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	SmokerRecipeBookFilterActive Bool // Boolean
	/* Number of elements in the following array. */
	ArraySize1 VarInt // VarInt
	RecipeIDs []Identifier // Array of Identifier
	/* Number of elements in the following array, only present if mode is 0 (init). */
	ArraySize2 VarInt // Optional VarInt
	/* Only present if mode is 0 (init) */
	RecipeIDs []Identifier // Optional Array of Identifier
}

var _ Packet = (*PlayUnlockRecipes_758_0)(nil)

func (p PlayUnlockRecipes_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.Action)
	b.Bool(p.CraftingRecipeBookOpen)
	b.Bool(p.CraftingRecipeBookFilterActive)
	b.Bool(p.SmeltingRecipeBookOpen)
	b.Bool(p.SmeltingRecipeBookFilterActive)
	b.Bool(p.BlastFurnaceRecipeBookOpen)
	b.Bool(p.BlastFurnaceRecipeBookFilterActive)
	b.Bool(p.SmokerRecipeBookOpen)
	b.Bool(p.SmokerRecipeBookFilterActive)
	b.VarInt(p.ArraySize1)
	TODO_Encode_Array(p.RecipeIDs)
	b.VarInt(p.ArraySize2)
	TODO_Encode_Array(p.RecipeIDs)
}

func (p *PlayUnlockRecipes_758_0)DecodeFrom(r *PacketReader)(err error){
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
	if p.BlastFurnaceRecipeBookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.BlastFurnaceRecipeBookFilterActive, ok = r.Bool(); !ok {
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

// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x37
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x36
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x34
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x34
type PlayUnlockRecipes_578_2 struct {
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
	/* Number of elements in the following array */
	ArraySize1 VarInt // VarInt
	RecipeIDs []Identifier // Array of Identifier
	/* Number of elements in the following array, only present if mode is 0 (init) */
	ArraySize2 VarInt // Optional VarInt
	RecipeIDs []Identifier // Optional Array of Identifier, only present if mode is 0 (init)
}

var _ Packet = (*PlayUnlockRecipes_578_2)(nil)

func (p PlayUnlockRecipes_578_2)Encode(b *PacketBuilder){
	b.VarInt(p.Action)
	b.Bool(p.CraftingRecipeBookOpen)
	b.Bool(p.CraftingRecipeBookFilterActive)
	b.Bool(p.SmeltingRecipeBookOpen)
	b.Bool(p.SmeltingRecipeBookFilterActive)
	b.VarInt(p.ArraySize1)
	TODO_Encode_Array(p.RecipeIDs)
	b.VarInt(p.ArraySize2)
	TODO_Encode_Array(p.RecipeIDs)
}

func (p *PlayUnlockRecipes_578_2)DecodeFrom(r *PacketReader)(err error){
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
	if p.ArraySize1, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.RecipeIDs)
	if p.ArraySize2, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.RecipeIDs)
}

// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x31
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x31
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x30
type PlayUnlockRecipes_340_3 struct {
	/* 0: init, 1: add, 2: remove */
	Action VarInt // VarInt
	/* If true, then the crafting book will be open when the player opens its inventory. */
	CraftingBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	FilteringCraftable Bool // Boolean
	/* Number of elements in the following array */
	ArraySize1 VarInt // VarInt
	RecipeIDs []VarInt // Array of VarInt
	/* Number of elements in the following array, only present if mode is 0 (init) */
	ArraySize2 VarInt // Optional VarInt
	RecipeIDs []VarInt // Optional Array of VarInt, only present if mode is 0 (init)
}

var _ Packet = (*PlayUnlockRecipes_340_3)(nil)

func (p PlayUnlockRecipes_340_3)Encode(b *PacketBuilder){
	b.VarInt(p.Action)
	b.Bool(p.CraftingBookOpen)
	b.Bool(p.FilteringCraftable)
	b.VarInt(p.ArraySize1)
	TODO_Encode_Array(p.RecipeIDs)
	b.VarInt(p.ArraySize2)
	TODO_Encode_Array(p.RecipeIDs)
}

func (p *PlayUnlockRecipes_340_3)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Action, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CraftingBookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.FilteringCraftable, ok = r.Bool(); !ok {
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

// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x9
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x9
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x9
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x9
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x9
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x9
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x9
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x9
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x9
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x9
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x35
type PlayUpdateBlockEntity_498_0 struct {
	Location Position // Position
	/* The type of update to perform, see below */
	Action UByte // Unsigned Byte
	/* If not present then it's a TAG_END (0) */
	NBTData NBT // Optional NBT Tag
}

var _ Packet = (*PlayUpdateBlockEntity_498_0)(nil)

func (p PlayUpdateBlockEntity_498_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.UByte(p.Action)
	b.Encode(p.NBTData)
}

func (p *PlayUpdateBlockEntity_498_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Action, ok = r.UByte(); !ok {
		return io.EOF
	}
	if err = p.NBTData.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x26
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x26
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x26
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x26
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x26
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x26
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x24
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x24
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x22
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x22
type PlayUpdateCommandBlock_758_0 struct {
	Location Position // Position
	Command String // String (32767)
	/* One of SEQUENCE (0), AUTO (1), or REDSTONE (2) */
	Mode VarInt // VarInt enum
	/* 0x01: Track Output (if false, the output of the previous command will not be stored within the command block); 0x02: Is conditional; 0x04: Automatic */
	Flags Byte // Byte
}

var _ Packet = (*PlayUpdateCommandBlock_758_0)(nil)

func (p PlayUpdateCommandBlock_758_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.String(p.Command)
	b.VarInt(p.Mode)
	b.Byte(p.Flags)
}

func (p *PlayUpdateCommandBlock_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Command, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x27
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x27
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x27
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x27
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x27
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x27
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x25
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x25
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x23
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x23
type PlayUpdateCommandBlockMinecart_758_0 struct {
	EntityID VarInt // VarInt
	Command String // String
	/* If false, the output of the previous command will not be stored within the command block. */
	TrackOutput Bool // Boolean
}

var _ Packet = (*PlayUpdateCommandBlockMinecart_758_0)(nil)

func (p PlayUpdateCommandBlockMinecart_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.String(p.Command)
	b.Bool(p.TrackOutput)
}

func (p *PlayUpdateCommandBlockMinecart_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Command, ok = r.String(); !ok {
		return io.EOF
	}
	if p.TrackOutput, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x2b
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x2b
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x27
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x28
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x26
type PlayUpdateEntityPosition_763_0 struct {
	EntityID VarInt // VarInt
	/* Change in X position as (currentX * 32 - prevX * 32) * 128. */
	DeltaX Short // Short
	/* Change in Y position as (currentY * 32 - prevY * 32) * 128. */
	DeltaY Short // Short
	/* Change in Z position as (currentZ * 32 - prevZ * 32) * 128. */
	DeltaZ Short // Short
	OnGround Bool // Boolean
}

var _ Packet = (*PlayUpdateEntityPosition_763_0)(nil)

func (p PlayUpdateEntityPosition_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.DeltaX)
	b.Short(p.DeltaY)
	b.Short(p.DeltaZ)
	b.Bool(p.OnGround)
}

func (p *PlayUpdateEntityPosition_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DeltaX, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaY, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaZ, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x2c
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x2c
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x28
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x29
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x27
type PlayUpdateEntityPositionAndRotation_763_0 struct {
	EntityID VarInt // VarInt
	/* Change in X position as (currentX * 32 - prevX * 32) * 128. */
	DeltaX Short // Short
	/* Change in Y position as (currentY * 32 - prevY * 32) * 128. */
	DeltaY Short // Short
	/* Change in Z position as (currentZ * 32 - prevZ * 32) * 128. */
	DeltaZ Short // Short
	/* New angle, not a delta. */
	Yaw Angle // Angle
	/* New angle, not a delta. */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayUpdateEntityPositionAndRotation_763_0)(nil)

func (p PlayUpdateEntityPositionAndRotation_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Short(p.DeltaX)
	b.Short(p.DeltaY)
	b.Short(p.DeltaZ)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayUpdateEntityPositionAndRotation_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.DeltaX, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaY, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.DeltaZ, ok = r.Short(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x2d
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x2d
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x29
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x2a
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x28
type PlayUpdateEntityRotation_763_0 struct {
	EntityID VarInt // VarInt
	/* New angle, not a delta. */
	Yaw Angle // Angle
	/* New angle, not a delta. */
	Pitch Angle // Angle
	OnGround Bool // Boolean
}

var _ Packet = (*PlayUpdateEntityRotation_763_0)(nil)

func (p PlayUpdateEntityRotation_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.Yaw)
	b.Encode(p.Pitch)
	b.Bool(p.OnGround)
}

func (p *PlayUpdateEntityRotation_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Yaw.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pitch.DecodeFrom(r); err != nil {
		return
	}
	if p.OnGround, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x52
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x52
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x52
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x52
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x49
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x49
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x49
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x48
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x44
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x44
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x41
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x41
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x40
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x3e
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x3e
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x3e
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x3e
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x6
type PlayUpdateHealth_758_0 struct {
	/* 0 or less = dead, 20 = full HP */
	Health Float // Float
	/* 0–20 */
	Food VarInt // VarInt
	/* Seems to vary from 0.0 to 5.0 in integer increments */
	FoodSaturation Float // Float
}

var _ Packet = (*PlayUpdateHealth_758_0)(nil)

func (p PlayUpdateHealth_758_0)Encode(b *PacketBuilder){
	b.Float(p.Health)
	b.VarInt(p.Food)
	b.Float(p.FoodSaturation)
}

func (p *PlayUpdateHealth_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Health, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Food, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.FoodSaturation, ok = r.Float(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x29
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x29
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x29
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x29
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x29
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x28
type PlayUpdateJigsawBlock_758_0 struct {
	/* Block entity location */
	Location Position // Position
	Name Identifier // Identifier
	Target Identifier // Identifier
	Pool Identifier // Identifier
	/* "Turns into" on the GUI, final_state in NBT */
	FinalState String // String
	/* rollable if the attached piece can be rotated, else aligned */
	JointType String // String
}

var _ Packet = (*PlayUpdateJigsawBlock_758_0)(nil)

func (p PlayUpdateJigsawBlock_758_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.Encode(p.Name)
	b.Encode(p.Target)
	b.Encode(p.Pool)
	b.String(p.FinalState)
	b.String(p.JointType)
}

func (p *PlayUpdateJigsawBlock_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Name.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Target.DecodeFrom(r); err != nil {
		return
	}
	if err = p.Pool.DecodeFrom(r); err != nil {
		return
	}
	if p.FinalState, ok = r.String(); !ok {
		return io.EOF
	}
	if p.JointType, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x27
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x27
type PlayUpdateJigsawBlock_578_1 struct {
	/* Block entity location */
	Location Position // Position
	AttachmentType Identifier // Identifier
	TargetPool Identifier // Identifier
	/* "Turns into" on the GUI, final_state in NBT */
	FinalState String // String
}

var _ Packet = (*PlayUpdateJigsawBlock_578_1)(nil)

func (p PlayUpdateJigsawBlock_578_1)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.Encode(p.AttachmentType)
	b.Encode(p.TargetPool)
	b.String(p.FinalState)
}

func (p *PlayUpdateJigsawBlock_578_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if err = p.AttachmentType.DecodeFrom(r); err != nil {
		return
	}
	if err = p.TargetPool.DecodeFrom(r); err != nil {
		return
	}
	if p.FinalState, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x58
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x58
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x54
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x56
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x53
type PlayUpdateObjectives_763_0 struct {
	/* A unique name for the objective. */
	ObjectiveName String // String (16)
	/* 0 to create the scoreboard. 1 to remove the scoreboard. 2 to update the display text. */
	Mode Byte // Byte
	/* Only if mode is 0 or 2. The text to be displayed for the score. */
	ObjectiveValue Object // Optional Chat
	/* Only if mode is 0 or 2. 0 = "integer", 1 = "hearts". */
	Type VarInt // Optional VarInt Enum
}

var _ Packet = (*PlayUpdateObjectives_763_0)(nil)

func (p PlayUpdateObjectives_763_0)Encode(b *PacketBuilder){
	b.String(p.ObjectiveName)
	b.Byte(p.Mode)
	b.JSON(p.ObjectiveValue)
	b.VarInt(p.Type)
}

func (p *PlayUpdateObjectives_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ObjectiveName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.Byte(); !ok {
		return io.EOF
	}
	if err = r.JSON(&p.ObjectiveValue); err != nil {
		return
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x3d
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x3d
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x39
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x3a
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x37
type PlayUpdateRecipeBook_763_0 struct {
	/* 0: init, 1: add, 2: remove. */
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
	BlastFurnaceRecipeBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	BlastFurnaceRecipeBookFilterActive Bool // Boolean
	/* If true, then the smoker recipe book will be open when the player opens its inventory. */
	SmokerRecipeBookOpen Bool // Boolean
	/* If true, then the filtering option is active when the players opens its inventory. */
	SmokerRecipeBookFilterActive Bool // Boolean
	/* Number of elements in the following array. */
	ArraySize1 VarInt // VarInt
	RecipeIDs []Identifier // Array of Identifier
	/* Number of elements in the following array, only present if mode is 0 (init). */
	ArraySize2 VarInt // Optional VarInt
	/* Only present if mode is 0 (init) */
	RecipeIDs []Identifier // Optional Array of Identifier
}

var _ Packet = (*PlayUpdateRecipeBook_763_0)(nil)

func (p PlayUpdateRecipeBook_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Action)
	b.Bool(p.CraftingRecipeBookOpen)
	b.Bool(p.CraftingRecipeBookFilterActive)
	b.Bool(p.SmeltingRecipeBookOpen)
	b.Bool(p.SmeltingRecipeBookFilterActive)
	b.Bool(p.BlastFurnaceRecipeBookOpen)
	b.Bool(p.BlastFurnaceRecipeBookFilterActive)
	b.Bool(p.SmokerRecipeBookOpen)
	b.Bool(p.SmokerRecipeBookFilterActive)
	b.VarInt(p.ArraySize1)
	TODO_Encode_Array(p.RecipeIDs)
	b.VarInt(p.ArraySize2)
	TODO_Encode_Array(p.RecipeIDs)
}

func (p *PlayUpdateRecipeBook_763_0)DecodeFrom(r *PacketReader)(err error){
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
	if p.BlastFurnaceRecipeBookOpen, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.BlastFurnaceRecipeBookFilterActive, ok = r.Bool(); !ok {
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

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x5b
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x5b
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x57
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x59
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x56
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x56
type PlayUpdateScore_763_0 struct {
	/* The entity whose score this is.  For players, this is their username; for other entities, it is their UUID. */
	EntityName String // String (40)
	/* 0 to create/update an item. 1 to remove an item. */
	Action VarInt // VarInt Enum
	/* The name of the objective the score belongs to. */
	ObjectiveName String // String (16)
	/* The score to be displayed next to the entry. Only sent when Action does not equal 1. */
	Value VarInt // Optional VarInt
}

var _ Packet = (*PlayUpdateScore_763_0)(nil)

func (p PlayUpdateScore_763_0)Encode(b *PacketBuilder){
	b.String(p.EntityName)
	b.VarInt(p.Action)
	b.String(p.ObjectiveName)
	b.VarInt(p.Value)
}

func (p *PlayUpdateScore_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Action, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ObjectiveName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Value, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x56
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x56
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x56
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x4d
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x4d
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x4d
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x4c
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x48
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x48
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x45
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x45
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x44
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x42
type PlayUpdateScore_757_1 struct {
	/* The entity whose score this is.  For players, this is their username; for other entities, it is their UUID. */
	EntityName String // String (40)
	/* 0 to create/update an item. 1 to remove an item. */
	Action Byte // Byte
	/* The name of the objective the score belongs to */
	ObjectiveName String // String (16)
	/* The score to be displayed next to the entry. Only sent when Action does not equal 1. */
	Value VarInt // Optional VarInt
}

var _ Packet = (*PlayUpdateScore_757_1)(nil)

func (p PlayUpdateScore_757_1)Encode(b *PacketBuilder){
	b.String(p.EntityName)
	b.Byte(p.Action)
	b.String(p.ObjectiveName)
	b.VarInt(p.Value)
}

func (p *PlayUpdateScore_757_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Action, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ObjectiveName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Value, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x42
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x42
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x42
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x3c
type PlayUpdateScore_315_2 struct {
	/* The name of the score to be updated or removed */
	ScoreName String // String
	/* 0 to create/update an item. 1 to remove an item. */
	Action Byte // Byte
	/* The name of the objective the score belongs to */
	ObjectiveName String // String
	/* The score to be displayed next to the entry. Only sent when Action does not equal 1. */
	Value VarInt // Optional VarInt
}

var _ Packet = (*PlayUpdateScore_315_2)(nil)

func (p PlayUpdateScore_315_2)Encode(b *PacketBuilder){
	b.String(p.ScoreName)
	b.Byte(p.Action)
	b.String(p.ObjectiveName)
	b.VarInt(p.Value)
}

func (p *PlayUpdateScore_315_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ScoreName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Action, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ObjectiveName, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Value, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x43
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x3f
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x40
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x3d
type PlayUpdateSectionBlocks_762_1 struct {
	/* Chunk section coordinate (encoded chunk x and z with each 22 bits, and section y with 20 bits, from left to right). */
	ChunkSectionPosition Long // Long
	/* Whether to ignore light updates caused by the contained changes. Always inverse the preceding Update Light packet's "Trust Edges" bool */
	SuppressLightUpdates Bool // Boolean
	/* Number of elements in the following array. */
	BlocksArraySize VarInt // VarInt
	/* Each entry is composed of the block state id, shifted left by 12, and the relative block position in the chunk section (4 bits for x, z, and y, from left to right). */
	Blocks []VarLong // Array of VarLong
}

var _ Packet = (*PlayUpdateSectionBlocks_762_1)(nil)

func (p PlayUpdateSectionBlocks_762_1)Encode(b *PacketBuilder){
	b.Long(p.ChunkSectionPosition)
	b.Bool(p.SuppressLightUpdates)
	b.VarInt(p.BlocksArraySize)
	TODO_Encode_Array(p.Blocks)
}

func (p *PlayUpdateSectionBlocks_762_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkSectionPosition, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.SuppressLightUpdates, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.BlocksArraySize, ok = r.VarInt(); !ok {
		return io.EOF
	}
	TODO_Decode_Array(p.Blocks)
}

// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x2e
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x2e
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x2e
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x2d
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x2b
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x2b
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x2b
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x2b
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x2b
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x2b
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x29
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x29
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x26
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x26
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x1c
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x1c
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x1c
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x19
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x19
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x19
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x19
type PlayUpdateSign_762_1 struct {
	/* Block Coordinates */
	Location Position // Position
	/* First line of text in the sign */
	Line1 String // String (384)
	/* Second line of text in the sign */
	Line2 String // String (384)
	/* Third line of text in the sign */
	Line3 String // String (384)
	/* Fourth line of text in the sign */
	Line4 String // String (384)
}

var _ Packet = (*PlayUpdateSign_762_1)(nil)

func (p PlayUpdateSign_762_1)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.String(p.Line1)
	b.String(p.Line2)
	b.String(p.Line3)
	b.String(p.Line4)
}

func (p *PlayUpdateSign_762_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Line1, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Line2, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Line3, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Line4, ok = r.String(); !ok {
		return io.EOF
	}
}

// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x33
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=server; ID=0x12
type PlayUpdateSign_47_2 struct {
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

var _ Packet = (*PlayUpdateSign_47_2)(nil)

func (p PlayUpdateSign_47_2)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.JSON(p.Line1)
	b.JSON(p.Line2)
	b.JSON(p.Line3)
	b.JSON(p.Line4)
}

func (p *PlayUpdateSign_47_2)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if err = r.JSON(&p.Line1); err != nil {
		return
	}
	if err = r.JSON(&p.Line2); err != nil {
		return
	}
	if err = r.JSON(&p.Line3); err != nil {
		return
	}
	if err = r.JSON(&p.Line4); err != nil {
		return
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x57
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x57
type PlayUpdateSimulationDistance_758_0 struct {
	/* The distance that the client will process specific things, such as entities. */
	SimulationDistance VarInt // VarInt
}

var _ Packet = (*PlayUpdateSimulationDistance_758_0)(nil)

func (p PlayUpdateSimulationDistance_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.SimulationDistance)
}

func (p *PlayUpdateSimulationDistance_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.SimulationDistance, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x2a
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x2a
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x2a
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x2a
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x2a
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x2a
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x28
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x28
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x25
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x25
type PlayUpdateStructureBlock_758_0 struct {
	/* Block entity location */
	Location Position // Position
	/* An additional action to perform beyond simply saving the given data; see below */
	Action VarInt // VarInt enum
	/* One of SAVE (0), LOAD (1), CORNER (2), DATA (3). */
	Mode VarInt // VarInt enum
	Name String // String
	/* Between -32 and 32 */
	OffsetX Byte // Byte
	/* Between -32 and 32 */
	OffsetY Byte // Byte
	/* Between -32 and 32 */
	OffsetZ Byte // Byte
	/* Between 0 and 32 */
	SizeX Byte // Byte
	/* Between 0 and 32 */
	SizeY Byte // Byte
	/* Between 0 and 32 */
	SizeZ Byte // Byte
	/* One of NONE (0), LEFT_RIGHT (1), FRONT_BACK (2). */
	Mirror VarInt // VarInt enum
	/* One of NONE (0), CLOCKWISE_90 (1), CLOCKWISE_180 (2), COUNTERCLOCKWISE_90 (3). */
	Rotation VarInt // VarInt enum
	Metadata String // String
	/* Between 0 and 1 */
	Integrity Float // Float
	Seed VarLong // VarLong
	/* 0x01: Ignore entities; 0x02: Show air; 0x04: Show bounding box */
	Flags Byte // Byte
}

var _ Packet = (*PlayUpdateStructureBlock_758_0)(nil)

func (p PlayUpdateStructureBlock_758_0)Encode(b *PacketBuilder){
	b.Encode(p.Location)
	b.VarInt(p.Action)
	b.VarInt(p.Mode)
	b.String(p.Name)
	b.Byte(p.OffsetX)
	b.Byte(p.OffsetY)
	b.Byte(p.OffsetZ)
	b.Byte(p.SizeX)
	b.Byte(p.SizeY)
	b.Byte(p.SizeZ)
	b.VarInt(p.Mirror)
	b.VarInt(p.Rotation)
	b.String(p.Metadata)
	b.Float(p.Integrity)
	b.VarLong(p.Seed)
	b.Byte(p.Flags)
}

func (p *PlayUpdateStructureBlock_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Action, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Mode, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Name, ok = r.String(); !ok {
		return io.EOF
	}
	if p.OffsetX, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.OffsetY, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.OffsetZ, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.SizeX, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.SizeY, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.SizeZ, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.Mirror, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Rotation, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Metadata, ok = r.String(); !ok {
		return io.EOF
	}
	if p.Integrity, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Seed, ok = r.VarLong(); !ok {
		return io.EOF
	}
	if p.Flags, ok = r.Byte(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x5e
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x5e
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x5a
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x5c
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x59
type PlayUpdateTime_763_0 struct {
	/* In ticks; not changed by server commands. */
	WorldAge Long // Long
	/* The world (or region) time, in ticks. If negative the sun will stop moving at the Math.abs of the time. */
	TimeOfDay Long // Long
}

var _ Packet = (*PlayUpdateTime_763_0)(nil)

func (p PlayUpdateTime_763_0)Encode(b *PacketBuilder){
	b.Long(p.WorldAge)
	b.Long(p.TimeOfDay)
}

func (p *PlayUpdateTime_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WorldAge, ok = r.Long(); !ok {
		return io.EOF
	}
	if p.TimeOfDay, ok = r.Long(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x4a
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x4a
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x4a
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x4a
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x41
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x41
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x42
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x41
type PlayUpdateViewDistance_758_0 struct {
	/* Render distance (2-32) */
	ViewDistance VarInt // VarInt
}

var _ Packet = (*PlayUpdateViewDistance_758_0)(nil)

func (p PlayUpdateViewDistance_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.ViewDistance)
}

func (p *PlayUpdateViewDistance_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ViewDistance, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x49
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x49
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x49
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x49
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x40
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x40
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x41
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x40
type PlayUpdateViewPosition_758_0 struct {
	/* Chunk X coordinate of the player's position */
	ChunkX VarInt // VarInt
	/* Chunk Z coordinate of the player's position */
	ChunkZ VarInt // VarInt
}

var _ Packet = (*PlayUpdateViewPosition_758_0)(nil)

func (p PlayUpdateViewPosition_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.ChunkX)
	b.VarInt(p.ChunkZ)
}

func (p *PlayUpdateViewPosition_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.ChunkX, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.ChunkZ, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x33
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x33
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x30
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x30
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x2f
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x2f
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x2f
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x2f
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x2f
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0xa
type PlayUseBed_404_0 struct {
	/* Sleeping player's EID */
	EntityID VarInt // VarInt
	/* Block location of the head part of the bed */
	Location Position // Position
}

var _ Packet = (*PlayUseBed_404_0)(nil)

func (p PlayUseBed_404_0)Encode(b *PacketBuilder){
	b.VarInt(p.EntityID)
	b.Encode(p.Location)
}

func (p *PlayUseBed_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.EntityID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0xd
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0xd
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0xa
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0xa
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0xb
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0xa
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0xa
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0xa
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0xa
type PlayUseEntity_404_0 struct {
	Target VarInt // VarInt
	/* 0: interact, 1: attack, 2: interact at */
	Type VarInt // VarInt Enum
	/* Only if Type is interact at */
	TargetX Float // Optional Float
	/* Only if Type is interact at */
	TargetY Float // Optional Float
	/* Only if Type is interact at */
	TargetZ Float // Optional Float
	/* Only if Type is interact or interact at; 0: main hand, 1: off hand */
	Hand VarInt // Optional VarInt Enum
}

var _ Packet = (*PlayUseEntity_404_0)(nil)

func (p PlayUseEntity_404_0)Encode(b *PacketBuilder){
	b.VarInt(p.Target)
	b.VarInt(p.Type)
	b.Float(p.TargetX)
	b.Float(p.TargetY)
	b.Float(p.TargetZ)
	b.VarInt(p.Hand)
}

func (p *PlayUseEntity_404_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Target, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Type, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.TargetX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TargetY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.TargetZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x32
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x32
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x32
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x32
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x31
type PlayUseItem_763_0 struct {
	/* Hand used for the animation. 0: main hand, 1: off hand. */
	Hand VarInt // VarInt Enum
	Sequence VarInt // VarInt
}

var _ Packet = (*PlayUseItem_763_0)(nil)

func (p PlayUseItem_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
	b.VarInt(p.Sequence)
}

func (p *PlayUseItem_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Sequence, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x2f
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x2f
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x2f
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x2f
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x2f
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x2f
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x2d
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x2d
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x2a
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x2a
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x20
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x20
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x20
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x1d
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x1d
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x1d
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x1d
type PlayUseItem_758_1 struct {
	/* Hand used for the animation. 0: main hand, 1: off hand. */
	Hand VarInt // VarInt Enum
}

var _ Packet = (*PlayUseItem_758_1)(nil)

func (p PlayUseItem_758_1)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
}

func (p *PlayUseItem_758_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=server; ID=0x31
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=server; ID=0x31
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=server; ID=0x31
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=server; ID=0x31
// Protocol=759; ProtocolName=1.19; State=play; Bound=server; ID=0x30
type PlayUseItemOn_763_0 struct {
	/* The hand from which the block is placed; 0: main hand, 1: off hand. */
	Hand VarInt // VarInt Enum
	/* Block position. */
	Location Position // Position
	/* The face on which the block is placed (as documented at Player Action). */
	Face VarInt // VarInt Enum
	/* The position of the crosshair on the block, from 0 to 1 increasing from west to east. */
	CursorPositionX Float // Float
	/* The position of the crosshair on the block, from 0 to 1 increasing from bottom to top. */
	CursorPositionY Float // Float
	/* The position of the crosshair on the block, from 0 to 1 increasing from north to south. */
	CursorPositionZ Float // Float
	/* True when the player's head is inside of a block. */
	InsideBlock Bool // Boolean
	Sequence VarInt // VarInt
}

var _ Packet = (*PlayUseItemOn_763_0)(nil)

func (p PlayUseItemOn_763_0)Encode(b *PacketBuilder){
	b.VarInt(p.Hand)
	b.Encode(p.Location)
	b.VarInt(p.Face)
	b.Float(p.CursorPositionX)
	b.Float(p.CursorPositionY)
	b.Float(p.CursorPositionZ)
	b.Bool(p.InsideBlock)
	b.VarInt(p.Sequence)
}

func (p *PlayUseItemOn_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Hand, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Face, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.CursorPositionX, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.CursorPositionY, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.CursorPositionZ, ok = r.Float(); !ok {
		return io.EOF
	}
	if p.InsideBlock, ok = r.Bool(); !ok {
		return io.EOF
	}
	if p.Sequence, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x2c
// Protocol=758; ProtocolName=1.18.2; State=play; Bound=server; ID=0x15
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x2c
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=server; ID=0x15
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x2c
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=server; ID=0x15
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x2c
// Protocol=755; ProtocolName=1.17; State=play; Bound=server; ID=0x15
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x2b
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x16
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x2b
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x16
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x2d
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x15
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x2c
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x15
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x2b
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=server; ID=0x13
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x2b
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=server; ID=0x13
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x29
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=server; ID=0x10
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x29
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=server; ID=0x10
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x29
// Protocol=335; ProtocolName=1.12; State=play; Bound=server; ID=0x11
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x29
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=server; ID=0x10
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x29
// Protocol=315; ProtocolName=1.11; State=play; Bound=server; ID=0x10
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x29
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=server; ID=0x10
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x29
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=server; ID=0x10
type PlayVehicleMove_758_0 struct {
	/* Absolute position (X coordinate) */
	X Double // Double
	/* Absolute position (Y coordinate) */
	Y Double // Double
	/* Absolute position (Z coordinate) */
	Z Double // Double
	/* Absolute rotation on the vertical axis, in degrees */
	Yaw Float // Float
	/* Absolute rotation on the horizontal axis, in degrees */
	Pitch Float // Float
}

var _ Packet = (*PlayVehicleMove_758_0)(nil)

func (p PlayVehicleMove_758_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Y)
	b.Double(p.Z)
	b.Float(p.Yaw)
	b.Float(p.Pitch)
}

func (p *PlayVehicleMove_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
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
}

// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x11
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=server; ID=0x7
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x11
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=server; ID=0x7
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x13
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=server; ID=0x7
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x12
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=server; ID=0x7
type PlayWindowConfirmation_754_0 struct {
	/* The ID of the window that the action occurred in */
	WindowID Byte // Byte
	/* Every action that is to be accepted has a unique number. This number is an incrementing integer (starting at 0) with separate counts for each window ID. */
	ActionNumber Short // Short
	/* Whether the action was accepted */
	Accepted Bool // Boolean
}

var _ Packet = (*PlayWindowConfirmation_754_0)(nil)

func (p PlayWindowConfirmation_754_0)Encode(b *PacketBuilder){
	b.Byte(p.WindowID)
	b.Short(p.ActionNumber)
	b.Bool(p.Accepted)
}

func (p *PlayWindowConfirmation_754_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.Byte(); !ok {
		return io.EOF
	}
	if p.ActionNumber, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.Accepted, ok = r.Bool(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x14
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x14
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x14
type PlayWindowItems_758_0 struct {
	/* The ID of window which items are being sent for. 0 for player inventory. */
	WindowID UByte // Unsigned Byte
	/* See State ID */
	StateID VarInt // VarInt
	/* Number of elements in the following array. */
	Count VarInt // VarInt
	SlotData []*Slot // Array of Slot
	/* Item held by player. */
	CarriedItem *Slot // Slot
}

var _ Packet = (*PlayWindowItems_758_0)(nil)

func (p PlayWindowItems_758_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.VarInt(p.StateID)
	p.Count = len(p.SlotData)
	b.VarInt(p.Count)
	for _, v := range p.SlotData {
		b.Encode(v)
	}
	b.Encode(p.CarriedItem)
}

func (p *PlayWindowItems_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.StateID, ok = r.VarInt(); !ok {
		return io.EOF
	}
	if p.Count, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.SlotData = make([]*Slot, p.Count)
	for i, _ := range p.SlotData {
		if err = p.SlotData[i].DecodeFrom(r); err != nil {
			return
		}
	}
	if err = p.CarriedItem.DecodeFrom(r); err != nil {
		return
	}
}

// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x14
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x13
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x13
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x15
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x14
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x15
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x15
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x14
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x14
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x14
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x14
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x14
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x14
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x14
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x30
type PlayWindowItems_755_1 struct {
	/* The ID of window which items are being sent for. 0 for player inventory. */
	WindowID UByte // Unsigned Byte
	/* Number of elements in the following array */
	Count Short // Short
	SlotData []*Slot // Array of Slot
}

var _ Packet = (*PlayWindowItems_755_1)(nil)

func (p PlayWindowItems_755_1)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	p.Count = len(p.SlotData)
	b.Short(p.Count)
	for _, v := range p.SlotData {
		b.Encode(v)
	}
}

func (p *PlayWindowItems_755_1)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Count, ok = r.Short(); !ok {
		return io.EOF
	}
	p.SlotData = make([]*Slot, p.Count)
	for i, _ := range p.SlotData {
		if err = p.SlotData[i].DecodeFrom(r); err != nil {
			return
		}
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x15
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x15
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x15
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x15
// Protocol=754; ProtocolName=1.16.5; State=play; Bound=client; ID=0x14
// Protocol=753; ProtocolName=1.16.3; State=play; Bound=client; ID=0x14
// Protocol=578; ProtocolName=1.15.2; State=play; Bound=client; ID=0x16
// Protocol=498; ProtocolName=1.14.4; State=play; Bound=client; ID=0x15
// Protocol=404; ProtocolName=1.13.2; State=play; Bound=client; ID=0x16
// Protocol=401; ProtocolName=1.13.1; State=play; Bound=client; ID=0x16
// Protocol=340; ProtocolName=1.12.2; State=play; Bound=client; ID=0x15
// Protocol=338; ProtocolName=1.12.1; State=play; Bound=client; ID=0x15
// Protocol=335; ProtocolName=1.12; State=play; Bound=client; ID=0x15
// Protocol=316; ProtocolName=1.11.2; State=play; Bound=client; ID=0x15
// Protocol=315; ProtocolName=1.11; State=play; Bound=client; ID=0x15
// Protocol=210; ProtocolName=1.10.2; State=play; Bound=client; ID=0x15
// Protocol=110; ProtocolName=1.9.4; State=play; Bound=client; ID=0x15
// Protocol=47; ProtocolName=1.8.9; State=play; Bound=client; ID=0x31
type PlayWindowProperty_758_0 struct {
	WindowID UByte // Unsigned Byte
	/* The property to be updated, see below */
	Property Short // Short
	/* The new value for the property, see below */
	Value Short // Short
}

var _ Packet = (*PlayWindowProperty_758_0)(nil)

func (p PlayWindowProperty_758_0)Encode(b *PacketBuilder){
	b.UByte(p.WindowID)
	b.Short(p.Property)
	b.Short(p.Value)
}

func (p *PlayWindowProperty_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WindowID, ok = r.UByte(); !ok {
		return io.EOF
	}
	if p.Property, ok = r.Short(); !ok {
		return io.EOF
	}
	if p.Value, ok = r.Short(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x42
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x42
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x42
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x42
type PlayWorldBorderCenter_758_0 struct {
	X Double // Double
	Z Double // Double
}

var _ Packet = (*PlayWorldBorderCenter_758_0)(nil)

func (p PlayWorldBorderCenter_758_0)Encode(b *PacketBuilder){
	b.Double(p.X)
	b.Double(p.Z)
}

func (p *PlayWorldBorderCenter_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.X, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Z, ok = r.Double(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x43
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x43
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x43
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x43
type PlayWorldBorderLerpSize_758_0 struct {
	/* Current length of a single side of the world border, in meters. */
	OldDiameter Double // Double
	/* Target length of a single side of the world border, in meters. */
	NewDiameter Double // Double
	/* Number of real-time milliseconds until New Diameter is reached. It appears that Notchian server does not sync world border speed to game ticks, so it gets out of sync with server lag. If the world border is not moving, this is set to 0. */
	Speed VarLong // VarLong
}

var _ Packet = (*PlayWorldBorderLerpSize_758_0)(nil)

func (p PlayWorldBorderLerpSize_758_0)Encode(b *PacketBuilder){
	b.Double(p.OldDiameter)
	b.Double(p.NewDiameter)
	b.VarLong(p.Speed)
}

func (p *PlayWorldBorderLerpSize_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.OldDiameter, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.NewDiameter, ok = r.Double(); !ok {
		return io.EOF
	}
	if p.Speed, ok = r.VarLong(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x44
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x44
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x44
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x44
type PlayWorldBorderSize_758_0 struct {
	/* Length of a single side of the world border, in meters. */
	Diameter Double // Double
}

var _ Packet = (*PlayWorldBorderSize_758_0)(nil)

func (p PlayWorldBorderSize_758_0)Encode(b *PacketBuilder){
	b.Double(p.Diameter)
}

func (p *PlayWorldBorderSize_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Diameter, ok = r.Double(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x45
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x45
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x45
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x45
type PlayWorldBorderWarningDelay_758_0 struct {
	/* In seconds as set by /worldborder warning time. */
	WarningTime VarInt // VarInt
}

var _ Packet = (*PlayWorldBorderWarningDelay_758_0)(nil)

func (p PlayWorldBorderWarningDelay_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.WarningTime)
}

func (p *PlayWorldBorderWarningDelay_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WarningTime, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=758; ProtocolName=1.18.2; State=play; Bound=client; ID=0x46
// Protocol=757; ProtocolName=1.18.1; State=play; Bound=client; ID=0x46
// Protocol=756; ProtocolName=1.17.1; State=play; Bound=client; ID=0x46
// Protocol=755; ProtocolName=1.17; State=play; Bound=client; ID=0x46
type PlayWorldBorderWarningReach_758_0 struct {
	/* In meters. */
	WarningBlocks VarInt // VarInt
}

var _ Packet = (*PlayWorldBorderWarningReach_758_0)(nil)

func (p PlayWorldBorderWarningReach_758_0)Encode(b *PacketBuilder){
	b.VarInt(p.WarningBlocks)
}

func (p *PlayWorldBorderWarningReach_758_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.WarningBlocks, ok = r.VarInt(); !ok {
		return io.EOF
	}
}

// Protocol=763; ProtocolName=1.20; State=play; Bound=client; ID=0x25
// Protocol=762; ProtocolName=1.19.4; State=play; Bound=client; ID=0x25
// Protocol=761; ProtocolName=1.19.3; State=play; Bound=client; ID=0x21
// Protocol=760; ProtocolName=1.19.2; State=play; Bound=client; ID=0x22
// Protocol=759; ProtocolName=1.19; State=play; Bound=client; ID=0x20
type PlayWorldEvent_763_0 struct {
	/* The event, see below. */
	Event Int // Int
	/* The location of the event. */
	Location Position // Position
	/* Extra data for certain events, see below. */
	Data Int // Int
	/* See above. */
	DisableRelativeVolume Bool // Boolean
}

var _ Packet = (*PlayWorldEvent_763_0)(nil)

func (p PlayWorldEvent_763_0)Encode(b *PacketBuilder){
	b.Int(p.Event)
	b.Encode(p.Location)
	b.Int(p.Data)
	b.Bool(p.DisableRelativeVolume)
}

func (p *PlayWorldEvent_763_0)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Event, ok = r.Int(); !ok {
		return io.EOF
	}
	if err = p.Location.DecodeFrom(r); err != nil {
		return
	}
	if p.Data, ok = r.Int(); !ok {
		return io.EOF
	}
	if p.DisableRelativeVolume, ok = r.Bool(); !ok {
		return io.EOF
	}
}
