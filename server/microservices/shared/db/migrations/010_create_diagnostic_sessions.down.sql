-- Drop indexes
DROP INDEX IF EXISTS idx_diagnostics_schedules_id_user_id;
DROP INDEX IF EXISTS idx_diagnostic_schedules_id_centre_user;
DROP INDEX IF EXISTS idx_diagnostic_schedules_centre_date;
DROP INDEX IF EXISTS idx_diagnostic_schedules_status;

-- Drop table
DROP TABLE IF EXISTS diagnostic_schedules CASCADE;

-- Drop ENUM types
DROP TYPE IF EXISTS schedule_status;
DROP TYPE IF EXISTS schedule_type;

-- Drop Extension
DROP EXTENSION IF EXISTS "uuid-ossp";
