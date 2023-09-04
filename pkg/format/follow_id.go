package format

const (
	FOLLOW_ID_PREFIX = "ifol"
)

type FollowIDType struct{}

func (t FollowIDType) IDMethod() IDMethod {
	return IDMETHOD_HASH
}

func (t FollowIDType) Prefix() string {
	return FOLLOW_ID_PREFIX
}

func (t FollowIDType) Size() uint {
	return 48
}

type FollowID = ID[FollowIDType]

func NewFollowID(followerID, followeeID UserID) FollowID {
	return NewID(FollowIDType{}, followerID, followeeID)
}

func ParseFollowID(id string) (FollowID, error) {
	return ParseID(FollowIDType{}, id)
}
