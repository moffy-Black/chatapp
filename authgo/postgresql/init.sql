\c chatuser
CREATE TABLE users(
  id serial PRIMARY KEY,
  name varchar(30) UNIQUE NOT NULL,
  password varchar(30) NOT NULL,
  created_by timestamp NOT NULL);