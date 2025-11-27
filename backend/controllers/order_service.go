package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"bakeflow/models"
)

// calculateDeliveryFee calculates delivery fee based on delivery type
func calculateDeliveryFee(deliveryType, address string) float64 {
	if deliveryType == "pickup" {
		return 0.00
	}

	// Simple distance-based fee (in production, use Google Maps API)
	addressLower := strings.ToLower(address)

	// Near locations - $3
	if strings.Contains(addressLower, "downtown") ||
		strings.Contains(addressLower, "yangon") ||
		strings.Contains(addressLower, "pickup at store") {
		return 3.00
	}

	// Far locations - $5
	if strings.Contains(addressLower, "airport") ||
		strings.Contains(addressLower, "suburb") {
		return 5.00
	}

	// Default delivery fee
	return 4.00
}

// calculateOrderTotals calculates subtotal, delivery fee, and total
func calculateOrderTotals(cart []CartItem, deliveryType, address string) (subtotal, deliveryFee, total float64) {
	// Calculate subtotal from cart
	for _, item := range cart {
		if product, exists := ProductCatalog[item.Product]; exists {
			priceStr := strings.ReplaceAll(product.Price, "$", "")
			if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
				subtotal += price * float64(item.Quantity)
			}
		}
	}

	// Calculate delivery fee
	deliveryFee = calculateDeliveryFee(deliveryType, address)

	// Total = subtotal + delivery fee
	total = subtotal + deliveryFee

	return subtotal, deliveryFee, total
}

// isBusinessOpen checks if current time is within business hours (8 AM - 8 PM)
func isBusinessOpen() bool {
	// TEMP: Always open for testing. Original logic (8AM-8PM) commented below.
	// now := time.Now()
	// hour := now.Hour()
	// return hour >= 8 && hour < 20
	return true
}

// getNextOpeningTime returns a formatted string of when the business opens next
func getNextOpeningTime() string {
	now := time.Now()
	hour := now.Hour()

	if hour < 8 {
		// Opens today at 8 AM
		return "8:00 AM today"
	} else {
		// Opens tomorrow at 8 AM
		tomorrow := now.Add(24 * time.Hour)
		return fmt.Sprintf("8:00 AM on %s", tomorrow.Format("Monday, Jan 2"))
	}
}

// confirmOrder saves the order to the database and sends confirmation
func confirmOrder(userID string) {
	state := GetUserState(userID)

	// Calculate total items
	totalItems := 0
	for _, item := range state.Cart {
		totalItems += item.Quantity
	}

	// Calculate totals (subtotal, delivery fee, total amount)
	subtotal, deliveryFee, totalAmount := calculateOrderTotals(state.Cart, state.DeliveryType, state.Address)

	// Create order in database (include Messenger sender ID for notifications)
	order := models.Order{
		CustomerName: state.CustomerName,
		DeliveryType: state.DeliveryType,
		Address:      state.Address,
		Status:       "pending",
		TotalItems:   totalItems,
		Subtotal:     subtotal,
		DeliveryFee:  deliveryFee,
		TotalAmount:  totalAmount,
		SenderID:     userID,
	}

	// Convert cart items to order items
	var orderItems []models.OrderItem
	for _, item := range state.Cart {
		// Get price from ProductCatalog
		price := 0.00
		if product, exists := ProductCatalog[item.Product]; exists {
			// Parse price string (e.g., "$25.00" → 25.00)
			priceStr := strings.ReplaceAll(product.Price, "$", "")
			if parsedPrice, err := strconv.ParseFloat(priceStr, 64); err == nil {
				price = parsedPrice
			}
		}

		orderItems = append(orderItems, models.OrderItem{
			Product:  item.Product,
			Quantity: item.Quantity,
			Price:    price,
		})
	}

	err := models.CreateOrder(&order, orderItems)
	if err != nil {
		log.Printf("❌ Error creating order: %v", err)
		SendMessage(userID, "😞 Sorry, there was an error placing your order. Please try again later.")
		ResetUserState(userID)
		return
	}

	deliveryIcon := "🏠"
	estimatedTime := "Ready in 15-20 minutes"
	if state.DeliveryType == "delivery" {
		deliveryIcon = "🚚"
		estimatedTime = "Delivered in 30-45 minutes"
	}

	// Build cart display with prices for confirmation
	cartDisplay := ""
	for _, item := range state.Cart {
		itemPrice := 0.00
		if product, exists := ProductCatalog[item.Product]; exists {
			priceStr := strings.ReplaceAll(product.Price, "$", "")
			if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
				itemPrice = price * float64(item.Quantity)
			}
		}
		cartDisplay += fmt.Sprintf("• %d× %s %s - $%.2f\n", item.Quantity, item.ProductEmoji, item.Product, itemPrice)
	}

	// Build pricing breakdown
	pricingBreakdown := fmt.Sprintf(
		"\n💰 **Pricing:**\n"+
			"Subtotal: $%.2f\n"+
			"Delivery Fee: $%.2f\n"+
			"━━━━━━━━━━━━\n"+
			"**Total: $%.2f**",
		order.Subtotal,
		order.DeliveryFee,
		order.TotalAmount,
	)

	// Send rich confirmation
	confirmation := fmt.Sprintf(
		"✅ **Order Confirmed!**\n\n"+
			"Order #%d\n\n"+
			"🛒 **Your Order:**\n"+
			"%s"+
			"%s\n\n"+
			"👤 %s\n"+
			"%s %s\n"+
			"📍 %s\n"+
			"📊 Status: %s\n\n"+
			"⏱ %s\n\n"+
			"Thank you for choosing BakeFlow! 🎉\n\n"+
			"Type 'menu' to order more, or 'orders' to view history.",
		order.ID,
		cartDisplay,
		pricingBreakdown,
		state.CustomerName,
		deliveryIcon, strings.Title(state.DeliveryType),
		order.Address,
		strings.Title(order.Status),
		estimatedTime,
	)
	SendMessage(userID, confirmation)

	// Reset state for next order
	ResetUserState(userID)
}

