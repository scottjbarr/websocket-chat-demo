package main

import (
	"net/http"
	"os"

	log "github.com/scottjbarr/websocket-chat-demo/Godeps/_workspace/src/github.com/Sirupsen/logrus"

	"github.com/scottjbarr/websocket-chat-demo/Godeps/_workspace/src/github.com/gorilla/mux"
)

var (
	rr redisReceiver
	rw redisWriter
)

func main() {
	redisURL := os.Getenv("REDIS_URL")
	redisPool, err := newRedisPool(redisURL)
	if err != nil {
		log.WithField("url", redisURL).Fatal("Unable to create Redis pool")
	}

	rr = newRedisReceiver(redisPool)
	go rr.run()
	rw = newRedisWriter(redisPool)
	go rw.run()

	port := os.Getenv("PORT")
	if port == "" {
		log.WithField("PORT", port).Fatal("$PORT must be set")
	}

	http.HandleFunc("/ws", handleWebsocket)

	r := mux.NewRouter()
	r.HandleFunc("/ws", handleWebsocket)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("public/"))))

	// mux := http.NewServeMux()

	// n := negroni.Classic()
	// n.UseHandler(mux)
	// n.Run(":" + port)
	http.ListenAndServe(":"+port, r)
}
