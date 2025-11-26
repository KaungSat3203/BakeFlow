# Facebook Messenger Webhook - Cheat Sheet

## 🚀 Start Commands

```bash
# Start server
cd backend && go run main.go

# Start ngrok (new terminal)
ngrok http 8080

# Kill port 8080
lsof -ti:8080 | xargs kill -9
```

## 🧪 Test Commands

```bash
# Test local webhook verification
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"
# Expected: test

# Test wrong token (should fail)
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=wrongtoken&hub.challenge=test"
# Expected: Forbidden (403)

# Test health check
curl http://localhost:8080/
# Expected: BakeFlow Bot is running! ✅

# Test POST webhook
curl -X POST http://localhost:8080/webhook \
  -H "Content-Type: application/json" \
  -d '{"object":"page","entry":[{"messaging":[{"sender":{"id":"123"},"message":{"text":"Hi"}}]}]}'
# Expected: EVENT_RECEIVED

# Run automated tests
./test_webhook.sh
```

## 📝 Environment Variables

```bash
# View all env vars
cat .env

# Check specific token
grep VERIFY_TOKEN .env

# Show token value
echo $VERIFY_TOKEN

# Edit .env
nano .env
# or
vim .env
```

## 🔍 Debugging Commands

```bash
# Check if server is running
ps aux | grep "go run"

# Check what's on port 8080
lsof -i :8080

# View server logs (if redirected)
tail -f /tmp/bakeflow.log

# Follow logs in real-time
tail -f /tmp/bakeflow.log | grep "WEBHOOK"

# Check ngrok status
curl http://localhost:4040/api/tunnels
```

## 📦 Go Commands

```bash
# Install dependencies
go mod download

# Tidy dependencies
go mod tidy

# Run with specific port
PORT=3000 go run main.go

# Build binary
go build -o bakeflow main.go

# Run binary
./bakeflow

# Check Go version
go version
```

## 🌐 ngrok Commands

```bash
# Start tunnel
ngrok http 8080

# Use specific subdomain (paid plan)
ngrok http 8080 --subdomain=mybakery

# Start with auth token
ngrok authtoken YOUR_TOKEN

# View ngrok web interface
open http://localhost:4040
```

## 🐛 Troubleshooting Commands

```bash
# Full diagnostic
./test_webhook.sh

# Check .env exists
ls -la .env

# Verify .env format
cat .env | grep -v '^#' | grep '='

# Test connection to database
curl $DATABASE_URL 2>&1 | head -5

# Restart everything
pkill -f "go run"
killall ngrok
cd backend && go run main.go &
ngrok http 8080

# Check network connectivity
curl https://graph.facebook.com/v18.0/me?access_token=$PAGE_ACCESS_TOKEN
```

## 📊 Meta Developer Console URLs

```bash
# Open in browser:
open https://developers.facebook.com/apps
open https://developers.facebook.com/tools/debug/
```

## 🔧 Quick Fixes

### Fix: Port in use
```bash
lsof -ti:8080 | xargs kill -9
```

### Fix: .env not found
```bash
cd backend
touch .env
echo "VERIFY_TOKEN=verifyme123" >> .env
```

### Fix: ngrok expired
```bash
# Get new URL
ngrok http 8080
# Copy new URL and update in Meta console
```

### Fix: Can't connect to DB
```bash
# Test connection
go run main.go 2>&1 | grep "Connected to PostgreSQL"
```

## 📋 Meta Console Configuration

| Field | Value |
|-------|-------|
| Callback URL | `https://YOUR-NGROK-URL.ngrok-free.app/webhook` |
| Verify Token | `verifyme123` (from .env) |
| Webhook Fields | `messages`, `messaging_postbacks` |

## 🔑 Token Locations

```bash
# VERIFY_TOKEN
Location: backend/.env
Used for: Webhook verification
Meta Console: Messenger → Settings → Webhooks → Verify Token

# PAGE_ACCESS_TOKEN
Location: backend/.env
Used for: Sending messages
Meta Console: Messenger → Settings → Access Tokens

# DATABASE_URL
Location: backend/.env
Used for: PostgreSQL connection
Format: postgresql://user:pass@host/db
```

