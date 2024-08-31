package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/guluzadehh/kode_test/apps/auth"
	"github.com/guluzadehh/kode_test/apps/note"
)

type APIServer struct {
	addr string
	db   *sql.DB // for future use
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	authService := auth.NewService(auth.NewMemoryStorage())
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(router)

	noteService := note.NewService(note.NewMemoryStorage())
	noteHandler := note.NewHandler(noteService, authHandler)
	noteHandler.RegisterRoutes(router)

	log.Printf("Server started at %s\n", s.addr)
	return http.ListenAndServe(s.addr, loggingMiddleware(router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
