package controllers

import (
	"fmt"
	"log"
	"strings"

	"bakeflow/models"
)

// handleMessage processes text messages from users
// handleMessage moved to `message_handler.go`.

// handlePostback processes button clicks (postback payloads)
// Full implementation moved to `postback_handler.go`.

// showHelp displays help information
// showHelp displays ordering instructions
func showHelp(userID string) {
	help := "🆘 *How to Order* / *မှာယူနည်း*\n\n" +
		"1️⃣ Choose what you'd like to order\n" +
		"    လိုချင်တဲ့ပစ္စည်းကို ရွေးပါ\n" +
		"2️⃣ Select quantity / အရေအတွက် ရွေးပါ\n" +
		"3️⃣ Enter your name / နာမည် ထည့်ပါ\n" +
		"4️⃣ Choose pickup or delivery\n" +
		"    ကိုယ်တိုင်ယူမလား ပို့မလား ရွေးပါ\n" +
		"5️⃣ Confirm your order / အတည်ပြုပါ\n\n" +
		"*You can type naturally:* / *သဘာဝအတိုင်း စာရိုက်နိုင်ပါတယ်*\n" +
		"• \"I want chocolate cake\" / \"ချောကလက်ကိတ်လိုချင်တယ်\"\n" +
		"• \"Give me 2\" / \"2 ခု ပေးပါ\"\n" +
		"• \"I want to cancel\" / \"ပယ်ဖျက်ချင်တယ်\"\n" +
		"• \"Show menu\" / \"မီနူး ပြပါ\"\n\n" +
		"*Quick Commands:*\n" +
		"• 'menu' - View products\n" +
		"• 'cancel' - Start over\n" +
		"• 'help' - Show this message"

	SendMessage(userID, help)

	// After showing help, start the ordering flow
	startOrderingFlow(userID)
}

