package format

const (
	POST_ID_PREFIX = "ipst"
)

type PostIDType int

func (id PostIDType) IDMethod() IDMethod {
	return IDMETHOD_RANDOM
}

func (id PostIDType) Prefix() string {
	return POST_ID_PREFIX
}

func (id PostIDType) Size() uint {
	return 48
}

type PostID string

func NewPostID() PostID {
	return PostID(NewID(PostIDType(0)).String())
}

func NewPostIDFromIdentifier(id string) PostID {
	return PostID(POST_ID_PREFIX + id)
}

func ParsePostID(id string) (PostID, error) {
	parsed, err := ParseID(PostIDType(0), id)
	if err != nil {
		return "", err
	}

	return PostID(parsed.String()), nil
}

func (u PostID) String() string {
	return string(u)
}

func (u PostID) Identifier() string {
	return string(u[4:])
}
