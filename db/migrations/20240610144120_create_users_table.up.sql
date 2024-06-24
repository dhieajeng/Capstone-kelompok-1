BEGIN;

CREATE TYPE gender AS ENUM ('F', 'M');

CREATE TABLE IF NOT EXISTS public.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(255) NOT NULL,
    phone VARCHAR(16) UNIQUE,
    address TEXT,
    avatar VARCHAR(255),
    birthdate DATE,
    gender gender,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

COMMIT;