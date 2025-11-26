import { useEffect } from 'react';
import Link from 'next/link';
import { useRouter } from 'next/router';

// Sidebar with collapse + bootstrap tooltips
export default function Sidebar({ open, toggle }) {
  const router = useRouter();
  
  useEffect(() => {
    // Init Bootstrap tooltips dynamically (guard for SSR)
    if (typeof window !== 'undefined' && window.bootstrap) {
      const tooltipTriggerList = [].slice.call(document.querySelectorAll('[data-bs-toggle="tooltip"]'));
      tooltipTriggerList.forEach(el => new window.bootstrap.Tooltip(el));
    }
  }, [open]);

  const navItems = [
    { href: '/admin', icon: 'speedometer2', label: 'Dashboard' },
    { href: '/admin/orders', icon: 'receipt', label: 'Orders' },
    { href: '/admin/products', icon: 'box-seam', label: 'Products' },
    { href: '/admin/customers', icon: 'people', label: 'Customers' },
    { href: '/admin/analytics', icon: 'graph-up', label: 'Analytics' },
    { href: '/admin/settings', icon: 'gear', label: 'Settings' }
  ];

  return (
    <aside className={`bf-sidebar bg-white border-end ${open ? 'expanded' : 'collapsed'}`}>
      <div className="bf-sidebar-header px-3 py-0 border-bottom d-flex align-items-center justify-content-center">
        {open ? (
          <>
            <div className="d-flex align-items-center gap-2 me-auto">
              <i className="bi bi-shop fs-3 text-primary-bake"></i>
              <span className="fs-4 fw-bold sidebar-brand">BakeFlow</span>
            </div>
            <button className="btn btn-sm btn-outline-secondary" onClick={toggle} aria-label="Toggle sidebar">
              <i className="bi bi-x-lg"></i>
            </button>
          </>
        ) : (
          <button className="btn btn-sm btn-outline-secondary" onClick={toggle} aria-label="Toggle sidebar">
            <i className="bi bi-list"></i>
          </button>
        )}
      </div>
      <nav className="flex-grow-1 p-3">
        <ul className="nav flex-column gap-2">
          {navItems.map(item => (
            <li key={item.href} className="nav-item">
              <Link href={item.href} className={`nav-link text-secondary d-flex align-items-center bf-nav-link ${router.pathname === item.href ? 'active' : ''}`} data-bs-toggle={!open ? 'tooltip' : undefined} data-bs-placement="right" title={!open ? item.label : undefined}>
                <i className={`bi bi-${item.icon} fs-5 me-2`}></i>
                {open && <span>{item.label}</span>}
              </Link>
            </li>
          ))}
        </ul>
      </nav>
      <div className="bf-sidebar-footer p-3 border-top">
        <div className="d-flex align-items-center">
          <div className="rounded-circle bg-primary-bake bg-opacity-10 p-2">
            <i className="bi bi-person-circle fs-5 text-primary-bake"></i>
          </div>
          {open && (
            <div className="ms-2 flex-grow-1">
              <div className="fw-semibold small text-dark">Admin User</div>
              <div className="text-muted small">admin@bakeflow.com</div>
            </div>
          )}
        </div>
      </div>
    </aside>
  );
}
