package resolver

import (
	"context"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type ChannelMessages struct {
	svc *services.Services
	getter[*messaging.GetMessagesResponse, func(context.Context) (*messaging.GetMessagesResponse, error)]
}

func NewChannelMessages(svc *services.Services, channelID format.ChannelID, page Pagination) *ChannelMessages {
	return &ChannelMessages{
		svc: svc,
		getter: NewGetter(
			func(ctx context.Context) (*messaging.GetMessagesResponse, error) {
				return svc.Messaging.GetMessages(ctx, &messaging.GetMessagesRequest{
					ChannelID: channelID,
					Cursor:    page.Cursor,
					Limit:     page.Limit,
				})
			},
		),
	}
}

func (m *ChannelMessages) Messages(ctx context.Context) ([]*Message, error) {
	reply, err := m.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	var toRet []*Message
	for _, message := range reply.Messages {
		toRet = append(toRet, NewMessageWithData(m.svc, message))
	}

	return toRet, nil
}

func (m *ChannelMessages) Next(ctx context.Context) (*string, error) {
	reply, err := m.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Next, nil
}
