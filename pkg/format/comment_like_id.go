package format

const (
	COMMENT_LIKE_ID_PREFIX = "ilke"
)

type CommentLikeIDType int

func (t CommentLikeIDType) IDMethod() IDMethod {
	return IDMETHOD_HASH
}

func (t CommentLikeIDType) Prefix() string {
	return COMMENT_LIKE_ID_PREFIX
}

func (t CommentLikeIDType) Size() uint {
	return 48
}

type CommentLikeID string

func NewCommentLikeID(commentID CommentID, likerID UserID) CommentLikeID {
	return CommentLikeID(NewID(CommentLikeIDType(0), commentID, likerID).String())
}

func NewCommentLikeIDFromIdentifier(id string) CommentLikeID {
	return CommentLikeID(COMMENT_LIKE_ID_PREFIX + id)
}

func ParseCommentLikeID(id string) (CommentLikeID, error) {
	parsed, err := ParseID(CommentLikeIDType(0), id)
	if err != nil {
		return "", err
	}

	return CommentLikeID(parsed.String()), nil
}

func (c CommentLikeID) String() string {
	return string(c)
}

func (c CommentLikeID) Identifier() string {
	return string(c[len(COMMENT_LIKE_ID_PREFIX):])
}
