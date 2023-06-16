SELECT c.name AS name, r.companies_id AS id, s.date AS date, r.reject reject, r.offer AS offer ,s.level AS level , s.adjusting AS adjusting
FROM recruits r 
JOIN companies c ON r.companies_id = c.id 
JOIN selections s ON r.companies_id = s.companies_id 
WHERE r.users_id =  '185ffaae-e320-11ed-8886-26359435711c'
AND s.level = (SELECT MAX(level) FROM selections WHERE companies_id = r.companies_id);
