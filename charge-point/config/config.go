package config

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var (
	Config *viper.Viper // Viper configuration object
	Env    string       // Environment name
)

func init() {
	// Get the environment name from environment variables
	Env = os.Getenv("ENV")
	if Env == "" {
		Env = "dev"
	}

	// Configure viper
	setupConfiguration()
}

// Set up configuration management
func setupConfiguration() {
	// Find the config file path
	_, filename, _, _ := runtime.Caller(1)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		panic(err)
	}
	dir = path.Join(dir, "..", "conf")

	// Set up the viper object
	Config = viper.New()
	Config.SetConfigName(Env)
	Config.SetConfigType("yaml")
	Config.AddConfigPath(dir)

	// Read the configuration
	err = Config.ReadInConfig()
	if err != nil {
		if err.Error() == "open : no such file or directory" {
			panic("Could not find the configuration file")
		}
		panic(err)
	}
}
