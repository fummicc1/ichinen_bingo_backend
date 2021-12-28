CREATE TABLE bingos (
  id        SERIAL primary key,
  title  varchar(255) NOT NULL
);

CREATE TABLE TODOS (
  id        SERIAL primary key,
  title     varchar(255) NOT NULL,
  index     integer NOT NULL,
  bingo_id  integer NOT NULL REFERENCES bingos(id),
  is_completed BOOLEAN DEFAULT FALSE,
  UNIQUE (bingo_id, index)
);
