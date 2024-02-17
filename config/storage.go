package config

import (
	"sync"
)

type ConfigStorage struct {
	Config
}

var lock = &sync.Mutex{}

var configStorageInstance *ConfigStorage

func SetConfig(conf Config) {
	if configStorageInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if configStorageInstance == nil {
			configStorageInstance = &ConfigStorage{
				Config: conf,
			}
		}
	}
}

func GetConfigInstance() *ConfigStorage {
	return configStorageInstance
}
