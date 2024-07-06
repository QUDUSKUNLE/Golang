-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
SET TIMEZONE="Africa/Lagos";

-- Create table users
CREATE TABLE IF NOT EXISTS users (
  id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  email VARCHAR (255) UNIQUE NOT NULL,
  pass VARCHAR (255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  user_type VARCHAR (10) NOT NULL
);


-- Add indexes
-- CREATE INDEX active_users ON users (title) WHERE book_status = 1;

-- Create shippings table
CREATE TABLE IF NOT EXISTS shippings (
  id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  user_id UUID,
  product_description VARCHAR(500) NOT NULL,
  pick_up_address JSON NOT NULL,
  delivery_address JSON NOT NULL,
  product_type VARCHAR (255) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  CONSTRAINT fk_users
    FOREIGN KEY(user_id)
      REFERENCES users(id)
      ON DELETE CASCADE
);
