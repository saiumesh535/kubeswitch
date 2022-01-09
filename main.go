package main

import (
	"log"

	"github.com/saiumesh535/kubectl-switch/cli"
)

func main() {
	if err := cli.RenderUI(); err != nil {
		log.Fatal(err.Error())
	}
}
