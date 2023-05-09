package bc

import (
	"github.com/sirupsen/logrus"
)

var LoggerInstance *logrus.Logger

// InitLogInstance will init the logger instance by config.
func InitLogInstance(config LogConfig) error {
	LoggerInstance = logrus.New()
	LoggerInstance.SetFormatter(&logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             config.DisableColor,
		EnvironmentOverrideColors: config.EnvironmentOverrideColors,
		DisableTimestamp:          config.DisableTimestamp,
		FullTimestamp:             config.FullTimestamp,
		TimestampFormat:           config.TimestampFormat,
	})

	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		return err
	}

	LoggerInstance.SetLevel(level)
	return nil
}
