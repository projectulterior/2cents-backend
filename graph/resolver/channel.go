package resolver

import (
	"context"
	"time"

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

func (c *Channel) Members(ctx context.Context) ([]*User, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*User
	for _, memberID := range reply.MemberIDs {
		toRet = append(toRet, NewUserByID(c.svc, memberID))
	}

	return toRet, nil
}

func (c *Channel) Messages(ctx context.Context, page Pagination) (*Messages, error) {
	return NewMessages(NewChannelMessages(c.svc, c.channelID, page)), nil
}

func (c *Channel) CreatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.CreatedAt, nil
}

func (c *Channel) UpdatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := c.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.UpdatedAt, nil
}
