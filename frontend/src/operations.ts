import type { Operation } from './api/types'

export interface OperationConfig {
  value: Operation
  label: string
  symbol: string
  // Whether this operation needs a second operand (b).
  needsB: boolean
}

export const OPERATIONS: OperationConfig[] = [
  { value: 'add', label: 'Addition', symbol: '+', needsB: true },
  { value: 'subtract', label: 'Subtraction', symbol: '−', needsB: true },
  { value: 'multiply', label: 'Multiplication', symbol: '×', needsB: true },
  { value: 'divide', label: 'Division', symbol: '÷', needsB: true },
  { value: 'power', label: 'Exponentiation', symbol: '^', needsB: true },
  { value: 'sqrt', label: 'Square root', symbol: '√', needsB: false },
  { value: 'percent', label: 'Percentage (a% of b)', symbol: '%', needsB: true },
]

export function operationConfig(op: Operation): OperationConfig {
  const config = OPERATIONS.find((o) => o.value === op)
  if (!config) throw new Error(`unknown operation: ${op}`)
  return config
}
