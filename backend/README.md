# BakeFlow - Facebook Messenger Chatbot Backend

A minimal Go backend for Facebook Messenger chatbot with webhook verification and message handling.

## 📁 Project Structure

```
backend/
├── main.go                 # Entry point with .env loading and setup instructions
├── go.mod                  # Go dependencies
├── .env                    # Environment variables (DO NOT COMMIT)
├── controllers/
│   ├── webhook.go         # Facebook webhook verification and message handling
│   └── orders.go          # Order management endpoints
├── routes/
│   └── routes.go          # HTTP router with logging middleware
├── configs/
│   └── config.go          # Database configuration
└── models/
    └── order.go           # Order data models
```

## 🚀 Quick Start

### 1. Install Dependencies

```bash
cd backend
go mod download
```

### 2. Configure Environment Variables

Make sure your `.env` file contains:

```env
# Facebook Messenger Config
VERIFY_TOKEN=verifyme123
PAGE_ACCESS_TOKEN=EAARCf5z54DUBP3sHZCcD2oZBZBrZBFcz68Dry...

# Database
DATABASE_URL=postgresql://user:pass@host/db?sslmode=require
```

**IMPORTANT:** The `VERIFY_TOKEN` must match exactly what you enter in Meta Developer Console.

### 3. Run the Server

```bash
go run main.go
```

You should see:
```
✅ .env file loaded successfully
✅ VERIFY_TOKEN: veri****e123 (length: 11)
✅ PAGE_ACCESS_TOKEN: EAARCf5z54DUBP3sHZCcD... (length: 195)
🚀 Server starting on port 8080...
```

### 4. Expose with ngrok

In a **new terminal window**:

```bash
ngrok http 8080
```

Copy the HTTPS URL (e.g., `https://abcd-123-45-67-89.ngrok-free.app`)

### 5. Configure Facebook Webhook

