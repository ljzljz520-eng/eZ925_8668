package service

import (
	"context"
	"privatealbum/internal/model"
	"privatealbum/internal/storage"
	"testing"
)

func TestBusinessChain26(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	v := New(s, 4)
	_, _ = v.RegisterAlbum(context.Background(), "a", "title", "1234")
	_, _ = v.AddPhoto(context.Background(), "a", model.Record{ID: "r1", Path: "p", Caption: "original"})
	old, _ := v.Snapshot(context.Background(), "a", "")
	_, _ = v.AddPhoto(context.Background(), "a", model.Record{ID: "r2", Path: "q", Caption: "new"})
	if old[0].Caption != "original" {
		t.Fatalf("old list mutated: %s", old[0].Caption)
	}
}
