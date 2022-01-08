package helper

import "encoding/json"

func MarshalingId(id int) ([]byte, error) {
	return json.Marshal(struct {
		ID int `json:"id"`
	}{
		ID: id,
	})
}
