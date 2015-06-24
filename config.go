package main

import (
	cfg "github.com/chalupaul/viper"
	log "github.com/Sirupsen/logrus"
	"time"
)

const CfgUrl string = "/julep/config"
const CfgProviderType string = "etcd"
const DefaultEtcdUrl string = "http://localhost:4001/"
const DefaultKeyFile string = "/etc/julep/.secring.gpg"

type config struct {
	hi string
}

func LoadCfg(url, key string) error {
	log.WithFields(log.Fields{
		"url": url,
		"key": key,
		"function": "LoadCfg",
	}).Debug("Loading config")
	if err := cfg.AddSecureRemoteProvider(CfgProviderType, url, CfgUrl, key); err != nil {
		log.WithFields(log.Fields{
			"function": "cfg.AddSecureRemoteProvider",
			"remote_provider": CfgProviderType,
			"url": url,
			"path": CfgUrl,
			"key": key,
		}).Fatal(err)
		return err
	}

	cfg.SetConfigType("json")
	if err := cfg.ReadRemoteConfig(); err != nil {
		log.WithFields(log.Fields{
			"function": "cfg.ReadRemoteConfig",
			"etcd_url": url,
			"path": CfgUrl,
			"key": key,
		}).Fatal(err)
		return err
	}
	
	// Read initial config and poll for changes
	var C config
	cfg.Marshal(&C)
	go func() {
		for {
			time.Sleep(time.Second * 5)
			if err := cfg.WatchRemoteConfig(); err != nil {
				log.Warn(err)
			} else {
				log.Warn("Config changed. Reloading.")
				cfg.Marshal(&C)
			}
		}
	}()
	
	return nil
}