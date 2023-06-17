CREATE TABLE recruits(
    /*就活生と企業の連想エンティティ*/
    users_id  UUID REFERENCES users(id),
    companies_id  UUID REFERENCES companies(id),
    /*ID*/
    PRIMARY KEY(users_id,companies_id),
    /*合否*/
    reject BOOLEAN NOT NULL DEFAULT  FALSE,
    offer BOOLEAN NOT NULL DEFAULT FALSE,
    /*企業研究*/
    motivation VARCHAR(300) NOT NULL DEFAULT '',
    good_point VARCHAR(200) NOT NULL DEFAULT '',
    concerns VARCHAR(200) NOT NULL DEFAULT ''
);





