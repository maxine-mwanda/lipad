package models

import (
	"lipad/entities"
	"database/sql"
	"lipad/utils"
	"time"
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
)

func CreateLoanRequest(data entities.LoanRequest, db *sql.DB) (err error) {
    // 1. Validate user exists
    var exists int
    queryUser := "SELECT COUNT(1) FROM user WHERE user_id = ?"
    err = db.QueryRow(queryUser, data.UserID).Scan(&exists)
    if err != nil {
        utils.Log("error checking if user exists", err)
        return
    }
    if exists == 0 {
        utils.Log("user does not exist")
        return fmt.Errorf("user does not exist")
    }

    // 2. Validate amount
    if data.Amount <= 0 || data.Amount > 1000000 {
        utils.Log("invalid loan amount")
        return fmt.Errorf("amount must be > 0 and < 1,000,000")
    }

    // 3. Reject duplicate pending requests
    var pending int
    queryPending := "SELECT COUNT(1) FROM loan_requests WHERE user_id = ? AND status = 'pending'"
    err = db.QueryRow(queryPending, data.UserID).Scan(&pending)
    if err != nil {
        utils.Log("error checking pending requests", err)
        return
    }
    if pending > 0 {
        utils.Log("user has existing pending request")
        return fmt.Errorf("duplicate pending request exists")
    }

    // 4. Insert loan request
    insertQuery := "INSERT INTO loan_requests (user_id, amount, status, reason, created_at, updated_at) VALUES (?,?,?,?,?,?)"
    now := time.Now().Format("2006-01-02 15:04:05")
    _, err = db.Exec(insertQuery, data.UserID, data.Amount, "pending", data.Reason, now, now)
    if err != nil {
        utils.Log("error inserting loan request", err)
        return
    }

    utils.Log("loan request created successfully")
	//line simulating callback
	go simulateCreditScoreCallback(int(data.LoanID), data.Amount, "Maxine Mwanda", "maxine@example.com")
    return
}

func GetLoanRequests(userID int, db *sql.DB) (loans []entities.LoanRequest, err error) {
    query := "SELECT loan_id, user_id, amount, status, reason, created_at, updated_at FROM loan_requests WHERE user_id = ?"
    rows, err := db.Query(query, userID)
    if err != nil {
        utils.Log("error fetching loan requests", err)
        return
    }
    defer rows.Close()

    for rows.Next() {
        var loan entities.LoanRequest
        err = rows.Scan(
            &loan.LoanID,
            &loan.UserID,
            &loan.Amount,
            &loan.Status,
            &loan.Reason,
            &loan.CreatedAt,
            &loan.UpdatedAt,
        )
        if err != nil {
            utils.Log("error scanning loan request", err)
            continue
        }
        loans = append(loans, loan)
    }

    utils.Log("loan requests fetched successfully")
    return
}

func UpdateLoanStatus(loanID string, status, reason string, db *sql.DB) (err error) {
    updateQuery := `
        UPDATE loan_requests
        SET status = ?, reason = ?, updated_at = ?
        WHERE loan_id = ?
    `
    now := time.Now().Format("2006-01-02 15:04:05")
    _, err = db.Exec(updateQuery, status, reason, now, loanID)
    if err != nil {
        utils.Log("error updating loan request", err)
        return fmt.Errorf("could not update loan request")
    }
    utils.Log("loan request updated successfully")
    return
}

func simulateCreditScoreCallback(loanID int, amount float64, name, email string) {
    time.Sleep(2 * time.Second) // api wait

	payload := entities.CreditScoreCallback{
		LoanID: loanID,
		Score:  720,
		Status: "APPROVED",
		Reason: "Good credit history",
	}

    body, _ := json.Marshal(payload)

    resp, err := http.Post("http://localhost:8080/webhook/credit-score", "application/json", bytes.NewBuffer(body))
    if err != nil {
        utils.Log("failed to call webhook", err)
        return
    }
    defer resp.Body.Close()

    utils.Log("simulating credit API callback for loan", loanID, "status code:", resp.StatusCode)
}