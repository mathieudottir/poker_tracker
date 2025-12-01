# 🃏 Winamax Expresso Tracker

Un tracker personnel complet pour les tournois Expresso Winamax (3 joueurs) avec suivi de résultats, analyse d'EV, comparaison avec les probabilités officielles, gestion du rake et rakeback.

## 🎯 Fonctionnalités

### ✅ Implémenté

- **Import de Hand Histories**
  - Import manuel de fichiers et dossiers
  - Surveillance en temps réel avec watcher automatique
  - Parsing intelligent des fichiers Winamax Expresso
  - Détection automatique des nouveaux tournois

- **Suivi de Résultats**
  - Tableau complet des tournois
  - Courbe de bankroll en temps réel
  - Courbe d'EV monétaire
  - Distribution des multiplicateurs obtenus vs attendus
  - Résultats par buy-in

- **Analyse Statistique**
  - Winrate et ROI (net et EV)
  - Chips EV par tournoi
  - Rake payé et rakeback calculé
  - Comparaison résultats réels vs EV théorique

- **Gestion Utilisateur**
  - Authentification sécurisée (bcrypt)
  - Mode dev avec login simplifié
  - Paramètres personnalisés
  - Reset complet de la base de données

- **Données Winamax**
  - Toutes les probabilités de multiplicateurs (€0.25 - €500)
  - Tous les statuts de rakeback (Aluminium → Red Diamond)
  - Calculs EV basés sur les données officielles

### 🔮 Futur (Structure Préparée)

- Tracking des adversaires (VPIP, PFR, Steal, Limp%, Check-raise%)
- Heatmaps préflop
- Replayer de mains
- Filtres avancés (3-way pot, HU, push/fold)

## 🏗️ Architecture

### Backend (Go)

```
backend/
├── cmd/server/          # Point d'entrée
├── internal/
│   ├── api/            # Handlers REST API
│   ├── auth/           # Service d'authentification
│   ├── calculator/     # Calculs EV, ROI, chips EV
│   ├── config/         # Configuration YAML
│   ├── models/         # Modèles de données
│   ├── parser/         # Parser de hand histories
│   ├── repository/     # Accès base de données
│   ├── services/       # Logique métier
│   └── watcher/        # Watcher temps réel
├── pkg/constants/      # Données Winamax hardcodées
└── migrations/         # Migrations SQL
```

### Frontend (Streamlit)

```
frontend/
├── app.py             # Application Streamlit
└── requirements.txt   # Dépendances Python
```

### Base de Données (PostgreSQL)

- `users` - Utilisateurs
- `tournaments` - Tournois
- `hands` - Mains
- `multipliers_expected` - Probabilités Winamax
- `rakeback_status` - Statuts rakeback
- `opponents` - Adversaires (future)
- `opponent_hands` - Mains adversaires (future)

## 🚀 Installation & Lancement

### Avec Docker (Recommandé)

```bash
# Construire les conteneurs
make build

# Lancer le système complet
make run

# Le système créera automatiquement un utilisateur dev:
# Username: mathieu
# Player Name: SIGRIDOTTIR
# Dev Mode: activé
```

Accès:
- Frontend: http://localhost:8501
- Backend API: http://localhost:8080

### Sans Docker

#### Prérequis

- Go 1.21+
- Python 3.11+
- PostgreSQL 15+

#### Backend

```bash
cd backend

# Installer les dépendances
go mod download

# Configurer PostgreSQL
createdb poker_tracker

# Éditer config.yaml avec vos paramètres DB

# Lancer le serveur
go run cmd/server/main.go
```

#### Frontend

```bash
cd frontend

# Installer les dépendances
pip install -r requirements.txt

# Lancer Streamlit
streamlit run app.py
```

## 📖 Utilisation

### 1. Login

Deux options:
- **Dev Mode**: Cliquer sur "Login as Mathieu" (uniquement si dev_mode activé)
- **Normal**: Username + Password

### 2. Import

**Onglet Import:**

1. Configurer le dossier HH dans Settings
2. Démarrer le watcher pour import automatique
3. OU importer manuellement depuis un dossier

Le watcher surveille le dossier en temps réel et importe automatiquement les nouveaux tournois.

### 3. Résultats

**Onglet My Results:**

- Tableau de tous les tournois
- Courbe de bankroll
- Courbe d'EV vs résultats réels
- Distribution des multiplicateurs
- Métriques globales (profit, ROI, EV)

### 4. Statistiques

**Onglet Statistics:**

- ROI net et EV ROI
- Rake payé et rakeback
- Résultats détaillés par buy-in
- Nombre de tournois et mains

### 5. Paramètres

**Onglet Settings:**

- Nom du joueur Winamax
- Dossier des hand histories
- Statut Winamax (pour calcul rakeback)
- Mode dev ON/OFF
- Reset complet de la base de données

## 📊 Calculs

### EV Monétaire

Calculé selon les probabilités officielles Winamax pour chaque buy-in:

