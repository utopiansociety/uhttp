package recovery

import (
	"log"
	"net/http"
)

// Handler is how to handle errors right now.
func Handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %#v", err)
				w.WriteHeader(500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
