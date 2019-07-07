package cmd

import (
	"log"
	"strings"

	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// AddLoggerFlags adds support to configure the level of the logger
func AddLoggerFlags(c *cobra.Command) {
	c.PersistentFlags().String("log.level", "info", "one of debug, info, warn, error or fatal")
	c.PersistentFlags().String("log.format", "text", "one of text or json")
	c.PersistentFlags().Bool("log.line", false, "enable filename and line in logs")
}

// AddGlobalFlags adds the persistent flags that will be added to all the
// commands
func AddGlobalFlags(c *cobra.Command) {
	c.PersistentFlags().String("conf", "", "configuration file to use")
	c.PersistentFlags().StringP("client_id", "i", "", "Define the SoundCloud Client ID")
	c.PersistentFlags().Bool("debug", false, "Enable or disable debug mode")
	if err := viper.BindPFlags(c.PersistentFlags()); err != nil {
		log.Fatal("Could not bind flags")
	}
}

// Initialize will be run when cobra finishes its initialization
func Initialize() {
	// Environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("depadofm")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Configuration file
	if viper.GetString("conf") != "" {
		viper.SetConfigFile(viper.GetString("conf"))
	} else {
		viper.SetConfigName("conf")
		viper.AddConfigPath(".")
		viper.AddConfigPath("/config/")
	}
	hasconf := viper.ReadInConfig() == nil

	// Set log level
	lvl := viper.GetString("log.level")
	if l, err := logrus.ParseLevel(lvl); err != nil {
		logrus.WithFields(logrus.Fields{"level": lvl, "fallback": "info"}).Warn("Invalid log level")
	} else {
		logrus.SetLevel(l)
	}

	// Set log format
	switch viper.GetString("log.format") {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableTimestamp: true,
			ForceColors:      true,
		})
	}

	// Defines if logrus should display filenames and line where the log ocured
	if viper.GetBool("log.line") {
		logrus.AddHook(filename.NewHook())
	}

	// Delays the log for once the logger has been setup
	if hasconf {
		logrus.WithField("file", viper.ConfigFileUsed()).Debug("Found configuration file")
	} else {
		logrus.Debug("No configuration file found")
	}
}
