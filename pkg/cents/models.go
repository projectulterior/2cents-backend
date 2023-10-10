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
	UserID    format.UserID `bson:"_id"`
	Total     int           `bson:"total"`
	Deposited int           `bson:"deposited"`
	Received  int           `bson:"received"`
	Sent      int           `bson:"sent"`
	UpdatedAt time.Time     `bson:"updated_at"`
}

func (s *Service) Setup(ctx context.Context) error {
	return nil
}
