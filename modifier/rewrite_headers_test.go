package modifier_test

import (
	"net/http"
	"testing"

	"github.com/amricko0b/joute/modifier"
)

func TestHeadersAreRewrittenWhenExist(t *testing.T) {

	response := http.Response{Header: map[string][]string{"X-Test-Header1": {"1"}, "X-Test-Header2": {"2"}}}

	actualHeaders := make(map[string][]string)
	rw := clientResponseMock{}
	rw.On("Header").Return(http.Header(actualHeaders))

	m := modifier.RewriteHeaders{}
	m.Modify(&response, &rw)

	rw.AssertCalled(t, "Header")
	rw.AssertNotCalled(t, "Write")
	rw.AssertNotCalled(t, "WriteHeader")
}

func TestMiddlewareSkippedWhenHeadersAreNil(t *testing.T) {

	response := http.Response{}
	rw := clientResponseMock{}

	m := modifier.RewriteHeaders{}
	m.Modify(&response, &rw)

	rw.AssertNotCalled(t, "Header")
	rw.AssertNotCalled(t, "Write")
	rw.AssertNotCalled(t, "WriteHeader")
}

func TestMiddlewareSkippedWhenNoHeaders(t *testing.T) {

	response := http.Response{Header: map[string][]string{}}
	rw := clientResponseMock{}

	m := modifier.RewriteHeaders{}
	m.Modify(&response, &rw)

	rw.AssertNotCalled(t, "Header")
	rw.AssertNotCalled(t, "Write")
	rw.AssertNotCalled(t, "WriteHeader")
}
