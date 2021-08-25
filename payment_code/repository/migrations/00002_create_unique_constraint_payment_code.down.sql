ALTER TABLE payment_codes
ADD CONSTRAINT uk_payment_code
UNIQUE KEY (payment_code)