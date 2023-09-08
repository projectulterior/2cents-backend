package users

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

type GetUsersRequest struct {
	Cursor string
	Limit  int
}

type GetUsersResponse struct {
	Users []*User
	Next  string
}

// GetUsers returns a list of all users in created_at order
func (s *Service) GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {
	type Cursor struct {
		CreatedAt time.Time     `json:"created_at"` // main sort field
		UserID    format.UserID `json:"user_id"`
	}

	filter := bson.M{}

	// Parse cursor
	if req.Cursor != "" {
		// not the first query

		// parse cursor
		var cursor Cursor
		err := json.Unmarshal([]byte(req.Cursor), &cursor)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		// add filters such that we only get documents after the cursor
		//
		// All documents returned should satified either of the two:
		// 1. cursor.CreatedAt > doc.created_at
		// 2. if created_at == cursor.CreatedAt => _id >= cursor.user_id

		filter["$or"] = bson.A{
			bson.M{
				// doc.created_at is less than cursor.Created_at
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
							"$gte": cursor.UserID.String(),
						},
					},
				},
			},
		}
	}

	// call mongodb with filter
	cusor, err := s.Collection(USERS_COLLECTION).
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

	// Decode all user document to accessible (array of) structs
	var users []*User
	err = cusor.All(ctx, &users)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// create "next" cursor, if there are more results than req.Limit
	var next []byte
	if len(users) > req.Limit {
		last := users[req.Limit]

		next, err = json.Marshal(Cursor{
			CreatedAt: last.CreatedAt,
			UserID:    last.UserID,
		})
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	// calculate the last element to be returned based on req.Limit
	end := req.Limit
	if len(users) < req.Limit {
		end = len(users)
	}

	return &GetUsersResponse{
		Users: users[:end],
		Next:  string(next),
	}, nil
}
