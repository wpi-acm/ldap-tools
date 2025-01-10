package ldap

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/wpi-acm/ldap-tools/internal/config"
	"golang.org/x/term"
)

func SetupLdap(conf *config.LdapConfig) *ldap.Conn {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Bind Username: ")
	uname, _ := reader.ReadString('\n')
	fmt.Printf("Bind password: ")
	pw, err := term.ReadPassword(int(os.Stdin.Fd()))
	conn, err := ldap.DialURL(conf.Host)

	if err != nil {
		fmt.Printf("Error connecting: %s", err)
		os.Exit(-1)
	}

	err = conn.Bind(fmt.Sprintf("uid=%s,%s", strings.TrimSpace(uname), conf.UserSuffix), string(pw))
	if err != nil {
		fmt.Printf("Error binding: %s", err)
		os.Exit(-1)
	}
	fmt.Println()

	return conn

}
