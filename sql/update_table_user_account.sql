DROP TABLE user_account;

CREATE TABLE user_account
(
    id        SERIAL PRIMARY KEY,
    firstname VARCHAR(40) CHECK ( firstname <> lastname ) NOT NULL,
    lastname  VARCHAR(70) CHECK ( lastname <> firstname ) NOT NULL,
    age       INT CHECK ( age >= 0 AND age <= 150 )       NOT NULL,
    login     VARCHAR(70) UNIQUE                          NOT NULL,
    password  VARCHAR(70)                                 NOT NULL
);
