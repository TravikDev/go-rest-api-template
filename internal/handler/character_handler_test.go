package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-api-template/internal/middleware"
	"go-rest-api-template/internal/models"
)

type stubCharRepo struct {
	createCalled bool
	lastChar     *models.Character
	char         *models.Character
	err          error
}

func (s *stubCharRepo) Create(c *models.Character) error {
	s.createCalled = true
	s.lastChar = c
	return s.err
}

func (s *stubCharRepo) GetByUserID(userID int) (*models.Character, error) {
	return s.char, s.err
}

func TestCharacterHandler_Create(t *testing.T) {
	repo := &stubCharRepo{}
	h := NewCharacterHandler(repo)

	body := bytes.NewBufferString(`{"user_id":1,"nickname":"Hero","level":1,"experience":0,"x":0,"y":0,"z":0}`)
	req := httptest.NewRequest(http.MethodPost, "/characters", body)
	req = req.WithContext(context.WithValue(req.Context(), middleware.UserIDKey, 1))
	rr := httptest.NewRecorder()

	h.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}
	if !repo.createCalled {
		t.Fatalf("expected Create to be called")
	}
}

func TestCharacterHandler_CreateForbidden(t *testing.T) {
	repo := &stubCharRepo{}
	h := NewCharacterHandler(repo)

	body := bytes.NewBufferString(`{"user_id":2,"nickname":"Hero","level":1,"experience":0,"x":0,"y":0,"z":0}`)
	req := httptest.NewRequest(http.MethodPost, "/characters", body)
	req = req.WithContext(context.WithValue(req.Context(), middleware.UserIDKey, 1))
	rr := httptest.NewRecorder()

	h.Create(rr, req)

	if rr.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d", rr.Code)
	}
	if repo.createCalled {
		t.Fatalf("expected Create not to be called")
	}
}
