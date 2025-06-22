CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL PRIMARY KEY,
  user_name VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(100),
  email VARCHAR(100) UNIQUE,
  admin BOOLEAN
);

INSERT INTO users (user_name,password,email,admin) 
VALUES ('admin','ZAQ!2wsx','admin@op.pl',true);