CREATE TABLE services (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    environment TEXT NOT NULL,
    status TEXT DEFAULT 'unknown',
    response_time_ms INT DEFAULT 0,
    last_checked_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);