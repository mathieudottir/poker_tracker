# ⚡ Démarrage Rapide - Poker Tracker

## 🎯 3 Étapes Simples

### 1️⃣ Démarrer PostgreSQL

```bash
sudo service postgresql start
```

### 2️⃣ Configurer la Base de Données (première fois seulement)

```bash
sudo -u postgres psql << EOF
CREATE DATABASE poker_tracker;
CREATE USER poker_user WITH PASSWORD 'poker_pass';
GRANT ALL PRIVILEGES ON DATABASE poker_tracker TO poker_user;
\q
EOF

# Appliquer les migrations
sudo -u postgres psql -d poker_tracker -f backend/migrations/001_initial_schema.sql
sudo -u postgres psql -d poker_tracker -f backend/migrations/002_seed_winamax_data.sql
```

### 3️⃣ Lancer l'Application

#### Terminal 1 - Backend
```bash
cd backend
go run cmd/server/main.go
```

#### Terminal 2 - Frontend
```bash
cd frontend
streamlit run app.py
```

#### Terminal 3 - Créer l'utilisateur (première fois)
```bash
sleep 5  # Attendre que le backend démarre

curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"mathieu","password":"dev","player_name":"SIGRIDOTTIR"}'

curl -X PUT http://localhost:8080/api/user/1/settings \
  -H "Content-Type: application/json" \
  -d '{"player_name":"SIGRIDOTTIR","wina_status":"Aluminium","hh_directory":"./examples","dev_mode":true}'
```

## 🌐 Accès

- **Frontend**: http://localhost:8501
- **Backend API**: http://localhost:8080

Cliquez sur **"Login as Mathieu"**

## 📦 Importer les Exemples

1. Aller dans l'onglet **Import**
2. Vérifier que le chemin est `./examples`
3. Cliquer sur **"Import from Directory"**

## 🛑 Arrêter

```bash
# Ctrl+C dans chaque terminal
# OU
pkill -f 'go run cmd/server'
pkill -f 'streamlit run'
```

---

**Note**: Pour un démarrage avec Docker (plus simple), suivez `SETUP_GUIDE.md`
