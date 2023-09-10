package posts

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

type GetPostsRequest struct {
	Cursor   string
	Limit    int
	AuthorID *format.UserID
}

type GetPostsResponse struct {
	Posts []*Post
	Next  string
}

func (s *Service) GetPosts(ctx context.Context, req *GetPostsRequest) (*GetPostsResponse, error) {
	type Cursor struct {
		CreatedAt time.Time
		PostID    format.PostID
	}

	filter := bson.M{}

	if req.AuthorID != nil {
		filter["author_id"] = req.AuthorID.String()
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
							"$gte": cursor.PostID.String(),
						},
					},
				},
			},
		}
	}

	cursor, err := s.Collection(POSTS_COLLECTION).
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

	var posts []*Post
	err = cursor.All(ctx, &posts)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	var next []byte
	if len(posts) > req.Limit {
		last := posts[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt: last.CreatedAt,
			PostID:    last.PostID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	end := req.Limit
	if len(posts) < req.Limit {
		end = len(posts)
	}

	return &GetPostsResponse{
		Posts: posts[:end],
		Next:  string(next),
	}, nil
}
