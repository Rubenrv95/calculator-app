import { describe, expect, it } from 'vitest'
import { parseOperand } from './validation'

describe('parseOperand', () => {
  it('parses a valid number', () => {
    expect(parseOperand('42', 'Number')).toEqual({ value: 42 })
  })

  it('parses a valid negative decimal', () => {
    expect(parseOperand('-3.5', 'Number')).toEqual({ value: -3.5 })
  })

  it('rejects an empty string', () => {
    expect(parseOperand('', 'Number')).toEqual({ error: 'Number is required.' })
  })

  it('rejects a whitespace-only string', () => {
    expect(parseOperand('   ', 'Number')).toEqual({ error: 'Number is required.' })
  })

  it('rejects a non-numeric string', () => {
    expect(parseOperand('abc', 'Number')).toEqual({ error: 'Number must be a valid number.' })
  })
})
