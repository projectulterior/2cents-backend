package format_test

import (
	"github.com/projectulterior/2cents-backend/pkg/format"
)

type RandomIDType struct{}

func (t RandomIDType) IDMethod() format.IDMethod {
	return format.IDMETHOD_RANDOM
}

func (t RandomIDType) Prefix() string {
	return TEST_ID_PREFIX
}

func (t RandomIDType) Size() uint {
	return TEST_ID_SIZE
}

type RandomID = format.ID[RandomIDType]

func NewRandomID() RandomID {
	return format.NewID(RandomIDType{})
}

func ParseRandomID(id string) (RandomID, error) {
	return format.ParseID(RandomIDType{}, id)
}
