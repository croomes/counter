package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HTTPServer returns a new http api server.
type HTTPServer struct {
	http    *http.Server
	router  *mux.Router
	display Displayer
}

// Displayer can display numbers.
type Displayer interface {
	SetBrightness(b uint8) error
	Colon(on bool) error
	Clear() error
	WriteDigit(pos uint8, d int) error
	WriteNumber(d int) error
	WriteBinary(pos uint8, b string) error
}

// New creates a new server instance.
func New(d Displayer) *HTTPServer {

	r := mux.NewRouter()

	return &HTTPServer{
		http: &http.Server{
			Addr:    ":8001",
			Handler: r,
		},
		router:  r,
		display: d,
	}
}

// Run the server, blocking.
func (s *HTTPServer) Run() error {

	// Add routes
	s.Routes()

	return s.http.ListenAndServe()
}

// Shutdown the server.
func (s *HTTPServer) Shutdown() error {
	return s.http.Close()
}

// Routes maps URI paths to handlers.
func (s *HTTPServer) Routes() {
	s.router.HandleFunc("/digit", s.handleDigit()).Methods("POST")
	s.router.HandleFunc("/number", s.handleNumber()).Methods("POST")
	s.router.HandleFunc("/", s.handleIndex())
}

func (s *HTTPServer) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%#v\n", s.router.GetRoute("/digit"))
	}

}
