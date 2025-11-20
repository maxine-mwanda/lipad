package middleware


import (
"encoding/json"
"net/http"
)

func ErrorResponse(w http.ResponseWriter, error string) {
	OkResponse(w, 400, map[string]string{"error": error})
}

func OkResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Read and log incoming request
        bodyBytes, _ := io.ReadAll(r.Body)
        r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // restore body for handler

        log.Printf("INCOMING REQUEST: %s %s\nHeaders: %v\nBody: %s\n",
            r.Method, r.URL.Path, r.Header, string(bodyBytes))

        // Capture outgoing requests by wrapping http.DefaultClient.Do if needed
        // For now, just call the next handler
        next.ServeHTTP(w, r)
    })
}
