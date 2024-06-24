-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="Africa/Lagos";

-- Create table books
CREATE TABLE IF NOT EXISTS books (
  id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_At TIMESTAMP NULL,
  title VARCHAR (255) NOT NULL,
  author VARCHAR (255) NOT NULL,
  book_status INT NOT NULL,
  book_attrs JSONB NOT NULL
);


-- Add indexes
CREATE INDEX active_books ON books (title) WHERE book_status = 1;
