-- Drop the TASK table if it exists
DROP TABLE IF EXISTS TASKS;

-- Create the TASK table
CREATE TABLE IF NOT EXISTS TASKS (
    drn_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(100) NOT NULL,
    description TEXT,
    status VARCHAR(20) NOT NULL DEFAULT 'Pending',
    due_date TIMESTAMP
);

-- Insert dummy tasks into the TASKS table
INSERT INTO TASKS (title, description, status, due_date) VALUES
    ('Task 1', 'Description for task 1', 'Pending', '2024-08-31 12:00:00'),
    ('Task 2', 'Description for task 2', 'In Progress', '2024-09-01 15:00:00'),
    ('Task 3', 'Description for task 3', 'Completed', '2024-08-15 09:00:00'),
    ('Task 4', 'Description for task 4', 'Pending', '2024-09-10 18:00:00'),
    ('Task 5', 'Description for task 5', 'Pending', NULL); -- Due date is NULL