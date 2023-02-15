
package main

import (
	"encoding/json"
	"flag"
	"io"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/kmcsr/go-logger"
	logrusl "github.com/kmcsr/go-logger/logrus"
	"github.com/kmcsr/go-liter"
)

var loger = initLogger()

func initLogger()(loger logger.Logger){
	loger = logrusl.New()
	loger.SetOutput(os.Stderr)
	logrusl.Unwrap(loger).SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp: true,
	})
	loger.SetLevel(logger.InfoLevel)
	return
}

var configDir = getConfigDir()

func getConfigDir()(dir string){
	_ = _after_load
	if runtime.GOOS != "windows" {
		dir = filepath.Join("/opt", "liter_proxy")
	}else{
		dir = "."
	}
	if fi, err := os.Stat(dir); err == nil {
		if !fi.IsDir() {
			loger.Fatalf("Path '%s' is not a dir", dir)
		}
	}else if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0755); err != nil {
			loger.Fatalf("Cannot create config dir '%s': %v", dir, err)
		}
	}
	return
}

type ProxyMapItem struct{
	Target   string `json:"target"`
	PassAddr string `json:"pass-addr,omitempty"`
	PassPort uint16 `json:"pass-port,omitempty"`
}

type Config struct{
	ServerIP   string `json:"server-ip"`
	ServerPort uint16 `json:"server-port"`

	ProxyMap map[string]ProxyMapItem `json:"proxy-items,omitempty"`
}

var cfg = readConfig()

func readConfig()(cfg Config){
	file := filepath.Join(configDir, "config.json")
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			cfg.ServerIP   = ""
			cfg.ServerPort = 25565
			cfg.ProxyMap = map[string]ProxyMapItem{
				"hypixel.example": ProxyMapItem{
					Target: "mc.hypixel.net",
					PassAddr: "mc.hypixel.net",
					PassPort: 25565,
				},
			}
			if data, err = json.MarshalIndent(cfg, "", "  "); err != nil {
				loger.Fatalf("Cannot encode config data: %v", err)
				return
			}
			if err = os.WriteFile(file, data, 0644); err != nil {
				loger.Fatalf("Cannot save config data: %v", err)
				return
			}
		}else{
			loger.Fatalf("Cannot read config file: %v", err)
		}
		return
	}
	if err = json.Unmarshal(data, &cfg); err != nil {
		loger.Fatalf("Cannot parse config file: %v", err)
		return
	}
	return
}

func main(){
	var (
		listener *net.TCPListener
		err error
	)
	laddr := &net.TCPAddr{
		Port: (int)(cfg.ServerPort),
	}
	if len(cfg.ServerIP) > 0 {
		if laddr.IP = net.ParseIP(cfg.ServerIP); laddr.IP == nil {
			loger.Warnf("Cannot parse server-ip")
		}
	}
	if listener, err = net.ListenTCP("tcp", laddr); err != nil {
		loger.Fatalf("Cannot listening at [%v]: %v", laddr, err)
	}
	loger.Infof("Server listening at [%v]", listener.Addr())
	var c net.Conn
	for {
		if c, err = listener.Accept(); err != nil {
			loger.Errorf("Accept error: %v", err)
			break
		}
		go handler(liter.NewConn(c))
	}
}

func handler(c *liter.Conn){
	defer c.Close()
	loger.Infof("client [%v] connected", c.RemoteAddr())
	var (
		p *liter.PacketReader
		err error
	)
	if p, err = c.Recv(); err != nil {
		loger.Errorf("client [%v] read handshake packet error: %v", c.RemoteAddr(), err)
		return
	}
	var hp liter.HandshakeP
	if hp, err = liter.ReadHandshakeP(p); err != nil {
		loger.Errorf("client [%v] read handshake packet error: %v", c.RemoteAddr(), err)
		return
	}
	loger.Debugf("client [%v] handshake packet: %v", c.RemoteAddr(), hp)

	item, ok := cfg.ProxyMap[hp.Addr]
	if !ok {
		return
	}

	var conn net.Conn
	if conn, err = net.Dial("tcp", item.Target); err != nil {
		loger.Errorf("Cannot dial to %q: %v", item.Target, err)
		return
	}
	np := liter.NewPacket(p.Id())
	if item.PassAddr != "" {
		hp.Addr = item.PassAddr
	}
	if item.PassPort != 0 {
		hp.Port = item.PassPort
	}
	hp.Encode(np)
	if _, err = conn.Write(np.Bytes()); err != nil {
		loger.Errorf("New connection handshake error: %v", err)
		return
	}
	rc := c.RawConn()
	go io.Copy(rc, conn)
	io.Copy(conn, rc)
}
