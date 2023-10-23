-- +goose Up
-- +goose StatementBegin
CREATE DATABASE aphrodite
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

COMMENT ON DATABASE aphrodite
    IS 'Aphrodite Monolite';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE IF EXISTS aphrodite;
-- +goose StatementEnd
