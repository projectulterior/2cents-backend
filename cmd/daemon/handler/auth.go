package handler

import (
	"encoding/json"
	"net/http"

	"github.com/projectulterior/2cents-backend/cmd/daemon/httputil"
	"github.com/projectulterior/2cents-backend/pkg/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleCreateToken(svc *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request auth.CreateTokenRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			httputil.JSONError(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		reply, err := svc.CreateToken(r.Context(), request)
		st, ok := status.FromError(err)
		if !ok {
			httputil.JSONError(w, http.StatusInternalServerError, "error in decoding response", nil)
			return
		}

		switch st.Code() {
		case codes.OK:
			httputil.JSONSuccess(w, http.StatusOK, reply)
		case codes.PermissionDenied:
			httputil.JSONError(w, http.StatusForbidden, err.Error(), nil)
		case codes.NotFound:
			httputil.JSONError(w, http.StatusNotFound, err.Error(), nil)
		default:
			httputil.JSONError(w, http.StatusInternalServerError, err.Error(), nil)
		}
	}
}

func HandleRefreshToken(svc *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request auth.RefreshTokenRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			httputil.JSONError(w, http.StatusBadRequest, err.Error(), nil)
			return
		}

		reply, err := svc.RefreshToken(r.Context(), request)
		st, ok := status.FromError(err)
		if !ok {
			httputil.JSONError(w, http.StatusInternalServerError, "error in decoding response", nil)
			return
		}

		switch st.Code() {
		case codes.OK:
			httputil.JSONSuccess(w, http.StatusOK, reply)
		case codes.PermissionDenied:
			httputil.JSONError(w, http.StatusForbidden, err.Error(), nil)
		case codes.InvalidArgument:
			httputil.JSONError(w, http.StatusBadRequest, err.Error(), nil)
		case codes.NotFound:
			httputil.JSONError(w, http.StatusNotFound, err.Error(), nil)
		default:
			httputil.JSONError(w, http.StatusInternalServerError, err.Error(), nil)
		}
	}
}
