package audit

import (
	"privatealbum/internal/storage"
	"testing"
)

func TestAuditEvent(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	l := New(s)
	if e := l.Record("u", "view", "a"); e != nil {
		t.Fatal(e)
	}
	if e := l.Event("a", "open", "ok"); e != nil {
		t.Fatal(e)
	}
}
