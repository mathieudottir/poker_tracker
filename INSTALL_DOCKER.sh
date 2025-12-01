#!/bin/bash

echo "🐳 Installation Docker - 2 minutes"
echo ""

# Installer Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

# Ajouter l'utilisateur au groupe docker
usermod -aG docker $USER

# Installer Docker Compose
curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

echo ""
echo "✅ Docker installé !"
echo ""
echo "⚠️  IMPORTANT: Redémarrez votre terminal/session puis lancez:"
echo "   ./start.sh"
echo ""
