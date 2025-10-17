import React from 'react';

const Dashboard = () => {
    return (
        <div className="dashboard">
            <h1>Dashboard</h1>
            <div className="dashboard-stats">
                <div className="stat-card">
                    <h3>Test Plans</h3>
                    <p>0</p>
                </div>
                <div className="stat-card">
                    <h3>Test Cases</h3>
                    <p>0</p>
                </div>
                <div className="stat-card">
                    <h3>Test Runs</h3>
                    <p>0</p>
                </div>
            </div>

            <div className="recent-activity">
                <h2>Recent Activity</h2>
                <p>No recent activity</p>
            </div>
        </div>
    );
};

export default Dashboard;