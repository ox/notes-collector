package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/notefile"
)

func getNotes(nf *notefile.Notefile, w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(nf)
	if err != nil {
		http.Error(w, "Could not marshal notes", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(b))
}
