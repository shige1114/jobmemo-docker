INSERT INTO users (id,name,email,future,pr,good_point,bad_point) VALUES (
    '185ffaae-e320-11ed-8886-26359435711c',
    '就活太郎',
    'syuukatu@gmail.com',
    'future',
    'pr',
    'good',
    'bad'
);
INSERT INTO cores (id,title,reason,users_id) VALUES(
    '075f3156-e321-11ed-8886-26359435711c',
    'title',
    'reason',
    '185ffaae-e320-11ed-8886-26359435711c'
);

INSERT INTO companies (
    id,name,industry
) VALUES (
    '21c39950-e322-11ed-8886-26359435711c',
    '株式会社採用します',
    'IT'
);

INSERT INTO recruits(
    users_id,companies_id,motivation,good_point,concerns
)VALUES (
    '185ffaae-e320-11ed-8886-26359435711c',
    '21c39950-e322-11ed-8886-26359435711c',
    'motivation',
    'good',
    'concerns'
);


INSERT INTO selections (
    users_id,companies_id
) VALUES (
    '185ffaae-e320-11ed-8886-26359435711c',
    '21c39950-e322-11ed-8886-26359435711c'
);

INSERT INTO selections (
    level,users_id,companies_id
) VALUES (

    1,
        '185ffaae-e320-11ed-8886-26359435711c',
    '21c39950-e322-11ed-8886-26359435711c'
);


