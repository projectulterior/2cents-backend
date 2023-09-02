package format

const (
	TOKEN_ID_PREFIX = "itoken"
)

type TokenIDType int

func (id TokenIDType) IDMethod() IDMethod {
	return IDMETHOD_RANDOM
}

func (id TokenIDType) Prefix() string {
	return TOKEN_ID_PREFIX
}

func (id TokenIDType) Size() uint {
	return 48
}

type TokenID string

func NewTokenID() TokenID {
	return TokenID(NewID(TokenIDType(0)).String())
}

func NewTokenIDFromIdentifier(id string) TokenID {
	return TokenID(TOKEN_ID_PREFIX + id)
}

func ParseTokenID(id string) (TokenID, error) {
	parsed, err := ParseID(TokenIDType(0), id)
	if err != nil {
		return "", err
	}

	return TokenID(parsed.String()), nil
}

func (u TokenID) String() string {
	return string(u)
}

func (u TokenID) Identifier() string {
	return string(u[4:])
}
