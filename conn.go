
package liter

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

const DefaultPort = 25565

var (
	ErrOldHandshake = errors.New("Old client(<=1.6) handshake received") // The handshake packet <= 1.6 which first byte is 0xFE
)

type PktIdAssertError struct {
	Require int32
	Got     int32
}

var _ error = (*PktIdAssertError)(nil)

func (e *PktIdAssertError)Error()(string){
	return fmt.Sprintf("PktIdAssertError: require=%d; got=%d", e.Require, e.Got)
}


type Conn struct{
	protocol int
	conn net.Conn
}

func WrapConn(c net.Conn)(*Conn){
	return &Conn{
		protocol: V_UNSET,
		conn: c,
	}
}

func ResloveAddrWithContext(ctx context.Context, target string)(addr *net.TCPAddr, err error){
	addr = &net.TCPAddr{
		Port: DefaultPort,
	}
	if strings.IndexByte(target, ':') == -1 {
		var srvs []*net.SRV
		if _, srvs, err = net.DefaultResolver.LookupSRV(ctx, "minecraft", "tcp", target); err == nil && len(srvs) > 0 {
			srv := srvs[0]
			target, addr.Port = srv.Target, (int)(srv.Port)
		}
		var ips []net.IP
		if ips, err = net.LookupIP(target); err != nil {
			return nil, err
		}
		addr.IP = ips[0]
	}else{
		if addr.IP = net.ParseIP(target); addr.IP == nil {
			if addr, err = net.ResolveTCPAddr("tcp", target); err != nil {
				return nil, err
			}
		}
	}
	return
}

func ResloveAddr(target string)(addr *net.TCPAddr, err error){
	return ResloveAddrWithContext(context.Background(), target)
}

func DialWithContext(ctx context.Context, target string)(c *Conn, err error){
	var c0 *net.TCPConn
	var addr *net.TCPAddr
	if addr, err = ResloveAddrWithContext(ctx, target); err != nil {
		return
	}
	if c0, err = net.DialTCP("tcp", nil, addr); err != nil {
		return
	}
	return WrapConn(c0), nil
}

func Dial(target string)(c *Conn, err error){
	return DialWithContext(context.Background(), target)
}

func (c *Conn)Protocol()(int){
	return c.protocol
}

func (c *Conn)RawConn()(net.Conn){
	return c.conn
}

func (c *Conn)LocalAddr()(net.Addr){
	return c.conn.LocalAddr()
}

func (c *Conn)RemoteAddr()(net.Addr){
	return c.conn.RemoteAddr()
}

func (c *Conn)Close()(err error){
	return c.conn.Close()
}

func (c *Conn)Send(p *PacketBuilder)(err error){
	if _, err = p.WriteTo(c.conn); err != nil {
		return
	}
	return
}

func (c *Conn)SendPkt(id int32, values ...any)(err error){
	p := NewPacket(c.protocol, (VarInt)(id))
	for _, v := range values {
		p.Encode(v)
	}
	return c.Send(p)
}

func (c *Conn)SendHandshakePkt(pkt *HandshakePkt)(err error){
	if c.protocol != V_UNSET {
		panic("The first handshake packet is already has been received or sent")
	}
	p := NewPacket(pkt.Protocol, 0x00)
	pkt.EncodeTo(p)
	if err = c.Send(p); err != nil {
		return
	}
	c.protocol = pkt.Protocol
	return
}

func (c *Conn)Recv()(r *PacketReader, err error){
	return ReadPacket(c.protocol, c.conn)
}

func (c *Conn)RecvPkt(id int32, pkt Decodable)(err error){
	var r *PacketReader
	if r, err = ReadPacket(c.protocol, c.conn); err != nil {
		return
	}
	if r.Id() != (VarInt)(id) {
		return &PktIdAssertError{ Require: id, Got: (int32)(r.Id()) }
	}
	if err = pkt.DecodeFrom(r); err != nil {
		return
	}
	return
}

func (c *Conn)RecvHandshakePkt()(pkt *HandshakePkt, err error){
	if c.protocol != V_UNSET {
		panic("The first handshake packet is already has been received or sent")
	}

	if pkt, err = readHandshakePacket(c.conn); err != nil {
		return
	}
	c.protocol = pkt.Protocol
	return
}

func readHandshakePacket(r io.Reader)(p *HandshakePkt, err error){
	var (
		varb [5]byte
		n int = 0
		size int32
	)
	{ // read varint
		var (
			i int = 0
			v0 uint32 = 0
		)
		for {
			if _, err = r.Read(varb[n:n + 1]); err != nil {
				return
			}
			b := varb[n]
			if n == 0 && b == 0xfe {
				return nil, ErrOldHandshake
			}
			n++
			v0 |= (uint32)(b & 0x7f) << i
			if b & 0x80 == 0 {
				size = (int32)(v0)
				break
			}
			if i += 7; i >= 32 {
				return nil, VarIntTooBig
			}
		}
	}
	pr := &PacketReader{
		protocol: V_UNSET, // we don't know the protocol before handshake
		buf: make([]byte, (int32)(n) + size),
		off: n,
	}
	copy(pr.buf[:n], varb[:n])
	if _, err = io.ReadFull(r, pr.buf[n:]); err != nil {
		return nil, err
	}
	var ok bool
	if pr.id, ok = pr.VarInt(); !ok {
		return nil, io.EOF
	}
	if pr.id != 0x00 {
		return nil, &PktIdAssertError{ Require: 0x00, Got: (int32)(pr.id) }
	}
	p = new(HandshakePkt)
	if err = p.Decode(pr); err != nil {
		p = nil
		return
	}
	return
}
