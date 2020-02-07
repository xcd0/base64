package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if terminal.IsTerminal(0) == false {
		stdin, _ := ioutil.ReadAll(os.Stdin)
		output := ""
		lines := strings.Split(string(stdin), "\n")

		for _, line := range lines {
			decode, _ := base64.StdEncoding.DecodeString(line)
			output += string(decode)
		}
		fmt.Println(output)
	}
}
