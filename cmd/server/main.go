package main

import (
	"log"
	"net/http"
	"privatealbum/internal/api"
	"privatealbum/internal/config"
	"privatealbum/internal/service"
	"privatealbum/internal/storage"
	"privatealbum/internal/workflow"
)

func main() {
	c := config.Load().Normalize()
	st, e := storage.Open(c.Database)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	svc := service.New(st, c.MaxPhotos)
	eng := workflow.New(svc)
	log.Fatal(http.ListenAndServe(c.Address, api.New(eng).Routes()))
}
