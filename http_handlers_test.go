package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNameHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/name?name=Someone", nil)

	w := httptest.NewRecorder()
	NameHandler(w,req)

	if w.Code != http.StatusOK {
		t.Errorf("%s returned with wrong wrong code: %v; wanted %v","NameHandler", w.Code, http.StatusOK)
	}

	if w.Body.String() != "Hello, Someone" {
		t.Errorf("%s returned %v; wanted %v", "NameHandler", w.Body.String(), "Hello, Someone")
	}
}

func TestNameHandler_TableDriven(t *testing.T) {
	tests := []struct {
		testName string
		name string
		expectedOutput string
		expectedCode int
	} {
		{"Name = Someone", "Someone", "Hello, Someone", 200},
		{" Name = \"\"", "", "Hello, Who", 200},
		{" Name = 1231 1234", "1231 1234", "Hello, 1231 1234", 200},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "" {
				tt.name = "Who" //default for empty name
			}

			req := httptest.NewRequest("GET", "/name?name="+tt.name, nil)
			w := httptest.NewRecorder()

			NameHandler(w,req)

			if status := w.Code; status != http.StatusOK {
				t.Errorf("%s returned code: %v. wanted %v", "NameHandler", status, http.StatusOK)
			}

			if body := w.Body.String(); body != "Hello, " + tt.name {
				t.Errorf("%s returned: %v. wanted %v", "NameHandler", body, "Hello, " + tt.name)
			}
		})
	}
}