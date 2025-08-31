CREATE TABLE donation_service.donations (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    campaign_id INT,
    donor_id TEXT NOT NULL,
    amount NUMERIC NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
