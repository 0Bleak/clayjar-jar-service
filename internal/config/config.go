package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	ServerPort   string
	MongoURI     string
	MongoDB      string
	KafkaBrokers []string
	KafkaTopic   string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		ServerPort:   mustGetEnv("SERVER_PORT"),
		MongoURI:     mustGetEnv("MONGO_URI"),
		MongoDB:      mustGetEnv("MONGO_DB"),
		KafkaBrokers: parseKafkaBrokers(mustGetEnv("KAFKA_BROKERS")),
		KafkaTopic:   mustGetEnv("KAFKA_TOPIC"),
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.ServerPort == "" {
		return fmt.Errorf("SERVER_PORT is required")
	}
	if c.MongoURI == "" {
		return fmt.Errorf("MONGO_URI is required")
	}
	if c.MongoDB == "" {
		return fmt.Errorf("MONGO_DB is required")
	}
	if len(c.KafkaBrokers) == 0 {
		return fmt.Errorf("KAFKA_BROKERS is required")
	}
	if c.KafkaTopic == "" {
		return fmt.Errorf("KAFKA_TOPIC is required")
	}
	return nil
}

// mustGetEnv retrieves an environment variable or panics if not found
func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Required environment variable %s is not set", key))
	}
	return value
}

func parseKafkaBrokers(brokers string) []string {
	return strings.Split(brokers, ",")
}
