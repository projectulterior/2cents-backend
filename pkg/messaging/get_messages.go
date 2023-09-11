package messaging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetMessagesRequest struct {
	ChannelID format.ChannelID
	Cursor    string
	Limit     int
}

type GetMessagesResponse struct {
	Messages []*Message
	Next     string
}

func (s *Service) GetMessages(ctx context.Context, req *GetMessagesRequest) (*GetMessagesResponse, error) {
	type Cursor struct {
		CreatedAt time.Time        `json:"created_at"`
		MessageID format.MessageID `json:"message_id"`
	}

	filter := bson.M{"channel_id": req.ChannelID.String()}

	if req.Cursor != "" {
		var cursor Cursor
		err := json.Unmarshal([]byte(req.Cursor), &cursor)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		filter["$or"] = bson.A{
			bson.M{
				"created_at": bson.M{
					"$lt": cursor.CreatedAt,
				},
			},
			bson.M{
				"$and": bson.A{
					bson.M{
						"created_at": bson.M{
							"$eq": cursor.CreatedAt,
						},
					},
					bson.M{
						"_id": bson.M{
							"$gte": cursor.MessageID.String(),
						},
					},
				},
			},
		}
	}

	cusor, err := s.Collection(MESSAGES_COLLECTION).
		Find(ctx,
			filter,
			options.Find().
				SetSort(bson.D{
					{Key: "created_at", Value: -1},
					{Key: "_id", Value: 1},
				}).SetLimit(int64(req.Limit+1)),
		)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var messages []*Message
	err = cusor.All(ctx, &messages)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var next []byte
	if len(messages) > req.Limit {
		last := messages[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt: last.CreatedAt,
			MessageID: last.MessageID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	end := req.Limit
	if len(messages) < req.Limit {
		end = len(messages)
	}

	return &GetMessagesResponse{
		Messages: messages[:end],
		Next:     string(next),
	}, nil
}
