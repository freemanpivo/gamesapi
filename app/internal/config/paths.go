package config

import (
	"log"
	"os"
)

func addConfigPath(v interface {
	AddConfigPath(string)
}) error {

	if _, err := os.Stat("./app/configs"); err == nil {
		log.Printf("[config] using config path: ./app/configs")
		v.AddConfigPath("./app/configs")
		return nil
	}

	if _, err := os.Stat("./configs"); err == nil {
		log.Printf("[config] using config path: ./configs")
		v.AddConfigPath("./configs")
		return nil
	}

	log.Printf("[config] no valid config path found")
	return os.ErrNotExist
}
