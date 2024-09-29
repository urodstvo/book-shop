package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

func Wrap(base http.Handler, wrappers ...Middleware) http.Handler {
	for _, wrapper := range wrappers {
		base = wrapper(base)
	}

	return base
}
