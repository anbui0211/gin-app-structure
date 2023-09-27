-- +migrate Up

-- Create table
CREATE TABLE IF NOT EXISTS users (
  id VARCHAR(50) NOT NULL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  phone VARCHAR(20) NOT NULL UNIQUE,
  email VARCHAR(50) NOT NULL UNIQUE,
  password VARCHAR(50) NOT NULL,
  isActive BOOLEAN NOT NULL,
  status_code SMALLINT NOT NULL,
  update_count INTEGER NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- Create trigger: Insert users_history when create user
DROP TRIGGER IF EXISTS tg_users_insert ON users;
CREATE FUNCTION tg_users_insert_function() RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO users_history SELECT * FROM users WHERE id = NEW.id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create trigger: Create users_history when update user
DROP TRIGGER IF EXISTS tg_users_update ON users;
CREATE FUNCTION tg_users_update_function() RETURNS TRIGGER AS $$
BEGIN
  INSERT INTO users_history SELECT * FROM users WHERE id = NEW.id AND update_count > OLD.update_count;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Excute trigger create
CREATE TRIGGER tg_users_insert AFTER INSERT ON users FOR EACH ROW
EXECUTE FUNCTION tg_users_insert_function();

-- Excute trigger update
CREATE TRIGGER tg_users_update AFTER UPDATE ON users FOR EACH ROW
EXECUTE FUNCTION tg_users_update_function();