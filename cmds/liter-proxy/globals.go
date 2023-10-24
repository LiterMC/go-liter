
package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"

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

var configDir = "."
// var configDir = getConfigDir()

func getConfigDir()(dir string){
	if runtime.GOOS == "windows" {
		dir = "."
	}else{
		dir = filepath.Join("/opt", "liter_proxy")
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

type ProxyMapItem struct {
	Target      string `json:"target"`
	ForwardAddr string `json:"forward-addr,omitempty"`
	ForwardPort uint16 `json:"forward-port,omitempty"`
}

type Config struct {
	Debug bool `json:"debug"`

	ProxyURL   string `json:"proxy-url"`
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
			cfg.ProxyURL   = ""
			cfg.ServerIP   = ""
			cfg.ServerPort = 25565
			cfg.ProxyMap = map[string]ProxyMapItem{
				"hypixel.localhost": ProxyMapItem{
					Target: "mc.hypixel.net", // `Target` is the actually connect address
					// the `ForwardAddr` and `ForwardPort` will rewrite `HandshakePkt`
					ForwardAddr: "mc.hypixel.net",
					ForwardPort: 25565,
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

var AuthClient = liter.DefaultAuthClient
