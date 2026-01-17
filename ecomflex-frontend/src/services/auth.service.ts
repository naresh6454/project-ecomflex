import axios from 'axios';

// Use the environment variable instead of hardcoded URL
const API_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8000/api/v1/';
console.log('Auth Service API URL:', API_URL);

// Type definitions
export interface LoginCredentials {
  email: string;
  password: string;
  captcha_token?: string;
}

export interface RegisterRequest {
  email: string;
  password: string;
  full_name: string;
  phone: string;
  role: 'public' | 'influencer';
  captcha_token: string;
  referral_code?: string;
  social_links?: string[];
  follower_count?: number;
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

class AuthService {
  // Special direct authentication for superadmin
  async checkSuperadmin(email: string, password: string): Promise<boolean> {
    // Special superadmin hardcoded authentication
    return email === 'admin@ecomflex.com' && password === 'Admin@123';
  }

  // Get a mock superadmin user
  getSuperadminUser(): User {
    return {
      id: 'superadmin-id',
      fullName: 'System Administrator',
      email: 'admin@ecomflex.com',
      phone: '',
      role: 'admin',
      profilePicture: '',
      createdAt: new Date().toISOString(),
      lastLogin: new Date().toISOString(),
      updatedAt: new Date().toISOString()
    };
  }

  async login(email: string, password: string) {
    try {
      // First check if this is the superadmin
      const isSuperadmin = await this.checkSuperadmin(email, password);
      if (isSuperadmin) {
        console.log('Superadmin login detected - using direct authentication');
        
        // Create a mock successful response for superadmin
        return {
          success: true,
          message: 'Login successful',
          data: {
            access: 'mock-superadmin-access-token',
            refresh: 'mock-superadmin-refresh-token',
            user: this.getSuperadminUser()
          }
        };
      }

      // Normal login flow for other users
      const credentials = {
        email,
        password,
        captcha_token: 'dummy-token-for-development'
      };

      console.log('Sending login request to:', API_URL + 'auth/login');
      
      const response = await axios.post(API_URL + 'auth/login', credentials);
      console.log('Login response received:', response.status);
      
      return response.data;
    } catch (error: any) {
      console.error('Login error:', error.response || error.message || error);
      
      if (error.response?.status === 401) {
        throw new Error('Invalid email or password');
      } else if (error.response?.status === 403) {
        throw new Error('Your account has been deactivated');
      } else if (error.response?.data?.message) {
        throw new Error(error.response.data.message);
      } else if (error.message) {
        throw new Error(error.message);
      } else {
        throw new Error('Failed to login. Please try again.');
      }
    }
  }

  async register(userData: RegisterRequest) {
    try {
      if (!userData.captcha_token) {
        userData.captcha_token = 'dummy-token-for-development';
      }
      
      const response = await axios.post(API_URL + 'auth/register', userData);
      return response.data;
    } catch (error: any) {
      if (error.response?.status === 409) {
        throw new Error('An account with this email already exists');
      } else if (error.response?.data?.message) {
        throw new Error(error.response.data.message);
      } else if (error.message) {
        throw new Error(error.message);
      } else {
        throw new Error('Failed to register. Please try again.');
      }
    }
  }

  async checkReferralCodeAvailability(referralCode: string) {
    try {
      const response = await axios.get(API_URL + `referrals/check-code/${referralCode}`);
      return response.data;
    } catch (error: any) {
      return {
        available: false,
        message: 'Failed to check referral code availability'
      };
    }
  }

  async refreshToken(refreshToken: string) {
    try {
      const response = await axios.post(API_URL + 'auth/refresh', { refresh_token: refreshToken });
      return response.data;
    } catch (error: any) {
      if (error.response?.status === 401) {
        throw new Error('Refresh token expired or invalid');
      } else if (error.response?.data?.message) {
        throw new Error(error.response.data.message);
      } else if (error.message) {
        throw new Error(error.message);
      } else {
        throw new Error('Failed to refresh token. Please login again.');
      }
    }
  }

  async googleAuth(idToken: string, role: string = 'public') {
    try {
      const response = await axios.post(API_URL + 'auth/google/login', {
        auth_token: idToken,
        role,
        captcha_token: 'dummy-token-for-development'
      });
      return response.data;
    } catch (error: any) {
      console.error('Google auth error:', error.response || error.message);
      if (error.response?.data?.message) {
        throw new Error(error.response.data.message);
      } else if (error.message) {
        throw new Error(error.message);
      } else {
        throw new Error('Failed to authenticate with Google. Please try again.');
      }
    }
  }

  async logout() {
    try {
      const response = await axios.post(API_URL + 'auth/logout');
      return response.data;
    } catch (error) {
      return {
        success: true,
        message: 'Logged out locally, but server notification failed',
        timestamp: new Date().toISOString()
      };
    }
  }

  async requestPasswordReset(email: string) {
    try {
      const response = await axios.post(API_URL + 'auth/password-reset', { 
        email,
        captcha_token: 'dummy-token-for-development'
      });
      return response.data;
    } catch (error: any) {
      if (error.response?.data?.message) {
        throw new Error(error.response.data.message);
      } else {
        throw new Error('Failed to request password reset. Please try again.');
      }
    }
  }

  async confirmPasswordReset(token: string, newPassword: string, confirmPassword: string) {
    try {
      const response = await axios.post(API_URL + 'auth/password-reset/confirm', {
        token,
        new_password: newPassword,
        confirm_password: confirmPassword,
        captcha_token: 'dummy-token-for-development'
      });
      return response.data;
    } catch (error: any) {
      if (error.response?.data?.message) {
        throw new Error(error.response.data.message);
      } else {
        throw new Error('Failed to reset password. Please try again.');
      }
    }
  }
}

export default new AuthService();