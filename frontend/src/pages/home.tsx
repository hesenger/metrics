import { useNavigate } from 'react-router-dom'
import { Container, Title, Text, Button, Group } from '@mantine/core'
import { useAuth } from '../contexts/auth-context'
import { useLogoutMutation } from '../api/auth'

export function Home() {
  const navigate = useNavigate()
  const { user, setUser } = useAuth()
  const logoutMutation = useLogoutMutation()

  const handleLogout = async () => {
    try {
      await logoutMutation.mutateAsync()
      setUser(null)
      navigate('/login')
    } catch (error) {
      console.error('Logout failed:', error)
    }
  }

  return (
    <Container>
      <Group justify="space-between" mb="xl">
        <div>
          <Title>Metrics</Title>
          <Text>Analytics and Feature Flags Platform</Text>
        </div>
        <div>
          <Text size="sm" c="dimmed" mb="xs">
            {user?.email}
          </Text>
          <Button onClick={handleLogout} loading={logoutMutation.isPending}>
            Logout
          </Button>
        </div>
      </Group>
    </Container>
  )
}
