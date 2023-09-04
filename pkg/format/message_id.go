package format

const (
	MESSAGE_ID_PREFIX = "imsg"
)

type MessageIDType int

func (id MessageIDType) IDMethod() IDMethod {
	return IDMETHOD_RANDOM
}

func (id MessageIDType) Prefix() string {
	return MESSAGE_ID_PREFIX
}

func (id MessageIDType) Size() uint {
	return 48
}

type MessageID string

func NewMessageID() MessageID {
	return MessageID(NewID(MessageIDType(0)).String())
}

func NewMessageIDFromIdentifier(id string) MessageID {
	return MessageID(MESSAGE_ID_PREFIX + id)
}

func ParseMessageID(id string) (MessageID, error) {
	parsed, err := ParseID(MessageIDType(0), id)
	if err != nil {
		return "", err
	}

	return MessageID(parsed.String()), nil
}

func (u MessageID) String() string {
	return string(u)
}

func (u MessageID) Identifier() string {
	return string(u[4:])
}
