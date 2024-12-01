package main

import (
	"casbin/handlers"
	"casbin/pkg"
	"casbin/repositories"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	var (
		r      = chi.NewRouter()
		ur     = repositories.NewUserRepository()
		uh     = handlers.NewUserHandler(ur)
		e, err = pkg.NewCasbinEnforcer()
	)

	if err != nil {
		panic(err)
	}

	r.Use(middleware.Logger)
	r.Use(pkg.NewMyMiddleware)
	r.Use(pkg.MyCasbinMiddleware(e, ur))

	r.Mount("/api/guest", newGuestRoute(uh))
	r.Mount("/api/member", newMemberRoute(uh))

	http.ListenAndServe(":8080", r)
}

func newGuestRoute(uh *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/me", uh.GetUser)
	return r
}

func newMemberRoute(uh *handlers.UserHandler) chi.Router {
	r := chi.NewRouter()
	r.Get("/me", uh.GetUser)
	r.Get("/users", uh.GetUsers)
	r.Post("/users", uh.CreateUser)
	r.Delete("/users/{id}", uh.DeleteUser)
	return r
}
