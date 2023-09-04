package likes

import (
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	LIKES_COLLECTION = "posts"
)

type Like struct {
	LikerID   format.UserID `bson:"liker_id"`
	CreatedAt time.Time     `bson:"created_at"`
}
