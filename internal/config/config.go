package config

import (
	"errors"
	"os"
	"strconv"
)

const (
	FireflyAppTimeout = 10               // 10 sec for fftc to app service timeout
	ModelFile         = "data/model.gob" //file name to store model
)

type Config struct {
	APIKey string
	FFApp  string
	Port   int
}

var envVars = []string{
	"FF_API_KEY",
	"FF_APP_URL",
	"PORT",
}

func EnvVarExist(varName string) bool {
	_, present := os.LookupEnv(varName)
	return present
}

func NewConfig() (*Config, error) {
	for _, val := range envVars {
		exist := EnvVarExist(val)
		if !exist {
			return nil, errors.New("env variable is not set: " + val)
		}
	}

	port, err := getenvInt("PORT")
	if err != nil {
	        return 0, err
        }
	
	cfg := Config{
		APIKey: os.Getenv("FF_API_KEY"),
		FFApp:  os.Getenv("FF_APP_URL"),
		Port:   port,
	}

	return &cfg, nil
}

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func getenvStr(key string) (string, error) {
    v := os.Getenv(key)
    if v == "" {
        return v, ErrEnvVarEmpty
    }
    return v, nil
}

func getenvInt(key string) (int, error) {
    s, err := getenvStr(key)
    if err != nil {
        return 0, err
    }
    v, err := strconv.Atoi(s)
    if err != nil {
        return 0, err
    }
    return v, nil
}
