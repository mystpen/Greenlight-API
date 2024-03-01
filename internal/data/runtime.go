package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) { //for Marshaler interface
	jsonVal := fmt.Sprintf("%d mins", r)

	// to wrap it in double quotes
	// needs to be surrounded by double quotes in order to be a valid *JSON string*.
	quotedJSONValue := strconv.Quote(jsonVal)

	return []byte(quotedJSONValue), nil
}
