// File: /ecomflex-frontend/src/stores/referral.ts

import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import influencerService from '@/services/influencer.service';

// Define types for our backend responses
interface ReferralStats {
  total_referrals: number;
  approved_bookings: number;
  pending_bookings: number;
  total_earnings: number;
}

interface ReferralRecord {
  id: string;
  user: string;
  product: string;
  date: string;
  status: string;
  earnings?: number;
}

interface DashboardData {
  referral_code: string;
  referral_link: string;
  stats: ReferralStats;
  recent_referrals: ReferralRecord[];
  user: Record<string, any>;
}

// Define frontend types
interface Stats {
  totalReferrals: number;
  approvedBookings: number;
  pendingBookings: number;
  totalEarnings: number;
}

interface Referral {
  id: string;
  user: string;
  product: string;
  date: string;
  status: string;
  earnings?: number;
}

export const useReferralStore = defineStore('referral', () => {
  // State
  const referralCode = ref('');
  const loading = ref(false);
  const error = ref('');
  const stats = ref<Stats>({
    totalReferrals: 0,
    approvedBookings: 0,
    pendingBookings: 0,
    totalEarnings: 0
  });
  const recentReferrals = ref<Referral[]>([]);
  const dashboardData = ref<DashboardData | null>(null);

  // Actions
  const fetchReferralCode = async (): Promise<string> => {
    loading.value = true;
    try {
      const code = await influencerService.getReferralCode();
      referralCode.value = code;
      return code;
    } catch (err: unknown) {
      console.error('Error fetching referral code:', err);
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch referral code';
      error.value = errorMessage;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const fetchReferralStats = async (): Promise<Stats> => {
    loading.value = true;
    try {
      const response = await influencerService.getReferralStats();
      // Map backend fields to frontend fields
      stats.value = {
        totalReferrals: response.total_referrals,
        approvedBookings: response.approved_bookings,
        pendingBookings: response.pending_bookings,
        totalEarnings: response.total_earnings
      };
      return stats.value;
    } catch (err: unknown) {
      console.error('Error fetching referral stats:', err);
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch referral statistics';
      error.value = errorMessage;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const fetchRecentReferrals = async (limit = 5): Promise<Referral[]> => {
    loading.value = true;
    try {
      const referrals = await influencerService.getRecentReferrals(limit);
      // Map backend data structure to frontend structure
      recentReferrals.value = referrals.map((ref: ReferralRecord) => ({
        id: ref.id,
        user: ref.user,
        product: ref.product,
        date: ref.date,
        status: ref.status,
        earnings: ref.earnings
      }));
      return recentReferrals.value;
    } catch (err: unknown) {
      console.error('Error fetching recent referrals:', err);
      const errorMessage = err instanceof Error ? err.message : 'Failed to fetch recent referrals';
      error.value = errorMessage;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const fetchDashboardData = async (): Promise<void> => {
    loading.value = true;
    error.value = '';
    try {
      const data = await influencerService.getDashboardData();
      dashboardData.value = data;
      
      // Update the store with the data
      referralCode.value = data.referral_code;
      
      // Map the stats
      stats.value = {
        totalReferrals: data.stats.total_referrals,
        approvedBookings: data.stats.approved_bookings,
        pendingBookings: data.stats.pending_bookings,
        totalEarnings: data.stats.total_earnings
      };
      
      // Map the referrals
      recentReferrals.value = data.recent_referrals.map((ref: ReferralRecord) => ({
        id: ref.id,
        user: ref.user,
        product: ref.product,
        date: ref.date,
        status: ref.status,
        earnings: ref.earnings
      }));
    } catch (err: unknown) {
      console.error('Error fetching dashboard data:', err);
      const errorMessage = err instanceof Error 
        ? err.message 
        : (err as any)?.response?.data?.error || 'Failed to fetch dashboard data';
      error.value = errorMessage;
      throw err;
    } finally {
      loading.value = false;
    }
  };

  const retryFetch = async (): Promise<void> => {
    try {
      await fetchDashboardData();
    } catch (err) {
      console.error('Retry failed:', err);
    }
  };

  // Computed properties
  const referralLink = computed(() => {
    // If dashboardData has a referral_link, use it directly
    if (dashboardData.value?.referral_link) {
      return dashboardData.value.referral_link;
    }
    
    // Otherwise construct from referralCode
    if (!referralCode.value) return '';
    return `${window.location.origin}/ref/${referralCode.value}`;
  });

  return {
    // State
    referralCode,
    loading,
    error,
    stats,
    recentReferrals,
    dashboardData,
    
    // Actions
    fetchReferralCode,
    fetchReferralStats,
    fetchRecentReferrals,
    fetchDashboardData,
    retryFetch,
    
    // Computed
    referralLink
  };
});