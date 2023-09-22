package search

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/projectulterior/2cents-backend/pkg/auth"
)

func (s *Service) ProcessUsernameUpdated(ctx context.Context, event auth.UserUpdatedEvent) error {
	b, err := json.Marshal(User{
		UserID:   event.User.UserID,
		Username: event.User.Username,
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
