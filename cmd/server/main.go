package main

import (
	"google.golang.org/grpc/grpclog"
	"os"
)

func main() {
	err := run()
	if err != nil {
		grpclog.Errorf("server was shutdown with errors: %v", err)
		os.Exit(1)
	}
}
