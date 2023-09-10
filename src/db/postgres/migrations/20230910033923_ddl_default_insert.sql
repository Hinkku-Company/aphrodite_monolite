-- +goose Up
-- +goose StatementBegin
-- DEFAULT
INSERT INTO hk.type_user
(id, "name")
VALUES
('c7d97580-21d9-4bb4-97a1-fe4ed7180fc3', 'cliente'),
('f8ff19a1-6548-4270-b0e5-5f3930499bd1', 'empleado'),
('dfc095ee-8ff1-448c-b975-1f800023ffce', 'owner');

INSERT INTO hk.access_rol
(id, "name")
VALUES
('1a4a3003-41d6-4b29-8bd4-718f192bc8ee', 'super'),
('c559aa06-12e6-4eea-a39d-9eb40e601c00', 'administrador_local'),
('f6c26c2b-377d-41ce-a526-b6f577d78222', 'cliente'),
('a1472917-f992-4804-b20a-301a7c4b0a2e', 'empleado'),
('4481c02c-776b-4d26-a65d-2c41347beeaa', 'publicidad'),
('9cc30d59-9449-44e5-be8e-f18adef8b60f', 'anonimo');

INSERT INTO hk.users_status
(id, "name")
VALUES
('96322f2f-c7b5-4c9c-9574-0d61d1842927', 'active'),
('ff86681e-67a2-4df7-88ec-97df1e4264c9', 'pending'),
('3654e50a-952b-4c75-9aa8-0af288f31c63', 'block'),
('7c39a363-4c54-413a-9e95-676bf09fc2ac', 'delete');

INSERT INTO hk.type_status_reservation
(id, "name")
VALUES
('9b86b2b3-3a6e-4ec5-9896-94094949cf27', 'en_curso'),
('7f6db2d3-d04c-4ee3-960b-c96da2e62740', 'movido'),
('3d4da924-7e0c-4229-bb84-1b61a29050a3', 'cancelado_usuario'),
('7c69eb80-fb47-4ff0-b3e8-90a653e4c59e', 'cancelado_negocio'),
('37a3fdc9-e29e-46a9-984a-e59b2fe7a4bc', 'terminado'),
('1c3d8efb-fca0-4d6e-b167-bfcd934a62f2', 'expirado_automatico');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE hk.type_user CASCADE;
TRUNCATE TABLE hk.access_rol CASCADE;
TRUNCATE TABLE hk.users_status CASCADE;
TRUNCATE TABLE hk.type_status_reservation CASCADE;
-- +goose StatementEnd
