CREATE TABLE IF NOT EXISTS Users (
  id SERIAL PRIMARY KEY,
  first_name VARCHAR(25),
  last_name VARCHAR(50),
  user_name VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(100),
  email VARCHAR(100) UNIQUE,
  admin BOOLEAN,
  user_inner_id VARCHAR(8),
  pfp VARCHAR(35)
);

CREATE TABLE IF NOT EXISTS Libraries (
  id SERIAL PRIMARY KEY,
  name VARCHAR(60) NOT NULL,
  location VARCHAR(60) NOT NULL,
  penalty_per_day FLOAT
);

CREATE TABLE IF NOT EXISTS Books (
  id SERIAL PRIMARY KEY,
  name VARCHAR(70) NOT NULL,
  author VARCHAR(80),
  price FLOAT,
  genre VARCHAR(15),
  library_id INT NOT NULL,
  is_avaible BOOLEAN,
  cover TEXT DEFAULT 'images/covers/def.png',
  CONSTRAINT fk_books FOREIGN KEY (library_id)
  REFERENCES Libraries(id)
);
CREATE TABLE IF NOT EXISTS Rented (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  book_id INT NOT NULL,
  to_return DATE,
  penalty FLOAT,
  is_extended BOOLEAN,
  CONSTRAINT fk_books2 FOREIGN KEY (book_id)
  REFERENCES Books(id) ,
  CONSTRAINT fk_users2 FOREIGN KEY (user_id)
  REFERENCES Users(id)
);
INSERT INTO users (first_name,last_name,user_name,password,email,admin,user_inner_id,pfp) 
VALUES ('Jan','Nowak','admin','ZAQ!2wsx','admin@op.pl',true,'00000000','profiles/admin.png');

INSERT INTO libraries (name,location,penalty_per_day) 
VALUES ('Biblioteka miejska w żorkach','żorki,Fiksalna 12',2.50);
INSERT INTO libraries (name,location,penalty_per_day) 
VALUES ('Filia Siemianowicach Śląskich','Siemianowice Śląskie,Powstańców 5',1.50);
INSERT INTO libraries (name,location,penalty_per_day) 
VALUES ('Biblioteka miejska w Katowicach','Katowice,Grzybna 7',5.50);

INSERT INTO Books (name,author,price,genre,library_id,is_avaible,cover)
VALUES ('Gillgamesh','Ben Bahsonn',45.99,'Akcja',2,true,'images/covers/Gillgamesh.png');
INSERT INTO Books (name,author,price,genre,library_id,is_avaible,cover)
VALUES ('Kokocimek','Tymoteusz Wariat',25.50,'Dramat',1,true,'images/covers/Kokocimek.png');

