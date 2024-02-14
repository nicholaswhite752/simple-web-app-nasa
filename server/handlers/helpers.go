package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (h *Handlers) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1024 * 1024 //1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	// Check that there is only 1 json value in the Body
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have a single json value")
	}

	return nil
}

func (h *Handlers) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	var output []byte
	if h.environment == "development" {
		out, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return err
		}
		output = out
	} else {
		out, err := json.Marshal(data)
		if err != nil {
			return err
		}
		output = out
	}

	// headers is variadic, so it is optional. If the len is > 0, then we will
	// take the first header and add it to the response
	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err := w.Write(output)

	return err
}
