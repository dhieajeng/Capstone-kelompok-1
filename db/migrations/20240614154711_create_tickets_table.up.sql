BEGIN;

CREATE TABLE IF NOT EXISTS public.tickets (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID NOT NULL,
    timetable_id UUID NOT NULL,
    no_ticket VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    personal_no VARCHAR(16),
    birthdate DATE,
    phone VARCHAR(16),
    email VARCHAR(255),
    gender gender,
    price INT,
    is_valid BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

COMMIT;