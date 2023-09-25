package cents

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
)

const (
	CENTS_COLLECTION = "cents"
)

type Cents struct {
	UserID      format.UserID `bson:"_id"`
	TotalCents  int           `bson:"total_cents"`
	Deposited   int           `bson:"deposited"`
	EarnedCents int           `bson:"earned_cents"`
	Given       int           `bson:"given"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}

func (s *Service) Setup(ctx context.Context) error {
	return nil
}
