package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// Digit describes a single digit on the display.
type Digit struct {
	Position uint8
	Value    int
}

// handleDigit sets a single digit on the display.
func (s *HTTPServer) handleDigit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var d Digit

		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&d); err != nil {
			log.Print(err.Error())
		}

		if err := s.display.WriteDigit(d.Position, d.Value); err != nil {
			log.Print(err.Error())
		}
	}
}
