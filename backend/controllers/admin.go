package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"bakeflow/models"
)

// AdminGetOrders returns all orders for admin dashboard
func AdminGetOrders(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Set JSON header before any response
	w.Header().Set("Content-Type", "application/json")
	
	// Get all orders from database
	orders, err := models.GetAllOrders()
	if err != nil {
		log.Printf("❌ Error fetching orders: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Error fetching orders",
			"details": err.Error(),
			"orders": []interface{}{},
			"total": 0,
		})
		return
	}

	// Return orders as JSON
	response := map[string]interface{}{
		"orders": orders,
		"total":  len(orders),
	}

	json.NewEncoder(w).Encode(response)
}

// AdminUpdateOrderStatus updates the status of an order
func AdminUpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Get order ID from URL path
	// URL format: /api/admin/orders/123/status
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	orderIDStr := pathParts[4]
	orderID, err := strconv.Atoi(orderIDStr)
	if err != nil {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	// Parse request body
	var requestBody struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate status
	validStatuses := map[string]bool{
		"pending":   true,
		"preparing": true,
		"ready":     true,
		"delivered": true,
	}

	if !validStatuses[requestBody.Status] {
		http.Error(w, "Invalid status", http.StatusBadRequest)
		return
	}

	// Update order status in database
	err = models.UpdateOrderStatus(orderID, requestBody.Status)
	if err != nil {
		log.Printf("❌ Error updating order status: %v", err)
		http.Error(w, "Error updating order status", http.StatusInternalServerError)
		return
	}

	log.Printf("✅ Order #%d status updated to: %s", orderID, requestBody.Status)

	// Prepare notification feedback
	notificationSent := false
	notificationError := ""

	// Get order details to send notification to customer
	order, err := models.GetOrderByID(orderID)
	if err != nil {
		log.Printf("⚠️ Could not load order for notification: %v", err)
		notificationError = "load_failed"
	} else if order.SenderID == "" {
		log.Printf("ℹ️ No SenderID stored for order #%d; skipping customer notification", orderID)
		notificationError = "missing_sender_id"
	} else {
		// Send status update message to customer (only for selected statuses)
		statusMessages := map[string]string{
			"pending":   "✅ Your order #%d has been received! We'll start preparing it soon.", // Added pending acknowledgement
			"preparing": "🍰 Great news! We've started preparing your order #%d. It will be ready soon!",
			"ready":     "✅ Your order #%d is ready! Please come pick it up or wait for delivery.",
			"delivered": "🎉 Your order #%d has been delivered! Enjoy your delicious treats!",
		}

		if message, exists := statusMessages[requestBody.Status]; exists {
			notificationText := fmt.Sprintf(message, orderID)
			if err := SendMessage(order.SenderID, notificationText); err != nil {
				log.Printf("⚠️ Failed to send notification to customer: %v", err)
				notificationError = err.Error()
			} else {
				log.Printf("📬 Status notification sent to customer for order #%d", orderID)
				notificationSent = true
			}
		} else {
			// Status not mapped for notifications
			log.Printf("ℹ️ Status '%s' not configured for notifications", requestBody.Status)
			notificationError = "status_not_configured"
		}
	}

	// Return success response with notification metadata
	response := map[string]interface{}{
		"success":            true,
		"message":            "Order status updated successfully",
		"order_id":           orderID,
		"new_status":         requestBody.Status,
		"notification_sent":  notificationSent,
		"notification_error": notificationError,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
