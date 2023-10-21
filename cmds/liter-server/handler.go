
package main

import (
	"io"
	"strings"
	"time"
	"net"

	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-liter"
	"github.com/kmcsr/go-liter/script"
)

func (s *Server)handle(c *liter.Conn, cfg *Config){
	preventClose := false
	wc := s.scripts.WrapConn(c)
	defer func(){
		if !preventClose {
			wc.Close()
		}
	}()
	rc := c.RawConn()

	c0 := &Conn{
		Conn: c,
		When: time.Now(),
	}
	cid := s.conns.Put(c0)
	wc.OnClose = func(){
		s.conns.Free(cid)
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

	var svr *ServerIns = nil
	F: for _, s := range cfg.Servers {
		for _, n := range s.ServerNames {
			if ismatch(hp.Addr, n) {
				svr = s
				break F
			}
		}
	}
	if svr == nil {
		ploger.Infof("Trying to connect with unexpected address %q", hp.Addr)
		return
	}

	noforward := <-s.scripts.Emit(script.NewEvent("handshake", Map{
		"client": wc.Exports(),
		"handshake": hp,
		"target": &svr, // to allow changes to the target
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
							"conn": wc.Exports(),
							"error": err.Error(),
						},
					})
				}
				return
			}
			<-wc.Emit(script.NewEvent("packet", Map{
				"conn": wc,
				"packet": r,
			}))
		}
		return
	}

	c0.SetLocalServer(hp.Addr, svr.Id)
	ploger.Infof("Connected with address [%s:%d], passing to server %q[%s]", hp.Addr, hp.Port, svr.Id, svr.Target)

	var lp liter.LoginStartPkt
	var player Map // only exists when login

	if isPing {
		if svr.HandlePing {
			ploger.Debugf("Handle ping connection for server %q", svr.Id)
			handleServerStatus(ploger, c, "Idle", svr.Motd)
			return
		}
	}else if isLogin {
		if err = c.RecvPkt(0x00, &lp); err != nil {
			ploger.Debugf("Read login start packet error: %v", err)
			return
		}

		var pl liter.PlayerInfo
		pl, err = AuthClient.GetPlayerInfo(lp.Name)
		if err != nil {
			if cfg.OnlineMode {
				c.SendPkt(0x00, &liter.DisconnectPkt{
					Reason: liter.NewChatFromString("Your username is not exists or auth server error"),
				})
				ploger.Debugf("Cannot get player info for %s: %v", lp.Name, err)
				return
			}
			pl = liter.PlayerInfo{
				Name: lp.Name,
			}
		}
		if lp.Id.Ok {
			id := lp.Id.V
			if id != pl.Id {
				c.SendPkt(0x00, &liter.DisconnectPkt{
					Reason: liter.NewChatFromString("Your user profile is not match"),
				})
				return
			}
		}
		if blacklist.HasPlayer(pl) {
			c.SendPkt(0x00, &liter.DisconnectPkt{
				Reason: liter.NewChatFromString("Your are in the blacklist"),
			})
			return
		}
		if cfg.EnableWhitelist && !whitelist.HasPlayer(pl) {
			c.SendPkt(0x00, &liter.DisconnectPkt{
				Reason: liter.NewChatFromString("Your are not in the whitelist"),
			})
			return
		}
		player = Map{
			"name": pl.Name,
			"id": pl.Id.String(),
		}
		c0.SetPlayer(pl)
	}else{
		// unknown type connection
		return
	}

	var conn *liter.Conn
	if conn, err = liter.Dial(svr.Target); err != nil {
		ploger.Errorf("Cannot dial to %q: %v", svr.Target, err)
		if isPing {
			ploger.Debugf("Handle ping connection for server %q", svr.Id)
			handleServerStatus(ploger, c, "Closed", svr.MotdFailed)
		}else if isLogin {
			c.SendPkt(0x00, &liter.DisconnectPkt{
				Reason: liter.NewChatFromString("Cannot connect to the subserver"),
			})
		}
		return
	}
	defer conn.Close()

	ploger.Debugf("Target %q connected", svr.Id)

	if err = conn.SendHandshakePkt(hp); err != nil {
		ploger.Errorf("Connection handshake error: %v", err)
		return
	}
	ploger.Debugf("Handshake sent successed")

	wconn := s.scripts.WrapConn(conn)

	if isLogin {
		if err = conn.SendPkt(0x00, lp); err != nil {
			ploger.Errorf("Connection login error: %v", err)
			return
		}
		ploger.Debugf("Login start packet sent successed")
	}

	cr := conn.RawConn()

	if !<-s.scripts.Emit(script.NewEvent("serve", Map{
		"client": wc.Exports(),
		"server": wconn.Exports(),
		"player": player,
		"handshake": hp,
	})) {
		if proxyRawConn(ploger, cr, rc) {
			// if connection reset by peer
			if isPing {
				ploger.Debugf("Handle ping connection for server %q", svr.Id)
				handleServerStatus(ploger, c, "Closed", svr.MotdFailed)
			}
		}
		return
	}

	errCh1 := parseAndForward(wconn, wc)
	errCh2 := parseAndForward(wc, wconn)
	select {
	case err := <-errCh1:
		ploger.Errorf("Error at client connection: %v", err)
		if !<-wc.Emit(script.NewEvent("close", Map{
			"conn": wc.Exports(),
			"error": err,
		})) {
			ploger.Infof("Server connection default close action prevented")
			wconn.Close()
		}
	case err := <-errCh2:
		ploger.Errorf("Error at server connection: %v", err)
		if <-wconn.Emit(script.NewEvent("close", Map{
			"conn": wconn.Exports(),
			"error": err,
		})) {
			ploger.Infof("Client connection default close action prevented")
			preventClose = true
		}
	}
}


