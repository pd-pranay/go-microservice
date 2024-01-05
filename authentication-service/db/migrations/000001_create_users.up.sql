CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
	id UUID NOT NULL DEFAULT uuid_generate_v4(),
  index SERIAL PRIMARY KEY,
  email TEXT UNIQUE NOT NULL,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  password TEXT NOT NULL,
  user_active INTEGER NOT NULL DEFAULT 0,
	created_at timestamptz NOT NULL DEFAULT current_timestamp,
	updated_at timestamptz NOT NULL DEFAULT current_timestamp
) WITH (
  OIDS=FALSE
);