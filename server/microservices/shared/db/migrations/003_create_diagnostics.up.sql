CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS diagnostics (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  diagnostic_centre_name VARCHAR(255) NOT NULL,
  latitude DOUBLE PRECISION NULL,
  longitude DOUBLE PRECISION NULL,
  address JSONB NULL,
  contact JSONB NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  -- Create a unique index on the combination of user_id and diagnostic_centre_name
  -- to ensure that a user cannot have two diagnostics with the same name
  UNIQUE (user_id, diagnostic_centre_name),
  -- Add a check constraint to ensure that latitude is between -90 and 90
  CHECK (latitude IS NULL OR (latitude >= -90 AND latitude <= 90)),
  -- Add a check constraint to ensure that longitude is between -180 and 180
  CHECK (longitude IS NULL OR (longitude >= -180 AND longitude <= 180)),
  -- Add a check constraint to ensure that the address field is a valid JSON object
  CHECK (address IS NULL OR jsonb_typeof(address) = 'object'),
  -- Add a check constraint to ensure that the contact field is a valid JSON object
  CHECK (contact IS NULL OR jsonb_typeof(contact) = 'object'),
  -- Add a check constraint to ensure that the created_at field is not in the future
  CHECK (created_at <= NOW()),
  -- Add a check constraint to ensure that the updated_at field is not in the future
  CHECK (updated_at <= NOW()),
  -- Add a check constraint to ensure that the updated_at field is greater than or equal to the created_at field
  CHECK (updated_at >= created_at),
  -- Add a check constraint to ensure that the user_id field is not null
  CHECK (user_id IS NOT NULL),
  -- Add a check constraint to ensure that the diagnostic_centre_name field is not null
  CHECK (diagnostic_centre_name IS NOT NULL)
  -- Add a check constraint to ensure that the latitude field is not null
);

CREATE INDEX idx_diagnostics_user_id ON diagnostics(user_id);
