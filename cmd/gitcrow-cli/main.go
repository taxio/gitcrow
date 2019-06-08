package main

import (
	"fmt"
	"os"

	"github.com/taxio/gitcrow/cmd/gitcrow-cli/cmd"
	"go.uber.org/zap"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		zap.L().Fatal(fmt.Sprintf("%+v\n", err))
	}
	os.Exit(0)
}
