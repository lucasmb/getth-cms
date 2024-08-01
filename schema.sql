
CREATE TABLE users(
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   first_name     TEXT      NOT NULL,
   last_name     TEXT      NOT NULL,
   username     TEXT      NOT NULL,
   email     TEXT      NOT NULL,
   active   BOOLEAN   DEFAULT 1
);

INSERT INTO users (first_name, last_name, username, email, active) VALUES ('Tom', 'Allen', 'TMan', 'tman@gmail.com', TRUE);
INSERT INTO users (first_name, last_name, username, email, active) VALUES ('Sam', 'Anta', 'SWom', 'swom@gmail.com', TRUE);
INSERT INTO users (first_name, last_name, username, email, active) VALUES ('Marta', 'Rodriguez', 'lamarta', 'lamarta@gmail.com', TRUE);