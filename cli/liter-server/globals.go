
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"

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
		dir = filepath.Join("/etc", "litermc")
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
	servers = loadServerIns()
}

type Config struct{
	Debug bool `json:"debug"`

	ServerIP   string `json:"server-ip"`
	ServerPort uint16 `json:"server-port"`

	EnableWhitelist   bool `json:"enable-whitelist"`
	EnableIPWhitelist bool `json:"enable-ip-whitelist"`
}

var config = readConfig()

func getConfig()(Config){
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func readConfig()(cfg Config){
	_ = _after_load
	path := filepath.Join(configDir, "config.json")
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			cfg.ServerIP   = ""
			cfg.ServerPort = 25565
			if data, err = json.MarshalIndent(cfg, "", "  "); err != nil {
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
	if err = json.Unmarshal(data, &cfg); err != nil {
		loger.Fatalf("Cannot parse config file: %v", err)
		return
	}
	loger.Infof("Loaded config file")
	return
}

type Whitelist struct{
	UUIDWhitelist []uuid.UUID `json:"uuid-whitelist"`
	IPWhitelist   []string    `json:"ip-whitelist"`
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
	sort.Slice(wl.UUIDWhitelist, func(i, j int)(bool){ return uuidLess(wl.UUIDWhitelist[i], wl.UUIDWhitelist[j]) })
	loger.Infof("Loaded whitelist")
	return
}

type Blacklist struct{
	NameBlacklist []string    `json:"name-blacklist"`
	UUIDBlacklist []uuid.UUID `json:"uuid-blacklist"`
	IPBlacklist   []string    `json:"ip-blacklist"`
}

var blacklist = loadBlacklist()

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
	loger.Infof("Loaded blacklist")
	return
}

func (cfg Config)CheckConn(ip net.IP, username string)(ok bool){
	wl := getWhitelist()
	if cfg.EnableIPWhitelist {
		for _, o := range wl.IPWhitelist {
			if matchIP(o, ip) {
				ok = true
				break
			}
		}
		if !ok {
			return
		}
	}
	if cfg.EnableWhitelist {
		_, uid, err := liter.DefaultAuthClient.GetPlayerUUID(username)
		if err != nil {
			loger.Errorf("Cannot get uuid of player '%s': %v", username, err)
			return false
		}
		uids := wl.UUIDWhitelist
		i := sort.Search(len(uids), func(i int)(bool){
			return !uuidLess(uids[i], uid)
		})
		if ok = i < len(uids) && uids[i] == uid; !ok {
			return
		}
	}
	blacklist
	return true
}

type ServerIns struct{
	Id          string
	Target      string
	ServerNames []string
	HandlePing  bool
	Motd        string
	MotdFailed  string
}

var servers = loadServerIns()

func getServers()([]*ServerIns){
	configLock.RLock()
	defer configLock.RUnlock()
	return servers
}

func loadServerIns()(svrs []*ServerIns){
	_ = _after_load
	dirpath := filepath.Join(configDir, "server-available")
	files, err := os.ReadDir(dirpath)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.Mkdir(dirpath, 0755); err != nil {
				loger.Fatalf("Cannot create config dir '%s': %v", dirpath, err)
			}
			return
		}
		loger.Fatalf("Error when loading server instances: %v", err)
	}
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".svr") {
			p := filepath.Join(dirpath, f.Name())
			data, err := os.ReadFile(p)
			if err != nil {
				loger.Errorf("Cannot load server instance at '%s': %v", p, err)
				continue
			}
			svr := new(ServerIns)
			sc := bufio.NewScanner(bytes.NewReader(data))
			for sc.Scan() {
				t := bytes.TrimSpace(sc.Bytes())
				if len(t) == 0 || t[0] == '#' {
					continue
				}
				k, v := split((string)(t), " \t")
				if v = strings.TrimSpace(v); len(v) == 0 {
					loger.Warnf("No value for key '%s' in '%s'", k, p)
					continue
				}
				switch strings.ToLower(k) {
				case "server_id":
					if len(svr.Id) != 0 {
						loger.Warnf("Already defined key 'server_id' in '%s'", p)
						continue
					}
					svr.Id = v
				case "server_name":
					svr.ServerNames = append(svr.ServerNames, v)
				case "target":
					svr.Target = v
				case "handle_ping":
					svr.HandlePing, _ = strconv.ParseBool(v)
				case "motd":
					svr.Motd = v
				case "motd_failed":
					svr.MotdFailed = v
				}
			}
			if len(svr.Id) == 0 {
				loger.Warnf("No id for server file at '%s'", p)
			}else{
				loger.Infof("Loaded server config id='%s'", svr.Id)
				svrs = append(svrs, svr)
			}
		}
	}
	return
}
