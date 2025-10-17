import React, { useState } from 'react';
import { useForm } from 'react-hook-form';

const TestCases = () => {
    const [showCreateForm, setShowCreateForm] = useState(false);

    return (
        <div className="test-cases">
            <div className="page-header">
                <h1>Test Cases</h1>
                <button
                    className="btn btn-primary"
                    onClick={() => setShowCreateForm(true)}
                >
                    Create Test Case
                </button>
            </div>

            {showCreateForm && (
                <CreateTestCaseForm
                    onCancel={() => setShowCreateForm(false)}
                />
            )}

            <div className="test-cases-list">
                <div className="empty-state">
                    <h3>No test cases yet</h3>
                    <p>Create your first test case to get started</p>
                </div>
            </div>
        </div>
    );
};

const CreateTestCaseForm = ({ onCancel }) => {
    const { register, handleSubmit, formState: { errors, isSubmitting } } = useForm();

    const onSubmit = async (data) => {
        // TODO: Implement create test case
        console.log('Create test case:', data);
        onCancel();
    };

    const [steps, setSteps] = useState([{ description: '', expectedResult: '' }]);

    const addStep = () => {
        setSteps([...steps, { description: '', expectedResult: '' }]);
    };

    const removeStep = (index) => {
        setSteps(steps.filter((_, i) => i !== index));
    };

    const updateStep = (index, field, value) => {
        const newSteps = [...steps];
        newSteps[index][field] = value;
        setSteps(newSteps);
    };

    return (
        <div className="card">
            <h3>Create New Test Case</h3>
            <form onSubmit={handleSubmit(onSubmit)}>
                <div className="form-group">
                    <label>Title *</label>
                    <input
                        type="text"
                        {...register('title', { required: 'Title is required' })}
                    />
                    {errors.title && <span className="error">{errors.title.message}</span>}
                </div>

                <div className="form-group">
                    <label>Description</label>
                    <textarea
                        {...register('description')}
                        rows="3"
                    />
                </div>

                <div className="form-group">
                    <label>Pre-Steps</label>
                    <textarea
                        {...register('preSteps')}
                        rows="3"
                        placeholder="Any setup steps required before the test..."
                    />
                </div>

                <div className="form-group">
                    <label>Test Steps</label>
                    {steps.map((step, index) => (
                        <div key={index} className="test-step">
                            <div className="step-header">
                                <h4>Step {index + 1}</h4>
                                {steps.length > 1 && (
                                    <button
                                        type="button"
                                        onClick={() => removeStep(index)}
                                        className="btn btn-danger"
                                    >
                                        Remove
                                    </button>
                                )}
                            </div>
                            <div className="form-group">
                                <label>Description *</label>
                                <input
                                    type="text"
                                    value={step.description}
                                    onChange={(e) => updateStep(index, 'description', e.target.value)}
                                    required
                                />
                            </div>
                            <div className="form-group">
                                <label>Expected Result</label>
                                <input
                                    type="text"
                                    value={step.expectedResult}
                                    onChange={(e) => updateStep(index, 'expectedResult', e.target.value)}
                                />
                            </div>
                        </div>
                    ))}
                    <button type="button" onClick={addStep} className="btn btn-secondary">
                        Add Step
                    </button>
                </div>

                <div className="form-group">
                    <label>Expected Result</label>
                    <textarea
                        {...register('expectedResult')}
                        rows="3"
                    />
                </div>

                <div className="form-actions">
                    <button type="button" onClick={onCancel} className="btn btn-secondary">
                        Cancel
                    </button>
                    <button type="submit" disabled={isSubmitting} className="btn btn-primary">
                        {isSubmitting ? 'Creating...' : 'Create Test Case'}
                    </button>
                </div>
            </form>
        </div>
    );
};

export default TestCases;