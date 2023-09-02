package middleware

import (
	"context"
	"net/http"

	"github.com/projectulterior/2cents-backend/cmd/daemon/httputil"
	"github.com/projectulterior/2cents-backend/pkg/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ContextKey string

const (
	AUTH_USER_CONTEXT_KEY ContextKey = "AUTH_USER"
)

type Auth struct {
	svc  *auth.Service
	next http.Handler
}

func NewAuth(svc *auth.Service, next http.Handler) *Auth {
	return &Auth{svc: svc, next: next}
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	protocol := r.Header.Get("Sec-Websocket-Protocol")
	if protocol == "graphql-ws" {
		a.next.ServeHTTP(w, r)
		return
	}

	token := r.Header.Get("Authorization")
	if token == "" {
		httputil.JSONError(w, http.StatusUnauthorized, "missing Authorization header -- must have a valid id token", nil)
		return
	}

	reply, err := a.svc.VerifyToken(r.Context(), auth.VerifyTokenRequest{
		Token: token,
	})
	st, ok := status.FromError(err)
	if !ok {
		httputil.JSONError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	switch st.Code() {
	case codes.OK:
	case codes.PermissionDenied:
		httputil.JSONError(w, http.StatusUnauthorized, st.Message(), nil)
	default:
		httputil.JSONError(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	ctx := context.WithValue(r.Context(), AUTH_USER_CONTEXT_KEY, reply.UserID)

	request := r.WithContext(ctx)

	a.next.ServeHTTP(w, request)
}
