package api

import (
	"net/http/httptest"
	"privatealbum/internal/service"
	"privatealbum/internal/storage"
	"privatealbum/internal/workflow"
	"testing"
)

func TestHealth(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	h := New(workflow.New(service.New(s, 2))).Routes()
	r := httptest.NewRecorder()
	h.ServeHTTP(r, httptest.NewRequest("GET", "/health", nil))
	if r.Code != 200 {
		t.Fatal(r.Code)
	}
}
