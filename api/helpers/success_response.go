package helpers


import (
    "net/http"
)

func SuccessResponse(w http.ResponseWriter, status int, message string) {
    response := map[string]interface{}{
        "error":   false,
        "message": message,
    }
    JSONResponse(w, status, response)
}