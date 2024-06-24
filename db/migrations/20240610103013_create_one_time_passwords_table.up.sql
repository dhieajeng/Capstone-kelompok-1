BEGIN;

CREATE TABLE IF NOT EXISTS public.one_time_passwords (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    otp_code VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    expires_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

COMMIT;