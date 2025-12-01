#!/bin/bash
# Watcher Client pour Mac/Linux

echo "================================"
echo "Winamax Expresso Tracker Watcher"
echo "================================"
echo ""

# Vérifier Python
if ! command -v python3 &> /dev/null; then
    echo "❌ ERREUR: Python 3 n'est pas installé"
    exit 1
fi

# Installer les dépendances si nécessaire
echo "📦 Installation des dépendances..."
pip3 install -q -r requirements.txt

# Lancer le watcher
echo ""
python3 watcher_client.py
