package handler

import (
	"encoding/json"
	"net/http"

	"go-rest-api-template/internal/auth"
	"go-rest-api-template/internal/repository"
)

type AuthHandler struct {
	repo   *repository.UserRepository
	secret string
}

func NewAuthHandler(repo *repository.UserRepository, secret string) *AuthHandler {
	return &AuthHandler{repo: repo, secret: secret}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.repo.GetByName(req.Username)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := auth.GenerateToken(user.ID, h.secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
