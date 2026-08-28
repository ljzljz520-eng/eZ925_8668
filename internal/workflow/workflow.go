package workflow

import (
	"context"
	"fmt"
	"privatealbum/internal/model"
	"privatealbum/internal/service"
)

type Engine struct{ s *service.Service }

func New(s *service.Service) *Engine { return &Engine{s: s} }
func (e *Engine) Intake(ctx context.Context, id, title, code string) (model.Album, error) {
	return e.s.RegisterAlbum(ctx, id, title, code)
}
func (e *Engine) Submit(ctx context.Context, album string, r model.Record) (model.Album, error) {
	p, er := e.s.Process(ctx, album, r)
	if er != nil {
		return model.Album{}, er
	}
	return e.s.AddPhoto(ctx, album, p)
}
func (e *Engine) Review(ctx context.Context, album, code string) (model.Album, error) {
	return e.s.ReadAlbum(ctx, album, code)
}
func (e *Engine) Archive(ctx context.Context, album, record string) error {
	return e.s.Archive(ctx, album, record)
}
func (e *Engine) Trace(ctx context.Context, album string) (string, error) {
	a, err := e.s.ReadAlbum(ctx, album, "")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", a.ID, a.Version), nil
}
