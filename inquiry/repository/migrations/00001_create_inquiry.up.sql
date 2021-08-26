CREATE TABLE IF NOT EXISTS inquiry (
   id uuid PRIMARY KEY,
   transaction_id text NOT NULL,
   payment_code text NOT NULL ,
   created_at timestamptz NOT NULL DEFAULT NOW(),
   updated_at timestamptz NOT NULL DEFAULT NOW()
);