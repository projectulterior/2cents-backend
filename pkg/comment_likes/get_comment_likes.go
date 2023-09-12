package comment_likes

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

type GetCommentLikesRequest struct {
	Cursor    string
	Limit     int
	LikerID   *format.UserID
	CommentID *format.CommentID
}

type GetCommentLikesResponse struct {
	CommentLikes []*CommentLike
	Next         string
}

func (s *Service) GetCommentLikes(ctx context.Context, req *GetCommentLikesRequest) (*GetCommentLikesResponse, error) {
	type Cursor struct {
		CreatedAt     time.Time
		CommentLikeID format.CommentLikeID
	}

	filter := bson.M{}

	if req.LikerID != nil {
		filter["liker_id"] = req.LikerID.String()
	}

	if req.CommentID != nil {
		filter["comment_id"] = req.CommentID.String()
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
							"$gte": cursor.CommentLikeID.String(),
						},
					},
				},
			},
		}
	}

	cursor, err := s.Collection(COMMENT_LIKES_COLLECTION).
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

	var commentLikes []*CommentLike
	err = cursor.All(ctx, &commentLikes)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var next []byte
	if len(commentLikes) > req.Limit {
		last := commentLikes[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt:     last.CreatedAt,
			CommentLikeID: last.CommentLikeID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	end := req.Limit
	if len(commentLikes) < req.Limit {
		end = len(commentLikes)
	}

	return &GetCommentLikesResponse{
		CommentLikes: commentLikes[:end],
		Next:         string(next),
	}, nil
}
