// parseOperand converts a raw text input into a number, or returns an
// error message if it is empty or not a valid number.
export function parseOperand(raw: string, fieldLabel: string): { value: number } | { error: string } {
  const trimmed = raw.trim()
  if (trimmed === '') {
    return { error: `${fieldLabel} is required.` }
  }
  const value = Number(trimmed)
  if (Number.isNaN(value)) {
    return { error: `${fieldLabel} must be a valid number.` }
  }
  return { value }
}
