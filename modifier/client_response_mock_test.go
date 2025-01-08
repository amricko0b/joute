package modifier_test

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

func (rw *clientResponseMock) Header() http.Header {
	args := rw.Called()
	return args.Get(0).(http.Header)
}

func (rw *clientResponseMock) Write(bytes []byte) (int, error) {
	args := rw.Called(bytes)
	return args.Int(0), args.Error(1)
}

func (rw *clientResponseMock) WriteHeader(statusCode int) {
	rw.Called(statusCode)
}

// clientResponseMock mocks HTTP's ResponseWriter
type clientResponseMock struct {
	mock.Mock
}
