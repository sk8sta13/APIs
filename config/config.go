package config

import (
	"github.com/spf13/viper"
)

type db struct {
	Drive string `mapstructure:"DB_DRIVE"`
	Host  string `mapstructure:"DB_HOST"`
	User  string `mapstructure:"DB_USER"`
	Pass  string `mapstructure:"DB_PASS"`
	Port  string `mapstructure:"DB_PORT"`
	Name  string `mapstructure:"DB_NAME"`
}

func LoadConfigs(path string) (*db, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg *db
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return cfg, nil
}
