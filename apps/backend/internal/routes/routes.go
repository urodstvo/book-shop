package routes

import (
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_admin"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_protected"
	"github.com/urodstvo/book-shop/apps/backend/internal/impl_unprotected"
	"github.com/urodstvo/book-shop/apps/backend/internal/middlewares"
	"go.uber.org/fx"
)

type Opts struct {
	fx.In

	ImplUnProtected *impl_unprotected.UnProtected
	ImplAdmin       *impl_admin.Admin
	ImplProtected   *impl_protected.Protected

	Session *scs.SessionManager
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

	books := v1.PathPrefix("/books").Subrouter()
	books.HandleFunc("", opts.ImplUnProtected.Books.GetBooks).Methods("GET")
	books.HandleFunc("/{BookId}", opts.ImplUnProtected.Books.GetByBookId).Methods("GET")
	books.HandleFunc("/{BookId}/preview", opts.ImplUnProtected.Books.BookPreview).Methods("GET")
	books.HandleFunc("/{BookId}/rate/{Rating}", middlewares.WithAuth(opts.Session, opts.ImplProtected.Books.Rate)).Methods("PUT")
	books.HandleFunc("/request", middlewares.WithAuth(opts.Session, opts.ImplProtected.Books.RequestBook)).Methods("POST")
	books.HandleFunc("/recomendations", middlewares.WithAuth(opts.Session, opts.ImplProtected.Books.Recomendations)).Methods("GET")

	users := v1.PathPrefix("/users").Subrouter()
	users.HandleFunc("/me", middlewares.WithAuth(opts.Session, opts.ImplProtected.Users.GetMe)).Methods("GET")
	users.HandleFunc("/me", middlewares.WithAuth(opts.Session, opts.ImplProtected.Users.PutMe)).Methods("PUT")

	orders := v1.PathPrefix("/orders").Subrouter()
	orders.HandleFunc("", middlewares.WithAuth(opts.Session, opts.ImplProtected.Orders.GetOrders)).Methods("GET")
	orders.HandleFunc("", middlewares.WithAuth(opts.Session, opts.ImplProtected.Orders.CreateOrder)).Methods("POST")
	orders.HandleFunc("/{OrderId}", middlewares.WithAuth(opts.Session, opts.ImplProtected.Orders.GetOrder)).Methods("GET")
	orders.HandleFunc("/{OrderId}", middlewares.WithAuth(opts.Session, opts.ImplProtected.Orders.DeleteOrder)).Methods("DELETE")
	orders.HandleFunc("/{OrderId}/status", middlewares.WithAuth(opts.Session, opts.ImplProtected.Orders.GetOrderStatus)).Methods("GET")
	orders.HandleFunc("/report", middlewares.WithAuth(opts.Session, opts.ImplProtected.Orders.GetReport)).Methods("GET")

	payments := v1.PathPrefix("/payments").Subrouter()
	payments.HandleFunc("", middlewares.WithAuth(opts.Session, opts.ImplProtected.GetPayments)).Methods("GET")
	payments.HandleFunc("", middlewares.WithAuth(opts.Session, opts.ImplProtected.AddPayment)).Methods("POST")
	payments.HandleFunc("/{PaymentId}", middlewares.WithAuth(opts.Session, opts.ImplProtected.DeletePayment)).Methods("DELETE")

	return router
}
