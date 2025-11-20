package controllers


import (
	"github.com/gorilla/mux"
	"lipad/middleware"
)


func (s *Server) initRoutes() {
	s.Router = mux.NewRouter()

	s.Router.HandleFunc("/", s.Home).Methods("GET")
	s.Router.HandleFunc("/users",middleware.JWTAuth(s.User)).Methods("POST")
	s.Router.HandleFunc("/loan-requests", middleware.JWTAuth(s.LoanRequests)).Methods("POST")
	s.Router.HandleFunc("/loan-requests/:id", middleware.JWTAuth(s.ListLoanRequests)).Methods("GET")
	s.Router.HandleFunc("/webhook/credit-score", middleware.JWTAuth(s.CreditScoreWebhook)).Methods("POST") // POST webhook

	// traceability of incoming/outgoing requests
	loggedRouter := LoggingMiddleware(s.Router)
	http.ListenAndServe(":8080", loggedRouter)
}