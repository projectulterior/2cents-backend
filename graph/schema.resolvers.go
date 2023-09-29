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
	"github.com/projectulterior/2cents-backend/pkg/auth"
	"github.com/projectulterior/2cents-backend/pkg/comment_likes"
	"github.com/projectulterior/2cents-backend/pkg/comments"
	"github.com/projectulterior/2cents-backend/pkg/follow"
	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/projectulterior/2cents-backend/pkg/likes"
	"github.com/projectulterior/2cents-backend/pkg/messaging"
	"github.com/projectulterior/2cents-backend/pkg/posts"
	"github.com/projectulterior/2cents-backend/pkg/pubsub/broker"
	"github.com/projectulterior/2cents-backend/pkg/users"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Total is the resolver for the total field.
func (r *centsResolver) Total(ctx context.Context, obj *resolver.Cents) (int, error) {
	panic(fmt.Errorf("not implemented: Total - total"))
}

// Earned is the resolver for the earned field.
func (r *centsResolver) Earned(ctx context.Context, obj *resolver.Cents) (int, error) {
	panic(fmt.Errorf("not implemented: Earned - earned"))
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
		reply, err := r.Follows.CreateFollow(ctx, follow.CreateFollowRequest{
			FollowerID: followerID,
			FolloweeID: followeeID,
		})
		if err != nil {
			return nil, err
		}

		return resolver.NewFollowWithData(r.Services, reply), nil
	}

	followID := format.NewFollowID(followerID, followeeID)

	_, err = r.Follows.DeleteFollow(ctx, follow.DeleteFollowRequest{
		FollowID: followID,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewFollowByID(r.Services, followID), nil
}

// PasswordUpdate is the resolver for the passwordUpdate field.
func (r *mutationResolver) PasswordUpdate(ctx context.Context, old string, new string) (bool, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return false, e(ctx, http.StatusForbidden, err.Error())
	}

	err = r.Auth.UpdatePassword(ctx, auth.UpdatePasswordRequest{
		UserID:      authID,
		OldPassword: old,
		NewPassword: new,
	})
	if err != nil {
		return false, e(ctx, http.StatusInternalServerError, err.Error())
	}

	return true, nil
}

// CentsUpdate is the resolver for the centsUpdate field.
func (r *mutationResolver) CentsUpdate(ctx context.Context, amount int) (*resolver.Cents, error) {
	panic(fmt.Errorf("not implemented: CentsUpdate - centsUpdate"))
}

