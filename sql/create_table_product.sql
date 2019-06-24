CREATE TABLE product
(
    id    SERIAL PRIMARY KEY                NOT NULL,
    name  VARCHAR(100) CHECK ( name <> '' ) NOT NULL,
    price REAL CHECK ( price >= 0 )         NOT NULL,
    type  VARCHAR(100) CHECK (type <> '' )  NOT NULL
);