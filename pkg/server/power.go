package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// Power toggles display output (onn/off).
type Power struct {
	Value bool
}

// handlePower toggles display output (onn/off).
func (s *HTTPServer) handlePower() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var p Power

		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&p); err != nil {
			log.Print(err.Error())
		}

		if err := s.display.SetDisplay(p.Value); err != nil {
			log.Print(err.Error())
		}
	}
}
