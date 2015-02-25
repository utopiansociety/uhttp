package httputils

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sample struct {
	Hello string `json:"hello"`
}

func TestParseJson(t *testing.T) {
	body := strings.NewReader(`{"hello":"world"}`)
	r := &http.Request{
		Body: ioutil.NopCloser(body),
	}
	var b sample
	err := ParseJson(r, &b)
	assert.Nil(t, err)
	assert.Equal(t, b.Hello, "world", "Hello should contain 'world'")
}

func TestJsonWriter(t *testing.T) {
	b := sample{
		Hello: "world",
	}
	w := httptest.NewRecorder()
	err := Json(w, 200, b)
	assert.Nil(t, err)
	assert.Equal(t, w.Body.String(), `{"hello":"world"}`, "should be equal to valid json string")
}
