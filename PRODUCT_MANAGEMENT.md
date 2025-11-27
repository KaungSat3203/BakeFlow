# Product Management System - Documentation

## Overview

This is a complete product management system for BakeFlow with:
- ✅ Full CRUD operations (Create, Read, Update, Delete)
- ✅ Product status management (Draft → Active → Inactive → Archived)
- ✅ Stock tracking with low stock alerts
- ✅ Product analytics (views, purchases)
- ✅ Audit logging for all changes
- ✅ Search and filtering
- ✅ Role-based access control (RBAC) support

## Database Setup

### 1. Apply the Migration

Run the migration to create all necessary tables:

```bash
cd backend
psql "$DATABASE_URL" -f migrations/004_create_products_system.sql
```

This creates:
- `products` - Main products table
- `product_logs` - Audit trail for all product changes
- `product_analytics` - View and purchase tracking
- `admin_roles` - Role-based access control
- `admins` - Admin users with roles

### 2. Verify Tables

```sql
-- Check tables were created
\dt

-- View sample products
SELECT * FROM products;

-- Check admin roles
SELECT * FROM admin_roles;
```

## Backend API

### Product Endpoints

#### GET /api/products
List all products with filtering and pagination

**Query Parameters:**
- `category` - Filter by category (e.g., "Cakes", "Cupcakes")
- `status` - Filter by status (draft, active, inactive, archived)
- `search` - Search in name and description
- `min_price` - Minimum price filter
- `max_price` - Maximum price filter
- `sort_by` - Sort field (name, price, stock, created_at, views, purchases)
- `sort_dir` - Sort direction (ASC, DESC)
- `limit` - Results per page (default: 50, max: 100)
- `offset` - Pagination offset

**Example:**
```bash
curl "http://localhost:8080/api/products?category=Cakes&status=active&limit=10"
```

**Response:**
```json
{
  "products": [
    {
      "id": 1,
      "name": "Chocolate Cake",
      "description": "Rich chocolate cake with ganache",
      "category": "Cakes",
      "price": 25.99,
      "stock": 10,
      "image_url": "https://example.com/cake.jpg",
      "status": "active",
      "views": 150,
      "purchases": 25,
      "low_stock": false,
      "out_of_stock": false,
      "created_at": "2025-01-01T10:00:00Z",
      "updated_at": "2025-01-15T14:30:00Z"
    }
  ],
  "count": 1
}
```

#### GET /api/products/:id
Get single product details

**Example:**
```bash
curl "http://localhost:8080/api/products/1"
```

#### POST /api/products
Create new product

**Request Body:**
```json
{
  "name": "Vanilla Cupcake",
  "description": "Classic vanilla cupcake with buttercream",
  "category": "Cupcakes",
  "price": 3.99,
  "stock": 50,
  "image_url": "https://example.com/cupcake.jpg",
  "status": "draft"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Product created successfully",
  "product": { ... }
}
```

#### PUT /api/products/:id
Update existing product

**Request Body:** Same as POST

#### PATCH /api/products/:id/status
Update product status only

**Request Body:**
```json
{
  "status": "active"
}
```

Valid statuses: `draft`, `active`, `inactive`, `archived`

#### DELETE /api/products/:id
Soft delete (archive) a product

**Response:**
```json
{
  "success": true,
  "message": "Product archived successfully"
}
```

#### GET /api/products/:id/logs
Get audit log for a product

**Response:**
```json
{
  "logs": [
    {
      "id": 1,
      "product_id": 1,
      "action": "UPDATE",
      "admin": "admin@bakeflow.com",
      "changes": {
        "old": { "price": 25.99 },
        "new": { "price": 29.99 }
      },
      "created_at": "2025-01-15T14:30:00Z"
    }
  ],
  "count": 1
}
```

#### GET /api/products/low-stock
Get products with low stock

**Query Parameters:**
- `threshold` - Stock threshold (default: 10)

**Example:**
```bash
curl "http://localhost:8080/api/products/low-stock?threshold=5"
```

## Frontend Pages

### 1. Products List Page
**URL:** `/admin/products`

Features:
- View all products in a table
- Search by name/description
- Filter by category and status
- Sort by various fields
- Quick actions: Edit, Publish, Deactivate, Archive
- Stock status badges
- View counts

### 2. Add/Edit Product Page
**URL:** `/admin/products/new` or `/admin/products/:id`

Features:
- Form validation
- Image preview
- Status selection
- Save as draft or publish directly
- Low stock warnings
- Stock alerts

### 3. Navigation
Products link is available in the sidebar under the "Products" menu item

## Product Workflow

### Draft → Active (Publish)

1. Create product with status = "draft"
2. Fill in all required fields
3. Click "Save & Publish" or change status to "active"
4. Product becomes visible to customers

### Active → Inactive (Deactivate)

