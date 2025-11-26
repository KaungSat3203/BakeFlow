package routes

import (
	"bakeflow/controllers"
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs all incoming requests (useful for debugging webhook issues)
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		log.Printf("➡️  %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		
		// Call the next handler
		next.ServeHTTP(w, r)
		
		log.Printf("⬅️  Completed in %v", time.Since(start))
	})
}

// CORSMiddleware adds CORS headers to allow cross-origin requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

func SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint (useful for monitoring)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("BakeFlow Bot is running! ✅"))
	})

	// Messenger webhook endpoint
	// GET: Facebook verification
	// POST: Receive messages from users
	mux.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			controllers.VerifyWebhook(w, r)
		} else if r.Method == "POST" {
			controllers.ReceiveWebhook(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Orders API
	mux.HandleFunc("/orders", controllers.GetOrders)
	
	// Admin API Routes
	mux.HandleFunc("/api/admin/orders", controllers.AdminGetOrders)
	mux.HandleFunc("/api/admin/orders/", controllers.AdminUpdateOrderStatus)

	// Wrap with middleware
	handler := LoggingMiddleware(mux)
	handler = CORSMiddleware(handler)

	return handler
}
