// loanevaluation/standard_strategy.go
package loanevaluation

import (
	"github.com/renankcb/honest-backend-challenge/entities"
	"github.com/renankcb/honest-backend-challenge/risk"
)

type StandardEvaluationStrategy struct{}

func (s *StandardEvaluationStrategy) Evaluate(applicant entities.Applicant) string {
	creditRisk := risk.CalculateCreditRisk(applicant.Age, applicant.NumberOfCreditCards)

	if applicant.Income > 100000 &&
		applicant.Age >= 18 &&
		applicant.NumberOfCreditCards <= 3 &&
		creditRisk == "LOW" &&
		!applicant.PoliticallyExposed &&
		IsAllowedAreaCode(applicant.PhoneNumber) {
		return "approved"
	}

	return "declined"
}
