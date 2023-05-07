package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/uberate/gom/cmd/stdcmd/cmds"
	"os"
)

func main() {
	cmd := cmds.RootCmd()
	cmd.AddCommand(VersionCmd())
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
}

// ----------- version flags.

// build with flag:
// -ldflags "-w -s -X 'main.Version=${VERSION}' -X 'main.HashTag=`git rev-parse HEAD`' -X 'main.BranchName=`git rev-parse --abbrev-ref HEAD`' -X 'main.BuildDate=`date -u '+%Y-%m-%d_%I:%M:%S%p'`' -X 'main.GoVersion=`go version`'"

const (
	VersionTagVersion    = "version"
	VersionTagHashTag    = "hash-tag"
	VersionTagBranchName = "branch-name"
	VersionTagBuildDate  = "build-date"
	VersionTagGoVersion  = "go-version"
)

var Version string
var HashTag string
var BranchName string
var BuildDate string
var GoVersion string

func GetVersionInfo() map[string]string {
	return map[string]string{
		VersionTagVersion:    Version,
		VersionTagHashTag:    HashTag,
		VersionTagBranchName: BranchName,
		VersionTagBuildDate:  BuildDate,
		VersionTagGoVersion:  GoVersion,
	}
}

func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "show version of 'GOM'",
		RunE: func(cmd *cobra.Command, args []string) error {
			v := GetVersionInfo()
			r, e := json.Marshal(v)
			if e != nil {
				err := fmt.Errorf("got version error: %v", e)
				cmd.PrintErr(err)
				return err
			}

			cmd.Println(string(r))
			return nil
		},
	}

	return cmd
}
