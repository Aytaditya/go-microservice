package service

import "net/http"

func GiveService() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi I am Aditya"))
	}
}
