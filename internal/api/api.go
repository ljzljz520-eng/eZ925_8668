package api

import (
	"context"
	"encoding/json"
	"net/http"
	"privatealbum/internal/model"
	"privatealbum/internal/workflow"
)

type Server struct{ engine *workflow.Engine }

func New(e *workflow.Engine) *Server { return &Server{engine: e} }
func (s *Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.health)
	m.HandleFunc("/albums", s.albums)
	return m
}
func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
func (s *Server) albums(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.create(w, r)
		return
	}
	http.Error(w, "method", http.StatusMethodNotAllowed)
}
func (s *Server) create(w http.ResponseWriter, r *http.Request) {
	var in struct{ ID, Title, Code string }
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "bad", 400)
		return
	}
	a, e := s.engine.Intake(context.Background(), in.ID, in.Title, in.Code)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(a)
}
func WriteJSON(w http.ResponseWriter, v any) {
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
func ParseRecord(r *http.Request) (model.Record, error) {
	var v model.Record
	e := json.NewDecoder(r.Body).Decode(&v)
	return v, e
}
