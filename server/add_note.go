package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/notefile"
	"time"
)

type AddNoteRequest struct {
	Link string
	Text string
}

func InvalidRequest(w http.ResponseWriter) {
	http.Error(w, "Invalid request", http.StatusBadRequest)
}

func ParseJSONBody(r io.ReadCloser, v any) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return fmt.Errorf("could not parse request body: %w", err)
	}

	if err := json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("could not unmarshal request body: %w", err)
	}

	return nil
}

func addNote(nf *notefile.Notefile, w http.ResponseWriter, r *http.Request) {
	req := AddNoteRequest{}
	if err := ParseJSONBody(r.Body, &req); err != nil {
		log.Printf("Error: %s", err)
		InvalidRequest(w)
		return
	}

	note := &notefile.Note{
		Timestamp: time.Now(),
		Link:      req.Link,
		Text:      req.Text,
	}

	nf.Notes = append(nf.Notes, note)
	fmt.Fprintf(w, "ok")
}
