package service

import (
	"context"
	"privatealbum/internal/model"
	"privatealbum/internal/storage"
	"testing"
)

func TestRegister(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	v := New(s, 3)
	if _, e := v.RegisterAlbum(context.Background(), "a", "title", "1234"); e != nil {
		t.Fatal(e)
	}
	if _, e := v.AddPhoto(context.Background(), "a", model.Record{ID: "r", Path: "p", Caption: "first"}); e != nil {
		t.Fatal(e)
	}
}
