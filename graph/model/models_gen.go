// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/projectulterior/2cents-backend/graph/resolver"
	"github.com/projectulterior/2cents-backend/pkg/format"
)

type Cents struct {
	Total     int `json:"total"`
	Deposited int `json:"deposited"`
	Earned    int `json:"earned"`
	Given     int `json:"given"`
}

type Channel struct {
	ID       string           `json:"id"`
	Members  []*resolver.User `json:"members,omitempty"`
	Messages *Messages        `json:"messages,omitempty"`
}

type Comment struct {
	ID        string         `json:"id"`
	Post      *resolver.Post `json:"post,omitempty"`
	Content   *string        `json:"content,omitempty"`
	Author    *resolver.User `json:"author,omitempty"`
	CreatedAt *time.Time     `json:"createdAt,omitempty"`
}

type CommentCreateInput struct {
	Content     *string             `json:"content,omitempty"`
	ContentType *format.ContentType `json:"contentType,omitempty"`
}

type CommentUpdateInput struct {
	Content     *string             `json:"content,omitempty"`
	ContentType *format.ContentType `json:"contentType,omitempty"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
	Next     string     `json:"next"`
}

type Follow struct {
	ID        string         `json:"id"`
	Follower  *resolver.User `json:"follower,omitempty"`
	Followee  *resolver.User `json:"followee,omitempty"`
	CreatedAt *time.Time     `json:"createdAt,omitempty"`
}

type Follows struct {
	Follows []*Follow `json:"follows"`
	Next    *string   `json:"next,omitempty"`
}

type Like struct {
	ID        string         `json:"id"`
	Post      *resolver.Post `json:"post,omitempty"`
	Liker     *resolver.User `json:"liker,omitempty"`
	CreatedAt *time.Time     `json:"createdAt,omitempty"`
}

type Likes struct {
	Likes []*Like `json:"likes"`
	Next  string  `json:"next"`
}

type Message struct {
	Content     *string             `json:"content,omitempty"`
	ContentType *format.ContentType `json:"contentType,omitempty"`
	CreatedAt   *time.Time          `json:"createdAt,omitempty"`
	Sender      *resolver.User      `json:"sender,omitempty"`
}

type Messages struct {
	Messages []*Message `json:"messages,omitempty"`
	Next     *string    `json:"next,omitempty"`
}

type Pagination struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

type PostCreateInput struct {
	Visibility  *format.Visibility  `json:"visibility,omitempty"`
	Content     *string             `json:"content,omitempty"`
	ContentType *format.ContentType `json:"contentType,omitempty"`
}

type PostUpdateInput struct {
	Visibility  *format.Visibility  `json:"visibility,omitempty"`
	Content     *string             `json:"content,omitempty"`
	ContentType *format.ContentType `json:"contentType,omitempty"`
}

type Posts struct {
	Posts []*resolver.Post `json:"posts"`
	Next  string           `json:"next"`
}

type UserUpdateInput struct {
	Name     *string          `json:"name,omitempty"`
	Email    *string          `json:"email,omitempty"`
	Bio      *string          `json:"bio,omitempty"`
	Birthday *format.Birthday `json:"birthday,omitempty"`
}

type Users struct {
	Users []*resolver.User `json:"users"`
	Next  string           `json:"next"`
}
