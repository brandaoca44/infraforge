CREATE INDEX IF NOT EXISTS idx_services_status ON services(status);
CREATE INDEX IF NOT EXISTS idx_services_environment ON services(environment);
CREATE INDEX IF NOT EXISTS idx_services_created_at ON services(created_at);