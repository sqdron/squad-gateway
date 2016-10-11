package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad/activation"
	"net/http"
	"log"
	"github.com/sqdron/squad/connect"
	"time"
	"github.com/sqdron/squad/configurator"
	"strings"
)

type GatewayOptions struct {
	AuthToken string `json:"auth_token"`
}

func main() {
	op := &GatewayOptions{}

	//TODO: this data should be loaded from hab
	configurator.New().ReadFromFile("app.json", &op)

	squad := squad.Client()
	squad.Activate(func(i activation.ServiceInfo) {
		println("Listening http on port :8080")
		http.ListenAndServe(":8080", &SquadMux{Options:op, Connect:connect.NewTransport(i.Endpoint)})
	})
}

type SquadMux struct {
	Connect connect.ITransport
	Options *GatewayOptions
}

func (s *SquadMux) checkAuth(r *http.Request) bool {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 {
		return false
	}
	return auth[1] == s.Options.AuthToken
}

//TODO: Refactor usin middleware approach
func (s *SquadMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if r.Method == "OPTIONS" {
		return
	}

	if (!s.checkAuth(r)) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	r.ParseForm()

	request := map[string]string{}
	var path = r.URL.Path[1: len(r.URL.Path)]
	if (r.Method == http.MethodGet) {
		for k, v := range r.URL.Query() {
			request[k] = v[0]
		}
	}

	for key, values := range r.Form {
		request[key] = values[0]
	}

	log.Printf("Requesting %s with params %s\n", path, request)
	res, e := s.Connect.RequestSync(path, request, 1 * time.Second)
	if (e != nil ) {

		http.Error(w, e.Error(), http.StatusInternalServerError)
		log.Println(e)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Accept", "application/json")
	w.Write(res.([]byte))
	w.WriteHeader(http.StatusOK)
}
