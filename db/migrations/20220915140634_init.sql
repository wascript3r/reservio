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
    id          uuid                                   NOT NULL
        PRIMARY KEY
        REFERENCES users,
    name        varchar(100)                           NOT NULL,
    address     varchar(200)                           NOT NULL,
    description text                                   NOT NULL,
    approved    boolean                  DEFAULT FALSE NOT NULL,
    created_at  timestamp WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE TABLE services
(
    id               uuid DEFAULT gen_random_uuid() NOT NULL
        PRIMARY KEY,
    company_id       uuid                           NOT NULL
        REFERENCES companies,
    title            varchar(100)                   NOT NULL,
    description      text                           NOT NULL,
    specialist_name  varchar(100),
    specialist_phone varchar(20),
    created_at       timestamp WITH TIME ZONE       NOT NULL
);


-- migrate:down

