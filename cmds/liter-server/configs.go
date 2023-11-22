
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/kmcsr/go-logger"
	logrusl "github.com/kmcsr/go-logger/logrus"
	"github.com/kmcsr/go-liter"
	"github.com/gin-gonic/gin"
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

var workingDir string = func()(string){
	dir, err := os.Getwd()
	if err != nil {
		loger.Panicf("Cannot get working dir", err)
	}
	return dir
}()

type ServerIns struct {
	Id          string   `json:"id" yaml:"id"`
	Target      string   `json:"target" yaml:"target"`
	ServerNames []string `json:"names" yaml:"names"`
	// HandlePing is useful if you want to hide online players from others who won't join the server
	HandlePing  bool     `json:"handlePing" yaml:"handle-ping"`
	// Motd only use when HandlePing is true
	Motd        string   `json:"motd" yaml:"motd"`
	// MotdFailed will be send back when the ping connection failed on the server
	MotdFailed  string   `json:"motdFailed" yaml:"motd-failed"`
}

type DevConfig struct {
	Debug bool `yaml:"debug"`
	WatchPlugin        bool `yaml:"watch-plugin"`
	WatchReloadTimeout int  `yaml:"watch-reload-timeout"` // in seconds
	// liter will run `${npm-bin} run build` to load a directory plugin
	DirectoryPlugin bool `yaml:"directory-plugin"`
	Npm     string   `yaml:"npm-bin"`
	NpmArgs []string `yaml:"npm-args"`
}

type Config struct {
	OnlineMode bool `yaml:"online-mode"`

	ServerAddr string `yaml:"server-addr"`

	Dashboard struct {
		Enable bool   `yaml:"enable"`
		Addr   string `yaml:"addr"`
	} `yaml:"dashboard"`

	EnableWhitelist   bool `yaml:"enable-whitelist"`
	EnableIPWhitelist bool `yaml:"enable-ip-whitelist"`

	Servers []*ServerIns `yaml:"servers"`

	Dev DevConfig `yaml:"dev"`
}

