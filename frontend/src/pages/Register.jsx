import React from 'react';
import { useForm } from 'react-hook-form';
import { authAPI } from '../services/api'; // Fixed import path
import { toast } from 'react-hot-toast';

const Register = ({ onRegister }) => {
    const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm();

    const onSubmit = async (data) => {
        try {
            const response = await authAPI.register(data);
            toast.success('Registration successful! Please login.');
            // Redirect to login after successful registration
            setTimeout(() => {
                window.location.href = '/login';
            }, 2000);
        } catch (error) {
            toast.error('Registration failed. Please try again.');
        }
    };

    return (
        <div className="auth-page">
            <div className="auth-container">
                <h1>Register for TMS</h1>
                <form onSubmit={handleSubmit(onSubmit)} className="auth-form">
                    <div className="form-group">
                        <label>Name</label>
                        <input
                            type="text"
                            {...register('name', { required: 'Name is required' })}
                        />
                        {errors.name && <span className="error">{errors.name.message}</span>}
                    </div>

                    <div className="form-group">
                        <label>Email</label>
                        <input
                            type="email"
                            {...register('email', {
                                required: 'Email is required',
                                pattern: {
                                    value: /^\S+@\S+$/i,
                                    message: 'Invalid email address'
                                }
                            })}
                        />
                        {errors.email && <span className="error">{errors.email.message}</span>}
                    </div>

                    <div className="form-group">
                        <label>Password</label>
                        <input
                            type="password"
                            {...register('password', {
                                required: 'Password is required',
                                minLength: {
                                    value: 6,
                                    message: 'Password must be at least 6 characters'
                                }
                            })}
                        />
                        {errors.password && <span className="error">{errors.password.message}</span>}
                    </div>

                    <button type="submit" disabled={isSubmitting} className="btn btn-primary">
                        {isSubmitting ? 'Registering...' : 'Register'}
                    </button>
                </form>

                <p>
                    Already have an account? <a href="/login">Login here</a>
                </p>
            </div>
        </div>
    );
};

export default Register;