func handleServerStatus(loger logger.Logger, c *liter.Conn, status string, motd string){
	var srp liter.StatusRequestPkt
	var err error
	if err = c.RecvPkt(0x00, &srp); err != nil {
		loger.Debugf("Read status request packet error: %v", err)
		return
	}
	if err = c.SendPkt(0x00, liter.StatusResponsePkt{
		Payload: liter.StatusResponsePayload{
			Version: liter.ProtocolVersion{
				Name: status,
				Protocol: c.Protocol(),
			},
			Players: liter.PlayerStatus{
				Max: 2,
				Online: 1,
				Sample: []liter.PlayerInfo{
					{ Name: status }, // to allow hover for the status
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

func proxyRawConn(ploger logger.Logger, cr, rc net.Conn)(bool){
	buf := make([]byte, 32 * 1024)
	cr.SetReadDeadline(time.Now().Add(time.Millisecond * 10))
	// try read to ensure the connection is ok
	if n, err := cr.Read(buf); err != nil {
		if strings.Contains(err.Error(), "reset by peer") {
			ploger.Errorf("Connection reset by peer")
			return true
		}
	}else if n != 0 {
		rc.Write(buf[:n])
	}
	cr.SetReadDeadline(time.Time{}) // clear read deadline

	go func(){
		defer cr.Close()
		defer rc.Close()
		buf := make([]byte, 32 * 1024)
		io.CopyBuffer(rc, cr, buf)
	}()
	io.CopyBuffer(cr, rc, buf)
	return false
}

func parseAndForward(dst, src *script.WrappedConn)(<-chan error){
	errCh := make(chan error, 1)
	go func(){
		var err error
		defer func(){
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
							"src": src.Exports(),
							"dst": dst.Exports(),
							"error": err.Error(),
						},
					})
				}
				return
			}
			if !<-src.Emit(script.NewEvent("packet", Map{
				"src": src.Exports(),
				"dst": dst.Exports(),
				"packet": r,
			})) {
				if err = dst.Send(pkt.Reset(r.Protocol(), r.Id()).ByteArray(r.Bytes())); err != nil {
					if !dst.Closed() {
						dst.Emit(&script.Event{
							Name: "error",
							Data: Map{
								"src": src.Exports(),
								"dst": dst.Exports(),
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
