package config

import "os"

type LdapConfig struct {
	Host       string
	BaseDN     string
	UserSuffix string
}

func ReadConfig(filename string) (LdapConfig, err) {
	file, err := os.Open(filename)

	if err != nil {

	}
}
