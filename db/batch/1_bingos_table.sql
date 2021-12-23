CREATE TABLE bingos (
  id        SERIAL primary key,
  title  varchar(255) NOT NULL
);

CREATE TABLE TODOS (
  id        SERIAL primary key,
  title     varchar(255) NOT NULL,
  bingo_id  INT NOT NULL REFERENCES bingos(id),
  is_completed BOOLEAN DEFAULT FALSE
);
