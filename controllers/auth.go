package middleware

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    err := json.NewDecoder(r.Body).Decode(&creds)
    if err != nil {
        middleware.ErrorResponse(w, "invalid data")
        return
    }

    // verify email + password with DB
    var userID int
    err = s.DB.QueryRow("SELECT user_id FROM user WHERE email=? AND password=?", creds.Email, creds.Password).Scan(&userID)
    if err != nil {
        middleware.ErrorResponse(w, "invalid credentials")
        return
    }

    // Generate JWT
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1 day expiration
    })

    tokenString, err := token.SignedString(middleware.JwtSecret)
    if err != nil {
        middleware.ErrorResponse(w, "could not generate token")
        return
    }

    middleware.OkResponse(w, 200, map[string]string{"token": tokenString})
}
