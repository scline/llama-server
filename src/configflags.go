package main

import (
	"flag"
	"os"
	"strconv"
)

// Struct for command line arguments
type Configflags struct {
	Port  int  // env: APP_PORT
	Debug bool // env: APP_DEBUG
}

// Struct for command line argument default values
type DefaultConfigflags struct {
	Port  int  // env: APP_PORT
	Debug bool // env: APP_DEBUG
}

// Function to load command line arguments for webserver and program configurations
func ParseConfigFromCommandLine() Configflags {
	var cfg Configflags
	var cfgDefaults DefaultConfigflags

	// If variables are defined within the environment, use them as defaults
	if value, present := os.LookupEnv("APP_PORT"); present {
		valueInt, err := strconv.Atoi(value)
		if err != nil {
			errorLogger.Fatalf("'APP_PORT' environment variable cannot be converted to 'int', value: '%s'", value)
		}
		cfgDefaults.Port = valueInt
	} else {
		cfgDefaults.Port = 8080
	}

	// As long as 'APP_DEBUG' is defined, we are going to assume we want to set as 'true'
	if _, present := os.LookupEnv("APP_DEBUG"); present {
		cfgDefaults.Debug = true
	} else {
		cfgDefaults.Debug = false
	}

	flag.IntVar(&cfg.Port, "port", cfgDefaults.Port, "Web server listening port")
	flag.BoolVar(&cfg.Debug, "debug", cfgDefaults.Debug, "Enable debug logging")

	flag.Parse()

	// If debug is enabled, dump configurations to logger
	if cfg.Debug {
		debugLogger.Printf("Imported command-line flags: %+v \n", cfg)
		debugLogger.Printf("Default command-line flags: %+v \n", cfgDefaults)
	}

	// return configuration
	return cfg
}
