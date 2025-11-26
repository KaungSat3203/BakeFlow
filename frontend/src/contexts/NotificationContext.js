import { createContext, useContext, useState, useEffect, useCallback } from 'react';

const NotificationContext = createContext();

export function NotificationProvider({ children }) {
  const [notifications, setNotifications] = useState([]);
  const [newOrderToast, setNewOrderToast] = useState({ show: false, items: [] });
  const [hasUnread, setHasUnread] = useState(false);

  // Load notifications from localStorage on mount
  useEffect(() => {
    const stored = localStorage.getItem('bakeflow_notifications');
    if (stored) {
      try {
        const parsed = JSON.parse(stored);
        // Filter out notifications older than 24 hours
        const oneDayAgo = Date.now() - (24 * 60 * 60 * 1000);
        const recent = parsed.filter(n => {
          const isRecent = (n.timestamp || 0) > oneDayAgo;
          // Also filter out notifications that have been read for more than 1 hour
          const isRead = n.read || n.isRead;
          if (isRead && n.readAt) {
            const oneHourAgo = Date.now() - (60 * 60 * 1000);
            return (n.readAt || 0) > oneHourAgo;
          }
          return isRecent;
        });
        console.log('📥 Loaded', recent.length, 'notifications from localStorage');
        setNotifications(recent);
        setHasUnread(recent.some(n => !n.read && !n.isRead));
      } catch (e) {
        console.error('Failed to parse stored notifications:', e);
        localStorage.removeItem('bakeflow_notifications');
      }
    }
  }, []);

  // Persist to localStorage whenever notifications change
  useEffect(() => {
    if (notifications.length === 0) {
      console.log('💾 Clearing localStorage (no notifications)');
      localStorage.removeItem('bakeflow_notifications');
    } else {
      console.log('💾 Saving', notifications.length, 'notifications to localStorage');
      localStorage.setItem('bakeflow_notifications', JSON.stringify(notifications));
    }
    setHasUnread(notifications.some(n => !n.read && !n.isRead));
  }, [notifications]);

  const addNotifications = useCallback((newOnes) => {
    if (!newOnes || newOnes.length === 0) return;
    console.log('🔔 Adding notifications:', newOnes);
    
    // Mark new notifications as unread and add timestamp
    const timestampedNotifications = newOnes.map(n => ({
      ...n,
      read: false,
      timestamp: n.timestamp || Date.now(),
      isRead: false,
      type: 'new_order'
    }));
    
    // Keep only last 20 notifications to prevent clutter
    setNotifications(prev => [...timestampedNotifications, ...prev].slice(0, 20));
    setNewOrderToast({ show: true, items: timestampedNotifications });
  }, []);

  const markAsRead = useCallback((notificationId) => {
    setNotifications(prev => 
      prev.map(n => n.id === notificationId ? { ...n, read: true, isRead: true, readAt: Date.now() } : n)
    );
  }, []);

  const markAllRead = useCallback(() => {
    const now = Date.now();
    setNotifications(prev => prev.map(n => ({ ...n, read: true, isRead: true, readAt: now })));
  }, []);

  const clearAll = useCallback(() => {
    setNotifications([]);
    localStorage.removeItem('bakeflow_notifications');
  }, []);

  const dismissToast = useCallback(() => {
    setNewOrderToast({ show: false, items: [] });
  }, []);

  // Auto-hide toast after 6s
  useEffect(() => {
    if (newOrderToast.show) {
      const t = setTimeout(dismissToast, 6000);
      return () => clearTimeout(t);
    }
  }, [newOrderToast.show, dismissToast]);

  const unreadCount = notifications.filter(n => !n.read && !n.isRead).length;

  return (
    <NotificationContext.Provider value={{
      notifications,
      newOrderToast,
      unreadCount,
      hasUnread,
      addNotifications,
      markAsRead,
      markAllRead,
      clearAll,
      dismissToast,
    }}>
      {children}
    </NotificationContext.Provider>
  );
}

export function useNotifications() {
  const context = useContext(NotificationContext);
  if (!context) {
    throw new Error('useNotifications must be used within NotificationProvider');
  }
  return context;
}
