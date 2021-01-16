-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE file_collection
(
    id         uuid PRIMARY KEY,
    status     VARCHAR(25) NOT NULL,
    system_id  VARCHAR(10) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE file
(
    id                 uuid PRIMARY KEY,
    file_collection_id uuid          NOT NULL,
    file_name          VARCHAR(1024) NOT NULL,
    file_type          VARCHAR(4),
    file_usage         VARCHAR(10)   NOT NULL,
    source_bucket      VARCHAR(63),
    source_key         VARCHAR(1024),
    status             VARCHAR(25)   NOT NULL,
    etag               VARCHAR(34),
    created_at         TIMESTAMPTZ   NOT NULL,
    updated_at         TIMESTAMPTZ   NOT NULL,
    FOREIGN KEY (file_collection_id) REFERENCES file_collection (id)
);

CREATE TABLE file_task
(
    id                uuid PRIMARY KEY,
    file_id           uuid        NOT NULL,
    task_type         VARCHAR(25) NOT NULL,
    status            VARCHAR(25) NOT NULL,
    source_meta       TEXT        NULL,
    result_meta       TEXT        NULL,
    error_message     VARCHAR(64),
    task_started_at   TIMESTAMPTZ NOT NULL,
    task_completed_at TIMESTAMPTZ,
    FOREIGN KEY (file_id) REFERENCES file (id)
);

CREATE TABLE file_collection_task
(
    id                 uuid PRIMARY KEY,
    file_collection_id uuid        NOT NULL,
    task_type          VARCHAR(25) NOT NULL,
    status             VARCHAR(25) NOT NULL,
    source_meta        TEXT        NULL,
    result_meta        TEXT        NULL,
    error_message      VARCHAR(64),
    task_started_at    TIMESTAMPTZ NOT NULL,
    task_completed_at  TIMESTAMPTZ,
    FOREIGN KEY (file_collection_id) REFERENCES file_collection (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.