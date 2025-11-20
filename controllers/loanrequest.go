package controllers
import(
	"lipad/entities"
	"net/http"
	"encoding/json"
	"lipad/middleware"
	"lipad/models"
	"fmt"
)

func (s *Server) LoanRequests (w http.ResponseWriter, r *http.Request) {
	var data entities.LoanRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	if err !=nil {
		middleware.ErrorResponse(w, "invalid data")
		return
	}
	err = models.CreateLoanRequest(data, s.Db)
	if err !=nil {
		middleware.ErrorResponse(w, "unable to save loan request")
		return
	}
	middleware.OkResponse(w, 200, data.Status)
}

func (s *Server) ListLoanRequests(w http.ResponseWriter, r *http.Request) {
    // Decode the request body to get the user ID
    var requestData struct {
        UserID int `json:"user_id"`
    }

    err := json.NewDecoder(r.Body).Decode(&requestData)
    if err != nil {
        middleware.ErrorResponse(w, "invalid data")
        return
    }

    // Fetch loan requests from DB
    loans, err := models.GetLoanRequests(requestData.UserID, s.Db)
    if err != nil {
        middleware.ErrorResponse(w, "could not fetch loan requests")
        return
    }
    // Respond with the fetched loan requests
    middleware.OkResponse(w, 200, loans)
}

//webhook logic
func (s *Server) CreditScoreWebhook(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        LoanID string    `json:"loan_id"`
        Status string `json:"status"`
        Reason string `json:"reason"`
    }

    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil {
        middleware.ErrorResponse(w, "invalid data")
        return
    }

    if payload.Status != "APPROVED" && payload.Status != "REJECTED" {
        middleware.ErrorResponse(w, "invalid status value")
        return
    }

    err = models.UpdateLoanStatus(payload.LoanID, payload.Status, payload.Reason, s.Db)
    if err != nil {
        middleware.ErrorResponse(w, "could not update loan request")
        return
    }

    middleware.OkResponse(w, 200, fmt.Sprintf("loan %d updated to %s", payload.LoanID, payload.Status))
}
