package main

import (
	"encoding/base32"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	// All we care about is the last argument which is the secret
	a := os.Args
	s := a[len(a)-1]
	s = strings.ToUpper(s)

	_, err := base32.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println("You must specify a valid base32 string")
		os.Exit(1)
	}

	// Generate and return the code
	c, err := totp.GenerateCode(s, time.Now())
	if err != nil {
		fmt.Println("Error generating code: " + err.Error())
		os.Exit(1)
	}
	fmt.Println(c)
}
