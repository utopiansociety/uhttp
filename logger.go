package uhttp

import (
	"log"
	"net/http"
	"time"
)

// Logger is the middleware to handle route logging.
func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))

	}
	return http.HandlerFunc(fn)
}
