package main

import (
	_ "github.com/hongjundu/go-embed-static-example/statik"
	"github.com/rakyll/statik/fs"
	"log"
	"net/http"
)

// Before build, run statik -src ./build

func main() {
	log.Printf("starts...")

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(statikFS))

	http.ListenAndServe(":3000", nil)
}
