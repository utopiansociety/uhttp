package uhttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Read parses the req body into an interface. The interface should be a pointer.
func Read(r *http.Request, body interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, body)
	return err
}

// Write uses the body and the status to respond to the request.
func Write(w http.ResponseWriter, body interface{}, status int) (err error) {
	w.WriteHeader(status)
	j, err := json.Marshal(body)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	return err
}
