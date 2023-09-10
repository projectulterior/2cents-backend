package follow

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

type GetFollowsRequest struct {
	Cursor     string
	Limit      int
	FollowerID *format.UserID
}

type GetFollowsResponse struct {
	Follows []*Follow
	Next    string
}

func (s *Service) GetFollows(ctx context.Context, req *GetFollowsRequest) (*GetFollowsResponse, error) {
	type Cursor struct {
		CreatedAt time.Time
		FollowID  format.FollowID
	}

	filter := bson.M{}

	if req.FollowerID != nil {
		filter["follower_id"] = req.FollowerID.String()
	}

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
							"$gte": cursor.FollowID.String(),
						},
					},
				},
			},
		}
	}

	cursor, err := s.Collection(FOLLOW_COLLECTION).
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

	var follows []*Follow
	err = cursor.All(ctx, &follows)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var next []byte
	if len(follows) > req.Limit {
		last := follows[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt: last.CreatedAt,
			FollowID:  last.FollowID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	end := req.Limit
	if len(follows) < req.Limit {
		end = len(follows)
	}

	return &GetFollowsResponse{
		Follows: follows[:end],
		Next:    string(next),
	}, nil
}
