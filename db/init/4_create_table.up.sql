/*選考に関するテーブル*/
CREATE TABLE selections(
    id UUID PRIMARY KEY,
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
    users_id UUID REFERENCES users(id),
    companies_id UUID REFERENCES companies(id),
    
    FOREIGN KEY(users_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(companies_id) REFERENCES companies(id) ON DELETE CASCADE
);

/*逆質問*/
CREATE TABLE questions(
    id UUID PRIMARY KEY,
    /*質問の本文*/
    title VARCHAR(20) NOT NULL,
    /*返答*/
    answer VARCHAR(100) NOT NULL DEFAULT '',

    selections_id UUID REFERENCES selections(id)
);