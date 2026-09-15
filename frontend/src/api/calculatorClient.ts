import { ApiError, type ApiErrorResponse, type CalculateRequest, type CalculateResponse } from './types'

const BASE_URL = import.meta.env.VITE_API_BASE_URL ?? 'http://localhost:8080'

// calculate calls the backend's POST /calculate endpoint and returns the
// numeric result. It throws an ApiError with the backend's message on any
// non-2xx response, or a generic message if the server is unreachable.
export async function calculate(request: CalculateRequest): Promise<number> {
  let response: Response
  try {
    response = await fetch(`${BASE_URL}/calculate`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(request),
    })
  } catch {
    throw new ApiError('Could not connect to the server. Check that the backend is running.')
  }

  if (!response.ok) {
    const body = (await response.json().catch(() => null)) as ApiErrorResponse | null
    throw new ApiError(body?.error ?? 'An unexpected error occurred while calculating.')
  }

  const body = (await response.json()) as CalculateResponse
  return body.result
}
