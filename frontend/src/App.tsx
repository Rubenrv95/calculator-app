import { Calculator } from './components/Calculator'
import { ThemeToggle } from './components/ThemeToggle'
import { useTheme } from './hooks/useTheme'
import './App.css'

function App() {
  const { theme, toggleTheme } = useTheme()

  return (
    <main className="app">
      <ThemeToggle theme={theme} onToggle={toggleTheme} />
      <Calculator />
    </main>
  )
}

export default App
