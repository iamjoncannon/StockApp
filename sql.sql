-- dropdb test && createdb test && psql -U jonathancannon -d test -a -f sql.sql

CREATE TABLE users (
  ID SERIAL PRIMARY KEY,
  name VARCHAR(30) NOT NULL,
  email VARCHAR(30) NOT NULL,
  password VARCHAR(300) NOT NULL,
  balance INT DEFAULT 5000,
  CONSTRAINT unique_email UNIQUE (email)
);

CREATE TABLE transactions (
  ID SERIAL NOT NULL PRIMARY KEY, 
  userID INT NOT NULL,
  TYPE TEXT NOT NULL,
  SYMBOL TEXT NOT NULL,
  QUANTITY INT NOT NULL,
  PRICE INT NOT NULL,
  DATE_CONDUCTED TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

