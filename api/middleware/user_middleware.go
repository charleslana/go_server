package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"go_server/api/helpers"
	"go_server/api/models"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type UserInput struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserNameInput struct {
	Name string `json:"name" validate:"required,min=3,max=20"`
}

func ValidateUserInput(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userInput UserInput
		err := json.NewDecoder(r.Body).Decode(&userInput)
		if err != nil {
			helpers.JSONResponse(w, http.StatusBadRequest, helpers.ErrorResponse(errors.New("Erro ao decodificar o JSON")))
			return
		}

		validate := validator.New()
		err = validate.Struct(&userInput)
		if err != nil {
			var validationErr validator.ValidationErrors
			if errors.As(err, &validationErr) {
				helpers.JSONResponse(w, http.StatusBadRequest, helpers.ErrorResponse(helpers.FormatValidationErrorMessage(validationErr)))
			} else {
				helpers.JSONResponse(w, http.StatusBadRequest, helpers.ErrorResponse(err))
			}
			return
		}

		user := models.User{
			Email:    userInput.Email,
			Password: userInput.Password,
		}

		ctx := context.WithValue(r.Context(), "user", &user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidateUserName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "id")
		id, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			apiError := helpers.ErrorResponse(errors.New("ID do usuário inválido"), http.StatusBadRequest)
			helpers.JSONResponse(w, apiError.StatusCode, apiError)
			return
		}

		var userNameInput UserNameInput
		err = json.NewDecoder(r.Body).Decode(&userNameInput)
		if err != nil {
			helpers.JSONResponse(w, http.StatusBadRequest, helpers.ErrorResponse(errors.New("Erro ao decodificar o JSON")))
			return
		}

		validate := validator.New()
		err = validate.Struct(&userNameInput)
		if err != nil {
			var validationErr validator.ValidationErrors
			if errors.As(err, &validationErr) {
				helpers.JSONResponse(w, http.StatusBadRequest, helpers.ErrorResponse(helpers.FormatValidationErrorMessage(validationErr)))
			} else {
				helpers.JSONResponse(w, http.StatusBadRequest, helpers.ErrorResponse(err))
			}
			return
		}

		user := &models.User{
			ID:   uint(id),
			Name: &userNameInput.Name,
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
