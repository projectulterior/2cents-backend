package format

const (
	TRANSACTION_ID_PREFIX = "itrans"
)

type TransactionIDType int

const (
	Deposit FollowIDType = iota + 1
	Withdrawal
	Post
	Like
	Follow
)

func (t TransactionIDType) IDMethod() IDMethod {
	return IDMETHOD_HASH
}

func (t TransactionIDType) Prefix() string {
	return TRANSACTION_ID_PREFIX
}

func (t TransactionIDType) Size() uint {
	return 48
}

type TransactionID string

func NewUpdateTransactionID(userID UserID) TransactionID {
	return TransactionID(NewID(TransactionIDType(0), userID).String())
}

func NewTransferTransactionID(userID UserID) TransactionID {
	return TransactionID(NewID(TransactionIDType(1), userID).String())
}

func NewTransactionIDFromIdentifier(id string) TransactionID {
	return TransactionID(TRANSACTION_ID_PREFIX + id)
}

func ParseTransactionID(id string) (TransactionID, error) {
	parsed, err := ParseID(TransactionIDType(0), id)
	if err != nil {
		return "", err
	}

	return TransactionID(parsed.String()), nil
}

func (u TransactionID) String() string {
	return string(u)
}

func (u TransactionID) Identifier() string {
	return string(u[len(TRANSACTION_ID_PREFIX):])
}
