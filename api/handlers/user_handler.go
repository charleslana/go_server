package handlers

import (
	"go_server/api/helpers"
	"go_server/api/models"
	"go_server/api/services"
	"net/http"

	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)

	userService := services.NewUserService(h.db)

	apiError := userService.CreateUser(user)
	if apiError.Error {
		helpers.JSONResponse(w, apiError.StatusCode, apiError)
		return
	}

	helpers.SuccessResponse(w, http.StatusCreated, "Usuário cadastrado com sucesso.")
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*models.User)

	userService := services.NewUserService(h.db)
	apiError := userService.UpdateUserName(user.ID, *user.Name)
	if apiError.Error {
		helpers.JSONResponse(w, apiError.StatusCode, apiError)
		return
	}

	helpers.SuccessResponse(w, http.StatusOK, "Usuário atualizado com sucesso.")
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userService := services.NewUserService(h.db)
	users, err := userService.GetAllUsers()
	if err != nil {
		apiError := helpers.ErrorResponse(err, http.StatusInternalServerError)
		helpers.JSONResponse(w, apiError.StatusCode, apiError)
		return
	}
	helpers.JSONResponse(w, http.StatusOK, users)
}
