package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"google.golang.org/api/idtoken"
)

func Auth(next http.Handler) http.Handler {
	GOOGLE_CLIENT_ID := os.Getenv("GOOGLE_CLIENT_ID")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Authorizationヘッダーの取得
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// IDトークンの取得
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}
		idToken := parts[1]

		// トークンの検証
		ctx := context.Background()
		payload, err := idtoken.Validate(ctx, idToken, GOOGLE_CLIENT_ID)
		if err != nil {
			http.Error(w, "Invalid ID token", http.StatusUnauthorized)
			return
		}
		// トークンのペイロードから必要な情報を取得（例：ユーザーID）
		userID := payload.Claims["sub"].(string)

		// リクエストのコンテキストにユーザー情報を追加
		ctx = context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
