
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"gopkg.in/yaml.v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/kmcsr/go-logger"
	logrusl "github.com/kmcsr/go-logger/logrus"
	"github.com/kmcsr/go-liter"
)

var loger = initLogger()

func initLogger()(loger logger.Logger){
	loger = logrusl.New()
	loger.SetOutput(os.Stdout)
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
		dir = filepath.Join("/opt", "litermc")
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

var configLock sync.RWMutex

func reloadConfigs(){
	loger.Infof("Reloading config...")

	configLock.Lock()
	defer configLock.Unlock()

	config = readConfig()
	whitelist = loadWhitelist()
	blacklist = loadBlacklist()
}

type Config struct {
	Debug bool `json:"debug" yaml:"debug"`
	OnlineMode bool `json:"online-mode" yaml:"online-mode"`

	ServerIP   string `json:"server-ip" yaml:"server-ip"`
	ServerPort uint16 `json:"server-port" yaml:"server-port"`

	EnableWhitelist   bool `json:"enable-whitelist" yaml:"enable-whitelist"`
	EnableIPWhitelist bool `json:"enable-ip-whitelist" yaml:"enable-ip-whitelist"`

	Servers []*ServerIns `json:"servers" yaml:"servers"`
}

var config = readConfig()

func getConfig()(Config){
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func readConfig()(cfg Config){
	_ = _after_load

	// set config default values
	cfg.OnlineMode = true
	cfg.ServerIP   = ""
	cfg.ServerPort = 25565
	cfg.Servers = []*ServerIns{
		{
			Id: "main",
			Target: "127.0.0.1:25665",
			ServerNames: []string{ "minecraft.example.com", "anotherdomain.example.com" },
			MotdFailed: "Server is closed",
		},
	}

	path := filepath.Join(configDir, "config.yml")
	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			loger.Fatalf("Cannot read config file: %v", err)
		}
	}else if err = yaml.Unmarshal(data, &cfg); err != nil {
		loger.Fatalf("Cannot parse config file: %v", err)
		return
	}
	if data, err = yaml.Marshal(cfg); err != nil {
		loger.Fatalf("Cannot encode config data: %v", err)
		return
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		loger.Fatalf("Cannot rewrite config file: %v", err)
		return
	}
	loger.Infof("Loaded config file")
	return
}

var AuthClient = liter.DefaultAuthClient

type PlayerInfo struct {
	liter.PlayerInfo
}

func (p *PlayerInfo)UnmarshalJSON(buf []byte)(err error){
	var v any
	if err = json.Unmarshal(buf, &v); err != nil {
		return
	}
	switch v := v.(type) {
	case string:
		var err error
		if p.Id, err = uuid.Parse(v); err == nil {
			profile, err := AuthClient.GetPlayerProfile(p.Id)
			if err != nil {
				return nil
			}
			p.Name = profile.Name
			break
		}
		p.Name = v
		info, err := AuthClient.GetPlayerInfo(p.Name)
		if err != nil {
			return nil
		}
		p.Id = info.Id
	case map[string]any:
		var ok bool
		var id string
		p.Name, _ = v["name"].(string)
		if config.OnlineMode {
			if id, ok = v["id"].(string); ok {
				var e error
				if p.Id, e = uuid.Parse(id); e != nil {
					ok = false
				}
				profile, err := AuthClient.GetPlayerProfile(p.Id)
				if err != nil {
					return nil
				}
				if p.Name != profile.Name {
					p.Name = profile.Name
				}
				return nil
			}
		}
		if p.Name == "" {
			return fmt.Errorf("Unknown player info")
		}else if config.OnlineMode {
			info, err := AuthClient.GetPlayerInfo(p.Name)
			if err != nil {
				return nil
			}
			p.Id = info.Id
		}
	default:
		return fmt.Errorf("Unexpected player info (%T)%v", v, v)
	}
	return
}

type Whitelist struct {
	Players []PlayerInfo `json:"players"`
	IPs     []string     `json:"ips"`
}

var whitelist = loadWhitelist()

