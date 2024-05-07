ALTER TABLE request_logs
ADD COLUMN IF NOT EXISTS  access_key text DEFAULT 0;
