import api from './api';

export const downloadExport = async (entityType, entityId, options = {}) => {
    const { includeHistory = false, includeComments = false } = options;

    const response = await api.post('/export', {
        entity_type: entityType,
        entity_id: entityId,
        format: 'markdown',
        include_history: includeHistory,
        include_comments: includeComments
    }, {
        responseType: 'blob'
    });

    // Create blob and download
    const blob = new Blob([response.data], { type: 'text/markdown' });
    const url = window.URL.createObjectURL(blob);
    const link = document.createElement('a');
    link.href = url;

    // Extract filename from content-disposition header
    const contentDisposition = response.headers['content-disposition'];
    let filename = `${entityType}_${entityId}.md`;
    if (contentDisposition) {
        const filenameMatch = contentDisposition.match(/filename="(.+)"/);
        if (filenameMatch) {
            filename = filenameMatch[1];
        }
    }

    link.download = filename;
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);
    window.URL.revokeObjectURL(url);
};

// Convenience functions for specific entity types
export const exportTestPlan = (planId, options) =>
    downloadExport('test_plan', planId, options);

export const exportTestCase = (testCaseId, options) =>
    downloadExport('test_case', testCaseId, options);

export const exportChecklist = (checklistId, options) =>
    downloadExport('checklist', checklistId, options);

export const exportTestStrategy = (strategyId, options) =>
    downloadExport('test_strategy', strategyId, options);

export const exportTestRun = (testRunId, options) =>
    downloadExport('test_run', testRunId, options);