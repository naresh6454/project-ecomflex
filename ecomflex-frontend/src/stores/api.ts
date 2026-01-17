import axios from 'axios'

// Create an axios instance with custom configuration
// Normalize base URL to avoid accidental double slashes
const BASE_API_URL = String(import.meta.env.VITE_API_URL || 'http://localhost:8000/api').replace(/\/+$/, '')
const apiClient = axios.create({
  baseURL: BASE_API_URL,
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  },
  timeout: 10000 // 10 seconds
})

// Request interceptor
apiClient.interceptors.request.use(
  config => {
    // Get token from localStorage
    const token = localStorage.getItem('accessToken')
    
    // If token exists, add it to request headers
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    
    return config
  },
  error => {
    return Promise.reject(error)
  }
)

// Response interceptor
apiClient.interceptors.response.use(
  response => {
    return response
  },
  async error => {
    const originalRequest = error.config
    
    // If error is 401 (Unauthorized) and the request hasn't been retried yet
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true
      
      try {
        // Try to refresh the token
        const refreshToken = localStorage.getItem('refreshToken')
        
        if (refreshToken) {
          const response = await axios.post(
            `${BASE_API_URL}/auth/refresh/`,
            { refresh: refreshToken }
          )
          
          const { access } = response.data
          
          // Update token in localStorage
          localStorage.setItem('accessToken', access)
          
          // Update Authorization header
          originalRequest.headers.Authorization = `Bearer ${access}`
          apiClient.defaults.headers.common.Authorization = `Bearer ${access}`
          
          // Retry the original request
          return apiClient(originalRequest)
        }
      } catch (refreshError) {
        console.error('Failed to refresh token:', refreshError)
        
        // Redirect to login page
        window.location.href = '/auth/login'
      }
    }
    
    // Extract error message
    let errorMessage = 'Something went wrong'
    
    if (error.response?.data?.message) {
      errorMessage = error.response.data.message
    } else if (error.message) {
      errorMessage = error.message
    }
    
    // Network error
    if (!error.response) {
      errorMessage = 'Network error. Please check your internet connection.'
    }
    
    // Handle different status codes
    switch (error.response?.status) {
      case 400:
        console.error('Bad Request:', errorMessage)
        break
      case 401:
        console.error('Unauthorized:', errorMessage)
        break
      case 403:
        console.error('Forbidden:', errorMessage)
        break
      case 404:
        console.error('Not Found:', errorMessage)
        break
      case 422:
        console.error('Validation Error:', errorMessage)
        break
      case 500:
        console.error('Server Error:', errorMessage)
        break
      default:
        console.error('Error:', errorMessage)
    }
    
    return Promise.reject(error)
  }
)

// API service with organized endpoints
const api = {
  // Auth endpoints
  auth: {
    login: (email: string, password: string) => 
      apiClient.post('/auth/login/', { email, password }),
    
    register: (userData: any) =>
      apiClient.post('/auth/register/public/', userData),
    
    registerInfluencer: (userData: any) =>
      apiClient.post('/auth/register/influencer/', userData),
    
    googleAuth: (idToken: string, role: string = 'public') =>
      apiClient.post('/auth/google/', { id_token: idToken, role }),
    
    refreshToken: (refreshToken: string) =>
      apiClient.post('/auth/refresh/', { refresh: refreshToken }),
    
    requestPasswordReset: (email: string) =>
      apiClient.post('/auth/password/reset/request/', { email }),
    
    confirmPasswordReset: (token: string, newPassword: string, confirmPassword: string) =>
      apiClient.post('/auth/password/reset/confirm/', { 
        token, 
        new_password: newPassword, 
        confirm_password: confirmPassword 
      }),
    
    logout: (refreshToken: string) =>
      apiClient.post('/auth/logout/', { refresh: refreshToken })
  },
  
  // User endpoints
  users: {
    getProfile: () => apiClient.get('/users/profile/'),
    
    updateProfile: (userData: any) => apiClient.patch('/users/profile/', userData),
    
    getInfluencerProfile: () => apiClient.get('/users/influencer-profile/'),
    
    updateInfluencerProfile: (profileData: any) => 
      apiClient.patch('/users/influencer-profile/', profileData)
  },

  // Product endpoints
  products: {
    getAll: (params?: any) => apiClient.get('/products/items/', { params }),
    
    get: (id: string) => apiClient.get(`/products/items/${id}/`)
  },
  
  // Booking endpoints
  bookings: {
    getAll: (params?: any) => apiClient.get('/bookings/items/', { params }),
    
    get: (id: string) => apiClient.get(`/bookings/items/${id}/`),
    
    create: (bookingData: any) => apiClient.post('/bookings/items/', bookingData)
  },
  
  // Referral endpoints
  referrals: {
    validateCode: (referralCode: string) => 
      apiClient.post('/referrals/validate-code/', { referral_code: referralCode })
  }
}

export default api