package routes

import (
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_admin"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_unprotected"
	"go.uber.org/fx"
)

type Opts struct {
	fx.In

	ImplUnProtected *impl_unprotected.UnProtected
	ImplAdmin       *impl_admin.Admin
	ImplProtected   *impl_protected.Protected

	session *scs.SessionManager
}

func New(opts Opts) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", opts.ImplUnProtected.HelloWorld.HelloWorld).Methods("GET")

	api := router.PathPrefix("/api").Subrouter()

	v1 := api.PathPrefix("/v1").Subrouter()

	auth := v1.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/login", opts.ImplUnProtected.Auth.Login).Methods("POST")
	auth.HandleFunc("/register", opts.ImplUnProtected.Auth.Register).Methods("POST")
	auth.HandleFunc("/logout", opts.ImplUnProtected.Auth.Logout).Methods("POST")

	return router
}
