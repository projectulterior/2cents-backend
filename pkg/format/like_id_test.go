package format_test

import (
	"testing"

	"github.com/projectulterior/2cents-backend/pkg/format"
	"github.com/stretchr/testify/assert"
)

func TestLikeID(t *testing.T) {
	postID := format.NewPostID()
	likerID1 := format.NewUserID()
	likerID2 := format.NewUserID()

	likeID1 := format.NewLikeID(postID, likerID1)
	likeID2 := format.NewLikeID(postID, likerID2)
	assert.NotEqual(t, likeID1, likeID2)

	sameID1 := format.NewLikeID(postID, likerID1)
	assert.Equal(t, likeID1, sameID1)
}
