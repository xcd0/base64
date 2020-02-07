package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	flag.Parse()

	input := flag.Arg(0)

	byteInput, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	output := ""
	lines := strings.Split(string(byteInput), "\n")

	for _, line := range lines {
		decode, _ := base64.StdEncoding.DecodeString(line)
		output += string(decode)
	}

	fmt.Println(output)

	return
}