```
EV = Σ (Probabilité × Gain moyen pour cette probabilité)
```

### ROI

```
Net ROI = (Profit / (Buy-in + Rake)) × 100
EV ROI = ((EV - Buy-in) / Buy-in) × 100
```

### Rakeback

Calculé selon le statut Winamax:
- Aluminium: 0%
- Bronze: 10%
- Silver: 12.5%
- Gold: 15%
- Platinum: 17.5%
- D1-D5: 20-30%
- Red Diamond: 33%

### Chips EV

Somme des EV de toutes les mains (implémentation simple, peut être améliorée avec équité préflop).

## 🗂️ Données Winamax Intégrées

### Multiplicateurs

Toutes les probabilités pour les buy-ins:
- €0.25, €0.50, €1, €2, €5, €10, €25, €50, €100, €250, €500

Avec multiplicateurs: 2x, 4x, 10x, 100x, 1000x

### Statuts Rakeback

Tous les statuts avec miles, rake équivalent, bonus annuel et % rakeback.

## 🛠️ Commandes Make

```bash
make help          # Afficher l'aide
make build         # Construire les conteneurs
make run           # Lancer le système
make stop          # Arrêter les conteneurs
make clean         # Nettoyer (supprime les volumes)
make logs          # Voir tous les logs
make logs-backend  # Logs du backend
make logs-frontend # Logs du frontend
make dev-user      # Créer l'utilisateur dev
```

## 🔧 Configuration

### config.yaml

```yaml
server:
  host: "0.0.0.0"
  port: 8080

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "postgres"
  dbname: "poker_tracker"
  sslmode: "disable"

logging:
  level: "info"
  file: "logs/poker_tracker.log"

dev_mode: true
```

## 📡 API Endpoints

### Auth
- `POST /api/auth/register` - Créer un compte
- `POST /api/auth/login` - Se connecter
- `POST /api/auth/dev-login` - Login dev

### User
- `GET /api/user/{id}` - Obtenir utilisateur
- `PUT /api/user/{id}/settings` - Mettre à jour paramètres
- `DELETE /api/user/{id}/data` - Supprimer données

### Tournaments
- `GET /api/tournaments/{userId}` - Liste des tournois
- `GET /api/tournament/{id}` - Détails tournoi
- `GET /api/tournament/{id}/hands` - Mains du tournoi

### Import
- `POST /api/import/tournament` - Importer un tournoi
- `POST /api/import/directory` - Importer un dossier
- `GET /api/import/logs` - Logs d'import

### Watcher
- `POST /api/watcher/start` - Démarrer le watcher
- `POST /api/watcher/stop` - Arrêter le watcher
- `GET /api/watcher/status/{userId}` - Statut du watcher

### Stats
- `GET /api/stats/{userId}` - Statistiques globales
- `GET /api/stats/{userId}/multipliers` - Distribution multiplicateurs
- `GET /api/stats/{userId}/bybuyin` - Résultats par buy-in
- `GET /api/stats/{userId}/bankroll` - Historique bankroll
- `GET /api/stats/{userId}/ev` - Historique EV

### Winamax Data
- `GET /api/winamax/multipliers` - Toutes les probabilités
- `GET /api/winamax/rakeback` - Tous les statuts

## 📁 Format des Hand Histories

Le parser accepte les fichiers Winamax au format:

```
20251130_Expresso Nitro(1025034573)_real_holdem_no-limit.txt
20251130_Expresso Nitro(1025034573)_real_holdem_no-limit_summary.txt
```

Chaque tournoi nécessite 2 fichiers:
1. Le fichier de mains (.txt)
2. Le fichier de résumé (_summary.txt)

## 🧪 Tests

```bash
# Tests backend
cd backend
go test ./...

# Tests avec examples
make run
# Puis dans l'interface, importer le dossier ./examples
```

## 🐛 Dépannage

### Le backend ne démarre pas

```bash
# Vérifier que PostgreSQL est démarré
make logs-db

# Vérifier les logs du backend
make logs-backend
```

### Le watcher ne détecte pas les fichiers

1. Vérifier que le dossier HH est correct dans Settings
2. Vérifier que le watcher est démarré
3. Vérifier les permissions du dossier

### Erreur d'import

Les logs d'import sont visibles dans l'onglet Import et dans les logs backend.

## 📝 TODO Future

- [ ] Tracking complet des adversaires
- [ ] Heatmaps préflop
- [ ] Replayer de mains avec visualisation
- [ ] Export des données (CSV, Excel)
- [ ] Filtres avancés
- [ ] Analyse ICM détaillée
- [ ] Graphiques de variance
- [ ] Comparaison avec d'autres joueurs

## 👤 Auteur

Développé pour le suivi personnel des Expressos Winamax.

## 📄 Licence

Usage personnel uniquement.

---

**Note**: Ce tracker est conçu pour un usage personnel d'analyse de résultats. Il utilise uniquement les données de vos propres hand histories.
