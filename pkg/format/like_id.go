package format

const (
	LIKE_ID_PREFIX = "ilke"
)

type LikeIDType struct{}

func (t LikeIDType) IDMethod() IDMethod {
	return IDMETHOD_HASH
}

func (t LikeIDType) Prefix() string {
	return LIKE_ID_PREFIX
}

func (t LikeIDType) Size() uint {
	return 48
}

type LikeID = ID[LikeIDType]

func NewLikeID(postID PostID, likerID UserID) LikeID {
	return NewID(LikeIDType{}, postID, likerID)
}

func ParseLikeID(id string) (LikeID, error) {
	return ParseID(LikeIDType{}, id)
}
