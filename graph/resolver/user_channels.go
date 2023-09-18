package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type UserChannels struct {
	svc *services.Services
	getter[*messaging.GetChannelsResponse, func(context.Context) (*messaging.GetChannelsResponse, error)]
}

func NewUserChannels(svc *services.Services, userID format.UserID, page Pagination) *UserChannels {
	return &UserChannels{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*messaging.GetChannelsResponse, error) {
				return svc.Messaging.GetChannels(ctx, messaging.GetChannelsRequest{
					MemberID: userID,
					Cursor:   page.Cursor,
					Limit:    page.Limit,
				})
			},
		),
	}
}

func (c *UserChannels) Channels(ctx context.Context) ([]*Channel, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*Channel
	for _, channel := range reply.Channels {
		toRet = append(toRet, NewChannelWithData(c.svc, channel))
	}

	return toRet, nil
}

func (c *UserChannels) Next(ctx context.Context) (*string, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
