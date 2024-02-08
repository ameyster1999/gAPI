package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/welcome", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getGreetingHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned worng status code got %v and expected %v", status, http.StatusOK)
	}
	expected := `{"Greeting":"Hello World!"}`

	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned %v and expected %v", rr.Body.String(), expected)
	}
	var msg Message
	err = json.Unmarshal(rr.Body.Bytes(), &msg)
	if err != nil {
		t.Errorf("Error decoding JSON response: %v", err)
	}

	if msg.Greeting != "Hello World!" {
		t.Errorf("Unexpected greeting value. Got %v, expected %v", msg.Greeting, "Hello World!")
	}

}
