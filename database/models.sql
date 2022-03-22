CREATE TABLE IF NOT EXISTS users (
    id serial NOT NULL,
    username VARCHAR(150) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    picture VARCHAR(256) NOT NULL,
    authorized BOOLEAN DEFAULT FALSE, 
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_users PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS roles (
    id int NOT NULL,
    name VARCHAR(150) NOT NULL UNIQUE,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_roles PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id int NOT NULL,
    role_id int NOT NULL, 
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    PRIMARY KEY(user_id, role_id),
    CONSTRAINT fk_role FOREIGN KEY(role_id) REFERENCES roles(id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS messages (
    id serial NOT NULL,
    firstname VARCHAR(50) NULL,
    lastname VARCHAR(50) NULL,
    email VARCHAR(150) NOT NULL,
    phone VARCHAR(100) NULL,
    message VARCHAR NULL,
    created_at timestamp DEFAULT now()
);

INSERT INTO roles VALUES (1, 'user', now(), now())  ON CONFLICT DO NOTHING;
INSERT INTO roles VALUES (2, 'moderator', now(), now())  ON CONFLICT DO NOTHING;
INSERT INTO roles VALUES (3, 'admin', now(), now())  ON CONFLICT DO NOTHING;

INSERT INTO users VALUES (1, 'freis', '$2a$10$n48Uhbhqmmb9rKf4IF7lYea0wPQOgZd0JC30Om4xceFOyXIoFm0/q', 'freis@test.com', 'https://placekitten.com/g/300/300', false, now(), now()) ON CONFLICT DO NOTHING;
INSERT INTO user_roles VALUES (1, 1, now(), now()) ON CONFLICT DO NOTHING;
INSERT INTO user_roles VALUES (1, 3, now(), now()) ON CONFLICT DO NOTHING;
