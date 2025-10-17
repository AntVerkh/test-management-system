import React, { useState } from 'react';
import { downloadExport } from '../services/api';

const ExportButton = ({
                          entityType,
                          entityId,
                          entityName,
                          variant = "secondary",
                          size = "md"
                      }) => {
    const [isExporting, setIsExporting] = useState(false);
    const [showOptions, setShowOptions] = useState(false);

    const handleExport = async (includeHistory = false, includeComments = false) => {
        setIsExporting(true);
        setShowOptions(false);

        try {
            await downloadExport(entityType, entityId, {
                includeHistory,
                includeComments
            });
        } catch (error) {
            console.error('Export failed:', error);
            alert('Export failed. Please try again.');
        } finally {
            setIsExporting(false);
        }
    };

    const sizeClasses = {
        sm: 'px-3 py-1 text-sm',
        md: 'px-4 py-2',
        lg: 'px-6 py-3 text-lg'
    };

    const variantClasses = {
        primary: 'bg-blue-600 hover:bg-blue-700 text-white',
        secondary: 'bg-gray-600 hover:bg-gray-700 text-white',
        outline: 'border border-gray-600 hover:bg-gray-100 text-gray-700'
    };

    return (
        <div className="relative inline-block">
            <button
                onClick={() => setShowOptions(!showOptions)}
                disabled={isExporting}
                className={`
          ${sizeClasses[size]} 
          ${variantClasses[variant]}
          rounded-md font-medium transition-colors
          disabled:opacity-50 disabled:cursor-not-allowed
          flex items-center gap-2
        `}
            >
                {isExporting ? (
                    <>
                        <svg className="animate-spin h-4 w-4" viewBox="0 0 24 24">
                            <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" fill="none"/>
                            <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"/>
                        </svg>
                        Exporting...
                    </>
                ) : (
                    <>
                        <svg className="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 10v6m0 0l-3-3m3 3l3-3m2 8H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                        </svg>
                        Export
                    </>
                )}
            </button>

            {showOptions && (
                <div className="absolute right-0 mt-2 w-64 bg-white rounded-md shadow-lg border z-10">
                    <div className="py-1">
                        <div className="px-4 py-2 text-sm text-gray-700 border-b">
                            Export Options for {entityName}
                        </div>

                        <button
                            onClick={() => handleExport(false, false)}
                            className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                        >
                            📄 Basic Export
                        </button>

                        <button
                            onClick={() => handleExport(true, false)}
                            className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                        >
                            📋 With History
                        </button>

                        <button
                            onClick={() => handleExport(false, true)}
                            className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                        >
                            💬 With Comments
                        </button>

                        <button
                            onClick={() => handleExport(true, true)}
                            className="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                        >
                            📊 Complete Export
                        </button>
                    </div>
                </div>
            )}
        </div>
    );
};

export default ExportButton;