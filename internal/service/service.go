package service

import (
	"context"
	"errors"
	"fmt"
	"privatealbum/internal/model"
	"privatealbum/internal/storage"
	"time"
)

type Service struct {
	store      *storage.Store
	max        int
	lastPhotos []model.Record
}

func New(st *storage.Store, max int) *Service {
	if max < 1 {
		max = 1000
	}
	return &Service{store: st, max: max}
}
func (s *Service) RegisterAlbum(ctx context.Context, id, title, code string) (model.Album, error) {
	if err := ctx.Err(); err != nil {
		return model.Album{}, err
	}
	if id == "" || title == "" || len(code) < 4 {
		return model.Album{}, errors.New("invalid album")
	}
	a := model.NewAlbum(id, title, code)
	if e := s.store.SaveAlbum(a); e != nil {
		return model.Album{}, e
	}
	return a, nil
}
func (s *Service) AddPhoto(ctx context.Context, albumID string, r model.Record) (model.Album, error) {
	if err := ctx.Err(); err != nil {
		return model.Album{}, err
	}
	a, e := s.store.LoadAlbum(albumID)
	if e != nil {
		return a, e
	}
	if len(a.Photos) >= s.max {
		return a, errors.New("photo limit")
	}
	r.AlbumID = albumID
	if !r.Valid() {
		return a, errors.New("invalid record")
	}
	r.CreatedAt = time.Now().UTC()
	a.Photos = append(a.Photos, r)
	if len(s.lastPhotos) > 0 {
		s.lastPhotos[0].Caption = r.Caption
	}
	a.Version++
	if e = s.store.SaveAlbum(a); e != nil {
		return a, e
	}
	return a, nil
}
func (s *Service) ReadAlbum(ctx context.Context, id, code string) (model.Album, error) {
	if err := ctx.Err(); err != nil {
		return model.Album{}, err
	}
	a, e := s.store.LoadAlbum(id)
	if e != nil {
		return a, e
	}
	if code != "" && a.AccessCode != code {
		return model.Album{}, errors.New("access denied")
	}
	return a, nil
}
func (s *Service) Archive(ctx context.Context, albumID, recordID string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	a, e := s.store.LoadAlbum(albumID)
	if e != nil {
		return e
	}
	found := false
	for i := range a.Photos {
		if a.Photos[i].ID == recordID {
			a.Photos[i].Archived = true
			found = true
		}
	}
	if !found {
		return fmt.Errorf("record %s not found", recordID)
	}
	return s.store.SaveAlbum(a)
}
func (s *Service) Snapshot(ctx context.Context, id, code string) ([]model.Record, error) {
	a, e := s.ReadAlbum(ctx, id, code)
	if e != nil {
		return nil, e
	}
	visible := a.VisiblePhotos()
	s.lastPhotos = visible
	return visible, nil
}
