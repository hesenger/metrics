import { useEffect, useRef } from 'react'
import { useNavigate, useSearchParams } from 'react-router-dom'
import { Container, Text, Loader } from '@mantine/core'
import { useAuth } from '../contexts/auth-context'
import type { User } from '../types/user'

export function AuthCallback() {
  const navigate = useNavigate()
  const [searchParams] = useSearchParams()
  const { setUser } = useAuth()
  const hasProcessed = useRef(false)

  useEffect(() => {
    if (hasProcessed.current) return
    hasProcessed.current = true

    const id = searchParams.get('id')
    const email = searchParams.get('email')
    const oauthProvider = searchParams.get('oauth_provider')

    console.log('OAuth callback params:', { id, email, oauthProvider })

    if (!id || !email) {
      console.error('Missing OAuth params')
      navigate('/login?error=oauth_failed')
      return
    }

    const user: User = {
      id: parseInt(id, 10),
      email,
      oauth_provider: oauthProvider || undefined,
      created_at: new Date().toISOString(),
    }

    console.log('Setting user:', user)
    setUser(user)
    navigate('/')
  }, [])

  return (
    <Container size={420} my={40} style={{ textAlign: 'center' }}>
      <Loader size="lg" />
      <Text mt="md">Completing authentication...</Text>
    </Container>
  )
}
