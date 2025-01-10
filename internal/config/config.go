package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type LdapConfig struct {
	Host       string `toml:"host"`
	BaseDN     string `toml:"base_dn"`
	UserSuffix string `toml:"user_suffix"`
}

func ReadConfig(filename string) (LdapConfig, error) {
	var conf LdapConfig
	_, err := toml.DecodeFile(filename, &conf)

	if err != nil {
		log.Fatalf("Invalid TOML configuration: %s", err)
		return LdapConfig{}, err
	}

	return conf, nil
}
