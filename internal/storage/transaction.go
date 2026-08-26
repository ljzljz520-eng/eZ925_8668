package storage

import (
	"privatealbum/internal/model"
	"time"
)

func (s *Store) SaveBundle(a model.Album, p model.Profile, e model.Event) error {
	if err := s.SaveAlbum(a); err != nil {
		return err
	}
	if err := s.SaveProfile(p); err != nil {
		return err
	}
	return s.SaveEvent(e)
}
func (s *Store) Touch(id string) error {
	a, e := s.LoadAlbum(id)
	if e != nil {
		return e
	}
	a.CreatedAt = time.Now().UTC()
	return s.SaveAlbum(a)
}
