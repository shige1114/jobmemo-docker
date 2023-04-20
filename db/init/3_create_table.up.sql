CREATE TABLE recruits(
    /*就活生と企業の連想エンティティ*/
    users_id  UUID REFERENCES users(id),
    companies_id  UUID REFERENCES companies(id),
    /*ID*/
    PRIMARY KEY(users_id,companies_id),

    
    reject BOOLEAN NOT NULL,
    offer BOOLEAN NOT NULL,

    motivation VARCHAR(300) NOT NULL,
    good_point VARCHAR(200) NOT NULL,
    concerns VARCHAR(200) NOT NULL
);
