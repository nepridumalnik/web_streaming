package main

import (
	"log"
	"net/http"
	"web_streaming/webstreaming"
)

func main() {
	streamer, err := webstreaming.MakeStreamer("/samples", "./samples")

	if err != nil {
		log.Fatalf("Failed with error: %s", err)
	}

	streamer.RegisterHandlers()

	err = http.ListenAndServe("127.0.0.1:8080", nil)

	if err != nil {
		log.Fatalf("Failed with error: %s", err)
	}
}
