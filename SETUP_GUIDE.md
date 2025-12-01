# 🚀 Guide de Configuration Complète

## État Actuel

✅ Go installé
✅ Python3 installé
✅ PostgreSQL installé
✅ Dépendances Python installées
✅ Code à jour

❌ Docker non installé
❌ PostgreSQL non démarré

## Option 1: Installer Docker (Recommandé)

### Ubuntu/Debian
```bash
# Installer Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sudo sh get-docker.sh
sudo usermod -aG docker $USER

# Installer Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Redémarrer la session pour appliquer les permissions
# Puis lancer le système
./start.sh
```

## Option 2: Lancer Localement (Sans Docker)

### Étape 1: Démarrer PostgreSQL

```bash
# Démarrer PostgreSQL
sudo service postgresql start

# OU si vous avez les droits sur pg_ctl
pg_ctl -D /var/lib/postgresql/data start
```

### Étape 2: Créer la Base de Données

```bash
# En tant que root ou avec sudo
sudo -u postgres psql -c "CREATE DATABASE poker_tracker;"
sudo -u postgres psql -c "CREATE USER poker_user WITH PASSWORD 'poker_pass';"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE poker_tracker TO poker_user;"

# Appliquer les migrations
sudo -u postgres psql -d poker_tracker -f backend/migrations/001_initial_schema.sql
sudo -u postgres psql -d poker_tracker -f backend/migrations/002_seed_winamax_data.sql
```

### Étape 3: Configurer l'Application

Éditez `config.yaml` si nécessaire (les valeurs par défaut devraient fonctionner).

### Étape 4: Lancer le Backend

Dans un terminal :

```bash
cd backend
go run cmd/server/main.go
```

Le backend sera sur **http://localhost:8080**

### Étape 5: Lancer le Frontend

Dans un autre terminal :

```bash
cd frontend
streamlit run app.py
```

Le frontend sera sur **http://localhost:8501**

### Étape 6: Créer l'Utilisateur Dev

Dans un troisième terminal :

```bash
# Attendre que le backend soit démarré (5 secondes)
sleep 5

# Créer l'utilisateur
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"mathieu","password":"dev","player_name":"SIGRIDOTTIR"}'

# Configurer
curl -X PUT http://localhost:8080/api/user/1/settings \
  -H "Content-Type: application/json" \
  -d '{"player_name":"SIGRIDOTTIR","wina_status":"Aluminium","hh_directory":"./examples","dev_mode":true}'
```

## Option 3: Script Automatique (Si PostgreSQL est déjà configuré)

Si PostgreSQL est déjà démarré et configuré:

```bash
# Démarrer PostgreSQL au préalable
sudo service postgresql start

# Puis utiliser le script
./start-local.sh
```

## Vérification

Une fois tout lancé, ouvrez votre navigateur:
- Frontend: http://localhost:8501
- Backend API: http://localhost:8080

Cliquez sur "Login as Mathieu" dans le frontend.

## Arrêter les Services

```bash
# Arrêter Backend et Frontend
pkill -f 'go run cmd/server/main.go'
pkill -f 'streamlit run app.py'

# Arrêter PostgreSQL
sudo service postgresql stop
```

## Troubleshooting

### "Connection refused" sur PostgreSQL

```bash
# Vérifier si PostgreSQL tourne
ps aux | grep postgres

# Démarrer si nécessaire
sudo service postgresql start
```

### "Cannot connect to backend"

```bash
# Vérifier que le backend tourne
curl http://localhost:8080/api/winamax/rakeback

# Si erreur, vérifier les logs
cd backend
go run cmd/server/main.go
```

### "Module not found" en Python

```bash
# Réinstaller les dépendances
pip3 install --user -r frontend/requirements.txt
```

## Quick Start (Si tout est configuré)

```bash
# Terminal 1
cd backend && go run cmd/server/main.go

# Terminal 2
cd frontend && streamlit run app.py

# Terminal 3 - Ouvrir le navigateur
xdg-open http://localhost:8501  # Linux
open http://localhost:8501       # macOS
```

## Support

Si vous rencontrez des problèmes, vérifiez :
1. PostgreSQL est démarré
2. Le backend se lance sans erreur
3. Le frontend se connecte au backend
4. Les ports 8080 et 8501 sont libres
