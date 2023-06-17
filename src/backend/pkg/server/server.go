package server

import (
	"net/http"

	"github.com/backend/pkg/handler"
	"github.com/gorilla/mux"
)

type Server struct {
	DataUri string
	BaseUri string

	Side *SideServer
	Main *MainServer
}
type ServerInterface interface {
	Router()
}

func (s *Server) Router() {
	servers := []ServerInterface{s.Side, s.Main}
	for _, server := range servers {
		server.Router()
	}
}
func NewServer(test bool, datauri string, baseuri string) (*Server, error) {
	var router *mux.Router
	var store *handler.Store
	var err error

	router = mux.NewRouter()
	store, err = handler.NewStore(test, datauri)
	if err != nil {
		return nil, err
	}

	return &Server{
		DataUri: datauri,
		BaseUri: baseuri,
		Side:    &SideServer{store: store, router: router},
		Main:    &MainServer{store: store, router: router},
	}, nil

}
func DefaultServer() (*Server, error) {
	var router *mux.Router
	var store *handler.Store
	var test bool
	var datauri, baseuri string
	var err error

	baseuri = ":5432"
	datauri = "postgres://root:password@db:5432/test_db"
	test = true
	router = mux.NewRouter()
	store, err = handler.NewStore(test, datauri)
	if err != nil {
		return nil, err
	}

	return &Server{
		DataUri: datauri,
		BaseUri: baseuri,
		Side:    &SideServer{store: store, router: router},
		Main:    &MainServer{store: store, router: router},
	}, nil

}

func ResponseError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
}
