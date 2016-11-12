package gateway

import (
	"net/http"
	"strings"
	"fmt"
)

func (s *Gateway) Validate(w http.ResponseWriter, r *http.Request) bool {

	if (r.Method != http.MethodPost) {
		http.Error(w, "Unauthorized", http.StatusMethodNotAllowed)
		return false
	}

	auth := r.Header.Get("Authorization")

	if (auth != "") {
		hd := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
		if (len(hd) != 2) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return false
		}

		token := hd[1]
		fmt.Println("token[1]", token)
		err := s.Auth.Validate(token)
		if (err != nil) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return false
		}
		return true
	} else {
		url := r.URL.Path[1: len(r.URL.Path)]
		var err error = nil
		switch url {
		case "signup" :
			err = s.Auth.Signup(r.Header.Get("Email"))
		case "login" :
			token, loginError := s.Auth.Login(r.Header.Get("User"), r.Header.Get("Password"))
			if(loginError != nil){
				http.Error(w, err.Error(), http.StatusBadRequest)
				return false
			}
			w.Write([]byte(token))
		}
		if (err != nil) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return false
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	return false
}
