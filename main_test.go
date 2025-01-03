package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMemberDetails(t *testing.T) {
	req, err := http.NewRequest("GET", "/member?name=Ponni", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(getMemberDetails)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK; got %v", rec.Code)
	}

	expectedBody := `{"Name":"Ponni","StudentID":"500228122","Image":"po.png","GitHub":"https://github.com/ponnisajeevan12/"}`
	if rec.Body.String() != expectedBody+"\n" {
		t.Errorf("Expected body %v; got %v", expectedBody, rec.Body.String())
	}
}

func TestGetMemberDetailsNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/member?name=NonExistent", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(getMemberDetails)
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status NotFound; got %v", rec.Code)
	}

	expectedBody := "Member not found\n"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body %v; got %v", expectedBody, rec.Body.String())
	}
}
