package config

import "github.com/spf13/viper"

type Config struct {
	EmailVerificationCodeLength uint8  `mapstructure:"EMAIL_VERIFICATION_CODE_LENGTH"`
	DBConnection                string `mapstructure:"DB_CONNECTION"`
	GRPCAddress                 string `mapstructure:"GRPC_SERVER_ADDRESS"`
	Environment                 string `mapstructure:"ENVIRONMENT""`
	ServiceName                 string `mapstructure:"SERVICE_NAME"`
	OtelCollectorAddr           string `mapstructure:"OTEL_COLLECTOR_STRUCTURE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return config, err
		}
	}

	// if file is just not found means we are running from docker
	viper.AutomaticEnv()
	viper.BindEnv("EMAIL_VERIFICATION_CODE_LENGTH")
	viper.BindEnv("DB_CONNECTION")
	viper.BindEnv("GRPC_SERVER_ADDRESS")
	viper.BindEnv("ENVIRONMENT")
	viper.BindEnv("SERVICE_NAME")
	viper.BindEnv("OTEL_COLLECTOR_STRUCTURE")
	err = viper.Unmarshal(&config)
	return config, err
}
