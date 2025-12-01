# 🃏 Installation sur Serveur Ubuntu Vierge

## Prérequis
- Ubuntu 20.04 ou plus récent
- Accès SSH au serveur (51.159.67.199)
- Accès root ou sudo

---

## 📋 ÉTAPES D'INSTALLATION

### 1️⃣ Se connecter au serveur

```bash
ssh root@51.159.67.199
# ou
ssh votre_user@51.159.67.199
```

### 2️⃣ Installer Git (si pas déjà installé)

```bash
apt-get update
apt-get install -y git
```

### 3️⃣ Cloner le projet

```bash
cd /root  # ou /home/votre_user
git clone https://github.com/mathieudottir/poker_tracker.git
cd poker_tracker
git checkout claude/winamax-expresso-tracker-01GjueeubKtruT2V4BRboKQW
```

### 4️⃣ Lancer l'installation automatique

```bash
chmod +x INSTALL_FRESH.sh
./INSTALL_FRESH.sh
```

Le script va installer automatiquement :
- PostgreSQL 16
- Go 1.21
- Python 3 et pip
- Streamlit et dépendances
- Toutes les configurations nécessaires

⏱️ Durée : ~5-10 minutes

### 5️⃣ Démarrer le système

```bash
chmod +x start-simple.sh
./start-simple.sh
```

Vous devriez voir :
```
✅✅✅ SYSTÈME DÉMARRÉ ! ✅✅✅

📍 Ouvrez votre navigateur:
   👉 Local:  http://localhost:8501
   👉 Remote: http://51.159.67.199:8501
```

### 6️⃣ Ouvrir dans le navigateur

Allez sur : **http://51.159.67.199:8501**

Cliquez sur **"🔧 Login as Mathieu (DEV)"**

---

## 🔥 FIREWALL - IMPORTANT !

Si rien ne s'affiche, ouvrez les ports :

```bash
# Si UFW (Ubuntu Firewall)
ufw allow 8501/tcp
ufw allow 8080/tcp

# Si iptables
iptables -A INPUT -p tcp --dport 8501 -j ACCEPT
iptables -A INPUT -p tcp --dport 8080 -j ACCEPT
iptables-save
```

---

## 🛑 Arrêter le système

```bash
pkill -f 'go run'
pkill -f streamlit
```

---

## 🔍 Vérifier que tout tourne

```bash
# Voir les processus
ps aux | grep -E 'streamlit|go run'

# Voir les ports ouverts
netstat -tlnp | grep -E '8080|8501'

# Voir les logs
tail -f logs/backend.log
tail -f logs/frontend.log
```

---

## ⚠️ Problèmes courants

### PostgreSQL ne démarre pas
```bash
systemctl status postgresql
systemctl start postgresql
```

### Port déjà utilisé
```bash
# Tuer les anciens processus
pkill -f streamlit
pkill -f 'go run'
# Puis relancer
./start-simple.sh
```

### Erreur "permission denied"
```bash
chmod +x start-simple.sh
chmod +x INSTALL_FRESH.sh
```

### Backend ne se connecte pas à la DB
```bash
# Vérifier que postgres password est configuré
sudo -u postgres psql -c "ALTER USER postgres PASSWORD 'postgres';"
```

---

## 📞 Support

Si problème, envoyez :
```bash
# Statut système
./start-simple.sh 2>&1 | tee install.log

# Logs
cat logs/backend.log
cat logs/frontend.log
```
