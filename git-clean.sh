#!/bin/bash

echo "🧹 Nettoyage Git Complet"
echo ""

# Annuler tout merge/rebase en cours
git merge --abort 2>/dev/null
git rebase --abort 2>/dev/null

# Récupérer les dernières modifications du remote
echo "📥 Récupération des modifications..."
git fetch origin

# Réinitialiser au dernier commit du remote
echo "🔄 Réinitialisation à l'état distant..."
git reset --hard origin/claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW

# Nettoyer les fichiers non suivis
echo "🗑️  Nettoyage des fichiers non suivis..."
git clean -fd

echo ""
echo "✅ Git est maintenant propre et à jour !"
echo ""
git status
