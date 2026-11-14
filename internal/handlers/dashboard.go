package handlers

import (
	_ "embed"
	"net/http"
)

//go:embed dashboard.html
var dashboardHTML []byte

// DashboardHandler serves the minimal endpoint health dashboard: a
// single HTML page with inline JavaScript and CSS and no external
// dependencies. The page itself needs no API key; the browser prompts
// for the key, stores it in localStorage, and sends it to the existing
// /v1/endpoints and /v1/endpoints/{id}/health APIs.
type DashboardHandler struct{}

// NewDashboardHandler creates a DashboardHandler.
func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{}
}

// Register wires the handler into the mux.
func (h *DashboardHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /dashboard", h.serve)
}

// serve returns the dashboard page.
func (h *DashboardHandler) serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(dashboardHTML)
}
