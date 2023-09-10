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

type GetChannelsRequest struct {
	Cursor string
	Limit  int
}

type GetChannelsResponse struct {
	Channels []*Channel
	Next     string
}

func (s *Service) GetChannels(ctx context.Context, req *GetChannelsRequest) (*GetChannelsResponse, error) {
	type Cursor struct {
		CreatedAt time.Time        `json:"created_at"`
		ChannelID format.ChannelID `json:"channel_id"`
	}

	filter := bson.M{}

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
							"$gte": cursor.ChannelID.String(),
						},
					},
				},
			},
		}
	}

	cusor, err := s.Collection(CHANNELS_COLLECTION).
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

	var channels []*Channel
	err = cusor.All(ctx, &channels)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var next []byte
	if len(channels) > req.Limit {
		last := channels[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt: last.CreatedAt,
			ChannelID: last.ChannelID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	end := req.Limit
	if len(channels) < req.Limit {
		end = len(channels)
	}

	return &GetChannelsResponse{
		Channels: channels[:end],
		Next:     string(next),
	}, nil
}
