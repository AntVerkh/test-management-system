import React from 'react';
import { useParams } from 'react-router-dom';
import ExportButton from '../components/ExportButton';
import { useTestPlan } from '../hooks/useTestPlan';

const TestPlanDetail = () => {
    const { id } = useParams();
    const { data: testPlan, isLoading } = useTestPlan(id);

    if (isLoading) return <div>Loading...</div>;
    if (!testPlan) return <div>Test plan not found</div>;

    return (
        <div className="container mx-auto px-4 py-8">
            <div className="flex justify-between items-center mb-6">
                <div>
                    <h1 className="text-3xl font-bold text-gray-900">{testPlan.name}</h1>
                    <p className="text-gray-600 mt-2">{testPlan.description}</p>
                </div>
                <div className="flex gap-2">
                    <ExportButton
                        entityType="test_plan"
                        entityId={testPlan.id}
                        entityName={testPlan.name}
                        variant="primary"
                    />
                    {/* Other action buttons */}
                </div>
            </div>

            {/* Test plan content */}
            <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
                <div className="lg:col-span-2">
                    {/* Test cases and checklists */}
                </div>
                <div className="space-y-4">
                    <div className="bg-white rounded-lg border p-4">
                        <h3 className="font-semibold mb-2">Quick Export</h3>
                        <div className="space-y-2">
                            <ExportButton
                                entityType="test_plan"
                                entityId={testPlan.id}
                                entityName={testPlan.name}
                                variant="outline"
                                size="sm"
                            />
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default TestPlanDetail;