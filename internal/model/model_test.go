package model

import "testing"

func TestAlbumAdd(t *testing.T) {
	a := NewAlbum("a", "t", "1234")
	a = a.AddPhoto(Record{ID: "r", AlbumID: "a", Path: "p"})
	if len(a.Photos) != 1 {
		t.Fatal()
	}
}
