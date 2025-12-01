.PHONY: help build run stop clean db-setup dev-user test

# Detect docker compose command
DOCKER_COMPOSE := $(shell command -v docker-compose 2> /dev/null)
ifndef DOCKER_COMPOSE
	DOCKER_COMPOSE := docker compose
endif

help:
	@echo "Winamax Expresso Tracker - Makefile Commands"
	@echo ""
	@echo "  make build       - Build Docker containers"
	@echo "  make run         - Run the complete system"
	@echo "  make stop        - Stop all containers"
	@echo "  make clean       - Clean up containers and volumes"
	@echo "  make db-setup    - Initialize database"
	@echo "  make dev-user    - Create dev user (Mathieu)"
	@echo "  make test        - Run tests"
	@echo "  make logs        - Show logs"
	@echo ""

build:
	$(DOCKER_COMPOSE) build

run:
	$(DOCKER_COMPOSE) up -d
	@echo "✅ System started!"
	@echo "Backend API: http://localhost:8080"
	@echo "Frontend UI: http://localhost:8501"
	@echo ""
	@echo "Creating dev user..."
	@sleep 5
	@make dev-user

stop:
	$(DOCKER_COMPOSE) down

clean:
	$(DOCKER_COMPOSE) down -v
	rm -rf logs/*

db-setup:
	$(DOCKER_COMPOSE) exec postgres psql -U postgres -d poker_tracker -c "SELECT version();"

dev-user:
	@curl -X POST http://localhost:8080/api/auth/register \
		-H "Content-Type: application/json" \
		-d '{"username":"mathieu","password":"dev","player_name":"SIGRIDOTTIR"}' \
		2>/dev/null || echo "User may already exist"
	@curl -X PUT http://localhost:8080/api/user/1/settings \
		-H "Content-Type: application/json" \
		-d '{"player_name":"SIGRIDOTTIR","wina_status":"Aluminium","hh_directory":"./examples","dev_mode":true}' \
		2>/dev/null || true
	@echo "✅ Dev user created: mathieu / SIGRIDOTTIR"

test:
	cd backend && go test ./...

logs:
	$(DOCKER_COMPOSE) logs -f

logs-backend:
	$(DOCKER_COMPOSE) logs -f backend

logs-frontend:
	$(DOCKER_COMPOSE) logs -f frontend

logs-db:
	$(DOCKER_COMPOSE) logs -f postgres

# Development shortcuts
dev-backend:
	cd backend && go run cmd/server/main.go

dev-frontend:
	cd frontend && streamlit run app.py

install-deps:
	cd backend && go mod tidy
	cd frontend && pip install -r requirements.txt
