package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("provide a json file path. usage path/to/json")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	buffer := []byte{}
	for scanner.Scan() {
		buffer = append(buffer, scanner.Bytes()...)
	}
}
