package util

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) (string, error) {
	envs, err := godotenv.Read(".env")
	if err != nil {
		env, ok := os.LookupEnv(key)
		if ok {
			return env, nil
		} else {
			return "", err
		}
	}

	return envs[key], nil
}
