BEGIN;

CREATE TABLE IF NOT EXISTS public.events (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    location_id INT NOT NULL,
    category_id INT NOT NULL,
    topic_id INT NOT NULL,
    name VARCHAR(255) NOT NULL,
    start TIMESTAMP NOT NULL,
    "end" TIMESTAMP NOT NULL,
    address TEXT NOT NULL,
    address_link VARCHAR(255),
    organizer VARCHAR(255) NOT NULL,
    organizer_logo VARCHAR(255),
    cover VARCHAR(255),
    description TEXT NOT NULL,
    term_and_condition TEXT NOT NULL,
    is_paid BOOLEAN NOT NULL DEFAULT true,
    is_public BOOLEAN NOT NULL DEFAULT true,
    is_approved BOOLEAN NOT NULL DEFAULT false,
    approved_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

COMMIT;