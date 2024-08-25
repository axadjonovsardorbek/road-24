DO $$ 
BEGIN 
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'roles') THEN
        CREATE TYPE roles AS ENUM(
            'yhxb', 
            'superadmin',
            'yidxp',
            'service'
        );
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'is_admin_enum') THEN
        CREATE TYPE is_admin_enum AS ENUM(
            'true', 
            'false'
        );
    END IF;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    first_name VARCHAR(128) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    username VARCHAR(64) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role roles NOT NULL,
    is_admin is_admin_enum NOT NULL DEFAULT 'false',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS car_types (
    id UUID PRIMARY KEY,
    name VARCHAR(255) PRIMARY KEY NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS cars (
    id UUID PRIMARY KEY,
    staf_id UUID REFERENCES users(id) NOT NULL,
    type VARCHAR(64) REFERENCES car_types(name) NOT NULL,
    model VARCHAR(64) NOT NULL,
    color VARCHAR(64) NOT NULL,
    year DATE NOT NULL,
    body_number VARCHAR(255) NOT NULL,
    horse_power VARCHAR(255) NOT NULL,
    number VARCHAR(10) NOT NULL,
    technical_passport VARCHAR(11) NOT NULL,
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

ALTER TABLE users
ADD CONSTRAINT users_unique_tg UNIQUE (username, deleted_at);

ALTER TABLE car_types
ADD CONSTRAINT car_types_unique_tg UNIQUE (name, deleted_at);