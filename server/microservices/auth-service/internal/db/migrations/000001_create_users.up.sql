CREATE TYPE user_enum AS ENUM ('user', 'organization', 'carrier', 'admin');

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY,
  email TEXT UNIQUE,
  nin text UNIQUE,
  password TEXT NOT NULL,
  user_type user_enum,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
