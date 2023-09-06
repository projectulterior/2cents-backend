package follow

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	FOLLOW_COLLECTION = "follows"
)

type Follow struct {
	FollowID   format.FollowID `bson:"_id"`
	FollowerID format.UserID   `bson:"follower_id"`
	FolloweeID format.UserID   `bson:"followee_id"`
	CreatedAt  time.Time       `bson:"created_at"`
}

func (s *Service) Setup(ctx context.Context) error {
	return nil
}
