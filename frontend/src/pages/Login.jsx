import React from 'react';
import { useForm } from 'react-hook-form';
import { authAPI } from '../services/api'; // Fixed import path
import { toast } from 'react-hot-toast';

const Login = ({ onLogin }) => {
    const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm();

    const onSubmit = async (data) => {
        try {
            const response = await authAPI.login(data);
            onLogin(response.data.user, response.data.token);
            toast.success('Login successful!');
        } catch (error) {
            toast.error('Login failed. Please check your credentials.');
        }
    };

    return (
        <div className="auth-page">
            <div className="auth-container">
                <h1>Login to TMS</h1>
                <form onSubmit={handleSubmit(onSubmit)} className="auth-form">
                    <div className="form-group">
                        <label>Email</label>
                        <input
                            type="email"
                            {...register('email', { required: 'Email is required' })}
                        />
                        {errors.email && <span className="error">{errors.email.message}</span>}
                    </div>

                    <div className="form-group">
                        <label>Password</label>
                        <input
                            type="password"
                            {...register('password', { required: 'Password is required' })}
                        />
                        {errors.password && <span className="error">{errors.password.message}</span>}
                    </div>

                    <button type="submit" disabled={isSubmitting} className="btn btn-primary">
                        {isSubmitting ? 'Logging in...' : 'Login'}
                    </button>
                </form>

                <p>
                    Don't have an account? <a href="/register">Register here</a>
                </p>
            </div>
        </div>
    );
};

export default Login;