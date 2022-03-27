package utils

import (
	"github.com/spf13/viper"
	"time"
)

// Config stores all necessary configuration for application
type Config struct {
	TwilioAccountSid     string        `mapstructure:"TWILIO_ACCOUNT_SID"`
	TwilioAuthToken      string        `mapstructure:"TWILIO_AUTH_TOKEN"`
	TwilioReceiver       string        `mapstructure:"TWILIO_RECEIVER"`
	TwilioSender         string        `mapstructure:"TWILIO_SENDER"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
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
