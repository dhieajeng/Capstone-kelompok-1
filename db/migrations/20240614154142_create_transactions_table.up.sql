BEGIN;

CREATE TABLE IF NOT EXISTS public.transactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    invoice VARCHAR(255) NOT NULL,
    grand_total INT NOT NULL DEFAULT 0,
    snap_token VARCHAR(255),
    status VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

COMMIT;