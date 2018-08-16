package chained

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
)

// HashStringSize represents the length of the hashes
// that will be generated in their string format.
const HashStringSize = sha512.Size * 2

// Global settings are defined below and should only
// be modified before any blockchains are created,
// and before any mining occurs in general.
var (
	// Target represents the mining difficulty of the
	// blockchain. If the target is changed after a
	// blockchain object is created, all previous blocks
	// mined with the former target will be considered
	// invalid.
	Target = 4

	// IDLength is the valid length of characters that
	// a blockchain must be represented by.
	IDLength = 16
)

// Errors are defined below and can be used to
// determine what error you may have received.
var (
	// ErrInvalidData is used when a struct is determined
	// to be invalid.
	ErrInvalidData = errors.New("Struct contains fields with invalid data")

	// ErrJSONFormat is used when a JSON string is unable
	// to be generated.
	ErrJSONFormat = errors.New("Unable to marshal JSON")
)

func toJSONString(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", ErrJSONFormat
	}

	return string(bs), nil
}

func hash512String(bs []byte) string {
	hash := sha512.Sum512(bs)
	hashString := hex.EncodeToString(hash[:])

	return hashString
}
