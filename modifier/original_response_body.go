package modifier

import (
	"io"
	"net/http"
)

// OriginalResponseBody keeps Downstream's response body as-is. Client will have it unmodified.
type OriginalResponseBody struct{}

func (or *OriginalResponseBody) Modify(downstreamResponse *http.Response, clientResponse http.ResponseWriter) error {
	clientResponse.WriteHeader(downstreamResponse.StatusCode)
	_, err := io.Copy(clientResponse, downstreamResponse.Body)

	return err
}
