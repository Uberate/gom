package bc

import "time"

var DefaultConfigDescribePrefix = `# ----------------------------------
# File name: default-web.config.yaml
#
# This file is default config of GOM
#

`

// DefaultConfig generate a default ApplicationConfig.
func DefaultConfig() ApplicationConfig {
	return ApplicationConfig{
		Log: LogConfig{
			Level:                     "INFO",
			DisableColor:              false,
			EnvironmentOverrideColors: false,
			DisableTimestamp:          false,
			FullTimestamp:             true,
			TimestampFormat:           time.RFC3339Nano,
		},
		Web: WebConfig{
			Host:       "0.0.0.0",
			ListenPort: "3000",
			DebugMode:  false,
		},
	}
}

// ApplicationConfig define the config of application.
type ApplicationConfig struct {
	Log LogConfig `json:"log" yaml:"log" mapstructure:"log"`
	Web WebConfig `json:"web" yaml:"web"`
}

// LogConfig about the config of log for app.
type LogConfig struct {
	Level string `json:"level" yaml:"level"`

	// DisableColor will disable color of output. Default false.
	// If not disable color, will auto check tty, if colorable, output will with color.
	DisableColor bool `json:"disable-color" yaml:"disable-color"`

	// Override coloring based on CLICOLOR and CLICOLOR_FORCE. - https://bixense.com/clicolors/
	EnvironmentOverrideColors bool `json:"environment-override-colors,omitempty" yaml:"environment-override-colors"`

	// Disable timestamp logging. useful when output is redirected to logging
	// system that already adds timestamps.
	DisableTimestamp bool `json:"disable-timestamp,omitempty" yaml:"disable-timestamp"`

	// Enable logging the full timestamp when a TTY is attached instead of just
	// the time passed since beginning of execution.
	FullTimestamp bool `json:"full-timestamp,omitempty" yaml:"full-timestamp"`

	// TimestampFormat to use for display when a full timestamp is printed.
	// The format to use is the same than for time.Format or time.Parse from the standard
	// library.
	// The standard Library already provides a set of predefined format.
	TimestampFormat string `json:"timestamp-format,omitempty" yaml:"timestamp-format"`
}

type WebConfig struct {
	Host       string `json:"host,omitempty" yaml:"host"`
	ListenPort string `json:"listen-port,omitempty" yaml:"listen-port"`
	DebugMode  bool   `json:"debug-mode,omitempty" yaml:"debug-mode"`
}
