package handler

import (
	"context"
	"encoding/json"
	"go_web_api/internal/auth"
	"go_web_api/internal/domain/user"
	"go_web_api/internal/infrastructure/db"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service       *user.Service
	authenticator *auth.Auth
}

func NewUserHandler(repo *db.UserRepository, authenticator *auth.Auth) *UserHandler {
	service := user.NewService(repo)
	return &UserHandler{
		service:       service,
		authenticator: authenticator,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeError(w, r, http.StatusBadRequest, "invalid JSON body")
		return
	}
	defer r.Body.Close()

	u := &user.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	ctx := context.Background()
	err = h.service.CreateUser(ctx, u)
	if err != nil {
		switch err {
		case user.ErrInvalidEmail, user.ErrInvalidPassword:
			writeError(w, r, http.StatusBadRequest, err.Error())
		case user.ErrEmailAlreadyUsed:
			writeError(w, r, http.StatusConflict, err.Error())
		default:
			writeError(w, r, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	writeJSON(w, http.StatusCreated, map[string]interface{}{
		"status": "success",
		"user": map[string]interface{}{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email,
		},
	})
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, r, http.StatusBadRequest, "invalid user ID")
		return
	}

	ctx := context.Background()
	u, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError, "database error")
		return
	}
	if u == nil {
		writeError(w, r, http.StatusNotFound, "user not found")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"created_at": u.CreatedAt,
	})
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, r, http.StatusBadRequest, "invalid user ID")
		return
	}

	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeError(w, r, http.StatusBadRequest, "invalid JSON body")
		return
	}
	defer r.Body.Close()

	u := &user.User{
		ID:    id,
		Name:  input.Name,
		Email: input.Email,
	}

	ctx := context.Background()
	err = h.service.UpdateUser(ctx, u)
	if err != nil {
		switch err {
		case user.ErrNotFound:
			writeError(w, r, http.StatusNotFound, err.Error())
		default:
			writeError(w, r, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		writeError(w, r, http.StatusBadRequest, "invalid user ID")
		return
	}

	ctx := context.Background()
	err = h.service.DeleteUser(ctx, id)
	if err != nil {
		switch err {
		case user.ErrNotFound:
			writeError(w, r, http.StatusNotFound, err.Error())
		default:
			writeError(w, r, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeError(w, r, http.StatusBadRequest, "invalid JSON body")
		return
	}
	defer r.Body.Close()

	u, err := h.service.Login(context.Background(), input.Email, input.Password)
	if err != nil {
		switch err {
		case user.ErrInvalidCredentials:
			writeError(w, r, http.StatusUnauthorized, err.Error())
		default:
			writeError(w, r, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	token, err := h.authenticator.GenerateToken(u)
	if err != nil {
		writeError(w, r, http.StatusInternalServerError, "failed to generate token")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}

