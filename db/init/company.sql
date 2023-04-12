CREATE TABLE IF NOT EXISTS campanies (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    industry TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
/*
CREATE TABLE IF NOT EXISTS company_reviews (
    id SERIAL PRIMARY KEY,
    company_id INTEGER NOT NULL,
    review TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    FOREIGN KEY (reviewer_id) REFERENCES users(id),
    FOREIGN KEY (company_id) REFERENCES company(id)
);
*/
