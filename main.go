package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/sqdron/squad-gateway/server"
	"github.com/nats-io/nats"
	"github.com/sqdron/squad/configurator"
)

type GateMux struct {
	Url string
}

type Message struct {
	Path string
}

func (p *GateMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "GET"){
		fmt.Fprintf(w, "Access denied for method GET!")
		return
	}

	auth:= r.Header.Get("Authorization");
	if (auth == "") {
		fmt.Fprintf(w, "Access denied !")
		return
	}

	log.Println(r.URL.Path + auth)
	nc, _ := nats.Connect(p.Url)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	m := &Message{Path:r.URL.Path}
	sendCh := make(chan *Message)
	ec.BindSendChan("hello", sendCh)

	sendCh <- m
	fmt.Fprintf(w, "success")
	nc.Close()
	return
}

func main() {

	opts := &server.Options{}

	cfg := configurator.New()
	cfg.ReadFlags(opts)

	nc, _ := nats.Connect(opts.Url)

	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	type person struct {
		Name    string
		Address string
		Age     int
	}

	recvCh := make(chan *person)
	ec.BindRecvChan("hello", recvCh)

	sendCh := make(chan *person)
	ec.BindSendChan("hello", sendCh)

	me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street"}

	// Send via Go channels
	sendCh <- me

	// Receive via Go channels
	who := <-recvCh
	log.Println(who)

	log.Println("Listening...")
	mux := &GateMux{Url:opts.Url}
	http.ListenAndServe(":3000", mux)
}