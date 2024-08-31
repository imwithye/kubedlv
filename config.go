package main

import (
	"os"
	"strconv"
)

type DlvConfig struct {
	APIVersion int
	Port       int
	Continue   bool
	Log        bool
}

func NewDlvConfig() *DlvConfig {
	return &DlvConfig{
		APIVersion: GetEnvInt("DLV_API_VERSION", 2),
		Port:       GetEnvInt("DLV_PORT", 2345),
		Continue:   GetEnvBool("DLV_CONTINUE", true),
		Log:        GetEnvBool("DLV_ENABLE_LOGGING", false),
	}
}

func GetEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return fallback
}

func GetEnvBool(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if i, err := strconv.ParseBool(value); err == nil {
			return i
		}
	}
	return fallback
}
