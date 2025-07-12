CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS wallet (
      id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
      amount NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
      CHECK (amount >= 0.00)
);