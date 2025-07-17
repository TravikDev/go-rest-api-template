package handler

import (
	"encoding/json"
	"net/http"

	"go-rest-api-template/internal/auth"
	"go-rest-api-template/internal/repository"
)

type AuthHandler struct {
	repo   repository.UserRepositoryInterface
	secret string
}

func NewAuthHandler(repo repository.UserRepositoryInterface, secret string) *AuthHandler {
	return &AuthHandler{repo: repo, secret: secret}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.repo.GetByUsername(req.Username)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	if user.Locked {
		http.Error(w, "user locked", http.StatusForbidden)
		return
	}

	if !auth.CheckPassword(user.PasswordHash, req.Password) {
		attempts := user.FailedAttempts + 1
		locked := attempts >= 10
		_ = h.repo.UpdateLoginState(user.ID, attempts, locked)
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	_ = h.repo.UpdateLoginState(user.ID, 0, false)
	token, err := auth.GenerateToken(user.ID, h.secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
