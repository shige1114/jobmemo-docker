package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/backend"
	"github.com/backend/pkg/handler"
	"github.com/gorilla/mux"
)

type SideServer struct {
	store  *handler.Store
	router *mux.Router
}

func (s *SideServer) Router() {
	s.router.HandleFunc("/sidebar", s.Get).Methods(http.MethodGet)

}
func (s *SideServer) Get(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var users_id string
	var sidebar []backend.Side

	vars := r.URL.Query()
	users_id = vars.Get("users_id")

	ctx = r.Context()

	err := s.store.SideStore.Side(ctx, users_id, &sidebar)
	if err != nil {
		ResponseError(w)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sidebar)
	w.WriteHeader(http.StatusOK)

}

func (s *SideServer) Post(w http.ResponseWriter, r *http.Request) {

}
func (s *SideServer) Patch(w http.ResponseWriter, r *http.Request) {

}
