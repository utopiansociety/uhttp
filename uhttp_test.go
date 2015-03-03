package uhttp

import (
	"errors"
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

func TestRead(t *testing.T) {
	body := strings.NewReader(`{"hello":"world"}`)
	r := &http.Request{
		Body: ioutil.NopCloser(body),
	}
	var b sample
	err := Read(r, &b)
	assert.Nil(t, err)
	assert.Equal(t, b.Hello, "world", "Hello should contain 'world'")
}

func TestWrite(t *testing.T) {
	b := sample{
		Hello: "world",
	}
	w := httptest.NewRecorder()
	err := Write(w, b, 200)
	assert.Nil(t, err)
	assert.Equal(t, w.Code, 200, "should be expected status code")
	assert.Equal(
		t,
		w.Body.String(),
		`{"hello":"world"}`,
		"should be equal to valid json string")
}

func TestError(t *testing.T) {
	e := errors.New("test")
	w := httptest.NewRecorder()
	err := Error(w, e, 200)
	assert.Nil(t, err)
	assert.Equal(t, w.Code, 200, "should be expected status code")
	assert.Equal(
		t,
		w.Body.String(),
		`{"error":"test"}`,
		"should be formatted as JSON error")
}

func TestStatus(t *testing.T) {
	w := httptest.NewRecorder()
	err := Status(w, 403)
	assert.Nil(t, err)
	assert.Equal(t, w.Code, 403, "should be expected status code")
	assert.Equal(
		t,
		w.Body.String(),
		`{"error":"Forbidden"}`,
		"should be formatted as JSON error")

	w = httptest.NewRecorder()
	err = Status(w, 204)
	assert.Nil(t, err)
	assert.Equal(t, w.Code, 204, "should be expected status code")
	assert.Equal(
		t,
		w.Body.String(),
		"",
		"should be an empty response")
}
