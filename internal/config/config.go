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
		ServerPort:   getEnv("SERVER_PORT", "8080"),
		MongoURI:     getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:      getEnv("MONGO_DB", "clayjar"),
		KafkaBrokers: parseKafkaBrokers(getEnv("KAFKA_BROKERS", "localhost:9092")),
		KafkaTopic:   getEnv("KAFKA_TOPIC", "jar-events"),
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

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func parseKafkaBrokers(brokers string) []string {
	return strings.Split(brokers, ",")
}
