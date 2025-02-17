CREATE TABLE code_snippets (
   id          SERIAL PRIMARY KEY,
   slug        TEXT UNIQUE NOT NULL,
   title       TEXT NOT NULL,
   snippet     TEXT NOT NULL,
   lang        TEXT NOT NULL,
   public      boolean NOT NULL DEFAULT FALSE,
   view_count  integer NOT NULL DEFAULT 0,
   created_time timestamptz NOT NULL DEFAULT (now()),
   updated_time timestamptz NOT NULL DEFAULT (now())
);