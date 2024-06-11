package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TreeServer struct {
	listenAddr string
	wd         string
}

type ServerError struct {
	Error string `json:"error"`
}

func NewTreeServer(listenAddr string, wd string) *TreeServer {
	return &TreeServer{
		listenAddr: listenAddr,
		wd:         wd,
	}
}

func (t *TreeServer) Run() {
	fmt.Println("Tree Server starting on port " + t.listenAddr)
	http.ListenAndServe(t.listenAddr, makeHTTPHandlerFunc(t.handleReq))
}

func (t *TreeServer) handleReq(w http.ResponseWriter, r *http.Request) error {
	fpath := t.wd + r.URL.Path
	file, err := os.Stat(fpath)
	if err != nil {
		return WriteJSON(w, http.StatusInternalServerError, ServerError{Error: "error reading file: " + fpath})
	}

	if file.IsDir() {
		return WriteJSON(w, http.StatusOK, readDir(fpath))
	}

	qp := r.URL.Query()
	// If dl=1 or download=1 are passed in the query string,
	// download the file to the user's browser.
	shouldDownload := qp.Get("dl") == "1" || qp.Get("download") == "1"
	if shouldDownload {
		// Set the appropriate header so the browser downloads the file.
		w.Header().Set("Content-Disposition", "attachment; filename="+file.Name())
	}

	http.ServeFile(w, r, fpath)
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type treeFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandlerFunc(f treeFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ServerError{Error: err.Error()})
		}
	}
}
