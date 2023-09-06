package resolver

import (
	"context"
	"time"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/services"
)

type Like struct {
	svc *services.Services

	likeID format.LikeID
	getter[*likes.Like, func(context.Context) (*likes.Like, error)]
}

func NewLikeByID(svc *services.Services, likeID format.LikeID) *Like {
	return &Like{
		svc:    svc,
		likeID: likeID,
		getter: NewGetter(
			func(ctx context.Context) (*likes.Like, error) {
				return svc.Likes.GetLike(ctx, likes.GetLikeRequest{
					LikeID: likeID,
				})
			},
		),
	}
}

func NewLikeWithData(svc *services.Services, data *likes.Like) *Like {
	return &Like{
		svc:    svc,
		likeID: data.LikeID,
		getter: NewGetter(
			func(ctx context.Context) (*likes.Like, error) {
				return data, nil
			},
		),
	}
}

func (l *Like) ID(ctx context.Context) (string, error) {
	return l.likeID.String(), nil
}

func (l *Like) Post(ctx context.Context) (*Post, error) {
	reply, err := l.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewPostByID(l.svc, reply.PostID), nil
}

func (l *Like) Liker(ctx context.Context) (*User, error) {
	reply, err := l.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return NewUserByID(l.svc, reply.LikerID), nil
}

func (l *Like) CreatedAt(ctx context.Context) (*time.Time, error) {
	reply, err := l.getter.Call(ctx)
	if err != nil {
		return nil, err
	}

	return &reply.CreatedAt, nil
}
