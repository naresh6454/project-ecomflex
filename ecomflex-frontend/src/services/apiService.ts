import axios from 'axios';

// API URL Configuration
// Ensure base URL ends with a slash for proper endpoint concatenation
const API_URL_RAW = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000/api/v1/';
const API_URL = String(API_URL_RAW).replace(/\/+$/, '') + '/';
console.log('API URL configured as:', API_URL);

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

// Create axios instance
const api = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json'
  },
  // Add a timeout to avoid long waits when hostname can't be resolved
  timeout: 15000 // Increased timeout for initial connections
});

// Add request interceptor for auth tokens
api.interceptors.request.use(
  (config) => {
    // Log the full URL being requested for debugging
    const fullUrl = (config.baseURL || '') + (config.url || '');
    console.log(`📤 Making ${config.method?.toUpperCase()} request to: ${fullUrl}`);
    
    // Check for superadmin mode first
    if (localStorage.getItem('isSuperadminMode') === 'true') {
      console.log('🔒 Using superadmin auth mode');
      config.headers['X-Superadmin-Auth'] = 'ecomflex-superadmin-direct-auth';
      return config;
    }
    
    // Standard auth token for regular users
    const token = localStorage.getItem('accessToken');
    if (token) {
      console.log('🔒 Using token auth mode');
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Add response interceptor for better error handling and URL fixing
api.interceptors.response.use(
  (response) => {
    // Fix S3 URLs in response data
    if (response.data) {
      response.data = fixS3Urls(response.data);
    }
    console.log(`✅ Response received:`, response.status);
    return response;
  },
  (error) => {
    // Enhanced error logging
    console.log('📛 Response error:', error);
    
    // Handle network errors specifically
    if (error.code === 'ERR_NETWORK' || error.code === 'ECONNABORTED') {
      console.error('Network error - check if the API server is running and accessible at:', API_URL);
      // You could implement a simple notification system here
      if (window && window.localStorage) {
        // Set a flag to show a connection error notification
        window.localStorage.setItem('apiConnectionError', 'true');
      }
    }
    return Promise.reject(error);
  }
);

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
  
  // Upload multiple booking documents
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
  
  // Upload review with optional media files
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
  
  // Get pre-signed URL for viewing a file
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
  delete: del
};