// CentsTransfer is the resolver for the centsTransfer field.
func (r *mutationResolver) CentsTransfer(ctx context.Context, amount int) (*resolver.Cents, error) {
	panic(fmt.Errorf("not implemented: CentsTransfer - centsTransfer"))
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

// PostLike is the resolver for the postLike field.
func (r *mutationResolver) PostLike(ctx context.Context, id string, isLike bool) (*resolver.Like, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	postID, err := format.ParsePostID(id)
	if err != nil {
		return nil, err
	}

	if isLike {
		reply, err := r.Likes.CreateLike(ctx, likes.CreateLikeRequest{
			PostID:  postID,
			LikerID: authID,
		})
		if err != nil {
			return nil, err
		}

		return resolver.NewLikeWithData(r.Services, reply), nil
	} else {
		likeID := format.NewLikeID(postID, authID)

		_, err = r.Likes.DeleteLike(ctx, likes.DeleteLikeRequest{
			LikeID: likeID,
		})
		if err != nil {
			return nil, err
		}

		return resolver.NewLikeByID(r.Services, likeID), nil
	}
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

// CommentLike is the resolver for the commentLike field.
func (r *mutationResolver) CommentLike(ctx context.Context, id string, isLike bool) (*resolver.CommentLike, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	commentID, err := format.ParseCommentID(id)
	if err != nil {
		return nil, err
	}

	if isLike {
		reply, err := r.CommentLikes.CreateCommentLike(ctx, comment_likes.CreateCommentLikeRequest{
			CommentID: commentID,
			LikerID:   authID,
		})
		if err != nil {
			return nil, err
		}

		return resolver.NewCommentLikeWithData(r.Services, reply), nil
	} else {
		likeID := format.NewCommentLikeID(commentID, authID)

		_, err = r.CommentLikes.DeleteCommentLike(ctx, comment_likes.DeleteCommentLikeRequest{
			CommentLikeID: likeID,
		})
		if err != nil {
			return nil, err
		}

		return resolver.NewCommentLikeByID(r.Services, likeID), nil
	}
}

// ChannelCreate is the resolver for the channelCreate field.
func (r *mutationResolver) ChannelCreate(ctx context.Context, input model.ChannelCreateInput) (*resolver.Channel, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	memberIDs := []format.UserID{authID}
	for _, id := range input.MemberIDs {
		memberID, err := format.ParseUserID(id)
		if err != nil {
			return nil, e(ctx, http.StatusBadRequest, err.Error())
		}
		memberIDs = append(memberIDs, memberID)
	}

	reply, err := r.Messaging.CreateChannel(ctx, messaging.CreateChannelRequest{
		MemberIDs: memberIDs,
	})
	if err != nil {
		return nil, e(ctx, http.StatusInternalServerError, err.Error())
	}

	return resolver.NewChannelWithData(r.Services, reply), nil
}

// ChannelAddMembers is the resolver for the channelAddMembers field.
func (r *mutationResolver) ChannelAddMembers(ctx context.Context, id string, input model.AddMembersInput) (*resolver.Channel, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, err
	}

	channelID, err := format.ParseChannelID(id)
	if err != nil {
		return nil, err
	}

	memberIDs := []format.UserID{}
	for _, id := range input.MemberIDs {
		memberID, err := format.ParseUserID(id)
		if err != nil {
			return nil, e(ctx, http.StatusBadRequest, err.Error())
		}
		memberIDs = append(memberIDs, memberID)
	}

	channel, err := r.Messaging.AddMembers(ctx, messaging.AddMembersRequest{
		ChannelID: channelID,
		MemberID:  authID,
		MemberIDs: memberIDs,
	})
	if err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, e(ctx, http.StatusInternalServerError, err.Error())
		}
		return nil, e(ctx, http.StatusNotFound, err.Error())
	}

	return resolver.NewChannelWithData(r.Services, channel), nil
}

