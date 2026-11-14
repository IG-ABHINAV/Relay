package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDashboardServesPage(t *testing.T) {
	mux := http.NewServeMux()
	NewDashboardHandler().Register(mux)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/dashboard", nil)
	mux.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("GET /dashboard status = %d, want %d", recorder.Code, http.StatusOK)
	}

	contentType := recorder.Header().Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		t.Fatalf("Content-Type = %q, want text/html", contentType)
	}

	body := recorder.Body.String()
	for _, want := range []string{
		"<!DOCTYPE html>",
		"<title>Relay — Endpoint Health</title>",
		"</html>",
		"/v1/endpoints",
		"/health",
		"localStorage",
		"relay_api_key",
		"setInterval",
		"visibilitychange",
	} {
		if !strings.Contains(body, want) {
			t.Errorf("dashboard page is missing %q", want)
		}
	}

	// The page must be self-contained: no external scripts, stylesheets,
	// or assets.
	for _, forbidden := range []string{
		"<script src=",
		"<link ",
		"http://",
		"https://",
	} {
		if strings.Contains(body, forbidden) {
			t.Errorf("dashboard page must not contain %q", forbidden)
		}
	}
}

func TestDashboardRejectsOtherMethods(t *testing.T) {
	mux := http.NewServeMux()
	NewDashboardHandler().Register(mux)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/dashboard", nil)
	mux.ServeHTTP(recorder, request)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf("POST /dashboard status = %d, want %d", recorder.Code, http.StatusMethodNotAllowed)
	}
}
