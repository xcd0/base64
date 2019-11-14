package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {

	flag.Parse()

	input := flag.Arg(0)

	byteInput, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}

	encode := base64.StdEncoding.EncodeToString([]byte(byteInput))
	output := string(encode)

	fmt.Println(output)
}
