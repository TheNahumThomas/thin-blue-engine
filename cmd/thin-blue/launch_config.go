package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	config_file string // Path to the configuration file
	debug_mode  bool
	test_mode   bool
	hot_reload  bool
	port        int
}

func (c *Config) getConfigFromFile(file string) error {

	os.Open(file)
	err := json.Unmarshal([]byte(file), &c)
	if err != nil {
		return fmt.Errorf("error reading config file %s: %w", file, err)
	}

	return nil

}

// Flags for the command line interface

func (c *Config) InitialConfig() error {

	flag.StringVar(&c.config_file, "config", "", "Name of a custom configuration file to be used by TBE all other flags will be ignored. Specified File must be present in the config/ directory.")
	flag.BoolVar(&c.debug_mode, "debug", false, "Enable debug mode. This will provide rule-tracing functionality. TBE will become more verbose and provide rule evaluation details.")
	flag.BoolVar(&c.test_mode, "test", false, "Enable built-in test mode. TBE will adopt a test ruleset and will not send alerts to the dispatcher. Can be configured to use custom test ruleset or simulated logs in the config file.")
	flag.BoolVar(&c.hot_reload, "hot reload", true, "Enable hot reloading of rules. TBE will automatically detect and load rule changes without restarting. This feature is enabled by default")
	flag.IntVar(&c.port, "port", 1470, "Specifies Port on which TBE will listen for incoming requests. Default is 1470.")

	flag.Parse()

	if c.config_file == "" {
		return nil
	}

	c.config_file = filepath.Join("config", c.config_file)
	_, err := os.Stat(c.config_file)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {

			return fmt.Errorf("TBE Couldn't find the config file at %s.\n These must be present in the config/ directory and match the specifcied name.\n Error: %w", c.config_file, err)

		} else if errors.Is(err, os.ErrPermission) {

			return fmt.Errorf("TBE does not have permission to access the config file at %s.\n Error: %w", c.config_file, err)

		} else {

			return fmt.Errorf("TBE encountered an error while trying to find the config file at %s. Please make sure it is a correctly formatted JSON file.\n Error: %w", c.config_file, err)

		}
	}

	err = c.getConfigFromFile(c.config_file)
	if err != nil {
		return fmt.Errorf("TBE encountered an error while trying to read the config file at %s. Please make sure it is a correctly formatted JSON file.\n Error: %w", c.config_file, err)
	}

	return nil

}
