package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go_web_api/internal/domain/user"
	"go_web_api/internal/infrastructure/db"
)

type UserHandler struct {
	service *user.Service
}

func NewUserHandler(repo *db.UserRepository) *UserHandler {
	service := user.NewService(repo)
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON body")
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
			writeError(w, http.StatusBadRequest, err.Error())
		case user.ErrEmailAlreadyUsed:
			writeError(w, http.StatusConflict, err.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error")
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
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Extract user ID from URL path
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		writeError(w, http.StatusBadRequest, "invalid URL format")
		return
	}

	id, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid user ID")
		return
	}

	ctx := context.Background()
	u, err := h.service.GetUserByID(ctx, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "database error")
		return
	}
	if u == nil {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"created_at": u.CreatedAt,
	})
}
