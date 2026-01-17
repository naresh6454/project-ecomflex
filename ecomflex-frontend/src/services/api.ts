// File: /ecomflex-frontend/src/services/api.ts
import axios, { AxiosInstance } from 'axios';

// API URL Configuration
// Normalize the base URL by removing any trailing slashes so requests
// that include a leading slash won't produce double-slashes.
const API_URL_RAW = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000/api/v1/';
const API_URL = String(API_URL_RAW).replace(/\/+$/, '');
console.log('API URL configured as:', API_URL);

// Create axios instance with improved configuration
const api: AxiosInstance = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json'
  },
  timeout: 15000,
  withCredentials: false
});

// S3 URL fixer function
const fixS3Urls = (data: any): any => {
  if (!data) return data;
  
  // Handle arrays
  if (Array.isArray(data)) {
    return data.map(item => fixS3Urls(item));
  }
  
  // Handle objects
  if (typeof data === 'object' && data !== null) {
    const fixed = {...data};
    
    // Fix image URLs
    if (fixed.image && typeof fixed.image === 'string') {
      if (fixed.image.includes('your-s3-bucket')) {
        console.log('Found placeholder S3 URL:', fixed.image);
        const urlParts = fixed.image.split('/');
        const filename = urlParts[urlParts.length - 1];
        fixed.image = `https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/uploads/${filename}`;
      }
    }
    
    // Fix image URLs in other fields
    if (fixed.imageUrl && typeof fixed.imageUrl === 'string') {
      if (fixed.imageUrl.includes('your-s3-bucket')) {
        console.log('Found placeholder S3 URL in imageUrl:', fixed.imageUrl);
        const urlParts = fixed.imageUrl.split('/');
        const filename = urlParts[urlParts.length - 1];
        fixed.imageUrl = `https://ecomflex-uploads-dev.s3.eu-north-1.amazonaws.com/uploads/${filename}`;
      }
    }
    
    // Process nested objects
    Object.keys(fixed).forEach(key => {
      if (typeof fixed[key] === 'object' && fixed[key] !== null) {
        fixed[key] = fixS3Urls(fixed[key]);
      }
    });
    
    return fixed;
  }
  
  return data;
};

// Helper function to check if token is expired
const isTokenExpired = (token: string): boolean => {
  // Handle mock/development tokens
  if (token === 'mock-development-token' || token.startsWith('mock-')) {
    console.log('🔧 Using mock token - never expires');
    return false;
  }
  
  try {
    const payload = JSON.parse(atob(token.split('.')[1]));
    const expiry = payload.exp * 1000; // Convert to milliseconds
    return Date.now() >= expiry;
  } catch (e) {
    console.error('Error checking token expiration:', e);
    return true; // Assume expired if we can't parse it
  }
};

// Token debugging helper
const debugToken = () => {
  const token = localStorage.getItem('accessToken');
  if (!token) {
    console.warn('⚠️ No access token found in localStorage');
    return false;
  }
  
  console.log('🔑 Token found:', token.substring(0, 15) + '...');
  
  if (isTokenExpired(token)) {
    console.warn('⚠️ Token appears to be expired');
    return false;
  }
  
  return true;
};

// Add request interceptor with enhanced token handling
api.interceptors.request.use(
  (config) => {
  console.log(`📤 Making ${config.method?.toUpperCase()} request to: ${config.baseURL}${config.url}`);
    
    // Check for superadmin mode first
    if (localStorage.getItem('isSuperadminMode') === 'true') {
      config.headers['X-Superadmin-Auth'] = 'ecomflex-superadmin-direct-auth';
      console.log('🔒 Using superadmin auth mode');
      return config;
    }
    
    // Enhanced token handling
    const token = localStorage.getItem('accessToken');
    if (token) {
      // Check if token is in the right format (Bearer token)
      if (!token.startsWith('Bearer ') && !token.startsWith('bearer ')) {
        config.headers.Authorization = `Bearer ${token}`;
        console.log('🔑 Adding Bearer prefix to token');
      } else {
        config.headers.Authorization = token;
        console.log('🔑 Token already has Bearer prefix');
      }
      
      // Log if token appears expired
      if (isTokenExpired(token.replace('Bearer ', ''))) {
        console.warn('⚠️ Using possibly expired token');
      }
    } else {
      console.warn('⚠️ No authentication token found in localStorage');
    }
    
    return config;
  },
  (error) => {
    console.error('❌ Request interceptor error:', error);
    return Promise.reject(error);
  }
);

