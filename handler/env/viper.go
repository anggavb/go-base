package load_config

import (
	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
//
// How to use:
//   - Create a struct that represents your configuration.
//   - Call LoadConfig with the path to your .env file and the struct type.
//   - The function will return the populated struct and any error encountered.
//
// Example:
//
//	type Config struct {
//	    Port int `mapstructure:"PORT"`
//	    DBHost string `mapstructure:"DB_HOST"`
//	}
//	config, err := env_config.LoadConfig[Config](".", ".env", "env")
func UseViper[T any](path, configFile, configType string) (config T, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(configFile)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
