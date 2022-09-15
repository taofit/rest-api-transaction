DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;

CREATE TABLE IF NOT EXISTS transactions (
    id TEXT NOT NULL PRIMARY KEY,
    account_id REAL NOT NULL, 
    amount REAL,
    created_at TEXT,
    FOREIGN KEY(account_id) REFERENCES accounts(id)
);

CREATE TABLE IF NOT EXISTS accounts (
    Id TEXT NOT NULL PRIMARY KEY,
    balance REAL
);
