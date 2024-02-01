package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/renankcb/honest-backend-challenge/entities"
	"github.com/renankcb/honest-backend-challenge/loanevaluation"
)

func TestLoanEvaluator_Evaluate(t *testing.T) {
	tests := []struct {
		name     string
		strategy loanevaluation.LoanEvaluationStrategy
		applicant entities.Applicant
		expected string
	}{
		{
			name:     "StandardEvaluation",
			strategy: &loanevaluation.StandardEvaluationStrategy{},
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
			name:     "PreApprovedEvaluation",
			strategy: loanevaluation.NewPreApprovedEvaluationStrategy([]string{"0123456789"}),
			applicant: entities.Applicant{
				PhoneNumber: "0123456789",
			},
			expected: "approved",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluator := loanevaluation.NewLoanEvaluator(tt.strategy)
			result := evaluator.Evaluate(tt.applicant)
			assert.Equal(t, tt.expected, result, "unexpected result")
		})
	}
}
