# 🚀 Quick Start Guide

## Démarrage Rapide (5 minutes)

### 1. Lancer le système

```bash
./start.sh
```

OU

```bash
make run
```

### 2. Accéder à l'interface

Ouvrir votre navigateur: **http://localhost:8501**

### 3. Se connecter

Cliquer sur le bouton: **"🔧 Login as Mathieu (DEV)"**

### 4. Importer des tournois de test

1. Aller dans l'onglet **"Import"**
2. Dans le champ "Directory Path", vérifier que `./examples` est affiché
3. Cliquer sur **"Import from Directory"**
4. Attendre quelques secondes

✅ Les tournois d'exemple sont maintenant importés !

### 5. Voir les résultats

1. Aller dans l'onglet **"My Results"**
2. Admirer vos courbes de bankroll et EV
3. Explorer les statistiques

### 6. Configuration pour vos propres HH

1. Aller dans **"Settings"**
2. Modifier:
   - **Winamax Player Name**: Votre pseudo Winamax
   - **Hand History Directory**: Le chemin vers votre dossier HH Winamax
   - **Winamax Status**: Votre statut actuel (pour le rakeback)
3. Cliquer sur **"Save Settings"**
4. Retourner dans **"Import"**
5. Cliquer sur **"Start Watcher"**

Le système importera automatiquement vos nouveaux tournois en temps réel !

## 🎮 Fonctionnalités Principales

### Import
- **Watcher temps réel**: Importe automatiquement les nouveaux tournois
- **Import manuel**: Importer un dossier entier en un clic
- **Logs**: Voir l'historique des imports

### Résultats
- **Courbe de bankroll**: Évolution de vos profits
- **Courbe EV**: Comparaison EV théorique vs résultats réels
- **Distribution multiplicateurs**: Voir vos multiplicateurs obtenus
- **Tableau détaillé**: Tous vos tournois en détail

### Statistiques
- **ROI global**: Net ROI et EV ROI
- **Rake & Rakeback**: Combien vous avez payé et récupéré
- **Par buy-in**: Vos résultats détaillés pour chaque limite

### Paramètres
- **Profil**: Nom de joueur et statut Winamax
- **Dossier HH**: Où trouver vos hand histories
- **Reset**: Supprimer toutes les données si besoin

## 📊 Comprendre les Métriques

### ROI (Return on Investment)
Votre rentabilité en pourcentage:
- **Net ROI**: Basé sur vos résultats réels
- **EV ROI**: Basé sur la valeur attendue théorique

### EV (Expected Value)
Ce que vous "devriez" gagner statistiquement selon les probabilités Winamax.

Si votre courbe réelle est **au-dessus** de l'EV: vous êtes chanceux ! 🍀
Si elle est **en-dessous**: malchance temporaire, ça reviendra 💪

### Multiplicateurs
Le système compare vos multiplicateurs obtenus avec les probabilités officielles Winamax pour voir si vous êtes dans la variance normale.

### Rakeback
Calculé automatiquement selon votre statut Winamax (Bronze, Silver, Gold, etc.)

## 🛠️ Commandes Utiles

```bash
# Voir les logs en temps réel
make logs

# Arrêter le système
make stop

# Redémarrer
make run

# Tout nettoyer et repartir de zéro
make clean
make run
```

## ❓ FAQ

### Le watcher ne détecte pas mes nouveaux tournois

1. Vérifier que le chemin du dossier HH est correct dans Settings
2. Vérifier que le watcher est bien démarré (voyant vert)
3. Vérifier les permissions du dossier

### Mes tournois n'apparaissent pas

1. Vérifier que vous avez bien les fichiers `.txt` ET `_summary.txt`
2. Vérifier que le nom du joueur dans Settings correspond exactement à votre pseudo Winamax
3. Regarder les logs d'import pour voir les erreurs

### Je veux recommencer à zéro

1. Aller dans **Settings**
2. Cliquer sur **"Reset My Database"**
3. Confirmer

### Comment arrêter le système ?

```bash
make stop
# ou
docker-compose down
```

## 🎯 Prochaines Étapes

1. Importer tous vos anciens tournois
2. Configurer le watcher pour les nouveaux
3. Analyser vos résultats par buy-in
4. Optimiser votre jeu selon les stats
5. Suivre votre progression !

---

**Bon run ! 🚀**
