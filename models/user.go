package models

import(
	"lipad/entities"
	"lipad/utils"
	"database/sql"
)
func CreateUser (data entities.User, db *sql.DB) (err error) {
	query := "insert into user (name, email, phone_number) values (?,?,?)"
	_, err = db.Exec(query, data.Name, data.Email, data.PhoneNumber)
	if err != nil {
		utils.Log("an error occured", err)
		return
	}
	utils.Log("success")
	return
}

