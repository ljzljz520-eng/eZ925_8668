package storage

import (
	"privatealbum/internal/model"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/a.db"
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	if e = s.SaveProfile(model.Profile{ID: "p", Name: "owner", AccessCode: "1234"}); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.LoadProfile("p"); e != nil {
		t.Fatal(e)
	}
}
