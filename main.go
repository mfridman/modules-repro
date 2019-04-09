package main

import (
	"fmt"
	"log"

	_ "cloud.google.com/go"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func main() {
	if err := yourFunction(); err != nil {
		log.Fatalln(err)
	}
	// yay!
	node := &yaml.Node{}
	fmt.Println(node)
}

func yourFunction() error {
	return errors.New("oh boi")
}
