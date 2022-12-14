package main

import (
	"github.com/nnaakkaaii/tododemo/internal/db"
	"github.com/nnaakkaaii/tododemo/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	os.Exit(run())
}

func run() int {
	termCh := make(chan os.Signal, 1)
	signal.Notify(termCh, syscall.SIGTERM, syscall.SIGINT)

	d := db.NewMemoryDB()
	s := server.NewServer(8080, d)
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	select {
	case <-termCh:
		return 0
	case <-errCh:
		return 1
	}
}
