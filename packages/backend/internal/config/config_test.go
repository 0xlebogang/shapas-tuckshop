package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ptr(s string) *string {
	return &s
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		value         *string
		fallbackValue string
		expectedValue string
	}{
		{
			name:          "Should return the correct value of the set environment variable",
			key:           "TEST_KEY",
			value:         ptr("test-value"),
			fallbackValue: "fallback-value",
			expectedValue: "test-value",
		},
		{
			name:          "Should return the fallback value when environment variable is not set",
			key:           "UNSET_TEST_KEY",
			value:         nil,
			fallbackValue: "fallback-value",
			expectedValue: "fallback-value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != nil {
				if err := os.Setenv(tt.key, *tt.value); err != nil {
					t.Error("Failed to set the variable")
				}
			}

			result := getEnv(tt.key, tt.fallbackValue)

			assert.Equal(t, tt.expectedValue, result)
		})
	}
}
