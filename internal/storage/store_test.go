package storage

import (
	"privatealbum/internal/model"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	s, e := Open(t.TempDir() + "/a.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	a := model.NewAlbum("a", "title", "1234")
	if e = s.SaveAlbum(a); e != nil {
		t.Fatal(e)
	}
	if _, e = s.LoadAlbum("a"); e != nil {
		t.Fatal(e)
	}
}
