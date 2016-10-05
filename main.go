package main

import (
	"net/http"
	"fmt"
	"github.com/sqdron/squad-gateway/server"
	"github.com/sqdron/squad/configurator"
	"log"
	"github.com/sqdron/squad/endpoint/nats"
)

type GateMux struct {
	Url string
}

type TokenRequest struct {
				 //	ClientId     string
				 //	ClientSecret string   //     `json:"client_secret"`
	State string
	Code  string // `json:"code"`
}

type TokenResponce struct {
	url string
}

func (p *GateMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//if (r.Method == "GET"){
	//	fmt.Fprintf(w, "Access denied for method GET!")
	//	return
	//}
	//
	//auth:= r.Header.Get("Authorization");
	//if (auth == "") {
	//	fmt.Fprintf(w, "Access denied !")
	//	return
	//}
	var path = r.URL.Path[1: len(r.URL.Path)]
	log.Println("Url... " + path)

	req := &TokenRequest{Code:"", State:""}
	nt := nats.NatsEndpoint(p.Url)
	resp := <-nt.Request(path, req)
	fmt.Println("Got responce")
	fmt.Println(resp)
	nt.Close()
	return
}

func main() {

	opts := &server.Options{}
	fmt.Println("Read configuration...")
	cfg := configurator.New()
	cfg.ReadFlags(opts)

	//squad.Client(opts.Hub, opts.ApplicationID).Activate()

	//client.Hub().
	//
	//client.Activate()
	//
	//nc, _ := nats.Connect(opts.Url)
	//
	//ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	//defer ec.Close()
	//
	//type person struct {
	//	Name    string
	//	Address string
	//	Age     int
	//}
	//
	//recvCh := make(chan *person)
	//ec.BindRecvChan("hello", recvCh)
	//
	//sendCh := make(chan *person)
	//ec.BindSendChan("hello", sendCh)
	//
	//me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street"}
	//
	//// Send via Go channels
	//sendCh <- me
	//
	//// Receive via Go channels
	//who := <-recvCh
	//log.Println(who)

	log.Println("Listening...")
	mux := &GateMux{Url:opts.Hub}
	http.ListenAndServe(":4000", mux)

	//nt := nats.NatsEndpoint(opts.Hub)
	//m := &endpoint.Message{}
	//nt.Publish("auth") <- m
	//nt.Close()
}