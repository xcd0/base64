package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
)

func main() {
	if terminal.IsTerminal(0) {
		fmt.Println("パイプ無し(FD値0)")
	} else {
		stdin, _ := ioutil.ReadAll(os.Stdin)
		encode := base64.StdEncoding.EncodeToString(stdin)
		output := string(encode)
		fmt.Println(output)
	}
}
