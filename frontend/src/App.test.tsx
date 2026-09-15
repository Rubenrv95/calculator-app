import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import App from './App'

beforeEach(() => {
  localStorage.clear()
  document.documentElement.removeAttribute('data-theme')
  vi.stubGlobal(
    'matchMedia',
    vi.fn().mockImplementation((query: string) => ({
      matches: false,
      media: query,
      addEventListener: vi.fn(),
      removeEventListener: vi.fn(),
    })),
  )
})

describe('App', () => {
  it('renders the calculator and the theme toggle', () => {
    render(<App />)

    expect(screen.getByRole('heading', { name: 'Calculator' })).toBeInTheDocument()
    expect(screen.getByRole('switch', { name: 'Toggle dark mode' })).toBeInTheDocument()
  })

  it('switches the document theme when the toggle is clicked', async () => {
    const user = userEvent.setup()
    render(<App />)

    expect(document.documentElement.dataset.theme).toBe('light')

    await user.click(screen.getByRole('switch', { name: 'Toggle dark mode' }))

    expect(document.documentElement.dataset.theme).toBe('dark')
  })
})
