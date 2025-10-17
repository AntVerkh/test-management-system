-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role VARCHAR(20) NOT NULL DEFAULT 'user',
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Projects table
CREATE TABLE projects (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          created_by UUID REFERENCES users(id),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Test Strategies table
CREATE TABLE test_strategies (
                                 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                 project_id UUID NOT NULL REFERENCES projects(id),
                                 name VARCHAR(255) NOT NULL,
                                 description TEXT,
                                 content TEXT,
                                 created_by UUID REFERENCES users(id),
                                 created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Checklists table
CREATE TABLE checklists (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            project_id UUID NOT NULL REFERENCES projects(id),
                            name VARCHAR(255) NOT NULL,
                            description TEXT,
                            created_by UUID REFERENCES users(id),
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Checklist Items table
CREATE TABLE checklist_items (
                                 id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                 checklist_id UUID NOT NULL REFERENCES checklists(id),
                                 description TEXT NOT NULL,
                                 expected_result TEXT,
                                 "order" INTEGER NOT NULL,
                                 created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Test Cases table
CREATE TABLE test_cases (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            project_id UUID NOT NULL REFERENCES projects(id),
                            title VARCHAR(255) NOT NULL,
                            description TEXT,
                            pre_steps TEXT,
                            expected_result TEXT,
                            created_by UUID REFERENCES users(id),
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Test Steps table
CREATE TABLE test_steps (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            test_case_id UUID NOT NULL REFERENCES test_cases(id),
                            description TEXT NOT NULL,
                            expected_result TEXT,
                            "order" INTEGER NOT NULL,
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Test Plans table
CREATE TABLE test_plans (
                            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                            project_id UUID NOT NULL REFERENCES projects(id),
                            name VARCHAR(255) NOT NULL,
                            description TEXT,
                            deadline TIMESTAMP WITH TIME ZONE,
                            status VARCHAR(50) DEFAULT 'draft',
                            created_by UUID REFERENCES users(id),
                            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Test Plan - Checklists junction table
CREATE TABLE test_plan_checklists (
                                      test_plan_id UUID NOT NULL REFERENCES test_plans(id),
                                      checklist_id UUID NOT NULL REFERENCES checklists(id),
                                      PRIMARY KEY (test_plan_id, checklist_id)
);

-- Test Plan - Test Cases junction table
CREATE TABLE test_plan_cases (
                                 test_plan_id UUID NOT NULL REFERENCES test_plans(id),
                                 test_case_id UUID NOT NULL REFERENCES test_cases(id),
                                 PRIMARY KEY (test_plan_id, test_case_id)
);

-- Test Runs table
CREATE TABLE test_runs (
                           id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                           test_plan_id UUID NOT NULL REFERENCES test_plans(id),
                           name VARCHAR(255) NOT NULL,
                           started_by UUID REFERENCES users(id),
                           started_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                           completed_at TIMESTAMP WITH TIME ZONE
);

-- Test Results table
CREATE TABLE test_results (
                              id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                              test_run_id UUID NOT NULL REFERENCES test_runs(id),
                              test_case_id UUID REFERENCES test_cases(id),
                              checklist_item_id UUID REFERENCES checklist_items(id),
                              status VARCHAR(50) NOT NULL,
                              comments TEXT,
                              executed_by UUID REFERENCES users(id),
                              executed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Attachments table
CREATE TABLE attachments (
                             id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             test_case_id UUID NOT NULL REFERENCES test_cases(id),
                             file_name VARCHAR(255) NOT NULL,
                             file_path VARCHAR(500) NOT NULL,
                             file_size BIGINT,
                             mime_type VARCHAR(100),
                             uploaded_by UUID REFERENCES users(id),
                             uploaded_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Comments table
CREATE TABLE comments (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          entity_id UUID NOT NULL,
                          entity_type VARCHAR(50) NOT NULL,
                          content TEXT NOT NULL,
                          created_by UUID REFERENCES users(id),
                          created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- History table
CREATE TABLE history (
                         id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                         entity_id UUID NOT NULL,
                         entity_type VARCHAR(50) NOT NULL,
                         action VARCHAR(50) NOT NULL,
                         changes JSONB,
                         changed_by UUID REFERENCES users(id),
                         changed_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Indexes for better performance
CREATE INDEX idx_test_plans_project_id ON test_plans(project_id);
CREATE INDEX idx_test_cases_project_id ON test_cases(project_id);
CREATE INDEX idx_checklists_project_id ON checklists(project_id);
CREATE INDEX idx_test_runs_test_plan_id ON test_runs(test_plan_id);
CREATE INDEX idx_test_results_test_run_id ON test_results(test_run_id);
CREATE INDEX idx_comments_entity ON comments(entity_id, entity_type);
CREATE INDEX idx_history_entity ON history(entity_id, entity_type);