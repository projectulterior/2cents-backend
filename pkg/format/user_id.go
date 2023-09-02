package format

const (
	USER_ID_PREFIX   = "iusr"
	ADMIN_ID_PREFIX  = USER_ID_PREFIX + "@"
	DEFAULT_ADMIN_ID = UserID(ADMIN_ID_PREFIX)
)

type UserIDType int

func (id UserIDType) IDMethod() IDMethod {
	return IDMETHOD_RANDOM
}

func (id UserIDType) Prefix() string {
	return USER_ID_PREFIX
}

func (id UserIDType) Size() uint {
	return 48
}

type UserID string

func NewUserID() UserID {
	return UserID(NewID(UserIDType(0)).String())
}

func NewAdminID(name string) UserID {
	return UserID(ADMIN_ID_PREFIX + name)
}

func NewUserIDFromIdentifier(id string) UserID {
	return UserID(USER_ID_PREFIX + id)
}

func ParseUserID(id string) (UserID, error) {
	parsed, err := ParseID(UserIDType(0), id)
	if err != nil {
		return "", err
	}

	return UserID(parsed.String()), nil
}

func (u UserID) String() string {
	return string(u)
}

func (u UserID) Identifier() string {
	return string(u[4:])
}

func (u UserID) IsAdmin() bool {
	return u[:len(ADMIN_ID_PREFIX)] == ADMIN_ID_PREFIX
}
