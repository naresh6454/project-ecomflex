// File: /ecomflex-frontend/src/services/influencer.service.ts

import api from './api';
import axios, { AxiosError } from 'axios';

// Type for API error responses
interface ApiErrorResponse {
  error?: string;
  message?: string;
}

export const influencerService = {
  // Get the dashboard data for an influencer
  getDashboardData: async () => {
    try {
      console.log('Fetching dashboard data from API...');
      const response = await api.get('influencer/dashboard');
      console.log('Dashboard data received:', response.data);
      return response.data.data;
    } catch (error) {
      console.error('Error fetching dashboard data:', error);
      // Type guard for AxiosError
      if (axios.isAxiosError(error)) {
        const axiosError = error as AxiosError<ApiErrorResponse>;
        if (axiosError.response) {
          console.error('Response status:', axiosError.response.status);
          console.error('Response data:', axiosError.response.data);
        } else if (axiosError.request) {
          console.error('No response received. Request:', axiosError.request);
        }
      }
      throw error;
    }
  },
  
  // Get the referral code
  getReferralCode: async () => {
    try {
      console.log('Fetching referral code...');
      const response = await api.get('influencer/referral-code');
      console.log('Referral code received:', response.data);
      return response.data.data.referral_code;
    } catch (error) {
      console.error('Error fetching referral code:', error);
      if (axios.isAxiosError(error)) {
        const axiosError = error as AxiosError<ApiErrorResponse>;
        if (axiosError.response) {
          console.error('Response status:', axiosError.response.status);
          console.error('Response data:', axiosError.response.data);
        }
      }
      throw error;
    }
  },
  
  // Get referral statistics
  getReferralStats: async () => {
    try {
      console.log('Fetching referral stats...');
      const response = await api.get('influencer/stats');
      console.log('Stats received:', response.data);
      return response.data.data;
    } catch (error) {
      console.error('Error fetching referral stats:', error);
      if (axios.isAxiosError(error)) {
        const axiosError = error as AxiosError<ApiErrorResponse>;
        if (axiosError.response) {
          console.error('Response status:', axiosError.response.status);
          console.error('Response data:', axiosError.response.data);
        }
      }
      throw error;
    }
  },
  
  // Get recent referrals
  getRecentReferrals: async (limit = 5) => {
    try {
      console.log(`Fetching recent referrals with limit ${limit}...`);
      const response = await api.get(`influencer/recent-referrals?limit=${limit}`);
      console.log('Referrals received:', response.data);
      return response.data.data.referrals;
    } catch (error) {
      console.error('Error fetching recent referrals:', error);
      if (axios.isAxiosError(error)) {
        const axiosError = error as AxiosError<ApiErrorResponse>;
        if (axiosError.response) {
          console.error('Response status:', axiosError.response.status);
          console.error('Response data:', axiosError.response.data);
        }
      }
      throw error;
    }
  }
};

export default influencerService;