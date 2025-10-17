import React from 'react';
import { Outlet, Link, useLocation } from 'react-router-dom';

const Layout = ({ user, onLogout }) => {
    const location = useLocation();

    return (
        <div className="layout">
            <nav className="navbar">
                <div className="container">
                    <div className="nav-brand">
                        <h2>Test Management System</h2>
                    </div>
                    <div className="nav-links">
                        <Link
                            to="/dashboard"
                            className={location.pathname === '/dashboard' ? 'active' : ''}
                        >
                            Dashboard
                        </Link>
                        <Link
                            to="/test-plans"
                            className={location.pathname.startsWith('/test-plans') ? 'active' : ''}
                        >
                            Test Plans
                        </Link>
                        <Link
                            to="/test-cases"
                            className={location.pathname.startsWith('/test-cases') ? 'active' : ''}
                        >
                            Test Cases
                        </Link>
                    </div>
                    <div className="nav-user">
                        <span>Welcome, {user?.email}</span>
                        <button onClick={onLogout} className="btn btn-secondary">
                            Logout
                        </button>
                    </div>
                </div>
            </nav>

            <main className="main-content">
                <div className="container">
                    <Outlet />
                </div>
            </main>
        </div>
    );
};

export default Layout;