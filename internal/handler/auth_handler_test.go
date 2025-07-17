package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-api-template/internal/auth"
	"go-rest-api-template/internal/models"
)

type stubAuthRepo struct {
	user *models.User
	err  error
}

func (s *stubAuthRepo) Create(u *models.User) error { return nil }

func (s *stubAuthRepo) GetByID(id int) (*models.User, error) { return s.user, s.err }

func (s *stubAuthRepo) List() ([]*models.User, error) { return nil, nil }

func (s *stubAuthRepo) GetByUsername(username string) (*models.User, error) {
	return s.user, s.err
}

func (s *stubAuthRepo) UpdateLoginState(id int, attempts int, locked bool) error {
	if s.user != nil {
		s.user.FailedAttempts = attempts
		s.user.Locked = locked
	}
	return nil
}

func TestAuthHandler_Login_Success(t *testing.T) {
	hash, _ := auth.HashPassword("pass")
	repo := &stubAuthRepo{user: &models.User{ID: 1, Username: "u", PasswordHash: hash}}
	h := NewAuthHandler(repo, "secret")

	body := bytes.NewBufferString(`{"username":"u","password":"pass"}`)
	req := httptest.NewRequest(http.MethodPost, "/login", body)
	rr := httptest.NewRecorder()

	h.Login(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if _, err := auth.ParseToken(resp["token"], "secret"); err != nil {
		t.Fatalf("invalid token: %v", err)
	}
}

func TestAuthHandler_Login_Invalid(t *testing.T) {
	hash, _ := auth.HashPassword("pass")
	repo := &stubAuthRepo{user: &models.User{ID: 1, Username: "u", PasswordHash: hash}}
	h := NewAuthHandler(repo, "secret")

	body := bytes.NewBufferString(`{"username":"u","password":"wrong"}`)
	req := httptest.NewRequest(http.MethodPost, "/login", body)
	rr := httptest.NewRecorder()

	h.Login(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rr.Code)
	}
}
