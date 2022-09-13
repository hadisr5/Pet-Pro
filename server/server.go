package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ws", wsHandler)
	server := http.Server{
		Addr:    ":9167",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.upgrade(w, r, nil)
	if err != nil {
		log.Println("Error in upgrad connection , err ", err)
		return
	}
	hub := newHub(conn)
	hub.run()
}
