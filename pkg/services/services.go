package services

import (
	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/comment_likes"
	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/users"
)

type Services struct {
	Auth         *auth.Service
	Users        *users.Service
	Posts        *posts.Service
	Comments     *comments.Service
	Likes        *likes.Service
	CommentLikes *comment_likes.Service
	Follows      *follow.Service
	Messaging    *messaging.Service
}
