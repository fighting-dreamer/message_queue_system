package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigurationIsSet(t *testing.T) {
	LoadTestConfig()
	expectedAppPort := 8080
	expectedLogLevel := "info"
	expectedAppName := "message_queue"

	assert.Equal(t, expectedAppPort, AppPort())
	assert.Equal(t, expectedAppName, AppName())
	assert.Equal(t, expectedLogLevel, LogLevel())
}
