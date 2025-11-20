package middleware

package middleware

import (
    "net/http"
    "strings"
    "github.com/golang-jwt/jwt/v5"
    "lipad/utils"
    "fmt"
)

var JwtSecret = os.Getenv("JWT")  // replace with env var in production

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            ErrorResponse(w, "missing Authorization header")
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            ErrorResponse(w, "invalid Authorization header format")
            return
        }

        tokenString := parts[1]

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method")
            }
            return JwtSecret, nil
        })

        if err != nil || !token.Valid {
            ErrorResponse(w, "invalid token")
            return
        }
        next.ServeHTTP(w, r)
    }
}
