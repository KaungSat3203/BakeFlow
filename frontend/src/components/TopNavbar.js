import { useState, useRef, useEffect } from 'react';

// Helper to format relative time
function getRelativeTime(timestamp) {
  const seconds = Math.floor((Date.now() - timestamp) / 1000);
  if (seconds < 60) return 'Just now';
  const minutes = Math.floor(seconds / 60);
  if (minutes < 60) return `${minutes} min ago`;
  const hours = Math.floor(minutes / 60);
  if (hours < 24) return `${hours}h ago`;
  const days = Math.floor(hours / 24);
  return `${days}d ago`;
}

export default function TopNavbar({
  toggleSidebar,
  notifications = [],
  unreadCount = 0,
  hasUnread = false,
  onBellClick,
  onMarkAllRead,
  onClearAll,
  onNotificationClick,
}) {
  const [open, setOpen] = useState(false);
  const panelRef = useRef(null);

  // Close when clicking outside
  useEffect(() => {
    function handleClickOutside(e) {
      if (open && panelRef.current && !panelRef.current.contains(e.target)) {
        setOpen(false);
      }
    }
    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, [open]);

  const togglePanel = () => {
    setOpen(o => !o);
    if (onBellClick) onBellClick();
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-white shadow-sm border-bottom bf-topnav py-0">
      <div className="container-fluid px-4 d-flex align-items-center position-relative">
        <span className="navbar-brand mb-0 fs-4 fw-bold">Order Management</span>
        <div className="ms-auto d-flex align-items-center gap-3">
          <div className="position-relative" ref={panelRef}>
            <button
              className={`btn btn-link text-secondary position-relative p-0 bf-bell-btn ${hasUnread ? 'pulse' : ''}`}
              aria-label={`Notifications${unreadCount ? `: ${unreadCount} new` : ''}`}
              onClick={togglePanel}
            >
              <i className={`bi bi-bell${hasUnread ? '-fill' : ''} fs-5`}></i>
              {unreadCount > 0 && (
                <span className="bf-badge-notification">
                  {unreadCount > 99 ? '99+' : unreadCount}
                </span>
              )}
            </button>

            {/* Notification dropdown panel */}
            <div className={`bf-notif-panel ${open ? 'show' : ''}`} role="region" aria-label="Notifications">
              
              {/* Header */}
              <div className="bf-notif-header">
                <div className="d-flex align-items-center justify-content-between">
                  <div>
                    <h6 className="mb-0 fw-bold">Notifications</h6>
                    {unreadCount > 0 && (
                      <small className="text-muted">{unreadCount} unread</small>
                    )}
                  </div>
                  <button 
                    className="btn btn-sm btn-link text-muted p-0"
                    onClick={() => setOpen(false)}
                    aria-label="Close"
                  >
                    <i className="bi bi-x-lg"></i>
                  </button>
                </div>
              </div>

              {/* Notification list */}
              <div className="bf-notif-body">
                {notifications.length > 0 ? (
                  notifications.map((notif) => (
                    <div
                      key={notif.id}
                      className={`bf-notif-item ${!notif.read ? 'unread' : ''}`}
                      onClick={() => {
                        if (onNotificationClick) onNotificationClick(notif.id);
                        // Navigate to orders page
                        window.location.href = '/admin/orders';
                      }}
                      style={{ cursor: 'pointer' }}
                    >
                      <div className="bf-notif-icon">
                        <i className="bi bi-bag-check-fill"></i>
                      </div>
                      <div className="bf-notif-content">
                        <div className="d-flex align-items-start justify-content-between">
                          <div className="flex-grow-1">
                            <div className="bf-notif-title">
                              <span className="badge bg-success-subtle text-success me-2">New Order</span>
                              <strong>#{notif.id}</strong>
                            </div>
                            <div className="bf-notif-text">
                              {notif.customer} ordered {notif.cake}
                            </div>
                            <div className="bf-notif-time">
                              <i className="bi bi-clock me-1"></i>
                              {getRelativeTime(notif.timestamp || Date.now())}
                            </div>
                          </div>
                          {!notif.read && (
                            <div className="bf-unread-dot"></div>
                          )}
                        </div>
                      </div>
                    </div>
                  ))
                ) : (
                  <div className="bf-notif-empty">
                    <i className="bi bi-bell-slash fs-1 text-muted mb-3"></i>
                    <p className="text-muted mb-0">No notifications yet</p>
                    <small className="text-muted">You're all caught up!</small>
                  </div>
                )}
              </div>

              {/* Footer actions */}
              {notifications.length > 0 && (
                <div className="bf-notif-footer">
                  <button
                    className="btn btn-sm btn-link text-muted"
                    onClick={() => {
                      if (onMarkAllRead) onMarkAllRead();
                    }}
                  >
                    <i className="bi bi-check2-all me-1"></i>
                    Mark all read
                  </button>
                  <button
                    className="btn btn-sm btn-link text-muted"
                    onClick={() => {
                      if (onClearAll) onClearAll();
                      setOpen(false);
                    }}
                  >
                    <i className="bi bi-trash me-1"></i>
                    Clear all
                  </button>
                </div>
              )}
            </div>
          </div>
        </div>
      </div>
    </nav>
  );
}
