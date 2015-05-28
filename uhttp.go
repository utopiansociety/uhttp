package uhttp

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Read parses the given Request body as JSON, storing the result in the given
// interface.
func Read(r *http.Request, body interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, body)
}

// Write serializes an interface to JSON and writes out the headers, given
// status code and serialized data to the given ResponseWriter.
func Write(w http.ResponseWriter, body interface{}, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	j, err := json.Marshal(body)
	if err != nil {
		return err
	}
	_, err = w.Write(j)
	return err
}

// Error serializes an error to JSON and writes out the headers, given status
// code and serialized data to the given ResponseWriter.
func Error(w http.ResponseWriter, err error, status int) error {
	body := make(map[string]string)
	body["error"] = err.Error()
	return Write(w, body, status)
}

// Status creates a response appropriate to the given status code, and writes
// it out to to the given ResponseWriter. If the status code represents an
// error, a JSON formatted error will be written, otherwise a blank response
// will be returned.
func Status(w http.ResponseWriter, status int) error {
	message := http.StatusText(status)

	if status >= 400 {
		return Error(w, errors.New(message), status)
	} else {
		w.WriteHeader(status)
		return nil
	}
}
