package workflow

import (
	"context"
	"privatealbum/internal/model"
	"privatealbum/internal/service"
	"privatealbum/internal/storage"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	e := New(service.New(s, 5))
	if _, x := e.Intake(context.Background(), "a", "title", "1234"); x != nil {
		t.Fatal(x)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	e := New(service.New(s, 5))
	_, _ = e.Intake(context.Background(), "a", "title", "1234")
	if _, x := e.Submit(context.Background(), "a", model.Record{ID: "r", Path: "p", Caption: "ok"}); x != nil {
		t.Fatal(x)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	e := New(service.New(s, 5))
	_, _ = e.Intake(context.Background(), "a", "title", "1234")
	if _, x := e.Review(context.Background(), "a", "1234"); x != nil {
		t.Fatal(x)
	}
}
