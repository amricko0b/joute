package modifier

import (
	"net/http"
)

// RewriteHeaders copies headers created by Downstream into response to Client
type RewriteHeaders struct{}

func (rh *RewriteHeaders) Modify(downstreamResponse *http.Response, clientResponse http.ResponseWriter) error {

	for key, values := range downstreamResponse.Header {
		for _, val := range values {
			clientResponse.Header().Add(key, val)
		}
	}

	return nil
}
