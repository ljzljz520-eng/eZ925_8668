package storage

import (
	"go.etcd.io/bbolt"
	"privatealbum/internal/model"
	"sort"
)

func (s *Store) ListRecords(album string) ([]model.Record, error) {
	all, e := s.ListAlbums()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, a := range all {
		if a.ID == album {
			out = append(out, a.Photos...)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return out, nil
}
func (s *Store) ArchiveRecord(id string) error {
	r, e := s.LoadRecord(id)
	if e != nil {
		return e
	}
	r.Archived = true
	return s.SaveRecord(r)
}
func (s *Store) Count(bucket string) (int, error) {
	n := 0
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return nil
		}
		n = b.Stats().KeyN
		return nil
	})
	return n, e
}
func (s *Store) Health() bool { return s != nil && s.db != nil }