func getWhitelist()(Whitelist){
	configLock.RLock()
	defer configLock.RUnlock()
	return whitelist
}

func loadWhitelist()(wl Whitelist){
	path := filepath.Join(configDir, "whitelist.json")

	wl.Players = make([]PlayerInfo, 0)
	wl.IPs = make([]string, 0)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			wl.Players = make([]PlayerInfo, 0)
			wl.IPs = make([]string, 0)
			if data, err = json.MarshalIndent(wl, "", "  "); err != nil {
				loger.Fatalf("Cannot encode config data: %v", err)
				return
			}
			if err = os.WriteFile(path, data, 0644); err != nil {
				loger.Fatalf("Cannot create config file: %v", err)
				return
			}
		}else{
			loger.Fatalf("Cannot read config file: %v", err)
		}
		return
	}
	if err = json.Unmarshal(data, &wl); err != nil {
		loger.Fatalf("Cannot parse whitelist file: %v", err)
		return
	}
	if data, err = json.MarshalIndent(wl, "", "  "); err != nil {
		loger.Fatalf("Cannot encode config data: %v", err)
		return
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		loger.Fatalf("Cannot rewrite config file: %v", err)
		return
	}
	loger.Infof("Whitelist is loaded")
	return
}

func (wl Whitelist)HasPlayer(player liter.PlayerInfo)(ok bool){
	for _, p := range wl.Players {
		if p.Id != uuid.Nil && player.Id != uuid.Nil { // player is online mode
			if p.Id == player.Id {
				return true
			}
		}else if p.Name == player.Name { // check when under offline mode (cannot access mojang api)
			return true
		}
	}
	return false
}

func (wl Whitelist)IncludeIP(ip net.IP)(ok bool){
	for _, o := range wl.IPs {
		if matchIP(o, ip) {
			return true
		}
	}
	return false
}


type Blacklist struct {
	Players []PlayerInfo `json:"players"`
	IPs     []string     `json:"ips"`
}

var blacklist = loadBlacklist()

func getBlacklist()(Blacklist){
	configLock.RLock()
	defer configLock.RUnlock()
	return blacklist
}

func loadBlacklist()(bl Blacklist){
	path := filepath.Join(configDir, "blacklist.json")

	bl.Players = make([]PlayerInfo, 0)
	bl.IPs = make([]string, 0)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			if data, err = json.MarshalIndent(bl, "", "  "); err != nil {
				loger.Fatalf("Cannot encode config data: %v", err)
				return
			}
			if err = os.WriteFile(path, data, 0644); err != nil {
				loger.Fatalf("Cannot create config file: %v", err)
				return
			}
		}else{
			loger.Fatalf("Cannot read config file: %v", err)
		}
		return
	}
	if err = json.Unmarshal(data, &bl); err != nil {
		loger.Fatalf("Cannot parse blacklist file: %v", err)
		return
	}
	if data, err = json.MarshalIndent(bl, "", "  "); err != nil {
		loger.Fatalf("Cannot encode config data: %v", err)
		return
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		loger.Fatalf("Cannot create config file: %v", err)
		return
	}
	loger.Infof("Blacklist is loaded")
	return
}

func (bl Blacklist)HasPlayer(player liter.PlayerInfo)(ok bool){
	for _, p := range bl.Players {
		if p.Name == player.Name || p.Id == player.Id { // for blacklist, we block both username and uuid
			return true
		}
	}
	return false
}

func (bl Blacklist)IncludeIP(ip net.IP)(ok bool){
	for _, o := range bl.IPs {
		if matchIP(o, ip) {
			return true
		}
	}
	return false
}

type ServerIns struct {
	Id          string   `json:"id" yaml:"id"`
	Target      string   `json:"target" yaml:"target"`
	ServerNames []string `json:"names" yaml:"names"`
	HandlePing  bool     `json:"handle-ping" yaml:"handle-ping"`
	Motd        string   `json:"motd" yaml:"motd"`
	MotdFailed  string   `json:"motd-failed" yaml:"motd-failed"`
}

func getServers()([]*ServerIns){
	return getConfig().Servers
}
