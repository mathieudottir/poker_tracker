# 🃏 Watcher Client - Surveillance Automatique depuis votre PC

Ce programme surveille votre dossier Winamax HandHistory **sur votre PC** et upload automatiquement les nouveaux fichiers vers le serveur.

## 📥 Installation

### Windows

1. Ouvrez PowerShell ou CMD
2. Installez Python si nécessaire: https://www.python.org/downloads/
3. Installez les dépendances:
```bash
cd C:\chemin\vers\poker_tracker\client
pip install -r requirements.txt
```

### Mac / Linux

1. Ouvrez le Terminal
2. Installez les dépendances:
```bash
cd /chemin/vers/poker_tracker/client
pip3 install -r requirements.txt
```

## 🚀 Utilisation

### Windows
```bash
python watcher_client.py
```

### Mac / Linux
```bash
python3 watcher_client.py
```

## 📋 Ce qui se passe

1. Le programme vous demande le chemin de votre dossier Winamax HandHistory
2. Il scanne tous les fichiers existants et vous demande si vous voulez les uploader
3. Il reste actif et surveille les nouveaux fichiers
4. Chaque nouveau fichier .txt est automatiquement uploadé vers le serveur

## 📂 Trouver votre dossier HandHistory

### Windows
```
C:\Users\VotreNom\AppData\Local\Winamax\HandHistory
```

### Mac
```
/Users/VotreNom/Library/Application Support/Winamax/HandHistory
```

### Linux (Wine)
```
~/.wine/drive_c/users/VotreNom/Local Settings/Application Data/Winamax/HandHistory
```

## 💡 Conseils

- **Laissez le programme tourner pendant que vous jouez** - vos tournois seront uploadés en temps réel
- **Cache intelligent** - les fichiers déjà uploadés sont mémorisés (`.poker_tracker_cache.json` dans votre dossier home)
- **Ctrl+C pour arrêter** - le programme se ferme proprement

## 🔧 Configuration

Par défaut, le serveur est configuré sur `http://51.159.67.199:8080`

Pour changer l'adresse du serveur, éditez `watcher_client.py` ligne 15:
```python
SERVER_URL = "http://VOTRE_SERVEUR:8080"
```

## 📊 Statistiques

Le programme affiche en temps réel:
- ✅ Nombre de fichiers uploadés
- ❌ Nombre d'erreurs
- 📂 Dossier surveillé
- 🔄 État de la surveillance

## ❓ Problèmes courants

### "Le dossier n'existe pas"
Vérifiez le chemin exact de votre dossier Winamax HandHistory

### "Connection refused"
Vérifiez que le serveur est accessible depuis votre PC:
```bash
ping 51.159.67.199
curl http://51.159.67.199:8080/health
```

### "Module 'watchdog' not found"
Installez les dépendances:
```bash
pip install -r requirements.txt
```
