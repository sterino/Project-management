CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT GEN_RANDOM_UUID(),
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    priority VARCHAR NOT NULL CHECK (priority IN ('low', 'medium', 'high')),
    status VARCHAR NOT NULL CHECK (status IN ('active', 'in_progress', 'completed')),
    user_id UUID NOT NULL REFERENCES users(id),
    project_id UUID NOT NULL REFERENCES projects(id),
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL
);
