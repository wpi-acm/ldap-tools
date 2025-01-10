package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/wpi-acm/ldap-tools/internal/config"
	acmldap "github.com/wpi-acm/ldap-tools/internal/ldap"
)

func main() {
	conf, err := config.ReadConfig("./config.toml")
	if err != nil {
		log.Fatalf("Couldn't read config: %s", err)
	}

	fmt.Println("LDAP Host: ", conf.Host)

	conn := acmldap.SetupLdap(&conf)
	var uid string
	if len(os.Args) == 2 {
		uid = os.Args[1]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter UID: ")
		uid, _ = reader.ReadString('\n')
		fmt.Println()
	}

	uid = strings.TrimSpace(uid)

	search := ldap.NewSearchRequest(conf.BaseDN, 2, 1, 10, 10, false, fmt.Sprintf("(uid=%s)", uid), nil, nil)

	res, err := conn.Search(search)

	if err != nil {
		fmt.Printf("Search error: %s\n", err)
		os.Exit(-1)
	}

	for i := 0; i < len(res.Entries); i++ {
		entry := res.Entries[i]
		fmt.Printf("--------------\n")
		fmt.Printf("DN: %s\n", entry.DN)
		for j := 0; j < len(entry.Attributes); j++ {
			attr := entry.Attributes[j]
			fmt.Printf("%s\n", attr.Name)
			for k := 0; k < len(attr.Values); k++ {
				fmt.Printf("\t%s\n", attr.Values[k])
			}
		}
	}
}
