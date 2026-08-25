package service

import (
	"encoding/json"
	"examarchive/internal/model"
	"net/http"
)

type Handler struct{ Archive *Archive }

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.list(w)
		return
	}
	if r.Method == http.MethodPost {
		h.create(w, r)
		return
	}
	http.Error(w, "method not allowed", 405)
}
func (h Handler) list(w http.ResponseWriter) {
	rs, e := h.Archive.List()
	if e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	json.NewEncoder(w).Encode(rs)
}
func (h Handler) create(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Candidate model.Candidate
		Score     model.Score
	}
	if json.NewDecoder(r.Body).Decode(&in) != nil {
		http.Error(w, "invalid json", 400)
		return
	}
	out, e := h.Archive.Create(in.Candidate, in.Score)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(out)
}
func NewHandler(a *Archive) http.Handler { return Handler{Archive: a} }
