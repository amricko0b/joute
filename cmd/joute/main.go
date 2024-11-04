package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/amricko0b/joute/handler"
)

func main() {
	var (
		port         int
		targetScheme string
		targetHost   string
		targetPort   int
	)

	flag.IntVar(&port, "p", 9000, "Http port to bind Joute on")
	flag.StringVar(&targetScheme, "target-scheme", "http", "Proxied host protocol scheme")
	flag.StringVar(&targetHost, "target-host", "127.0.0.1", "Proxied host")
	flag.IntVar(&targetPort, "target-port", 80, "Proxied port")
	flag.Parse()

	joute := handler.JouteHandler{
		TargetScheme:   targetScheme,
		TargetHostPort: fmt.Sprintf("%s:%d", targetHost, targetPort),
		Client:         &http.Client{},
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), &joute); err != nil {
		panic(err)
	}
}