// handleReorder pre-fills cart with items from previous order
func handleReorder(userID string, orderID int) {
	// Get the order
	order, err := models.GetOrderByID(orderID)
	if err != nil {
		log.Printf("❌ Error fetching order for reorder: %v", err)
		SendMessage(userID, "😞 Sorry, couldn't load that order. Please try again.")
		return
	}

	// Reset state and pre-fill cart
	state := GetUserState(userID)
	state.Cart = []CartItem{}

	// Convert order items to cart items
	for _, item := range order.Items {
		emoji := "🍰"
		if product, exists := ProductCatalog[item.Product]; exists {
			emoji = product.Emoji
		}

		state.Cart = append(state.Cart, CartItem{
			Product:      item.Product,
			ProductEmoji: emoji,
			Quantity:     item.Quantity,
		})
	}

	// Calculate total items
	totalItems := 0
	for _, item := range state.Cart {
		totalItems += item.Quantity
	}

	// Send confirmation message
	SendMessage(userID, fmt.Sprintf("🔄 **Reordering from Order #%d**\n\n✅ Added %d items to your cart!", order.ID, totalItems))

	// Show cart
	showCart(userID)

	// Ask for checkout
	time.Sleep(1 * time.Second)
	askName(userID)
}

// askForRating sends rating request with star buttons
func askForRating(userID string, orderID int) {
	state := GetUserState(userID)
	state.State = "awaiting_rating"
	state.CurrentProduct = strconv.Itoa(orderID) // Temporarily store orderID

	ratingMsg := "⭐ **How was your order?**\n\n" +
		"We'd love to hear your feedback!\n" +
		"Please rate your experience:"

	if state.Language == "my" {
		ratingMsg = "⭐ **အော်ဒါက ဘယ်လိုလဲ?**\n\n" +
			"သင့်ရဲ့ အကြံပြုချက်ကို ကြားလိုပါတယ်!\n" +
			"သင့်အတွေ့အကြုံကို အဆင့်သတ်မှတ်ပေးပါ:"
	}

	quickReplies := []QuickReply{
		{ContentType: "text", Title: "⭐ 1 Star - Poor", Payload: "RATING_1"},
		{ContentType: "text", Title: "⭐⭐ 2 Stars", Payload: "RATING_2"},
		{ContentType: "text", Title: "⭐⭐⭐ 3 Stars", Payload: "RATING_3"},
		{ContentType: "text", Title: "⭐⭐⭐⭐ 4 Stars", Payload: "RATING_4"},
		{ContentType: "text", Title: "⭐⭐⭐⭐⭐ 5 Stars - Excellent!", Payload: "RATING_5"},
		{ContentType: "text", Title: "Skip", Payload: "SKIP_RATING"},
	}

	SendQuickReplies(userID, ratingMsg, quickReplies)
}

