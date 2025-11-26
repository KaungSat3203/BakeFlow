# 🔔 BakeFlow Admin Notification System

## Overview
A professional, big-tech style notification system for the BakeFlow admin dashboard with real-time order notifications, read/unread tracking, and elegant animations.

## ✨ Features

### 1. **Bell Icon with Badge**
- 🔴 Red gradient badge showing unread notification count
- 🫀 Pulsing animation when there are unread notifications
- 🎯 Bell icon fills when unread notifications exist
- 📊 Real-time count updates (99+ limit)

### 2. **Modern Dropdown Panel**
- 📏 400px width with smooth scale animation
- 🎨 Enhanced shadows and 16px border radius
- 📜 Scrollable notification list (max 400px height)
- 🔍 Header shows total unread count
- ❌ Close button in header

### 3. **Notification Items**
- 🟡 Yellow gradient highlight for unread items
- 🔴 Red dot indicator on unread items
- 🎨 Gradient bakery-colored icon (40×40px)
- 🏷️ "New Order" badge on each item
- ⏰ Relative timestamps ("Just now", "2 min ago", "5h ago")
- 📝 Order details (ID, customer name, cake type)
- 🖱️ Hover effects with smooth transitions
- ✅ Click to mark as read

### 4. **Empty State**
- 🔕 Bell-slash icon
- 💬 "You're all caught up!" message
- 🎨 Subtle styling to indicate no notifications

### 5. **Preview Cards (Facebook-style)**
- 📍 Bottom-right corner positioning
- 🎬 Slide-up animation on appear
- 🎬 Slide-down animation on dismiss
- ⏱️ Auto-dismisses after 6 seconds
- 🎨 Gradient header with bakery colors
- 👁️ "View Order" button to navigate
- ❌ "Dismiss" button to close
- 📱 Shows first new order only

### 6. **Persistence**
- 💾 LocalStorage-based persistence (`bakeflow_notifications`)
- 🔄 Notifications survive page navigation
- 📊 Read/unread states tracked per notification
- 🕐 Timestamps stored for relative time display

### 7. **Actions**
- ✅ "Mark all read" - marks all notifications as read
- 🗑️ "Clear all" - removes all notifications
- 👆 Click notification to mark as read and navigate

## 🏗️ Architecture

### Context API (`NotificationContext.js`)
```javascript
{
  notifications: [
    {
      id: number,
      customer: string,
      cake: string,
      time: string,
      read: boolean,
      timestamp: number,
      type: 'new_order'
    }
  ],
  unreadCount: number,
  hasUnread: boolean,
  addNotifications: (items) => void,
  markAsRead: (id) => void,
  markAllRead: () => void,
  clearAll: () => void
}
```

### Components
1. **TopNavbar** - Bell icon, badge, dropdown panel
2. **NotificationPreviewCard** - Bottom-right toast notification
3. **Dashboard** - Polls for new orders, triggers notifications
4. **Orders** - Displays notifications across pages

## 🎨 Styling Classes

### Notification Panel
- `.bf-notif-panel` - Main dropdown container
- `.bf-bell-btn.pulse` - Pulsing bell animation
- `.bf-badge-notification` - Red gradient badge
- `.bf-notif-header` - Dropdown header
- `.bf-notif-body` - Scrollable notification list
- `.bf-notif-footer` - Action buttons area

### Notification Items
- `.bf-notif-item` - Individual notification
- `.bf-notif-item.unread` - Yellow highlight for unread
- `.bf-notif-icon` - Gradient bakery icon
- `.bf-unread-dot` - Red dot indicator
- `.bf-notif-badge` - "New Order" badge
- `.bf-notif-empty` - Empty state container

### Preview Card
- `.bf-preview-card` - Bottom-right toast
- `.bf-preview-header` - Gradient header
- `.bf-preview-body` - Content area
- `.bf-preview-footer` - Action buttons
- `.bf-preview-card.hiding` - Slide-out animation

## 🎬 Animations

### Bell Pulse
```css
@keyframes bellPulse {
  0%, 100% { transform: scale(1); }
  50% { transform: scale(1.1); }
}
```

### Dropdown Scale
```css
@keyframes scaleIn {
  from { transform: scale(0.9); opacity: 0; }
  to { transform: scale(1); opacity: 1; }
}
```

