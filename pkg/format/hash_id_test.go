package format_test

import (
	"github.com/projectulterior/2cents-backend/pkg/format"
)

type HashIDType struct{}

func (t HashIDType) IDMethod() format.IDMethod {
	return format.IDMETHOD_HASH
}

func (t HashIDType) Prefix() string {
	return TEST_ID_PREFIX
}

func (t HashIDType) Size() uint {
	return TEST_ID_SIZE
}

type HashID = format.ID[HashIDType]

func NewHashID(item interface{}) HashID {
	return format.NewID(HashIDType{}, item)
}

func ParseHashID(id string) (HashID, error) {
	return format.ParseID(HashIDType{}, id)
}
