CREATE TABLE campaign_service.campaigns (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    creator_id TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT,
    goal_amount NUMERIC NOT NULL,
    raised_amount NUMERIC NOT NULL DEFAULT 0,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
