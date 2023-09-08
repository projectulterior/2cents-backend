// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/projectulterior/2cents-backend/graph/resolver"
	"github.com/projectulterior/2cents-backend/pkg/format"
)

type AddMembersInput struct {
	ChannelID string   `json:"channelID"`
	MemberID  string   `json:"memberID"`
	MemberIDs []string `json:"memberIDs"`
}

type Cents struct {
	Total     int `json:"total"`
	Deposited int `json:"deposited"`
	Earned    int `json:"earned"`
	Given     int `json:"given"`
}

type ChannelCreateInput struct {
	MemberIDs []string `json:"memberIDs"`
}

type CommentCreateInput struct {
	PostID      string             `json:"postID"`
	Content     string             `json:"content"`
	ContentType format.ContentType `json:"contentType"`
}

type CommentUpdateInput struct {
	Content     *string             `json:"content,omitempty"`
	ContentType *format.ContentType `json:"contentType,omitempty"`
}

type Comments struct {
	Comments []*resolver.Comment `json:"comments"`
	Next     string              `json:"next"`
}

type Follows struct {
	Follows []*resolver.Follow `json:"follows"`
	Next    *string            `json:"next,omitempty"`
}

type Likes struct {
	Likes []*resolver.Like `json:"likes"`
	Next  string           `json:"next"`
}

type MessageCreateInput struct {
	ChannelID   string              `json:"channelID"`
	SenderID    string              `json:"senderID"`
	Content     *string             `json:"content,omitempty"`
	ContentType *format.ContentType `json:"contentType,omitempty"`
}

type Messages struct {
	Messages []*resolver.Message `json:"messages,omitempty"`
	Next     *string             `json:"next,omitempty"`
}

type Pagination struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}

type PostCreateInput struct {
	Visibility  format.Visibility  `json:"visibility"`
	Content     string             `json:"content"`
	ContentType format.ContentType `json:"contentType"`
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
	Profile  *string          `json:"profile,omitempty"`
	Cover    *string          `json:"cover,omitempty"`
}

type Users struct {
	Users []*resolver.User `json:"users"`
	Next  string           `json:"next"`
}
