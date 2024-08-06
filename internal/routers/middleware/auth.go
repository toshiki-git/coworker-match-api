package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/coworker-match-api/internal/common"
	"google.golang.org/api/idtoken"
)

func Auth(next http.Handler) http.Handler {
	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// Authorizationヘッダーの取得
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Authorization header required"})
			return
		}

		// IDトークンの取得
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Authorization header format"})
			return
		}
		idToken := parts[1]

		// トークンの検証
		ctx := context.Background()
		payload, err := idtoken.Validate(ctx, idToken, GOOGLE_CLIENT_ID)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		// トークンのペイロードから必要な情報を取得（例：ユーザーID）
		userId, ok := payload.Claims["sub"].(string)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid token payload"})
			return
		}

		// リクエストのコンテキストにユーザー情報を追加
		key := common.UserIdKey
		ctx = context.WithValue(r.Context(), key, userId)
		r = r.WithContext(ctx)

		// デバッグ用のヘッダーを追加
		w.Header().Set("X-Debug-UserId", userId)

		next.ServeHTTP(w, r)
	})
}
