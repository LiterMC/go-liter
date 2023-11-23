package liter

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"sync"
)

const DefaultPort = 25565

var (
	ErrNoRecordFound = errors.New("No DNS record was found")
	ErrOldHandshake  = errors.New("Old client(<=1.6) handshake received") // The handshake packet <= 1.6 which first byte is 0xFE
)

type PktIdAssertError struct {
	Require int32
	Got     int32
}

var _ error = (*PktIdAssertError)(nil)

func (e *PktIdAssertError) Error() string {
	return fmt.Sprintf("PktIdAssertError: require=%d; got=%d", e.Require, e.Got)
}

type Conn struct {
	conn      net.Conn
	protocol  int
	threshold int
	sendMux   sync.Mutex
}

// WrapConn wraps a raw net.Conn as a minecraft connection
func WrapConn(c net.Conn) *Conn {
	return &Conn{
		protocol:  V_UNSET,
		conn:      c,
		threshold: -1,
	}
}

// Resolve the minecraft server's hostport with the given address and cancel context
func ResloveAddrWithContext(ctx context.Context, target string) (addr *net.TCPAddr, err error) {
	addr = &net.TCPAddr{
		Port: DefaultPort,
	}
	host, port, err := net.SplitHostPort(target)
	if err != nil {
		if strings.Contains(err.Error(), "missing port in address") {
			host = target
			port = ""
			err = nil
		} else {
			return
		}
	}
	if len(port) == 0 {
		// if target is host only, then we lookup srv
		var srvs []*net.SRV
		if _, srvs, err = net.DefaultResolver.LookupSRV(ctx, "minecraft", "tcp", host); err == nil && len(srvs) > 0 {
			srv := srvs[0]
			host, addr.Port = srv.Target, (int)(srv.Port)
		}
	} else if addr.Port, err = strconv.Atoi(port); err != nil {
		return
	}
	var ips []net.IP
	if ips, err = net.DefaultResolver.LookupIP(ctx, "ip", host); err != nil {
		return nil, err
	}
	if len(ips) == 0 {
		return nil, ErrNoRecordFound
	}
	addr.IP = ips[0]
	return
}

// Resolve the minecraft server's hostport with the given address
func ResloveAddr(target string) (addr *net.TCPAddr, err error) {
	return ResloveAddrWithContext(context.Background(), target)
}

// Dial to a minecraft server with given context
func DialWithContext(ctx context.Context, target string) (c *Conn, err error) {
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

// Dial to a minecraft server
func Dial(target string) (c *Conn, err error) {
	return DialWithContext(context.Background(), target)
}

// Protocol returns the cached protocol from HandshakePkt
func (c *Conn) Protocol() int {
	return c.protocol
}

// RawConn returns the underlying connection
func (c *Conn) RawConn() net.Conn {
	return c.conn
}

// LocalAddr returns the LocalAddr of the underlying connection
func (c *Conn) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr returns the RemoteAddr of the underlying connection
func (c *Conn) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// Close will directly close the underlying connection
func (c *Conn) Close() (err error) {
	return c.conn.Close()
}

// Compressed returns whether the connection is compressed or not (threshold >= 0)
func (c *Conn) Compressed() bool {
	return c.threshold >= 0
}

// Threshold returns the current connection's data compress threshold.
// Less than zero means the threshold is unset
func (c *Conn) Threshold() int {
	return c.threshold
}

// SetThreshold will set the current connection's data compress threshold.
// It should only be called once
func (c *Conn) SetThreshold(threshold int) {
	if c.threshold >= 0 && threshold < 0 {
		panic("liter: Conn.SetThreshold: Wrong usage: cannot cancel a compressed connection")
	}
	c.threshold = threshold
}

// Send sends data from a PacketBuilder
// It's thread-safe
func (c *Conn) Send(p *PacketBuilder) (err error) {
	c.sendMux.Lock()
	defer c.sendMux.Unlock()
	if c.Compressed() {
		if _, err = p.WriteCompressedTo(c.conn, c.threshold); err != nil {
			return
		}
	} else if _, err = p.WriteTo(c.conn); err != nil {
		return
	}
	return
}

// SendPkt encode the encodable value into a PacketBuilder with given ID and cached protocol.
// It will pass the builder to method Send.
// It's thread-safe
func (c *Conn) SendPkt(id int32, value Encodable) (err error) {
	p := NewPacket(c.protocol, (VarInt)(id))
	value.Encode(p)
	return c.Send(p)
}

// SendHandshakePkt will encode the HandshakePkt into a PacketBuilder.
// For each connection, only one of the SendHandshakePkt or RecvHandshakePkt can be called, and it should be only called once.
// It's thread-safe, but you shouldn't call it more than one time
func (c *Conn) SendHandshakePkt(pkt *HandshakePkt) (err error) {
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

// Recv read a packet and returns a PacketReader.
// The PacketReader is safely to read concurrently, but this method cannot be called concurrently.
func (c *Conn) Recv() (r *PacketReader, err error) {
	return ReadPacket(c.protocol, c.conn, c.Compressed())
}

// RecvPkt read a packet and parse it into the Decodable.
// It will assert the id while parsing the packet.
// This method cannot be called concurrently
func (c *Conn) RecvPkt(id int32, pkt Decodable) (err error) {
	var r *PacketReader
	if r, err = c.Recv(); err != nil {
		return
	}
	if r.Id() != (VarInt)(id) {
		return &PktIdAssertError{Require: id, Got: (int32)(r.Id())}
	}
	if err = pkt.DecodeFrom(r); err != nil {
		return
	}
	return
}

// For each connection, only one of the SendHandshakePkt or RecvHandshakePkt can be called, and it should be only called once.
// This method cannot be called concurrently
func (c *Conn) RecvHandshakePkt() (pkt *HandshakePkt, err error) {
	if c.protocol != V_UNSET {
		panic("The first handshake packet is already has been received or sent")
	}

	if pkt, err = readHandshakePacket(c.conn); err != nil {
		return
	}
	c.protocol = pkt.Protocol
	return
}

func readHandshakePacket(r io.Reader) (p *HandshakePkt, err error) {
	var (
		varb [5]byte
		n    int = 0
		size int32
	)
	{ // read varint
		var (
			i  int    = 0
			v0 uint32 = 0
		)
		for {
			if _, err = r.Read(varb[n : n+1]); err != nil {
				return
			}
			b := varb[n]
			if n == 0 && b == 0xfe {
				return nil, ErrOldHandshake
			}
			n++
			v0 |= (uint32)(b&0x7f) << i
			if b&0x80 == 0 {
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
		buf:      make([]byte, (int32)(n)+size),
		off:      n,
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
		return nil, &PktIdAssertError{Require: 0x00, Got: (int32)(pr.id)}
	}
	p = new(HandshakePkt)
	if err = p.Decode(pr); err != nil {
		p = nil
		return
	}
	return
}
