import React, { useState } from 'react';
import { useQuery, useMutation, useQueryClient } from 'react-query';
import { useForm } from 'react-hook-form'; // Add this import
import { testPlansAPI } from '../services/api'; // Fixed import path
import { toast } from 'react-hot-toast';

const TestPlans = () => {
    const [showCreateForm, setShowCreateForm] = useState(false);
    const queryClient = useQueryClient();

    const { data: testPlans, isLoading } = useQuery(
        'testPlans',
        () => testPlansAPI.getAll('demo-project-id'),
        { enabled: false } // Disable auto-fetch for demo
    );

    const createMutation = useMutation(testPlansAPI.create, {
        onSuccess: () => {
            queryClient.invalidateQueries('testPlans');
            setShowCreateForm(false);
            toast.success('Test plan created successfully!');
        },
        onError: () => {
            toast.error('Failed to create test plan');
        }
    });

    const handleCreate = (data) => {
        createMutation.mutate({
            ...data,
            project_id: 'demo-project-id'
        });
    };

    return (
        <div className="test-plans">
            <div className="page-header">
                <h1>Test Plans</h1>
                <button
                    className="btn btn-primary"
                    onClick={() => setShowCreateForm(true)}
                >
                    Create Test Plan
                </button>
            </div>

            {showCreateForm && (
                <CreateTestPlanForm
                    onSubmit={handleCreate}
                    onCancel={() => setShowCreateForm(false)}
                    loading={createMutation.isLoading}
                />
            )}

            <div className="test-plans-list">
                {isLoading ? (
                    <p>Loading...</p>
                ) : testPlans?.data?.length > 0 ? (
                    testPlans.data.map(plan => (
                        <TestPlanCard key={plan.id} plan={plan} />
                    ))
                ) : (
                    <div className="empty-state">
                        <h3>No test plans yet</h3>
                        <p>Create your first test plan to get started</p>
                    </div>
                )}
            </div>
        </div>
    );
};

const CreateTestPlanForm = ({ onSubmit, onCancel, loading }) => {
    const { register, handleSubmit, formState: { errors } } = useForm();

    return (
        <div className="card">
            <h3>Create New Test Plan</h3>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div className="form-group">
                    <label>Name</label>
                    <input
                        type="text"
                        {...register('name', { required: 'Name is required' })}
                    />
                    {errors.name && <span className="error">{errors.name.message}</span>}
                </div>

                <div className="form-group">
                    <label>Description</label>
                    <textarea
                        {...register('description')}
                    />
                </div>

                <div className="form-group">
                    <label>Deadline</label>
                    <input
                        type="datetime-local"
                        {...register('deadline')}
                    />
                </div>

                <div className="form-actions">
                    <button type="button" onClick={onCancel} className="btn btn-secondary">
                        Cancel
                    </button>
                    <button type="submit" disabled={loading} className="btn btn-primary">
                        {loading ? 'Creating...' : 'Create'}
                    </button>
                </div>
            </form>
        </div>
    );
};

const TestPlanCard = ({ plan }) => {
    return (
        <div className="card test-plan-card">
            <h3>{plan.name}</h3>
            <p>{plan.description}</p>
            <div className="test-plan-meta">
                <span>Status: {plan.status}</span>
                {plan.deadline && (
                    <span>Deadline: {new Date(plan.deadline).toLocaleDateString()}</span>
                )}
            </div>
            <div className="test-plan-actions">
                <button className="btn btn-primary">View</button>
                <button className="btn btn-secondary">Export</button>
            </div>
        </div>
    );
};

export default TestPlans;