package main
import (
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestHandleIndexReturnsWithStatusOK(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	cityHandler(response, request)
	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v", "200", response.Code)
	}
}

Commit the file.

6.	Create a file named Dockerfile.
FROM golang  
WORKDIR /work
ADD . .
RUN go test ./...
RUN go build -o /bin/myapp .
WORKDIR /
RUN rm -r /work
CMD ["/bin/myapp"]
