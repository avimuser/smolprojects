package handler

import (
	"io"
	"log"
	"net/http"

	"github.com/vinaykandagatla/backend-projects/caching-server/internal/cache"
)

type Handler struct {
	Origin string
}

func expect(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	URL := h.Origin + request.URL.Path

	cache.Init()
	defer cache.DB.Close()

	fetchedCache, cached := cache.Get(URL)
	header := fetchedCache.Header
	body := fetchedCache.Body

	if cached {
		header["X-Cache"] = []string{"HIT"}
	} else {
		response, err := http.Get(URL)
		expect(err)

		header = response.Header
		body, err = io.ReadAll(response.Body)
		expect(err)

		newCache := cache.Cache{Header: header, Body: body}
		cache.Store(URL, newCache)

		header["X-Cache"] = []string{"MISS"}
	}

	for key, values := range header {
		for _, value := range values {
			writer.Header().Add(key, value)
		}
	}
	writer.Write(body)
}
