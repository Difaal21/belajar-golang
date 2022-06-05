package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func TestHtpp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:5000", nil)

	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	bodyString := string(body)
	fmt.Println(bodyString)
}
