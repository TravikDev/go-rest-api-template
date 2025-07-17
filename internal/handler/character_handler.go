package handler

import (
	"encoding/json"
	"net/http"

	"go-rest-api-template/internal/models"
	"go-rest-api-template/internal/repository"
)

// CharacterHandler handles character endpoints.
type CharacterHandler struct {
	repo repository.CharacterRepositoryInterface
}

// NewCharacterHandler creates a new CharacterHandler.
func NewCharacterHandler(repo repository.CharacterRepositoryInterface) *CharacterHandler {
	return &CharacterHandler{repo: repo}
}

// Create adds a new character for a user.
func (h *CharacterHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID     int     `json:"user_id"`
		Nickname   string  `json:"nickname"`
		Level      int     `json:"level"`
		Experience int     `json:"experience"`
		X          float64 `json:"x"`
		Y          float64 `json:"y"`
		Z          float64 `json:"z"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	char := models.Character{
		UserID:     req.UserID,
		Nickname:   req.Nickname,
		Level:      req.Level,
		Experience: req.Experience,
		PosX:       req.X,
		PosY:       req.Y,
		PosZ:       req.Z,
	}
	if err := h.repo.Create(&char); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(char)
}
