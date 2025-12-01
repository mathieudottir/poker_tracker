#!/usr/bin/env python3
"""
🃏 Winamax Expresso Tracker - Client Watcher
Surveille un dossier sur VOTRE PC et upload automatiquement les nouveaux fichiers
"""

import os
import time
import requests
from pathlib import Path
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
import json

# Configuration
SERVER_URL = "http://51.159.67.199:8080"  # Adresse de votre serveur
USER_ID = 1  # ID de l'utilisateur (Mathieu)

# Statistiques
stats = {
    "uploaded": 0,
    "errors": 0,
    "watching": False
}


class HandHistoryHandler(FileSystemEventHandler):
    """Gestionnaire d'événements pour les fichiers Hand History"""

    def __init__(self):
        self.uploaded_files = set()
        self.load_uploaded_cache()

    def load_uploaded_cache(self):
        """Charge la liste des fichiers déjà uploadés depuis le cache"""
        cache_file = Path.home() / ".poker_tracker_cache.json"
        if cache_file.exists():
            try:
                with open(cache_file, 'r') as f:
                    data = json.load(f)
                    self.uploaded_files = set(data.get("uploaded_files", []))
                print(f"📂 Cache chargé: {len(self.uploaded_files)} fichiers déjà traités")
            except Exception as e:
                print(f"⚠️  Erreur chargement cache: {e}")

    def save_uploaded_cache(self):
        """Sauvegarde la liste des fichiers uploadés dans le cache"""
        cache_file = Path.home() / ".poker_tracker_cache.json"
        try:
            with open(cache_file, 'w') as f:
                json.dump({"uploaded_files": list(self.uploaded_files)}, f)
        except Exception as e:
            print(f"⚠️  Erreur sauvegarde cache: {e}")

    def on_created(self, event):
        """Appelé quand un nouveau fichier est créé"""
        if event.is_directory:
            return

        if event.src_path.endswith('.txt'):
            print(f"\n📥 Nouveau fichier détecté: {Path(event.src_path).name}")
            time.sleep(1)  # Attendre que le fichier soit complètement écrit
            self.upload_file(event.src_path)

    def on_modified(self, event):
        """Appelé quand un fichier est modifié"""
        if event.is_directory:
            return

        if event.src_path.endswith('.txt') and event.src_path not in self.uploaded_files:
            print(f"\n📝 Fichier modifié: {Path(event.src_path).name}")
            time.sleep(1)  # Attendre que le fichier soit complètement écrit
            self.upload_file(event.src_path)

    def upload_file(self, file_path):
        """Upload un fichier vers le serveur"""
        try:
            # Vérifier si déjà uploadé
            if file_path in self.uploaded_files:
                print(f"⏭️  Fichier déjà uploadé, ignoré")
                return

            # Lire le contenu du fichier
            with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
                content = f.read()

            # Préparer la requête
            files = {
                'file': (Path(file_path).name, content, 'text/plain')
            }
            data = {
                'user_id': USER_ID
            }

            # Envoyer au serveur
            print(f"📤 Upload vers {SERVER_URL}/api/import/file...")
            response = requests.post(
                f"{SERVER_URL}/api/import/file",
                files=files,
                data=data,
                timeout=30
            )

            if response.status_code == 200:
                result = response.json()
                print(f"✅ Succès! {result.get('tournaments_imported', 0)} tournois importés")
                stats["uploaded"] += 1
                self.uploaded_files.add(file_path)
                self.save_uploaded_cache()
            else:
                print(f"❌ Erreur {response.status_code}: {response.text}")
                stats["errors"] += 1

        except Exception as e:
            print(f"❌ Erreur upload: {str(e)}")
            stats["errors"] += 1


def scan_existing_files(watch_path, handler):
    """Scanne les fichiers existants au démarrage"""
    print(f"\n🔍 Scan des fichiers existants dans {watch_path}...")

    txt_files = list(Path(watch_path).glob("*.txt"))
    print(f"📊 {len(txt_files)} fichiers .txt trouvés")

    # Demander à l'utilisateur s'il veut uploader les fichiers existants
    if txt_files:
        response = input("\n❓ Voulez-vous uploader tous les fichiers existants ? (o/N): ").lower()
        if response == 'o' or response == 'oui':
            print("\n📤 Upload des fichiers existants...")
            for i, file_path in enumerate(txt_files, 1):
                print(f"\n[{i}/{len(txt_files)}] {file_path.name}")
                handler.upload_file(str(file_path))
                time.sleep(0.5)  # Petit délai entre chaque upload

            print(f"\n✅ Scan terminé: {stats['uploaded']} fichiers uploadés, {stats['errors']} erreurs")
        else:
            print("⏭️  Fichiers existants ignorés, seuls les nouveaux fichiers seront uploadés")


def main():
    """Fonction principale"""
    print("=" * 60)
    print("🃏 WINAMAX EXPRESSO TRACKER - CLIENT WATCHER")
    print("=" * 60)
    print(f"\n🌐 Serveur: {SERVER_URL}")
    print(f"👤 Utilisateur: {USER_ID} (Mathieu)")

    # Demander le chemin du dossier à surveiller
    print("\n📁 Entrez le chemin de votre dossier Winamax HandHistory:")
    print("   Exemple Windows: C:\\Users\\VotreNom\\AppData\\Local\\Winamax\\HandHistory")
    print("   Exemple Mac: /Users/VotreNom/Library/Application Support/Winamax/HandHistory")

    watch_path = input("\n📂 Chemin: ").strip().strip('"').strip("'")

    # Vérifier que le dossier existe
    if not os.path.exists(watch_path):
        print(f"❌ Erreur: Le dossier '{watch_path}' n'existe pas !")
        return

    if not os.path.isdir(watch_path):
        print(f"❌ Erreur: '{watch_path}' n'est pas un dossier !")
        return

    print(f"\n✅ Dossier trouvé: {watch_path}")

    # Créer le gestionnaire et l'observateur
    event_handler = HandHistoryHandler()
    observer = Observer()
    observer.schedule(event_handler, watch_path, recursive=False)

    # Scanner les fichiers existants
    scan_existing_files(watch_path, event_handler)

    # Démarrer la surveillance
    observer.start()
    stats["watching"] = True

    print("\n" + "=" * 60)
    print("👀 SURVEILLANCE ACTIVE")
    print("=" * 60)
    print(f"📂 Dossier surveillé: {watch_path}")
    print(f"🔄 Les nouveaux fichiers .txt seront automatiquement uploadés")
    print(f"\n📊 Statistiques:")
    print(f"   ✅ Fichiers uploadés: {stats['uploaded']}")
    print(f"   ❌ Erreurs: {stats['errors']}")
    print("\n⏸️  Appuyez sur Ctrl+C pour arrêter")
    print("=" * 60)

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        print("\n\n🛑 Arrêt de la surveillance...")
        observer.stop()
        observer.join()
        print(f"\n📊 Statistiques finales:")
        print(f"   ✅ Fichiers uploadés: {stats['uploaded']}")
        print(f"   ❌ Erreurs: {stats['errors']}")
        print("\n👋 Au revoir!")


if __name__ == "__main__":
    main()
