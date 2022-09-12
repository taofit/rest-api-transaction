DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;

CREATE TABLE IF NOT EXISTS transactions (
    id TEXT NOT NULL PRIMARY KEY,
    account_id INTEGER NOT NULL, 
    amount INTEGER,
    created_at TEXT,
    FOREIGN KEY(account_id) REFERENCES accounts(id)
);

CREATE TABLE IF NOT EXISTS accounts (
    Id TEXT NOT NULL PRIMARY KEY,
    balance INTEGER
);

-- INSERT INTO Categories (Id, Name) VALUES
-- (1, "Watersports"),database
-- (2, "Soccer");

-- INSERT INTO Products (Id, Name, Category, Price) VALUES 
-- (1, "Kayak", 1, 279),
-- (2, "Lifejacket", 1, 48.95),
-- (3, "Soccer Ball", 2, 19.50),
-- (4, "Corner Flags", 2, 34.95);