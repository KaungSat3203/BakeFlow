# Product Management System - Quick Setup Guide

## 🚀 Quick Start (5 minutes)

### Step 1: Apply Database Migration

```bash
cd /Users/zuuji/Desktop/BakeFlow/backend
psql "$DATABASE_URL" -f migrations/004_create_products_system.sql
```

Expected output:
```
CREATE TABLE
CREATE INDEX
INSERT 0 4  (admin roles)
INSERT 0 5  (sample products)
```

### Step 2: Install Go Dependencies

```bash
cd /Users/zuuji/Desktop/BakeFlow/backend
go get github.com/gorilla/mux
go mod tidy
```

### Step 3: Restart Backend Server

```bash
# Stop current server (Ctrl+C)
# Start it again
go run main.go
```

You should see:
```
✅ .env file loaded successfully
⚙️  Setting up Facebook Messenger features...
✅ Facebook Messenger setup complete
🚀 Server running on http://localhost:8080
```

### Step 4: Test API

```bash
# Test products endpoint
curl "http://localhost:8080/api/products"
```

You should see JSON with 5 sample products!

### Step 5: Open Frontend

```bash
cd /Users/zuuji/Desktop/BakeFlow/frontend
npm run dev
```

Visit: http://localhost:3000/admin/products

## ✅ What You Get

### Backend (Go)
- ✅ `/api/products` - List products with filters
- ✅ `/api/products/:id` - Get single product
- ✅ `/api/products` (POST) - Create product
- ✅ `/api/products/:id` (PUT) - Update product
- ✅ `/api/products/:id/status` (PATCH) - Change status
- ✅ `/api/products/:id` (DELETE) - Archive product
- ✅ `/api/products/:id/logs` - View change history
- ✅ `/api/products/low-stock` - Get low stock alerts

### Frontend (Next.js)
- ✅ Product listing with search/filter
- ✅ Add new product form
- ✅ Edit product form
- ✅ Draft → Publish workflow
- ✅ Stock alerts
- ✅ Analytics display

### Database
- ✅ `products` table
- ✅ `product_logs` table (audit trail)
- ✅ `product_analytics` table (views, purchases)
- ✅ `admin_roles` table (RBAC)
- ✅ `admins` table (user management)

## 📱 How to Use

### Create a Product

1. Go to http://localhost:3000/admin/products
2. Click "Add New Product"
3. Fill in:
   - Name: "Lemon Tart"
   - Description: "Tangy lemon curd in buttery pastry"
   - Category: Tarts
   - Price: 15.99
   - Stock: 20
4. Click "Create Product" (saves as draft)
5. Or click "Save & Publish" (makes it active immediately)

### Publish a Draft

1. Go to products list
2. Find draft product
3. Click green checkmark icon
4. Product now shows as "active"

### Edit a Product

1. Click pencil icon on any product
2. Make changes
3. Click "Update Product"
4. All changes are logged in `product_logs`

### Archive a Product

1. Click red archive icon
2. Confirm
3. Product soft-deleted (can still see in logs)

## 🔍 Testing

### Test Backend API

```bash
# List all products
curl "http://localhost:8080/api/products"

# Filter by category
curl "http://localhost:8080/api/products?category=Cakes"

# Filter by status
curl "http://localhost:8080/api/products?status=active"

# Search
curl "http://localhost:8080/api/products?search=chocolate"

# Get one product
curl "http://localhost:8080/api/products/1"

# Create product
curl -X POST "http://localhost:8080/api/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Strawberry Cake",
    "description": "Fresh strawberry cake",
    "category": "Cakes",
    "price": 29.99,
    "stock": 8,
    "status": "active"
  }'

# Update status
curl -X PATCH "http://localhost:8080/api/products/1/status" \
  -H "Content-Type: application/json" \
  -d '{"status": "active"}'

# Get low stock
curl "http://localhost:8080/api/products/low-stock?threshold=10"

# Get product logs
curl "http://localhost:8080/api/products/1/logs"
```

