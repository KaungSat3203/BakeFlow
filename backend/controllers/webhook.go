package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"bakeflow/models"
)

// getProductElements returns all product carousel elements
func getProductElements() []Element {
	return []Element{
		{
			Title:    ProductCatalog["Chocolate Cake"].Emoji + " " + ProductCatalog["Chocolate Cake"].Name,
			ImageURL: ProductCatalog["Chocolate Cake"].ImageURL,
			Subtitle: ProductCatalog["Chocolate Cake"].Description + " • " + ProductCatalog["Chocolate Cake"].Price,
			Buttons: []Button{
				{Type: "postback", Title: "🛒 Order", Payload: "ORDER_CHOCOLATE_CAKE"},
			},
		},
		{
			Title:    ProductCatalog["Vanilla Cake"].Emoji + " " + ProductCatalog["Vanilla Cake"].Name,
			ImageURL: ProductCatalog["Vanilla Cake"].ImageURL,
			Subtitle: ProductCatalog["Vanilla Cake"].Description + " • " + ProductCatalog["Vanilla Cake"].Price,
			Buttons: []Button{
				{Type: "postback", Title: "🛒 Order", Payload: "ORDER_VANILLA_CAKE"},
			},
		},
		{
			Title:    ProductCatalog["Red Velvet Cake"].Emoji + " " + ProductCatalog["Red Velvet Cake"].Name,
			ImageURL: ProductCatalog["Red Velvet Cake"].ImageURL,
			Subtitle: ProductCatalog["Red Velvet Cake"].Description + " • " + ProductCatalog["Red Velvet Cake"].Price,
			Buttons: []Button{
				{Type: "postback", Title: "🛒 Order", Payload: "ORDER_RED_VELVET"},
			},
		},
		{
			Title:    ProductCatalog["Coffee"].Emoji + " " + ProductCatalog["Coffee"].Name,
			ImageURL: ProductCatalog["Coffee"].ImageURL,
			Subtitle: ProductCatalog["Coffee"].Description + " • " + ProductCatalog["Coffee"].Price,
			Buttons: []Button{
				{Type: "postback", Title: "🛒 Order", Payload: "ORDER_COFFEE"},
			},
		},
		{
			Title:    ProductCatalog["Croissant"].Emoji + " " + ProductCatalog["Croissant"].Name,
			ImageURL: ProductCatalog["Croissant"].ImageURL,
			Subtitle: ProductCatalog["Croissant"].Description + " • " + ProductCatalog["Croissant"].Price,
			Buttons: []Button{
				{Type: "postback", Title: "🛒 Order", Payload: "ORDER_CROISSANT"},
			},
		},
		{
			Title:    ProductCatalog["Cinnamon Roll"].Emoji + " " + ProductCatalog["Cinnamon Roll"].Name,
			ImageURL: ProductCatalog["Cinnamon Roll"].ImageURL,
			Subtitle: ProductCatalog["Cinnamon Roll"].Description + " • " + ProductCatalog["Cinnamon Roll"].Price,
			Buttons: []Button{
				{Type: "postback", Title: "🛒 Order", Payload: "ORDER_CINNAMON_ROLL"},
			},
		},
		{
			Title:    ProductCatalog["Chocolate Cupcake"].Emoji + " " + ProductCatalog["Chocolate Cupcake"].Name,
			ImageURL: ProductCatalog["Chocolate Cupcake"].ImageURL,
			Subtitle: ProductCatalog["Chocolate Cupcake"].Description + " • " + ProductCatalog["Chocolate Cupcake"].Price,
			Buttons: []Button{
				{Type: "postback", Title: "🛒 Order", Payload: "ORDER_CHOCOLATE_CUPCAKE"},
			},
		},
		{
			Title:    ProductCatalog["Bread"].Emoji + " " + ProductCatalog["Bread"].Name,
			ImageURL: ProductCatalog["Bread"].ImageURL,
			Subtitle: ProductCatalog["Bread"].Description + " • " + ProductCatalog["Bread"].Price,
			Buttons: []Button{
				{Type: "postback", Title: "🛒 Order", Payload: "ORDER_BREAD"},
			},
		},
	}
}

