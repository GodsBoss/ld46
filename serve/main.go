package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}

const defaultAddr = "127.0.0.1:8080"

func run(args []string) error {
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = defaultAddr
	}
	serveDir := os.Getenv("SERVE_DIR")
	if serveDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		serveDir = cwd
	}
	fmt.Printf("Serve %s on %s...\n", serveDir, addr)
	fs := http.FileServer(http.Dir(serveDir))
	http.ListenAndServe(
		addr,
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Printf("Serve %s\n", r.URL)
				fs.ServeHTTP(w, r)
			},
		),
	)
	return nil
}
