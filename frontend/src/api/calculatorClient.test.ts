import { afterEach, describe, expect, it, vi } from 'vitest'
import { calculate } from './calculatorClient'
import { ApiError } from './types'

afterEach(() => {
  vi.unstubAllGlobals()
})

describe('calculate', () => {
  it('returns the result on a successful response', async () => {
    const fetchMock = vi.fn().mockResolvedValue({
      ok: true,
      json: async () => ({ result: 5 }),
    })
    vi.stubGlobal('fetch', fetchMock)

    const result = await calculate({ operation: 'add', a: 2, b: 3 })

    expect(result).toBe(5)
    expect(fetchMock).toHaveBeenCalledWith(
      expect.stringContaining('/calculate'),
      expect.objectContaining({
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ operation: 'add', a: 2, b: 3 }),
      }),
    )
  })

  it('throws an ApiError with the backend message on a non-ok response', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockResolvedValue({
        ok: false,
        json: async () => ({ error: 'division by zero' }),
      }),
    )

    await expect(calculate({ operation: 'divide', a: 10, b: 0 })).rejects.toThrow(
      new ApiError('division by zero'),
    )
  })

  it('throws a generic ApiError when the network request fails', async () => {
    vi.stubGlobal(
      'fetch',
      vi.fn().mockRejectedValue(new TypeError('network error')),
    )

    await expect(calculate({ operation: 'add', a: 1, b: 2 })).rejects.toThrow(ApiError)
  })
})
