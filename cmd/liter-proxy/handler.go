package main

import (
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/script"
	"github.com/kmcsr/go-logger"
	"golang.org/x/net/proxy"
)

func (s *ProxyServer) handle(c *liter.Conn) {
	preventCliSideClose := false
	wc := manager.WrapConn(c)
	defer func() {
		if !preventCliSideClose {
			wc.Close()
		}
	}()
	rc := c.RawConn()

	wc.OnClose = func() {
		wc.Emit(&script.Event{Name: "close", Data: Map{"conn": wc.Exports()}})
	}

	ploger := logger.NewPrefixLogger(loger, "client [%v]:", c.RemoteAddr())
	ploger.Debugf("Connected!")
	var err error
	var hp *liter.HandshakePkt
	rc.SetReadDeadline(time.Now().Add(time.Second * 5))
	if hp, err = c.RecvHandshakePkt(); err != nil {
		ploger.Debugf("Read handshake packet error: %v", err)
		return
	}
	rc.SetReadDeadline(time.Time{})
	ploger.Debugf("Handshake packet: %v", hp)
	isPing := hp.NextState == liter.NextPingState
	isLogin := hp.NextState == liter.NextLoginState

	item, ok := cfg.ProxyMap[hp.Addr]
	if !ok {
		ploger.Infof("Trying to connect with unexpected address %q", hp.Addr)
		return
	}

	outHp := new(liter.HandshakePkt)
	*outHp = *hp
	if item.ForwardAddr != "" {
		outHp.Addr = item.ForwardAddr
	}
	if item.ForwardPort != 0 {
		outHp.Port = item.ForwardPort
	}

	noforward := <-manager.Emit(script.NewEvent("handshake", Map{
		"client":    wc.Exports(),
		"handshake": hp,
		"target":    item.Target,
	}))
	if wc.Closed() {
		return
	}
	if noforward {
		for !wc.Closed() {
			var r *script.WrappedPacketReader
			if r, err = wc.Recv(); err != nil {
				if !wc.Closed() {
					wc.Emit(&script.Event{
						Name: "error",
						Data: Map{
							"conn":  wc.Exports(),
							"error": err.Error(),
						},
					})
				}
				return
			}
			wc.Emit(script.NewEvent("packet", Map{
				"conn":   wc,
				"packet": r,
			}))
		}
		return
	}

	var lp liter.LoginStartPkt
	var player Map // only exists when login

	if isLogin {
		if err = c.RecvPkt(0x00, &lp); err != nil {
			ploger.Debugf("Read login start packet error: %v", err)
			return
		}
		player = Map{
			"name": lp.Name,
		}
		if lp.Id.Ok {
			player["id"] = lp.Id.V.String()
		}
	} else if !isPing {
		// unknown type connection
		return
	}

	var addr *net.TCPAddr
	if addr, err = liter.ResloveAddrWithContext(s.ctx, item.Target); err != nil {
		ploger.Errorf("Cannot resolve addr of %q: %v", item.Target, err)
		return
	}
	var cr net.Conn
	if s.Dialer == nil {
		cr, err = proxy.Dial(s.ctx, "tcp", addr.String())
	} else {
		cr, err = s.Dialer.DialContext(s.ctx, "tcp", addr.String())
	}
	if err != nil {
		ploger.Errorf("Cannot dial to %q: %v", item.Target, err)
		return
	}
	conn := liter.WrapConn(cr)
	wconn := manager.WrapConn(conn)
	preventSvrSideClose := false
	defer func() {
		if !preventSvrSideClose {
			wconn.Close()
		}
	}()
	wconn.OnClose = func() {
		wconn.Emit(&script.Event{Name: "close", Data: Map{"conn": wconn.Exports()}})
	}

	if err = conn.SendHandshakePkt(hp); err != nil {
		ploger.Errorf("Connection handshake error: %v", err)
		return
	}
	ploger.Debugf("Handshake sent successed")

	if isLogin {
		if err = conn.SendPkt(0x00, lp); err != nil {
			ploger.Errorf("Connection login error: %v", err)
			return
		}
		ploger.Debugf("Login start packet sent successed")
	}

	if !<-manager.Emit(script.NewEvent("serve", Map{
		"client":    wc.Exports(),
		"server":    wconn.Exports(),
		"player":    player,
		"handshake": hp,
	})) {
		if proxyRawConn(ploger, cr, rc) {
			// if connection reset by peer
			if isLogin {
				loginDisconnect(c, "Connection reset by peer")
			}
		}
		return
	}

	loger.Debugf("Proxied login process")

	var sp *liter.LoginSuccessPkt
	if sp, err = proxyLoginPackets(s, conn, c); err != nil {
		ploger.Debugf("Cannot login: %v", err)
		return
	}
	_ = sp

	errCh1 := parseAndForward(wconn, wc)
	errCh2 := parseAndForward(wc, wconn)
	select {
	case err := <-errCh1:
		ploger.Errorf("Error at client connection: %v", err)
		if <-wconn.Emit(script.NewEvent("before_close", Map{
			"conn":  wconn.Exports(),
			"error": err.Error(),
		})) {
			ploger.Infof("Server connection default close action prevented")
			preventSvrSideClose = true
		}
	case err := <-errCh2:
		ploger.Errorf("Error at server connection: %v", err)
		if <-wc.Emit(script.NewEvent("before_close", Map{
			"conn":  wc.Exports(),
			"error": err.Error(),
		})) {
			ploger.Infof("Client connection default close action prevented")
			preventCliSideClose = true
		}
	}
}

