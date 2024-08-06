package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coworker-match-api/internal/common"
)

func respondWithJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		writeError(w, fmt.Sprintf("Error encoding response: %v", err), http.StatusInternalServerError)
	}
}

// writeErrorはエラーメッセージをJSON形式でレスポンスに書き込むヘルパー関数
func writeError(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": errorMessage})
}

// コンテキストからユーザーIDを取得するヘルパー関数
func GetUserID(ctx context.Context) (string, bool) {
	key := common.UserIdKey
	userID, ok := ctx.Value(key).(string)
	return userID, ok
}
