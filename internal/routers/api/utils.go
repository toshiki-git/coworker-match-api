package api

import (
	"encoding/json"
	"net/http"
)

// writeErrorはエラーメッセージをJSON形式でレスポンスに書き込むヘルパー関数
func writeError(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": errorMessage})
}
