package search

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/totalhitsrelation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUsersRequest struct {
	Query  string
	Cursor string
	Limit  int
}

type GetUsersResponse struct {
	Users []*User
	Next  string
}

func (s *Service) GetUsers(ctx context.Context, req GetUsersRequest) (*GetUsersResponse, error) {
	type Cursor struct {
		PID    string             `json:"pid"`
		After  []types.FieldValue `json:"after"`
		Offset int                `json:"offset"`
	}

	var cursor Cursor
	if req.Cursor != "" {
		err := json.Unmarshal([]byte(req.Cursor), &cursor)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	} else {
		resp, err := s.OpenPointInTime(s.UsersIndex).Index(s.UsersIndex).KeepAlive("30s").Do(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		cursor.PID = resp.Id
	}

	resp, err := s.Search().
		Request(&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Should: []types.Query{
						{
							Match: map[string]types.MatchQuery{
								"username": {
									Query: req.Query,
								},
							},
						},
						{
							Match: map[string]types.MatchQuery{
								"name": {
									Query: req.Query,
								},
							},
						},
					},
				},
			},
			Size:           &req.Limit,
			TrackTotalHits: true,
			Pit: &types.PointInTimeReference{
				Id: cursor.PID,
				// KeepAlive: 30 * time.Second,
			},
			SearchAfter: cursor.After,
			Sort: []types.SortCombinations{
				types.SortOptions{Score_: &types.ScoreSort{Order: &sortorder.Desc}},
				types.SortOptions{Doc_: &types.ScoreSort{Order: &sortorder.Desc}},
			},
		}).
		Do(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	toRet := GetUsersResponse{
		Users: make([]*User, 0),
	}
	for _, hit := range resp.Hits.Hits {
		var user User
		err := json.Unmarshal(hit.Source_, &user)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		toRet.Users = append(toRet.Users, &user)
	}

	// next (cursor)
	if resp.Hits.Total.Relation == totalhitsrelation.Gte || cursor.Offset+len(resp.Hits.Hits) < int(resp.Hits.Total.Value) {
		if len(resp.Hits.Hits) > 0 {
			b, err := json.Marshal(Cursor{
				PID:    cursor.PID,
				After:  resp.Hits.Hits[len(resp.Hits.Hits)-1].Sort,
				Offset: cursor.Offset + len(resp.Hits.Hits),
			})
			if err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			toRet.Next = string(b)
		}
	} else {
		// close pit
		_, err := s.ClosePointInTime().
			Id(cursor.PID).
			Do(ctx)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &toRet, nil
}
