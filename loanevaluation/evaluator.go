package loanevaluation

import "github.com/renankcb/honest-backend-challenge/entities"

type LoanEvaluationStrategy interface {
	Evaluate(applicant entities.Applicant) string
}

type LoanEvaluator struct {
	strategy LoanEvaluationStrategy
}

func NewLoanEvaluator(strategy LoanEvaluationStrategy) *LoanEvaluator {
	return &LoanEvaluator{strategy: strategy}
}

func (e *LoanEvaluator) SetStrategy(strategy LoanEvaluationStrategy) {
	e.strategy = strategy
}

func (e *LoanEvaluator) Evaluate(applicant entities.Applicant) string {
	return e.strategy.Evaluate(applicant)
}
