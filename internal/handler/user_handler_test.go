package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-api-template/internal/models"
)

type stubRepo struct {
	createCalled bool
	lastUser     *models.User
	user         *models.User
	users        []*models.User
	err          error
}

func (s *stubRepo) Create(u *models.User) error {
	s.createCalled = true
	s.lastUser = u
	return s.err
}

func (s *stubRepo) GetByID(id int) (*models.User, error) {
	return s.user, s.err
}

func (s *stubRepo) List() ([]*models.User, error) {
	return s.users, s.err
}

func TestUserHandler_Create(t *testing.T) {
	repo := &stubRepo{}
	h := NewUserHandler(repo)

	body := bytes.NewBufferString(`{"name":"Tom","email":"tom@example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/users", body)
	rr := httptest.NewRecorder()

	h.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", rr.Code)
	}
	if !repo.createCalled {
		t.Fatalf("expected Create to be called")
	}
	if repo.lastUser.Name != "Tom" || repo.lastUser.Email != "tom@example.com" {
		t.Fatalf("unexpected user passed to repo: %+v", repo.lastUser)
	}
}

func TestUserHandler_GetByID(t *testing.T) {
	expected := &models.User{ID: 1, Name: "Tom", Email: "tom@example.com"}
	repo := &stubRepo{user: expected}
	h := NewUserHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/users/show?id=1", nil)
	rr := httptest.NewRecorder()

	h.GetByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	var got models.User
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if got != *expected {
		t.Fatalf("expected %+v, got %+v", expected, &got)
	}
}

func TestUserHandler_List(t *testing.T) {
	users := []*models.User{{ID: 1, Name: "Tom"}}
	repo := &stubRepo{users: users}
	h := NewUserHandler(repo)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rr := httptest.NewRecorder()

	h.List(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	var got []*models.User
	if err := json.NewDecoder(rr.Body).Decode(&got); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if len(got) != 1 || got[0].Name != "Tom" {
		t.Fatalf("unexpected response: %+v", got)
	}
}