### Test Frontend

1. **Products List:** http://localhost:3000/admin/products
   - Should show 5 sample products
   - Try search: type "chocolate"
   - Try filter: select "Cakes"
   - Try status filter: select "active"

2. **Add Product:** http://localhost:3000/admin/products/new
   - Fill form and click "Create Product"
   - Should redirect to products list

3. **Edit Product:** http://localhost:3000/admin/products/1
   - Change price to 27.99
   - Click "Update Product"
   - Check that it saved

### Verify Database

```bash
psql "$DATABASE_URL"
```

```sql
-- Check products
SELECT id, name, category, price, stock, status FROM products;

-- Check logs
SELECT id, product_id, action, created_at FROM product_logs ORDER BY created_at DESC LIMIT 5;

-- Check analytics
SELECT p.name, pa.views, pa.purchases 
FROM products p 
LEFT JOIN product_analytics pa ON p.id = pa.product_id;

-- Check roles
SELECT name, permissions FROM admin_roles;
```

## 🎨 Features Overview

### Draft → Publish Workflow
1. Create product as "draft" (not visible)
2. Review and test
3. Change status to "active" (goes live)
4. Can change back to "inactive" (hide) or "archived" (delete)

### Stock Management
- **Green badge:** In Stock (>= 10)
- **Yellow badge:** Low Stock (< 10)
- **Red badge:** Out of Stock (0)
- Alerts shown on product form

### Audit Logging
Every change is logged:
- Who made the change (admin_id)
- What changed (JSON diff)
- When it happened (timestamp)

### Analytics
- **Views:** Tracked when product is fetched
- **Purchases:** Manually tracked when order is placed
- Shown on product list

### Role-Based Access
- **Viewer:** Can only view products
- **Editor:** Can create/edit products
- **Manager:** Can delete products
- **Owner:** Full access + role management

## 📂 File Structure

```
backend/
├── migrations/
│   └── 004_create_products_system.sql  ← Database schema
├── models/
│   └── product.go                      ← Product models
├── controllers/
│   └── product_controller.go           ← API handlers
└── routes/
    └── routes.go                        ← API routes (updated)

frontend/
└── src/
    └── pages/
        └── admin/
            ├── products.js              ← Product list
            └── products/
                └── [id].js              ← Add/Edit form
```

## 🚨 Troubleshooting

### "relation 'products' does not exist"
→ Run the migration:
```bash
psql "$DATABASE_URL" -f migrations/004_create_products_system.sql
```

### "cannot find package github.com/gorilla/mux"
→ Install dependency:
```bash
cd backend
go get github.com/gorilla/mux
```

### Backend won't start
→ Check for port conflicts:
```bash
lsof -i :8080
kill -9 <PID>
```

### Products page shows nothing
→ Check backend is running:
```bash
curl "http://localhost:8080/api/products"
```

### CORS errors
→ Already handled in routes.go with `CORSMiddleware`

## 📚 Documentation

Full documentation: `/Users/zuuji/Desktop/BakeFlow/PRODUCT_MANAGEMENT.md`

Contains:
- Complete API reference
- Database schema details
- Frontend component docs
- RBAC implementation guide
- Analytics integration
- Advanced features

## 🎯 Next Steps

1. **Test the system** with the commands above
2. **Customize categories** in the dropdown
3. **Add image upload** instead of URLs
4. **Implement authentication** for admin_id tracking
5. **Connect to orders** to auto-update stock
6. **Build analytics dashboard** to visualize data

## 💡 Tips

- Use "draft" status while setting up products
- Publish when ready to go live
- Monitor low stock with `/api/products/low-stock`
- Review logs regularly for audit trail
- Set up alerts for out-of-stock products

---

🎉 **You're all set!** The product management system is ready to use.

Questions? Check PRODUCT_MANAGEMENT.md for detailed docs!
