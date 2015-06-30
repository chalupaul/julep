package main

import (
	"github.com/chalupaul/viper"
	log "github.com/Sirupsen/logrus"
	"time"
	"reflect"
	"errors"
	"fmt"
)

const DefaultCfgUrl string = "/julep/config"
const DefaultTreeUrl string = "/julep/hostgroups.json"
const CfgProviderType string = "etcd"
const DefaultEtcdUrl string = "http://localhost:4001/"
const DefaultKeyFile string = "/etc/julep/.secring.gpg"

type ConfigOption struct {
	Url string
	Key string
	CfgUrl string
	ProviderType string
	ContentType string
}

// CfgOpt is a wrapper function to set options for LoadCfg()
type CfgOpt func(*ConfigOption) error

// LoadCfg builds a config object from etcd values. It takes a url to an etcd server
// and a keyfile to decrypt the contents. It takes the rest of the values from 
// constants declared in this file. It will return any error that is bubbled up.
func LoadCfg(options ...CfgOpt) (*viper.Viper, error) {
	// Set defaults
	co := &ConfigOption{}
	co.Url = DefaultEtcdUrl
	co.Key = DefaultKeyFile
	co.CfgUrl = DefaultCfgUrl
	co.ProviderType = "etcd"
	co.ContentType = "json"
	
	// Override defaults
	for _, op := range options {
		err := op(co)
		if err != nil {
			return nil, err
		}
	}
	
	// Make the connections
	cfg := viper.New()
	log.WithFields(log.Fields{
		"url": co.Url,
		"key": co.Key,
		"function": "LoadCfg",
	}).Debug("Loading config")
	if err := cfg.AddSecureRemoteProvider(co.ProviderType, co.Url, co.CfgUrl, co.Key); err != nil {
		log.WithFields(log.Fields{
			"function": "cfg.AddSecureRemoteProvider",
			"remote_provider": co.ProviderType,
			"url": co.Url,
			"path": co.CfgUrl,
			"key": co.Key,
		}).Fatal(err)
		return nil, err
	}

	// Read config from etcd
	cfg.SetConfigType(co.ContentType)
	if err := cfg.ReadRemoteConfig(); err != nil {
		log.WithFields(log.Fields{
			"function": "cfg.ReadRemoteConfig",
			"etcd_url": co.Url,
			"path": co.CfgUrl,
			"key": co.Key,
		}).Fatal(err)
		return nil, err
	}
	
	// poll for changes
	go func() {
		for {
			time.Sleep(time.Second * 5)
			if err := cfg.WatchRemoteConfig(); err != nil {
				log.Warn(err)
			} else {
				log.Warn("Config changed. Reloading.")
			}
		}
	}()
	
	return cfg, nil
}
// OptionGeneric takes a string for value and name, then reflects the corresponding
// attribute from the ConfigOption struct and sets the variable. This should not be
// used directly.
func OptionGeneric(o, n string) func (c *ConfigOption) error {
	return func(c *ConfigOption) error {
		f := reflect.ValueOf(c).Elem().FieldByName(n)
		if f.IsValid() && f.CanSet() {
			f.SetString(o)
			return nil
		} else {
			return errors.New(fmt.Sprintf("Failed to set config field %s: Invalid Option.", n))
		}
		reflect.ValueOf(c).Elem().FieldByName(n).SetString(o)
		return nil
	}
}

func OptionUrl(o string) func(c *ConfigOption) error {
	return OptionGeneric(o, "Url")
}
func OptionKey(o string) func(c *ConfigOption) error {
	return OptionGeneric(o, "Key")
}
func OptionCfgUrl(o string) func(c *ConfigOption) error {
	return OptionGeneric(o, "CfgUrl")
}
func OptionProviderType(o string) func(c *ConfigOption) error {
	return OptionGeneric(o, "ProviderType")
}
