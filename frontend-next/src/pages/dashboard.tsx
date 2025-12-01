import { useState, useEffect } from 'react'
import { useRouter } from 'next/router'
import Link from 'next/link'
import { api, Tournament } from '@/lib/api'
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
  BarChart,
  Bar,
} from 'recharts'
import { format } from 'date-fns'

interface PageProps {
  user: any
  setUser: (user: any) => void
}

export default function Dashboard({ user, setUser }: PageProps) {
  const router = useRouter()
  const [tournaments, setTournaments] = useState<Tournament[]>([])
  const [loading, setLoading] = useState(true)
  const [startDate, setStartDate] = useState('')
  const [endDate, setEndDate] = useState('')

  useEffect(() => {
    if (!user) {
      router.push('/')
      return
    }
    loadData()
  }, [user, router])

  const loadData = async () => {
    try {
      setLoading(true)
      const data = await api.getTournaments(user.id)
      setTournaments(data)

      if (data.length > 0) {
        const dates = data.map((t) => new Date(t.start_time))
        setStartDate(format(new Date(Math.min(...dates.map((d) => d.getTime()))), 'yyyy-MM-dd'))
        setEndDate(format(new Date(Math.max(...dates.map((d) => d.getTime()))), 'yyyy-MM-dd'))
      }
    } catch (err) {
      console.error('Error loading tournaments:', err)
    } finally {
      setLoading(false)
    }
  }

  const handleLogout = () => {
    setUser(null)
    localStorage.removeItem('user')
    router.push('/')
  }

  if (!user) return null

  // Filter tournaments by date range
  const filteredTournaments = tournaments.filter((t) => {
    if (!startDate || !endDate) return true
    const tournamentDate = new Date(t.start_time)
    return (
      tournamentDate >= new Date(startDate) &&
      tournamentDate <= new Date(endDate + 'T23:59:59')
    )
  })

  // Calculate stats
  const totalTournaments = filteredTournaments.length
  const totalProfit = filteredTournaments.reduce(
    (sum, t) => sum + t.net_result_cents / 100,
    0
  )
  const totalRake = filteredTournaments.reduce(
    (sum, t) => sum + t.rake_cents / 100,
    0
  )
  const totalEV = filteredTournaments.reduce(
    (sum, t) => sum + t.ev_cents / 100,
    0
  )

  // Rakeback calculation
  const rakebackPct = {
    Aluminium: 0,
    Bronze: 10,
    Argent: 15,
    Or: 20,
    Platine: 25,
    Diamant: 30,
    'Red Diamond': 33,
  }[user.status || 'Aluminium'] || 0

  const totalRakeback = (totalRake * rakebackPct) / 100
  const profitWithRakeback = totalProfit + totalRakeback

  // Prepare chart data
  const sortedTournaments = [...filteredTournaments].sort(
    (a, b) => new Date(a.start_time).getTime() - new Date(b.start_time).getTime()
  )

  let cumulative = 0
  let cumulativeWithRB = 0
  let cumulativeEV = 0

  const chartData = sortedTournaments.map((t, i) => {
    cumulative += t.net_result_cents / 100
    const rakeback = (t.rake_cents / 100) * (rakebackPct / 100)
    cumulativeWithRB += t.net_result_cents / 100 + rakeback
    cumulativeEV += t.ev_cents / 100

    return {
      tournament: i + 1,
      profit: Number(cumulative.toFixed(2)),
      profitWithRB: Number(cumulativeWithRB.toFixed(2)),
      ev: Number(cumulativeEV.toFixed(2)),
      date: format(new Date(t.start_time), 'dd/MM HH:mm'),
    }
  })

  // Multiplier distribution
  const multiplierCounts: Record<number, number> = {}
  filteredTournaments.forEach((t) => {
    multiplierCounts[t.multiplier] = (multiplierCounts[t.multiplier] || 0) + 1
  })

  const multiplierData = Object.entries(multiplierCounts)
    .map(([mult, count]) => ({
      multiplier: `${mult}x`,
      count,
    }))
    .sort((a, b) => parseInt(a.multiplier) - parseInt(b.multiplier))

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-900 flex items-center justify-center">
        <div className="text-white text-xl">Chargement...</div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-900">
      {/* Header */}
      <header className="bg-gray-800 border-b border-gray-700">
        <div className="container mx-auto px-4 py-4">
          <div className="flex justify-between items-center">
            <div>
              <h1 className="text-2xl font-bold text-white">
                🃏 Winamax Expresso Tracker
              </h1>
              <p className="text-gray-400">
                Bienvenue, {user.player_name} ({user.status})
              </p>
            </div>
            <div className="flex gap-4">
              <Link
                href="/import"
                className="bg-green-600 hover:bg-green-700 text-white px-6 py-2 rounded-lg font-semibold transition-colors"
              >
                📥 Import
              </Link>
              <button
                onClick={handleLogout}
                className="bg-red-600 hover:bg-red-700 text-white px-6 py-2 rounded-lg font-semibold transition-colors"
              >
                Déconnexion
              </button>
            </div>
          </div>
        </div>
      </header>

      <div className="container mx-auto px-4 py-8">
        {/* Date Range Filter */}
        <div className="bg-gray-800 rounded-xl p-6 mb-8 border border-gray-700">
          <h2 className="text-xl font-bold text-white mb-4">🗓️ Filtres</h2>
          <div className="grid grid-cols-2 gap-4">
            <div>
              <label className="block text-gray-300 text-sm font-medium mb-2">
                Date de début
              </label>
              <input
                type="date"
                value={startDate}
                onChange={(e) => setStartDate(e.target.value)}
                className="w-full px-4 py-2 rounded-lg bg-gray-700 border border-gray-600 text-white focus:outline-none focus:ring-2 focus:ring-green-500"
              />
            </div>
            <div>
              <label className="block text-gray-300 text-sm font-medium mb-2">
                Date de fin
              </label>
              <input
                type="date"
                value={endDate}
                onChange={(e) => setEndDate(e.target.value)}
                className="w-full px-4 py-2 rounded-lg bg-gray-700 border border-gray-600 text-white focus:outline-none focus:ring-2 focus:ring-green-500"
              />
            </div>
          </div>
        </div>

        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-5 gap-4 mb-8">
          <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
            <div className="text-gray-400 text-sm mb-1">Total Tournois</div>
            <div className="text-3xl font-bold text-white">
              {totalTournaments}
            </div>
          </div>

          <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
            <div className="text-gray-400 text-sm mb-1">Profit</div>
            <div
              className={`text-3xl font-bold ${
                totalProfit >= 0 ? 'text-green-500' : 'text-red-500'
              }`}
            >
              €{totalProfit.toFixed(2)}
            </div>
          </div>

          <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
            <div className="text-gray-400 text-sm mb-1">
              Rakeback ({user.status} {rakebackPct}%)
            </div>
            <div className="text-3xl font-bold text-green-400">
              €{totalRakeback.toFixed(2)}
            </div>
          </div>

          <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
            <div className="text-gray-400 text-sm mb-1">Profit + Rakeback</div>
            <div
              className={`text-3xl font-bold ${
                profitWithRakeback >= 0 ? 'text-green-500' : 'text-red-500'
              }`}
            >
              €{profitWithRakeback.toFixed(2)}
            </div>
          </div>

          <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
            <div className="text-gray-400 text-sm mb-1">EV Total</div>
            <div
              className={`text-3xl font-bold ${
                totalEV >= 0 ? 'text-blue-500' : 'text-red-500'
              }`}
            >
              €{totalEV.toFixed(2)}
            </div>
          </div>
        </div>

        {/* Bankroll Evolution Chart */}
        <div className="bg-gray-800 rounded-xl p-6 mb-8 border border-gray-700">
          <h2 className="text-xl font-bold text-white mb-4">
            📈 Évolution de la Bankroll
          </h2>
          <ResponsiveContainer width="100%" height={400}>
            <LineChart data={chartData}>
              <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
              <XAxis dataKey="tournament" stroke="#9CA3AF" />
              <YAxis stroke="#9CA3AF" />
              <Tooltip
                contentStyle={{
                  backgroundColor: '#1F2937',
                  border: '1px solid #374151',
                  borderRadius: '8px',
                  color: '#fff',
                }}
              />
              <Legend />
              <Line
                type="monotone"
                dataKey="profit"
                stroke="#f97316"
                name="Sans Rakeback"
                strokeWidth={2}
              />
              <Line
                type="monotone"
                dataKey="profitWithRB"
                stroke="#22c55e"
                name={`Avec Rakeback (${rakebackPct}%)`}
                strokeWidth={2}
              />
            </LineChart>
          </ResponsiveContainer>
        </div>

        {/* EV vs Results Chart */}
        <div className="bg-gray-800 rounded-xl p-6 mb-8 border border-gray-700">
          <h2 className="text-xl font-bold text-white mb-4">
            📊 EV vs Résultats Réels (avec Rakeback)
          </h2>
          <ResponsiveContainer width="100%" height={400}>
            <LineChart data={chartData}>
              <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
              <XAxis dataKey="tournament" stroke="#9CA3AF" />
              <YAxis stroke="#9CA3AF" />
              <Tooltip
                contentStyle={{
                  backgroundColor: '#1F2937',
                  border: '1px solid #374151',
                  borderRadius: '8px',
                  color: '#fff',
                }}
              />
              <Legend />
              <Line
                type="monotone"
                dataKey="ev"
                stroke="#3b82f6"
                name="EV Théorique"
                strokeWidth={2}
              />
              <Line
                type="monotone"
                dataKey="profit"
                stroke="#f97316"
                name="Résultats (sans RB)"
                strokeWidth={2}
                strokeDasharray="5 5"
              />
              <Line
                type="monotone"
                dataKey="profitWithRB"
                stroke="#22c55e"
                name={`Résultats + Rakeback (${rakebackPct}%)`}
                strokeWidth={2}
              />
            </LineChart>
          </ResponsiveContainer>
        </div>

        {/* Multiplier Distribution */}
        <div className="bg-gray-800 rounded-xl p-6 border border-gray-700">
          <h2 className="text-xl font-bold text-white mb-4">
            🎲 Distribution des Multiplicateurs
          </h2>
          <ResponsiveContainer width="100%" height={300}>
            <BarChart data={multiplierData}>
              <CartesianGrid strokeDasharray="3 3" stroke="#374151" />
              <XAxis dataKey="multiplier" stroke="#9CA3AF" />
              <YAxis stroke="#9CA3AF" />
              <Tooltip
                contentStyle={{
                  backgroundColor: '#1F2937',
                  border: '1px solid #374151',
                  borderRadius: '8px',
                  color: '#fff',
                }}
              />
              <Bar dataKey="count" fill="#22c55e" name="Nombre" />
            </BarChart>
          </ResponsiveContainer>
        </div>
      </div>
    </div>
  )
}
