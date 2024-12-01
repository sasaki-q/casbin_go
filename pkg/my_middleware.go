package pkg

import (
	"casbin/repositories"
	"context"
	"net/http"
	"strconv"

	"github.com/casbin/casbin/v2"
)

type contextKey string

var ContextKey = contextKey("user-id-key")

func NewMyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid := r.Header.Get("uid")
		if uid == "" {
			ErrorResp(w, r, http.StatusUnauthorized, ErrHeaderValueNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), ContextKey, uid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func MyCasbinMiddleware(enforcer *casbin.Enforcer, ur repositories.IUserRepository) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctxv := r.Context().Value(ContextKey)
			if ctxv == nil {
				ErrorResp(w, r, http.StatusBadRequest, ErrContextValueNotFound)
				return
			}

			uid, err := strconv.Atoi(ctxv.(string))
			if err != nil {
				ErrorResp(w, r, http.StatusBadRequest, err)
				return
			}

			user := ur.GetUser(uid)
			if user == nil {
				ErrorResp(w, r, http.StatusNotFound, ErrUserNotFound)
				return
			}

			allowed, err := enforcer.Enforce(string(user.Type), r.URL.Path, r.Method)
			if err != nil {
				ErrorResp(w, r, http.StatusForbidden, err)
				return
			}

			if !allowed {
				ErrorResp(w, r, http.StatusForbidden, ErrNotAllowed)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
