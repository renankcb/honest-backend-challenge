package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/renankcb/honest-backend-challenge/entities"
	"github.com/renankcb/honest-backend-challenge/loanevaluation"
)

func TestStandardEvaluationStrategy_Evaluate(t *testing.T) {
	tests := []struct {
		name     string
		applicant entities.Applicant
		expected string
	}{
		{
			name: "ApprovedScenario",
			applicant: entities.Applicant{
				Income:              120000,
				NumberOfCreditCards: 2,
				Age:                 25,
				PoliticallyExposed:  false,
				JobIndustryCode:     "XYZ",
				PhoneNumber:         "0123456789",
			},
			expected: "approved",
		},
		{
			name: "DeclinedScenario",
			applicant: entities.Applicant{
				Income:              80000,
				NumberOfCreditCards: 4,
				Age:                 17,
				PoliticallyExposed:  true,
				JobIndustryCode:     "ABC",
				PhoneNumber:         "5678901234",
			},
			expected: "declined",
		},
	}

	strategy := &loanevaluation.StandardEvaluationStrategy{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := strategy.Evaluate(tt.applicant)
			assert.Equal(t, tt.expected, result, "unexpected result")
		})
	}
}
