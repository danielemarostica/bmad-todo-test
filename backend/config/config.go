package config

import "os"

type Config struct {
	Port       string
	MongoURI   string
	CORSOrigin string
}

func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "8080"),
		MongoURI:   getEnv("MONGO_URI", "mongodb://mongo:27017/todos"),
		CORSOrigin: getEnv("CORS_ORIGIN", "http://localhost:3000"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
