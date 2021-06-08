package helpers

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// EncodeCursor is function for encode id to base64 string
// This function will return base64 string
func EncodeCursor(id int64) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", id)))
}

// DecodeCursor is function for decode string to id
// This function will return id and error
func DecodeCursor(cursor string) (int64, error) {
	bytes, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, err
	}

	id, err := strconv.Atoi(string(bytes))
	if err != nil {
		return 0, err
	}

	return int64(id), nil
}
