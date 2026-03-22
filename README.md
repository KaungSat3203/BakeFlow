# BakeFlow: Social Commerce Automation for Modern Bakeries

**BakeFlow** is a full-stack social commerce solution designed to bridge the gap between social media engagement and structured order management. Built with a performance-first **Go** backend and a dynamic **Next.js** frontend, it automates the end-to-end ordering process—from Facebook Messenger interactions to secure payment verification.

---

## 🚀 The Vision
Small businesses often struggle to manage orders coming through fragmented social media channels. BakeFlow provides a centralized platform that transforms chat conversations into actionable orders, offering customers a seamless "chat-to-checkout" experience while giving business owners a robust admin suite to manage inventory, promotions, and deliveries.

## 🏗️ Technical Architecture

### Backend: High-Performance Go
The backend is engineered for reliability and speed, leveraging Go's efficient concurrency model.
- **Framework**: Built with `gorilla/mux` for clean, modular RESTful routing.
- **Concurrency**: Implements background workers for non-blocking tasks, such as automated stock reservation cleanup and Messenger persistent menu synchronization.
- **Database**: PostgreSQL (hosted on **Neon**) with a resilient connection strategy designed for serverless environments (handling cold-starts and connection pooling via PgBouncer).
- **Integrations**: Robust **Facebook Messenger Webhook** implementation with multi-stage message parsing and automated response logic.

### Frontend: Modern Next.js Interface
A responsive, high-fidelity web application built for both customers and administrators.
- **Core**: React 19 + Next.js 16/15 (using App Router patterns).
- **Styling**: Tailwind CSS 4 for a custom, premium design system with dark mode support.
- **State & Data**: Real-time stock validation and dynamic order tracking.
- **Assets**: Cloud-based image management via **Cloudinary** for lightning-fast delivery and storage.

---

## 🛠️ Key Technical Challenges & Solutions

### 1. High-Concurrency Stock Integrity
**Challenge**: Ensuring stock consistency during high-traffic "flash sales" or when multiple users interact with the bot simultaneously.
**Solution**: Developed a reservation-based stock system. Orders are temporarily "held" during the checkout phase, and a custom background cleanup job in Go (running every minute) automatically releases expired reservations, ensuring inventory is always accurate without locking the database for long periods.

### 2. Messenger-to-Web Context Switching
**Challenge**: Passing user context (like `SenderID` or `PSID`) securely from a chat interface to a webview order form.
**Solution**: Implemented a secure tokenization system that authenticates users across platforms, allowing for a personalized "Saved Orders" and "Active Order" experience within the Messenger webview without requiring traditional username/password login.

### 3. Serverless Database Resilience
**Challenge**: Managing database connections in a serverless environment (Neon/Vercel/Render) where idle connections are frequently terminated.
**Solution**: Implemented a custom database configuration with specific pool settings (`SetConnMaxLifetime`, `SetMaxOpenConns`) and a retry mechanism for critical queries, ensuring the application remains responsive even during "cold start" scenarios.

---

## 💼 Professional Engineering Practices
- **Environment Driven**: 100% configurable via environment variables for easy staging/production parity.
- **Modular Design**: Separation of concerns between controllers, routes, models, and configurations for high maintainability.
- **Security First**: Middleware for CORS, admin authentication, and payload verification.
---

## 🛠️ Getting Started

### Prerequisites
- Go 1.25+
- Node.js 18+
- PostgreSQL (Neon for serverless workflows)
- Cloudinary Account & Facebook Developer App (for Messenger)

### Setup Instructions
1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/BakeFlow.git
   cd BakeFlow
   ```

2. **Backend Setup**:
   ```bash
   cd backend
   cp .env.example .env # Add your DATABASE_URL, PAGE_ACCESS_TOKEN, etc.
   go run main.go
   ```

3. **Frontend Setup**:
   ```bash
   cd frontend
   npm install
   npm run dev
   ```

---

## 📈 Future Roadmap
- [ ] Payment Gateway Integration
- [ ] AI-powered recommendation engine based on past orders.
- [ ] Advanced analytics dashboard for revenue and inventory trends.

---

**Developed by The Tripods**  
*Turning social conversations into successful transactions.*
