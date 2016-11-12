package main

import (
	"github.com/sqdron/squad"
	"net/http"
	"github.com/sqdron/squad-gateway/gateway"
	"github.com/sqdron/squad-gateway/remote"
)

type GatewayOptions struct {
	Port string
}

type authClient struct {
}

func main() {
	op := &GatewayOptions{}
	//TODO: this data should be loaded from hab
	squad := squad.Client("squad.gateway", op)

	auth := remote.RemoteAuth(squad.Api)

	exit := squad.Run()
	controller := &gateway.Gateway{Api:squad.Api, Auth:auth}
	println("Listening http on port :" + op.Port)
	http.ListenAndServe(":" + op.Port, controller)
	<-exit
}


//
//func (s *Gateway) HandleAuth(w http.ResponseWriter, r *http.Request) {
//	url := r.URL.Path[1: len(r.URL.Path)]
//	switch url {
//	case "login":
//		login := r.Form.Get("user")
//		if (login == "") {
//			http.Error(w, "user is not defined", http.StatusBadRequest)
//			return
//		}
//
//		password := r.Form.Get("password")
//		if (password == "") {
//			http.Error(w, "password is not defined", http.StatusBadRequest)
//			return
//		}
//
//		token, e := s.Auth.Login(&auth.LoginRequest{User:r.Form.Get("user")})
//		if (e != nil ) {
//			http.Error(w, e.Error(), http.StatusBadRequest)
//			log.Println(e)
//			return
//		}
//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
//		w.Header().Add("Accept", "application/json")
//		w.Write([]byte(token))
//		w.WriteHeader(http.StatusOK)
//	}
//}

