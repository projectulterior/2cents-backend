package services

import (
	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/cents"
	"github.com/projectulterior/2cents-backend/pkg/comment_likes"
	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/search"
	"github.com/projectulterior/2cents-backend/pkg/users"
)

type Services struct {
	Auth         *auth.Service
	Comments     *comments.Service
	CommentLikes *comment_likes.Service
	Follows      *follow.Service
	Likes        *likes.Service
	Messaging    *messaging.Service
	Posts        *posts.Service
	Cents        *cents.Service
	Search       *search.Service
	Users        *users.Service
}
