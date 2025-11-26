# 🧪 Quick Testing Guide - New Features

## ⚡ 5-Minute Test Plan

### **Step 1: Apply Database Migration** (CRITICAL - Do this first!)

1. Go to your **Neon Dashboard**: https://console.neon.tech
2. Select your **BakeFlow** database
3. Click **SQL Editor**
4. Copy and paste this SQL:

```sql
-- Add new columns to orders table
ALTER TABLE orders 
  ADD COLUMN IF NOT EXISTS delivery_fee DECIMAL(10,2) DEFAULT 0.00,
  ADD COLUMN IF NOT EXISTS subtotal DECIMAL(10,2) DEFAULT 0.00,
  ADD COLUMN IF NOT EXISTS total_amount DECIMAL(10,2) DEFAULT 0.00,
  ADD COLUMN IF NOT EXISTS reordered_from INT REFERENCES orders(id) ON DELETE SET NULL,
  ADD COLUMN IF NOT EXISTS rating_id INT REFERENCES ratings(id) ON DELETE SET NULL,
  ADD COLUMN IF NOT EXISTS completed_at TIMESTAMP;

-- Create ratings table
CREATE TABLE IF NOT EXISTS ratings (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  user_id TEXT NOT NULL,
  stars INT NOT NULL CHECK (stars >= 1 AND stars <= 5),
  comment TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_ratings_order_id ON ratings(order_id);
CREATE INDEX IF NOT EXISTS idx_ratings_user_id ON ratings(user_id);
```

5. Click **Run** ✅

---

### **Step 2: Test Order History** 📋

**In Messenger, type:**
```
orders
```

**Expected result:**
- If you have orders: Beautiful cards showing past orders
- If no orders: "No Orders Yet!" message with emoji

**What to check:**
- [ ] Order cards display correctly
- [ ] Status badges show (⏳ Pending)
- [ ] Items list shows with emojis
- [ ] Total amount displays
- [ ] Reorder and Rate buttons appear

---

### **Step 3: Test Delivery Fees** 💰

**Place a new order:**
1. Type `menu`
2. Select any product (e.g., Chocolate Cake)
3. Choose quantity: `2`
4. Add to cart
5. Type your name: `Test User`
6. Choose **Delivery**
7. Enter address: `123 Main St, Yangon`

**Expected result at confirmation:**
```
✅ Order Confirmed!

Order #X

🛒 Your Order:
• 2× 🍫 Chocolate Cake - $50.00

💰 Pricing:
Subtotal: $50.00
Delivery Fee: $3.00
━━━━━━━━━━━━
Total: $53.00
```

**What to check:**
- [ ] Subtotal shows correct item prices
- [ ] Delivery fee calculated ($3 for Yangon)
- [ ] Total = Subtotal + Delivery Fee
- [ ] Pricing breakdown looks clean

---

### **Step 4: Test Pickup (Free Delivery)** 🏠

**Place pickup order:**
1. Type `menu`
2. Select Coffee
3. Quantity: `1`
4. Name: `Pickup Test`
5. Choose **Pickup**

**Expected result:**
```
💰 Pricing:
Subtotal: $5.00
Delivery Fee: $0.00
━━━━━━━━━━━━
Total: $5.00
```

**What to check:**
- [ ] Delivery fee is $0.00 for pickup
- [ ] Address shows "Pickup at store"

---

### **Step 5: Test Reorder** 🔄

1. Type `orders` to see order history
2. Click **🔄 Reorder** button on any order
3. Should see: "Reordering from Order #X"
4. Cart pre-filled with items
5. Just confirm name/address

**What to check:**
- [ ] Cart has all items from original order
- [ ] Quantities match original order
- [ ] Can complete checkout quickly

---

### **Step 6: Test Business Hours** ⏰

**To test closed hours:**
1. Temporarily change the business hours in code:
   - Open `backend/controllers/webhook.go`
   - Find `func isBusinessOpen()`
   - Change to return `false` for testing:
   ```go
   func isBusinessOpen() bool {
       return false  // Test closed mode
   }
   ```
