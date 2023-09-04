package format

const (
	LIKE_ID_PREFIX = "ilke"
)

type LikeIDType int

func (t LikeIDType) IDMethod() IDMethod {
	return IDMETHOD_HASH
}

func (t LikeIDType) Prefix() string {
	return LIKE_ID_PREFIX
}

func (t LikeIDType) Size() uint {
	return 48
}

type LikeID string

func NewLikeID(postID PostID, likerID UserID) LikeID {
	return LikeID(NewID(LikeIDType(0), postID, likerID).String())
}

func NewLikeIDFromIdentifier(id string) LikeID {
	return LikeID(LIKE_ID_PREFIX + id)
}

func ParseLikeID(id string) (LikeID, error) {
	parsed, err := ParseID(LikeIDType(0), id)
	if err != nil {
		return "", err
	}

	return LikeID(parsed.String()), nil
}

func (u LikeID) String() string {
	return string(u)
}

func (u LikeID) Identifier() string {
	return string(u[len(LIKE_ID_PREFIX):])
}
