# Facebook Webhook Verification Troubleshooting Guide

## 🔴 Error: "Callback URL or verify token couldn't be validated"

This is the most common error. Here's how to diagnose and fix it:

### Step 1: Check Your .env File

```bash
cd backend
cat .env
```

Verify that `VERIFY_TOKEN` is set:
```env
VERIFY_TOKEN=verifyme123
```

**CRITICAL:** The token in `.env` must EXACTLY match what you enter in Meta Developer Console (case-sensitive, no extra spaces).

### Step 2: Verify Server is Running

```bash
# Check if process is running
ps aux | grep "go run main.go"

# Or check if port 8080 is in use
lsof -i :8080
```

If nothing appears, start the server:
```bash
cd backend
go run main.go
```

You should see:
```
✅ .env file loaded successfully
✅ VERIFY_TOKEN: veri****e123 (length: 11)
🚀 Server starting on port 8080...
```

### Step 3: Test Locally First

Before using ngrok, test the webhook locally:

```bash
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test123"
```

**Expected output:** `test123`

**If you get "Forbidden":**
- The VERIFY_TOKEN doesn't match
- Check for typos, extra spaces, or wrong case
- Verify with: `echo $VERIFY_TOKEN` or `cat .env | grep VERIFY_TOKEN`

### Step 4: Check ngrok

```bash
# Start ngrok (if not already running)
ngrok http 8080
```

You should see:
```
Forwarding    https://abcd-1234.ngrok-free.app -> http://localhost:8080
```

**Test the ngrok URL:**
```bash
curl "https://YOUR-NGROK-URL.ngrok-free.app/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test123"
```

**Expected output:** `test123`

### Step 5: Common Mistakes in Meta Developer Console

#### ❌ Wrong Callback URL Format

**WRONG:**
```
http://your-url/webhook         # ❌ HTTP not allowed
https://your-url/webhook/       # ❌ Trailing slash
https://your-url                # ❌ Missing /webhook
localhost:8080/webhook          # ❌ Must use ngrok URL
```

**CORRECT:**
```
https://abcd-1234.ngrok-free.app/webhook
```

#### ❌ Token Mismatch

Make sure the token in Meta console EXACTLY matches your .env:

**.env file:**
```env
VERIFY_TOKEN=mySecretToken123
```

**Meta Developer Console:**
```
Verify Token: mySecretToken123
```

These must be IDENTICAL (case-sensitive).

### Step 6: Check Server Logs

When you click "Verify and Save" in Meta console, check your server logs:

```bash
tail -f /tmp/bakeflow.log
```

**Success looks like:**
```
➡️  GET /webhook from 34.67.89.123
========== WEBHOOK VERIFICATION ATTEMPT ==========
Mode: subscribe
Token received: verifyme123
Token expected: verifyme123
Challenge: AbCd1234EfGh5678
✅ Webhook verified successfully!
```

**Failure looks like:**
```
❌ Webhook verification FAILED
   - Token mismatch!
   - Received: 'wrongtoken'
   - Expected: 'verifyme123'
```

---

## 🔴 Error: "The URL couldn't be validated"

This means Meta can't reach your server at all.

### Checklist:

1. **Server Running?**
   ```bash
   curl http://localhost:8080/
   # Should return: "BakeFlow Bot is running! ✅"
   ```

2. **ngrok Running?**
   ```bash
   curl https://your-ngrok-url.ngrok-free.app/
   # Should return: "BakeFlow Bot is running! ✅"
   ```

3. **Firewall Blocking?**
   - Check if your firewall is blocking port 8080
   - On macOS: System Settings → Network → Firewall

4. **ngrok Expired?**
   - Free ngrok URLs change every time you restart ngrok
   - Get a new URL with `ngrok http 8080`
   - Update the URL in Meta Developer Console

---

## 🔴 Webhook Verified But Not Receiving Messages

### Issue: Webhook shows "Active" but no POST requests arrive

#### Solution 1: Subscribe Page to Webhook

1. Go to Meta Developer Console → Messenger → Settings
2. Scroll to **Webhooks** section
3. Find your page in the list
4. Click **Subscribe**
5. Make sure **messages** and **messaging_postbacks** are checked

#### Solution 2: Check App Mode

