CREATE TABLE IF NOT EXISTS wallets
(
    "id"         serial PRIMARY KEY,
    "created_at" timestamp NOT NULL DEFAULT now(),
    "updated_at" timestamp          DEFAULT NULL,
    "deleted_at" timestamp          DEFAULT NULL,
    "address"    text      NOT NULL,
    "network"    text      NOT NULL,
    unique (address, network)
);