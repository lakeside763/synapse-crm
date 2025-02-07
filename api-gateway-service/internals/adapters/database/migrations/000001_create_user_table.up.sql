CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SEQUENCE IF NOT EXISTS user_serial_id START 1;

CREATE TYPE user_role AS ENUM ('superadmin', 'admin', 'sales', 'customer');

CREATE TABLE "users" (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    serial_id INT NOT NULL DEFAULT nextval('user_serial_id'),
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role user_role NOT NULL DEFAULT 'admin',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX user_serial_id_index ON "users" (serial_id);
CREATE INDEX user_full_name_index ON "users" (first_name, last_name);
CREATE INDEX user_email_index ON "users" (email);
CREATE INDEX user_role_index ON "users" (role);
CREATE INDEX user_created_at_index ON "users" (created_at);
CREATE INDEX user_active_index ON "users" (email) WHERE deleted_at IS NULL;

