#!/bin/bash

echo "🃏 Winamax Expresso Tracker - Startup Script"
echo ""

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    exit 1
fi

# Detect docker compose command
if command -v docker-compose &> /dev/null; then
    DOCKER_COMPOSE="docker-compose"
elif docker compose version &> /dev/null; then
    DOCKER_COMPOSE="docker compose"
else
    echo "❌ Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

echo "✅ Docker and Docker Compose are installed"
echo ""

# Build containers
echo "📦 Building containers..."
$DOCKER_COMPOSE build

echo ""
echo "🚀 Starting services..."
$DOCKER_COMPOSE up -d

echo ""
echo "⏳ Waiting for services to be ready..."
sleep 10

echo ""
echo "👤 Creating dev user..."
curl -X POST http://localhost:8080/api/auth/register \
    -H "Content-Type: application/json" \
    -d '{"username":"mathieu","password":"dev","player_name":"SIGRIDOTTIR"}' \
    2>/dev/null

sleep 2

curl -X PUT http://localhost:8080/api/user/1/settings \
    -H "Content-Type: application/json" \
    -d '{"player_name":"SIGRIDOTTIR","wina_status":"Aluminium","hh_directory":"./examples","dev_mode":true}' \
    2>/dev/null

echo ""
echo ""
echo "✅ System is ready!"
echo ""
echo "📍 Access points:"
echo "   Frontend:  http://localhost:8501"
echo "   Backend:   http://localhost:8080"
echo ""
echo "👤 Dev user credentials:"
echo "   Click 'Login as Mathieu' button"
echo "   (or use username: mathieu, password: dev)"
echo ""
echo "📂 Example tournaments are available in ./examples directory"
echo ""
echo "💡 Useful commands:"
echo "   make logs          - View all logs"
echo "   make stop          - Stop all services"
echo "   make clean         - Clean and reset everything"
echo ""
