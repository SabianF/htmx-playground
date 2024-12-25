package common

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	Handler http.Handler
}

func Notify() Adapter {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				start_time := time.Now()
				handler.ServeHTTP(w, r)
				Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start_time))
			},
		)
	}
}

func Printf(format string, v ...any) {
	log.Printf(format, v...)
}
