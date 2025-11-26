# Shopping Cart Migration Instructions

## Step 1: Run the database migration

You need to run this migration **BEFORE** starting the new server:

```bash
# Connect to your Neon PostgreSQL database
psql "YOUR_DATABASE_URL_HERE"

# Or if you have the connection string in your .env:
# psql $(grep DATABASE_URL .env | cut -d '=' -f2)
```

Then run the migration SQL:

```sql
-- Migration: Add order_items table for shopping cart functionality

-- Update orders table structure
ALTER TABLE orders 
  DROP COLUMN IF EXISTS product,
  DROP COLUMN IF EXISTS quantity,
  ADD COLUMN IF NOT EXISTS delivery_type TEXT,
  ADD COLUMN IF NOT EXISTS total_items INT DEFAULT 0;

-- Create order_items table
CREATE TABLE IF NOT EXISTS order_items (
  id SERIAL PRIMARY KEY,
  order_id INT NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
  product TEXT NOT NULL,
  quantity INT NOT NULL,
  price DECIMAL(10,2) DEFAULT 0.00,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for faster queries
CREATE INDEX IF NOT EXISTS idx_order_items_order_id ON order_items(order_id);

-- Comments for documentation
COMMENT ON TABLE order_items IS 'Individual items in each order (supports multiple items per order)';
COMMENT ON COLUMN orders.total_items IS 'Total quantity of all items in the order';
COMMENT ON COLUMN orders.delivery_type IS 'pickup or delivery';
```

## Step 2: Verify the migration

```sql
-- Check orders table structure
\d orders

-- Check order_items table structure
\d order_items

-- Should see both tables
\dt
```

## Step 3: Start the server

```bash
go run main.go
```

## New Features:

### Shopping Cart Flow:
1. **Select Product** → Choose item (e.g., Chocolate Cake)
2. **Select Quantity** → Choose how many (1-5)
3. **Add to Cart** → Item added with confirmation
4. **Add More or Checkout?** → Continue shopping or proceed
5. **Repeat** → Add more items to cart
6. **Checkout** → View cart summary
7. **Enter Name** → Provide customer name
8. **Delivery Type** → Pickup or Delivery
9. **Address** (if delivery) → Enter delivery address
10. **Confirm Order** → Review all items and confirm

### Database Structure:
- `orders` table: Main order info (customer, delivery type, address, status, total_items)
- `order_items` table: Individual items in each order (product, quantity, price)

### Example Order:
```
orders:
  id: 42
  customer_name: "John Doe"
  delivery_type: "delivery"
  address: "123 Main St"
  status: "pending"
  total_items: 6

order_items:
  {order_id: 42, product: "Chocolate Cake", quantity: 2}
  {order_id: 42, product: "Coffee", quantity: 1}
  {order_id: 42, product: "Croissant", quantity: 3}
```

## Testing Checklist:

- [ ] Run migration successfully
- [ ] Start server without errors
- [ ] Add first item to cart
- [ ] Click "Add More Items" and add second item
- [ ] Click "Add More Items" and add third item
- [ ] Click "Checkout"
- [ ] Enter name
- [ ] Select delivery type
- [ ] Enter address (if delivery)
- [ ] Confirm order
- [ ] Check database for order and order_items records
- [ ] Verify all items are saved correctly

## Notes:

- Old orders with `product` and `quantity` columns will be dropped
- Make sure to backup data if you have important old orders
- The migration is idempotent (can be run multiple times safely)
