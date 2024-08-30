package helpers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetBody(r *http.Request) (map[string]interface{}, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	var jsonData map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &jsonData); err != nil {
		return nil, err
	}

	return jsonData, nil
}
