
package liter

import (
	"strings"
	"fmt"
	"io"
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


type LoginStartPacket struct {
	Name string

	// if protocol >= 1.19 && protocol <= 1.19.2
	HasSign bool
	Timestamp Long
	PublicKey ByteArray
	Sign ByteArray

	// if protocol >= 1.19.1
	Id Optional[UUID]
}

var _ Packet = (*LoginStartPacket)(nil)

func (p LoginStartPacket)String()(string){
	if p.Id.Ok {
		return fmt.Sprintf("<LoginStartPacket name=%s uuid=%v>", p.Name, p.Id.V)
	}
	return fmt.Sprintf("<LoginStartPacket name=%s>", p.Name)
}

func (p LoginStartPacket)Encode(b *PacketBuilder){
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

func (p *LoginStartPacket)DecodeFrom(r *PacketReader)(err error){
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
