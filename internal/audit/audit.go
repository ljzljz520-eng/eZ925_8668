package audit

import (
	"fmt"
	"privatealbum/internal/model"
	"privatealbum/internal/storage"
	"time"
)

type Logger struct{ st *storage.Store }

func New(st *storage.Store) *Logger { return &Logger{st: st} }
func (l *Logger) Record(actor, action, target string) error {
	v := model.Audit{ID: fmt.Sprintf("audit-%d", time.Now().UnixNano()), Actor: actor, Action: action, Target: target, At: time.Now().UTC()}
	return l.st.SaveAudit(v)
}
func (l *Logger) Event(album, kind, detail string) error {
	v := model.Event{ID: fmt.Sprintf("event-%d", time.Now().UnixNano()), AlbumID: album, Kind: kind, Detail: detail, At: time.Now().UTC()}
	return l.st.SaveEvent(v)
}
func (l *Logger) Describe(v model.Event) string { return v.Kind + ":" + v.Detail }
