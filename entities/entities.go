package entities

const (
    StatusPending  Status = "PENDING"
    StatusApproved Status = "APPROVED"
    StatusRejected Status = "REJECTED"
)

type User struct{
	UserID int `json:"user_id"`
	Name string `json:"string"`
	Email string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	//id, name, email, phone_number
}

type Status string

type LoanRequest struct{
	LoanID int `json:"loan_id"`
	UserID int `json:"user_id"`
	Amount float64 `json:"amount"`
	Status Status `json:"status"`
	Reason string `json:"reason"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"` 

	//id, user_id, amount, status (PENDING, APPROVED, REJECTED), reason, created_at, updated_at
}

type CreditScoreCallback struct {
	LoanID int    `json:"loan_id"`
	Score  int    `json:"score"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}