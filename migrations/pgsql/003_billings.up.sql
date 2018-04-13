DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'bills_types') THEN
        CREATE TYPE bills_types AS ENUM ('pay','hold');
    END IF;
END$$;

CREATE TABLE billings (
  id SERIAL PRIMARY KEY,
  sender INTEGER,
  reciever INTEGER,
  amount INTEGER,
  time_bill timestamp DEFAULT NULL,
  task_id INTEGER,
  btype bills_types
) ;