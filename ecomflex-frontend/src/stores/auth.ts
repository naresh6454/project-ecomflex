import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import apiService from '../services/apiService';
import axios from 'axios';

// Existing interfaces
export interface LoginCredentials {
  email: string;
  password: string;
}

export interface User {
  id: string;
  fullName: string;
  email: string;
  phone: string;
  role: 'admin' | 'influencer' | 'public';
  profilePicture: string;
  createdAt: string;
  lastLogin: string;
  updatedAt?: string;
}

export interface PublicUserRegistration {
  fullName: string;
  email: string;
  password: string;
  phone: string;
  agreeToTerms: boolean;
}

export interface InfluencerRegistration {
  fullName: string;
  email: string;
  password: string;
  phone: string;
  socialMedia: Record<string, string>;
  followers: number;
  niche: string;
  agreeToTerms: boolean;
  referralCode?: string;
}

export interface PasswordResetRequest {
  email: string;
}

export interface VerifyResetCodeRequest {
  email: string;
  code: string;
}

export interface ResetPasswordRequest {
  email: string;
  code: string;
  newPassword: string;
}

// Create a special superadmin API instance that bypasses normal authentication
const createSuperadminApiInstance = () => {
  // Use the same base URL as the regular API
  const API_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000/api/v1/';
  
  // Create a custom axios instance for superadmin
  const superadminApi = axios.create({
    baseURL: API_URL,
    headers: {
      'Content-Type': 'application/json',
      // Include a special header that your backend can check for superadmin
      'X-Superadmin-Auth': 'ecomflex-superadmin-direct-auth'
    },
    timeout: 15000
  });
  
  return superadminApi;
};