// Add response interceptor for better error handling and URL fixing
api.interceptors.response.use(
  (response) => {
    console.log(`✅ Response from ${response.config.url}: ${response.status}`);
    
    // Fix S3 URLs in response data
    if (response.data) {
      response.data = fixS3Urls(response.data);
    }
    return response;
  },
  (error) => {
    // Handle auth errors
    if (error.response && error.response.status === 401) {
      console.error('🚫 Authentication error (401):', error.response.data);
      console.log('🔄 Attempting to refresh token or redirect to login...');
      
      // Don't clear tokens in superadmin mode or during development
      const isSuperadminMode = localStorage.getItem('isSuperadminMode') === 'true';
      const token = localStorage.getItem('accessToken');
      const isMockToken = token === 'mock-development-token' || token?.startsWith('mock-');
      
      if (!isSuperadminMode && !isMockToken) {
        // Clear invalid token only in production mode
        localStorage.removeItem('accessToken');
      } else {
        console.log('🔧 Keeping token in development/superadmin mode');
      }
      
      // You could implement token refresh here or redirect to login
      // window.location.href = '/login';
    }
    else if (error.response && error.response.status === 403) {
      console.error('🚫 Authorization error (403):', error.response.data);
    }
    // Handle network errors
    else if (error.code === 'ERR_NETWORK' || error.code === 'ECONNABORTED') {
      console.error('🌐 Network error - API server not accessible at:', API_URL);
      if (window && window.localStorage) {
        window.localStorage.setItem('apiConnectionError', 'true');
      }
    }
    
    // Log detailed error information
    if (error.response) {
      console.error('📛 Response error:', {
        data: error.response.data,
        status: error.response.status,
        headers: error.response.headers
      });
    } else if (error.request) {
      console.error('📛 No response received:', error.request);
    } else {
      console.error('📛 Request setup error:', error.message);
    }
    
    return Promise.reject(error);
  }
);

// Add a method to manually validate/debug auth token
const checkAuth = async () => {
  if (!debugToken()) {
    console.warn('Token validation failed');
    return false;
  }
  
  try {
    // Try a request to an authenticated endpoint
    const response = await api.get('auth/verify');
    console.log('✅ Auth verification successful:', response.data);
    return true;
  } catch (error) {
    console.error('❌ Auth verification failed:', error);
    return false;
  }
};

// Generic HTTP methods
const get = (url: string) => api.get(url);
const post = (url: string, data?: any) => api.post(url, data);
const put = (url: string, data?: any) => api.put(url, data);
const del = (url: string) => api.delete(url);

// Auth service
const auth = {
  register: (userData: any) => api.post('auth/register', userData),
  registerInfluencer: (userData: any) => api.post('auth/register/influencer', userData),
  login: (email: string, password: string, recaptchaToken: string = '') => 
    api.post('auth/login', { 
      email, 
      password, 
      recaptcha_token: recaptchaToken || 'dummy-token-replace-with-actual-recaptcha' 
    }),
  googleAuth: (idToken: string, role: string = 'public') =>
    api.post('auth/google/login', { auth_token: idToken, role }),
  refreshToken: (refreshToken: string) =>
    api.post('auth/refresh', { refresh_token: refreshToken }),
  logout: (refreshToken: string) =>
    api.post('auth/logout', { refresh_token: refreshToken }),
  requestPasswordReset: (email: string) =>
    api.post('auth/password-reset', { email }),
  confirmPasswordReset: (token: string, newPassword: string, confirmPassword: string) =>
    api.post('auth/password-reset/confirm', { 
      token, 
      new_password: newPassword, 
      confirm_password: confirmPassword 
    }),
};

// User service
const users = {
  getProfile: () => api.get('users/me'),
  updateProfile: (userData: any) => api.put('users/me', userData),
};

// Referral service
const referrals = {
  validateCode: (code: string) => api.get(`referrals/check-referral-code/${code}`),
  listReferrals: () => api.get('referrals'),
};

