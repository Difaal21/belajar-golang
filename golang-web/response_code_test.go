package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name != "" {
		// writer.WriteHeader(400)
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(writer, "name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("response.StatusCode", response.StatusCode)
	fmt.Println("response.Status", response.Status)
	fmt.Println("body", string(body))
}