1. Go to products list
2. Click pause icon on active product
3. Product hidden from customers but not deleted

### Any Status → Archived (Delete)

1. Click archive icon
2. Product soft-deleted (deleted_at timestamp set)
3. Still in database for reporting

## Stock Management

### Stock Alerts

- **Low Stock:** Stock < 10 (configurable)
- **Out of Stock:** Stock = 0

Alerts appear on:
- Product list page (badges)
- Edit product page (warning box)
- Low stock endpoint

### Updating Stock

When orders are placed, update stock:

```go
// In your order processing code
_, err := db.Exec(`
    UPDATE products 
    SET stock = stock - $1 
    WHERE id = $2 AND stock >= $1
`, quantity, productID)

// Increment purchase count
models.IncrementPurchases(db, productID)
```

## Analytics

### View Tracking

Views are automatically tracked when fetching a single product:

```go
// Automatically called in GetProduct handler
models.IncrementViews(db, productID)
```

### Purchase Tracking

Call when an order is completed:

```go
models.IncrementPurchases(db, productID)
```

### Viewing Analytics

Analytics are included in product API responses:
- `views` - Total product views
- `purchases` - Total purchases

## Audit Logging

All product changes are logged automatically:

- CREATE - Product created
- UPDATE - Product updated
- DELETE - Product archived
- STATUS_CHANGE - Status updated

View logs via API:
```bash
curl "http://localhost:8080/api/products/1/logs"
```

## Role-Based Access Control (RBAC)

### Default Roles

1. **Viewer** - Read-only access
   ```json
   {"products": ["read"], "analytics": ["read"]}
   ```

2. **Editor** - Can create and edit
   ```json
   {"products": ["read", "create", "update"], "analytics": ["read"]}
   ```

3. **Manager** - Full product access
   ```json
   {"products": ["read", "create", "update", "delete"], "analytics": ["read", "manage"]}
   ```

4. **Owner** - Full system access
   ```json
   {"products": ["read", "create", "update", "delete"], "analytics": ["read", "manage"], "roles": ["manage"]}
   ```

### Implementing RBAC Middleware

Add to your authentication flow:

```go
func CheckPermission(requiredPermission string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Get admin from context/session
            adminID := getAdminIDFromSession(r)
            
            // Check permission
            var permissions json.RawMessage
            err := db.QueryRow(`
                SELECT ar.permissions 
                FROM admins a 
                JOIN admin_roles ar ON a.role_id = ar.id 
                WHERE a.id = $1
            `, adminID).Scan(&permissions)
            
            if err != nil || !hasPermission(permissions, requiredPermission) {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}
```

## Testing

### 1. Test with Sample Data

The migration includes sample products. Test the API:

```bash
# List products
curl "http://localhost:8080/api/products"

# Get single product
curl "http://localhost:8080/api/products/1"

# Create product
curl -X POST "http://localhost:8080/api/products" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Cake",
    "category": "Cakes",
    "price": 19.99,
    "stock": 15,
    "status": "draft"
  }'

# Update status
curl -X PATCH "http://localhost:8080/api/products/1/status" \
  -H "Content-Type: application/json" \
  -d '{"status": "active"}'

# Check low stock
curl "http://localhost:8080/api/products/low-stock"
```

### 2. Test Frontend

```bash
cd frontend
npm run dev
```

Visit:
- http://localhost:3000/admin/products - Products list
- http://localhost:3000/admin/products/new - Add product
- http://localhost:3000/admin/products/1 - Edit product

### 3. Verify Logging

```sql
-- Check product logs
SELECT * FROM product_logs ORDER BY created_at DESC LIMIT 10;

-- Check analytics
SELECT p.name, pa.views, pa.purchases 
FROM products p 
LEFT JOIN product_analytics pa ON p.id = pa.product_id
ORDER BY pa.views DESC;
```

## Troubleshooting

### Products Not Showing

1. Check database connection:
```bash
psql "$DATABASE_URL" -c "SELECT COUNT(*) FROM products;"
```

2. Check backend logs for errors

3. Verify CORS headers are set

### Can't Create Products

1. Check validation errors in browser console
2. Verify all required fields are filled
3. Check backend logs for SQL errors

### Images Not Loading

1. Verify image URLs are accessible
2. Check CORS on image host
3. Use placeholder if image fails to load

## Next Steps

1. **Image Upload:** Implement file upload instead of URLs
2. **Categories:** Make categories dynamic (database table)
3. **Variants:** Add product variants (size, flavor)
4. **Inventory:** More advanced stock management
5. **Discounts:** Add pricing and discount rules
6. **Bulk Operations:** Import/export products via CSV

## Support

For issues or questions:
- Check backend logs: `backend/logs/`
- Check browser console for frontend errors
- Review API responses in Network tab
- See TROUBLESHOOTING.md for common issues
