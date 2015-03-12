package recovery

import (
	"log"
	"net/http"
)

// Handler is how to handle errors right now.
func Handler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			log.Println("RECOVERY START")
			if err := recover(); err != nil {
				log.Println("RECOVERY ERROR")
				//log.Panic(err)
				log.Printf("panic: %#v", err)
				w.WriteHeader(500)
				w.Write([]byte("hello!!!"))
				log.Println("lalalalala")
			}
			log.Println("RECOVERY END")
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
