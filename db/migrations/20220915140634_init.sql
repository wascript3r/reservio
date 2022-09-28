-- migrate:up


CREATE TYPE role_name AS enum ('client', 'company', 'admin');

CREATE TABLE users
(
    id         uuid                     DEFAULT gen_random_uuid() NOT NULL
        PRIMARY KEY,
    email      varchar(200)                                       NOT NULL,
    password   varchar(70)                                        NOT NULL,
    role       role_name                                          NOT NULL,
    created_at timestamp WITH TIME ZONE DEFAULT NOW()             NOT NULL
);

CREATE UNIQUE INDEX users_email_uindex
    ON users (LOWER(email::text));

CREATE TABLE companies
(
    company_id  uuid                                   NOT NULL
        PRIMARY KEY
        REFERENCES users,
    name        varchar(100)                           NOT NULL,
    address     varchar(200)                           NOT NULL,
    description text                                   NOT NULL,
    approved    boolean                  DEFAULT FALSE NOT NULL,
    created_at  timestamp WITH TIME ZONE DEFAULT NOW() NOT NULL
);


-- migrate:down

