
package liter

import (
	"context"
	"fmt"
	"net"
	"strings"
)

const DefaultPort = 25565

type PktIdAssertError struct {
	Require int32
	Got     int32
}

var _ error = (*PktIdAssertError)(nil)

func (e *PktIdAssertError)Error()(string){
	return fmt.Sprintf("PktIdAssertError: require=%d; got=%d", e.Require, e.Got)
}


type Conn struct{
	conn net.Conn
}

func NewConn(c net.Conn)(*Conn){
	return &Conn{
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
	return NewConn(c0), nil
}

func Dial(target string)(c *Conn, err error){
	return DialWithContext(context.Background(), target)
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

func (c *Conn)Send(id int32, values ...any)(err error){
	p := NewPacket(id)
	for _, v := range values {
		switch w := v.(type) {
		case Bool:
			p.Bool(w)
		case Byte:
			p.Byte(w)
		case Short:
			p.Short(w)
		case Int:
			p.Int(w)
		case Long:
			p.Long(w)
		case Float:
			p.Float(w)
		case Double:
			p.Double(w)
		case VarInt:
			p.VarInt((int32)(w))
		case VarLong:
			p.VarLong((int64)(w))
		case ByteArray:
			p.ByteArray(w)
		case String:
			p.String(w)
		case Object:
			p.JSON(w)
		case UUID:
			p.UUID(w)
		case Encodable:
			if err = w.Encode(p); err != nil {
				return
			}
		default:
			panic("Unknown non-encodable type")
		}
	}
	return c.SendPkt(p)
}

func (c *Conn)SendPkt(p *PacketBuilder)(err error){
	if _, err = p.WriteTo(c.conn); err != nil {
		return
	}
	return
}

func (c *Conn)Recv()(r *PacketReader, err error){
	return ReadPacket(c.conn)
}

func (c *Conn)RecvPkt(id int32, pkt Decodable)(err error){
	var r *PacketReader
	if r, err = ReadPacket(c.conn); err != nil {
		return
	}
	if r.Id() != id {
		return &PktIdAssertError{ Require: id, Got: r.Id() }
	}
	if err = pkt.DecodeFrom(r); err != nil {
		return
	}
	return
}
