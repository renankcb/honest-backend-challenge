package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/renankcb/honest-backend-challenge/entities"
	"github.com/renankcb/honest-backend-challenge/loanevaluation"
)

func TestPreApprovedEvaluationStrategy_Evaluate(t *testing.T) {
	preApprovedNumbers := []string{"0123456789", "9876543210"}
	strategy := loanevaluation.NewPreApprovedEvaluationStrategy(preApprovedNumbers)

	tests := []struct {
		name     string
		applicant entities.Applicant
		expected string
	}{
		{
			name: "ApprovedScenario",
			applicant: entities.Applicant{
				PhoneNumber: "0123456789",
			},
			expected: "approved",
		},
		{
			name: "DeclinedScenario",
			applicant: entities.Applicant{
				PhoneNumber: "1112233444",
			},
			expected: "declined",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := strategy.Evaluate(tt.applicant)
			assert.Equal(t, tt.expected, result, "unexpected result")
		})
	}
}
