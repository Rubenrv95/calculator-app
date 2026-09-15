import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { beforeEach, describe, expect, it, vi } from 'vitest'
import { Calculator } from './Calculator'
import { calculate } from '../api/calculatorClient'
import { ApiError } from '../api/types'

vi.mock('../api/calculatorClient', () => ({
  calculate: vi.fn(),
}))

const calculateMock = vi.mocked(calculate)

beforeEach(() => {
  calculateMock.mockReset()
})

describe('Calculator', () => {
  it('renders the form fields', () => {
    render(<Calculator />)

    expect(screen.getByRole('heading', { name: 'Calculator' })).toBeInTheDocument()
    expect(screen.getByLabelText('Operation')).toBeInTheDocument()
    expect(screen.getByLabelText('First number')).toBeInTheDocument()
    expect(screen.getByLabelText('Second number')).toBeInTheDocument()
    expect(screen.getByRole('button', { name: 'Calculate' })).toBeInTheDocument()
  })

  it('hides the second operand field for sqrt', async () => {
    const user = userEvent.setup()
    render(<Calculator />)

    await user.selectOptions(screen.getByLabelText('Operation'), 'sqrt')

    expect(screen.getByLabelText('Number')).toBeInTheDocument()
    expect(screen.queryByLabelText('Second number')).not.toBeInTheDocument()
  })

  it('shows a validation error without calling the API when a field is empty', async () => {
    const user = userEvent.setup()
    render(<Calculator />)

    await user.type(screen.getByLabelText('First number'), '5')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(await screen.findByRole('alert')).toHaveTextContent('The second number is required.')
    expect(calculateMock).not.toHaveBeenCalled()
  })

  it('shows a validation error for non-numeric input', async () => {
    const user = userEvent.setup()
    render(<Calculator />)

    await user.type(screen.getByLabelText('First number'), 'abc')
    await user.type(screen.getByLabelText('Second number'), '2')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(await screen.findByRole('alert')).toHaveTextContent('The first number must be a valid number.')
    expect(calculateMock).not.toHaveBeenCalled()
  })

  it('displays the result on a successful calculation', async () => {
    calculateMock.mockResolvedValue(8)
    const user = userEvent.setup()
    render(<Calculator />)

    await user.type(screen.getByLabelText('First number'), '5')
    await user.type(screen.getByLabelText('Second number'), '3')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(await screen.findByText('Result: 8')).toBeInTheDocument()
    expect(calculateMock).toHaveBeenCalledWith({ operation: 'add', a: 5, b: 3 })
  })

  it('displays the backend error message when the API call fails', async () => {
    calculateMock.mockRejectedValue(new ApiError('division by zero'))
    const user = userEvent.setup()
    render(<Calculator />)

    await user.selectOptions(screen.getByLabelText('Operation'), 'divide')
    await user.type(screen.getByLabelText('First number'), '10')
    await user.type(screen.getByLabelText('Second number'), '0')
    await user.click(screen.getByRole('button', { name: 'Calculate' }))

    expect(await screen.findByRole('alert')).toHaveTextContent('division by zero')
  })
})
