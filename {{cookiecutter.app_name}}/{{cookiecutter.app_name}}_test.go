package main

import (
	"encoding/json"
<<<<<<< HEAD
	"github.com/{{cookiecutter.github_handle}}/{{cookiecutter.app_name}}/handlers"
=======
	"{{cookiecutter.app_name}}/handlers/healthcheck"
>>>>>>> 8b5d5dbe68327d37f3cfddbe57e078ff65fa90e8
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	type ApiMessage struct {
		Message string
	}
	w := httptest.NewRecorder()
	r := gin.Default()
	r.GET("/health", healthcheck.HealthCheck)
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP response code should be 200, was: %d", w.Code)
	}
	var message ApiMessage
	err := json.Unmarshal(w.Body.Bytes(), &message)
	if err != nil {
		t.Errorf("JSON parse error: %d", w.Code)
	}

	if message.Message != "OK" {
		t.Errorf("JSON message should be OK, but is: %s", message.Message)
	}
}
