package likes

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

type GetLikesRequest struct {
	Cursor string
	Limit  int
}

type GetLikesResponse struct {
	Likes []*Like
	Next  string
}

func (s *Service) GetLikes(ctx context.Context, req *GetLikesRequest) (*GetLikesResponse, error) {
	type Cursor struct {
		CreatedAt time.Time
		LikeID    format.LikeID
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
				"cr,eated_at": bson.M{
					"$lt": cursor.CreatedAt,
				},
			},
			bson.M{
				"$and": bson.A{
					bson.M{
						"created_at": bson.M{
							"$gte": cursor.LikeID.String(),
						},
					},
				},
			},
		}
	}

	cursor, err := s.Collection(LIKES_COLLECTION).
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

	var likes []*Like
	err = cursor.All(ctx, &likes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var next []byte
	if len(likes) > req.Limit {
		last := likes[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt: last.CreatedAt,
			LikeID:    last.LikeID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	end := req.Limit
	if len(likes) < req.Limit {
		end = len(likes)
	}

	return &GetLikesResponse{
		Likes: likes[:end],
		Next:  string(next),
	}, nil
}