If your app is in **Development Mode**, only admins/testers can message the bot.

**To add test users:**
1. Go to App Settings → Roles
2. Add Test Users or Testers
3. Or switch app to **Live Mode** (requires app review for production)

#### Solution 3: Test with ngrok Inspector

1. Open ngrok web interface: http://localhost:4040
2. Send a message to your Facebook Page
3. Check if POST request appears in ngrok inspector
4. If it appears in ngrok but not in your logs, check your code

---

## 🔴 Error: "Error loading .env file"

### Cause: .env file not found or in wrong location

**Fix:**
```bash
# Make sure .env is in backend/ directory
cd backend
ls -la .env

# If .env doesn't exist, create it:
cat > .env << EOF
VERIFY_TOKEN=verifyme123
PAGE_ACCESS_TOKEN=your_token_here
DATABASE_URL=your_db_url_here
EOF

# Restart server
go run main.go
```

---

## 🔴 Port 8080 Already in Use

### Error: `listen tcp :8080: bind: address already in use`

**Find and kill the process:**
```bash
# Find process using port 8080
lsof -ti:8080

# Kill it
lsof -ti:8080 | xargs kill -9

# Or use a different port
PORT=3000 go run main.go
```

---

## 🔴 ngrok Connection Failed

### Error: `ngrok: command not found`

**Install ngrok:**
```bash
# macOS with Homebrew
brew install ngrok

# Or download from: https://ngrok.com/download
```

### Error: ngrok "ERR_NGROK_108"

**Sign up and authenticate:**
```bash
# Sign up at https://dashboard.ngrok.com/signup
# Get your auth token from https://dashboard.ngrok.com/get-started/your-authtoken

# Authenticate
ngrok authtoken YOUR_AUTH_TOKEN
```

---

## 🧪 Testing Checklist

Use this checklist to verify everything works:

- [ ] `.env` file exists in `backend/` directory
- [ ] `VERIFY_TOKEN` is set in `.env`
- [ ] Server starts without errors
- [ ] Local webhook test returns challenge: 
  ```bash
  curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"
  # Output: test
  ```
- [ ] ngrok is running and shows forwarding URL
- [ ] ngrok URL test returns challenge:
  ```bash
  curl "https://your-ngrok-url/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"
  # Output: test
  ```
- [ ] Webhook verified in Meta Developer Console (green checkmark)
- [ ] Page subscribed to webhook
- [ ] Webhook events selected: `messages`, `messaging_postbacks`
- [ ] Test message sent to page appears in server logs

---

## 🔍 Debug Mode

To see detailed request/response information:

### 1. Enable Verbose Logging

Add this to `main.go`:
```go
log.SetFlags(log.LstdFlags | log.Lshortfile)
```

### 2. Use ngrok Inspector

Open http://localhost:4040 to see all HTTP requests in real-time.

### 3. Test with Postman

Import this cURL as a Postman request:
```bash
curl -X GET "https://your-ngrok-url/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test123"
```

### 4. Check Meta Developer Tools

Go to: https://developers.facebook.com/tools/debug/

Enter your webhook URL and check for errors.

---

## 📞 Still Having Issues?

If you've tried everything above:

1. **Check server logs** for specific error messages
2. **Check ngrok web interface** (http://localhost:4040) for incoming requests
3. **Verify tokens match** by comparing side-by-side:
   ```bash
   echo "From .env: $(grep VERIFY_TOKEN .env | cut -d= -f2)"
   # Compare with what you entered in Meta console
   ```
4. **Try the test script:**
   ```bash
   ./test_webhook.sh
   ```

---

## 🎯 Quick Fix Checklist

Try these in order:

```bash
# 1. Restart everything
lsof -ti:8080 | xargs kill -9
killall ngrok
cd backend && go run main.go &
ngrok http 8080

# 2. Test locally
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=$(grep VERIFY_TOKEN .env | cut -d= -f2)&hub.challenge=test"

# 3. Get new ngrok URL
ngrok http 8080
# Copy the HTTPS URL

# 4. Update webhook in Meta Developer Console
# Use the new ngrok URL + /webhook
# Use the VERIFY_TOKEN from your .env file

# 5. Click "Verify and Save"
```

---

**Last Updated:** 2025-11-17
