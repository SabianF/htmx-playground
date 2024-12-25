package common

import "net/http"

type Adapter func(http.Handler) http.Handler

func Adapt(handler http.Handler, adapter Adapter) http.Handler {
	return adapter(handler)
}
