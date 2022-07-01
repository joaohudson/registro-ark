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
-- Dados gerais do dinossauro
CREATE TABLE dino(
    id_dino BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name_dino VARCHAR(80) NOT NULL,
    id_region BIGINT REFERENCES region(id_region),
    id_food BIGINT REFERENCES food(id_food),
    id_locomotion BIGINT REFERENCES locomotion(id_locomotion),
    utility_dino VARCHAR(300) NOT NULL,
    training_dino TEXT NOT NULL
);