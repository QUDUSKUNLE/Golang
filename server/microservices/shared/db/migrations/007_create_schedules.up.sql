-- Create ENUM type for schedule status
CREATE TYPE schedule_status AS ENUM (
  'SCHEDULED',
  'COMPLETED',
  'CANCELED'
);

-- Create ENUM type for schedule type
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
  'OTHER'
);

-- Ensure the UUID extension is available
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Create the diagnostic_schedules table
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'diagnostic_schedules') THEN
        CREATE TABLE diagnostic_schedules (
            id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
            user_id UUID NOT NULL,
            diagnostic_centre_id UUID NOT NULL,
            date TIMESTAMP WITH TIME ZONE NOT NULL,
            time TIMESTAMP WITH TIME ZONE NOT NULL,
            test_type schedule_type NOT NULL DEFAULT 'OTHER',
            status schedule_status NOT NULL DEFAULT 'SCHEDULED',
            notes TEXT,
            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
            FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
            FOREIGN KEY (diagnostic_centre_id) REFERENCES diagnostics(id) ON DELETE CASCADE,
            UNIQUE (user_id, diagnostic_centre_id, date, time), -- Prevent duplicate schedules
            CHECK (date >= NOW()) -- Ensure date is in the future
        );
    END IF;
END $$;

-- Create indexes for faster lookups
CREATE INDEX IF NOT EXISTS idx_diagnostics_schedules_id_user_id
ON diagnostic_schedules (id, user_id);

CREATE INDEX IF NOT EXISTS idx_diagnostic_schedules_id_centre_user
ON diagnostic_schedules (id, diagnostic_centre_id, user_id);

CREATE INDEX IF NOT EXISTS idx_diagnostic_schedules_centre_date
ON diagnostic_schedules (diagnostic_centre_id, date);

CREATE INDEX IF NOT EXISTS idx_diagnostic_schedules_status
ON diagnostic_schedules (status);

CREATE INDEX IF NOT EXISTS idx_diagnostic_schedules_centre_status_date
ON diagnostic_schedules (diagnostic_centre_id, status, date DESC, time DESC);
