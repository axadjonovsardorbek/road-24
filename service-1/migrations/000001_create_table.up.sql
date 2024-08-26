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
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'categories') THEN
        CREATE TYPE categories AS ENUM(
            'a', 
            'b',
            'c',
            'd',
            'e'
        );
    END IF;
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'fine_status') THEN
        CREATE TYPE fine_status AS ENUM(
            'unpaid', 
            'pending',
            'paid'
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
    id UUID NOT NULL,
    name VARCHAR(255) PRIMARY KEY NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS cars (
    id UUID NOT NULL,
    staf_id UUID REFERENCES users(id) NOT NULL,
    type VARCHAR(64) REFERENCES car_types(name) NOT NULL,
    model VARCHAR(64) NOT NULL,
    color VARCHAR(64) NOT NULL,
    year INTEGER NOT NULL CHECK (year > 0),
    body_number VARCHAR(255) NOT NULL,
    horse_power VARCHAR(255) NOT NULL,
    number VARCHAR(10) NOT NULL UNIQUE,
    technical_passport VARCHAR(11) NOT NULL UNIQUE,
    image_url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0,
    PRIMARY KEY (id, technical_passport, number),
    FOREIGN KEY (type) REFERENCES car_types(name)
);

CREATE TABLE IF NOT EXISTS service_types (
    id UUID NOT NULL,
    name VARCHAR(128) PRIMARY KEY NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS services (
    id UUID PRIMARY KEY,
    type VARCHAR(128) REFERENCES service_types(name) NOT NULL,
    name VARCHAR(128) NOT NULL,
    sertificat_number VARCHAR(64) NOT NULL UNIQUE,
    owner_name VARCHAR(128) NOT NULL,
    address VARCHAR(128) NOT NULL,
    phone_number VARCHAR(15) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS car_services (
    id UUID PRIMARY KEY,
    car_number VARCHAR(10) REFERENCES cars(number) NOT NULL,
    service_type VARCHAR(64) REFERENCES service_types(name) NOT NULL,
    service_date TIMESTAMP WITH TIME ZONE NOT NULL,
    exp_date TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS driver_licences (
    id UUID PRIMARY KEY,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    third_name VARCHAR(64) NOT NULL,
    address VARCHAR(64) NOT NULL,
    got_date TIMESTAMP WITH TIME ZONE NOT NULL,
    exp_date TIMESTAMP WITH TIME ZONE NOT NULL,
    category categories NOT NULL,
    got_address VARCHAR(64) NOT NULL,
    licence_number VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS fines (
    id UUID PRIMARY KEY,
    staf_id UUID REFERENCES users(id) NOT NULL,
    technical_number VARCHAR(64) REFERENCES cars(technical_passport) NOT NULL,
    car_number VARCHAR(15) REFERENCES cars(number) NOT NULL,
    licence_number VARCHAR(20) REFERENCES driver_licences(licence_number) NOT NULL,
    fine_date TIMESTAMP WITH TIME ZONE NOT NULL,
    fine_amount BIGINT NOT NULL,
    payment_date TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    status fine_status NOT NULL DEFAULT 'unpaid',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);

ALTER TABLE users
ADD CONSTRAINT users_unique_tg UNIQUE (username, deleted_at);

ALTER TABLE car_types
ADD CONSTRAINT car_types_unique_tg UNIQUE (name, deleted_at);

ALTER TABLE service_types
ADD CONSTRAINT service_types_unique_tg UNIQUE (name, deleted_at);

ALTER TABLE car_services
ADD CONSTRAINT car_services_unique_tg UNIQUE (service_type, deleted_at); 