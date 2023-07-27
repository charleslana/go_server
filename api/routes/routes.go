package routes

import (
	"go_server/api/handlers"
	"go_server/api/middleware"

	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	handler := handlers.NewHandler(db)

	r.Get("/", handler.GetRoot)

	userHandler := handlers.NewUserHandler(db)

	r.With(middleware.ValidateUserInput).Post("/user", userHandler.CreateUser)
	r.With(middleware.ValidateUserName).Put("/user/{id}", userHandler.UpdateUser)
	r.Get("/user", userHandler.GetAllUsers)

	return r
}
