package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestHealthCheck(t *testing.T) {
	router := router()

	tester := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthcheck", nil)
	router.ServeHTTP(tester, req)

	assert.Equal(t, 200, tester.Code)
	assert.Equal(t, "{\"hostname\":\"localhost\",\"serviceStatus\":\"OK\"}", tester.Body.String())
}
