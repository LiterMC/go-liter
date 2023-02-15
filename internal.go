
package liter

import (
	"strings"
	"fmt"
	"io"
)

const (
	NextPingState  int32 = 1
	NextLoginState int32 = 2
)

type HandshakeP struct{
	Protocol  int32
	Addr      string
	Addition  string
	Port      uint16
	NextState int32
}

var _ Encodable = HandshakeP{}

func (p HandshakeP)Encode(b *PacketBuilder)(err error){
	b.
		VarInt(p.Protocol).
		String(p.Addr + p.Addition).
		Short((Short)(p.Port)).
		VarInt(p.NextState)
	return
}

func (p HandshakeP)String()(string){
	if len(p.Addition) == 0 {
		return fmt.Sprintf("<HandshakeP protocol=%d addr=%s port=%d nextState=%d>", p.Protocol, p.Addr, p.Port, p.NextState)
	}
	return fmt.Sprintf("<HandshakeP protocol=%d addr=%s + %q port=%d nextState=%d>", p.Protocol, p.Addr, p.Addition, p.Port, p.NextState)
}

func ReadHandshakeP(r *PacketReader)(p HandshakeP, err error){
	var ok bool
	if p.Protocol, ok = r.VarInt(); !ok {
		err = io.EOF
		return
	}
	if p.Addr, ok = r.String(); !ok {
		err = io.EOF
		return
	}
	if i := strings.IndexByte(p.Addr, '\x00'); i != -1 {
		p.Addition = p.Addr[i:]
		p.Addr = p.Addr[:i]
	}
	if p.Port, ok = r.Uint16(); !ok {
		err = io.EOF
		return
	}
	if p.NextState, ok = r.VarInt(); !ok {
		err = io.EOF
		return
	}
	return
}

type PingRequestP struct{
}

func (p PingRequestP)Encode(b *PacketBuilder)(err error){
	return
}

func (p PingRequestP)String()(string){
	return fmt.Sprintf("<PingRequestP>")
}

func ReadPingRequestP(r *PacketReader)(err error){
	// TODO?
	return
}
