package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"
	"net/http"

	"github.com/projectulterior/2cents-backend/graph/model"
	"github.com/projectulterior/2cents-backend/graph/resolver"
	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/users"
)

// ID is the resolver for the id field.
func (r *channelResolver) ID(ctx context.Context, obj *resolver.Channel) (string, error) {
	panic(fmt.Errorf("not implemented: ID - id"))
}

// Members is the resolver for the members field.
func (r *channelResolver) Members(ctx context.Context, obj *resolver.Channel) ([]*resolver.User, error) {
	panic(fmt.Errorf("not implemented: Members - members"))
}

// Messages is the resolver for the messages field.
func (r *channelResolver) Messages(ctx context.Context, obj *resolver.Channel, page model.Pagination) (*model.Messages, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// UserUpdate is the resolver for the userUpdate field.
func (r *mutationResolver) UserUpdate(ctx context.Context, input model.UserUpdateInput) (*resolver.User, error) {
	userID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	user, err := r.Users.UpdateUser(ctx, users.UpdateUserRequest{
		UserID:   userID,
		Name:     input.Name,
		Email:    input.Email,
		Bio:      input.Bio,
		Birthday: input.Birthday,
		Profile:  input.Profile,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewUserWithData(r.Services, user), nil
}

// UserDelete is the resolver for the userDelete field.
func (r *mutationResolver) UserDelete(ctx context.Context) (*resolver.User, error) {
	userID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	_, err = r.Users.DeleteUser(ctx, users.DeleteUserRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewUserByID(r.Services, userID), nil
}

// UserFollow is the resolver for the userFollow field.
func (r *mutationResolver) UserFollow(ctx context.Context, id string, isFollow bool) (*resolver.Follow, error) {
	followerID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	followeeID, err := format.ParseUserID(id)
	if err != nil {
		return nil, err
	}

	if isFollow {
		reply, err := r.Services.Follows.CreateFollow(ctx, follow.CreateFollowRequest{
			FollowerID: followerID,
			FolloweeID: followeeID,
		})
		if err != nil {
			return nil, err
		}

		return resolver.NewFollowWithData(r.Services, reply), nil
	}

	followID := format.NewFollowID(followerID, followeeID)

	_, err = r.Services.Follows.DeleteFollow(ctx, follow.DeleteFollowRequest{
		FollowID: followID,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewFollowByID(r.Services, followID), nil
}

// PostCreate is the resolver for the postCreate field.
func (r *mutationResolver) PostCreate(ctx context.Context, input model.PostCreateInput) (*resolver.Post, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	reply, err := r.Posts.CreatePost(ctx, posts.CreatePostRequest{
		AuthorID:    authID,
		Visibility:  input.Visibility,
		Content:     input.Content,
		ContentType: input.ContentType,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewPostWithData(r.Services, reply), nil
}

// PostUpdate is the resolver for the postUpdate field.
func (r *mutationResolver) PostUpdate(ctx context.Context, id string, input model.PostUpdateInput) (*resolver.Post, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, err
	}

	postID, err := format.ParsePostID(id)
	if err != nil {
		return nil, err
	}

	post, err := r.Posts.UpdatePost(ctx, posts.UpdatePostRequest{
		PostID:      postID,
		AuthorID:    authID,
		Visibility:  input.Visibility,
		Content:     input.Content,
		ContentType: input.ContentType,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewPostWithData(r.Services, post), nil
}

// PostDelete is the resolver for the postDelete field.
func (r *mutationResolver) PostDelete(ctx context.Context, id string) (*resolver.Post, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	postID, err := format.ParsePostID(id)
	if err != nil {
		return nil, err
	}

	_, err = r.Posts.DeletePost(ctx, posts.DeletePostRequest{
		PostID:   postID,
		AuthorID: authID,
	})
	if err != nil {
		return nil, err
	}
	return resolver.NewPostByID(r.Services, postID), nil
}

// CommentCreate is the resolver for the commentCreate field.
func (r *mutationResolver) CommentCreate(ctx context.Context, input model.CommentCreateInput) (*resolver.Comment, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	postID, err := format.ParsePostID(input.PostID)
	if err != nil {
		return nil, e(ctx, http.StatusBadRequest, err.Error())
	}

	reply, err := r.Comments.CreateComment(ctx, comments.CreateCommentRequest{
		AuthorID: authID,
		PostID:   postID,
		Content:  input.Content,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewCommentWithData(r.Services, reply), nil
}

// CommentUpdate is the resolver for the commentUpdate field.
func (r *mutationResolver) CommentUpdate(ctx context.Context, id string, input model.CommentUpdateInput) (*resolver.Comment, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, err
	}

	commentID, err := format.ParseCommentID(id)
	if err != nil {
		return nil, err
	}

	comment, err := r.Comments.UpdateComment(ctx, comments.UpdateCommentRequest{
		CommentID: commentID,
		AuthorID:  authID,
		Content:   *input.Content,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewCommentWithData(r.Services, comment), nil
}

// CommentDelete is the resolver for the commentDelete field.
func (r *mutationResolver) CommentDelete(ctx context.Context, id string) (*resolver.Comment, error) {
	commentID, err := format.ParseCommentID(id)
	if err != nil {
		return nil, err
	}

	authID, err := authUserID(ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.Comments.DeleteComment(ctx, comments.DeleteCommentRequest{
		CommentID: commentID,
		DeleterID: authID,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewCommentByID(r.Services, commentID), nil
}

// LikeCreate is the resolver for the likeCreate field.
func (r *mutationResolver) LikeCreate(ctx context.Context, id string) (*resolver.Like, error) {
	postID, err := format.ParsePostID(id)
	if err != nil {
		return nil, err
	}

	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	reply, err := r.Likes.CreateLike(ctx, likes.CreateLikeRequest{
		PostID:  postID,
		LikerID: authID,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewLikeWithData(r.Services, reply), nil
}

// LikeDelete is the resolver for the likeDelete field.
func (r *mutationResolver) LikeDelete(ctx context.Context, id string) (*resolver.Like, error) {
	likeID, err := format.ParseLikeID(id)
	if err != nil {
		return nil, err
	}

	_, err = r.Likes.DeleteLike(ctx, likes.DeleteLikeRequest{
		LikeID: likeID,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewLikeByID(r.Services, likeID), nil
}

// MessageCreate is the resolver for the messageCreate field.
func (r *mutationResolver) MessageCreate(ctx context.Context, input model.MessageCreateInput) (*resolver.Message, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	channelID, err := format.ParseChannelID(input.ChannelID)
	if err != nil {
		return nil, err
	}

	reply, err := r.Messaging.CreateMessage(ctx, messaging.CreateMessageRequest{
		ChannelID:   channelID,
		SenderID:    authID,
		Content:     *input.Content,
		ContentType: *input.ContentType,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewMessageByID(r.Services, reply.MessageID), nil
}

// MessageDelete is the resolver for the messageDelete field.
func (r *mutationResolver) MessageDelete(ctx context.Context, id string) (*resolver.Comment, error) {
	panic(fmt.Errorf("not implemented: MessageDelete - messageDelete"))
}

// Likes is the resolver for the likes field.
func (r *postResolver) Likes(ctx context.Context, obj *resolver.Post, page model.Pagination) (*model.Likes, error) {
	panic(fmt.Errorf("not implemented: Likes - likes"))
}

// Comments is the resolver for the comments field.
func (r *postResolver) Comments(ctx context.Context, obj *resolver.Post, page model.Pagination) (*model.Comments, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *string) (*resolver.User, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	if id != nil {
		// not their own user, but someone else
		userID, err := format.ParseUserID(*id)
		if err != nil {
			return nil, e(ctx, http.StatusBadRequest, err.Error())
		}

		return resolver.NewUserByID(r.Services, userID), nil
	}

	// their own user
	return resolver.NewMyUser(r.Services, authID), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, page model.Pagination) (*model.Users, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*resolver.Post, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	postID, err := format.ParsePostID(id)
	if err != nil {
		return nil, e(ctx, http.StatusBadRequest, err.Error())
	}

	return resolver.NewPostByID(r.Services, postID), nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, page model.Pagination) (*model.Posts, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Comment is the resolver for the comment field.
func (r *queryResolver) Comment(ctx context.Context, id string) (*resolver.Comment, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	commentID, err := format.ParseCommentID(id)
	if err != nil {
		return nil, e(ctx, http.StatusBadRequest, err.Error())
	}

	return resolver.NewCommentByID(r.Services, commentID), nil
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, page model.Pagination) (*model.Comments, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Like is the resolver for the like field.
func (r *queryResolver) Like(ctx context.Context, id string) (*resolver.Like, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	likeID, err := format.ParseLikeID(id)
	if err != nil {
		return nil, e(ctx, http.StatusBadRequest, err.Error())
	}

	return resolver.NewLikeByID(r.Services, likeID), nil
}

// Likes is the resolver for the likes field.
func (r *queryResolver) Likes(ctx context.Context, page model.Pagination) (*model.Likes, error) {
	panic(fmt.Errorf("not implemented: Likes - likes"))
}

// Channel is the resolver for the channel field.
func (r *queryResolver) Channel(ctx context.Context, id string) (*resolver.Channel, error) {
	panic(fmt.Errorf("not implemented: Channel - channel"))
}

// ChannelByMembers is the resolver for the channelByMembers field.
func (r *queryResolver) ChannelByMembers(ctx context.Context, members []string) (*resolver.Channel, error) {
	panic(fmt.Errorf("not implemented: ChannelByMembers - channelByMembers"))
}

// Message is the resolver for the message field.
func (r *queryResolver) Message(ctx context.Context, id string) (*resolver.Message, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	messageID, err := format.ParseMessageID(id)
	if err != nil {
		return nil, e(ctx, http.StatusBadRequest, err.Error())
	}

	return resolver.NewMessageByID(r.Services, messageID), nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context, page model.Pagination) (*model.Messages, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// OnUserUpdated is the resolver for the onUserUpdated field.
func (r *subscriptionResolver) OnUserUpdated(ctx context.Context, id *string) (<-chan *resolver.User, error) {
	panic(fmt.Errorf("not implemented: OnUserUpdated - onUserUpdated"))
}

// Email is the resolver for the email field.
func (r *userResolver) Email(ctx context.Context, obj *resolver.User) (*string, error) {
	panic(fmt.Errorf("not implemented: Email - email"))
}

// Birthday is the resolver for the birthday field.
func (r *userResolver) Birthday(ctx context.Context, obj *resolver.User) (*format.Birthday, error) {
	panic(fmt.Errorf("not implemented: Birthday - birthday"))
}

// Cents is the resolver for the cents field.
func (r *userResolver) Cents(ctx context.Context, obj *resolver.User) (*model.Cents, error) {
	panic(fmt.Errorf("not implemented: Cents - cents"))
}

// Follows is the resolver for the follows field.
func (r *userResolver) Follows(ctx context.Context, obj *resolver.User, page *model.Pagination) (*model.Follows, error) {
	panic(fmt.Errorf("not implemented: Follows - follows"))
}

// Posts is the resolver for the posts field.
func (r *userResolver) Posts(ctx context.Context, obj *resolver.User, page *model.Pagination) (*model.Posts, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Likes is the resolver for the likes field.
func (r *userResolver) Likes(ctx context.Context, obj *resolver.User, page *model.Pagination) (*model.Likes, error) {
	panic(fmt.Errorf("not implemented: Likes - likes"))
}

// Channel returns ChannelResolver implementation.
func (r *Resolver) Channel() ChannelResolver { return &channelResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type channelResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
