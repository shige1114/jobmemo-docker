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
INSERT INTO companies (
    id,name,industry
) VALUES (
    '21c39950-e322-22ed-8886-26359435711c',
    '株式会社採用です',
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
INSERT INTO recruits(
    users_id,companies_id,motivation,good_point,concerns
)VALUES (
    '185ffaae-e320-11ed-8886-26359435711c',
    '21c39950-e322-22ed-8886-26359435711c',
    'motivation',
    'good',
    'concerns'
);
INSERT INTO selections (
    users_id,companies_id,level
) VALUES (
    '185ffaae-e320-11ed-8886-26359435711c',
    '21c39950-e322-22ed-8886-26359435711c',
    1
);

INSERT INTO selections (
    users_id,companies_id,level
) VALUES (
    '185ffaae-e320-11ed-8886-26359435711c',
    '21c39950-e322-11ed-8886-26359435711c',
    1
);

INSERT INTO selections (
    users_id,companies_id,level
) VALUES (
    '185ffaae-e320-11ed-8886-26359435711c',
    '21c39950-e322-11ed-8886-26359435711c',
    2
);
INSERT INTO selections (
    users_id,companies_id,level
) VALUES (
    '185ffaae-e320-11ed-8886-26359435711c',
    '21c39950-e322-11ed-8886-26359435711c',
    3
);
INSERT INTO questions (
    id,title
) VALUES (
    'ec35b76a-e325-11ed-8886-26359435711c',
    'title'
);

SELECT c.name as companies_name, r.offer as recruits_offer,
s.level as selections_level
FROM recruits r
JOIN companies c ON r.companies_id = c.id
JOIN selections s ON r.companies_id = s.companies_id 
WHERE r.users_id = '185ffaae-e320-11ed-8886-26359435711c';



SELECT c.name, r.users_id, r.companies_id, r.reject, r.offer, s.id AS selection_id, s.level, s.type, s.person, s.adjusting, s.date, s.pass, s.fail, s.good_point, s.memo
FROM recruits r
JOIN companies c ON c.id = r.companies_id
JOIN selections s ON r.selections_id = s.id
WHERE r.users_id = '185ffaae-e320-11ed-8886-26359435711c' AND r.companies_id = '21c39950-e322-11ed-8886-26359435711c';
