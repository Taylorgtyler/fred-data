package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()

	// CORS configuration
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsOptions).Handler)

	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Get("/api/real-gdp", s.GetRealGDPHandler)
	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to create response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

// GetRealGDPHandler is a handler that returns the real GDP data
func (s *Server) GetRealGDPHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	transaction, err := s.db.GetRealGDP(startDate, endDate)
	if err != nil {
		log.Printf("Failed to fetch real GDP data: %v", err)
		http.Error(w, "Failed to fetch real GDP data", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		http.Error(w, "No GDP data points found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// GetRealGDPPerCapitaHandler is a handler that returns the real GDP per capita data
func (s *Server) GetRealGDPPerCapitaHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	transaction, err := s.db.GetRealGDPPerCapita(startDate, endDate)
	if err != nil {
		log.Printf("Failed to fetch real GDP per capita data: %v", err)
		http.Error(w, "Failed to fetch real GDP per capita data", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		http.Error(w, "No GDP per capita data points found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
