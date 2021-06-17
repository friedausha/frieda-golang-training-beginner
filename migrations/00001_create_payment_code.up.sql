CREATE TABLE IF NOT EXISTS payment_codes (
   id uuid PRIMARY KEY,
   payment_code TEXT NOT NULL,
   name TEXT NOT NULL,
   status varchar(20) NOT NULL,
   expiration_date timestamptz NOT NULL,
   created_at timestamptz NOT NULL DEFAULT NOW(),
   updated_at timestamptz NOT NULL DEFAULT NOW()
);