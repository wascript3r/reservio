-- migrate:up


CREATE TYPE role_name AS enum ('client', 'company', 'admin');

CREATE TABLE users
(
    id       uuid DEFAULT gen_random_uuid() NOT NULL
        PRIMARY KEY,
    email    varchar(200)                   NOT NULL,
    password varchar(70)                    NOT NULL,
    role     role_name                      NOT NULL
);

CREATE UNIQUE INDEX users_email_uindex
    ON users (LOWER(email::text));


-- migrate:down