// handleRating saves customer rating
func handleRating(userID string, stars int) {
	state := GetUserState(userID)

	// Get orderID from temporary storage
	orderID, err := strconv.Atoi(state.CurrentProduct)
	if err != nil {
		SendMessage(userID, "😞 Sorry, something went wrong. Please try again.")
		ResetUserState(userID)
		return
	}

	// Save rating to database
	rating := models.Rating{
		OrderID: orderID,
		UserID:  userID,
		Stars:   stars,
		Comment: "",
	}

	err = models.CreateRating(&rating)
	if err != nil {
		log.Printf("❌ Error saving rating: %v", err)
		SendMessage(userID, "😞 Sorry, couldn't save your rating. Please try again later.")
		return
	}

	// Send thank you message
	thankYouMsg := ""
	if stars >= 4 {
		thankYouMsg = "🎉 **Thank you so much!**\n\n" +
			"We're thrilled you loved your order! ⭐⭐⭐⭐⭐\n\n" +
			"Your feedback means the world to us. Looking forward to serving you again! 🍰"

		if state.Language == "my" {
			thankYouMsg = "🎉 **အရမ်းကျေးဇူးတင်ပါတယ်!**\n\n" +
				"သင့် အော်ဒါကို နှစ်သက်တာ သိရတာ အရမ်းဝမ်းသာပါတယ်! ⭐⭐⭐⭐⭐\n\n" +
				"သင့်ရဲ့ အကြံပြုချက်က ကျွန်ုပ်တို့အတွက် အရမ်းအရေးကြီးပါတယ်။ နောက်တစ်ခါ ထပ်ဆောင်ရွက်ပေးဖို့ မျှော်လင့်နေပါတယ်! 🍰"
		}
	} else if stars == 3 {
		thankYouMsg = "😊 **Thank you for your feedback!**\n\n" +
			"We appreciate your honesty. We're always working to improve!\n\n" +
			"Type 'menu' to order again! 🍰"

		if state.Language == "my" {
			thankYouMsg = "😊 **သင့်အကြံပြုချက်အတွက် ကျေးဇူးတင်ပါတယ်!**\n\n" +
				"သင့်ရိုးသားမှုကို တန်ဖိုးထားပါတယ်။ ကျွန်ုပ်တို့ အမြဲတမ်း တိုးတက်အောင် လုပ်ဆောင်နေပါတယ်!\n\n" +
				"'မီနူး' လို့ရိုက်ပြီး ထပ်မှာလိုက်ပါ! 🍰"
		}
	} else {
		thankYouMsg = "😔 **We're sorry you weren't satisfied.**\n\n" +
			"Your feedback is important to us. We'll do better next time!\n\n" +
			"Please give us another chance. Type 'menu' to order! 🍰"

		if state.Language == "my" {
			thankYouMsg = "😔 **သင် မကျေနပ်မှုအတွက် တောင်းပန်ပါတယ်။**\n\n" +
				"သင့်အကြံပြုချက်က ကျွန်ုပ်တို့အတွက် အရေးကြီးပါတယ်။ နောက်တစ်ခါ ပိုကောင်းအောင် လုပ်ပါမယ်!\n\n" +
				"နောက်တစ်ခါ အခွင့်အရေးပေးပါ။ 'မီနူး' လို့ရိုက်ပြီး မှာလိုက်ပါ! 🍰"
		}
	}

	SendMessage(userID, thankYouMsg)
	ResetUserState(userID)
}

// checkBusinessHours checks if ordering is allowed (business hours check)
func checkBusinessHours(userID string) bool {
	if isBusinessOpen() {
		return true
	}

	state := GetUserState(userID)
	closedMsg := "🔒 **We're Currently Closed**\n\n" +
		"Business Hours: 8:00 AM - 8:00 PM\n\n" +
		fmt.Sprintf("We'll be open again at %s.\n\n", getNextOpeningTime()) +
		"You can browse our menu, but ordering is temporarily unavailable.\n\n" +
		"See you soon! 🍰"

	if state.Language == "my" {
		closedMsg = "🔒 **ကျွန်ုပ်တို့ လောလောဆယ် ပိတ်နေပါတယ်**\n\n" +
			"စီးပွားရေး အချိန်: နံနက် 8:00 - ည 8:00\n\n" +
			fmt.Sprintf("ကျွန်ုပ်တို့ %s မှာ ပြန်ဖွင့်ပါမယ်။\n\n", getNextOpeningTime()) +
			"မီနူးကို ကြည့်နိုင်ပေမယ့် မှာယူခြင်းကို ယာယီ မရနိုင်ပါဘူး။\n\n" +
			"မကြာခင် တွေ့ရအောင်! 🍰"
	}

	SendMessage(userID, closedMsg)
	return false
}
