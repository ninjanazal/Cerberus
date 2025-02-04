package config

import (
	"bufio"
	logger "cerberus/internal/tools"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ConfigData holds the configuration settings for the application.
// It includes the server port and the debug flag.
type ConfigData struct {
	ServerPort int
	Debug      bool
}

// DefaultCfg is the default configuration that is loaded at initialization.
var DefaultCfg ConfigData

func init() {
	DefaultCfg.ServerPort = 8080
	DefaultCfg.Debug = true
}

// ========================================
// region Public

func (m_config *ConfigData) GetAddressStr() string {
	return fmt.Sprintf(":%d", m_config.ServerPort)
}

// LoadEnvFile loads environment variables from the specified `.env` file.
// It parses the file, checks for valid key-value pairs, and returns a ConfigData struct
// with the values of the configuration settings.
//
// Parameters:
//   - path: The file path to the `.env` file.
//
// Returns:
//   - A pointer to a ConfigData struct populated with the values from the `.env` file.
//   - An error if there was an issue opening or reading the file.
func LoadEnvFile(p_path string) (*ConfigData, error) {
	if p_path == "" {
		logger.Log("Empty config file path, aborting", logger.WARN)
		return nil, errors.New("empty config file path")
	}

	file, err := os.Open(p_path)
	cfg := &ConfigData{}

	if err != nil {
		logger.Log(fmt.Sprintf("Error opening .env file: %v", err), logger.INFO)
		return nil, err

	} else {
		sc := bufio.NewScanner(file)
		for sc.Scan() {
			line := strings.TrimSpace(sc.Text())

			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}

			parts := strings.SplitN(line, "=", 2)
			if len(parts) != 2 {
				logger.Log(fmt.Sprintf("Skipping malformed line - %s", line), logger.INFO)

				continue
			}

			key := strings.TrimSpace(parts[0])
			value := strings.Trim(strings.TrimSpace(parts[1]), `"'`)

			// Set the corresponding field in the Config struct
			switch key {
			case "SERVER_PORT":
				v, err := strconv.Atoi(value)
				if err != nil {
					v = 8080
				}
				cfg.ServerPort = v

			case "DEBUG":
				// Convert string to bool
				cfg.Debug = value == "true"

			}
		}

		if err := sc.Err(); err != nil {
			logger.Log(fmt.Sprintf("Error reading .env file: %v", err), logger.INFO)
			return nil, err
		}
	}

	file.Close()
	return cfg, nil
}

// endregion Public
// ========================================
