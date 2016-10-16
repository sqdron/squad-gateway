package main

import (
	"github.com/sqdron/squad"
	"github.com/sqdron/squad/activation"
	"net/http"
	"log"
	"github.com/sqdron/squad/connect"
	"time"
	"strings"
)

type GatewayOptions struct {
	Auth string
	Port string
}

func main() {
	op := &GatewayOptions{}
	//TODO: this data should be loaded from hab
	squad := squad.Client(op)
	squad.Activate(func(i activation.ServiceInfo) {
		println("Listening http on port :" + op.Port)
		http.ListenAndServe(":" + op.Port, &SquadMux{Options:op, Connect:connect.NatsTransport(i.Endpoint)})
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
	return auth[1] == s.Options.Auth
}

//TODO: Refactor usin middleware approach
func (s *SquadMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	if r.Method == "OPTIONS" {
		return
	}

	if (r.Method != http.MethodPost) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if (!s.checkAuth(r)) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	r.ParseForm()

	request := map[string]string{}

	subject := r.Form.Get("subject")
	if (subject == "") {
		subject = r.URL.Path[1: len(r.URL.Path)]
	}

	if (subject == "") {
		http.Error(w, http.ErrNotSupported.Error(), http.StatusBadRequest)

		return
	}

	for key, values := range r.Form {
		request[key] = values[0]
	}

	res, e := s.Connect.RequestSync(subject, request, 3 * time.Second)
	if (e != nil ) {
		http.Error(w, e.Error(), http.StatusBadRequest)
		log.Println(e)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Add("Accept", "application/json")
	w.Write(res.([]byte))
	w.WriteHeader(http.StatusOK)
}
