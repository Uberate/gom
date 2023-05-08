package mian

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// main function will bootstrap the application
func main() {

}

func init() {

	// init config info
	configPath := os.Getenv("OM_CONFIG_PATH")
	if len(configPath) == 0 {

		// If none env setting, parse flag.
		flag.StringVar(&configPath, "c", "./conf/config.yaml", "-c config-path or set env OM_CONFIG_PATH")
		flag.Parse()
	}

	if len(configPath) == 0 {

		// ignore conf info, use default config
		fmt.Println("load default config")
	} else {
		parseConfig(configPath)
	}

	// init logger info

	// print the version info
}

func parseConfig(configPath string) {

	v := viper.New()

	// trans A_B-C to A.B-C
	v.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

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