1. Go to [Meta Developer Console](https://developers.facebook.com/apps)
2. Select your app → **Messenger** → **Settings**
3. Under **Webhooks**, click **Add Callback URL**

   **Callback URL:**
   ```
   https://YOUR-NGROK-URL.ngrok-free.app/webhook
   ```

   **Verify Token:**
   ```
   verifyme123
   ```
   (or whatever you set in `.env`)

4. Click **Verify and Save**

5. Subscribe to webhook events:
   - ✅ `messages`
   - ✅ `messaging_postbacks`

6. Subscribe your page to the webhook

## 🧪 Testing

### Test Webhook Verification Locally

```bash
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test123"
```

**Expected output:** `test123`

### Test with ngrok URL

```bash
curl "https://your-ngrok-url.ngrok-free.app/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test123"
```

### Send a Test Message

After webhook is verified, message your Facebook Page. Check server logs:

```
========== INCOMING WEBHOOK POST ==========
📨 Message from 1234567890: Hello!
✅ Message sent to 1234567890
```

## ❌ Troubleshooting Common Issues

### 1. "Callback URL or verify token couldn't be validated"

**Causes & Solutions:**

| Problem | Solution |
|---------|----------|
| ❌ Token mismatch | Ensure `VERIFY_TOKEN` in `.env` **exactly matches** Meta console |
| ❌ .env not loaded | Check for `✅ .env file loaded successfully` in logs |
| ❌ Server not running | Run `go run main.go` and look for `🚀 Server starting...` |
| ❌ Wrong ngrok URL | Get fresh ngrok URL with `ngrok http 8080` |
| ❌ Not using HTTPS | ngrok automatically provides HTTPS ✅ |

**Debug Steps:**

1. Check server logs when clicking "Verify and Save" in Meta console
2. Look for this in logs:
   ```
   ========== WEBHOOK VERIFICATION ATTEMPT ==========
   Mode: subscribe
   Token received: verifyme123
   Token expected: verifyme123
   ✅ Webhook verified successfully!
   ```

3. If tokens don't match:
   ```
   ❌ Webhook verification FAILED
      - Token mismatch!
      - Received: 'wrongtoken'
      - Expected: 'verifyme123'
   ```

### 2. "The URL couldn't be validated"

**Causes:**

- ❌ Server crashed or not responding
- ❌ ngrok tunnel closed
- ❌ Firewall blocking port 8080
- ❌ Server returning wrong HTTP status

**Solution:**

```bash
# Terminal 1: Start server
cd backend
go run main.go

# Terminal 2: Start ngrok
ngrok http 8080

# Terminal 3: Test manually
curl "https://YOUR-NGROK-URL/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"
```

### 3. Webhook verified but not receiving messages

**Causes:**

- ❌ Page not subscribed to webhook
- ❌ Webhook fields not selected (`messages`, `messaging_postbacks`)
- ❌ App not in production mode (only admins can message during dev)

**Solution:**

1. In Meta Developer Console → Messenger → Settings
2. Under **Webhooks**, find your page
3. Click **Edit** → ensure these fields are checked:
   - ✅ `messages`
   - ✅ `messaging_postbacks`
4. Add yourself as a test user if app is in dev mode

### 4. "Error loading .env file"

**Causes:**

- ❌ `.env` file not in `backend/` directory
- ❌ Wrong working directory

**Solution:**

```bash
# Check current directory
pwd
# Should output: /path/to/BakeFlow/backend

# Check .env exists
ls -la .env

# Run from correct directory
cd backend
go run main.go
```

### 5. Environment Variables Not Set

If you see:
```
❌ VERIFY_TOKEN: NOT SET - Webhook verification will fail!
```

**Solutions:**

1. Create `.env` file in `backend/` directory:
   ```bash
   touch .env
   ```

2. Add variables:
   ```env
   VERIFY_TOKEN=verifyme123
   PAGE_ACCESS_TOKEN=your_token_here
   DATABASE_URL=your_db_url_here
   ```

3. Restart server:
   ```bash
   go run main.go
   ```

## 📝 Key Files Explained

### `controllers/webhook.go`

- **`VerifyWebhook()`**: Handles GET requests for Facebook verification
  - Checks `hub.mode`, `hub.verify_token`, `hub.challenge`
  - Returns challenge if token matches
  - Includes detailed debug logging

- **`ReceiveWebhook()`**: Handles POST requests with user messages
  - Parses JSON payload
  - Extracts sender ID and message text
  - Returns `EVENT_RECEIVED` to Facebook

- **`SendMessage()`**: Sends replies via Graph API
  - Uses `PAGE_ACCESS_TOKEN`
  - Posts to `https://graph.facebook.com/v18.0/me/messages`

### `routes/routes.go`

- **`SetupRoutes()`**: Configures HTTP endpoints
  - `/` - Health check
  - `/webhook` - GET (verify) and POST (messages)
  - `/orders` - Orders API

- **`LoggingMiddleware`**: Logs all requests (useful for debugging)

- **`CORSMiddleware`**: Adds CORS headers for testing

### `main.go`

- Loads `.env` file with `godotenv`
- Verifies environment variables are set
- Prints setup instructions
- Starts HTTP server on port 8080

## 🔐 Security Notes

1. **Never commit `.env` file** - Add to `.gitignore`
2. **Rotate tokens regularly** - Get new `PAGE_ACCESS_TOKEN` from Meta
3. **Use HTTPS in production** - ngrok provides this automatically
4. **Validate webhook signatures** - (TODO: implement for production)

## 🛠️ Development Workflow

```bash
# Terminal 1: Run server with auto-reload
go run main.go

# Terminal 2: Tunnel with ngrok
ngrok http 8080

# Terminal 3: Watch logs
tail -f logs.txt  # (if you add file logging)

# Or use air for hot reload:
go install github.com/air-verse/air@latest
air
```

## 📚 Facebook Messenger Resources

- [Messenger Platform Docs](https://developers.facebook.com/docs/messenger-platform)
- [Webhook Reference](https://developers.facebook.com/docs/messenger-platform/webhooks)
- [Send API Reference](https://developers.facebook.com/docs/messenger-platform/reference/send-api)
- [Testing & Troubleshooting](https://developers.facebook.com/docs/messenger-platform/testing)

## 🐛 Debug Mode

To see all incoming requests, check server logs:

```
➡️  GET /webhook from 127.0.0.1:12345
========== WEBHOOK VERIFICATION ATTEMPT ==========
Mode: subscribe
Token received: verifyme123
Token expected: verifyme123
Challenge: test123
✅ Webhook verified successfully!
⬅️  Completed in 1.234ms
```

## 📞 Support

If you're still having issues:

1. Check Meta Developer Console → Webhooks → "Test" button
2. Review server logs for error messages
3. Verify ngrok tunnel is active: `curl https://your-ngrok-url.ngrok-free.app/`
4. Test locally first before using ngrok

## ✅ Success Checklist

- [ ] `.env` file created with correct tokens
- [ ] Server starts without errors
- [ ] ngrok tunnel active and HTTPS URL copied
- [ ] Webhook verified in Meta Developer Console (green checkmark)
- [ ] Webhook events subscribed (`messages`, `messaging_postbacks`)
- [ ] Page subscribed to webhook
- [ ] Test message sent to page and received in server logs

---

**Last Updated:** 2025-11-17
