# uhttp

> uhttp -> utopia + http

As working with the [net/http](https://golang.org/pkg/net/http) package,
we have found that it is easier to use core packages rather than some
framework but we do need some utility helper functions like Read and Write.
This impliments JSON read and Write functionality but could easily be written
to handle content negotiation.


All of this is meant to be used with [Alice](https://github.com/justinas/alice).

## Example Write:

```go
package main

import (
	"stash.itriagehealth.com/go/uhttp"
	"net/http"
)

type Profile struct {
	Name    string
	Hobbies []string
}

func main() {
  	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	var profile Profile
	err := uhttp.Read(r, &profile)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
	profile = Profile{"Alex", []string{"snowboarding", "programming"}}
	err = uhttp.Write(w, profile, 200)
	if err != nil {
		http.Error(w, err.Error(), 422)
		return
	}
}
```

## Recovery

Recovery is middleware to help with if your applications is crashing, what do you do?
It will catch the error and respond with http.Error for you. Usually this is custom
and will be modified to you needs but if you want a generic one, here it is.


## Logger

Need to log requests. Here you go. Thats all this middleware does. Again, if you need
it to be custom. Please fork and change to your desires. This is meant to be generic.


## Routing

We suggest using (httpRouter)[https://github.com/julienschmidt/httprouter].

# Contributing

1. Tests must continue to pass.

```shell
	go test -v
```

2. go lint, go vet and go fmt must be ran on the code
