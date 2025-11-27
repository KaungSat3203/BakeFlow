package controllers

import (
	"fmt"
	"strings"
)

// handleMessage processes text messages from users
func handleMessage(userID, messageText string) {
	state := GetUserState(userID)
	msgLower := strings.ToLower(strings.TrimSpace(messageText))

	// ========== SMART TEXT MATCHING (English + Burmese) ==========

	// Cancel/Reset - Natural language understanding
	if strings.Contains(msgLower, "cancel") ||
		strings.Contains(msgLower, "ပယ်ဖျက်") ||
		strings.Contains(msgLower, "reset") ||
		strings.Contains(msgLower, "start over") ||
		strings.Contains(msgLower, "ပြန်စမယ်") {
		ResetUserState(userID)
		SendMessage(userID, "❌ Order cancelled.")
		SendMessage(userID, "━━━━━━━━━━━━━━━━━")
		SendMessage(userID, "Ready to start fresh? Type 'menu' to see our products!")
		return
	}

	// Menu/Catalog
	if strings.Contains(msgLower, "menu") ||
		strings.Contains(msgLower, "catalog") ||
		strings.Contains(msgLower, "product") ||
		strings.Contains(msgLower, "show me") ||
		strings.Contains(msgLower, "မီနူး") ||
		strings.Contains(msgLower, "ပစ္စည်း") {
		showMenu(userID)
		return
	}

	// Help
	if strings.Contains(msgLower, "help") ||
		msgLower == "?" ||
		strings.Contains(msgLower, "how") ||
		strings.Contains(msgLower, "ကူညီ") {
		showHelp(userID)
		return
	}

	// Order History
	if strings.Contains(msgLower, "order") && (strings.Contains(msgLower, "history") || strings.Contains(msgLower, "my")) ||
		strings.Contains(msgLower, "ငါ့မှာတာ") {
		showOrderHistory(userID)
		return
	}

	// Product name matching - English + Burmese
	if state.State == "awaiting_product" || state.State == "greeting" {
		// Chocolate Cake
		if strings.Contains(msgLower, "chocolate") || strings.Contains(msgLower, "choco") || strings.Contains(msgLower, "ချောကလက်") {
			handlePostback(userID, "ORDER_CHOCOLATE_CAKE")
			return
		}
		// Vanilla Cake
		if strings.Contains(msgLower, "vanilla") || strings.Contains(msgLower, "ဗနီလာ") {
			handlePostback(userID, "ORDER_VANILLA_CAKE")
			return
		}
		// Red Velvet
		if strings.Contains(msgLower, "red velvet") || strings.Contains(msgLower, "velvet") || strings.Contains(msgLower, "အနီရောင်") {
			handlePostback(userID, "ORDER_RED_VELVET")
			return
		}
		// Coffee
		if strings.Contains(msgLower, "coffee") || strings.Contains(msgLower, "ကော်ဖီ") {
			handlePostback(userID, "ORDER_COFFEE")
			return
		}
		// Croissant
		if strings.Contains(msgLower, "croissant") || strings.Contains(msgLower, "ခရို့ဆန့်") {
			handlePostback(userID, "ORDER_CROISSANT")
			return
		}
		// Cinnamon Roll
		if strings.Contains(msgLower, "cinnamon") || strings.Contains(msgLower, "roll") || strings.Contains(msgLower, "ဆင်နမွန်") {
			handlePostback(userID, "ORDER_CINNAMON_ROLL")
			return
		}
		// Cupcake
		if strings.Contains(msgLower, "cupcake") || strings.Contains(msgLower, "cup cake") || strings.Contains(msgLower, "ကပ်ကိတ်") {
			handlePostback(userID, "ORDER_CHOCOLATE_CUPCAKE")
			return
		}
		// Bread
		if strings.Contains(msgLower, "bread") || strings.Contains(msgLower, "ပေါင်မုန့်") {
			handlePostback(userID, "ORDER_BREAD")
			return
		}
	}

	// Quantity matching - Natural language
	if state.State == "awaiting_quantity" {
		// Extract numbers from text: "I want 2", "give me 3", "၂ ခု"
		if strings.Contains(msgLower, "1") || strings.Contains(msgLower, "one") || strings.Contains(msgLower, "တစ်") {
			handlePostback(userID, "QTY_1")
			return
		}
		if strings.Contains(msgLower, "2") || strings.Contains(msgLower, "two") || strings.Contains(msgLower, "နှစ်") {
			handlePostback(userID, "QTY_2")
			return
		}
		if strings.Contains(msgLower, "3") || strings.Contains(msgLower, "three") || strings.Contains(msgLower, "သုံး") {
			handlePostback(userID, "QTY_3")
			return
		}
		if strings.Contains(msgLower, "4") || strings.Contains(msgLower, "four") || strings.Contains(msgLower, "လေး") {
			handlePostback(userID, "QTY_4")
			return
		}
		if strings.Contains(msgLower, "5") || strings.Contains(msgLower, "five") || strings.Contains(msgLower, "ငါး") {
			handlePostback(userID, "QTY_5")
			return
		}
	}

	// Delivery type matching
	if state.State == "awaiting_delivery_type" {
		if strings.Contains(msgLower, "pickup") || strings.Contains(msgLower, "pick up") || strings.Contains(msgLower, "ကိုယ်တိုင်ယူ") {
			handlePostback(userID, "PICKUP")
			return
		}
		if strings.Contains(msgLower, "delivery") || strings.Contains(msgLower, "deliver") || strings.Contains(msgLower, "ပို့") {
			handlePostback(userID, "DELIVERY")
			return
		}
	}

	// ========== END SMART MATCHING ==========

	// Handle special commands at any time (keep for exact matches)
	if msgLower == "menu" || msgLower == "catalog" {
		showMenu(userID)
		return
	}

	if msgLower == "help" || msgLower == "?" {
		showHelp(userID)
		return
	}

	if msgLower == "cancel" || msgLower == "reset" {
		ResetUserState(userID)
		SendMessage(userID, "Order cancelled. Type anything to start a new order!")
		return
	}

	if msgLower == "orders" || msgLower == "history" || msgLower == "my orders" {
		showOrderHistory(userID)
		return
	}

	// Process based on state
	switch state.State {
	case "language_selection":
		// User typed something instead of clicking language button
		showLanguageSelection(userID)
		return

	case "awaiting_name":
		// Validate name
		if len(messageText) < 2 {
			SendMessage(userID, "Please enter a valid name (at least 2 characters).")
			return
		}

		// User is providing their name
		state.CustomerName = messageText

		// Show typing indicator for better UX
		SendTypingIndicator(userID, true)

		// Ask: Pickup or Delivery?
		state.State = "awaiting_delivery_type"
		quickReplies := []QuickReply{
			{ContentType: "text", Title: "🏠 Pickup", Payload: "PICKUP"},
			{ContentType: "text", Title: "🚚 Delivery", Payload: "DELIVERY"},
			{ContentType: "text", Title: "⬅️ Back", Payload: "GO_BACK"},
			{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
		}
		SendQuickReplies(userID, fmt.Sprintf("Thanks %s! Would you like pickup or delivery?", state.CustomerName), quickReplies)

	case "awaiting_address":
		// Validate address
		if len(messageText) < 5 {
			SendMessage(userID, "Please enter a complete delivery address.")
			return
		}

		// User is providing delivery address
		state.Address = messageText
		state.State = "confirming"

		SendTypingIndicator(userID, true)
		showOrderSummary(userID)

	default:
		// For any other text input during button/quick-reply steps, guide them back
		if state.State == "language_selection" {
			// Show language selection again
			showLanguageSelection(userID)
		} else if state.State == "greeting" {
			// First message → start ordering flow
			startOrderingFlow(userID)
		} else if state.State == "awaiting_product" {
			// Re-show products if they type instead of clicking
			SendMessage(userID, "Please select a product using the buttons:")
			showProducts(userID)
		} else if state.State == "awaiting_quantity" {
			// Re-show quantity options
			SendMessage(userID, "Please select quantity using the buttons:")
			askQuantity(userID)
		} else if state.State == "awaiting_cart_decision" {
			// Re-show add more or checkout buttons
			SendMessage(userID, "Please choose an option:")
			askAddMore(userID)
		} else if state.State == "awaiting_delivery_type" {
			// Re-show delivery type options
			SendMessage(userID, "Please select pickup or delivery:")
			quickReplies := []QuickReply{
				{ContentType: "text", Title: "🏪 Pickup", Payload: "PICKUP"},
				{ContentType: "text", Title: "🚚 Delivery", Payload: "DELIVERY"},
				{ContentType: "text", Title: "⬅️ Back", Payload: "GO_BACK"},
				{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
			}
			SendQuickReplies(userID, fmt.Sprintf("Thanks %s! Would you like pickup or delivery?", state.CustomerName), quickReplies)
		} else if state.State == "confirming" {
			// Re-show order confirmation
			SendMessage(userID, "Please confirm your order:")
			showOrderSummary(userID)
		} else {
			SendMessage(userID, "Type 'menu' to see products or 'help' for assistance.")
		}
	}
}
