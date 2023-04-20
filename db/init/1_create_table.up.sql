CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    /*個人の特定*/
    name TEXT NOT NULL,
    email TEXT NOT NULL,

    /*自己分析*/
    future VARCHAR(300) NOT NULL,
    pr VARCHAR(300) NOT NULL,
    good_point VARCHAR(200) NOT NULL,
    bad_point VARCHAR(200) NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE cores (
    /*就活軸*/
    id SERIAL PRIMARY KEY,
    title VARCHAR(8) NOT NULL,
    reason VARCHAR(100) NOT NULL
);

CREATE TABLE users_cores (
    /*就活軸とユーザの連想エンティティ*/
    users_id  SERIAL REFERENCES users(id),
    cores_id  SERIAL REFERENCES cores(id),

    PRIMARY KEY (users_id,cores_id)
);
