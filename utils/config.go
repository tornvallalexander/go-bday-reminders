package utils

import "github.com/spf13/viper"

// Config stores all necessary configuration for application
type Config struct {
	AccountSid    string `mapstructure:"ACCOUNT_SID"`
	AuthToken     string `mapstructure:"AUTH_TOKEN"`
	Receiver      string `mapstructure:"RECEIVER"`
	Sender        string `mapstructure:"SENDER"`
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig loads all environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
