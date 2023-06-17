package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/backend/pkg/handler"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestSideServer(t *testing.T) {

	store, err := handler.NewStore(true, "postgres://root:password@db:5432/test_db")
	if err != nil {
		t.Fatal(err)
	}
	s := SideServer{store: store}
	req := httptest.NewRequest(http.MethodGet, "http://localhost:5432/sidebar/", nil)
	req = mux.SetURLVars(req, map[string]string{"users_id": "185ffaae-e320-11ed-8886-26359435711c"})

	rr := httptest.NewRecorder()

	s.Get(rr, req)

	if rr.Code == http.StatusInternalServerError {
		t.Fatal(rr.Body)
	}

}
