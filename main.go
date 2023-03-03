package main

import (
	"fmt"
	"os"
	"time"

	"github.com/pquerna/otp/totp"
)

func main() {
	// All we care about is the last argument which is the secret
	a := os.Args
	s := a[len(a)-1]

	// Generate and return the code
	c, err := totp.GenerateCode(s, time.Now())
	if err != nil {
		fmt.Println("Error generating code: " + err.Error())
		os.Exit(1)
	}
	fmt.Println(c)
}
