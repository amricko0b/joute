package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/amricko0b/joute"
	"github.com/amricko0b/joute/modifier"
)

func main() {
	var (
		port             int
		downstreamScheme string
		downstreamHost   string
		downstreamPort   int
	)

	flag.IntVar(&port, "p", 9000, "Http port to bind Joute on")
	flag.StringVar(&downstreamScheme, "downstream-scheme", "http", "Downstream protocol scheme (http/https)")
	flag.StringVar(&downstreamHost, "downstream-host", "127.0.0.1", "HTTP Downstream server host")
	flag.IntVar(&downstreamPort, "downstream-port", 80, "HTTP Downstream server port")
	flag.Parse()

	joute := joute.Handler{
		RequestModifiers: []joute.RequestModifier{
			&modifier.AddHeaders{},
			&modifier.Redirect{
				DownstreamScheme:   downstreamScheme,
				DownstreamHostPort: fmt.Sprintf("%s:%d", downstreamHost, downstreamPort),
			},
		},

		ResponseModifiers: []joute.ResponseModifier{
			&modifier.OriginalResponseBody{},
			&modifier.RewriteHeaders{},
		},

		Client: &http.Client{},
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), &joute); err != nil {
		panic(err)
	}
}
