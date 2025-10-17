import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from 'react-query';
import { Toaster } from 'react-hot-toast';
import Layout from './components/Layout';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import TestPlans from './pages/TestPlans';
import TestCases from './pages/TestCases';
import { authAPI } from './services/api'; // Fixed import path

const queryClient = new QueryClient();

function App() {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
            authAPI.getProfile()
                .then(response => {
                    setUser(response.data);
                })
                .catch(() => {
                    localStorage.removeItem('token');
                })
                .finally(() => {
                    setLoading(false);
                });
        } else {
            setLoading(false);
        }
    }, []);

    const login = (userData, token) => {
        localStorage.setItem('token', token);
        setUser(userData);
    };

    const logout = () => {
        localStorage.removeItem('token');
        setUser(null);
    };

    if (loading) {
        return <div className="loading">Loading...</div>;
    }

    return (
        <QueryClientProvider client={queryClient}>
            <Router>
                <div className="App">
                    <Toaster position="top-right" />
                    <Routes>
                        <Route
                            path="/login"
                            element={!user ? <Login onLogin={login} /> : <Navigate to="/dashboard" />}
                        />
                        <Route
                            path="/register"
                            element={!user ? <Register onRegister={login} /> : <Navigate to="/dashboard" />}
                        />
                        <Route
                            path="/*"
                            element={user ? <Layout user={user} onLogout={logout} /> : <Navigate to="/login" />}
                        >
                            <Route path="dashboard" element={<Dashboard />} />
                            <Route path="test-plans" element={<TestPlans />} />
                            <Route path="test-cases" element={<TestCases />} />
                            <Route path="" element={<Navigate to="/dashboard" />} />
                        </Route>
                    </Routes>
                </div>
            </Router>
        </QueryClientProvider>
    );
}

export default App;