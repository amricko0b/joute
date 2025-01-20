package joute

import (
	"net/http"
)

type (
	EndpointMap     map[string]*Endpoint
	EndpointRouting string
)

// Endpoint is exported by Joute to client.
// Client exchanges JSON-RPC messages with endpoint.
type Endpoint struct {
	Routing EndpointRouting
	RouteTo string
}

const (
	RoutingDirect    EndpointRouting = "direct"
	RoutingUseMethod EndpointRouting = "useMethod"
)

func (endpoint *Endpoint) MakeHandlerFunc(app *App) http.HandlerFunc {
	return func(clientResponse http.ResponseWriter, clientRequest *http.Request) {

		clientMessage, err := unmarshallBodyIntact(clientRequest)
		if err != nil {
			// TODO specification is violated here...
			// TODO when bad JSON-RPC received from client we must answer with appropriate JSON-RPC response ("Parse error" or "Invalid Request")
			http.Error(clientResponse, "Joute: not a JSON-RPC request", http.StatusBadRequest)
			return
		}

		ds := app.Downstreams[endpoint.RouteTo]
		if ds == nil {
			http.Error(clientResponse, "Joute: no such downstream", http.StatusServiceUnavailable)
			return
		}

		var (
			downstreamResponse *http.Response
			downstreamError    error
		)

		switch endpoint.Routing {
		case RoutingDirect:
			downstreamResponse, downstreamError = ds.CallDirect(clientRequest)
		case RoutingUseMethod:
			downstreamResponse, downstreamError = ds.CallMethod(clientRequest, clientMessage.Method)
		}

		if downstreamError != nil {
			http.Error(clientResponse, "Joute: downstream interaction error", http.StatusInternalServerError)
			return
		}

		defer downstreamResponse.Body.Close()

		// TODO complete the sequence
		clientResponse.WriteHeader(200)
	}
}
