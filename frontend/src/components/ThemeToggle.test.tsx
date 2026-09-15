import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { describe, expect, it, vi } from 'vitest'
import { ThemeToggle } from './ThemeToggle'

describe('ThemeToggle', () => {
  it('renders as unchecked when the theme is light', () => {
    render(<ThemeToggle theme="light" onToggle={vi.fn()} />)

    expect(screen.getByRole('switch', { name: 'Toggle dark mode' })).toHaveAttribute(
      'aria-checked',
      'false',
    )
  })

  it('renders as checked when the theme is dark', () => {
    render(<ThemeToggle theme="dark" onToggle={vi.fn()} />)

    expect(screen.getByRole('switch', { name: 'Toggle dark mode' })).toHaveAttribute(
      'aria-checked',
      'true',
    )
  })

  it('calls onToggle when clicked', async () => {
    const onToggle = vi.fn()
    const user = userEvent.setup()
    render(<ThemeToggle theme="light" onToggle={onToggle} />)

    await user.click(screen.getByRole('switch', { name: 'Toggle dark mode' }))

    expect(onToggle).toHaveBeenCalledTimes(1)
  })
})