// showAbout displays company information and help instructions in user's language
func showAbout(userID string) {
	state := GetUserState(userID)
	var aboutMsg, helpMsg string
	
	if state.Language == "my" {
		aboutMsg = "🏪 ကျွန်ုပ်တို့အကြောင်း\n\n" +
			"BakeFlow သည် လတ်ဆတ်သော မုန့်များကို နေ့စဉ် ဖုတ်လုပ်သော မုန့်ဆိုင်ဖြစ်ပါသည်။\n\n" +
			"🎂 ကျွန်ုပ်တို့၏ အထူးမုန့်များ:\n" +
			"• ချောကလက် ကိတ်မုန့်\n" +
			"• ဗနီလာ ကိတ်မုန့်\n" +
			"• ဆော့ဘီ ကိတ်မုန့်\n" +
			"• ချိစ်ကိတ်မုန့်\n" +
			"• နီမုန့်\n" +
			"• ချောကလက် ကွတ်ကီး\n" +
			"• ဗာတာကွတ်ကီး\n" +
			"• အာလုမွန့်\n\n" +
			"📍 တည်နေရာ: ရန်ကုန်မြို့\n" +
			"⏰ ဖွင့်ချိန်: နံနက် 8:00 - ညနေ 8:00\n" +
			"📞 ဆက်သွယ်ရန်: +95 9 XXX XXX XXX"
		
		helpMsg = "\n\n❓ အသုံးပြုနည်း\n\n" +
			"သဘာဝဘာသာစကားဖြင့် ရိုက်နိုင်ပါတယ်:\n\n" +
			"• \"မီနူး\" သို့မဟုတ် \"မုန့်များ\"\n" +
			"• \"ချောကလက်ကိတ်မုန့်လိုချင်တယ်\"\n" +
			"• \"နှစ်ခု\" သို့မဟုတ် \"၂\"\n" +
			"• \"ပို့ပေးပါ\" သို့မဟုတ် \"ကိုယ်တိုင်ယူမယ်\"\n" +
			"• \"ပယ်ဖျက်\" သို့မဟုတ် \"အစကနေစမယ်\"\n\n" +
			"🛒 အော်ဒါမှာရန် 'မီနူး' လို့ရိုက်ပါ!"
	} else {
		aboutMsg = "🏪 About Us\n\n" +
			"BakeFlow is your neighborhood bakery, baking fresh daily!\n\n" +
			"🎂 Our Specialties:\n" +
			"• Chocolate Cake\n" +
			"• Vanilla Cake\n" +
			"• Strawberry Cake\n" +
			"• Cheesecake\n" +
			"• Red Velvet Cake\n" +
			"• Chocolate Cookies\n" +
			"• Butter Cookies\n" +
			"• Almond Croissant\n\n" +
			"📍 Location: Yangon, Myanmar\n" +
			"⏰ Hours: 8:00 AM - 8:00 PM\n" +
			"📞 Contact: +95 9 XXX XXX XXX"
		
		helpMsg = "\n\n❓ How to Use\n\n" +
			"You can type naturally:\n\n" +
			"• \"menu\" or \"show products\"\n" +
			"• \"I want chocolate cake\"\n" +
			"• \"two\" or \"2\"\n" +
			"• \"delivery please\" or \"pickup\"\n" +
			"• \"cancel\" or \"start over\"\n\n" +
			"🛒 Type 'menu' to start ordering!"
	}
	
	SendMessage(userID, aboutMsg + helpMsg)
}

