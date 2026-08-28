package model

import "time"

type Record struct {
	ID, AlbumID, Caption, Path string
	CreatedAt                  time.Time
	Archived                   bool
}
type Profile struct {
	ID, Name, Email, AccessCode string
	CreatedAt                   time.Time
}
type Event struct {
	ID, AlbumID, Kind, Detail string
	At                        time.Time
}
type Audit struct {
	ID, Actor, Action, Target string
	At                        time.Time
}
type Album struct {
	ID, Title, AccessCode string
	Photos                []Record
	Version               int
	CreatedAt             time.Time
}

func NewAlbum(id, title, code string) Album {
	return Album{ID: id, Title: title, AccessCode: code, Photos: make([]Record, 0), CreatedAt: time.Now().UTC()}
}
func (a Album) AddPhoto(r Record) Album { a.Photos = append(a.Photos, r); a.Version++; return a }
func (a Album) VisiblePhotos() []Record {
	out := make([]Record, 0, len(a.Photos))
	for _, p := range a.Photos {
		if !p.Archived {
			out = append(out, p)
		}
	}
	return out
}
func (r Record) Valid() bool  { return r.ID != "" && r.AlbumID != "" && r.Path != "" }
func (p Profile) Valid() bool { return p.ID != "" && p.Name != "" && len(p.AccessCode) >= 4 }
