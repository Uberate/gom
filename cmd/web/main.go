package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/uberate/gom/cmd/web/bc"
	"gopkg.in/yaml.v3"
	"os"
)

// main function will bootstrap the application
func main() {
	println("init done")
}

var configInstance *bc.ApplicationConfig

func init() {
	showDefaultConfig := false

	// init config info
	configPath := os.Getenv("OM_CONFIG_PATH")
	if len(configPath) == 0 {
		flag.BoolVar(&showDefaultConfig, "show-default-config", false, "--show-default-config to "+
			"show the default config, this flags will stop the process.")
		// If none env setting, parse flag.
		flag.StringVar(&configPath, "config", "./conf/web.conf.yaml", "--config config-path or set env OM_CONFIG_PATH")
		flag.Parse()
	}

	if showDefaultConfig {
		c := bc.DefaultConfig()
		configYamlValue, err := yaml.Marshal(c)
		if err != nil {
			panic(err)
		}
		fmt.Print(bc.DefaultConfigDescribePrefix)
		fmt.Println(string(configYamlValue))
		os.Exit(0)
	}

	if len(configPath) == 0 {
		// ignore conf info, use default config
		fmt.Println("load default config")
		c := bc.DefaultConfig()
		configInstance = &c
	} else {
		c, err := parseConfig(configPath)
		if err != nil {
			panic(err)
		}
		configInstance = c
	}

	// init logger info
	if err := bc.InitLogInstance(configInstance.Log); err != nil {
		panic(err)
	}
	bc.LoggerInstance.Trace("log init done")

	// print the version info
	versionJsonBytes, err := json.Marshal(GetVersionInfo())
	if err != nil {
		bc.LoggerInstance.Error(err)
		bc.LoggerInstance.Warn("skip version check")
	}

	bc.LoggerInstance.Info(string(versionJsonBytes))
}

func parseConfig(configPath string) (*bc.ApplicationConfig, error) {

	v := viper.NewWithOptions(viper.KeyDelimiter("_"))
	c := bc.DefaultConfig()

	// merge default config
	defaultConfig := map[string]interface{}{}
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "yaml",
		Result:  &defaultConfig,
	})
	if err != nil {
		return nil, err
	}
	if err := decoder.Decode(c); err != nil {
		return nil, err
	}
	if err := v.MergeConfigMap(defaultConfig); err != nil {
		return nil, err
	}

	// read from config file
	v.SetConfigFile(configPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	// sync config value from env
	v.AutomaticEnv()

	// unmarshal value to config instance
	if err := v.Unmarshal(&c, func(config *mapstructure.DecoderConfig) {
		config.TagName = "yaml"
	}); err != nil {
		return nil, err
	}

	return &c, nil
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