// showLanguageSelection shows language choice at the beginning
func showLanguageSelection(userID string) {
	state := GetUserState(userID)
	state.State = "language_selection"
	
	welcomeMsg := "Hi there! 👋 မင်္ဂလာပါ! 👋\n\n" +
		"I'm BakeFlow Bot, your virtual bakery assistant (Beta). " +
		"I'm still learning, so I might not have all the answers yet, but I'll try to assist you the best I can! 🍰\n\n" +
		"ကျွန်တော် BakeFlow Bot ပါ၊ သင့်ရဲ့ မုန့်ဆိုင် အကူအညီပေး စက်ရုပ်ပါ (စမ်းသပ်ဗားရှင်း)။ " +
		"ကျွန်တော် ယခုတော့ သင်ယူနေဆဲဖြစ်တဲ့အတွက် အားလုံးကို မဖြေနိုင်သေးပေမယ့် တတ်နိုင်သမျှ အကောင်းဆုံး ကူညီပေးပါမယ်နော်! 🍰\n\n" +
		"Please select your language to get started.\n" +
		"စတင်ဖို့ ဘာသာစကားကို ရွေးချယ်ပါ။"
	
	SendMessage(userID, welcomeMsg)
	
	quickReplies := []QuickReply{
		{ContentType: "text", Title: "🇬🇧 English", Payload: "LANG_EN"},
		{ContentType: "text", Title: "🇲🇲 မြန်မာ", Payload: "LANG_MY"},
	}
	SendQuickReplies(userID, "Choose your language / ဘာသာစကား ရွေးပါ:", quickReplies)
}

// startOrderingFlow begins the ordering process with welcome message and simple menu
func startOrderingFlow(userID string) {
	state := GetUserState(userID)
	state.State = "main_menu"
	
	// Send welcome message with simple button menu
	if state.Language == "my" {
		SendMessage(userID, "🍰 BakeFlow မှ ကြိုဆိုပါတယ်!")
		showMainMenuSimple(userID)
	} else {
		SendMessage(userID, "🍰 Welcome to BakeFlow!")
		showMainMenuSimple(userID)
	}
}

// showMainMenu displays main menu as cards (like your screenshot)
func showMainMenu(userID string) {
	state := GetUserState(userID)
	
	var elements []Element
	
	if state.Language == "my" {
		elements = []Element{
			{
				Title:    "🛒 အော်ဒါမှာမယ်",
				Subtitle: "ကျွန်ုပ်တို့၏ လတ်ဆတ်သော မုန့်များကို ကြည့်ရှုပါ",
				ImageURL: "https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=300&h=200&fit=crop",
				Buttons: []Button{
					{
						Type:    "postback",
						Title:   "လုပ်ဆောင်မည်",
						Payload: "MENU_ORDER_PRODUCTS",
					},
				},
			},
			{
				Title:    "ℹ️ အကြောင်းနှင့်အကူအညီ",
				Subtitle: "ကျွန်ုပ်တို့အကြောင်းနှင့် အသုံးပြုနည်း",
				ImageURL: "https://images.unsplash.com/photo-1556910103-1c02745aae4d?w=300&h=200&fit=crop",
				Buttons: []Button{
					{
						Type:    "postback",
						Title:   "ဖတ်ရှုမည်",
						Payload: "MENU_ABOUT",
					},
				},
			},
			{
				Title:    "🌐 ဘာသာပြောင်းမယ်",
				Subtitle: "English သို့ ပြောင်းလဲရန်",
				ImageURL: "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=300&h=200&fit=crop",
				Buttons: []Button{
					{
						Type:    "postback",
						Title:   "ပြောင်းမည်",
						Payload: "MENU_CHANGE_LANG",
					},
				},
			},
		}
	} else {
		elements = []Element{
			{
				Title:    "� Order Now",
				Subtitle: "Browse our fresh baked goods",
				ImageURL: "https://images.unsplash.com/photo-1578985545062-69928b1d9587?w=300&h=200&fit=crop",
				Buttons: []Button{
					{
						Type:    "postback",
						Title:   "Start Order",
						Payload: "MENU_ORDER_PRODUCTS",
					},
				},
			},
			{
				Title:    "ℹ️ About & Help",
				Subtitle: "Learn about us and how to order",
				ImageURL: "https://images.unsplash.com/photo-1556910103-1c02745aae4d?w=300&h=200&fit=crop",
				Buttons: []Button{
					{
						Type:    "postback",
						Title:   "Learn More",
						Payload: "MENU_ABOUT",
					},
				},
			},
			{
				Title:    "🌐 Change Language",
				Subtitle: "Switch to Myanmar language",
				ImageURL: "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=300&h=200&fit=crop",
				Buttons: []Button{
					{
						Type:    "postback",
						Title:   "Switch",
						Payload: "MENU_CHANGE_LANG",
					},
				},
			},
		}
	}
	
	SendGenericTemplate(userID, elements)
}