## 📝 File Paths

```
/Users/zuuji/Desktop/BakeFlow/backend/
├── main.go                        # Entry point
├── .env                           # YOUR TOKENS
├── controllers/webhook.go         # Webhook handlers
├── routes/routes.go               # HTTP routes
├── README.md                      # Full docs
├── QUICKSTART.md                  # 5-min guide
├── TROUBLESHOOTING.md             # Debug guide
└── test_webhook.sh                # Test script
```

## 🎯 Common Curl Patterns

```bash
# GET with query params
curl "http://localhost:8080/webhook?param1=value1&param2=value2"

# POST with JSON
curl -X POST http://localhost:8080/webhook \
  -H "Content-Type: application/json" \
  -d '{"key":"value"}'

# With verbose output
curl -v "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=verifyme123&hub.challenge=test"

# Show HTTP status code
curl -w "\n%{http_code}\n" "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=wrongtoken&hub.challenge=test"

# Follow redirects
curl -L "http://localhost:8080/webhook"

# Test with timeout
curl --max-time 5 "http://localhost:8080/webhook"
```

## 💾 Git Commands (Don't commit secrets!)

```bash
# Check what's staged
git status

# Make sure .env is ignored
echo ".env" >> .gitignore

# Check if .env is tracked
git ls-files | grep .env
# Should return nothing

# If .env is tracked, remove it:
git rm --cached .env
git commit -m "Remove .env from tracking"
```

## 🔐 Security Checklist

```bash
# ✅ Check .gitignore includes .env
cat .gitignore | grep .env

# ✅ Use .env.example instead
cp .env .env.example
# Edit .env.example to remove real tokens

# ✅ Never log full tokens
# Code already masks them: veri****e123

# ✅ Use HTTPS (ngrok does this automatically)
curl https://your-url.ngrok-free.app  # ✅
curl http://your-url.ngrok-free.app   # ❌
```

## 📞 Quick Reference Links

```
Documentation:
- QUICKSTART.md        → 5-minute setup
- README.md           → Full reference
- TROUBLESHOOTING.md  → Debug guide
- ARCHITECTURE.md     → System diagrams
- examples.go         → Code samples

External:
- Meta Developers: https://developers.facebook.com/apps
- Messenger Docs:  https://developers.facebook.com/docs/messenger-platform
- ngrok:          https://ngrok.com
- Go Docs:        https://golang.org/doc/
```

## 🎨 Log Emoji Reference

```
✅ Success
❌ Error
⚠️  Warning
📨 Incoming message
🚀 Server starting
➡️  Request received
⬅️  Request completed
🔧 Configuration
📋 Instructions
💡 Tip
🔐 Security
```

## ⚡ Quick Troubleshoot

```bash
# 1. Is server running?
curl http://localhost:8080/ && echo "✅ Server running" || echo "❌ Server down"

# 2. Is .env loaded?
go run main.go 2>&1 | grep "VERIFY_TOKEN"

# 3. Is ngrok working?
curl https://your-ngrok-url.ngrok-free.app/ && echo "✅ ngrok ok" || echo "❌ ngrok down"

# 4. Can we verify?
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=$(grep VERIFY_TOKEN .env | cut -d= -f2)&hub.challenge=test"

# 5. Full system check
./test_webhook.sh
```

---

## 💡 Pro Tips

```bash
# Auto-restart on file changes (install air first)
go install github.com/air-verse/air@latest
air

# Run in background and save logs
nohup go run main.go > logs.txt 2>&1 &

# Pretty print JSON logs
go run main.go 2>&1 | jq

# Watch for specific log patterns
tail -f logs.txt | grep --line-buffered "MESSAGE"

# Multiple terminals in tmux
tmux new -s bakeflow
# Ctrl+B then " to split horizontally
# Ctrl+B then arrow keys to navigate
```

---

**Keep this cheat sheet handy!** Bookmark or print for quick reference.

**Last Updated:** 2025-11-17
