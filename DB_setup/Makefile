init-postgres-container:
	docker run -d --name city_letterbox \
		-e POSTGRES_PASSWORD=super_secret \
		-v ./CreateTables.sql:/tmp/CreateTables.sql:ro \
		postgres:latest

execute-sql-init-file:
	docker exec -i city_letterbox psql -U postgres -d postgres < ./CreateTables.sql

execute-sql-dummy-data:
	docker exec -i city_blog psql -U postgres -d postgres < ./InsertDummyData.sql

access-postgres-container:
	docker exec -it city_blog psql -U postgres -d postgres
