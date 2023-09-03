
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

type StatusResponsePkt struct {
	Payload Object
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

type PingResponsePkt PingRequestPkt

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
