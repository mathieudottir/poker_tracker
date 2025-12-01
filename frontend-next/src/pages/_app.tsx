import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { useState, useEffect } from 'react'

export interface User {
  id: number
  username: string
  player_name: string
  status: string
}

export default function App({ Component, pageProps }: AppProps) {
  const [user, setUser] = useState<User | null>(null)

  useEffect(() => {
    const storedUser = localStorage.getItem('user')
    if (storedUser) {
      setUser(JSON.parse(storedUser))
    }
  }, [])

  return <Component {...pageProps} user={user} setUser={setUser} />
}
