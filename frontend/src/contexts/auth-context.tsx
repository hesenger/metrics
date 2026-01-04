import { createContext, useContext, useState, useEffect, ReactNode } from 'react'
import { useMeQuery } from '../api/auth'
import type { User } from '../types/user'

interface AuthContextType {
  user: User | null
  isAuthenticated: boolean
  setUser: (user: User | null) => void
}

const AuthContext = createContext<AuthContextType | undefined>(undefined)

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUserState] = useState<User | null>(() => {
    const stored = localStorage.getItem('user')
    return stored ? JSON.parse(stored) : null
  })

  const hasStoredUser = !!localStorage.getItem('user')
  const { data, error } = useMeQuery(hasStoredUser)

  useEffect(() => {
    if (data) {
      localStorage.setItem('user', JSON.stringify(data))
      setUserState(data)
    }
  }, [data])

  useEffect(() => {
    if (error) {
      localStorage.removeItem('user')
      setUserState(null)
    }
  }, [error])

  const setUser = (newUser: User | null) => {
    if (newUser) {
      localStorage.setItem('user', JSON.stringify(newUser))
      setUserState(newUser)
    } else {
      localStorage.removeItem('user')
      setUserState(null)
    }
  }

  return (
    <AuthContext.Provider
      value={{
        user,
        isAuthenticated: !!user,
        setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  )
}

export function useAuth() {
  const context = useContext(AuthContext)
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider')
  }
  return context
}
