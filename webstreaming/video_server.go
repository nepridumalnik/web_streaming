package webstreaming

import (
	"errors"
	"net/http"
	"os"
)

type streamer struct {
	folder string
	route  string
}

func MakeStreamer(route string, folder string) (*streamer, error) {
	if _, err := os.Stat(folder); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(folder, os.ModePerm)

		if err != nil {
			return nil, err
		}
	}

	return &streamer{route: route, folder: folder}, nil
}

func (vs *streamer) StreamHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !r.URL.Query().Has("v") {
		http.Error(w, "Parameter \"v\" required", http.StatusBadRequest)
		return
	}

	fileName := r.URL.Query().Get("v")
	http.ServeFile(w, r, vs.folder+"/"+fileName)
}

func (vs *streamer) RegisterHandlers() {
	http.HandleFunc(vs.route, vs.StreamHandler)
}