func handleServerStatus(loger logger.Logger, c *liter.Conn, status string, motd string) {
	var srp liter.StatusRequestPkt
	var err error
	if err = c.RecvPkt(0x00, &srp); err != nil {
		loger.Debugf("Read status request packet error: %v", err)
		return
	}
	if err = c.SendPkt(0x00, liter.StatusResponsePkt{
		Payload: liter.StatusResponsePayload{
			Version: liter.ProtocolVersion{
				Name:     status,
				Protocol: c.Protocol(),
			},
			Players: liter.PlayerStatus{
				Max:    2,
				Online: 1,
				Sample: []liter.PlayerInfo{
					{Name: status}, // to allow hover for the status
				},
			},
			Description: liter.NewChatFromString(motd),
		},
	}); err != nil {
		loger.Debugf("Send packet error: %v", err)
		return
	}
	var prp liter.PingRequestPkt
	if err = c.RecvPkt(0x01, &prp); err != nil {
		loger.Debugf("Read ping request packet error: %v", err)
		return
	}
	if err = c.SendPkt(0x01, (liter.PingResponsePkt)(prp)); err != nil {
		loger.Debugf("Send ping response packet error: %v", err)
		return
	}
}

func proxyRawConn(ploger logger.Logger, cr, rc net.Conn) bool {
	buf := make([]byte, 32*1024)
	cr.SetReadDeadline(time.Now().Add(time.Millisecond * 10))
	// try read to ensure the connection is ok
	if n, err := cr.Read(buf); err != nil {
		if strings.Contains(err.Error(), "reset by peer") {
			ploger.Errorf("Connection reset by peer")
			return true
		}
	} else if n != 0 {
		rc.Write(buf[:n])
	}
	cr.SetReadDeadline(time.Time{}) // clear read deadline

	go func() {
		defer cr.Close()
		defer rc.Close()
		buf := make([]byte, 32*1024)
		io.CopyBuffer(rc, cr, buf)
	}()
	io.CopyBuffer(cr, rc, buf)
	return false
}

func loginDisconnectByErr(c *liter.Conn, e error) (err error) {
	return c.SendPkt(0x00, &liter.DisconnectPkt{
		Reason: liter.NewChatFromString(e.Error()),
	})
}

func loginDisconnect(c *liter.Conn, reason string) (err error) {
	return c.SendPkt(0x00, &liter.DisconnectPkt{
		Reason: liter.NewChatFromString(reason),
	})
}

type DisconnectError struct {
	Reason *liter.Chat
}

func (e *DisconnectError) Error() string {
	return e.Reason.Plain()
}

func proxyLoginPackets(s *ProxyServer, svr, cli *liter.Conn) (res *liter.LoginSuccessPkt, err error) {
	res = new(liter.LoginSuccessPkt)
	var r *liter.PacketReader
	if r, err = svr.Recv(); err != nil {
		loginDisconnectByErr(cli, err)
		return
	}
	switch r.Id() {
	case 0x00: // Disconnect
		var pkt liter.DisconnectPkt
		if err = pkt.DecodeFrom(r); err != nil {
			loginDisconnectByErr(cli, err)
			return
		}
		if err = cli.SendPkt(0x00, &pkt); err != nil {
			return
		}
		return nil, &DisconnectError{Reason: pkt.Reason}
	case 0x03: // Set Compression
		var pkt liter.LoginSetCompressionPkt
		if err = pkt.DecodeFrom(r); err != nil {
			loginDisconnectByErr(cli, err)
			return
		}
		if err = cli.SendPkt(0x03, &pkt); err != nil {
			return
		}
		if pkt.Threshold >= 0 {
			cli.SetThreshold((int)(pkt.Threshold))
			svr.SetThreshold((int)(pkt.Threshold))
		}
		if r, err = svr.Recv(); err != nil {
			loginDisconnectByErr(cli, err)
			return
		}
		if r.Id() != 0x02 {
			err = &liter.PktIdAssertError{Require: 0x02, Got: (int32)(r.Id())}
			return
		}
		fallthrough
	case 0x02: // Login Success
		if err = res.DecodeFrom(r); err != nil {
			loginDisconnectByErr(cli, err)
			return
		}
		if err = cli.SendPkt(0x02, res); err != nil {
			return
		}
		return
	// case 0x01: // Encryption Request
	default:
		err = fmt.Errorf("Unexpected packet id %d", r.Id())
		loginDisconnectByErr(cli, err)
		return
	}
	return
}

func parseAndForward(dst, src *script.WrappedConn) <-chan error {
	errCh := make(chan error, 1)
	go func() {
		var err error
		defer func() {
			errCh <- err
		}()
		var pkt liter.PacketBuilder
		for !src.Closed() {
			var r *script.WrappedPacketReader
			if r, err = src.Recv(); err != nil {
				if !src.Closed() {
					src.Emit(&script.Event{
						Name: "error",
						Data: Map{
							"conn":  src.Exports(),
							"error": err.Error(),
						},
					})
				}
				return
			}
			if !<-src.Emit(script.NewEvent("packet", Map{
				"conn":   src.Exports(),
				"packet": r,
			})) {
				if err = dst.Send(pkt.Reset(r.Protocol(), r.Id()).ByteArray(r.Bytes())); err != nil {
					if !dst.Closed() {
						dst.Emit(&script.Event{
							Name: "error",
							Data: Map{
								"conn":  dst.Exports(),
								"error": err.Error(),
							},
						})
					}
					return
				}
			}
		}
	}()
	return errCh
}
