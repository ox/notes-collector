package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/notefile"
	"syscall"
)

var port = 0
var help = false
var notesLocalPath string

func init() {
	flag.StringVar(&notesLocalPath, "local", "./notes", "Path to local notes file")
	flag.IntVar(&port, "port", 9872, "Port to listen on")
	flag.BoolVar(&help, "help", false, "Print usage")
	flag.Parse()
}

func setupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func main() {
	if help {
		flag.PrintDefaults()
		return
	}

	setupCloseHandler()

	var nf *notefile.Notefile
	var err error

	if _, err = os.Stat(notesLocalPath); err != nil {
		if os.IsNotExist(err) {
			log.Printf("Notefile does not exist, generating...")
			nf = notefile.NewNotefile()
			if err = nf.Save(notesLocalPath); err != nil {
				log.Fatalf("Could not save new notefile: %s", err)
			}
		}
	} else {
		nf, err = notefile.Load(notesLocalPath)
		if err != nil {
			log.Fatalf("Could not load notefile: %s\n", err)
		}
	}

	log.Printf("Loaded notefile at %s with %d notes\n", notesLocalPath, len(nf.Notes))

	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s /notes\n", r.Method)
		if r.Method == http.MethodGet {
			getNotes(nf, w, r)
		} else if r.Method == http.MethodPost {
			addNote(nf, w, r)
			nf.Save(notesLocalPath)
		} else {
			http.Error(w, "Invalid HTTP method", http.StatusBadRequest)
		}
	})

	log.Printf("Listening on :%d\n", port)

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
