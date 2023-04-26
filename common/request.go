package common

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func MakePostRequest(url string, payload []byte) (map[string]interface{}, error) {
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(payload))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
	return res, nil
}
