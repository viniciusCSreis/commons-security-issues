CREATE TABLE IF NOT EXISTS public.book
(
    id   CHARACTER varying(36)  NOT NULL,
    name CHARACTER varying(200) NOT NULL,

    PRIMARY KEY (id)
    );

INSERT INTO book(id, name)
VALUES ('1', 'Harry Potter'),
       ('2', 'Star Wars'),
       ('3', 'Batman'),
       ('4', 'Star lord');

CREATE TABLE IF NOT EXISTS public.users
(
    id       CHARACTER varying(36)  NOT NULL,
    name     CHARACTER varying(200) NOT NULL,
    password CHARACTER varying(200) NOT NULL,

    PRIMARY KEY (id)
    );
INSERT INTO users(id, name, password)
VALUES ('1', 'Vinicius', '123'),
       ('2', 'Jef', 'password'),
       ('3', 'Mario', 'askdjqwoijh');