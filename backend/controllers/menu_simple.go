package controllers

// showMainMenuSimple displays main menu as one simple box with 3 buttons (no images)
func showMainMenuSimple(userID string) {
	state := GetUserState(userID)
	
	// Create one card with 3 buttons (no image, just clean text)
	var element Element
	
	if state.Language == "my" {
		element = Element{
			Title:    "ဘာလုပ်ချင်လဲ?",
			Subtitle: "အောက်ပါရွေးချယ်စရာများမှ ရွေးချယ်ပါ",
			Buttons: []Button{
				{
					Type:    "postback",
					Title:   "🛒 အော်ဒါစတင်မယ်",
					Payload: "MENU_ORDER_PRODUCTS",
				},
				{
					Type:    "postback",
					Title:   "ℹ️ အကြောင်းအရာ",
					Payload: "MENU_ABOUT",
				},
				{
					Type:    "postback",
					Title:   "❓ အကူအညီ",
					Payload: "MENU_HELP",
				},
			},
		}
	} else {
		element = Element{
			Title:    "What would you like to do?",
			Subtitle: "Choose an option below",
			Buttons: []Button{
				{
					Type:    "postback",
					Title:   "🛒 Start Order",
					Payload: "MENU_ORDER_PRODUCTS",
				},
				{
					Type:    "postback",
					Title:   "ℹ️ About",
					Payload: "MENU_ABOUT",
				},
				{
					Type:    "postback",
					Title:   "❓ Help",
					Payload: "MENU_HELP",
				},
			},
		}
	}
	
	SendGenericTemplate(userID, []Element{element})
}
