package resolver

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type Message struct {
	svc *services.Services

	messageID format.MessageID
	getter[*messaging.Message, func(context.Context) (*messaging.Message, error)]
}

func NewMessageByID(svc *services.Services, messageID format.MessageID) *Message {
	return &Message{
		svc:       svc,
		messageID: messageID,
		getter: NewGetter(
			func(ctx context.Context) (*messaging.Message, error) {
				return svc.Messaging.GetMessage(ctx, messaging.GetMessageRequest{
					MessageID: messageID,
				})
			},
		),
	}
}

func (m *Message) ID(ctx context.Context) (string, error) {
	return m.messageID.String(), nil
}

func (m *Message) Channel(ctx context.Context) (*Channel, error) {
	reply, err := m.getter.Call(ctx)
	if err != nil {
		return nil, err
	}
	return NewChannelByID(m.svc, reply.ChannelID), nil
}

func (m *Message) Content(ctx context.Context) (*string, error) {
	reply, err := m.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.Content, nil
}

func (m *Message) ContentType(ctx context.Context) (*format.ContentType, error) {
	reply, err := m.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.ContentType, nil
}

func (m *Message) CreatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := m.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.CreatedAt, nil
}

func (m *Message) Sender(ctx context.Context) (*User, error) {
	reply, err := m.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewUserByID(m.svc, reply.SenderID), nil
}
