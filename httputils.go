package httputils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseJson(r *http.Request, body interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, body)
	return err
}

func Json(w http.ResponseWriter, status int, body interface{}) (err error) {
	w.WriteHeader(status)
	j, err := json.Marshal(body)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(j)
	return err
}
