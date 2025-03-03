package config

import "github.com/spf13/viper"

type Config struct {
	DBConnection string `mapstructure:"DB_CONNECTION"`
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
	viper.BindEnv("DB_CONNECTION")
	err = viper.Unmarshal(&config)
	return config, err
}
