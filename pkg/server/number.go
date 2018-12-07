package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// Number describes a 4-digit number on the display.
type Number struct {
	Value int
}

// handleNumber sets a number on the whole display.
func (s *HTTPServer) handleNumber() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var n Number

		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&n); err != nil {
			log.Print(err.Error())
		}

		if err := s.display.WriteNumber(n.Value); err != nil {
			log.Print(err.Error())
		}
	}
}
