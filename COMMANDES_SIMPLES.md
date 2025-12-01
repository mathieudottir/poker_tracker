# 🎯 Commandes Simples - Poker Tracker

## 🚀 Démarrer l'Application

```bash
./start-simple.sh
```

**C'est tout !** Ouvre ensuite http://localhost:8501 et clique sur "Login as Mathieu"

---

## 🔧 Si Problèmes Git

```bash
./git-clean.sh
```

Ce script nettoie tous les conflits et remet Git à jour automatiquement.

---

## 🛑 Arrêter l'Application

```bash
pkill -f 'go run' && pkill -f streamlit
```

---

## 📦 Installer Docker (optionnel, plus simple)

```bash
./INSTALL_DOCKER.sh
```

Redémarre ton terminal, puis :

```bash
./start.sh
```

---

## 📋 Résumé

| Action | Commande |
|--------|----------|
| **Lancer** | `./start-simple.sh` |
| **Arrêter** | `pkill -f 'go run' && pkill -f streamlit` |
| **Nettoyer Git** | `./git-clean.sh` |
| **Installer Docker** | `./INSTALL_DOCKER.sh` |
| **Avec Docker** | `./start.sh` |

---

## 🆘 Aide

### L'app ne démarre pas ?

1. Lance `./git-clean.sh` pour mettre à jour
2. Relance `./start-simple.sh`

### Erreur Git ?

Lance juste `./git-clean.sh` - ça règle tout automatiquement !

### Rien ne fonctionne ?

Lance Docker à la place :
```bash
./INSTALL_DOCKER.sh
# Redémarre ton terminal
./start.sh
```

---

**Note** : PostgreSQL doit tourner (il tourne déjà sur ta machine)
