package common

import (
	"errors"
	"os"
)

func GetRequiredEnvVariable(name string) string {
	value := os.Getenv(name)
	if value == "" {
		panic(errors.New("Missing environment variable " + name))
	} else {
		return value
	}
	return ""
}
