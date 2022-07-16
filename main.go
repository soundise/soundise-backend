package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/soundise/soundise/handlers"
)

func main() {
	var router *mux.Router = mux.NewRouter()

	staticServer := http.FileServer(http.Dir("static/"))

	handlers.UsersHandler{}.New(router)
	handlers.WebsocketGateway{}.New(router)

	router.PathPrefix("/").Handler(staticServer).Methods("GET")

	var server = http.Server{
		Addr:         ":8000",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,

		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,

			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		},
	}
	log.Println("Listening...")

	log.Fatal(server.ListenAndServe())
}