### Preview Card Slide
```css
@keyframes slideInUp {
  from { transform: translateY(100%) scale(0.9); opacity: 0; }
  to { transform: translateY(0) scale(1); opacity: 1; }
}
```

## 🔄 Data Flow

```
Backend Polling (10s interval)
    ↓
Detect New Pending Orders
    ↓
addNotifications(newOrders)
    ↓
NotificationContext Updates
    ↓
├─→ TopNavbar (bell, badge, dropdown)
├─→ Preview Card (bottom-right toast)
└─→ localStorage (persistence)
```

## 📱 User Interactions

1. **New Order Arrives**
   - Bell pulses with animation
   - Badge shows count
   - Preview card slides up from bottom-right
   - Notification appears in dropdown as unread (yellow)

2. **User Opens Dropdown**
   - Dropdown scales in smoothly
   - Unread items highlighted in yellow
   - Relative timestamps shown

3. **User Clicks Notification**
   - Marks as read (yellow highlight removed)
   - Scrolls to orders section
   - Updates localStorage

4. **User Marks All Read**
   - All notifications marked as read
   - Yellow highlights removed
   - Bell stops pulsing
   - Badge disappears

5. **User Clears All**
   - All notifications removed
   - Empty state displayed
   - localStorage cleared

## 🎨 Design Philosophy

### Colors
- **Primary Bakery**: `#D8A35D`
- **Accent Bakery**: `#F4C27F`
- **Soft Background**: `#FFF8F0`
- **Unread Highlight**: Yellow gradient (`#fffbea` → `#fff8dc`)
- **Badge**: Red gradient (`#ff4757` → `#ff6348`)

### Typography
- **Header**: 16px, bold
- **Notification Title**: 15px, semi-bold
- **Notification Text**: 14px, regular
- **Timestamp**: 12px, muted

### Spacing
- **Panel Width**: 400px
- **Border Radius**: 16px (panel), 12px (items)
- **Icon Size**: 40×40px
- **Badge**: 20px height, 8px indicator dot

### Shadows
- **Panel**: `0 16px 48px rgba(0,0,0,0.15)`
- **Badge**: `0 2px 8px rgba(255,71,87,0.4)`
- **Icon**: `0 2px 8px rgba(0,0,0,0.1)`

## 🚀 Usage

### In Dashboard
```javascript
const { notifications, unreadCount, hasUnread, addNotifications, markAsRead, markAllRead, clearAll } = useNotifications();

<TopNavbar
  notifications={notifications}
  unreadCount={unreadCount}
  hasUnread={hasUnread}
  onMarkAllRead={markAllRead}
  onClearAll={clearAll}
  onNotificationClick={(id) => markAsRead(id)}
/>

<NotificationPreviewCard
  notification={previewCard}
  onClose={() => setPreviewCard(null)}
  onView={(id) => markAsRead(id)}
/>
```

## 🔧 Backend Integration

### Order Status Update Response
```json
{
  "message": "Order status updated",
  "notification_sent": true,
  "notification_error": null
}
```

### Migration Required
Run migration to add `sender_id` column:
```bash
psql "$DATABASE_URL" -f backend/migrations/003_add_sender_id.sql
```

## ✅ Completed Features

- ✅ Bell icon with pulsing animation
- ✅ Gradient red badge with count
- ✅ Modern dropdown panel design
- ✅ Read/unread state tracking
- ✅ Yellow highlight for unread items
- ✅ Relative timestamps
- ✅ Empty state design
- ✅ Preview card component
- ✅ Auto-dismiss after 6 seconds
- ✅ Mark all read functionality
- ✅ Clear all functionality
- ✅ Click to mark as read
- ✅ LocalStorage persistence
- ✅ Cross-page consistency
- ✅ Smooth animations and transitions
- ✅ Responsive design

## 🎯 Next Steps

1. ✅ Test end-to-end flow with real orders
2. ✅ Verify Messenger notifications work
3. ✅ Apply database migration
4. 🔄 Consider adding sound notifications (optional)
5. 🔄 Add notification settings page (optional)

## 📝 Notes

- Preview card only shows for the first new order to avoid spam
- Notifications persist across dashboard and orders pages
- Bell stops pulsing when all notifications are read
- Badge shows "99+" for counts over 99
- Relative time updates dynamically in dropdown
