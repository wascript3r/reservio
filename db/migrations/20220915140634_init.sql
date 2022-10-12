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
    id               uuid                     DEFAULT gen_random_uuid() NOT NULL
        PRIMARY KEY,
    company_id       uuid                                               NOT NULL
        REFERENCES companies,
    title            varchar(100)                                       NOT NULL,
    description      text                                               NOT NULL,
    specialist_name  varchar(100),
    specialist_phone varchar(20),
    visit_duration   integer                                            NOT NULL,
    work_schedule    jsonb                                              NOT NULL,
    created_at       timestamp WITH TIME ZONE DEFAULT NOW()             NOT NULL
);

CREATE TABLE clients
(
    id         uuid                                   NOT NULL
        PRIMARY KEY
        REFERENCES users,
    first_name varchar(50)                            NOT NULL,
    last_name  varchar(50)                            NOT NULL,
    phone      varchar(20)                            NOT NULL,
    created_at timestamp WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE TABLE reservations
(
    id         uuid                     DEFAULT gen_random_uuid() NOT NULL
        PRIMARY KEY,
    service_id uuid                                               NOT NULL
        REFERENCES services,
    client_id  uuid                                               NOT NULL
        REFERENCES clients,
    date       timestamp WITH TIME ZONE                           NOT NULL,
    comment    varchar(200),
    approved   boolean                  DEFAULT FALSE             NOT NULL,
    created_at timestamp WITH TIME ZONE DEFAULT NOW()             NOT NULL
);

CREATE TABLE refresh_tokens
(
    id         uuid DEFAULT gen_random_uuid() NOT NULL
        PRIMARY KEY,
    user_id    uuid                           NOT NULL
        REFERENCES users,
    expires_at timestamp WITH TIME ZONE       NOT NULL
);


-- migrate:down

