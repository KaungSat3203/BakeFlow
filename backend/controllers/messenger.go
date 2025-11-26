package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

// SendMessage sends a text message to a user via Messenger API
func SendMessage(recipientID, messageText string) error {
	pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")
	if pageAccessToken == "" {
		return fmt.Errorf("PAGE_ACCESS_TOKEN not set in .env")
	}

	// Construct the message payload
	payload := map[string]interface{}{
		"recipient": map[string]string{"id": recipientID},
		"message":   map[string]string{"text": messageText},
	}

	payloadBytes, _ := json.Marshal(payload)

	// Send to Facebook Graph API
	url := fmt.Sprintf("https://graph.facebook.com/v18.0/me/messages?access_token=%s", pageAccessToken)

	resp, err := http.Post(url, "application/json", strings.NewReader(string(payloadBytes)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("❌ Error sending message: %s", string(body))
		return fmt.Errorf("failed to send message: %s", string(body))
	}

	log.Printf("✅ Message sent to %s", recipientID)
	return nil
}

// SendQuickReplies sends a message with quick reply buttons
func SendQuickReplies(recipientID, messageText string, quickReplies []QuickReply) error {
	pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")
	if pageAccessToken == "" {
		return fmt.Errorf("PAGE_ACCESS_TOKEN not set in .env")
	}

	payload := map[string]interface{}{
		"recipient": map[string]string{"id": recipientID},
		"message": map[string]interface{}{
			"text":          messageText,
			"quick_replies": quickReplies,
		},
	}

	payloadBytes, _ := json.Marshal(payload)
	url := fmt.Sprintf("https://graph.facebook.com/v18.0/me/messages?access_token=%s", pageAccessToken)

	resp, err := http.Post(url, "application/json", strings.NewReader(string(payloadBytes)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("❌ Error sending quick replies: %s", string(body))
		return fmt.Errorf("failed to send quick replies: %s", string(body))
	}

	log.Printf("✅ Quick replies sent to %s", recipientID)
	return nil
}

// SendTypingIndicator shows typing indicator for better UX
func SendTypingIndicator(recipientID string, on bool) error {
	pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")
	if pageAccessToken == "" {
		return fmt.Errorf("PAGE_ACCESS_TOKEN not set in .env")
	}

	action := "typing_off"
	if on {
		action = "typing_on"
	}

	payload := map[string]interface{}{
		"recipient":     map[string]string{"id": recipientID},
		"sender_action": action,
	}

	payloadBytes, _ := json.Marshal(payload)
	url := fmt.Sprintf("https://graph.facebook.com/v18.0/me/messages?access_token=%s", pageAccessToken)

	http.Post(url, "application/json", strings.NewReader(string(payloadBytes)))
	return nil
}

// SendGenericTemplate sends image-based product cards (carousel)
func SendGenericTemplate(recipientID string, elements []Element) error {
	pageAccessToken := os.Getenv("PAGE_ACCESS_TOKEN")
	if pageAccessToken == "" {
		return fmt.Errorf("PAGE_ACCESS_TOKEN not set in .env")
	}

	payload := map[string]interface{}{
		"recipient": map[string]string{"id": recipientID},
		"message": map[string]interface{}{
			"attachment": map[string]interface{}{
				"type": "template",
				"payload": GenericTemplate{
					TemplateType: "generic",
					Elements:     elements,
				},
			},
		},
	}

	payloadBytes, _ := json.Marshal(payload)
	url := fmt.Sprintf("https://graph.facebook.com/v18.0/me/messages?access_token=%s", pageAccessToken)

	resp, err := http.Post(url, "application/json", strings.NewReader(string(payloadBytes)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("❌ Error sending generic template: %s", string(body))
		return fmt.Errorf("failed to send generic template: %s", string(body))
	}

	log.Printf("✅ Generic template sent to %s", recipientID)
	return nil
}
