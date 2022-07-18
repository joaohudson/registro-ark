-- Forma como o dino se locomove
CREATE TABLE locomotion(
    id_locomotion BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name_locomotion VARCHAR(80) NOT NULL
);
-- Região onde se pode encontrar o dino
CREATE TABLE region(
    id_region BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name_region VARCHAR(80) NOT NULL
);
-- Tipo de Alimentação
CREATE TABLE food(
    id_food BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name_food VARCHAR(80) NOT NULL
);
-- Dados dos usuários adms
CREATE TABLE adm(
    id_adm BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name_adm VARCHAR(80) NOT NULL,
    password_adm VARCHAR(80) NOT NULL,
    permission_manager_dino BOOLEAN NOT NULL,
    permission_manager_category BOOLEAN NOT NULL,
    permission_manager_adm BOOLEAN NOT NULL,
    main_adm BOOLEAN NOT NULL
);
-- Dados gerais do dinossauro
CREATE TABLE dino(
    id_dino BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name_dino VARCHAR(80) NOT NULL,
    id_region BIGINT REFERENCES region(id_region),
    id_food BIGINT REFERENCES food(id_food),
    id_locomotion BIGINT REFERENCES locomotion(id_locomotion),
    id_adm BIGINT REFERENCES adm(id_adm),
    dt_creation TIMESTAMP NOT NULL,
    utility_dino VARCHAR(300) NOT NULL,
    training_dino TEXT NOT NULL
);

INSERT INTO adm(name_adm, password_adm, permission_manager_dino, permission_manager_category, permission_manager_adm, main_adm)
VALUES('Admin', 'Admin', true, true, true, true);