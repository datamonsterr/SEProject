package main

import (
	"log"
	"net/http"

	"example.com/test/handler"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handler.GetHomePage)
	router.HandleFunc("/click/open/", handler.GetBookingView)

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Server started at :8080")
	server.ListenAndServe()
}
