package format

import (
	"github.com/projectulterior/2cents-backend/pkg/utils"
)

// UniqueID implementation from firestore
func UniqueID() string {
	return utils.Random(20)
}

// Hash hashes the argument
//
// It will panic if the hashing function
// errors
func Hash(toHash interface{}) string {
	identifier, err := utils.SHA256Base64(toHash)
	if err != nil {
		panic(err)
	}

	return identifier
}
