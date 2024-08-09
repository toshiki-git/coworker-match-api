package common

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// ExtractPathParam パスパラメータから指定されたキーの値を抽出し、バリデーションを行う
func ExtractPathParam(r *http.Request, w http.ResponseWriter, paramName string) (string, error) {
	vars := mux.Vars(r)
	paramValue := vars[paramName]
	if paramValue == "" {
		RespondWithError(w, http.StatusBadRequest, "Missing "+paramName)
		return "", errors.New("missing " + paramName)
	}
	return paramValue, nil
}

// ExtractUserIdFromContext コンテキストからユーザーIDを抽出し、バリデーションを行う
func ExtractUserIdFromContext(r *http.Request, w http.ResponseWriter) (string, error) {
	key := UserIdKey
	userId, ok := r.Context().Value(key).(string)
	if !ok {
		RespondWithError(w, http.StatusInternalServerError, "Failed to get userId from context")
		return "", errors.New("failed to get userId from context")
	}
	return userId, nil
}
