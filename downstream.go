package joute

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type (
	DownstreamMap     map[string]*Downstream
	DownstreamURL     url.URL
	DownstreamTimeout time.Duration
)

// Downstream is server behind Joute.
// Each downstream may serve JSON-RPC API - thus Joute acts like simple gateway.
// Each downstream may serve JSON API - thus Joute provides more complicated functions on adapting incoming requests and outgoing responses.
type Downstream struct {
	URL     *DownstreamURL `json:"url"` // Only Scheme, Host and Path fields are specified
	Timeout DownstreamTimeout
}

// CallMethod calls downstream's method exported on dedicated endpoint (<downstream.URL>/<method>).
// Target URL will be modified to change request direction
func (d *Downstream) CallMethod(clientRequest *http.Request, method string) (*http.Response, error) {
	cli := http.Client{Timeout: time.Duration(d.Timeout)}

	clientRequest.URL.Scheme = d.URL.Scheme
	clientRequest.URL.Host = d.URL.Host
	clientRequest.URL.Path = fmt.Sprintf("%s/%s", d.URL.Path, method)

	// This is mandatory when proxying HTTP request to another server!
	clientRequest.Host = ""
	clientRequest.RequestURI = ""

	return cli.Do(clientRequest)
}

// CallDirect calls downstream directly.
// Target URL will be modified to change request direction
func (d *Downstream) CallDirect(clientRequest *http.Request) (*http.Response, error) {
	cli := http.Client{Timeout: time.Duration(d.Timeout)}

	clientRequest.URL.Scheme = d.URL.Scheme
	clientRequest.URL.Host = d.URL.Host
	clientRequest.URL.Path = d.URL.Path

	// This is mandatory when proxying HTTP request to another server!
	clientRequest.Host = ""
	clientRequest.RequestURI = ""

	return cli.Do(clientRequest)
}

// Golang "enconding/json" doesn't implement unmarshaling for time.Duration
func (timeout *DownstreamTimeout) UnmarshalJSON(payload []byte) error {

	var formatted string
	if err := json.Unmarshal(payload, &formatted); err != nil {
		return nil
	}

	if duration, err := time.ParseDuration(formatted); err == nil {
		*timeout = DownstreamTimeout(duration)
		return nil
	} else {
		return err
	}
}

// Golang "enconding/json" doesn't implement unmarshaling for url.URL
func (u *DownstreamURL) UnmarshalJSON(payload []byte) error {

	var formatted string
	if err := json.Unmarshal(payload, &formatted); err != nil {
		return nil
	}

	if parsed, err := url.Parse(formatted); err == nil {
		*u = DownstreamURL(*parsed)
		return nil
	} else {
		return err
	}
}
