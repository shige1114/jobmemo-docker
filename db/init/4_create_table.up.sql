CREATE TABLE selections(
    id SERIAL PRIMARY KEY,
    /*先行の種類*/
    level INTEGER NOT NULL,
    type INTEGER NOT NULL DEFAULT 1,
    /*面接官*/
    person TEXT NOT NULL,
    /*日程*/
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    result BOOLEAN DEFAULT FALSE,
    /*面接に関するメモ*/
    good_point VARCHAR(200) NOT NULL,
    memo VARCHAR(200) NOT NULL
);

CREATE TABLE questions(
    id SERIAL PRIMARY KEY,
    /*質問の本文*/
    title VARCHAR(20) NOT NULL,
    /*返答*/
    answer VARCHAR(100) NOT NULL
);

CREATE TABLE preparations(
    /*選考と逆質問に関する連想エンティティ*/
    selections_id SERIAL REFERENCES selections (id),
    questions_id  SERIAL REFERENCES questions (id),

    PRIMARY KEY (selections_id,questions_id)
);

CREATE TABLE status(
    /*採用と選考の連想エンティティ*/
    selections_id  SERIAL REFERENCES selections(id) , 
    users_id SERIAL REFERENCES users(id),
    companies_id SERIAL REFERENCES companies(id),

    PRIMARY KEY (users_id,selections_id,companies_id)
);