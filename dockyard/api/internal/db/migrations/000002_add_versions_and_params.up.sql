CREATE TABLE image_versions (
    id          BIGSERIAL PRIMARY KEY,
    image_id    BIGINT NOT NULL REFERENCES images(id) ON DELETE CASCADE,
    tag         VARCHAR(255) NOT NULL,
    digest      VARCHAR(255),
    is_latest   BOOLEAN NOT NULL DEFAULT false,
    released_at DATE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    UNIQUE (image_id, tag)
);

CREATE TABLE parameters (
    id               BIGSERIAL PRIMARY KEY,
    image_version_id BIGINT NOT NULL REFERENCES image_versions(id) ON DELETE CASCADE,
    env_var_name     VARCHAR(255) NOT NULL,
    type             VARCHAR(50) NOT NULL DEFAULT 'string'
                     CHECK (type IN ('string', 'int', 'bool', 'enum')),
    default_value    TEXT,
    required         BOOLEAN NOT NULL DEFAULT false,
    description      TEXT,
    UNIQUE (image_version_id, env_var_name)
);

CREATE TABLE ports (
    id               BIGSERIAL PRIMARY KEY,
    image_version_id BIGINT NOT NULL REFERENCES image_versions(id) ON DELETE CASCADE,
    container_port   INTEGER NOT NULL,
    protocol         VARCHAR(10) NOT NULL DEFAULT 'tcp'
                     CHECK (protocol IN ('tcp', 'udp')),
    description      TEXT
);

CREATE TABLE volumes (
    id               BIGSERIAL PRIMARY KEY,
    image_version_id BIGINT NOT NULL REFERENCES image_versions(id) ON DELETE CASCADE,
    path             VARCHAR(500) NOT NULL,
    description      TEXT
);

CREATE INDEX idx_image_versions_image_id ON image_versions (image_id);
CREATE INDEX idx_parameters_image_version_id ON parameters (image_version_id);