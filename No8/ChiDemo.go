package No8

import (
	"learn-go/No8/handler"
	"learn-go/No8/middleware"
	"learn-go/No8/store"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ChiDemoTest() {
	store.InitDB()
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
	})
	r.Route("/order", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/", handler.CreateOrder)
		r.Get("/", handler.MyOrders)

	})
	http.ListenAndServe(":8080", r)
}
