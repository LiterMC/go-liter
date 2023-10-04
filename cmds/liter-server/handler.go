
package main

import (
	"io"
	"strings"
	"time"

	"github.com/kmcsr/go-logger"
	"github.com/kmcsr/go-liter"
)

func (s *Server)handle(c *liter.Conn){
	defer c.Close()
	rc := c.RawConn()

	cfg := getConfig()

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
	ploger.Tracef("Handshake packet: %v", hp)

	var svr *ServerIns = nil
	servers := getServers()
	F: for _, s := range servers {
		for _, n := range s.ServerNames {
			if ismatch(hp.Addr, n) {
				svr = s
				break F
			}
		}
	}
	if svr == nil {
		ploger.Infof("Trying connected with unexpected address %q", hp.Addr)
		return
	}

	ploger.Infof("Connected with address [%s:%d], passing to server %q[%s]", hp.Addr, hp.Port, svr.Id, svr.Target)

	var lp liter.LoginStartPacket

	if hp.NextState == liter.NextPingState {
		if svr.HandlePing {
			ploger.Debugf("Handle ping connection for server %q", svr.Id)
			handleServerStatus(ploger, c, "Idle", svr.Motd)
			return
		}
	}else if hp.NextState == liter.NextLoginState {
		if err = c.RecvPkt(0x00, &lp); err != nil {
			ploger.Debugf("Read login start packet error: %v", err)
			return
		}

		player, err := AuthClient.GetPlayerInfo(lp.Name)
		if err != nil {
			if cfg.OnlineMode {
				c.SendPkt(0x00, &liter.DisconnectPkt{
					Reason: liter.NewChatFromString("Your username is not exists or auth server error"),
				})
				ploger.Debugf("Cannot get player info for %s: %v", lp.Name, err)
				return
			}
			player = liter.PlayerInfo{
				Name: lp.Name,
			}
		}
		if lp.Id.Ok {
			id := lp.Id.V
			if id != player.Id {
				c.SendPkt(0x00, &liter.DisconnectPkt{
					Reason: liter.NewChatFromString("Your user profile is not match"),
				})
				return
			}
		}
		if blacklist.HasPlayer(player) {
			c.SendPkt(0x00, &liter.DisconnectPkt{
				Reason: liter.NewChatFromString("Your are in the blacklist"),
			})
			return
		}
		if cfg.EnableWhitelist && !whitelist.HasPlayer(player) {
			c.SendPkt(0x00, &liter.DisconnectPkt{
				Reason: liter.NewChatFromString("Your are not in the whitelist"),
			})
			return
		}
	}else{
		// unknown type connection
		return
	}

	var conn *liter.Conn
	if conn, err = liter.Dial(svr.Target); err != nil {
		ploger.Errorf("Cannot dial to %q: %v", svr.Target, err)
		if hp.NextState == liter.NextPingState {
			ploger.Debugf("Handle ping connection for server %q", svr.Id)
			handleServerStatus(ploger, c, "Closed", svr.MotdFailed)
		}else if hp.NextState == liter.NextLoginState {
			c.SendPkt(0x00, &liter.DisconnectPkt{
				Reason: liter.NewChatFromString("Cannot connect to the subserver"),
			})
		}
		return
	}
	ploger.Debugf("Target %q connected", svr.Id)
	if err = conn.SendHandshakePkt(hp); err != nil {
		ploger.Errorf("Connection handshake error: %v", err)
		return
	}
	ploger.Tracef("Handshake sent successed")

	if hp.NextState == liter.NextLoginState {
		if err = conn.SendPkt(0x00, lp); err != nil {
			ploger.Errorf("Connection login error: %v", err)
			return
		}
		ploger.Tracef("Login start packet sent successed")
	}

	cr := conn.RawConn()

	buf := make([]byte, 32 * 1024)
	cr.SetReadDeadline(time.Now().Add(time.Millisecond * 10))
	// try read to ensure the connection is ok
	if n, err := cr.Read(buf); err != nil {
		if strings.Contains(err.Error(), "connection reset by peer") {
			ploger.Errorf("Connection reset by peer")
			if hp.NextState == liter.NextPingState {
				ploger.Debugf("Handle ping connection for server %q", svr.Id)
				handleServerStatus(ploger, c, "Closed", svr.MotdFailed)
			}
			return
		}
	}else if n != 0 {
		rc.Write(buf[:n])
	}
	cr.SetReadDeadline(time.Time{}) // clear read deadline

	go func(){
		defer c.Close()
		defer conn.Close()
		buf := make([]byte, 32 * 1024)
		io.CopyBuffer(rc, cr, buf)
	}()
	io.CopyBuffer(cr, rc, buf)
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
