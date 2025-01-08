package modifier_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/amricko0b/joute/modifier"
)

func TestOriginalResponsePreserved(t *testing.T) {
	buff := bytes.NewBufferString("{}")
	response := http.Response{StatusCode: 200, Body: io.NopCloser(buff)}

	cientResponse := clientResponseMock{}
	cientResponse.On("Write", mock.Anything).Return(buff.Len(), nil)
	cientResponse.On("WriteHeader", mock.AnythingOfType("int"))

	m := modifier.OriginalResponseBody{}
	m.Modify(&response, &cientResponse)

	cientResponse.AssertCalled(t, "Write", mock.Anything)
	cientResponse.AssertCalled(t, "WriteHeader", response.StatusCode)
	cientResponse.AssertNotCalled(t, "Header")
}
