/*選考に関するテーブル*/
CREATE TABLE selections(
    /*先行の種類*/
    level INTEGER NOT NULL DEFAULT 0,
    type TEXT NOT NULL DEFAULT '',
    /*面接官*/
    person TEXT NOT NULL DEFAULT '',
    /*日程*/
    adjusting BOOLEAN DEFAULT FALSE,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    /*合否*/
    pass BOOLEAN DEFAULT FALSE,
    fail BOOLEAN DEFAULT FALSE,

    /*面接に関するメモ*/
    good_point VARCHAR(200) NOT NULL DEFAULT '',
    memo VARCHAR(200) NOT NULL DEFAULT '',
    /*ユーザと会社の関係*/
    users_id UUID REFERENCES recruits(users_id),
    companies_id UUID REFERENCES recruits(companies_id),
    FOREIGN KEY (users_id, companies_id) REFERENCES recruits (users_id, companies_id) ON DELETE CASCADE,

    PRIMARY KEY (level,users_id,companies_id)
);

/*逆質問*/
CREATE TABLE questions(
    id UUID PRIMARY KEY,
    /*質問の本文*/
    title VARCHAR(50) NOT NULL DEFAULT '',
    /*返答*/
    answer VARCHAR(100) NOT NULL DEFAULT '',

    users_id UUID NOT NULL,
    companies_id UUID NOT NULL,
    level INTEGER NOT NULL,
    FOREIGN KEY(users_id,companies_id,level) REFERENCES selections(users_id,companies_id,level) ON DELETE CASCADE
);