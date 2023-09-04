package format

const (
	CHANNEL_ID_PREFIX = "ichn"
)

type ChannelIDType int

func (id ChannelIDType) IDMethod() IDMethod {
	return IDMETHOD_RANDOM
}

func (id ChannelIDType) Prefix() string {
	return CHANNEL_ID_PREFIX
}

func (id ChannelIDType) Size() uint {
	return 48
}

type ChannelID string

func NewChannelID() ChannelID {
	return ChannelID(NewID(ChannelIDType(0)).String())
}

func NewChannelIDFromIdentifier(id string) ChannelID {
	return ChannelID(CHANNEL_ID_PREFIX + id)
}

func ParseChannelID(id string) (ChannelID, error) {
	parsed, err := ParseID(ChannelIDType(0), id)
	if err != nil {
		return "", err
	}

	return ChannelID(parsed.String()), nil
}

func (u ChannelID) String() string {
	return string(u)
}

func (u ChannelID) Identifier() string {
	return string(u[4:])
}
