
package liter

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"io"
	"strings"
)

const (
	NextPingState  = 1
	NextLoginState = 2
)

// It's on purpose that `HandshakePkt` didn't implement `Packet` interface
// To prevent accidently call with `Conn.SendPkt` or `Conn.RecvPkt`
// If you want to handle it, please use `Conn.SendHandshakePkt` or `Conn.RecvHandshakePkt` instead
type HandshakePkt struct {
	Protocol  int
	Addr      string
	Addition  string
	Port      uint16
	NextState VarInt
}

func (p HandshakePkt)String()(string){
	if len(p.Addition) == 0 {
		return fmt.Sprintf("<HandshakePkt protocol=%d addr=%s port=%d nextState=%d>", p.Protocol, p.Addr, p.Port, p.NextState)
	}
	return fmt.Sprintf("<HandshakePkt protocol=%d addr=%s + %q port=%d nextState=%d>", p.Protocol, p.Addr, p.Addition, p.Port, p.NextState)
}

func (p HandshakePkt)EncodeTo(b *PacketBuilder){
	b.
		VarInt((VarInt)(p.Protocol)).
		String(p.Addr + p.Addition).
		Short((Short)(p.Port)).
		VarInt(p.NextState)
}

func (p *HandshakePkt)Decode(r *PacketReader)(err error){
	var ok bool
	var v VarInt
	if v, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Protocol = (int)(v)
	if p.Addr, ok = r.String(); !ok {
		return io.EOF
	}
	if i := strings.IndexByte(p.Addr, '\x00'); i != -1 {
		p.Addition = p.Addr[i:]
		p.Addr = p.Addr[:i]
	}
	if p.Port, ok = r.Uint16(); !ok {
		return io.EOF
	}
	if p.NextState, ok = r.VarInt(); !ok {
		return io.EOF
	}
	return
}

////////////////////////////
//// BEGIN PING PACKETS ////
////////////////////////////

type StatusRequestPkt struct{}

var _ Packet = (*StatusRequestPkt)(nil)

func (p StatusRequestPkt)String()(string){ return "<StatusRequestPkt>" }
func (p StatusRequestPkt)Encode(b *PacketBuilder){}
func (p *StatusRequestPkt)DecodeFrom(r *PacketReader)(err error){ return }

type PlayerInfo struct {
	Name string `json:"name"`
	Id   UUID   `json:"id"`
}

type (
	StatusResponsePayload struct {
		Version ProtocolVersion `json:"version"`
		Players PlayerStatus `json:"players"`
		Description *Chat `json:"description"`
	}
	ProtocolVersion struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	}
	PlayerStatus struct {
		Max    int `json:"max"`
		Online int `json:"online"`
		Sample []PlayerInfo `json:"sample,omitempty"`
	}
)

type StatusResponsePkt struct {
	Payload StatusResponsePayload
}

var _ Packet = (*StatusResponsePkt)(nil)

func (p StatusResponsePkt)String()(string){
	return "<StatusResponsePkt>"
}

func (p StatusResponsePkt)Encode(b *PacketBuilder){
	b.
		JSON(p.Payload)
}

func (p *StatusResponsePkt)DecodeFrom(r *PacketReader)(err error){
	if err = r.JSON(&p.Payload); err != nil {
		return
	}
	return
}

type PingRequestPkt struct {
	Payload int64
}

var _ Packet = (*PingRequestPkt)(nil)

func (p PingRequestPkt)String()(string){
	return fmt.Sprintf("<PingRequestPkt payload=%d>", p.Payload)
}

func (p PingRequestPkt)Encode(b *PacketBuilder){
	b.
		Long(p.Payload)
}

func (p *PingRequestPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Payload, ok = r.Long(); !ok {
		return io.EOF
	}
	return
}

type PingResponsePkt struct {
	Payload int64
}

var _ Packet = (*PingResponsePkt)(nil)

func (p PingResponsePkt)String()(string){
	return fmt.Sprintf("<PingResponsePkt payload=%d>", p.Payload)
}

func (p PingResponsePkt)Encode(b *PacketBuilder){
	b.
		Long(p.Payload)
}

func (p *PingResponsePkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.Payload, ok = r.Long(); !ok {
		return io.EOF
	}
	return
}

/////////////////////////////
//// BEGIN LOGIN PACKETS ////
/////////////////////////////

// See: https://wiki.vg/Protocol#Login
// See: https://wiki.vg/Protocol_Encryption

type DisconnectPkt struct {
	Reason *Chat
}

var _ Packet = (*DisconnectPkt)(nil)

func (p DisconnectPkt)Encode(b *PacketBuilder){
	b.Chat(p.Reason)
}

func (p *DisconnectPkt)DecodeFrom(r *PacketReader)(err error){
	if err = r.JSON(&p.Reason); err != nil {
		return
	}
	return
}

type LoginStartPkt struct {
	Name string

	// if protocol >= 1.19 && protocol <= 1.19.2
	HasSign bool
	Timestamp Long
	PublicKey ByteArray
	Sign ByteArray

	// if protocol >= 1.19.1
	Id Optional[UUID]
}

var _ Packet = (*LoginStartPkt)(nil)

func (p LoginStartPkt)String()(string){
	if p.Id.Ok {
		return fmt.Sprintf("<LoginStartPkt name=%s uuid=%v>", p.Name, p.Id.V)
	}
	return fmt.Sprintf("<LoginStartPkt name=%s>", p.Name)
}

