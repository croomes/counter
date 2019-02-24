package server

import (
	"encoding/json"
	"log"
	"net/http"
)

// Brightness contols the LED brightness level, ranging from 0 (off) to 15
// (full power).
type Brightness struct {
	Value uint8
}

// handleBrightness sets a number on the whole display.
func (s *HTTPServer) handleBrightness() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var n Brightness

		dec := json.NewDecoder(r.Body)
		if err := dec.Decode(&n); err != nil {
			log.Print(err.Error())
		}

		if err := s.display.SetBrightness(n.Value); err != nil {
			log.Print(err.Error())
		}
	}
}
