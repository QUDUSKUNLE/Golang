CREATE TYPE user_enum AS ENUM (
  'USER',
  'DIAGNOSTIC_CENTRE',
  'HOSPITAL',
  'ADMIN'
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Check if the table exists before creating it
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') THEN
        CREATE TABLE users (
            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
            email VARCHAR(255) UNIQUE,
            nin VARCHAR(11) NULL,
            password VARCHAR(255) NOT NULL,
            user_type user_enum NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
        );
    END IF;
END $$;
