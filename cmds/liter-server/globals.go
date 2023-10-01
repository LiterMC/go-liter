
package main

import (
	"encoding/json"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	path := filepath.Join(configDir, "config.yml")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
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
			if data, err = yaml.Marshal(cfg); err != nil {
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
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		loger.Fatalf("Cannot parse config file: %v", err)
		return
	}
	loger.Infof("Loaded config file")
	return
}

type Whitelist struct {
	UUIDWhitelist []uuid.UUID `json:"uuids"`
	IPWhitelist   []string    `json:"ips"`
}

var whitelist = loadWhitelist()

func getWhitelist()(Whitelist){
	configLock.RLock()
	defer configLock.RUnlock()
	return whitelist
}

func loadWhitelist()(wl Whitelist){
	path := filepath.Join(configDir, "whitelist.json")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			wl.UUIDWhitelist = make([]uuid.UUID, 0)
			wl.IPWhitelist = make([]string, 0)
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
	sort.Slice(wl.UUIDWhitelist, func(i, j int)(bool){
		return uuidLess(wl.UUIDWhitelist[i], wl.UUIDWhitelist[j])
	})
	loger.Infof("Whitelist is loaded")
	return
}

func (wl Whitelist)HasUUID(id uuid.UUID)(ok bool){
	l := len(wl.UUIDWhitelist)
	i := sort.Search(l, func(i int)(bool){
		return !uuidLess(wl.UUIDWhitelist[i], id)
	})
	return i < l && wl.UUIDWhitelist[i] == id
}

func (wl Whitelist)IncludeIP(ip net.IP)(ok bool){
	for _, o := range wl.IPWhitelist {
		if matchIP(o, ip) {
			return true
		}
	}
	return false
}


type Blacklist struct {
	NameBlacklist []string    `json:"names"`
	UUIDBlacklist []uuid.UUID `json:"uuids"`
	IPBlacklist   []string    `json:"ips"`
}

var blacklist = loadBlacklist()

func getBlacklist()(Blacklist){
	configLock.RLock()
	defer configLock.RUnlock()
	return blacklist
}

func loadBlacklist()(bl Blacklist){
	path := filepath.Join(configDir, "blacklist.json")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			bl.NameBlacklist = make([]string, 0)
			bl.UUIDBlacklist = make([]uuid.UUID, 0)
			bl.IPBlacklist = make([]string, 0)
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
	sort.Strings(bl.NameBlacklist)
	sort.Slice(bl.UUIDBlacklist, func(i, j int)(bool){
		return uuidLess(bl.UUIDBlacklist[i], bl.UUIDBlacklist[j])
	})
	loger.Infof("Blacklist is loaded")
	return
}

func (bl Blacklist)HasName(name string)(ok bool){
	i := sort.SearchStrings(bl.NameBlacklist, name)
	return i < len(bl.NameBlacklist) && bl.NameBlacklist[i] == name
}

func (bl Blacklist)HasUUID(id uuid.UUID)(ok bool){
	l := len(bl.UUIDBlacklist)
	i := sort.Search(l, func(i int)(bool){
		return !uuidLess(bl.UUIDBlacklist[i], id)
	})
	return i < l && bl.UUIDBlacklist[i] == id
}

func (bl Blacklist)IncludeIP(ip net.IP)(ok bool){
	for _, o := range bl.IPBlacklist {
		if matchIP(o, ip) {
			return true
		}
	}
	return false
}


func (cfg Config)CheckConn(ip net.IP, username string)(ok bool){
	wl := getWhitelist()
	bl := getBlacklist()

	if bl.HasName(username) {
		return false
	}
	if bl.IncludeIP(ip) {
		return false
	}
	if cfg.EnableIPWhitelist {
		if !wl.IncludeIP(ip) {
			return false
		}
	}
	_, uid, err := liter.DefaultAuthClient.GetPlayerUUID(username)
	if err != nil {
		loger.Errorf("Cannot get uuid of player '%s': %v", username, err)
		return false
	}
	if cfg.EnableWhitelist {
		if !wl.HasUUID(uid) {
			return false
		}
	}
	return true
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
