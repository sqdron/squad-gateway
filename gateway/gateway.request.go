package gateway

import (
	"fmt"
	"net/http"
	"log"
)

func (s *Gateway) HandleRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	request := map[string]string{}
	subject := r.Form.Get("subject")
	fmt.Println(subject)

	if (subject == "") {
		http.Error(w, http.ErrNotSupported.Error(), http.StatusBadRequest)

		return
	}

	for key, values := range r.Form {
		request[key] = values[0]
	}
	res, e := s.Api.Request(subject, request)
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