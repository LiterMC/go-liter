
package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
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

type ProxyMapItem struct {
	Target      string `yaml:"target"`
	ForwardAddr string `yaml:"forward-addr,omitempty"`
	ForwardPort uint16 `yaml:"forward-port,omitempty"`
}

type Config struct {
	Debug bool `yaml:"debug"`

	ProxyURL   string `yaml:"proxy-url"`
	ServerIP   string `yaml:"server-ip"`
	ServerPort uint16 `yaml:"server-port"`

	ProxyMap map[string]ProxyMapItem `yaml:"proxy-items,omitempty"`
}

var cfg = loadConfig()

func loadConfig()(cfg Config){
	// set config default values
	cfg.ProxyURL   = ""
	cfg.ServerIP   = "127.0.0.1"
	cfg.ServerPort = 25565
	cfg.ProxyMap = map[string]ProxyMapItem{
		"hypixel.localhost": ProxyMapItem{
			Target: "mc.hypixel.net", // `Target` is the actually connect address
			// the `ForwardAddr` and `ForwardPort` will rewrite `HandshakePkt`
			ForwardAddr: "mc.hypixel.net",
			ForwardPort: 25565,
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
		loger.Fatalf("Cannot save config file: %v", err)
		return
	}
	loger.Infof("Loaded config file")

	if cfg.Debug {
		loger.SetLevel(logger.TraceLevel)
		loger.Debug("Debug log enabled")
	}else{
		loger.SetLevel(logger.InfoLevel)
	}
	return
}

var AuthClient = liter.DefaultAuthClient
