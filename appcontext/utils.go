package appcontext

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

func getIntOrPanic(key string) int {
	intValue, err := strconv.Atoi(getConfigStringValue(key))
	if err != nil {
		panicForkey(key, err)
	}
	return intValue
}

func getBoolOrPanic(key string) bool {
	boolValue, err := strconv.ParseBool(getConfigStringValue(key))
	if err != nil {
		panicForkey(key, err)
	}
	return boolValue
}

func getStringOrPanic(key string) string {
	value := getConfigStringValue(key)
	if value == "" {
		panicForkey(key, errors.New("config is not set"))
	}
	return value
}

func getConfigStringValue(key string) string {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		fmt.Printf("config %s is not set\n", key)
	}
	value := os.Getenv(key)
	if value == "" {
		value = viper.GetString(key)
	}
	return value
}

func panicForkey(key string, err error) {
	panic(fmt.Sprintf("Error %v occured while reading config %s", err.Error(), key))
}
