package main

import (
	cfg "github.com/chalupaul/viper"
	log "github.com/Sirupsen/logrus"
)

const CfgUrl string = "/ginos/config.json"
const CfgProviderType string = "etcd"
const DefaultEtcdUrl string = "http://localhost:4001/"
const DefaultKeyFile string = "$HOME/ginos/.secring.gpg"

func LoadCfg(url, key string) error {
	if err := cfg.AddSecureRemoteProvider(CfgProviderType, url, CfgUrl, key); err != nil {
		log.WithFields(log.Fields{
			"function": "cfg.AddSecureRemoteProvider",
			"remote_provider": "etcd",
			"url": url,
			"path": "/ginos/config.json",
			"key": key,
		}).Fatal(err)
		return err
	}

	cfg.SetConfigType("json")
	if err := cfg.ReadRemoteConfig(); err != nil {
		log.WithFields(log.Fields{
			"function": "cfg.ReadRemoteConfig",
			"etcd_url": url,
			"path": "/ginos/config.json",
			"key": key,
		}).Fatal(err)
		return err
	}
	return nil
}