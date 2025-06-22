CREATE TABLE IF NOT EXISTS users (
  user_id SERIAL PRIMARY KEY,
  first_name VARCHAR(25),
  last_name VARCHAR(50),
  user_name VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(100),
  email VARCHAR(100) UNIQUE,
  admin BOOLEAN,
  user_inner_id VARCHAR(8),
  pfp VARCHAR(35)
);

INSERT INTO users (user_name,password,email,admin) 
VALUES ('admin','ZAQ!2wsx','admin@op.pl',true);