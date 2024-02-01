// loanevaluation/process_data.go
package loanevaluation

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/renankcb/honest-backend-challenge/entities"
)

func ProcessData(resp http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		processPostRequest(resp, req)
	default:
		handleNotFound(resp)
	}
}

func processPostRequest(resp http.ResponseWriter, req *http.Request) {
	var applicant entities.Applicant
	err := json.NewDecoder(req.Body).Decode(&applicant)
	if err != nil {
		handleError(resp, err, http.StatusBadRequest, "Bad Request")
		return
	}

	strategy := &StandardEvaluationStrategy{}
	loanEvaluator := NewLoanEvaluator(strategy)
	status := loanEvaluator.Evaluate(applicant)

	response := entities.ApplicationStatus{Status: status}
	respondWithJSON(resp, response)
}

func handleError(resp http.ResponseWriter, err error, statusCode int, message string) {
	log.Println("Error:", err)
	http.Error(resp, message, statusCode)
}

func handleNotFound(resp http.ResponseWriter) {
	log.Println("Error 404: Not Found")
	resp.WriteHeader(http.StatusNotFound)
	fmt.Fprint(resp, "Not Found")
}

func respondWithJSON(resp http.ResponseWriter, data interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(data)
}
