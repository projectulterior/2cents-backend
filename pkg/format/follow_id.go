package format

const (
	FOLLOW_ID_PREFIX = "ifol"
)

type FollowIDType int

func (t FollowIDType) IDMethod() IDMethod {
	return IDMETHOD_HASH
}

func (t FollowIDType) Prefix() string {
	return FOLLOW_ID_PREFIX
}

func (t FollowIDType) Size() uint {
	return 48
}

type FollowID string

func NewFollowID(followerID, followeeID UserID) FollowID {
	return FollowID(NewID(FollowIDType(0), followerID, followeeID).String())
}

func NewFollowIDFromIdentifier(id string) FollowID {
	return FollowID(FOLLOW_ID_PREFIX + id)
}

func ParseFollowID(id string) (FollowID, error) {
	parsed, err := ParseID(FollowIDType(0), id)
	if err != nil {
		return "", err
	}

	return FollowID(parsed.String()), nil
}

func (u FollowID) String() string {
	return string(u)
}

func (u FollowID) Identifier() string {
	return string(u[len(FOLLOW_ID_PREFIX):])
}
