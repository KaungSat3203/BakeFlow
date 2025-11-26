# Facebook Messenger Webhook Flow

## 📊 Architecture Diagram

```
┌──────────────┐         ┌──────────┐         ┌───────────────┐         ┌──────────┐
│   Facebook   │         │  ngrok   │         │   Go Server   │         │ Database │
│  Messenger   │         │ Tunnel   │         │  (Port 8080)  │         │ (Neon)   │
└──────┬───────┘         └────┬─────┘         └───────┬───────┘         └────┬─────┘
       │                      │                        │                      │
       │  1. Verify Webhook   │                        │                      │
       │ GET /webhook?        │                        │                      │
       │ hub.verify_token=... │                        │                      │
       ├─────────────────────>│                        │                      │
       │                      │  2. Forward Request    │                      │
       │                      ├───────────────────────>│                      │
       │                      │                        │  3. Load .env        │
       │                      │                        │     VERIFY_TOKEN     │
       │                      │                        │                      │
       │                      │                        │  4. Compare Tokens   │
       │                      │                        │     ✅ Match?        │
       │                      │                        │                      │
       │                      │  5. Return Challenge   │                      │
       │                      │<───────────────────────┤                      │
       │  6. Challenge String │                        │                      │
       │<─────────────────────┤                        │                      │
       │                      │                        │                      │
       │  ✅ Verified!        │                        │                      │
       │                      │                        │                      │
       │                      │                        │                      │
       │  7. User Sends Msg   │                        │                      │
       │     "Hello!"         │                        │                      │
       │                      │                        │                      │
       │  8. POST /webhook    │                        │                      │
       │     {message: ...}   │                        │                      │
       ├─────────────────────>│                        │                      │
       │                      │  9. Forward POST       │                      │
       │                      ├───────────────────────>│                      │
       │                      │                        │ 10. Parse JSON       │
       │                      │                        │     Extract text     │
       │                      │                        │                      │
       │                      │                        │ 11. Process Message  │
       │                      │                        │     (Your logic)     │
       │                      │                        │                      │
       │                      │                        │ 12. Save to DB       │
       │                      │                        ├─────────────────────>│
       │                      │                        │                      │
       │                      │ 13. Return OK          │                      │
       │                      │<───────────────────────┤                      │
       │  14. 200 OK          │                        │                      │
       │<─────────────────────┤                        │                      │
       │                      │                        │                      │
       │                      │                        │ 15. Send Reply       │
       │                      │                        │     Graph API        │
       │  16. Reply Message   │                        │                      │
       │<──────────────────────────────────────────────┤                      │
       │                      │                        │                      │
```

## 🔄 Request Flow Details

### Phase 1: Webhook Verification (One-time Setup)

```
Facebook → GET Request → ngrok → Your Server
                                     │
                                     ├─ Load VERIFY_TOKEN from .env
                                     ├─ Check hub.verify_token == VERIFY_TOKEN
                                     ├─ ✅ Match? Return hub.challenge
                                     └─ ❌ No match? Return 403 Forbidden
```

### Phase 2: Message Reception (Every message)

```
User sends message → Facebook → POST Request → ngrok → Your Server
                                                           │
                                                           ├─ Parse JSON payload
                                                           ├─ Extract sender ID
                                                           ├─ Extract message text
                                                           ├─ Process message
                                                           ├─ Save to database (optional)
                                                           ├─ Send reply (optional)
                                                           └─ Return 200 OK
```

## 📂 Code Flow

### File: `main.go`
```
main()
  │
  ├─ Load .env file (godotenv.Load())
  ├─ Verify environment variables exist
  ├─ Print setup instructions
  ├─ Connect to database
  ├─ Setup HTTP routes
  └─ Start server on port 8080
```

### File: `routes/routes.go`
```
SetupRoutes()
  │
  ├─ Create HTTP mux
  ├─ Add middleware:
  │   ├─ LoggingMiddleware (logs all requests)
  │   └─ CORSMiddleware (adds CORS headers)
  │
  ├─ Register routes:
  │   ├─ GET /webhook → VerifyWebhook()
  │   └─ POST /webhook → ReceiveWebhook()
  │
  └─ Return configured handler
```

