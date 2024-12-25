package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/vinaykandagatla/backend-projects/caching-server/internal/cache"
	"github.com/vinaykandagatla/backend-projects/caching-server/internal/handler"
)

var port int
var origin string
var clear bool

func init() {
	flag.IntVar(&port, "port", 8080, "The port where the proxy server will be started")
	flag.BoolVar(&clear, "clear-cache", false, "Should the cache be cleared")
	flag.StringVar(&origin, "origin", "", "The URL of the origin server")
}

func main() {
	flag.Parse()

	if clear {
		cache.Clear()
		log.Print("Cleared cache")
	}

	if origin != "" {
		log.Printf("Hosted at %d\n", port)
		log.Printf("Mirroring %s\n", origin)
		http.ListenAndServe(fmt.Sprintf(":%d", port), handler.Handler{Origin: origin})
	}
}
