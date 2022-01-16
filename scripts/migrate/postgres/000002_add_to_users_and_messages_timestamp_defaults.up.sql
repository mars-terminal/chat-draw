ALTER TABLE users ALTER COLUMN created_at SET DEFAULT now();
ALTER TABLE users ALTER COLUMN updated_at SET DEFAULT now();

ALTER TABLE messages ALTER COLUMN created_at SET DEFAULT now();
ALTER TABLE messages ALTER COLUMN updated_at SET DEFAULT now();
