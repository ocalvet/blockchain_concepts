package server

import (
	"log"
	"net/http"
	"os"
	"time"
)

func Run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("APP_PORT")
	log.Println("Listening on ", httpAddr)
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
