#!/bin/bash

echo "🃏 Winamax Expresso Tracker - Local Startup"
echo ""
echo "⚠️  This script runs the system WITHOUT Docker"
echo ""

# Check Go
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed"
    echo "   Install from: https://golang.org/dl/"
    exit 1
fi

# Check Python
if ! command -v python3 &> /dev/null; then
    echo "❌ Python 3 is not installed"
    exit 1
fi

# Check PostgreSQL
if ! command -v psql &> /dev/null; then
    echo "❌ PostgreSQL is not installed"
    echo "   Ubuntu/Debian: sudo apt install postgresql"
    echo "   macOS: brew install postgresql@15"
    exit 1
fi

echo "✅ Prerequisites installed"
echo ""

# Create database if not exists
echo "📊 Setting up database..."
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'poker_tracker'" | grep -q 1 || \
    psql -U postgres -c "CREATE DATABASE poker_tracker;"

# Run migrations
echo "🔄 Running migrations..."
psql -U postgres -d poker_tracker -f backend/migrations/001_initial_schema.sql -q 2>/dev/null
psql -U postgres -d poker_tracker -f backend/migrations/002_seed_winamax_data.sql -q 2>/dev/null

echo ""
echo "🚀 Starting Backend..."
cd backend
go run cmd/server/main.go &
BACKEND_PID=$!
cd ..

echo "⏳ Waiting for backend to start..."
sleep 5

echo ""
echo "🎨 Starting Frontend..."
cd frontend
python3 -m streamlit run app.py &
FRONTEND_PID=$!
cd ..

echo ""
echo "✅ System is running!"
echo ""
echo "📍 Access points:"
echo "   Frontend:  http://localhost:8501"
echo "   Backend:   http://localhost:8080"
echo ""
echo "👤 Creating dev user..."
sleep 3

curl -X POST http://localhost:8080/api/auth/register \
    -H "Content-Type: application/json" \
    -d '{"username":"mathieu","password":"dev","player_name":"SIGRIDOTTIR"}' \
    2>/dev/null

curl -X PUT http://localhost:8080/api/user/1/settings \
    -H "Content-Type: application/json" \
    -d '{"player_name":"SIGRIDOTTIR","wina_status":"Aluminium","hh_directory":"./examples","dev_mode":true}' \
    2>/dev/null

echo ""
echo "✅ Dev user created: mathieu / SIGRIDOTTIR"
echo ""
echo "💡 To stop the servers:"
echo "   kill $BACKEND_PID $FRONTEND_PID"
echo ""
echo "   Or press Ctrl+C and run:"
echo "   pkill -f 'go run cmd/server/main.go'"
echo "   pkill -f 'streamlit run app.py'"
echo ""

# Wait for user to stop
wait
