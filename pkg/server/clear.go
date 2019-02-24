package server

import (
	"log"
	"net/http"
)

// handleClear controls clearning contents from the display.
func (s *HTTPServer) handleClear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.display.Clear(); err != nil {
			log.Print(err.Error())
		}
	}
}
