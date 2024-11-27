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
	r.Get("/api/real-gdp-per-capita", s.GetRealGDPPerCapitaHandler)
	r.Get("/api/federal-funds-effective-rate", s.GetFederalFundsEffectiveRateHandler)
	r.Get("/api/labor-force-participation-rate", s.GetLaborForceParticipationRateHandler)
	r.Get("/api/unemployment-rate", s.GetUnemploymentRateHandler)
	r.Get("/api/mean-unemployment-rate", s.GetMeanUnemploymentRateHandler)
	r.Get("/api/real-median-personal-income", s.GetRealMedianPersonalIncomeHandler)
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

// GetFederalFundsEffectiveRateHandler is a handler that returns the federal funds effective rate data
func (s *Server) GetFederalFundsEffectiveRateHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	transaction, err := s.db.GetFederalFundsEffectiveRate(startDate, endDate)
	if err != nil {
		log.Printf("Failed to fetch federal funds effective rate data: %v", err)
		http.Error(w, "Failed to fetch federal funds effective rate data", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		http.Error(w, "No federal funds effective rate data points found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// GetLaborForceParticipationRateHandler is a handler that returns the labor force participation rate data
func (s *Server) GetLaborForceParticipationRateHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	transaction, err := s.db.GetLaborForceParticipationRate(startDate, endDate)
	if err != nil {
		log.Printf("Failed to fetch labor force participation rate data: %v", err)
		http.Error(w, "Failed to fetch labor force participation rate data", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		http.Error(w, "No labor force participation rate data points found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// GetUnemploymentRateHandler is a handler that returns the unemployment rate data
func (s *Server) GetUnemploymentRateHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	transaction, err := s.db.GetUnemploymentRate(startDate, endDate)
	if err != nil {
		log.Printf("Failed to fetch unemployment rate data: %v", err)
		http.Error(w, "Failed to fetch unemployment rate data", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		http.Error(w, "No unemployment rate data points found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// GetRealMedianPersonalIncomeHandler is a handler that returns the real median personal income data
func (s *Server) GetRealMedianPersonalIncomeHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	transaction, err := s.db.GetRealMedianPersonalIncome(startDate, endDate)
	if err != nil {
		log.Printf("Failed to fetch real median personal income data: %v", err)
		http.Error(w, "Failed to fetch real median personal income data", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		http.Error(w, "No real median personal income data points found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// GetMeanUnemploymentRateHandler is a handler that returns the mean unemployment rate data
func (s *Server) GetMeanUnemploymentRateHandler(w http.ResponseWriter, r *http.Request) {
	startDate := r.URL.Query().Get("startDate")
	endDate := r.URL.Query().Get("endDate")

	transaction, err := s.db.GetMeanUnemploymentRate(startDate, endDate)
	if err != nil {
		log.Printf("Failed to fetch mean unemployment rate data: %v", err)
		http.Error(w, "Failed to fetch mean unemployment rate data", http.StatusInternalServerError)
		return
	}

	if transaction == nil {
		http.Error(w, "No mean unemployment rate data points found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
