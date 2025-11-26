# 🚀 Quick Start Guide - 5 Minutes to Live Bot

Follow these exact steps to get your Facebook Messenger bot running:

## Step 1: Start the Server (30 seconds)

```bash
cd backend
go run main.go
```

✅ You should see:
```
✅ .env file loaded successfully
✅ VERIFY_TOKEN: veri****e123
🚀 Server starting on port 8080...
```

## Step 2: Start ngrok (30 seconds)

**Open a NEW terminal window** and run:

```bash
ngrok http 8080
```

✅ Copy the HTTPS URL that looks like:
```
https://abcd-1234-5678.ngrok-free.app
```

## Step 3: Configure Facebook Webhook (2 minutes)

1. Go to: https://developers.facebook.com/apps
2. Click your app → **Messenger** → **Settings**
3. Scroll to **Webhooks** section
4. Click **Add Callback URL**

**Enter these values:**

| Field | Value |
|-------|-------|
| Callback URL | `https://YOUR-NGROK-URL.ngrok-free.app/webhook` |
| Verify Token | `verifyme123` (from your .env file) |

5. Click **Verify and Save**

✅ You should see a green checkmark!

## Step 4: Subscribe to Events (30 seconds)

Still in the Webhooks section:

1. Find your Facebook Page in the list
2. Click **Subscribe**
3. Check these boxes:
   - ✅ `messages`
   - ✅ `messaging_postbacks`
4. Click **Subscribe**

## Step 5: Test It! (1 minute)

1. Go to your Facebook Page
2. Send a message: "Hello!"
3. Check your server logs

✅ You should see:
```
========== INCOMING WEBHOOK POST ==========
📨 Message from 1234567890: Hello!
```

---

## 🎉 Success! What's Next?

### Add Auto-Reply

Edit `backend/controllers/webhook.go`, find this section:

```go
if event.Message.Text != "" {
    log.Printf("📨 Message from %s: %s", senderID, event.Message.Text)
    
    // TODO: Add your message processing logic here
```

Add this below the TODO:

```go
// Send automatic reply
SendMessage(senderID, "Thanks for your message! How can I help you today?")
```

Restart the server and test again!

### Add Command Handler

Replace the TODO section with:

```go
messageText := strings.ToLower(event.Message.Text)

switch messageText {
case "menu":
    SendMessage(senderID, "🍰 Our Menu:\n1. Chocolate Cake - $25\n2. Vanilla Cupcake - $3\n3. Red Velvet - $28")
    
case "help":
    SendMessage(senderID, "Commands:\n• menu - View our menu\n• order - Place an order\n• help - Show this message")
    
case "order":
    SendMessage(senderID, "Great! What would you like to order? Reply with the item name.")
    
default:
    SendMessage(senderID, "I received: " + event.Message.Text + "\n\nType 'help' for available commands.")
}
```

---

## ⚠️ Troubleshooting

### "URL couldn't be validated"

```bash
# Test locally first
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"

# Should return: test
```

If this works but ngrok doesn't:
1. Make sure ngrok is running
2. Copy the CORRECT ngrok URL (changes every restart)
3. Use the HTTPS URL, not HTTP

### "Token mismatch"

Your `.env` file token must EXACTLY match Meta console:

```bash
# Check your token
cat backend/.env | grep VERIFY_TOKEN

# Output should be: VERIFY_TOKEN=verifyme123
```

Use this EXACT value in Meta Developer Console.

### Not receiving messages

1. Check if page is subscribed:
   - Meta Developer Console → Messenger → Settings → Webhooks
   - Your page should be listed with "Subscribed" status
   
2. Check if app is in Dev Mode:
   - Only admins/testers can message during development
   - Add yourself as a tester: App Settings → Roles → Testers

---

## 📁 Project Structure Reference

```
backend/
├── main.go                    # ✅ Entry point
├── .env                       # ✅ Your tokens (NEVER commit!)
├── controllers/
│   └── webhook.go            # ✅ Message handling logic
├── routes/
│   └── routes.go             # ✅ HTTP routing
└── README.md                  # Full documentation
```

---

## 🔗 Useful Commands

```bash
# Start server
cd backend && go run main.go

# Start ngrok (new terminal)
ngrok http 8080

# Test webhook locally
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"

# Kill server on port 8080
lsof -ti:8080 | xargs kill -9

# Check server logs
tail -f /tmp/bakeflow.log

# Run test script
./test_webhook.sh
```

---

## 📚 Next Steps

1. ✅ **Basic bot working?** Add more commands in `webhook.go`
2. ✅ **Want buttons?** See `examples.go` for Quick Reply buttons
3. ✅ **Database integration?** Orders are already set up in `controllers/orders.go`
4. ✅ **Deploy to production?** See deployment guide (coming soon)

---

## 🆘 Still Stuck?

1. Read: `TROUBLESHOOTING.md` (comprehensive guide)
2. Check: `README.md` (full documentation)
3. Run: `./test_webhook.sh` (automated testing)

---

**Total Time: ~5 minutes** ⏱️

**Difficulty: Easy** 🟢

**Last Updated: 2025-11-17**
