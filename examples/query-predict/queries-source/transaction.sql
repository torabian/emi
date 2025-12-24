--skip
BEGIN TRANSACTION;

INSERT INTO users(name, email) VALUES ('Alice', 'alice@example.com');
INSERT INTO users(name, email) VALUES ('Bob', 'bob@example.com');

COMMIT;