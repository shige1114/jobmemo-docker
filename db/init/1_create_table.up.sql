CREATE TABLE users (
    id UUID PRIMARY KEY,
    /*個人の特定*/
    name TEXT NOT NULL,
    email TEXT NOT NULL,

    /*自己分析*/
    future VARCHAR(300) NOT NULL  DEFAULT '',
    pr VARCHAR(300) NOT NULL  DEFAULT '',
    good_point VARCHAR(200) NOT NULL  DEFAULT '',
    bad_point VARCHAR(200) NOT NULL  DEFAULT '',
    
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);
/*就活軸*/
CREATE TABLE cores (
    id UUID PRIMARY KEY,
    title VARCHAR(8) NOT NULL,
    reason VARCHAR(100) NOT NULL,

    users_id UUID REFERENCES users(id),
    
    FOREIGN KEY(users_id) REFERENCES users(id) ON DELETE CASCADE
);