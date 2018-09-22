package main

import (
	"context"
	"log"
	"time"

	"crypto/rand"
	"crypto/tls"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/net/http2"
)

func main() {
	ctx := context.Background()

	m := autocert.Manager{
		Cache:      autocert.DirCache("cert"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: nil,
	}

	tlsConfig := &tls.Config{
		Rand:           rand.Reader,
		Time:           time.Now,
		NextProtos:     []string{http2.NextProtoTLS, "http/1.1"},
		MinVersion:     tls.VersionTLS12,
		GetCertificate: m.GetCertificate,
	}

	ln, err := tls.Listen("tcp", ":https", tlsConfig)
	if err != nil {
		log.Fatalf("ssl listener %v", err)
	}

	server := http.Server{
		Addr:        "localhost:8080",
		TLSConfig:   tlsConfig,
		ReadTimeout: 15 * time.Second,
	}

	server.ListenAndServe()
}
