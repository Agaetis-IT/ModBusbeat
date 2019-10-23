package main

import (
	"modbusbeat/cmd"
	_ "modbusbeat/include"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
