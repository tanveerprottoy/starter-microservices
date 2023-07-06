package httppkg

import (
	"errors"
	"io"
	"net/http"

	"github.com/tanveerprottoy/starter-microservices/gateway/pkg/adapter"
)

func Request[T any](method string, url string, header http.Header, body io.Reader, httpClient *HTTPClient) (int, T, error) {
	var d T
	code, resBody, err := httpClient.Request(method, url, header, body)
	if err != nil {
		return -1, d, err
	}
	if code >= http.StatusOK && code < http.StatusMultipleChoices {
		// res ok, parse response body to type
		d, err = adapter.BytesToType[T](resBody)
		if err != nil {
			return code, d, err
		}
		return code, d, nil
	} else {
		// res not ok, parse error
		errBody, err := adapter.BytesToType[ErrorBody](resBody)
		if err != nil {
			return code, d, err
		}
		errMessage := "Something went wrong"
		if errBody.Message != "" {
			errMessage = errBody.Message
		} else if errBody.Error != "" {
			errMessage = errBody.Error
		}
		return code, d, errors.New(errMessage)
	}
}

func RequestRaw(method string, url string, header http.Header, body io.Reader, httpClient *HTTPClient) (int, []byte, error) {
	return httpClient.Request(method, url, header, body)
}
