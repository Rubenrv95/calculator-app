import { useState, type FormEvent } from 'react'
import { calculate } from '../api/calculatorClient'
import { ApiError, type Operation } from '../api/types'
import { OPERATIONS, operationConfig } from '../operations'
import { parseOperand } from '../validation'

export function Calculator() {
  const [operation, setOperation] = useState<Operation>('add')
  const [aInput, setAInput] = useState('')
  const [bInput, setBInput] = useState('')
  const [result, setResult] = useState<number | null>(null)
  const [error, setError] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(false)

  const needsB = operationConfig(operation).needsB

  async function handleSubmit(event: FormEvent) {
    event.preventDefault()
    setResult(null)
    setError(null)

    const a = parseOperand(aInput, 'The first number')
    if ('error' in a) {
      setError(a.error)
      return
    }

    let b: number | undefined
    if (needsB) {
      const parsedB = parseOperand(bInput, 'The second number')
      if ('error' in parsedB) {
        setError(parsedB.error)
        return
      }
      b = parsedB.value
    }

    setIsLoading(true)
    try {
      const value = await calculate({ operation, a: a.value, b })
      setResult(value)
    } catch (err) {
      setError(err instanceof ApiError ? err.message : 'An unexpected error occurred.')
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <form className="calculator" onSubmit={handleSubmit}>
      <h1>Calculator</h1>

      <label className="field">
        <span>Operation</span>
        <select
          value={operation}
          onChange={(e) => setOperation(e.target.value as Operation)}
        >
          {OPERATIONS.map((op) => (
            <option key={op.value} value={op.value}>
              {op.label}
            </option>
          ))}
        </select>
      </label>

      <label className="field">
        <span>{needsB ? 'First number' : 'Number'}</span>
        <input
          type="text"
          inputMode="decimal"
          value={aInput}
          onChange={(e) => setAInput(e.target.value)}
          placeholder="e.g. 10"
        />
      </label>

      {needsB && (
        <label className="field">
          <span>Second number</span>
          <input
            type="text"
            inputMode="decimal"
            value={bInput}
            onChange={(e) => setBInput(e.target.value)}
            placeholder="e.g. 5"
          />
        </label>
      )}

      <button type="submit" disabled={isLoading}>
        {isLoading ? 'Calculating…' : 'Calculate'}
      </button>

      {error && (
        <p className="result result--error" role="alert">
          {error}
        </p>
      )}

      {result !== null && !error && (
        <p className="result result--success">Result: {result}</p>
      )}
    </form>
  )
}