// showOrderHistory displays recent orders (mock implementation)
// goBack handles the "Go Back" navigation
func goBack(userID string) {
	state := GetUserState(userID)

	switch state.State {
	case "awaiting_quantity":
		// Go back to product selection
		showProducts(userID)

	case "awaiting_cart_decision":
		// Go back to cart
		showCart(userID)

	case "awaiting_name":
		// Go back to cart decision
		askAddMore(userID)

	case "awaiting_delivery_type":
		// Go back to name input
		state.State = "awaiting_name"
		quickReplies := []QuickReply{
			{ContentType: "text", Title: "⬅️ Back to Cart", Payload: "GO_BACK"},
			{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
		}
		SendQuickReplies(userID, "What's your name?", quickReplies)

	case "awaiting_address":
		// Go back to pickup/delivery selection
		state.State = "awaiting_delivery_type"
		quickReplies := []QuickReply{
			{ContentType: "text", Title: "🏠 Pickup", Payload: "PICKUP"},
			{ContentType: "text", Title: "🚚 Delivery", Payload: "DELIVERY"},
			{ContentType: "text", Title: "⬅️ Back", Payload: "GO_BACK"},
			{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
		}
		SendQuickReplies(userID, fmt.Sprintf("Thanks %s! Would you like pickup or delivery?", state.CustomerName), quickReplies)

	case "confirming":
		// Go back to address or delivery type
		if state.DeliveryType == "delivery" {
			state.State = "awaiting_address"
			quickReplies := []QuickReply{
				{ContentType: "text", Title: "⬅️ Back", Payload: "GO_BACK"},
				{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
			}
			SendQuickReplies(userID, "Please type your delivery address:\n(Street, City, ZIP)", quickReplies)
		} else {
			state.State = "awaiting_delivery_type"
			quickReplies := []QuickReply{
				{ContentType: "text", Title: "🏠 Pickup", Payload: "PICKUP"},
				{ContentType: "text", Title: "🚚 Delivery", Payload: "DELIVERY"},
				{ContentType: "text", Title: "⬅️ Back", Payload: "GO_BACK"},
				{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
			}
			SendQuickReplies(userID, fmt.Sprintf("Thanks %s! Would you like pickup or delivery?", state.CustomerName), quickReplies)
		}

	default:
		// If no clear back action, go to main menu
		startOrderingFlow(userID)
	}
}

// ========== NEW FEATURES ==========

// Business logic moved to `order_service.go`.

// showOrderHistory displays user's past orders with beautiful card design
func showOrderHistory(userID string) {
	// Get all orders (in future, filter by userID)
	orders, err := models.GetAllOrders()
	if err != nil {
		log.Printf("❌ Error fetching orders: %v", err)
		SendMessage(userID, "😞 Sorry, couldn't load your order history. Please try again later.")
		return
	}

	// Check if empty
	if len(orders) == 0 {
		state := GetUserState(userID)
		emptyMsg := "🛒 **No Orders Yet!**\n\n" +
			"You haven't placed any orders with us.\n\n" +
			"Ready to try our delicious baked goods?\n\n" +
			"Type 'menu' to start ordering! 🍰"

		if state.Language == "my" {
			emptyMsg = "🛒 **မှာထားမှုမရှိသေးပါ!**\n\n" +
				"သင် ကျွန်ုပ်တို့နှင့် မှာထားမှုမလုပ်ရသေးပါ။\n\n" +
				"ကျွန်ုပ်တို့ရဲ့ အရသာရှိတဲ့ မုန့်တွေကို စမ်းကြည့်ဖို့ အဆင်သင့်လား?\n\n" +
				"'မီနူး' လို့ရိုက်ပြီး မှာယူလိုက်ပါ! 🍰"
		}

		SendMessage(userID, emptyMsg)
		return
	}

	// Show recent orders as cards (limit to 5 most recent)
	displayOrders := orders
	if len(orders) > 5 {
		displayOrders = orders[:5]
	}

	var elements []Element
	for _, order := range displayOrders {
		// Build items list
		itemsList := ""
		for i, item := range order.Items {
			if i < 3 {
				emoji := "🍰"
				if product, exists := ProductCatalog[item.Product]; exists {
					emoji = product.Emoji
				}
				itemsList += fmt.Sprintf("%d× %s %s\n", item.Quantity, emoji, item.Product)
			}
		}
		if len(order.Items) > 3 {
			itemsList += fmt.Sprintf("...and %d more items\n", len(order.Items)-3)
		}

		// Status badge
		statusEmoji := "⏳"
		statusText := "Pending"
		switch order.Status {
		case "pending":
			statusEmoji = "⏳"
			statusText = "Pending"
		case "preparing":
			statusEmoji = "👨‍🍳"
			statusText = "Preparing"
		case "ready":
			statusEmoji = "✅"
			statusText = "Ready"
		case "delivered":
			statusEmoji = "🎉"
			statusText = "Delivered"
		case "completed":
			statusEmoji = "✔️"
			statusText = "Completed"
		}

		// Delivery icon
		deliveryIcon := "🏠"
		if order.DeliveryType == "delivery" {
			deliveryIcon = "🚚"
		}

		// Format date
		dateStr := order.CreatedAt.Format("Jan 2, 3:04 PM")

		// Build subtitle
		subtitle := fmt.Sprintf("%s %s • %s %s\n%s\nTotal: $%.2f",
			statusEmoji, statusText,
			deliveryIcon, strings.Title(order.DeliveryType),
			dateStr,
			order.TotalAmount)

		element := Element{
			Title:    fmt.Sprintf("Order #%d - %s", order.ID, order.CustomerName),
			Subtitle: subtitle + "\n\n" + itemsList,
			Buttons: []Button{
				{
					Type:    "postback",
					Title:   "🔄 Reorder",
					Payload: fmt.Sprintf("REORDER_%d", order.ID),
				},
				{
					Type:    "postback",
					Title:   "⭐ Rate",
					Payload: fmt.Sprintf("RATE_ORDER_%d", order.ID),
				},
			},
		}

		elements = append(elements, element)
	}

	SendMessage(userID, fmt.Sprintf("📋 **Your Recent Orders** (Showing %d of %d)", len(displayOrders), len(orders)))
	SendGenericTemplate(userID, elements)
}

// Rating handling moved to `order_service.go`.

// Business-hour checks moved to `order_service.go`.
