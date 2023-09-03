package users

import (
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	USERS_COLLECTION = "users"
)

type User struct {
	UserID    format.UserID `bson:"_id"`
	Name      string        `bson:"name"`
	Email     string        `bson:"email"`
	Bio       string        `bson:"bio"`
	CreatedAt time.Time     `bson:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at"`
}
