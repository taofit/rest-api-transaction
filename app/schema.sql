DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;

CREATE TABLE IF NOT EXISTS transactions (
    id TEXT NOT NULL PRIMARY KEY,
    account_id TEXT NOT NULL, 
    amount int,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(account_id) REFERENCES accounts(id)
);

CREATE TABLE IF NOT EXISTS accounts (
    id TEXT NOT NULL PRIMARY KEY,
    balance int
);
