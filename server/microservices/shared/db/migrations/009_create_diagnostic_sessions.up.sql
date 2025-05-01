CREATE TYPE schedule_status AS ENUM (
  'SCHEDULED',
  'COMPLETED',
  'CANCELED'
);

CREATE TYPE schedule_type AS ENUM (
  'BLOOD_TEST',
  'URINE_TEST',
  'X_RAY',
  'MRI',
  'CT_SCAN',
  'ULTRASOUND',
  'ECG',
  'COVID_TEST',
  'DNA_TEST',
  'ALLERGY_TEST',
  'GENETIC_TEST',
  'OTHER');

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE diagnostic_schedules (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL,
  diagnostic_centre_id UUID NOT NULL,
  date TIMESTAMP WITH TIME ZONE NOT NULL,
  time TIMESTAMP WITH TIME ZONE NOT NULL,
  test_type schedule_type NOT NULL,
  status schedule_status NOT NULL,
  notes TEXT,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (diagnostic_centre_id) REFERENCES diagnostics(id) ON DELETE CASCADE);
