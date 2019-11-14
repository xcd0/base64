package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {

	flag.Parse()

	input := flag.Arg(0)

	inputFilePath, _ := filepath.Abs(input)
	inputFileDir := filepath.Dir(inputFilePath)
	inputFileNameWithExt := filepath.Base(inputFilePath)
	inputFileExt := filepath.Ext(inputFilePath)
	inputFileName := inputFileNameWithExt[:len(inputFileNameWithExt)-len(inputFileExt)]
	outputDecodeFileName := filepath.Join(inputFileDir, inputFileName+"_decoded"+inputFileExt)
	outputEncodeFileName := filepath.Join(inputFileDir, inputFileName+"_encoded"+inputFileExt)

	//fmt.Println(inputFileName)
	//fmt.Println(outputFileName)

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

	if ioutil.WriteFile(outputDecodeFileName, []byte(output), 0666) != nil {
		panic(err)
	}

	encode := base64.StdEncoding.EncodeToString([]byte(byteInput))
	output = string(encode)

	if ioutil.WriteFile(outputEncodeFileName, []byte(output), 0666) != nil {
		panic(err)
	}

	return
}
