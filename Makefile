
buildback:
	docker-compose -f docker-compose.backend.yml -p backend up --build
startback:
	docker-compose -f docker-compose.backend.yml -p backend up
back:
	docker exec -it backend-golang /bin/bash


buildfront:
	docker-compose -f docker-compose.frontend.yml -p frontend up --build
startfront:
	docker-compose -f docker-compose.frontend.yml -p frontend up
front:
	docker exec -it frontend-nextjs /bin/bash
delete:
	docker-compose -f docker-compose.backend.yml -p backend rm
	docker-compose -f docker-compose.backend.yml -p backend rm
