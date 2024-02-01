package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/renankcb/honest-backend-challenge/loanevaluation"
)

func TestIsAllowedAreaCode(t *testing.T) {
	tests := []struct {
		name        string
		phoneNumber string
		expected    bool
	}{
		{
			name:        "AllowedAreaCode",
			phoneNumber: "0123456789",
			expected:    true,
		},
		{
			name:        "AllowedAreaCode",
			phoneNumber: "2678901234",
			expected:    true,
		},
		{
			name:        "AllowedAreaCode",
			phoneNumber: "5123456789",
			expected:    true,
		},
		{
			name:        "AllowedAreaCode",
			phoneNumber: "8678901234",
			expected:    true,
		},
		{
			name:        "DisallowedAreaCode",
			phoneNumber: "1123456789",
			expected:    false,
		},
		{
			name:        "DisallowedAreaCode",
			phoneNumber: "3678901234",
			expected:    false,
		},
		{
			name:        "DisallowedAreaCode",
			phoneNumber: "4123456789",
			expected:    false,
		},
		{
			name:        "DisallowedAreaCode",
			phoneNumber: "6678901234",
			expected:    false,
		},
		{
			name:        "DisallowedAreaCode",
			phoneNumber: "7123456789",
			expected:    false,
		},
		{
			name:        "DisallowedAreaCode",
			phoneNumber: "9678901234",
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := loanevaluation.IsAllowedAreaCode(tt.phoneNumber)
			assert.Equal(t, tt.expected, result, "unexpected result")
		})
	}
}
