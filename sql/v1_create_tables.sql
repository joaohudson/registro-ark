CREATE TABLE dino(
    id_dino BIGINT GENERATED ALWAYS AS IDENTITY,
    name_dino VARCHAR(80) NOT NULL,
    food_dino VARCHAR(80) NOT NULL,
    locomotion_dino VARCHAR(30) NOT NULL,
    region_dino VARCHAR(80) NOT NULL,
    utility_dino VARCHAR(300) NOT NULL,
    training_dino TEXT NOT NULL
);