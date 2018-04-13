
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_types') THEN
        CREATE TYPE user_types AS ENUM ('client','executor');
    END IF;
END$$;

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name varchar(255) NOT NULL,
  utype user_types,
  balance INTEGER
) ;