/**
 * Data formatting utilities
 */

// Format date to readable string
export const formatDate = (dateString: string, format: 'short' | 'medium' | 'long' = 'medium'): string => {
  const date = new Date(dateString)
  
  if (isNaN(date.getTime())) {
    return 'Invalid date'
  }
  
  const options: Intl.DateTimeFormatOptions = {}
  
  switch (format) {
    case 'short':
      options.year = 'numeric'
      options.month = 'short'
      options.day = 'numeric'
      break
    case 'medium':
      options.year = 'numeric'
      options.month = 'short'
      options.day = 'numeric'
      options.hour = '2-digit'
      options.minute = '2-digit'
      break
    case 'long':
      options.year = 'numeric'
      options.month = 'long'
      options.day = 'numeric'
      options.weekday = 'long'
      options.hour = '2-digit'
      options.minute = '2-digit'
      options.second = '2-digit'
      break
  }
  
  return new Intl.DateTimeFormat('en-US', options).format(date)
}

// Format relative time
export const formatRelativeTime = (dateString: string): string => {
  const date = new Date(dateString)
  const now = new Date()
  const diffInSeconds = Math.floor((now.getTime() - date.getTime()) / 1000)
  
  if (diffInSeconds < 60) {
    return `${diffInSeconds} second${diffInSeconds !== 1 ? 's' : ''} ago`
  }
  
  const diffInMinutes = Math.floor(diffInSeconds / 60)
  if (diffInMinutes < 60) {
    return `${diffInMinutes} minute${diffInMinutes !== 1 ? 's' : ''} ago`
  }
  
  const diffInHours = Math.floor(diffInMinutes / 60)
  if (diffInHours < 24) {
    return `${diffInHours} hour${diffInHours !== 1 ? 's' : ''} ago`
  }
  
  const diffInDays = Math.floor(diffInHours / 24)
  if (diffInDays < 30) {
    return `${diffInDays} day${diffInDays !== 1 ? 's' : ''} ago`
  }
  
  const diffInMonths = Math.floor(diffInDays / 30)
  if (diffInMonths < 12) {
    return `${diffInMonths} month${diffInMonths !== 1 ? 's' : ''} ago`
  }
  
  const diffInYears = Math.floor(diffInMonths / 12)
  return `${diffInYears} year${diffInYears !== 1 ? 's' : ''} ago`
}

// Format currency
export const formatCurrency = (amount: number, currency: string = 'USD'): string => {
  return new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency,
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  }).format(amount)
}

// Format percentage
export const formatPercentage = (value: number, decimals: number = 1): string => {
  return `${value.toFixed(decimals)}%`
}

// Format file size
export const formatFileSize = (bytes: number): string => {
  if (bytes === 0) return '0 Bytes'
  
  const k = 1024
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  
  return `${parseFloat((bytes / Math.pow(k, i)).toFixed(2))} ${sizes[i]}`
}

// Truncate text with ellipsis
export const truncateText = (text: string, maxLength: number): string => {
  if (text.length <= maxLength) return text
  
  return `${text.substring(0, maxLength)}...`
}

// Format booking status
export const formatBookingStatus = (status: string): string => {
  switch (status) {
    case 'initiated':
      return 'Initiated'
    case 'pending':
      return 'Pending Verification'
    case 'approved':
      return 'Approved'
    case 'rejected':
      return 'Rejected'
    default:
      return status.charAt(0).toUpperCase() + status.slice(1)
  }
}

// Format cashback status
export const formatCashbackStatus = (status: string): string => {
  switch (status) {
    case 'not_paid':
      return 'Pending'
    case 'paid':
      return 'Paid'
    default:
      return status.charAt(0).toUpperCase() + status.slice(1)
  }
}

// Format phone number
export const formatPhoneNumber = (phoneNumber: string): string => {
  // Remove all non-digit characters
  const cleaned = phoneNumber.replace(/\D/g, '')
  
  // Format based on length
  if (cleaned.length === 10) {
    // US format: (xxx) xxx-xxxx
    return `(${cleaned.substring(0, 3)}) ${cleaned.substring(3, 6)}-${cleaned.substring(6, 10)}`
  } else {
    // Default format for other lengths
    return phoneNumber
  }
}

// Format name to initials
export const getInitials = (name: string): string => {
  if (!name) return ''
  
  const parts = name.split(' ')
  if (parts.length === 1) return parts[0].charAt(0).toUpperCase()
  
  return (parts[0].charAt(0) + parts[parts.length - 1].charAt(0)).toUpperCase()
}

// Format S3 image URL
export const getS3ImageUrl = (url: string): string => {
  if (!url) return '';
  
  console.log('Processing URL:', url);
  
  // If it's already a full URL with the correct bucket name
  if (url.includes('ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com')) {
    return url;
  }
  
  // If it's just a key or path
  if (url.includes('/products/')) {
    // Extract just the filename if it has a path structure
    const parts = url.split('/');
    const filename = parts[parts.length - 1];
    return `https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/uploads/${filename}`;
  }
  
  // If it's already a full URL but with wrong bucket
  if (url.includes('your-s3-bucket.s3.amazonaws.com')) {
    const urlParts = url.split('/');
    const filename = urlParts[urlParts.length - 1];
    return `https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/uploads/${filename}`;
  }
  
  // Default case - assume it's just a filename
  return `https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/uploads/${url}`;
}