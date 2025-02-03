DOCKERCOMPOSE_PATH := ../docker/docker-compose-db.yml
POSTGRES_CONTAINER := postgres
POSTGRES_USER := myuser
POSTGRES_DB := mydb

# Цель для старта базы данных
start_dbpsq:
	docker-compose -f $(DOCKERCOMPOSE_PATH) up -d

# Цель для остановки базы данных
stop_dbpsq:
	docker-compose -f $(DOCKERCOMPOSE_PATH) down

# Цель для удаления контейнера и образа PostgreSQL
rm_dbpsq:
	docker rmi $(POSTGRES_CONTAINER)

# Цель для подключения к базе данных
connect_dbpsq:
	docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(POSTGRES_DB)

# Проверка на правильное количество аргументов
help:
	@echo "Usage:"
	@echo "  make start_dbpsq   - Start the PostgreSQL database"
	@echo "  make stop_dbpsq    - Stop the PostgreSQL database"
	@echo "  make rm_dbpsq      - Remove the PostgreSQL container and image"
	@echo "  make connect_dbpsq - Connect to PostgreSQL database"