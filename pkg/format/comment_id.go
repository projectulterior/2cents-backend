package format

const (
	COMMENT_ID_PREFIX = "icom"
)

type CommentIDType int

func (id CommentIDType) IDMethod() IDMethod {
	return IDMETHOD_RANDOM
}

func (id CommentIDType) Prefix() string {
	return COMMENT_ID_PREFIX
}

func (id CommentIDType) Size() uint {
	return 48
}

type CommentID string

func NewCommentID() CommentID {
	return CommentID(NewID(CommentIDType(0)).String())
}

func NewCommentIDFromIdentifier(id string) CommentID {
	return CommentID(COMMENT_ID_PREFIX + id)
}

func ParseCommentID(id string) (CommentID, error) {
	parsed, err := ParseID(CommentIDType(0), id)
	if err != nil {
		return "", err
	}

	return CommentID(parsed.String()), nil
}

func (u CommentID) String() string {
	return string(u)
}

func (u CommentID) Identifier() string {
	return string(u[4:])
}
