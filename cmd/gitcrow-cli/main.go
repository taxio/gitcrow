package main

import (
	"log"
	"os"

	"github.com/taxio/gitcrow/cmd/gitcrow-cli/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	os.Exit(0)
}
