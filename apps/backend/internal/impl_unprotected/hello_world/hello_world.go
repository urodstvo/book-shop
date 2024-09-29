package hello_world

import (
	"net/http"

	"github.com/urodstvo/book-shop/apps/backend/internal/impl_deps"
)

type HelloWorld struct {
	*impl_deps.Deps
}

func (*HelloWorld) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