### File: `controllers/webhook.go`
```
VerifyWebhook(w, r)
  │
  ├─ Get query parameters:
  │   ├─ hub.mode
  │   ├─ hub.verify_token
  │   └─ hub.challenge
  │
  ├─ Get VERIFY_TOKEN from env
  ├─ Log all values (debug)
  │
  ├─ Check: mode == "subscribe" && token == verifyToken
  │   ├─ ✅ Yes: Return challenge
  │   └─ ❌ No: Return 403 + log detailed error
  │
  └─ Done

ReceiveWebhook(w, r)
  │
  ├─ Read request body
  ├─ Parse JSON → WebhookPayload
  │
  ├─ For each entry:
  │   └─ For each messaging event:
  │       ├─ Extract sender.id
  │       ├─ Extract message.text
  │       ├─ Log message
  │       └─ TODO: Your processing logic
  │
  └─ Return "EVENT_RECEIVED"

SendMessage(recipientID, text)
  │
  ├─ Get PAGE_ACCESS_TOKEN from env
  ├─ Build JSON payload
  ├─ POST to Graph API:
  │   URL: https://graph.facebook.com/v18.0/me/messages
  │
  └─ Return success/error
```

## 🧪 Testing Flow

### Local Testing
```
Terminal 1: Start Server
  $ cd backend
  $ go run main.go
  → Server listening on :8080

Terminal 2: Test Webhook
  $ curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"
  → test ✅
```

### ngrok Testing
```
Terminal 1: Server (already running)
Terminal 2: Start ngrok
  $ ngrok http 8080
  → Forwarding: https://abc123.ngrok-free.app → localhost:8080

Terminal 3: Test ngrok URL
  $ curl "https://abc123.ngrok-free.app/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"
  → test ✅
```

### Facebook Testing
```
1. Configure webhook in Meta Developer Console:
   URL: https://abc123.ngrok-free.app/webhook
   Token: verifyme123

2. Click "Verify and Save"
   → Facebook sends GET request
   → Server returns challenge
   → ✅ Webhook verified!

3. Send test message to page:
   → Facebook sends POST request
   → Server logs: "📨 Message from [user]: Hello!"
   → Server can send reply
```

## 🔍 Debug Flow

### When Verification Fails

```
1. Check server logs:
   ========== WEBHOOK VERIFICATION ATTEMPT ==========
   Mode: subscribe
   Token received: wrongtoken
   Token expected: verifyme123
   ❌ Webhook verification FAILED
      - Token mismatch!

2. Fix token in Meta console:
   Use: verifyme123

3. Retry verification:
   ✅ Webhook verified successfully!
```

## 🎯 Data Structures

### Incoming Webhook (POST)
```json
{
  "object": "page",
  "entry": [
    {
      "id": "page_id",
      "time": 1234567890,
      "messaging": [
        {
          "sender": {"id": "user_id"},
          "recipient": {"id": "page_id"},
          "timestamp": 1234567890,
          "message": {
            "mid": "message_id",
            "text": "Hello!"
          }
        }
      ]
    }
  ]
}
```

### Outgoing Message (SendMessage)
```json
{
  "recipient": {"id": "user_id"},
  "message": {"text": "Your reply here"}
}
```

## 🔐 Environment Variables Flow

```
.env file
  │
  ├─ VERIFY_TOKEN=verifyme123
  ├─ PAGE_ACCESS_TOKEN=EAA...
  └─ DATABASE_URL=postgresql://...
      │
      ├─ Loaded by godotenv.Load()
      │
      ├─ Accessed by:
      │   ├─ os.Getenv("VERIFY_TOKEN")    → Webhook verification
      │   ├─ os.Getenv("PAGE_ACCESS_TOKEN") → Send messages
      │   └─ os.Getenv("DATABASE_URL")     → Database connection
      │
      └─ Validated at startup (main.go)
```

## ⚡ Performance Flow

```
Request received
  │
  ├─ LoggingMiddleware: Log start time
  │
  ├─ Route to handler (VerifyWebhook or ReceiveWebhook)
  │   └─ Process request (<1ms typical)
  │
  ├─ LoggingMiddleware: Calculate duration
  │
  └─ Return response
      └─ Log: "⬅️ Completed in 126.042µs"
```

## 🚦 Error Handling Flow

```
Error Occurs
  │
  ├─ Check error type:
  │   ├─ .env not found
  │   │   └─ ⚠️ Warning (continue with system env)
  │   │
  │   ├─ VERIFY_TOKEN empty
  │   │   └─ ❌ Return 500 Internal Server Error
  │   │
  │   ├─ Token mismatch
  │   │   └─ ❌ Return 403 Forbidden + detailed log
  │   │
  │   ├─ JSON parse error
  │   │   └─ ⚠️ Log error but return 200 OK
  │   │
  │   └─ Send message error
  │       └─ ❌ Log error + return error to caller
  │
  └─ All errors logged with ❌ emoji for easy filtering
```

---

**This diagram shows the complete flow from Facebook through ngrok to your Go server and back.**
