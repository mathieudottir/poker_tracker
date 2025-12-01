import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import Link from 'next/link'
import { api } from '@/lib/api'

interface PageProps {
  user: any
  setUser: (user: any) => void
}

const WINA_STATUSES = [
  { name: 'Aluminium', rakeback: 0 },
  { name: 'Bronze', rakeback: 20 },
  { name: 'Argent', rakeback: 25 },
  { name: 'Or', rakeback: 30 },
  { name: 'Platine', rakeback: 35 },
  { name: 'Diamond 1', rakeback: 40 },
  { name: 'Diamond 2', rakeback: 44 },
  { name: 'Diamond 3', rakeback: 48 },
  { name: 'Diamond 4', rakeback: 52 },
]

export default function Settings({ user, setUser }: PageProps) {
  const router = useRouter()
  const [loading, setLoading] = useState(false)
  const [message, setMessage] = useState('')
  const [error, setError] = useState('')
  const [playerName, setPlayerName] = useState('')
  const [winaStatus, setWinaStatus] = useState('Aluminium')
  const [hhDirectory, setHhDirectory] = useState('')
  const [confirmDelete, setConfirmDelete] = useState(false)

  useEffect(() => {
    if (!user) {
      router.push('/')
      return
    }
    setPlayerName(user.player_name || '')
    setWinaStatus(user.wina_status || 'Aluminium')
    setHhDirectory(user.hh_directory || '')
  }, [user, router])

  // Return null during SSR or if no user
  if (!user) return null

  const handleUpdateSettings = async (e: React.FormEvent) => {
    e.preventDefault()
    try {
      setLoading(true)
      setError('')
      setMessage('')

      await api.updateSettings(user.id, playerName, winaStatus, hhDirectory, user.dev_mode)

      // Update local user
      const updatedUser = { ...user, player_name: playerName, wina_status: winaStatus, hh_directory: hhDirectory }
      setUser(updatedUser)
      localStorage.setItem('user', JSON.stringify(updatedUser))

      setMessage('Paramètres mis à jour avec succès !')
    } catch (err: any) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  const handleDeleteData = async () => {
    if (!confirmDelete) {
      setError('Veuillez cocher la case de confirmation')
      return
    }

    try {
      setLoading(true)
      setError('')
      setMessage('')

      await api.deleteUserData(user.id)

      setMessage('Toutes vos données de tournois ont été supprimées !')
      setConfirmDelete(false)

      // Refresh page after 2 seconds
      setTimeout(() => {
        router.reload()
      }, 2000)
    } catch (err: any) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header */}
      <header className="bg-gray-800 border-b border-gray-700">
        <div className="container mx-auto px-4 py-4">
          <div className="flex justify-between items-center">
            <div>
              <h1 className="text-2xl font-bold text-white">
                ⚙️ Paramètres
              </h1>
              <p className="text-gray-400">
                {user.player_name} ({user.wina_status})
              </p>
            </div>
            <div className="flex gap-4">
              <Link
                href="/dashboard"
                className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-semibold transition-colors"
              >
                ← Dashboard
              </Link>
              <Link
                href="/import"
                className="bg-green-600 hover:bg-green-700 text-white px-6 py-2 rounded-lg font-semibold transition-colors"
              >
                📥 Import
              </Link>
            </div>
          </div>
        </div>
      </header>

      <div className="container mx-auto px-4 py-8 max-w-4xl">
        {message && (
          <div className="bg-green-500/20 border border-green-500 text-green-200 px-4 py-3 rounded-lg mb-6">
            {message}
          </div>
        )}

        {error && (
          <div className="bg-red-500/20 border border-red-500 text-red-200 px-4 py-3 rounded-lg mb-6">
            {error}
          </div>
        )}

        {/* User Settings */}
        <div className="bg-gray-800 rounded-xl p-6 mb-8 border border-gray-700">
          <h2 className="text-xl font-bold text-white mb-4">👤 Informations Utilisateur</h2>

          <form onSubmit={handleUpdateSettings} className="space-y-4">
            <div>
              <label className="block text-white text-sm font-medium mb-2">
                Nom de joueur Winamax
              </label>
              <input
                type="text"
                value={playerName}
                onChange={(e) => setPlayerName(e.target.value)}
                className="w-full px-4 py-3 rounded-lg bg-gray-700 border border-gray-600 text-white focus:outline-none focus:ring-2 focus:ring-green-500"
                placeholder="SIGRIDOTTIR"
                required
              />
            </div>

            <div>
              <label className="block text-white text-sm font-medium mb-2">
                Statut Winamax (pour le calcul du rakeback)
              </label>
              <select
                value={winaStatus}
                onChange={(e) => setWinaStatus(e.target.value)}
                className="w-full px-4 py-3 rounded-lg bg-gray-700 border border-gray-600 text-white focus:outline-none focus:ring-2 focus:ring-green-500"
              >
                {WINA_STATUSES.map((status) => (
                  <option key={status.name} value={status.name}>
                    {status.name} - {status.rakeback}% rakeback
                  </option>
                ))}
              </select>
              <p className="text-gray-400 text-sm mt-2">
                Le statut détermine le pourcentage de rakeback utilisé pour les calculs d'EV
              </p>
            </div>

            <div>
              <label className="block text-white text-sm font-medium mb-2">
                Répertoire Hand History (optionnel)
              </label>
              <input
                type="text"
                value={hhDirectory}
                onChange={(e) => setHhDirectory(e.target.value)}
                className="w-full px-4 py-3 rounded-lg bg-gray-700 border border-gray-600 text-white focus:outline-none focus:ring-2 focus:ring-green-500"
                placeholder="/chemin/vers/dossier/hand-histories"
              />
            </div>

            <button
              type="submit"
              disabled={loading}
              className="w-full bg-green-600 hover:bg-green-700 text-white font-semibold py-3 rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {loading ? 'Enregistrement...' : '💾 Enregistrer les paramètres'}
            </button>
          </form>
        </div>

        {/* Danger Zone */}
        <div className="bg-red-900/20 rounded-xl p-6 border border-red-700">
          <h2 className="text-xl font-bold text-red-400 mb-4">⚠️ Zone dangereuse</h2>

          <div className="bg-gray-800 rounded-lg p-4 mb-4">
            <h3 className="text-white font-semibold mb-2">Vider toutes les données de tournois</h3>
            <p className="text-gray-400 text-sm mb-4">
              Cette action supprimera tous vos tournois et mains importés. Vos paramètres utilisateur seront conservés.
              Cette action est irréversible !
            </p>

            <div className="mb-4">
              <label className="flex items-center text-white">
                <input
                  type="checkbox"
                  checked={confirmDelete}
                  onChange={(e) => setConfirmDelete(e.target.checked)}
                  className="mr-2 h-4 w-4"
                />
                <span className="text-sm">
                  Je comprends que cette action est irréversible et que toutes mes données de tournois seront supprimées
                </span>
              </label>
            </div>

            <button
              onClick={handleDeleteData}
              disabled={loading || !confirmDelete}
              className="bg-red-600 hover:bg-red-700 text-white font-semibold py-3 px-6 rounded-lg transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {loading ? 'Suppression...' : '🗑️ Vider toutes les données'}
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}
