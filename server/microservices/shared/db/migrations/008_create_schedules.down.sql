-- Drop indexes
DROP INDEX IF EXISTS idx_diagnostics_schedules_id_user_id;
DROP INDEX IF EXISTS idx_diagnostic_schedules_id_centre_user;
DROP INDEX IF EXISTS idx_diagnostic_schedules_centre_date;
DROP INDEX IF EXISTS idx_diagnostic_schedules_status;

-- Drop ENUM types
DROP TYPE IF EXISTS schedule_status CASCADE;
DROP TYPE IF EXISTS schedule_type CASCADE;

-- Drop Extension
DROP EXTENSION IF EXISTS "uuid-ossp" CASCADE;
-- Drop table
DROP TABLE IF EXISTS diagnostic_schedules CASCADE;
