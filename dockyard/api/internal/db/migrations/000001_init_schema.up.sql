CREATE TABLE images (
    id           BIGSERIAL PRIMARY KEY,
    name         VARCHAR(255) NOT NULL,
    registry     VARCHAR(255) NOT NULL DEFAULT 'docker.io',
    description  TEXT,
    source_url   VARCHAR(500),
    logo_url     VARCHAR(500),
    created_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at   TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (registry, name)
);

CREATE INDEX idx_images_name ON images (name);