// Admin service
const admin = {
  getInfluencers: async (params: any) => {
    const queryString = new URLSearchParams(params).toString();
    return api.get(`/admin/influencers?${queryString}`);
  },
  getInfluencer: (id: string) => api.get(`/admin/influencers/${id}`),
  updateInfluencerStatus: (id: string, status: string, reason?: string) => 
    api.put(`/admin/influencers/${id}/status`, { status, reason }),
};

// File upload service
const uploads = {
  // Upload booking document
  uploadBookingDocument: (productId: string, file: File, onProgress?: (percentage: number) => void) => {
    const formData = new FormData();
    formData.append('productId', productId);
    formData.append('file', file);
    
  return axios.post(`${API_URL}/files/booking`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        ...(localStorage.getItem('accessToken') ? { 
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        } : {})
      },
      onUploadProgress: onProgress ? (progressEvent) => {
        const percentCompleted = progressEvent.total ? 
          Math.round((progressEvent.loaded * 100) / progressEvent.total) : 0;
        onProgress(percentCompleted);
      } : undefined
    });
  },
  
  uploadBookingDocuments: (productId: string, files: File[], onProgress?: (percentage: number) => void) => {
    const formData = new FormData();
    formData.append('productId', productId);
    
    files.forEach((file, index) => {
      formData.append(`files[${index}]`, file);
    });
    
  return axios.post(`${API_URL}/files/booking/batch`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        ...(localStorage.getItem('accessToken') ? { 
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        } : {})
      },
      onUploadProgress: onProgress ? (progressEvent) => {
        const percentCompleted = progressEvent.total ? 
          Math.round((progressEvent.loaded * 100) / progressEvent.total) : 0;
        onProgress(percentCompleted);
      } : undefined
    });
  },
  
  uploadReview: (productId: string, text: string, rating: number, files: File[] = [], onProgress?: (percentage: number) => void) => {
    const formData = new FormData();
    formData.append('productId', productId);
    formData.append('text', text);
    formData.append('rating', rating.toString());
    
    files.forEach((file, index) => {
      formData.append(`files[${index}]`, file);
    });
    
  return axios.post(`${API_URL}/files/review`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        ...(localStorage.getItem('accessToken') ? { 
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        } : {})
      },
      onUploadProgress: onProgress ? (progressEvent) => {
        const percentCompleted = progressEvent.total ? 
          Math.round((progressEvent.loaded * 100) / progressEvent.total) : 0;
        onProgress(percentCompleted);
      } : undefined
    });
  },
  
  getFilePreSignedUrl: (fileKey: string) => api.get(`files/presigned-url?key=${encodeURIComponent(fileKey)}`),
};

// Products service
const products = {
  getAllProducts: (params: any = {}) => {
    const queryString = new URLSearchParams(params).toString();
    return api.get(`products?${queryString}`);
  },
  getProduct: (id: string) => api.get(`products/${id}`),
  getTrendingProducts: (limit: number = 10) => api.get(`products/trending?limit=${limit}`),
  getCategories: () => api.get('products/categories'),
  createProduct: (productData: any) => api.post('products', productData),
  updateProduct: (id: string, productData: any) => api.put(`products/${id}`, productData),
  deleteProduct: (id: string) => api.delete(`products/${id}`),
  toggleProductStatus: (id: string, status: boolean) => api.patch(`products/${id}/status`, { active: status }),
  uploadProductImage: (id: string, imageFile: File, onProgress?: (percentage: number) => void) => {
    const formData = new FormData();
    formData.append('image', imageFile);
    
  return axios.post(`${API_URL}/products/${id}/image`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        ...(localStorage.getItem('accessToken') ? { 
          Authorization: `Bearer ${localStorage.getItem('accessToken')}`
        } : {})
      },
      onUploadProgress: onProgress ? (progressEvent) => {
        const percentCompleted = progressEvent.total ? 
          Math.round((progressEvent.loaded * 100) / progressEvent.total) : 0;
        onProgress(percentCompleted);
      } : undefined
    });
  }
};

export default { 
  auth, 
  users, 
  referrals,
  admin,
  uploads,
  products,
  get,
  post,
  put,
  delete: del,
  checkAuth,
  debugToken
};