func loadConfig()(cfg Config){
	// set config default values
	cfg.OnlineMode = true
	cfg.ServerAddr = ":25565"
	cfg.Servers = []*ServerIns{
		{
			Id: "main",
			Target: "127.0.0.1:25665",
			ServerNames: []string{ "minecraft.example.com", "anotherdomain.example.com" },
			MotdFailed: "Server is closed",
		},
	}
	cfg.Dashboard.Addr = "127.0.0.1:25580"
	cfg.Dev = DevConfig{
		WatchReloadTimeout: 5,
		Npm: "npm",
	}

	path := filepath.Join(workingDir, "config.yml")
	data, err := os.ReadFile(path)
	if err != nil {
		if !os.IsNotExist(err) {
			loger.Panicf("Cannot read config file: %v", err)
		}
	}else if err = yaml.Unmarshal(data, &cfg); err != nil {
		loger.Panicf("Cannot parse config file: %v", err)
		return
	}
	if err = cfg.Save(); err != nil {
		loger.Panicf("Cannot save config file: %v", err)
		return
	}
	loger.Infof("Loaded config file")

	if cfg.Dev.Debug {
		loger.SetLevel(logger.TraceLevel)
		loger.Debug("Debug log enabled")
		gin.SetMode(gin.DebugMode)
	}else{
		loger.SetLevel(logger.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}
	return
}

func (cfg Config)Save()(err error){
	path := filepath.Join(workingDir, "config.yml")

	var data []byte
	if data, err = yaml.Marshal(cfg); err != nil {
		return
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		return
	}
	return
}

var AuthClient = liter.DefaultAuthClient

type PlayerInfo struct {
	onlineMode bool
	liter.PlayerInfo
}

func ParsePlayerInfoFromString(v string, onlineMode bool)(p PlayerInfo, err error){
	p.onlineMode = onlineMode
	if p.Id, err = uuid.Parse(v); err == nil {
		var profile *liter.PlayerProfile
		profile, err = AuthClient.GetPlayerProfile(p.Id)
		if err != nil {
			return
		}
		p.Name = profile.Name
	}else{
		err = nil
		if onlineMode {
			var info liter.PlayerInfo
			info, err = AuthClient.GetPlayerInfo(v)
			if err != nil {
				return
			}
			p.Name = info.Name
			p.Id = info.Id
		}else{
			p.Name = v
		}
	}
	return
}

func (p PlayerInfo)IsOffline()(ok bool){
	return p.Id == uuid.Nil
}

func (p *PlayerInfo)update()(err error){
	if p.IsOffline() {
		return
	}
	if p.onlineMode {
		var profile *liter.PlayerProfile
		profile, err = AuthClient.GetPlayerProfile(p.Id)
		if err != nil {
			return
		}
		p.Name = profile.Name
	}
	if p.Name == "" {
		return fmt.Errorf("Unknown player info")
	}else if p.onlineMode {
		var info liter.PlayerInfo
		info, err = AuthClient.GetPlayerInfo(p.Name)
		if err != nil {
			return
		}
		p.Name = info.Name
		p.Id = info.Id
	}
	return
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
		p.Name = info.Name
		p.Id = info.Id
	case Map:
		var ok bool
		var id string
		p.Name, _ = v["name"].(string)
		if p.onlineMode {
			if id, ok = v["id"].(string); ok {
				var e error
				if p.Id, e = uuid.Parse(id); e != nil {
					ok = false
				}
				profile, err := AuthClient.GetPlayerProfile(p.Id)
				if err != nil {
					return nil
				}
				p.Name = profile.Name
				return nil
			}
		}
		if p.Name == "" {
			return fmt.Errorf("Unknown player info")
		}else if p.onlineMode {
			info, err := AuthClient.GetPlayerInfo(p.Name)
			if err != nil {
				return nil
			}
			p.Name = info.Name
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

func loadWhitelist()(wl Whitelist){
	path := filepath.Join(workingDir, "whitelist.json")

	wl.Players = make([]PlayerInfo, 0)
	wl.IPs = make([]string, 0)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			wl.Players = make([]PlayerInfo, 0)
			wl.IPs = make([]string, 0)
			if data, err = json.MarshalIndent(wl, "", "  "); err != nil {
				loger.Panicf("Cannot encode whitelist data: %v", err)
				return
			}
			if err = os.WriteFile(path, data, 0644); err != nil {
				loger.Panicf("Cannot create whitelist file: %v", err)
				return
			}
		}else{
			loger.Panicf("Cannot read whitelist file: %v", err)
		}
		return
	}
	if err = json.Unmarshal(data, &wl); err != nil {
		loger.Panicf("Cannot parse whitelist file: %v", err)
		return
	}
	if err = wl.Save(); err != nil {
		loger.Panicf("Cannot save whitelist: %v", err)
		return
	}
	loger.Infof("Whitelist is loaded")
	return
}

func (wl Whitelist)Save()(err error){
	path := filepath.Join(workingDir, "whitelist.json")
	var data []byte
	if data, err = json.MarshalIndent(wl, "", "  "); err != nil {
		return
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		return
	}
	return
}

func (wl Whitelist)HasPlayer(player liter.PlayerInfo, onlineMode bool)(ok bool){
	for _, p := range wl.Players {
		if !p.IsOffline() && player.Id != uuid.Nil { // player is online mode
			if p.Id == player.Id {
				return true
			}
		}else if !onlineMode && p.Name == player.Name { // check when under offline mode
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


// Clean the player list: update player informations and remove duplicated players
func (wl *Whitelist)cleanPlayers(){
	pm := make(map[PlayerInfo]struct{}, len(wl.Players))
	players := make([]PlayerInfo, 0, len(wl.Players))
	for _, p := range wl.Players {
		p.update()
		if _, ok := pm[p]; ok {
			continue
		}
		pm[p] = struct{}{}
		players = append(players, p)
	}
	wl.Players = players
}

func (wl *Whitelist)AddPlayer(v string, onlineMode bool)(err error){
	p, err := ParsePlayerInfoFromString(v, onlineMode)
	if err != nil {
		return
	}
	wl.Players = append(wl.Players, p)
	return
}

func (wl *Whitelist)RemovePlayer(v string)(err error){
	if id, err := uuid.Parse(v); err == nil {
		for i, p := range wl.Players {
			if p.Id == id {
				remove(wl.Players, i)
				break
			}
		}
	}else{
		v = strings.ToLower(v)
		for i := 0; i < len(wl.Players); {
			p := wl.Players[i]
			if strings.ToLower(p.Name) == v {
				remove(wl.Players, i)
			}else{
				i++
			}
		}
	}
	return
}

func (wl *Whitelist)AddIP(v string)(err error){
	wl.IPs = append(wl.IPs, v)
	return
}


type Blacklist struct {
	Players []PlayerInfo `json:"players"`
	IPs     []string     `json:"ips"`
}

func loadBlacklist()(bl Blacklist){
	path := filepath.Join(workingDir, "blacklist.json")

	bl.Players = make([]PlayerInfo, 0)
	bl.IPs = make([]string, 0)

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			if data, err = json.MarshalIndent(bl, "", "  "); err != nil {
				loger.Panicf("Cannot encode blacklist data: %v", err)
				return
			}
			if err = os.WriteFile(path, data, 0644); err != nil {
				loger.Panicf("Cannot create blacklist file: %v", err)
				return
			}
		}else{
			loger.Panicf("Cannot read blacklist file: %v", err)
		}
		return
	}
	if err = json.Unmarshal(data, &bl); err != nil {
		loger.Panicf("Cannot parse blacklist file: %v", err)
		return
	}
	if err = bl.Save(); err != nil {
		loger.Panicf("Cannot save blacklist: %v", err)
		return
	}
	loger.Infof("Blacklist is loaded")
	return
}

func (bl Blacklist)Save()(err error){
	path := filepath.Join(workingDir, "blacklist.json")
	var data []byte
	if data, err = json.MarshalIndent(bl, "", "  "); err != nil {
		return
	}
	if err = os.WriteFile(path, data, 0644); err != nil {
		return
	}
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

// Clean the player list: update player informations and remove duplicated players
func (bl *Blacklist)cleanPlayers(){
	pm := make(map[PlayerInfo]struct{}, len(bl.Players))
	players := make([]PlayerInfo, 0, len(bl.Players))
	for _, p := range bl.Players {
		p.update()
		if _, ok := pm[p]; ok {
			continue
		}
		pm[p] = struct{}{}
		players = append(players, p)
	}
	bl.Players = players
}

func (bl *Blacklist)AddPlayer(v string, onlineMode bool)(err error){
	p, err := ParsePlayerInfoFromString(v, onlineMode)
	if err != nil {
		return
	}
	bl.Players = append(bl.Players, p)
	return
}

func (bl *Blacklist)RemovePlayer(v string)(err error){
	if id, err := uuid.Parse(v); err == nil {
		for i, p := range bl.Players {
			if p.Id == id {
				remove(bl.Players, i)
				break
			}
		}
	}else{
		v = strings.ToLower(v)
		for i := 0; i < len(bl.Players); {
			p := bl.Players[i]
			if strings.ToLower(p.Name) == v {
				remove(bl.Players, i)
			}else{
				i++
			}
		}
	}
	return
}

func (bl *Blacklist)AddIP(v string)(err error){
	bl.IPs = append(bl.IPs, v)
	return
}
