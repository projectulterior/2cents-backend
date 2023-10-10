package resolver

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type Cents struct {
	svc *services.Services

	userID format.UserID
	getter[*cents.Cents, func(context.Context) (*cents.Cents, error)]
}

func NewCentsByID(svc *services.Services, userID format.UserID) *Cents {
	return &Cents{
		svc:    svc,
		userID: userID,
		getter: NewGetter(
			func(ctx context.Context) (*cents.Cents, error) {
				return svc.Cents.GetCents(ctx, cents.GetCentsRequest{
					UserID: userID,
				})
			},
		),
	}
}

func NewCentsWithData(svc *services.Services, data *cents.Cents) *Cents {
	return &Cents{
		svc:    svc,
		userID: data.UserID,
		getter: NewGetter(
			func(ctx context.Context) (*cents.Cents, error) {
				return data, nil
			},
		),
	}
}

func (c *Cents) ID(ctx context.Context) (*Cents, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewCentsByID(c.svc, reply.UserID), nil
}

func (c *Cents) TotalCents(ctx context.Context) (*int, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Total, nil
}

func (c *Cents) Deposited(ctx context.Context) (*int, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Deposited, nil
}

func (c *Cents) EarnedCents(ctx context.Context) (*int, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Received, nil
}

func (c *Cents) Given(ctx context.Context) (*int, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Sent, nil
}

func (c *Cents) UpdatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.UpdatedAt, nil
}
