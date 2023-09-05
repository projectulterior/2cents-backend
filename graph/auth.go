package graph

import (
	"context"
	"net/http"

	"github.com/projectulterior/2cents-backend/cmd/daemon/middleware"
	"github.com/projectulterior/2cents-backend/pkg/format"
)

func authUserID(ctx context.Context) (format.UserID, error) {
	userID, ok := ctx.Value(middleware.AUTH_USER_CONTEXT_KEY).(format.UserID)
	if !ok {
		return "", e(ctx, http.StatusUnauthorized, "no auth user")
	}

	return userID, nil
}
