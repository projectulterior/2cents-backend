package comments

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

type GetCommentsRequest struct {
	Cursor string
	Limit  int
}

type GetCommentsResponse struct {
	Comments []*Comment
	Next     string
}

func (s *Service) GetComments(ctx context.Context, req *GetCommentsRequest) (*GetCommentsResponse, error) {
	type Cursor struct {
		CreatedAt time.Time        `json:"created_at"`
		CommentID format.CommentID `json:"comment_id"`
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
							"$gte": cursor.CommentID.String(),
						},
					},
				},
			},
		}
	}

	cusor, err := s.Collection(COMMENTS_COLLECTION).
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

	var comments []*Comment
	err = cusor.All(ctx, &comments)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var next []byte
	if len(comments) > req.Limit {
		last := comments[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt: last.CreatedAt,
			CommentID: last.CommentID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	end := req.Limit
	if len(comments) < req.Limit {
		end = len(comments)
	}

	return &GetCommentsResponse{
		Comments: comments[:end],
		Next:     string(next),
	}, nil
}
