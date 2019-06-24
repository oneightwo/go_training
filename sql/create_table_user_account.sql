CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE user_account
(
  id          SERIAL PRIMARY KEY,
  firstname   VARCHAR(40) CHECK ( firstname <> lastname ) NOT NULL,
  lastname    VARCHAR(70) CHECK ( lastname <> firstname ) NOT NULL,
  age         INT CHECK ( age >= 0 AND age <= 150 )       NOT NULL,
  external_id UUID DEFAULT uuid_generate_v4() NOT NULL
);

