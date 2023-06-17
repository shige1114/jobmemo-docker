package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/backend"
	"github.com/backend/pkg/handler"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

type MainServer struct {
	store  *handler.Store
	router *mux.Router
}
type MainServerInterface interface {
	Get(c *gin.Context)
	Patch(c *gin.Context)
	Offer(c *gin.Context)
	Reject(c *gin.Context)
	NewSelect(c *gin.Context)
}

func (s *MainServer) Router() {
	s.router.HandleFunc("/main", s.Get).Methods(http.MethodGet)

}
func (s *MainServer) Get(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var users_id, companies_id string
	var main backend.Main

	vars := r.URL.Query()

	users_id = vars.Get("users_id")
	companies_id = vars.Get("companies_id")

	ctx = r.Context()

	err := s.store.MainStore.Main(ctx, users_id, companies_id, &main)
	if err != nil {
		ResponseError(w)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(main)
	w.WriteHeader(http.StatusOK)
}