func (p LoginStartPkt)Encode(b *PacketBuilder){
	protocol := b.Protocol()
	b.String(p.Name)
	if protocol >= V1_19 {
		if protocol <= V1_19_2 {
			b.Bool(p.HasSign)
			if p.HasSign {
				b.
					Long(p.Timestamp).
					VarInt((VarInt)(len(p.PublicKey))).
					ByteArray(p.PublicKey).
					VarInt((VarInt)(len(p.Sign))).
					ByteArray(p.Sign)
			}
		}
		if protocol >= V1_19_2 {
			b.Bool(p.Id.Ok)
			if p.Id.Ok {
				b.UUID(p.Id.V)
			}
		}
	}
}

func (p *LoginStartPkt)DecodeFrom(r *PacketReader)(err error){
	protocol := r.Protocol()
	var ok bool
	if p.Name, ok = r.String(); !ok {
		return io.EOF
	}
	if protocol >= V1_19 {
		if protocol <= V1_19_2 {
			if p.HasSign, ok = r.Bool(); !ok {
				return io.EOF
			}
			if p.HasSign {
				if p.Timestamp, ok = r.Long(); !ok {
					return io.EOF
				}
				var n VarInt
				if n, ok = r.VarInt(); !ok {
					return io.EOF
				}
				p.PublicKey = make(ByteArray, n)
				if ok = r.ByteArray(p.PublicKey); !ok {
					return io.EOF
				}
				if n, ok = r.VarInt(); !ok {
					return io.EOF
				}
				p.Sign = make(ByteArray, n)
				if ok = r.ByteArray(p.Sign); !ok {
					return io.EOF
				}
			}
		}
		if protocol >= V1_19_2 {
			if p.Id.Ok, ok = r.Bool(); !ok {
				return io.EOF
			}
			if p.Id.Ok {
				if p.Id.V, ok = r.UUID(); !ok {
					return io.EOF
				}
			}
		}
	}
	return
}

type LoginEncryptionRequestPkt struct {
	ServerID string // Appears to be empty.
	PublicKey *rsa.PublicKey // encoded in ASN.1 DER format
	/* A sequence of random bytes generated by the server.
	 * Length of Verify Token is always 4 for Notchian servers. */
	VerifyToken ByteArray
}

var _ Packet = (*LoginEncryptionRequestPkt)(nil)

func (p LoginEncryptionRequestPkt)Encode(b *PacketBuilder){
	b.String(p.ServerID)
	pubKey, _ := x509.MarshalPKIXPublicKey(p.PublicKey)
	b.VarInt((VarInt)(len(pubKey)))
	b.ByteArray(pubKey)
	b.VarInt((VarInt)(len(p.VerifyToken)))
	b.ByteArray(p.VerifyToken)
}

func (p *LoginEncryptionRequestPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	var n VarInt
	if p.ServerID, ok = r.String(); !ok {
		return io.EOF
	}
	if n, ok = r.VarInt(); !ok {
		return io.EOF
	}
	pubKey := make(ByteArray, n)
	if ok = r.ByteArray(pubKey); !ok {
		return io.EOF
	}
	var key any
	if key, err = x509.ParsePKIXPublicKey(pubKey); err != nil {
		return
	}
	if p.PublicKey, ok = key.(*rsa.PublicKey); !ok {
		return fmt.Errorf("Unexpected public key type %T when parsing login encryption request", key)
	}
	if n, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.VerifyToken = make(ByteArray, n)
	if ok = r.ByteArray(p.VerifyToken); !ok {
		return io.EOF
	}
	return
}

type LoginEncryptionResponsePkt struct {
	SharedSecret ByteArray // Shared Secret value, encrypted with the server's public key.
	VerifyToken ByteArray // Verify Token value, encrypted with the same public key as the shared secret.
}

var _ Packet = (*LoginEncryptionResponsePkt)(nil)

func (p LoginEncryptionResponsePkt)Encode(b *PacketBuilder){
	b.VarInt((VarInt)(len(p.SharedSecret)))
	b.ByteArray(p.SharedSecret)
	b.VarInt((VarInt)(len(p.VerifyToken)))
	b.ByteArray(p.VerifyToken)
}

func (p *LoginEncryptionResponsePkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	var n VarInt
	if n, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.SharedSecret = make(ByteArray, n)
	if ok = r.ByteArray(p.SharedSecret); !ok {
		return io.EOF
	}
	if n, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.VerifyToken = make(ByteArray, n)
	if ok = r.ByteArray(p.VerifyToken); !ok {
		return io.EOF
	}
	return
}

type LoginSuccessPkt struct {
	UUID UUID
	Username string
	Properties []*Property
}

var _ Packet = (*LoginSuccessPkt)(nil)

func (p *LoginSuccessPkt)Encode(b *PacketBuilder){
	b.UUID(p.UUID)
	b.String(p.Username)
	b.VarInt((VarInt)(len(p.Properties)))
	for _, v := range p.Properties {
		v.Encode(b)
	}
}

func (p *LoginSuccessPkt)DecodeFrom(r *PacketReader)(err error){
	var ok bool
	if p.UUID, ok = r.UUID(); !ok {
		return io.EOF
	}
	if p.Username, ok = r.String(); !ok {
		return io.EOF
	}
	var n VarInt
	if n, ok = r.VarInt(); !ok {
		return io.EOF
	}
	p.Properties = make([]*Property, n)
	for i, _ := range p.Properties {
		v := new(Property)
		if err = v.DecodeFrom(r); err != nil {
			return
		}
		p.Properties[i] = v
	}
	return
}

type LoginAcknowledgedPkt struct{}

var _ Packet = (*LoginAcknowledgedPkt)(nil)

func (p LoginAcknowledgedPkt)String()(string){ return "<LoginAcknowledgedPkt>" }
func (p LoginAcknowledgedPkt)Encode(b *PacketBuilder){}
func (p *LoginAcknowledgedPkt)DecodeFrom(r *PacketReader)(err error){ return }