2. Restart server
3. Try typing `menu`

**Expected result:**
```
🔒 We're Currently Closed

Business Hours: 8:00 AM - 8:00 PM
We'll be open again at 8:00 AM tomorrow.
```

**What to check:**
- [ ] Blocks ordering when closed
- [ ] Shows next opening time
- [ ] Friendly message

**After testing, revert the change!**

---

### **Step 7: Test Rating System** ⭐

1. Type `orders`
2. Click **⭐ Rate** button on any order
3. Select a star rating (try 5 stars first)

**Expected result:**
```
🎉 Thank you so much!
We're thrilled you loved your order! ⭐⭐⭐⭐⭐
```

**Try different ratings:**
- **5 stars:** Thrilled message
- **3 stars:** Thank you + improvement message
- **1 star:** Sorry + we'll do better message

**What to check:**
- [ ] Rating buttons appear (1-5 stars)
- [ ] Thank you message changes based on rating
- [ ] Can skip rating
- [ ] Check database has new row in `ratings` table

---

### **Step 8: Verify Database** 🗄️

**In Neon SQL Editor, run:**

```sql
-- Check orders have new columns
SELECT id, customer_name, subtotal, delivery_fee, total_amount 
FROM orders 
ORDER BY id DESC 
LIMIT 5;

-- Check ratings saved
SELECT * FROM ratings ORDER BY id DESC;
```

**What to check:**
- [ ] Orders table has subtotal, delivery_fee, total_amount filled
- [ ] Ratings table has entries
- [ ] Prices match what you ordered

---

## 🎯 Full Flow Test (End-to-End)

**Complete order journey:**

1. **Start:** Type `menu` or click "Get Started"
2. **Browse:** See product cards
3. **Select:** Choose Chocolate Cake
4. **Quantity:** Pick 2
5. **Add more:** Click "Add More" → Choose Coffee → Quantity 1
6. **Checkout:** Click "Checkout"
7. **Name:** Enter name
8. **Delivery:** Choose Delivery
9. **Address:** Enter address
10. **Review:** See order summary with pricing breakdown
11. **Confirm:** Click "Confirm Order"
12. **Success:** See order confirmation with #ID and total
13. **History:** Type `orders` - see your order
14. **Reorder:** Click Reorder button - cart pre-fills
15. **Rate:** Go back to history, click Rate, give 5 stars
16. **Done:** See thank you message!

---

## 📸 Screenshots to Check

Take screenshots of:
1. ✅ Order history cards
2. ✅ Pricing breakdown in order summary
3. ✅ Order confirmation with total
4. ✅ Reorder working (pre-filled cart)
5. ✅ Business hours closed message
6. ✅ Rating buttons
7. ✅ Thank you message after rating

---

## 🐛 Common Issues

| Issue | Solution |
|-------|----------|
| "Orders not showing" | Run SQL migration in Neon |
| "Prices are $0.00" | Restart server (already fixed in code) |
| "Reorder button not working" | Check you have orders with items |
| "Business hours not blocking" | Check server time matches Myanmar |
| "Rating not saving" | Run SQL migration to create ratings table |

---

## ✅ Success Checklist

- [ ] Database migration applied in Neon
- [ ] Server restarted successfully
- [ ] Order history displays beautifully
- [ ] Delivery fees calculate correctly
- [ ] Pickup shows $0.00 fee
- [ ] Reorder pre-fills cart
- [ ] Business hours blocks when closed
- [ ] Rating system works
- [ ] Different rating messages show
- [ ] Database has new columns and ratings table
- [ ] Bilingual support works (test in Myanmar language)

---

## 🚀 Ready for Production?

Once all tests pass:

1. ✅ All 5 features working
2. ✅ Database properly migrated
3. ✅ No errors in server logs
4. ✅ Tested in both languages
5. ✅ ngrok tunnel working

**You're ready to go live!** 🎉

Consider deploying to Render/Railway for production instead of ngrok.

---

**Need help? Check the logs:**
```bash
# In terminal where server is running
# You'll see output like:
# ✅ Order created: #123
# ⭐ Rating saved for order #123
```
