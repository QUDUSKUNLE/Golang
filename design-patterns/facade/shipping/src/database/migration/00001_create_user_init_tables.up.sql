-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="Africa/Lagos";

-- Create table books
CREATE TABLE IF NOT EXISTS users (
  id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  email VARCHAR (255) UNIQUE NOT NULL,
  pass VARCHAR (255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL,
);


-- Add indexes
-- CREATE INDEX active_users ON users (title) WHERE book_status = 1;
