-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id serial primary key,
    name varchar,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE TABLE book (
    id serial primary key,
    title varchar,
    description varchar,
    image_url varchar,
    release_year smallint,
    price varchar,
    total_page smallint,
    thickness varchar,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    category_id smallint NOT NULL,
    CONSTRAINT fk_category_id FOREIGN KEY (category_id) REFERENCES category (id)
);

-- +migrate StatementEnd