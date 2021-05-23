package main

import (
	"log"
	"net/http"
)

type (
	M = map[string]interface{}
)

func listen(env *Env) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("OK"))
	})
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("code: " + r.URL.Query().Get("code")))
	})

	log.Println("server listen on " + env.ListenAddr)
	log.Fatalln(http.ListenAndServe(env.ListenAddr, nil))
}