package store

import (
	"github.com/gorilla/mux"
	"net/http/httptest"
	"testing"
)

func TestPetRequestAvailable(t *testing.T) {
	p := Pet{}
	router := mux.NewRouter()
	p.ConfigureRoute(router)
	localServer := httptest.NewServer(router)

	assert := func(t *testing.T, s *httptest.Server, c int) {
		resp, err := s.Client().Get(s.URL + "/products?status=available")
		if err != nil {
			t.Fatalf("unexpected error getting from server: %v", err)
		}
		if resp.StatusCode != c {
			t.Fatalf("expected a status code of 200, got %v", resp.StatusCode)
		}
	}

	defer localServer.Close()
	assert(t, localServer, 200)
}

func TestPetRequestSold(t *testing.T) {
	p := Pet{}
	router := mux.NewRouter()
	p.ConfigureRoute(router)
	localServer := httptest.NewServer(router)

	assert := func(t *testing.T, s *httptest.Server, c int) {
		resp, err := s.Client().Get(s.URL + "/products?status=sold")
		if err != nil {
			t.Fatalf("unexpected error getting from server: %v", err)
		}
		if resp.StatusCode != c {
			t.Fatalf("expected a status code of 200, got %v", resp.StatusCode)
		}
	}

	defer localServer.Close()
	assert(t, localServer, 200)
}

func TestPetRequestPending(t *testing.T) {
	p := Pet{}
	router := mux.NewRouter()
	p.ConfigureRoute(router)
	localServer := httptest.NewServer(router)

	assert := func(t *testing.T, s *httptest.Server, c int) {
		resp, err := s.Client().Get(s.URL + "/products?status=pending")
		if err != nil {
			t.Fatalf("unexpected error getting from server: %v", err)
		}
		if resp.StatusCode != c {
			t.Fatalf("expected a status code of 200, got %v", resp.StatusCode)
		}
	}

	defer localServer.Close()
	assert(t, localServer, 200)
}

func TestPetRequestWrongStatus(t *testing.T) {
	p := Pet{}
	router := mux.NewRouter()
	p.ConfigureRoute(router)
	localServer := httptest.NewServer(router)

	assert := func(t *testing.T, s *httptest.Server, c int) {
		resp, err := s.Client().Get(s.URL + "/products?status=notavailable")
		if err != nil {
			t.Fatalf("unexpected error getting from server: %v", err)
		}
		if resp.StatusCode != c {
			t.Fatalf("expected a status code of 200, got %v", resp.StatusCode)
		}
	}

	defer localServer.Close()
	assert(t, localServer, 400)
}
