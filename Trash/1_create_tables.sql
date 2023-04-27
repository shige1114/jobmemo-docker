CREATE TABLE IF NOT EXISTS campany (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    industry TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
);

CREATE TABLE IF NOT EXISTS company_statuses(
    id SERIAL PRIMARY KEY,
    /*選考段階を表す*/
    recruitment_status TEXT NOT NULL,
    is_offered BOOLEAN NULL,
    
    /*面接に関する状態*/
    is_scheduling BOOLEAN NOT NULL DEFAULT FALSE,
    interveiw_date DATE,

    /*面接合否に関する状態*/
    is_accepting BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (company_id) REFERENCES company(id),
);

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    mail TEXT NOT NULL,
    image TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
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
