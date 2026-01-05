package config

import (
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	cfg   Config
	mutex sync.RWMutex
)

func AppSettings() (*Config, error) {
	v := viper.New()

	v.SetConfigName("config")
	v.SetConfigType("yaml")

	if err := addConfigPath(v); err != nil {
		return nil, err
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	startHotReload(v)

	return &cfg, nil
}

func startHotReload(v *viper.Viper) {
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config changed: %s", e.Name)

		var newCfg Config
		if err := v.Unmarshal(&newCfg); err != nil {
			log.Printf("failed to reload config: %v", err)
			return
		}

		mutex.Lock()
		cfg = newCfg
		mutex.Unlock()

		log.Println("configuration reloaded successfully")
	})
}

func Get() Config {
	mutex.RLock()
	defer mutex.RUnlock()
	return cfg
}
