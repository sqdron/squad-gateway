package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad/activation"
	"net/http"
	"log"
	"github.com/nats-io/nats"
	"github.com/sqdron/squad/connect"
	"time"
)

func main() {
	squad := squad.Client()
	squad.Activate(func(i activation.ServiceInfo) {
		con, e := nats.Connect(i.Endpoint)
		if (e != nil) {
			panic(e)
		}
		c := connect.NewTransport(i.Endpoint)
		http.ListenAndServe(":4000", &SquadMux{connect:con, c:c})
	})
}

type SquadMux struct {
	connect *nats.Conn
	c       connect.ITransport
}

func (s *SquadMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	var path = r.URL.Path[1: len(r.URL.Path)]
	if (r.Method == http.MethodGet) {
		for k, v := range r.URL.Query() {
			request[k] = v[0]
		}
	}
	log.Printf("Requesting %s with params %s\n", path, request)
	res, e := s.c.RequestSync(path, request, 1 * time.Second)
	if (e != nil ) {

		http.Error(w, e.Error(), http.StatusInternalServerError)
		log.Println(e)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(res.([]byte))
	w.WriteHeader(http.StatusOK)
}