// showProducts displays the product catalog
func showProducts(userID string) {
	// Check business hours before showing products
	if !checkBusinessHours(userID) {
		return
	}
	
	state := GetUserState(userID)
	state.State = "awaiting_product"
	SendGenericTemplate(userID, getProductElements())
}

// askQuantity asks how many items the user wants
func askQuantity(userID string) {
	state := GetUserState(userID)
	
	quickReplies := []QuickReply{
		{ContentType: "text", Title: "1", Payload: "QTY_1"},
		{ContentType: "text", Title: "2", Payload: "QTY_2"},
		{ContentType: "text", Title: "3", Payload: "QTY_3"},
		{ContentType: "text", Title: "4", Payload: "QTY_4"},
		{ContentType: "text", Title: "5", Payload: "QTY_5"},
		{ContentType: "text", Title: "⬅️ Back", Payload: "GO_BACK"},
		{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
	}
	SendQuickReplies(userID, fmt.Sprintf("How many %s %s would you like?", state.CurrentEmoji, state.CurrentProduct), quickReplies)
}

// askName asks for the customer's name
func askName(userID string) {
	state := GetUserState(userID)
	state.State = "awaiting_name"
	
	// Send a message with quick reply options to go back
	quickReplies := []QuickReply{
		{ContentType: "text", Title: "⬅️ Back to Cart", Payload: "GO_BACK"},
		{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
	}
	SendQuickReplies(userID, "Great! What's your name?", quickReplies)
}

// addToCart adds the current product to the cart
func addToCart(userID string) {
	state := GetUserState(userID)
	
	// Add current product to cart
	cartItem := CartItem{
		Product:      state.CurrentProduct,
		ProductEmoji: state.CurrentEmoji,
		Quantity:     state.CurrentQuantity,
	}
	state.Cart = append(state.Cart, cartItem)
	
	// Clear current product
	state.CurrentProduct = ""
	state.CurrentEmoji = ""
	state.CurrentQuantity = 0
	
	// Ask if they want to add more
	askAddMore(userID)
}

// askAddMore asks if customer wants to add more items or checkout
func askAddMore(userID string) {
	state := GetUserState(userID)
	
	// Calculate total items in cart
	totalItems := 0
	for _, item := range state.Cart {
		totalItems += item.Quantity
	}
	
	// Show what was just added
	lastItem := state.Cart[len(state.Cart)-1]
	message := fmt.Sprintf("✅ %d× %s %s added\n\nCart: %d items", 
		lastItem.Quantity, lastItem.ProductEmoji, lastItem.Product, totalItems)
	
	quickReplies := []QuickReply{
		{ContentType: "text", Title: "Add More", Payload: "ADD_MORE_ITEMS"},
		{ContentType: "text", Title: fmt.Sprintf("Checkout (%d)", totalItems), Payload: "CHECKOUT"},
		{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
	}
	
	state.State = "awaiting_cart_decision"
	SendQuickReplies(userID, message, quickReplies)
}

// showCart displays current cart contents
func showCart(userID string) {
	state := GetUserState(userID)
	
	if len(state.Cart) == 0 {
		SendMessage(userID, "🛒 Your cart is empty!\n\nLet's start ordering!")
		startOrderingFlow(userID)
		return
	}
	
	// Build cart display
	cartDisplay := "🛒 **Your Cart:**\n\n"
	totalItems := 0
	
	for _, item := range state.Cart {
		cartDisplay += fmt.Sprintf("• %d× %s %s\n", item.Quantity, item.ProductEmoji, item.Product)
		totalItems += item.Quantity
	}
	
	cartDisplay += fmt.Sprintf("\n**Total Items:** %d", totalItems)
	
	SendMessage(userID, cartDisplay)
}

// VerifyWebhook handles Facebook Messenger webhook verification (GET requests)
// 
// Facebook sends a GET request with these query parameters:
// - hub.mode=subscribe
// - hub.verify_token=<your_verify_token>
// - hub.challenge=<random_string>
//
// COMMON VERIFICATION FAILURES AND HOW TO FIX:
// 1. "Callback URL or verify token couldn't be validated"
//    - VERIFY_TOKEN in .env doesn't match the one in Meta Developer Console
//    - .env file not loaded properly (check godotenv.Load() in main.go)
//    - ngrok URL is wrong or expired (get new URL with `ngrok http 8080`)
//    - Server not running on the correct port
//
// 2. "URL is not available"
//    - Server not running (run `go run main.go`)
//    - Firewall blocking port 8080
//    - ngrok not forwarding to localhost:8080
//
// 3. "The URL couldn't be validated"
//    - Server returning wrong status code
//    - Not returning the hub.challenge value
//    - HTTPS required (ngrok provides this automatically)
func VerifyWebhook(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")
	
	// Get the verify token from environment
	verifyToken := os.Getenv("VERIFY_TOKEN")
	
	// Debug logging (helps troubleshoot verification issues)
	log.Printf("========== WEBHOOK VERIFICATION ATTEMPT ==========")
	log.Printf("Mode: %s", mode)
	log.Printf("Token received: %s", token)
	log.Printf("Token expected: %s", verifyToken)
	log.Printf("Challenge: %s", challenge)
	log.Printf("Full URL: %s", r.URL.String())
	
	// Check if verify token is loaded from .env
	if verifyToken == "" {
		log.Println("❌ ERROR: VERIFY_TOKEN is empty! Check your .env file")
		http.Error(w, "Server configuration error", http.StatusInternalServerError)
		return
	}
	
	// Verify that mode and token are correct
	if mode == "subscribe" && token == verifyToken {
		log.Println("✅ Webhook verified successfully!")
		
		// Respond with the challenge token from the request
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
		return
	}
	
	// Verification failed
	log.Println("❌ Webhook verification FAILED")
	if mode != "subscribe" {
		log.Printf("   - Wrong mode: got '%s', expected 'subscribe'", mode)
	}
	if token != verifyToken {
		log.Printf("   - Token mismatch!")
		log.Printf("   - Received: '%s'", token)
		log.Printf("   - Expected: '%s'", verifyToken)
	}
	
	http.Error(w, "Forbidden", http.StatusForbidden)
}

// ReceiveWebhook handles incoming messages from Facebook Messenger (POST requests)
//
// Facebook sends POST requests when users message your page with this structure:
// {
//   "object": "page",
//   "entry": [{
//     "id": "page_id",
//     "time": 1234567890,
//     "messaging": [{
//       "sender": {"id": "user_id"},
//       "recipient": {"id": "page_id"},
//       "timestamp": 1234567890,
//       "message": {
//         "mid": "message_id",
//         "text": "Hello!"
//       }
//     }]
//   }]
// }
func ReceiveWebhook(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("❌ Error reading request body: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	
	log.Println("========== INCOMING WEBHOOK POST ==========")
	log.Printf("Raw body: %s", string(body))
	
	// Parse the webhook payload
	var webhook WebhookPayload
	if err := json.Unmarshal(body, &webhook); err != nil {
		log.Printf("❌ Error parsing JSON: %v", err)
		// Still return 200 OK to Facebook so they don't retry
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("EVENT_RECEIVED"))
		return
	}
	
	// Verify this is a page subscription
	if webhook.Object != "page" {
		log.Printf("⚠️  Not a page webhook: %s", webhook.Object)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("EVENT_RECEIVED"))
		return
	}
	
	// Process each entry
	for _, entry := range webhook.Entry {
		log.Printf("Processing entry from page ID: %s", entry.ID)
		
		// Process each messaging event
		for _, event := range entry.Messaging {
			senderID := event.Sender.ID
			
			// Check if this is a quick reply (button click from quick reply)
			if event.Message.QuickReply != nil && event.Message.QuickReply.Payload != "" {
				log.Printf("⚡ Quick Reply from %s: %s", senderID, event.Message.QuickReply.Payload)
				handlePostback(senderID, event.Message.QuickReply.Payload)
				continue
			}
			
			// Check if this is a message event (text input)
			if event.Message.Text != "" {
				log.Printf("📨 Message from %s: %s", senderID, event.Message.Text)
				handleMessage(senderID, strings.TrimSpace(event.Message.Text))
				continue
			}
			
			// Check for postback (button clicks from structured messages)
			if event.Postback.Payload != "" {
				log.Printf("🔘 Postback from %s: %s", senderID, event.Postback.Payload)
				handlePostback(senderID, event.Postback.Payload)
			}
		}
	}
	
	// Always return 200 OK to Facebook within 20 seconds
	// Otherwise Facebook will retry the webhook multiple times
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("EVENT_RECEIVED"))
}





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

// handlePostback processes button clicks (postback payloads)
func handlePostback(userID, payload string) {
	state := GetUserState(userID)
	
	switch payload {
	// Language selection
	case "LANG_EN":
		state.Language = "en"
		state.State = "greeting"
		SendMessage(userID, "✅ English selected!")
		startOrderingFlow(userID)
		
	case "LANG_MY":
		state.Language = "my"
		state.State = "greeting"
		SendMessage(userID, "✅ မြန်မာဘာသာ ရွေးချယ်ပြီးပါပြီ!")
		startOrderingFlow(userID)
		
	// Persistent Menu Actions (from ☰ menu)
	case "MENU_ORDER":
		startOrderingFlow(userID)
	
	case "MENU_ORDER_HISTORY":
		showOrderHistory(userID)
		
	case "MENU_ABOUT":
		showAbout(userID)  // Shows both About and Help combined
		
	case "MENU_CHANGE_LANG":
		showLanguageSelection(userID)
	
	// Main Menu Actions (from card buttons)
	case "MENU_ORDER_PRODUCTS":
		showProducts(userID)
	
	case "MENU_HELP":
		showHelp(userID)
	
	case "GET_STARTED":
		showLanguageSelection(userID)
		
	// Product selection
	case "ORDER_CHOCOLATE_CAKE":
		// Ignore if user is not in product selection state
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Chocolate Cake"].Name
		state.CurrentEmoji = ProductCatalog["Chocolate Cake"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
		
	case "ORDER_VANILLA_CAKE":
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Vanilla Cake"].Name
		state.CurrentEmoji = ProductCatalog["Vanilla Cake"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
		
	case "ORDER_RED_VELVET":
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Red Velvet Cake"].Name
		state.CurrentEmoji = ProductCatalog["Red Velvet Cake"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
		
	case "ORDER_CROISSANT":
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Croissant"].Name
		state.CurrentEmoji = ProductCatalog["Croissant"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
		
	case "ORDER_CINNAMON_ROLL":
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Cinnamon Roll"].Name
		state.CurrentEmoji = ProductCatalog["Cinnamon Roll"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
		
	case "ORDER_CUPCAKE":
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Chocolate Cupcake"].Name
		state.CurrentEmoji = ProductCatalog["Chocolate Cupcake"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
		
	case "ORDER_COFFEE":
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Coffee"].Name
		state.CurrentEmoji = ProductCatalog["Coffee"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
		
	case "ORDER_BREAD":
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Bread"].Name
		state.CurrentEmoji = ProductCatalog["Bread"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
	
	case "ORDER_CHOCOLATE_CUPCAKE":
		if state.State != "awaiting_product" {
			SendMessage(userID, "⚠️ Please complete your current step first, or type 'cancel' to start over.")
			return
		}
		state.CurrentProduct = ProductCatalog["Chocolate Cupcake"].Name
		state.CurrentEmoji = ProductCatalog["Chocolate Cupcake"].Emoji
		state.State = "awaiting_quantity"
		SendTypingIndicator(userID, true)
		askQuantity(userID)
		
	// Quantity selection
	case "QTY_1":
		if state.State != "awaiting_quantity" {
			SendMessage(userID, "⚠️ Please select a product first!")
			return
		}
		state.CurrentQuantity = 1
		SendTypingIndicator(userID, true)
		addToCart(userID)
		
	case "QTY_2":
		if state.State != "awaiting_quantity" {
			SendMessage(userID, "⚠️ Please select a product first!")
			return
		}
		state.CurrentQuantity = 2
		SendTypingIndicator(userID, true)
		addToCart(userID)
		
	case "QTY_3":
		if state.State != "awaiting_quantity" {
			SendMessage(userID, "⚠️ Please select a product first!")
			return
		}
		state.CurrentQuantity = 3
		SendTypingIndicator(userID, true)
		addToCart(userID)
		
	case "QTY_4":
		if state.State != "awaiting_quantity" {
			SendMessage(userID, "⚠️ Please select a product first!")
			return
		}
		state.CurrentQuantity = 4
		SendTypingIndicator(userID, true)
		addToCart(userID)
		
	case "QTY_5":
		if state.State != "awaiting_quantity" {
			SendMessage(userID, "⚠️ Please select a product first!")
			return
		}
		state.CurrentQuantity = 5
		SendTypingIndicator(userID, true)
		addToCart(userID)
		
	// Cart actions
	case "ADD_MORE_ITEMS":
		showProducts(userID)
		
	case "CHECKOUT":
		// Show cart and ask for name
		showCart(userID)
		SendTypingIndicator(userID, true)
		askName(userID)
		
	// Navigation
	case "GO_BACK":
		goBack(userID)
		
	case "MAIN_MENU":
		ResetUserState(userID)
		startOrderingFlow(userID)
		
	// Delivery type
	case "PICKUP":
		state.DeliveryType = "pickup"
		state.Address = "Pickup at store"
		state.State = "confirming"
		SendTypingIndicator(userID, true)
		showOrderSummary(userID)
		
	case "DELIVERY":
		state.DeliveryType = "delivery"
		state.State = "awaiting_address"
		
		// Add navigation options when asking for address
		quickReplies := []QuickReply{
			{ContentType: "text", Title: "⬅️ Back", Payload: "GO_BACK"},
			{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
		}
		SendQuickReplies(userID, "Perfect! Please type your delivery address:\n(Street, City, ZIP)", quickReplies)
		
	// Order confirmation
	case "CONFIRM_ORDER":
		SendTypingIndicator(userID, true)
		confirmOrder(userID)
		
	case "CANCEL_ORDER":
		ResetUserState(userID)
		SendMessage(userID, "❌ Order cancelled.")
		SendMessage(userID, "━━━━━━━━━━━━━━━━━")
		SendMessage(userID, "Ready to start fresh? Type 'menu' to see our products!")
		
	// Special actions
	case "SHOW_MENU":
		showMenu(userID)
	
	// Rating actions
	case "RATING_1":
		handleRating(userID, 1)
	case "RATING_2":
		handleRating(userID, 2)
	case "RATING_3":
		handleRating(userID, 3)
	case "RATING_4":
		handleRating(userID, 4)
	case "RATING_5":
		handleRating(userID, 5)
	case "SKIP_RATING":
		SendMessage(userID, "No problem! Feel free to rate us anytime.\n\nType 'menu' to order again! 🍰")
		ResetUserState(userID)
		
	default:
		// Check for dynamic payloads (REORDER_123, RATE_ORDER_123)
		if strings.HasPrefix(payload, "REORDER_") {
			orderIDStr := strings.TrimPrefix(payload, "REORDER_")
			if orderID, err := strconv.Atoi(orderIDStr); err == nil {
				// Check business hours before reordering
				if !checkBusinessHours(userID) {
					return
				}
				handleReorder(userID, orderID)
				return
			}
		}
		
		if strings.HasPrefix(payload, "RATE_ORDER_") {
			orderIDStr := strings.TrimPrefix(payload, "RATE_ORDER_")
			if orderID, err := strconv.Atoi(orderIDStr); err == nil {
				askForRating(userID, orderID)
				return
			}
		}
		
		SendMessage(userID, "Sorry, I didn't understand that. Let's start over!")
		ResetUserState(userID)
	}
}

// showOrderSummary displays the order summary and asks for confirmation
func showOrderSummary(userID string) {
	state := GetUserState(userID)
	
	deliveryIcon := "🏠"
	if state.DeliveryType == "delivery" {
		deliveryIcon = "🚚"
	}
	
	// Build cart items display with pricing
	cartDisplay := ""
	totalItems := 0
	for _, item := range state.Cart {
		itemPrice := 0.00
		if product, exists := ProductCatalog[item.Product]; exists {
			priceStr := strings.ReplaceAll(product.Price, "$", "")
			if price, err := strconv.ParseFloat(priceStr, 64); err == nil {
				itemPrice = price * float64(item.Quantity)
			}
		}
		cartDisplay += fmt.Sprintf("• %d× %s %s - $%.2f\n", item.Quantity, item.ProductEmoji, item.Product, itemPrice)
		totalItems += item.Quantity
	}
	
	// Calculate totals
	subtotal, deliveryFee, totalAmount := calculateOrderTotals(state.Cart, state.DeliveryType, state.Address)
	
	// Pricing breakdown
	pricingInfo := fmt.Sprintf(
		"\n💰 **Pricing:**\n"+
			"Subtotal: $%.2f\n"+
			"Delivery Fee: $%.2f\n"+
			"━━━━━━━━━━━━\n"+
			"**Total: $%.2f**",
		subtotal,
		deliveryFee,
		totalAmount,
	)
	
	summary := fmt.Sprintf(
		"📋 **Order Summary**\n\n"+
			"🛒 **Your Items:**\n"+
			"%s"+
			"%s\n\n"+
			"👤 **Customer:** %s\n"+
			"%s **%s**\n"+
			"📍 **Address:** %s\n\n"+
			"Everything look good?",
		cartDisplay,
		pricingInfo,
		state.CustomerName,
		deliveryIcon, strings.Title(state.DeliveryType),
		state.Address,
	)
	
	quickReplies := []QuickReply{
		{ContentType: "text", Title: "✅ Confirm Order", Payload: "CONFIRM_ORDER"},
		{ContentType: "text", Title: "❌ Cancel", Payload: "CANCEL_ORDER"},
	}
	SendQuickReplies(userID, summary, quickReplies)
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
		SenderID:     userID, // Persist sender ID so we can send status updates later
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

// showMenu displays the product menu
func showMenu(userID string) {
	// Show text menu first
	menu := "🍰 **BakeFlow Menu**\n\n" +
		"🎂 **Cakes**\n" +
		"  • Chocolate Cake - $25\n" +
		"  • Vanilla Cake - $24\n" +
		"  • Red Velvet Cake - $28\n\n" +
		"🥐 **Pastries**\n" +
		"  • Croissant - $4.50\n" +
		"  • Cinnamon Roll - $5\n\n" +
		"🧁 **Others**\n" +
		"  • Chocolate Cupcake - $3.50\n" +
		"  • Fresh Bread - $6\n" +
		"  • Coffee - $5\n\n" +
		"👇 Click the buttons below to order!"
	
	SendMessage(userID, menu)
	
	// Automatically show product cards with order buttons
	showProducts(userID)
}

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
			if i < 3 { // Show max 3 items
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
		Comment: "", // Could add comment feature later
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

