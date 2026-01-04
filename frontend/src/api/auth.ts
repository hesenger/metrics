import { useMutation, useQuery } from '@tanstack/react-query'
import type { User } from '../types/user'

interface AuthResponse {
  user: User
}

interface LoginData {
  email: string
  password: string
}

interface RegisterData {
  email: string
  password: string
}

async function login(data: LoginData): Promise<User> {
  const response = await fetch('/api/auth/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
    credentials: 'include',
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.error || 'Login failed')
  }

  const result: AuthResponse = await response.json()
  return result.user
}

async function register(data: RegisterData): Promise<User> {
  const response = await fetch('/api/auth/register', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data),
    credentials: 'include',
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.error || 'Registration failed')
  }

  const result: AuthResponse = await response.json()
  return result.user
}

async function logout(): Promise<void> {
  const response = await fetch('/api/auth/logout', {
    method: 'POST',
    credentials: 'include',
  })

  if (!response.ok) {
    throw new Error('Logout failed')
  }
}

async function me(): Promise<User> {
  const response = await fetch('/api/auth/me', {
    method: 'GET',
    credentials: 'include',
  })

  if (!response.ok) {
    const error = await response.json()
    throw new Error(error.error || 'Failed to fetch user')
  }

  const result: AuthResponse = await response.json()
  return result.user
}

export function useLoginMutation() {
  return useMutation({
    mutationFn: login,
  })
}

export function useRegisterMutation() {
  return useMutation({
    mutationFn: register,
  })
}

export function useLogoutMutation() {
  return useMutation({
    mutationFn: logout,
  })
}

export function useMeQuery(enabled: boolean) {
  return useQuery({
    queryKey: ['me'],
    queryFn: me,
    enabled,
    retry: false,
  })
}
