DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'tasks_types') THEN
        CREATE TYPE tasks_types AS ENUM ('done','not done');
    END IF;
END$$;

CREATE TABLE tasks (
  id SERIAL PRIMARY KEY,
  description varchar(255) NOT NULL,
  price INTEGER,
  status tasks_types  
) ;