DROP TABLE IF EXISTS users;

DROP SEQUENCE IF EXISTS user_serial_id;

DROP TYPE IF EXISTS user_role;

DROP EXTENSION IF EXISTS "uuid-ossp";

DROP INDEX IF EXISTS user_serial_id_index;

DROP INDEX IF EXISTS user_full_name_index;

DROP INDEX IF EXISTS user_email_index;

DROP INDEX IF EXISTS user_role_index;

DROP INDEX IF EXISTS user_created_at_index;

DROP INDEX IF EXISTS user_active_index;