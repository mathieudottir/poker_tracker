#!/bin/bash

echo "🃏 Démarrage Simple - Sans Docker"
echo ""

# Arrêter les anciens processus
pkill -f 'go run' 2>/dev/null
pkill -f 'streamlit' 2>/dev/null
sleep 1

# PostgreSQL devrait déjà tourner
if ! ps aux | grep -v grep | grep postgres > /dev/null; then
    echo "⚠️  Démarrage de PostgreSQL..."
    pg_ctlcluster 16 main start
    sleep 2
fi

echo "✅ PostgreSQL tourne"

# Créer la DB en tant qu'utilisateur actuel (pas postgres)
echo "📊 Configuration de la base de données..."

# S'assurer que le mot de passe postgres est configuré
psql -U postgres -d postgres -c "ALTER USER postgres PASSWORD 'postgres';" 2>/dev/null || true

# Se connecter en tant que superuser (l'utilisateur système courant)
psql -U postgres -d postgres -c "CREATE DATABASE poker_tracker;" 2>/dev/null || echo "DB existe déjà"

# Appliquer les migrations
psql -U postgres -d poker_tracker -f backend/migrations/001_initial_schema.sql -q 2>/dev/null
psql -U postgres -d poker_tracker -f backend/migrations/002_seed_winamax_data.sql -q 2>/dev/null

echo "✅ Base de données prête"
echo ""

# Lancer le backend en arrière-plan
echo "🚀 Démarrage du backend..."
cd backend
nohup go run cmd/server/main.go > ../logs/backend.log 2>&1 &
BACKEND_PID=$!
cd ..
sleep 2

echo "⏳ Attente du backend..."
sleep 8

# Créer l'utilisateur dev
echo "👤 Création de l'utilisateur..."
curl -s -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"mathieu","password":"dev","player_name":"SIGRIDOTTIR"}' > /dev/null 2>&1

curl -s -X PUT http://localhost:8080/api/user/1/settings \
  -H "Content-Type: application/json" \
  -d '{"player_name":"SIGRIDOTTIR","wina_status":"Aluminium","hh_directory":"./examples","dev_mode":true}' > /dev/null 2>&1

echo "✅ Utilisateur créé"
echo ""

# Lancer le frontend
echo "🎨 Démarrage du frontend..."
cd frontend
nohup streamlit run app.py --server.address 0.0.0.0 > ../logs/frontend.log 2>&1 &
FRONTEND_PID=$!
cd ..

sleep 3

echo ""
echo "✅✅✅ SYSTÈME DÉMARRÉ ! ✅✅✅"
echo ""
echo "📍 Ouvrez votre navigateur:"
echo "   👉 Local:  http://localhost:8501"
echo "   👉 Remote: http://$(curl -s ifconfig.me 2>/dev/null || echo 'YOUR_IP'):8501"
echo ""
echo "🔑 Cliquez sur 'Login as Mathieu'"
echo ""
echo "📝 Pour voir les logs:"
echo "   tail -f logs/backend.log"
echo "   tail -f logs/frontend.log"
echo ""
echo "🛑 Pour arrêter:"
echo "   kill $BACKEND_PID $FRONTEND_PID"
echo "   OU: pkill -f 'go run' && pkill -f streamlit"
echo ""
