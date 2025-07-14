package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Subscription struct {
	UserID      string    `json:"user_id"`
	Active      bool      `json:"active"`
	StartedAt   time.Time `json:"started_at,omitempty"`
	CancelledAt time.Time `json:"cancelled_at,omitempty"`
}

var subscriptions = make(map[string]*Subscription)

func main() {
	http.HandleFunc("/api/subscribe", subscribeHandler)
	http.HandleFunc("/api/status", statusHandler)
	http.HandleFunc("/api/cancel", cancelHandler)
	http.HandleFunc("/health", healthHandler)

	log.Println("PremiumAccess API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var sub Subscription
	if err := json.NewDecoder(r.Body).Decode(&sub); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Simulated payment processing
	time.Sleep(1 * time.Second) // Processing delay

	sub.Active = true
	sub.StartedAt = time.Now()
	subscriptions[sub.UserID] = &sub

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Subscription activated",
		"data":    sub,
	})
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	sub, ok := subscriptions[userID]
	if !ok {
		http.Error(w, "Subscription not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sub)
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req struct{ UserID string `json:"user_id"` }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	sub, ok := subscriptions[req.UserID]
	if !ok || !sub.Active {
		http.Error(w, "Active subscription not found", http.StatusNotFound)
		return
	}
	sub.Active = false
	sub.CancelledAt = time.Now()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Subscription cancelled",
		"data":    sub,
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
