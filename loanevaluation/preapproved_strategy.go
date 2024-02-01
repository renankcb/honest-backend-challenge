// loanevaluation/preapproved_strategy.go
package loanevaluation

import "github.com/renankcb/honest-backend-challenge/entities"

type PreApprovedEvaluationStrategy struct {
	preApprovedNumbers map[string]bool
}

func NewPreApprovedEvaluationStrategy(numbers []string) *PreApprovedEvaluationStrategy {
	preApprovedNumbers := make(map[string]bool)
	for _, num := range numbers {
		preApprovedNumbers[num] = true
	}
	return &PreApprovedEvaluationStrategy{preApprovedNumbers: preApprovedNumbers}
}

func (s *PreApprovedEvaluationStrategy) Evaluate(applicant entities.Applicant) string {
	if s.preApprovedNumbers[applicant.PhoneNumber] {
		return "approved"
	}
	return "declined"
}
