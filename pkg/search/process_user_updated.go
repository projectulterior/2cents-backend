package search

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/projectulterior/2cents-backend/pkg/users"
)

func (s *Service) ProcessUserUpdated(ctx context.Context, event users.UserUpdatedEvent) error {
	b, err := json.Marshal(User{
		UserID: event.User.UserID,
		Name:   event.User.Name,
	})
	if err != nil {
		return err
	}

	docAsUpsert := true
	_, err = s.Update(s.UsersIndex, event.User.UserID.String()).
		Request(&update.Request{
			Doc:         b,
			DocAsUpsert: &docAsUpsert,
		}).Do(ctx)
	return err
}
