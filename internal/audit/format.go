package audit

import (
	"privatealbum/internal/model"
	"strings"
)

func NormalizeActor(v string) string { return strings.TrimSpace(strings.ToLower(v)) }
func Filter(events []model.Event, kind string) []model.Event {
	out := []model.Event{}
	for _, e := range events {
		if kind == "" || e.Kind == kind {
			out = append(out, e)
		}
	}
	return out
}
