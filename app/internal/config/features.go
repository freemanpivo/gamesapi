package config

import (
	"log"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	features   FeatureFlags
	featuresMu sync.RWMutex
)

func AppFeatures() (*FeatureFlags, error) {
	v := viper.New()

	v.SetConfigName("features")
	v.SetConfigType("yaml")

	if err := addConfigPath(v); err != nil {
		return nil, err
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	loaded := FeatureFlags{
		Features: map[string]bool{},
		Limits:   map[string]int{},
	}
	if err := v.Unmarshal(&loaded); err != nil {
		return nil, err
	}

	featuresMu.Lock()
	features = loaded
	featuresMu.Unlock()

	startFeaturesHotReload(v)

	return &features, nil
}

func startFeaturesHotReload(v *viper.Viper) {
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("features config changed: %s", e.Name)

		var updated FeatureFlags
		if err := v.Unmarshal(&updated); err != nil {
			log.Printf("failed to reload features: %v", err)
			return
		}

		featuresMu.Lock()
		features = updated
		featuresMu.Unlock()

		log.Println("feature toggles reloaded")
	})
}

func IsEnabled(name string, defaultValue bool) bool {
	featuresMu.RLock()
	defer featuresMu.RUnlock()

	if features.Features == nil {
		return defaultValue
	}

	val, exists := features.Features[name]
	if !exists {
		return defaultValue
	}

	return val
}

func Limit(name string, defaultValue int) int {
	featuresMu.RLock()
	defer featuresMu.RUnlock()

	if features.Limits == nil {
		return defaultValue
	}

	val, exists := features.Limits[strings.ToLower(name)]
	if !exists {
		return defaultValue
	}

	return val
}
