package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"privatealbum/internal/model"
	"time"
)

var buckets = []string{"records", "profiles", "events", "audits", "albums"}

type Store struct{ db *bbolt.DB }

func Open(path string) (*Store, error) {
	db, e := bbolt.Open(path, 0600, &bbolt.Options{Timeout: time.Second})
	if e != nil {
		return nil, e
	}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, n := range buckets {
			if _, x := tx.CreateBucketIfNotExists([]byte(n)); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return &Store{db: db}, nil
}
func (s *Store) Close() error {
	if s == nil || s.db == nil {
		return nil
	}
	return s.db.Close()
}
func put[T any](s *Store, bucket, key string, v T) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), raw) })
}
func get[T any](s *Store, bucket, key string, v *T) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return errors.New("missing bucket")
		}
		raw := b.Get([]byte(key))
		if raw == nil {
			return fmt.Errorf("not found: %s", key)
		}
		return json.Unmarshal(raw, v)
	})
}
func (s *Store) SaveAlbum(a model.Album) error { return put(s, "albums", a.ID, a) }
func (s *Store) LoadAlbum(id string) (model.Album, error) {
	var a model.Album
	e := get(s, "albums", id, &a)
	return a, e
}
func (s *Store) SaveRecord(r model.Record) error { return put(s, "records", r.ID, r) }
func (s *Store) LoadRecord(id string) (model.Record, error) {
	var r model.Record
	e := get(s, "records", id, &r)
	return r, e
}
func (s *Store) SaveProfile(p model.Profile) error { return put(s, "profiles", p.ID, p) }
func (s *Store) LoadProfile(id string) (model.Profile, error) {
	var p model.Profile
	e := get(s, "profiles", id, &p)
	return p, e
}
func (s *Store) SaveEvent(v model.Event) error { return put(s, "events", v.ID, v) }
func (s *Store) SaveAudit(v model.Audit) error { return put(s, "audits", v.ID, v) }
func (s *Store) ListAlbums() ([]model.Album, error) {
	out := []model.Album{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("albums")).ForEach(func(_, v []byte) error {
			var a model.Album
			if e := json.Unmarshal(v, &a); e != nil {
				return e
			}
			out = append(out, a)
			return nil
		})
	})
	return out, e
}
