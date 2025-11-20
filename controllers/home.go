package controllers


import (
	"net/http"
	"lipad/utils"
	"lipad/middleware"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	utils.Log("welcome home")
	middleware.OkResponse(w, 200, "Ok")
}