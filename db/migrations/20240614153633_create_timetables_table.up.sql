BEGIN;

CREATE TABLE IF NOT EXISTS public.timetables (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    event_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    start TIMESTAMP NOT NULL,
    "end" TIMESTAMP NOT NULL,
    description TEXT,
    stock INT NOT NULL DEFAULT 0,
    price INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

COMMIT;