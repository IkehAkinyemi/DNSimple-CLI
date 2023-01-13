package utils

import (
	"encoding/json"
	"net/http"
)

// readJSON reads/parses DNSimple response body. Also handles any possible error
func ReadJSON(r *http.Response, dst interface{}) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(dst)

	return err
}
