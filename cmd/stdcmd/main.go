package main

import (
	"fmt"
	"github.com/uberate/gom/cmd/stdcmd/cmds"
	"os"
)

func main() {
	cmd := cmds.RootCmd()
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
}
