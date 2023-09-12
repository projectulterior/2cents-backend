package resolver

// import (
// 	"context"

// 	"github.com/projectulterior/2cents-backend/pkg/format"
// 	"github.com/projectulterior/2cents-backend/pkg/messaging"
// 	"github.com/projectulterior/2cents-backend/pkg/services"
// )

// type UserChannels struct {
// 	svc *services.Services
// 	getter[*messaging.GetChannelsResponse, func(context.Context) (*messaging.GetChannelsResponse, error)]
// }

// func NewUserChannels(svc *services.Services, userID format.UserID, page Pagination) *UserChannels {
// 	return &UserChannels{
// 		svc: svc,
// 		getter: NewGetter(
// 			func(ctx context.Context) (*messaging.GetChannelResponse, error) {
// 				return svc.Messaging.GetChannels(ctx, &messaging.GetChannelsRequest{

// 				})
// 			}
// 		)
// 	}
// }
