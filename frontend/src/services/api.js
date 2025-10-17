import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1';

const api = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

// Add token to requests
api.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
});

// Response interceptor to handle errors
api.interceptors.response.use(
    (response) => response,
    (error) => {
        if (error.response?.status === 401) {
            localStorage.removeItem('token');
            window.location.href = '/login';
        }
        return Promise.reject(error);
    }
);

// Auth API
export const authAPI = {
    login: (credentials) => api.post('/auth/login', credentials),
    register: (userData) => api.post('/auth/register', userData),
    getProfile: () => api.get('/profile'),
};

// Test Plans API
export const testPlansAPI = {
    getAll: (projectId) => api.get(`/test-plans?project_id=${projectId}`),
    getById: (id) => api.get(`/test-plans/${id}`),
    create: (data) => api.post('/test-plans', data),
    update: (id, data) => api.put(`/test-plans/${id}`, data),
    addTestCase: (planId, testCaseId) => api.post(`/test-plans/${planId}/test-cases`, { test_case_id: testCaseId }),
};

// Test Cases API
export const testCasesAPI = {
    getAll: (projectId) => api.get(`/test-cases?project_id=${projectId}`),
    getById: (id) => api.get(`/test-cases/${id}`),
    create: (data) => api.post('/test-cases', data),
    update: (id, data) => api.put(`/test-cases/${id}`, data),
};

// Projects API
export const projectsAPI = {
    getAll: () => api.get('/projects'),
    create: (data) => api.post('/projects', data),
};

// Export API
export const exportAPI = {
    exportTestPlan: (id, options = {}) =>
        api.get(`/test-plans/${id}/export?include_history=${options.includeHistory || false}&include_comments=${options.includeComments || false}`, {
            responseType: 'blob'
        }),
    exportTestCase: (id, options = {}) =>
        api.get(`/test-cases/${id}/export?include_history=${options.includeHistory || false}&include_comments=${options.includeComments || false}`, {
            responseType: 'blob'
        }),
};

export default api;