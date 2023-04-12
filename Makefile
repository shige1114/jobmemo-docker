
buildback:
	docker-compose -f docker-compose.backend.yml -p backend up --build
startback:
	docker-compose -f docker-compose.backend.yml -p backend up -b
back:
	docker exec -it backend-golang /bin/sh


buildfront:
	docker-compose -f docker-compose.frontend.yml -p frontend up --build
startfront:
	docker-compose -f docker-compose.frontend.yml -p frontend up  -b
front:
	docker exec -it frontend-nextjs /bin/sh
delete:
	docker-compose -f docker-compose.backend.yml rm 
	docker-compose -f docker-compose.frontend.yml rm
	rm -rf ./src/backend/*
	rm -rf ./src/frontend/*
