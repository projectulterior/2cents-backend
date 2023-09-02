package format_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_ID_PREFIX = "itst"
	TEST_ID_SIZE   = 10
)

func TestID(t *testing.T) {
	randomID := NewRandomID()
	hashID := NewHashID("helloworld")

	t.Run("parse id", func(t *testing.T) {
		t.Run("random id", func(t *testing.T) {
			parsedID, err := ParseRandomID(randomID.String())
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, randomID, parsedID)
		})

		t.Run("hash id", func(t *testing.T) {
			parsedID, err := ParseHashID(hashID.String())
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, hashID, parsedID)
		})
	})

	t.Run("json", func(t *testing.T) {
		t.Run("random id", func(t *testing.T) {
			type Body struct {
				TestID RandomID `json:"test_id"`
			}

			expectedBody := Body{TestID: randomID}
			bytes, err := json.Marshal(expectedBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, fmt.Sprintf("{\"test_id\":%q}", randomID.String()), string(bytes))

			var actualBody Body
			err = json.Unmarshal(bytes, &actualBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expectedBody, actualBody)
		})

		t.Run("hash id", func(t *testing.T) {
			type Body struct {
				TestID HashID `json:"test_id"`
			}

			expectedBody := Body{TestID: hashID}
			bytes, err := json.Marshal(expectedBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, fmt.Sprintf("{\"test_id\":%q}", hashID.String()), string(bytes))

			var actualBody Body
			err = json.Unmarshal(bytes, &actualBody)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expectedBody, actualBody)
		})
	})
}