// Define the auth store
export const useAuthStore = defineStore('auth', () => {
  // State
  const user = ref<User | null>(null);
  const accessToken = ref<string | null>(localStorage.getItem('accessToken'));
  const refreshToken = ref<string | null>(localStorage.getItem('refreshToken'));
  const isLoading = ref(false);
  const error = ref<string | null>(null);
  const isSuperadminMode = ref(false);

  // Create a superadmin API instance
  const superadminApi = createSuperadminApiInstance();

  // Getters
  const isAuthenticated = computed(() => !!accessToken.value || isSuperadminMode.value);
  const isSuperAdmin = computed(() => user.value?.role === 'admin' || isSuperadminMode.value);
  const isInfluencer = computed(() => user.value?.role === 'influencer');
  const isPublicUser = computed(() => user.value?.role === 'public');
  const isApprovedInfluencer = computed(() => {
    return user.value?.role === 'influencer';
  });

  // Actions
  const setAccessToken = (token: string) => {
    accessToken.value = token;
    localStorage.setItem('accessToken', token);
    // Also store token with the name the router checks for
    localStorage.setItem('token', token);
  };

  const setRefreshToken = (token: string) => {
    refreshToken.value = token;
    localStorage.setItem('refreshToken', token);
  };

  const setUser = (userData: User) => {
    user.value = userData;
    // Store the user role in localStorage for the router to access
    if (userData.role) {
      localStorage.setItem('userRole', userData.role);
      console.log('User role set in localStorage:', userData.role);
    }
  };

  const clearTokens = () => {
    accessToken.value = null;
    refreshToken.value = null;
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
    localStorage.removeItem('token');
    localStorage.removeItem('userRole');
  };

  // Modified API methods for superadmin mode
  const getSuperadminData = async (url: string) => {
    if (!isSuperadminMode.value) {
      throw new Error('Not in superadmin mode');
    }
    
    try {
      const response = await superadminApi.get(url);
      return response;
    } catch (err: any) {
      console.error('Superadmin API error:', err);
      
      // For superadmin, return mock data if the API fails
      if (url.includes('/admin/influencers')) {
        return {
          data: {
            status: 'success',
            data: {
              data: [
                {
                  id: 'mock-id-1',
                  email: 'influencer1@example.com',
                  full_name: 'Mock Influencer 1',
                  profile_picture: null,
                  role: 'influencer',
                  status: 'pending_approval',
                  referral_code: 'INFL1',
                  social_links: ['instagram:https://instagram.com/mock1'],
                  follower_count: 12500,
                  phone: '+11234567890',
                  created_at: new Date().toISOString(),
                  updated_at: new Date().toISOString(),
                }
              ],
              total: 1,
              page: 1,
              page_size: 10,
              total_pages: 1
            }
          }
        };
      }
      
      throw err;
    }
  };

  const login = async (email: string, password: string) => {
    isLoading.value = true;
    error.value = null;
    
    console.log(`Login attempt for email: ${email}`);
    
    try {
      // Special handling for superadmin login
      if (email === 'admin@ecomflex.com' && password === 'Admin@123') {
        console.log('SUPERADMIN LOGIN DETECTED - DIRECT AUTHENTICATION');
        
        // Set superadmin mode
        isSuperadminMode.value = true;
        
        // Create a hardcoded superadmin user
        const superadminUser: User = {
          id: 'superadmin-id',
          email: 'admin@ecomflex.com',
          fullName: 'System Administrator',
          phone: '',
          role: 'admin',
          profilePicture: '',
          createdAt: new Date().toISOString(),
          lastLogin: new Date().toISOString(),
          updatedAt: new Date().toISOString()
        };
        
        // Set the user with admin role
        setUser(superadminUser);
        
        // Force the userRole to admin
        localStorage.setItem('userRole', 'admin');
        localStorage.setItem('isSuperadminMode', 'true');
        console.log('Superadmin role set in localStorage: admin');
        
        return true;
      }
      
      // Normal login flow for other users
      // Ensure we're not in superadmin mode for regular users
      isSuperadminMode.value = false;
      localStorage.removeItem('isSuperadminMode');
      
      const response = await apiService.auth.login(email, password);
      
      let token, refreshTokenValue, userData;
      
      // Handle different response formats
      if (response.data.access) {
        token = response.data.access;
        refreshTokenValue = response.data.refresh;
        userData = response.data.user;
      } else if (response.data.token) {
        token = response.data.token.access;
        refreshTokenValue = response.data.token.refresh;
        userData = response.data.user;
      } else if (response.data.data && response.data.data.token) {
        token = response.data.data.token.access_token || response.data.data.token.access;
        refreshTokenValue = response.data.data.token.refresh_token || response.data.data.token.refresh;
        userData = response.data.data.user;
      }

      if (token) {
        setAccessToken(token);
        setRefreshToken(refreshTokenValue);
        setUser(userData);
        
        // Additional check to ensure token is set properly
        console.log('Token set successfully:', !!localStorage.getItem('token'));
        console.log('UserRole set successfully:', localStorage.getItem('userRole'));
        
        return true;
      } else {
        throw new Error('Invalid token format in response');
      }
    } catch (err: any) {
      console.error('Login error:', err);
      error.value = err.response?.data?.message || err.response?.data?.detail || 'Login failed. Please check your credentials.';
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const register = async (userData: PublicUserRegistration) => {
    isLoading.value = true;
    error.value = null;

    try {
      // Transform the userData to match API expectations
      const apiData = {
        email: userData.email,
        password: userData.password,
        full_name: userData.fullName,
        phone: userData.phone,
        role: 'public',
        captcha_token: '123456', // Replace with actual captcha token
      };

      const response = await apiService.auth.register(apiData);
      console.log('Registration response:', response);
      
      // Check if the response has success property
      if (response.data && (response.data.success === true || response.data.status === 'success')) {
        // Registration was successful
        console.log('Registration successful');
        
        // If tokens are provided, set them
        if (response.data.token || (response.data.data && response.data.data.token)) {
          const tokenData = response.data.token || response.data.data.token;
          if (tokenData.access_token || tokenData.access) {
            setAccessToken(tokenData.access_token || tokenData.access);
            setRefreshToken(tokenData.refresh_token || tokenData.refresh);
          }
        }
        
        // If user data is provided, set it
        if (response.data.user || (response.data.data && response.data.data.user)) {
          const userData = response.data.user || response.data.data.user;
          setUser(userData);
        }
        
        return true;
      } else {
        // Handle case where API returns success: false
        error.value = response.data?.message || 'Registration failed. Please try again.';
        return false;
      }
    } catch (err: any) {
      console.error('Registration error:', err);
      
      // Check if it's a 409 conflict - treat as success
      if (err.response?.status === 409) {
        console.log('Email already exists - treating as success');
        error.value = "An account with this email already exists. You can login with your credentials.";
        return true; // Show success modal
      }
      
      // Explicitly check if the response has success=true despite being an error
      if (err.response?.data?.success === true) {
        console.log('Response marked as success despite error');
        return true;
      }
      
      error.value = err.response?.data?.message || err.response?.data?.detail || 'Registration failed. Please try again.';
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const registerInfluencer = async (userData: InfluencerRegistration) => {
    isLoading.value = true;
    error.value = null;

    try {
      // Create social links array from social media object
      const socialLinks = Object.entries(userData.socialMedia)
        .filter(([_, url]) => url) // Filter out empty URLs
        .map(([platform, url]) => `${platform}:${url}`);
      
      // Transform the userData to match API expectations
      const apiData = {
        email: userData.email,
        password: userData.password,
        full_name: userData.fullName,
        phone: userData.phone,
        role: 'influencer',
        captcha_token: '123456', // Replace with actual captcha token
        social_links: socialLinks,
        follower_count: userData.followers,
        referral_code: userData.referralCode || '', 
      };

      console.log('Sending registration data:', apiData);
      
      const response = await apiService.auth.register(apiData);
      console.log('Registration response:', response);
      
      // Check if the response has success property
      if (response.data && (response.data.success === true || response.data.status === 'success')) {
        // Registration was successful
        console.log('Registration successful');
        
        // If tokens are provided, set them
        if (response.data.token || (response.data.data && response.data.data.token)) {
          const tokenData = response.data.token || response.data.data.token;
          if (tokenData.access_token || tokenData.access) {
            setAccessToken(tokenData.access_token || tokenData.access);
            setRefreshToken(tokenData.refresh_token || tokenData.refresh);
          }
        }
        
        // If user data is provided, set it
        if (response.data.user || (response.data.data && response.data.data.user)) {
          const userData = response.data.user || response.data.data.user;
          setUser(userData);
        }
        
        return true;
      } else {
        // Handle case where API returns success: false
        error.value = response.data?.message || 'Registration failed. Please try again.';
        return false;
      }
    } catch (err: any) {
      console.error('Influencer registration error:', err);
      console.error('Error response:', err.response);
      
      // Check for 409 conflict - treat as success
      if (err.response?.status === 409) {
        console.log('Email already exists - treating as success');
        error.value = "An account with this email already exists. You can login with your credentials.";
        return true; // Show success modal
      }
      
      // Check if the error response itself contains success: true
      if (err.response?.data?.success === true) {
        console.log('Response marked as success despite error');
        return true;
      }
      
      error.value = err.response?.data?.message || err.response?.data?.detail || 'Registration failed. Please try again.';
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const googleAuth = async (idToken: string, role: string = 'public') => {
    isLoading.value = true;
    error.value = null;

    try {
      const response = await apiService.auth.googleAuth(idToken, role);
      const { access, refresh, user: userData } = response.data;

      setAccessToken(access);
      setRefreshToken(refresh);
      setUser(userData);

      return true;
    } catch (err: any) {
      error.value = err.response?.data?.detail || 'Google authentication failed. Please try again.';
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const refreshAuth = async () => {
    // If in superadmin mode, skip token refresh
    if (isSuperadminMode.value) {
      return true;
    }
    
    // Handle mock/development tokens
    const currentToken = accessToken.value;
    if (currentToken === 'mock-development-token' || currentToken?.startsWith('mock-')) {
      console.log('🔧 Using mock token - skipping refresh');
      return true;
    }

    if (!refreshToken.value) return false;

    try {
      const response = await apiService.auth.refreshToken(refreshToken.value);
      
      // Handle different response formats
      const token = response.data.access || response.data.token?.access || 
                  (response.data.data && response.data.data.token ? 
                   response.data.data.token.access_token || response.data.data.token.access : null);

      if (token) {
        setAccessToken(token);
        return true;
      }
      return false;
    } catch (err) {
      // Don't logout in development mode with mock tokens
      const isMockToken = accessToken.value === 'mock-development-token' || accessToken.value?.startsWith('mock-');
      if (!isMockToken) {
        logout();
      } else {
        console.log('🔧 Mock token refresh failed, but keeping session active');
      }
      return false;
    }
  };

  const fetchUserProfile = async () => {
    // If in superadmin mode, skip profile fetch
    if (isSuperadminMode.value) {
      return true;
    }
    
    if (!accessToken.value) return false;

    isLoading.value = true;

    try {
      const response = await apiService.users.getProfile();
      setUser(response.data);
      return true;
    } catch (err) {
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const requestPasswordReset = async (email: string) => {
    isLoading.value = true;
    error.value = null;

    try {
      await apiService.auth.requestPasswordReset(email);
      return true;
    } catch (err: any) {
      error.value = err.response?.data?.detail || 'Failed to request password reset. Please try again.';
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const verifyResetCode = async (_data: VerifyResetCodeRequest) => {
    isLoading.value = true;
    error.value = null;

    try {
      // This would need to be implemented in your API service
      // For now, just return true as a placeholder
      return true;
    } catch (err: any) {
      error.value = err.response?.data?.detail || 'Invalid or expired code. Please try again.';
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const resetPassword = async (data: ResetPasswordRequest) => {
    isLoading.value = true;
    error.value = null;

    try {
      // Transform to match API expectations
      const resetData = {
        token: `${data.email}/${data.code}`, // Adjust as needed for your API
        new_password: data.newPassword,
        confirm_password: data.newPassword
      };

      await apiService.auth.confirmPasswordReset(
        resetData.token, 
        resetData.new_password, 
        resetData.confirm_password
      );
      return true;
    } catch (err: any) {
      error.value = err.response?.data?.detail || 'Failed to reset password. Please try again.';
      return false;
    } finally {
      isLoading.value = false;
    }
  };

  const logout = async () => {
    // If in superadmin mode, just clear the flag
    if (isSuperadminMode.value) {
      isSuperadminMode.value = false;
      localStorage.removeItem('isSuperadminMode');
    }
    
    if (refreshToken.value) {
      try {
        await apiService.auth.logout(refreshToken.value);
      } catch (err) {
        console.error('Logout error:', err);
      }
    }

    user.value = null;
    clearTokens();
    error.value = null;
  };

  // Initialize superadmin mode from localStorage only if user is actually superadmin
  if (localStorage.getItem('isSuperadminMode') === 'true' && 
      localStorage.getItem('userRole') === 'admin' &&
      user.value?.email === 'admin@ecomflex.com') {
    isSuperadminMode.value = true;
  } else {
    // Clear superadmin mode if conditions aren't met
    isSuperadminMode.value = false;
    localStorage.removeItem('isSuperadminMode');
  }

  return {
    // State
    user,
    accessToken,
    refreshToken,
    isLoading,
    error,
    isSuperadminMode,
    
    // Getters
    isAuthenticated,
    isSuperAdmin,
    isInfluencer,
    isPublicUser,
    isApprovedInfluencer,
    
    // Actions
    login,
    register,
    registerInfluencer,
    googleAuth,
    refreshAuth,
    fetchUserProfile,
    requestPasswordReset,
    verifyResetCode,
    resetPassword,
    logout,
    setAccessToken,
    setRefreshToken,
    setUser,
    clearTokens,
    
    // Special superadmin API
    getSuperadminData
  };
});