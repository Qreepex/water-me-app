package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type ctxKey string

const userIDKey ctxKey = "userID"

func WithUserID(r *http.Request, userID string) *http.Request {
    ctx := context.WithValue(r.Context(), userIDKey, userID)
    return r.WithContext(ctx)
}

// GetUserID extracts userID from request context.
func GetUserID(r *http.Request) (string, bool) {
    id, ok := r.Context().Value(userIDKey).(string)
    return id, ok
}

// AuthMiddleware validates Bearer JWT and injects userID into context.
func AuthMiddleware(secret string, next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        auth := r.Header.Get("Authorization")
        if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        tokenStr := strings.TrimPrefix(auth, "Bearer ")
        token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) { return []byte(secret), nil })
        if err != nil || !token.Valid {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        uidAny := claims["uid"]
        uid, ok := uidAny.(string)
        if !ok || uid == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next(w, WithUserID(r, uid))
    }
}
