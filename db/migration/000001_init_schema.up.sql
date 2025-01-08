-- Table: accounts
CREATE TABLE accounts (
    id BIGSERIAL PRIMARY KEY,
    owner VARCHAR NOT NULL,
    balance BIGINT NOT NULL,
    currency VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Indexes for accounts
CREATE INDEX idx_accounts_owner ON accounts (owner);

-- Table: entries
CREATE TABLE entries (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts (id),
    amount BIGINT NOT NULL, -- can be negative or positive
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Indexes for entries
CREATE INDEX idx_entries_account_id ON entries (account_id);

-- Table: transfers
CREATE TABLE transfers (
    id BIGSERIAL PRIMARY KEY,
    from_account_id BIGINT NOT NULL REFERENCES accounts (id),
    to_account_id BIGINT NOT NULL REFERENCES accounts (id),
    amount BIGINT NOT NULL, -- must be positive
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Indexes for transfers
CREATE INDEX idx_transfers_from_account_id ON transfers (from_account_id);
CREATE INDEX idx_transfers_to_account_id ON transfers (to_account_id);
CREATE INDEX idx_transfers_from_to_accounts ON transfers (from_account_id, to_account_id);
