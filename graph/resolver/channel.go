package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type Channel struct {
	svc *services.Services

	channelID format.ChannelID
	getter[*messaging.Channel, func(context.Context) (*messaging.Channel, error)]
}

func NewChannelByID(svc *services.Services, channelID format.ChannelID) *Channel {
	return &Channel{
		svc:       svc,
		channelID: channelID,
		getter: NewGetter(
			func(ctx context.Context) (*messaging.Channel, error) {
				return svc.Messaging.GetChannel(ctx, messaging.GetChannelRequest{
					ChannelID: channelID,
				})
			},
		),
	}
}

func NewChannelWithData(svc *services.Services, data *messaging.Channel) *Channel {
	return &Channel{
		svc:       svc,
		channelID: data.ChannelID,
		getter: NewGetter(
			func(ctx context.Context) (*messaging.Channel, error) {
				return data, nil
			},
		),
	}
}

func (c *Channel) ID(ctx context.Context) (string, error) {
	return c.channelID.String(), nil
}