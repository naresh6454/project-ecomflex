/**
 * Form validation utilities
 */

// Email validation
export const isValidEmail = (email: string): boolean => {
  const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
  return re.test(String(email).toLowerCase())
}

// Password validation - at least 8 chars, 1 uppercase, 1 lowercase, 1 number
export const isValidPassword = (password: string): boolean => {
  const re = /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$/
  return re.test(password)
}

// Phone number validation - basic format check
export const isValidPhone = (phone: string): boolean => {
  const re = /^\+?[0-9]{10,15}$/
  return re.test(phone.replace(/\s+/g, ''))
}

// Required field validation
export const isRequired = (value: any): boolean => {
  if (value === null || value === undefined) return false
  
  // Check for empty string
  if (typeof value === 'string') {
    return value.trim().length > 0
  }
  
  // Check for empty array
  if (Array.isArray(value)) {
    return value.length > 0
  }
  
  // Check for empty object
  if (typeof value === 'object') {
    return Object.keys(value).length > 0
  }
  
  return true
}

// Min length validation
export const minLength = (value: string, min: number): boolean => {
  return value.length >= min
}

// Max length validation
export const maxLength = (value: string, max: number): boolean => {
  return value.length <= max
}

// URL validation
export const isValidUrl = (url: string): boolean => {
  try {
    new URL(url)
    return true
  } catch (error) {
    return false
  }
}

// ASIN validation (Amazon Standard Identification Number)
export const isValidASIN = (asin: string): boolean => {
  // ASIN is typically 10 characters, containing letters and numbers
  const re = /^[A-Z0-9]{10}$/
  return re.test(asin)
}

// Referral code validation
export const isValidReferralCode = (code: string): boolean => {
  // Allow alphanumeric with optional hyphens or underscores, 4-20 chars
  const re = /^[a-zA-Z0-9_-]{4,20}$/
  return re.test(code)
}

// Validate file size
export const isValidFileSize = (file: File, maxSizeInMB: number): boolean => {
  const maxSizeInBytes = maxSizeInMB * 1024 * 1024
  return file.size <= maxSizeInBytes
}

// Validate file type
export const isValidFileType = (file: File, acceptedTypes: string[]): boolean => {
  return acceptedTypes.includes(file.type)
}

// Form validator utility
export type ValidationRule = {
  rule: (value: any, ...args: any[]) => boolean
  message: string
  args?: any[]
}

export const validateField = (value: any, rules: ValidationRule[]): string | null => {
  for (const { rule, message, args } of rules) {
    const isValid = args ? rule(value, ...args) : rule(value)
    if (!isValid) {
      return message
    }
  }
  return null
}