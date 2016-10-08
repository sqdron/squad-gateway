package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad/activation"
	"net/http"
	"log"
	"fmt"
	"github.com/nats-io/nats"
	"encoding/json"
	"time"
)

func main() {
	squad := squad.Client()
	squad.Activate(func(i activation.ServiceInfo) {
		connect, _ := nats.Connect(i.Endpoint)
		http.ListenAndServe(":4000", &SquadMux{connect})
	})
}

type SquadMux struct {
	connect *nats.Conn
}

func (s *SquadMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := map[string]string{}
	var path = r.URL.Path[1: len(r.URL.Path)]
	fmt.Println(r)
	log.Printf("Route to %s \n", path)

	if (r.Method == http.MethodGet) {
		for k, v := range r.URL.Query() {
			request[k] = v[0]
		}
	}
	log.Printf("Requesting %s with params %s\n", path, request)

	data, encodeError := json.Marshal(request)
	if (encodeError != nil) {
		panic(encodeError)
	}

	res, _ := s.connect.Request("get_auth_url", data, 10 * time.Millisecond)
	log.Printf("Requestis %s \n", res)
}