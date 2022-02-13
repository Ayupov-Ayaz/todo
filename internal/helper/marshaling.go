package helper

import (
	"strconv"
)

func MarshalingId(id int) []byte {
	return []byte(`{"id":` + strconv.Itoa(id) + `}`)
}
