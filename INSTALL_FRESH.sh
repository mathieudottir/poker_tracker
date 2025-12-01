#!/bin/bash
#
# 🃏 Installation complète Winamax Expresso Tracker
# Pour Ubuntu 20.04+ sur serveur vierge
#

set -e  # Arrêt si erreur

echo "════════════════════════════════════════════════"
echo "🃏 INSTALLATION WINAMAX EXPRESSO TRACKER"
echo "════════════════════════════════════════════════"
echo ""

# 1. Mise à jour système
echo "📦 [1/6] Mise à jour du système..."
sudo apt-get update -qq

# 2. Installation PostgreSQL
echo "🐘 [2/6] Installation PostgreSQL..."
sudo apt-get install -y postgresql postgresql-contrib > /dev/null

# Démarrer PostgreSQL
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Configurer PostgreSQL
echo "🔧 Configuration PostgreSQL..."
sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';" 2>/dev/null || true

# Configuration pg_hba.conf pour permettre connexions locales
PG_VERSION=$(ls /etc/postgresql/ | head -1)
sudo sed -i "s/local   all             postgres                                peer/local   all             postgres                                trust/" /etc/postgresql/$PG_VERSION/main/pg_hba.conf
sudo sed -i "s/local   all             all                                     peer/local   all             all                                     trust/" /etc/postgresql/$PG_VERSION/main/pg_hba.conf
sudo systemctl restart postgresql

echo "✅ PostgreSQL configuré"

# 3. Installation Go
echo "🐹 [3/6] Installation Go 1.21..."
if ! command -v go &> /dev/null; then
    cd /tmp
    wget -q https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
    sudo rm -rf /usr/local/go
    sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    export PATH=$PATH:/usr/local/go/bin
    echo "✅ Go installé"
else
    echo "✅ Go déjà installé"
fi

# 4. Installation Python et pip
echo "🐍 [4/6] Installation Python et dépendances..."
sudo apt-get install -y python3 python3-pip git curl > /dev/null
echo "✅ Python installé"

# 5. Installation Streamlit
echo "🎨 [5/6] Installation Streamlit..."
pip3 install -q streamlit plotly requests pandas
echo "✅ Streamlit installé"

# 6. Cloner/Mettre à jour le projet
echo "📥 [6/6] Récupération du code..."
cd /home/$(whoami)

if [ -d "poker_tracker" ]; then
    echo "⚠️  Dossier poker_tracker existe, mise à jour..."
    cd poker_tracker
    git fetch origin
    git checkout claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW
    git pull origin claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW
else
    echo "📦 Clonage du projet..."
    # À remplacer par votre URL Git
    echo "⚠️  Vous devez cloner le projet manuellement:"
    echo "    git clone <VOTRE_REPO_URL> poker_tracker"
    echo "    cd poker_tracker"
    echo "    git checkout claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW"
    exit 1
fi

# Créer les dossiers nécessaires
mkdir -p logs examples

# Télécharger les dépendances Go
echo "📦 Installation dépendances Go..."
cd backend
go mod download > /dev/null 2>&1
cd ..

echo ""
echo "════════════════════════════════════════════════"
echo "✅ INSTALLATION TERMINÉE !"
echo "════════════════════════════════════════════════"
echo ""
echo "🚀 Pour démarrer le système:"
echo "   cd /home/$(whoami)/poker_tracker"
echo "   ./start-simple.sh"
echo ""
echo "🌐 Accès web:"
echo "   http://51.159.67.199:8501"
echo ""
echo "════════════════════════════════════════════════"
