import { useState, useEffect } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import { TextInput, PasswordInput, Button, Paper, Title, Text, Container, Anchor } from '@mantine/core'
import { useRegisterMutation } from '../api/auth'
import { useAuth } from '../contexts/auth-context'

export function Register() {
  const navigate = useNavigate()
  const { isAuthenticated, setUser } = useAuth()
  const [email, setEmail] = useState('')
  const [password, setPassword] = useState('')
  const [errors, setErrors] = useState({ email: '', password: '', general: '' })

  const registerMutation = useRegisterMutation()

  useEffect(() => {
    if (isAuthenticated) {
      navigate('/')
    }
  }, [isAuthenticated, navigate])

  const validateForm = () => {
    const newErrors = { email: '', password: '', general: '' }
    let isValid = true

    if (!email) {
      newErrors.email = 'Email is required'
      isValid = false
    } else if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
      newErrors.email = 'Invalid email format'
      isValid = false
    }

    if (!password) {
      newErrors.password = 'Password is required'
      isValid = false
    } else if (password.length < 8) {
      newErrors.password = 'Password must be at least 8 characters'
      isValid = false
    }

    setErrors(newErrors)
    return isValid
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    if (!validateForm()) {
      return
    }

    try {
      const user = await registerMutation.mutateAsync({ email, password })
      setUser(user)
      navigate('/')
    } catch (error) {
      setErrors({
        email: '',
        password: '',
        general: error instanceof Error ? error.message : 'Registration failed',
      })
    }
  }

  const handleGoogleRegister = () => {
    window.location.href = '/api/auth/google'
  }

  return (
    <Container size={420} my={40}>
      <Title ta="center">Create account</Title>
      <Text c="dimmed" size="sm" ta="center" mt={5}>
        Already have an account?{' '}
        <Anchor component={Link} to="/login" size="sm">
          Sign in
        </Anchor>
      </Text>

      <Paper withBorder shadow="md" p={30} mt={30} radius="md">
        <form onSubmit={handleSubmit}>
          <TextInput
            label="Email"
            placeholder="your@email.com"
            required
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            error={errors.email}
          />
          <PasswordInput
            label="Password"
            placeholder="Your password"
            required
            mt="md"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            error={errors.password}
          />
          {errors.general && (
            <Text c="red" size="sm" mt="sm">
              {errors.general}
            </Text>
          )}
          <Button fullWidth mt="xl" type="submit" loading={registerMutation.isPending}>
            Register
          </Button>
        </form>

        <Button
          fullWidth
          mt="md"
          variant="default"
          onClick={handleGoogleRegister}
        >
          Sign in with Google
        </Button>
      </Paper>
    </Container>
  )
}
