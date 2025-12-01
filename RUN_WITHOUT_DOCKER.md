# 🚀 Lancer sans Docker

Si Docker n'est pas disponible, vous pouvez lancer le système localement.

## Prérequis

- Go 1.21+
- Python 3.11+
- PostgreSQL 15+

## 1. Installer PostgreSQL

### Ubuntu/Debian
```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
sudo systemctl start postgresql
```

### macOS
```bash
brew install postgresql@15
brew services start postgresql@15
```

## 2. Créer la base de données

```bash
# Se connecter à PostgreSQL
sudo -u postgres psql

# Créer la base de données
CREATE DATABASE poker_tracker;
CREATE USER poker_user WITH PASSWORD 'poker_pass';
GRANT ALL PRIVILEGES ON DATABASE poker_tracker TO poker_user;
\q
```

## 3. Configurer l'application

Éditer `config.yaml` :

```yaml
server:
  host: "0.0.0.0"
  port: 8080

database:
  host: "localhost"
  port: 5432
  user: "poker_user"
  password: "poker_pass"
  dbname: "poker_tracker"
  sslmode: "disable"

logging:
  level: "info"
  file: "logs/poker_tracker.log"

dev_mode: true
```

## 4. Lancer le Backend

```bash
cd backend

# Installer les dépendances
go mod download

# Lancer le serveur
go run cmd/server/main.go
```

Le backend sera disponible sur http://localhost:8080

## 5. Lancer le Frontend (dans un autre terminal)

```bash
cd frontend

# Créer un environnement virtuel
python -m venv venv
source venv/bin/activate  # Sur Windows: venv\Scripts\activate

# Installer les dépendances
pip install -r requirements.txt

# Définir l'URL de l'API
export API_URL=http://localhost:8080/api

# Lancer Streamlit
streamlit run app.py
```

Le frontend sera disponible sur http://localhost:8501

## 6. Créer l'utilisateur dev

```bash
# Enregistrer l'utilisateur
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"mathieu","password":"dev","player_name":"SIGRIDOTTIR"}'

# Configurer les paramètres
curl -X PUT http://localhost:8080/api/user/1/settings \
  -H "Content-Type: application/json" \
  -d '{"player_name":"SIGRIDOTTIR","wina_status":"Aluminium","hh_directory":"./examples","dev_mode":true}'
```

## 7. Importer les exemples

Dans l'interface web:
1. Cliquer sur "Login as Mathieu"
2. Aller dans "Import"
3. Entrer `./examples` comme chemin
4. Cliquer sur "Import from Directory"

## Arrêter les services

```bash
# Backend: Ctrl+C dans le terminal
# Frontend: Ctrl+C dans le terminal

# PostgreSQL
sudo systemctl stop postgresql  # Linux
brew services stop postgresql@15  # macOS
```

## Développement

### Backend uniquement
```bash
make dev-backend
```

### Frontend uniquement
```bash
make dev-frontend
```

## Troubleshooting

### Le backend ne se connecte pas à PostgreSQL

Vérifier que PostgreSQL est démarré:
```bash
sudo systemctl status postgresql  # Linux
brew services list  # macOS
```

Vérifier les credentials dans `config.yaml`

### Le frontend ne se connecte pas au backend

Vérifier que la variable d'environnement API_URL est définie:
```bash
echo $API_URL
# Devrait afficher: http://localhost:8080/api
```

### Erreur de migration

Exécuter manuellement les migrations:
```bash
psql -U poker_user -d poker_tracker -f backend/migrations/001_initial_schema.sql
psql -U poker_user -d poker_tracker -f backend/migrations/002_seed_winamax_data.sql
```
