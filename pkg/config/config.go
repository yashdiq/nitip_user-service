package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	CASSANDRA_URL string `mapstructure:"CASSANDRA_URL"`
	CASSANDRA_PORT int `mapstructure:"CASSANDRA_PORT"`
	CASSANDRA_KEYSPACE string `mapstructure:"CASSANDRA_KEYSPACE"`
	CASSANDRA_USER string `mapstructure:"CASSANDRA_USER"`
	CASSANDRA_PASS string `mapstructure:"CASSANDRA_PASS"`
}

var envs =[]string{
	"CASSANDRA_URL", "CASSANDRA_PORT", "CASSANDRA_KEYSPACE", "CASSANDRA_USER", "CASSANDRA_PASS",
}

func LoadConfig() (Config, error) {
	var cfg Config
	viper.AddConfigPath("/")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return cfg, err
		}
	}
	cfgerr := viper.Unmarshal(&cfg)

	LoadEnv()

	return cfg, cfgerr
}

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(".env not loaded, error!")
	}
}