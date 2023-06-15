package config

import "os"

const (
	ApiKeyEnvVar = "LINEAR_API_KEY"
)

func ApiKey() string {
	return os.Getenv(ApiKeyEnvVar)
}