// ChannelDelete is the resolver for the channelDelete field.
func (r *mutationResolver) ChannelDelete(ctx context.Context, id string) (*resolver.Channel, error) {
	channelID, err := format.ParseChannelID(id)
	if err != nil {
		return nil, err
	}

	_, err = r.Messaging.DeleteChannel(ctx, messaging.DeleteChannelRequest{
		ChannelID: channelID,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewChannelByID(r.Services, channelID), nil
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

// MessageUpdate is the resolver for the messageUpdate field.
func (r *mutationResolver) MessageUpdate(ctx context.Context, id string, input model.MessageUpdateInput) (*resolver.Message, error) {
	messageID, err := format.ParseMessageID(id)
	if err != nil {
		return nil, err
	}

	senderID, err := authUserID(ctx)
	if err != nil {
		return nil, err
	}

	message, err := r.Messaging.UpdateMessage(ctx, messaging.UpdateMessageRequest{
		MessageID:   messageID,
		SenderID:    senderID,
		Content:     input.Content,
		ContentType: input.ContentType,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewMessageWithData(r.Services, message), nil
}

// MessageDelete is the resolver for the messageDelete field.
func (r *mutationResolver) MessageDelete(ctx context.Context, id string) (*resolver.Message, error) {
	messageID, err := format.ParseMessageID(id)
	if err != nil {
		return nil, err
	}

	authID, err := authUserID(ctx)
	if err != nil {
		return nil, err
	}

	_, err = r.Messaging.DeleteMessage(ctx, messaging.DeleteMessageRequest{
		MessageID: messageID,
		SenderID:  authID,
	})
	if err != nil {
		return nil, err
	}

	return resolver.NewMessageByID(r.Services, messageID), nil
}

// MessageRead is the resolver for the messageRead field.
func (r *mutationResolver) MessageRead(ctx context.Context, id string) (*resolver.Message, error) {
	panic(fmt.Errorf("not implemented: MessageRead - messageRead"))
}

// ExportUsers is the resolver for the exportUsers field.
func (r *mutationResolver) ExportUsers(ctx context.Context) (bool, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return false, e(ctx, http.StatusForbidden, err.Error())
	}

	err = r.Auth.ExportUsers(ctx)
	if err != nil {
		return false, e(ctx, http.StatusForbidden, err.Error())
	}

	return true, nil
}

// Like is the resolver for the like field.
func (r *postResolver) Like(ctx context.Context, obj *resolver.Post) (*resolver.Like, error) {
	panic(fmt.Errorf("not implemented: Like - like"))
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
func (r *queryResolver) Users(ctx context.Context, page resolver.Pagination) (*resolver.Users, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// Cents is the resolver for the cents field.
func (r *queryResolver) Cents(ctx context.Context, id *string) (*resolver.Cents, error) {
	panic(fmt.Errorf("not implemented: Cents - cents"))
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
func (r *queryResolver) Posts(ctx context.Context, id *string, page resolver.Pagination) (*resolver.Posts, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	if id != nil {
		userID, err := format.ParseUserID(*id)
		if err != nil {
			return nil, e(ctx, http.StatusBadRequest, err.Error())
		}

		return resolver.NewPosts(resolver.NewUserPosts(r.Services, userID, page)), nil
	}

	return resolver.NewPosts(resolver.NewAllPosts(r.Services, page)), nil
}

// PostsFollowing is the resolver for the postsFollowing field.
func (r *queryResolver) PostsFollowing(ctx context.Context, page resolver.Pagination) (*resolver.Posts, error) {
	panic(fmt.Errorf("not implemented: PostsFollowing - postsFollowing"))
}

// PostsDiscovery is the resolver for the postsDiscovery field.
func (r *queryResolver) PostsDiscovery(ctx context.Context, page resolver.Pagination) (*resolver.Posts, error) {
	panic(fmt.Errorf("not implemented: PostsDiscovery - postsDiscovery"))
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
func (r *queryResolver) Comments(ctx context.Context, page resolver.Pagination) (*resolver.Comments, error) {
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
func (r *queryResolver) Likes(ctx context.Context, page resolver.Pagination) (*resolver.Likes, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	return resolver.NewLikes(resolver.NewUserLikes(r.Services, authID, page)), nil
}

// CommentLike is the resolver for the commentLike field.
func (r *queryResolver) CommentLike(ctx context.Context, id string) (*resolver.CommentLike, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	commentLikeID, err := format.ParseCommentLikeID(id)
	if err != nil {
		return nil, e(ctx, http.StatusBadRequest, err.Error())
	}

	return resolver.NewCommentLikeByID(r.Services, commentLikeID), nil
}

// CommentLikes is the resolver for the commentLikes field.
func (r *queryResolver) CommentLikes(ctx context.Context, page resolver.Pagination) (*resolver.CommentLikes, error) {
	panic(fmt.Errorf("not implemented: CommentLikes - commentLikes"))
}

// Follow is the resolver for the follow field.
func (r *queryResolver) Follow(ctx context.Context, id string) (*resolver.Follow, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	followID, err := format.ParseFollowID(id)
	if err != nil {
		return nil, e(ctx, http.StatusBadRequest, err.Error())
	}

	return resolver.NewFollowByID(r.Services, followID), nil
}

// Follows is the resolver for the follows field.
func (r *queryResolver) Follows(ctx context.Context, page resolver.Pagination) (*resolver.Follows, error) {
	panic(fmt.Errorf("not implemented: Follows - follows"))
}

// Channel is the resolver for the channel field.
func (r *queryResolver) Channel(ctx context.Context, id string) (*resolver.Channel, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	channelID, err := format.ParseChannelID(id)
	if err != nil {
		return nil, e(ctx, http.StatusBadRequest, err.Error())
	}

	return resolver.NewChannelByID(r.Services, channelID), nil
}

// ChannelByMembers is the resolver for the channelByMembers field.
func (r *queryResolver) ChannelByMembers(ctx context.Context, members []string) (*resolver.Channel, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	memberIDs := []format.UserID{authID}

	for _, mid := range members {
		memberID, err := format.ParseUserID(mid)
		if err != nil {
			return nil, e(ctx, http.StatusBadRequest, err.Error())
		}

		memberIDs = append(memberIDs, memberID)
	}

	channel, err := r.Messaging.GetChannelByMembers(ctx, messaging.GetChannelByMembersRequest{
		MemberIDs: memberIDs,
	})
	if err != nil {
		if status.Code(err) != codes.NotFound {
			return nil, e(ctx, http.StatusInternalServerError, err.Error())
		}
		return nil, e(ctx, http.StatusNotFound, err.Error())
	}

	return resolver.NewChannelWithData(r.Services, channel), nil
}

// Channels is the resolver for the channels field.
func (r *queryResolver) Channels(ctx context.Context, page resolver.Pagination) (*resolver.Channels, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	return resolver.NewChannels(resolver.NewUserChannels(r.Services, authID, page)), nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context, id string, page resolver.Pagination) (*resolver.Messages, error) {
	panic(fmt.Errorf("not implemented: Messages - messages"))
}

// Notifications is the resolver for the notifications field.
func (r *queryResolver) Notifications(ctx context.Context, page *resolver.Pagination) (*model.Notifications, error) {
	panic(fmt.Errorf("not implemented: Notifications - notifications"))
}

// SearchUsers is the resolver for the searchUsers field.
func (r *queryResolver) SearchUsers(ctx context.Context, query string, page resolver.Pagination) (*resolver.Users, error) {
	_, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	return resolver.NewUsers(resolver.NewSearchUsers(r.Services, query, page)), nil
}

// OnUserUpdated is the resolver for the onUserUpdated field.
func (r *subscriptionResolver) OnUserUpdated(ctx context.Context, id *string) (<-chan *resolver.User, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	ch := make(chan *resolver.User)

	go func() {
		listener := broker.Exchange(users.UserUpdatedEvent{}).Listener()
		defer listener.Close(ctx)

		for {
			event, err := listener.Next(ctx)
			if err != nil {
				r.Error("error in listener", zap.Error(err))
				return
			}

			if event.User.UserID == authID {
				ch <- resolver.NewUserWithData(r.Services, &event.User)
			}
		}
	}()

	return ch, nil
}

// OnChannelUpdated is the resolver for the onChannelUpdated field.
func (r *subscriptionResolver) OnChannelUpdated(ctx context.Context) (<-chan *resolver.Channel, error) {
	authID, err := authUserID(ctx)
	if err != nil {
		return nil, e(ctx, http.StatusForbidden, err.Error())
	}

	ch := make(chan *resolver.Channel)

	go func() {
		listener := broker.Exchange(messaging.ChannelUpdatedEvent{}).Listener()
		defer listener.Close(ctx)

		for {
			event, err := listener.Next(ctx)
			if err != nil {
				r.Error("error in listener", zap.Error(err))
				return
			}

			for _, memberID := range event.Channel.MemberIDs {
				if memberID == authID {
					ch <- resolver.NewChannelWithData(r.Services, &event.Channel)
				}
			}
		}
	}()

	return ch, nil
}

// Cents returns CentsResolver implementation.
func (r *Resolver) Cents() CentsResolver { return &centsResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type centsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
