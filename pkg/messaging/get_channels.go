package messaging

import (
	"context"
	"encoding/json"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetChannelsRequest struct {
	MemberID format.UserID
	Cursor   string
	Limit    int
}

type GetChannelsResponse struct {
	Channels []*Channel
	Next     string
}

func (s *Service) GetChannels(ctx context.Context, req GetChannelsRequest) (*GetChannelsResponse, error) {
	type Cursor struct {
		ChannelID format.ChannelID `json:"channel_id"`
		CreatedAt time.Time        `json:"created_at"`
	}

	filter := bson.M{
		"latest": bson.M{
			"$ne": bson.A{},
		},
	}

	if req.Cursor != "" {
		var cursor Cursor
		err := json.Unmarshal([]byte(req.Cursor), &cursor)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		filter["$or"] = bson.A{
			bson.M{
				"latest.created_at": bson.M{
					"$lt": cursor.CreatedAt,
				},
			},
			bson.M{
				"$and": bson.A{
					bson.M{
						"latest.created_at": bson.M{
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

	cursor, err := s.Collection(CHANNELS_COLLECTION).
		Aggregate(ctx,
			bson.A{
				bson.M{
					"$match": bson.M{
						"member_ids": req.MemberID.String(),
					},
				},
				bson.M{
					"$lookup": bson.M{
						"from": MESSAGES_COLLECTION,
						"let": bson.M{
							"channel_id": "$_id",
						},
						"pipeline": bson.A{
							bson.M{
								"$match": bson.M{
									"$expr": bson.M{
										"$and": bson.A{
											bson.M{
												"$eq": bson.A{
													"$channel_id", "$$channel_id",
												},
											},
										},
									},
								},
							},
							bson.M{
								"$sort": bson.M{
									"created_at": -1,
								},
							},
							bson.M{
								"$limit": 1,
							},
						},
						"as": "latest",
					},
				},
				bson.M{
					"$match": filter,
				},
				bson.M{
					"$sort": bson.D{
						{Key: "latest.created_at", Value: -1},
						{Key: "_id", Value: 1},
					},
				},
				bson.M{
					"$limit": req.Limit + 1,
				},
			},
		)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var channels []*Channel
	err = cursor.All(ctx, &channels)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var next []byte
	if len(channels) > req.Limit {
		last := channels[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt: last.Latest[0].CreatedAt,
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
