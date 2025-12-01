# 📥 Comment récupérer le code

## Si vous êtes sur une autre machine

```bash
# Cloner le repository
git clone https://github.com/mathieudottir/poker_tracker.git
cd poker_tracker

# Récupérer la branche
git checkout claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW

# Lancer le projet
./start.sh
```

## Si vous êtes déjà dans le repository

```bash
# Récupérer toutes les branches
git fetch origin

# Basculer sur la branche
git checkout claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW

# Mettre à jour (si besoin)
git pull origin claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW
```

## Si vous voulez fusionner avec votre branche main

```bash
# D'abord, récupérer la branche
git checkout claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW

# Puis fusionner avec main (ou master)
git checkout main
git merge claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW
git push origin main
```

## Vérifier que vous avez le bon code

```bash
# Vérifier la branche actuelle
git branch

# Vérifier les derniers commits
git log --oneline -5

# Vous devriez voir le commit "Complete Winamax Expresso Tracker Implementation"
```

## Structure des fichiers que vous devriez avoir

```
poker_tracker/
├── backend/
│   ├── cmd/server/main.go
│   ├── internal/
│   ├── pkg/
│   ├── migrations/
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── app.py
│   └── requirements.txt
├── examples/                    (15 fichiers HH)
├── docker-compose.yml
├── Dockerfile.backend
├── Dockerfile.frontend
├── Makefile
├── README.md
├── QUICKSTART.md
├── config.yaml
├── start.sh
├── .gitignore
└── .env.example
```

## Une fois le code récupéré

```bash
# Lancer le système
./start.sh

# OU
make run

# Puis ouvrir http://localhost:8501
```

## En cas de problème

```bash
# Forcer la mise à jour
git fetch origin
git reset --hard origin/claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW

# Vérifier le statut
git status
```
