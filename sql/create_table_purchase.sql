CREATE TABLE purchase
(
    id            SERIAL PRIMARY KEY                                                              NOT NULL,
    user_id       INT REFERENCES user_account (id)                                                NOT NULL,
    product_id    INT REFERENCES product (id)                                                     NOT NULL,
    purchase_date TIMESTAMP CHECK ( purchase_date >= NOW()) DEFAULT NOW() NOT NULL
);

