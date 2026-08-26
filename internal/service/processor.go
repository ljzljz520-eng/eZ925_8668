package service

import (
	"context"
	"errors"
	"privatealbum/internal/model"
	"strings"
)

func (s *Service) ValidateCaption(c string) error {
	if strings.TrimSpace(c) == "" || len(c) > 240 {
		return errors.New("caption invalid")
	}
	return nil
}
func (s *Service) Process(ctx context.Context, albumID string, r model.Record) (model.Record, error) {
	if err := ctx.Err(); err != nil {
		return r, err
	}
	if e := s.ValidateCaption(r.Caption); e != nil {
		return r, e
	}
	r.Caption = strings.TrimSpace(r.Caption)
	return r, nil
}
func (s *Service) Find(ctx context.Context, albumID, term string) ([]model.Record, error) {
	rows, e := s.Snapshot(ctx, albumID, "")
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range rows {
		if term == "" || strings.Contains(strings.ToLower(r.Caption), strings.ToLower(term)) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Service) Delete(ctx context.Context, albumID, recordID string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	a, e := s.store.LoadAlbum(albumID)
	if e != nil {
		return e
	}
	next := make([]model.Record, 0, len(a.Photos))
	for _, r := range a.Photos {
		if r.ID != recordID {
			next = append(next, r)
		}
	}
	if len(next) == len(a.Photos) {
		return errors.New("not found")
	}
	a.Photos = next
	return s.store.SaveAlbum(a)
}
