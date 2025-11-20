package controllers
import(
	"net/http"
	"lipad/entities"
	"lipad/middleware"
	"lipad/models"
	"encoding/json"
)

func (s *Server) User (w http.ResponseWriter, r *http.Request) {
	var data entities.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err !=nil {
		middleware.ErrorResponse(w, "invalid data")
		return
	}
	err = models.CreateUser(data, s.Db)
	if err !=nil {
		middleware.ErrorResponse(w, "unable to save user")
		return
	}
	middleware.OkResponse(w, 200, data.Name)
}