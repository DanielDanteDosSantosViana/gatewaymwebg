package build

import (
	"net/http"
	"testing"
)

func TestRequestBuilder(t *testing.T) {
	req, _ := http.NewRequest("POST", "", nil)
	builder := NewRequestBuilder(req)
	requestRedirect := builder.Build()
	if requestRedirect == nil {
		t.Errorf("Error request return is %v", requestRedirect)
	}
}
