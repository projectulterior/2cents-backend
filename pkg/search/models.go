package search

import (
	"context"

	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/projectulterior/2cents-backend/pkg/format"
)

var USERS_MAPPING = types.TypeMapping{
	Properties: map[string]types.Property{
		"user_id":  types.NewKeywordProperty(),
		"username": types.NewTextProperty(),
		"name":     types.NewTextProperty(),
	},
}

type User struct {
	UserID   format.UserID `json:"user_id,omitempty"`
	Username string        `json:"username,omitempty"`
	Name     string        `json:"name,omitempty"`
}

func (s *Service) Setup(ctx context.Context) error {
	_, err := s.Indices.Create(s.UsersIndex).
		Request(&create.Request{
			Mappings: &USERS_MAPPING,
		}).Do(ctx)
	if err != nil {
		// check if duplicate error
		// return err
	}

	return nil
}
