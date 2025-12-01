import { useState, useRef } from 'react'
import { useRouter } from 'next/router'
import Link from 'next/link'
import { api } from '@/lib/api'

interface PageProps {
  user: any
  setUser: (user: any) => void
}

export async function getServerSideProps() {
  return { props: {} }
}

export default function Import({ user, setUser }: PageProps) {
  const router = useRouter()
  const fileInputRef = useRef<HTMLInputElement>(null)
  const [files, setFiles] = useState<File[]>([])
  const [uploading, setUploading] = useState(false)
  const [progress, setProgress] = useState(0)
  const [result, setResult] = useState<any>(null)
  const [error, setError] = useState('')

  if (!user) {
    router.push('/')
    return null
  }

  const handleFolderSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const fileList = Array.from(e.target.files)
      const txtFiles = fileList.filter((f) => f.name.endsWith('.txt'))
      setFiles(txtFiles)
      setResult(null)
      setError('')
    }
  }

  const handleUpload = async () => {
    if (files.length === 0) return

    try {
      setUploading(true)
      setError('')
      setResult(null)
      setProgress(0)

      // Upload files using API
      const data = await api.importFiles(user.id, files)
      setResult(data)
      setProgress(100)
      setFiles([])
      if (fileInputRef.current) {
        fileInputRef.current.value = ''
      }
    } catch (err: any) {
      setError(err.message)
    } finally {
      setUploading(false)
    }
  }

  const handleDownloadWatcher = () => {
    const watcherScript = `#!/usr/bin/env python3
"""
🃏 Winamax Expresso Tracker - Auto Watcher
Surveillance automatique de votre dossier HandHistory
"""

import os
import time
import requests
from pathlib import Path
from watchdog.observers import Observer
from watchdog.events import FileSystemEventHandler
import json

# Configuration personnalisée
SERVER_URL = "http://51.159.67.199:8080"
USER_ID = ${user.id}

class HandHistoryHandler(FileSystemEventHandler):
    def __init__(self):
        self.uploaded_files = set()
        self.cache_file = Path.home() / ".poker_tracker_cache.json"
        self.load_cache()

    def load_cache(self):
        if self.cache_file.exists():
            try:
                with open(self.cache_file, 'r') as f:
                    data = json.load(f)
                    self.uploaded_files = set(data.get("uploaded_files", []))
                print(f"📂 {len(self.uploaded_files)} fichiers déjà traités")
            except:
                pass

    def save_cache(self):
        try:
            with open(self.cache_file, 'w') as f:
                json.dump({"uploaded_files": list(self.uploaded_files)}, f)
        except:
            pass

    def on_created(self, event):
        if not event.is_directory and event.src_path.endswith('.txt'):
            print(f"\\n📥 Nouveau: {Path(event.src_path).name}")
            time.sleep(1)
            self.upload_file(event.src_path)

    def on_modified(self, event):
        if not event.is_directory and event.src_path.endswith('.txt'):
            if event.src_path not in self.uploaded_files:
                print(f"\\n📝 Modifié: {Path(event.src_path).name}")
                time.sleep(1)
                self.upload_file(event.src_path)

    def upload_file(self, file_path):
        if file_path in self.uploaded_files:
            return

        try:
            with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
                content = f.read()

            files = {'file': (Path(file_path).name, content, 'text/plain')}
            data = {'user_id': USER_ID}

            response = requests.post(
                f"{SERVER_URL}/api/import/file",
                files=files,
                data=data,
                timeout=30
            )

            if response.status_code == 200:
                result = response.json()
                print(f"✅ {result.get('tournaments_imported', 0)} tournois importés")
                self.uploaded_files.add(file_path)
                self.save_cache()
            else:
                print(f"❌ Erreur {response.status_code}")
        except Exception as e:
            print(f"❌ {str(e)}")

def main():
    print("="*60)
    print("🃏 WINAMAX EXPRESSO TRACKER - AUTO WATCHER")
    print("="*60)
    print(f"🌐 Serveur: {SERVER_URL}")
    print(f"👤 Utilisateur ID: {USER_ID}")

    print("\\n📁 Entrez le chemin de votre dossier HandHistory:")
    print(f"   Windows: C:\\\\Users\\\\VotreNom\\\\AppData\\\\Local\\\\Programs\\\\WinamaxPoker\\\\HandHistory")
    print(f"   Mac: ~/Library/Application Support/WinamaxPoker/HandHistory")

    watch_path = input("\\n📂 Chemin: ").strip().strip('"').strip("'")

    if not os.path.exists(watch_path):
        print(f"❌ Dossier introuvable: {watch_path}")
        return

    print(f"\\n✅ Dossier trouvé!")

    # Scanner fichiers existants
    txt_files = list(Path(watch_path).glob("*.txt"))
    print(f"🔍 {len(txt_files)} fichiers .txt trouvés")

    if txt_files:
        response = input("\\nUploader les fichiers existants ? (o/N): ").lower()
        if response in ['o', 'oui']:
            handler = HandHistoryHandler()
            print("\\n📤 Upload en cours...")
            for i, file_path in enumerate(txt_files, 1):
                print(f"[{i}/{len(txt_files)}] {file_path.name}")
                handler.upload_file(str(file_path))
                time.sleep(0.5)
            print("\\n✅ Upload terminé!")

    # Démarrer surveillance
    handler = HandHistoryHandler()
    observer = Observer()
    observer.schedule(handler, watch_path, recursive=False)
    observer.start()

    print("\\n" + "="*60)
    print("👀 SURVEILLANCE ACTIVE")
    print("="*60)
    print(f"📂 {watch_path}")
    print("🔄 Les nouveaux fichiers seront automatiquement uploadés")
    print("\\n⏸️  Ctrl+C pour arrêter")
    print("="*60)

    try:
        while True:
            time.sleep(1)
    except KeyboardInterrupt:
        print("\\n\\n🛑 Arrêt...")
        observer.stop()
        observer.join()
        print("👋 Au revoir!")

if __name__ == "__main__":
    main()
`

    const blob = new Blob([watcherScript], { type: 'text/plain' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'watcher.py'
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    URL.revokeObjectURL(url)
  }

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header */}
      <header className="bg-gray-800 border-b border-gray-700">
        <div className="container mx-auto px-4 py-4">
          <div className="flex justify-between items-center">
            <h1 className="text-2xl font-bold text-white">
              📥 Import Hand Histories
            </h1>
            <Link
              href="/dashboard"
              className="bg-gray-700 hover:bg-gray-600 text-white px-6 py-2 rounded-lg font-semibold transition-colors"
            >
              ← Retour
            </Link>
          </div>
        </div>
      </header>

      <div className="container mx-auto px-4 py-8 max-w-4xl">
        {/* Watcher Download */}
        <div className="bg-gradient-to-r from-green-900 to-green-800 rounded-xl p-6 mb-8 border border-green-700">
          <h2 className="text-2xl font-bold text-white mb-2">
            🤖 Surveillance Automatique (Recommandé)
          </h2>
          <p className="text-green-100 mb-4">
            Téléchargez le watcher pour upload automatique pendant que vous jouez !
          </p>
          <button
            onClick={handleDownloadWatcher}
            className="bg-white hover:bg-gray-100 text-green-900 px-6 py-3 rounded-lg font-bold transition-colors"
          >
            📥 Télécharger le Watcher Auto
          </button>
          <div className="mt-4 text-sm text-green-200">
            <p className="font-semibold mb-2">Installation:</p>
            <code className="bg-black/30 px-3 py-1 rounded block">
              pip install watchdog requests
            </code>
            <code className="bg-black/30 px-3 py-1 rounded block mt-2">
              python watcher.py
            </code>
          </div>
        </div>

        {/* Manual Import */}
        <div className="bg-gray-800 rounded-xl p-8 border border-gray-700">
          <h2 className="text-2xl font-bold text-white mb-4">
            📁 Import Manuel (Dossier)
          </h2>

          <div className="bg-blue-900/30 border border-blue-700 rounded-lg p-4 mb-6">
            <h3 className="font-semibold text-blue-200 mb-2">
              📖 Comment importer :
            </h3>
            <ol className="list-decimal list-inside text-blue-100 space-y-1 text-sm">
              <li>Cliquez sur "Sélectionner un dossier"</li>
              <li>Choisissez votre dossier Winamax HandHistory</li>
              <li>Le système trouvera automatiquement tous les fichiers .txt</li>
              <li>Cliquez sur "Importer" pour uploader</li>
            </ol>
            <p className="text-xs text-blue-300 mt-3">
              💡 Windows:{' '}
              <code className="bg-black/30 px-2 py-1 rounded">
                C:\Users\VotreNom\AppData\Local\Programs\WinamaxPoker\HandHistory
              </code>
            </p>
          </div>

          {/* Folder Selection */}
          <div className="mb-6">
            <input
              ref={fileInputRef}
              type="file"
              // @ts-ignore
              webkitdirectory=""
              directory=""
              multiple
              onChange={handleFolderSelect}
              className="hidden"
              id="folder-input"
            />
            <label
              htmlFor="folder-input"
              className="cursor-pointer block w-full"
            >
              <div className="border-2 border-dashed border-gray-600 rounded-lg p-12 text-center hover:border-green-500 transition-colors">
                <div className="text-6xl mb-4">📁</div>
                <div className="text-xl font-semibold text-white mb-2">
                  Sélectionner un dossier
                </div>
                <div className="text-gray-400">
                  Cliquez pour choisir votre dossier HandHistory
                </div>
              </div>
            </label>
          </div>

          {/* Files Preview */}
          {files.length > 0 && (
            <div className="bg-gray-700 rounded-lg p-4 mb-6">
              <div className="flex justify-between items-center mb-2">
                <h3 className="font-semibold text-white">
                  ✅ {files.length} fichiers .txt trouvés
                </h3>
                <button
                  onClick={() => {
                    setFiles([])
                    if (fileInputRef.current) {
                      fileInputRef.current.value = ''
                    }
                  }}
                  className="text-red-400 hover:text-red-300 text-sm"
                >
                  ✕ Annuler
                </button>
              </div>
              <div className="max-h-48 overflow-y-auto text-sm text-gray-300">
                {files.slice(0, 10).map((file, i) => (
                  <div key={i} className="py-1">
                    {file.name}
                  </div>
                ))}
                {files.length > 10 && (
                  <div className="text-gray-500 py-1">
                    ... et {files.length - 10} autres fichiers
                  </div>
                )}
              </div>
            </div>
          )}

          {/* Progress */}
          {uploading && (
            <div className="mb-6">
              <div className="flex justify-between text-sm text-gray-400 mb-2">
                <span>Upload en cours...</span>
                <span>{progress}%</span>
              </div>
              <div className="w-full bg-gray-700 rounded-full h-2">
                <div
                  className="bg-green-500 h-2 rounded-full transition-all duration-300"
                  style={{ width: `${progress}%` }}
                ></div>
              </div>
            </div>
          )}

          {/* Error */}
          {error && (
            <div className="bg-red-900/30 border border-red-700 text-red-200 px-4 py-3 rounded-lg mb-6">
              ❌ {error}
            </div>
          )}

          {/* Success */}
          {result && (
            <div className="bg-green-900/30 border border-green-700 text-green-200 px-4 py-3 rounded-lg mb-6">
              🎉 <strong>Import réussi !</strong>
              <br />
              📊 {result.count} tournois importés
            </div>
          )}

          {/* Upload Button */}
          <button
            onClick={handleUpload}
            disabled={files.length === 0 || uploading}
            className="w-full bg-green-600 hover:bg-green-700 disabled:bg-gray-600 disabled:cursor-not-allowed text-white font-bold py-4 rounded-lg transition-colors text-lg"
          >
            {uploading ? '⏳ Import en cours...' : `🚀 Importer ${files.length} fichiers`}
          </button>
        </div>
      </div>
    </div>
  )
}
