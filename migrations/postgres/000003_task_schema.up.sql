CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    description TEXT,
    priority VARCHAR NOT NULL CHECK (priority IN ('low', 'medium', 'high')),
    status VARCHAR NOT NULL CHECK (status IN ('active', 'in_progress', 'completed')),
    assignee_id UUID NOT NULL REFERENCES users(id),
    project_id UUID NOT NULL REFERENCES projects(id),
    creation_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completion_date TIMESTAMP
);
