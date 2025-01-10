package config

import (
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type LdapConfig struct {
	Host       string
	BaseDN     string
	UserSuffix string
}

func ReadConfig(filename string) (LdapConfig, error) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("Fatal error: %s", err)
		return LdapConfig{}, err
	}

	defer file.Close()

	var contents []byte
	_, err = file.Read(contents)
	if err != nil {
		log.Fatalf("Couldn't read file: %s", err)
		return LdapConfig{}, err
	}

	var conf LdapConfig
	err = toml.Unmarshal(contents, conf)

	if err != nil {
		log.Fatalf("Invalid TOML configuration: %s", err)
		return LdapConfig{}, err
	}

	return conf, nil
}
