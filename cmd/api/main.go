package main

import (
	"github.com/wen-flower/easy-douyin/pkg/command"
	"os"
)

func main() {
	cmd := command.NewCommand("api-service", func() error {
		return nil
	})
	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
