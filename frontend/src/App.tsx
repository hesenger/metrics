import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { Home } from './pages/home'
import { Login } from './pages/login'
import { Register } from './pages/register'
import { AuthCallback } from './pages/auth-callback'
import { ProtectedRoute } from './components/protected-route'

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/auth/callback" element={<AuthCallback />} />
        <Route
          path="/"
          element={
            <ProtectedRoute>
              <Home />
            </ProtectedRoute>
          }
        />
      </Routes>
    </BrowserRouter>
  )
}

export default App
