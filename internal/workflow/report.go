package workflow

import (
	"context"
	"privatealbum/internal/model"
	"sort"
)

func SortRecords(rows []model.Record) []model.Record {
	out := append([]model.Record(nil), rows...)
	sort.SliceStable(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out
}
func (e *Engine) Report(ctx context.Context, album, code string) (map[string]int, error) {
	rows, er := e.s.Snapshot(ctx, album, code)
	if er != nil {
		return nil, er
	}
	m := map[string]int{"total": len(rows)}
	for _, r := range rows {
		if r.Archived {
			m["archived"]++
		} else {
			m["active"]++
		}
	}
	return m, nil
}
