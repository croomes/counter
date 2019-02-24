package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// Colon contols the whether the colon on the left of the display is shown.
type Colon struct {
	Value bool
}

// handleColon controls displaying the colon.
func (s *HTTPServer) handleColon() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var c Colon

		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&c); err != nil {
			log.Print(err.Error())
		}

		if err := s.display.Colon(c.Value); err != nil {
			log.Print(err.Error())
		}
	}
}
