<template>
  <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
    <!-- Loading State -->
    <div v-if="loading" class="flex justify-center items-center h-64">
      <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
      <strong class="font-bold">Error!</strong>
      <span class="block sm:inline"> {{ error }}</span>
      <button @click="retryFetch" class="mt-2 bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 rounded">
        Retry
      </button>
    </div>

    <!-- Content (when loaded) -->
    <div v-else>
      <!-- Page Header -->
      <div 
        v-motion
        :initial="{ opacity: 0, y: -20 }"
        :enter="{ opacity: 1, y: 0 }"
        class="md:flex md:items-center md:justify-between mb-8">
        <div class="flex-1 min-w-0">
          <h1 class="text-2xl font-bold leading-7 text-gray-900 sm:text-3xl sm:truncate">
            Influencer Dashboard
          </h1>
          <p class="mt-1 text-sm text-gray-500">
            Welcome back, {{ user.fullName }}! Here's an overview of your referral performance.
          </p>
        </div>
        <div class="mt-4 flex md:mt-0 md:ml-4">
          <button 
            type="button" 
            class="inline-flex items-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
            <svg class="-ml-1 mr-2 h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
            </svg>
            Download Report
          </button>
        </div>
      </div>

      <!-- Referral Code Card -->
      <div 
        v-motion
        :initial="{ opacity: 0, scale: 0.95 }"
        :enter="{ opacity: 1, scale: 1 }"
        class="bg-white shadow sm:rounded-lg mb-8">
        <div class="px-4 py-5 sm:p-6">
          <h3 class="text-lg leading-6 font-medium text-gray-900">
            Your Referral Code
          </h3>
          <div class="mt-2 max-w-xl text-sm text-gray-500">
            <p>Share this code with your audience. They'll get 100% cashback, and you'll earn commission.</p>
          </div>
          <div class="mt-3">
            <div class="flex items-center">
              <div 
                class="text-2xl font-bold text-blue-600 bg-blue-50 px-4 py-2 rounded-md border border-blue-200"
                v-motion
                :initial="{ opacity: 0, x: -20 }"
                :enter="{ opacity: 1, x: 0, transition: { delay: 200 } }">
                {{ referralCode }}
              </div>
              <button 
                @click="copyReferralCode"
                class="ml-3 inline-flex items-center px-3 py-2 border border-transparent shadow-sm text-sm leading-4 font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors">
                <svg class="-ml-0.5 mr-2 h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z" />
                </svg>
                {{ copied ? 'Copied!' : 'Copy Code' }}
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="mb-8">
        <h2 class="text-lg font-medium text-gray-900 mb-4">Performance Overview</h2>
        <div class="grid grid-cols-1 gap-5 sm:grid-cols-2 lg:grid-cols-4">
          <StatsCard
            title="Total Referrals"
            :value="stats.totalReferrals"
            :icon="LinkIcon"
            type="default"
            format="number"
            actionText="View details"
            actionUrl="/influencer/referrals"
            :delay="0"
          />
          
          <StatsCard
            title="Approved Bookings"
            :value="stats.approvedBookings"
            :icon="CheckIcon"
            type="success"
            format="number"
            actionText="View approved"
            actionUrl="/influencer/referrals?status=approved"
            :delay="1"
          />
          
          <StatsCard
            title="Pending Bookings"
            :value="stats.pendingBookings"
            :icon="ClockIcon"
            type="warning"
            format="number"
            actionText="View pending"
            actionUrl="/influencer/referrals?status=pending"
            :delay="2"
          />
          
          <StatsCard
            title="Total Earnings"
            :value="stats.totalEarnings"
            :icon="CashIcon"
            type="primary"
            format="currency"
            actionText="View earnings"
            actionUrl="/influencer/earnings"
            :delay="3"
          />
        </div>
      </div>

      <!-- Share Your Referral Link -->
      <div class="mb-8">
        <h2 class="text-lg font-medium text-gray-900 mb-4">Share Your Referral Link</h2>
        <ShareOptions :referral-link="referralLink" />
      </div>

      <!-- Recent Referrals -->
      <div class="mb-8">
        <div class="flex items-center justify-between mb-4">
          <h2 class="text-lg font-medium text-gray-900">Recent Referrals</h2>
          <router-link 
            to="/influencer/referrals" 
            class="text-sm font-medium text-blue-600 hover:text-blue-500 transition-colors">
            View all
          </router-link>
        </div>
        
        <div class="bg-white shadow overflow-hidden sm:rounded-md">
          <ul class="divide-y divide-gray-200">
            <li v-for="(referral, index) in recentReferrals" :key="referral.id"
              v-motion
              :initial="{ opacity: 0, y: 20 }"
              :enter="{ opacity: 1, y: 0, transition: { delay: index * 100 } }">
              <div class="px-4 py-4 flex items-center sm:px-6">
                <div class="min-w-0 flex-1 sm:flex sm:items-center sm:justify-between">
                  <div>
                    <div class="flex text-sm">
                      <p class="font-medium text-blue-600 truncate">{{ referral.user }}</p>
                      <p class="ml-1 flex-shrink-0 font-normal text-gray-500">
                        booked <span class="text-gray-900 font-medium">{{ referral.product }}</span>
                      </p>
                    </div>
                    <div class="mt-2 flex">
                      <div class="flex items-center text-sm text-gray-500">
                        <svg class="flex-shrink-0 mr-1.5 h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                          <path fill-rule="evenodd" d="M6 2a1 1 0 00-1 1v1H4a2 2 0 00-2 2v10a2 2 0 002 2h12a2 2 0 002-2V6a2 2 0 00-2-2h-1V3a1 1 0 10-2 0v1H7V3a1 1 0 00-1-1zm0 5a1 1 0 000 2h8a1 1 0 100-2H6z" clip-rule="evenodd" />
                        </svg>
                        <span>{{ referral.date }}</span>
                      </div>
                    </div>
                  </div>
                  <div class="mt-4 flex-shrink-0 sm:mt-0">
                    <div class="flex overflow-hidden">
                      <span 
                        :class="[
                          referral.status === 'Approved' ? 'bg-green-100 text-green-800' : 
                          referral.status === 'Pending' ? 'bg-yellow-100 text-yellow-800' : 
                          'bg-red-100 text-red-800',
                          'px-2 py-1 text-xs font-medium rounded-full'
                        ]">
                        {{ referral.status }}
                      </span>
                    </div>
                  </div>
                </div>
              </div>
            </li>
            <li v-if="recentReferrals.length === 0" class="px-4 py-5 text-center text-sm text-gray-500">
              No referrals yet. Share your referral code to start earning!
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed, h } from 'vue'
import StatsCard from '@/components/influencer/StatsCard.vue'
import ShareOptions from '@/components/influencer/ShareOptions.vue'
import { useReferralStore } from '@/stores/referral'
import { useAuthStore } from '@/stores/auth'
import { storeToRefs } from 'pinia'
import api from '@/services/api'

// Define interfaces for our data types
interface Referral {
  id: string;
  user: string;
  product: string;
  date: string;
  status: string;
  earnings?: number;
}

interface User {
  id: string;
  fullName: string;
  email: string;
  phone: string;
  role: "public" | "admin" | "influencer";
  profilePicture: string;
  createdAt: string;
  lastLogin: string;
  updatedAt?: string;
}

// Icon components
const LinkIcon = (props: any) => {
  return h('svg', {
    xmlns: "http://www.w3.org/2000/svg",
    fill: "none",
    viewBox: "0 0 24 24",
    "stroke-width": "1.5",
    stroke: "currentColor",
    ...props
  }, [
    h('path', {
      "stroke-linecap": "round",
      "stroke-linejoin": "round",
      d: "M13.19 8.688a4.5 4.5 0 011.242 7.244l-4.5 4.5a4.5 4.5 0 01-6.364-6.364l1.757-1.757m13.35-.622l1.757-1.757a4.5 4.5 0 00-6.364-6.364l-4.5 4.5a4.5 4.5 0 001.242 7.244"
    })
  ])
}

const CheckIcon = (props: any) => {
  return h('svg', {
    xmlns: "http://www.w3.org/2000/svg",
    fill: "none",
    viewBox: "0 0 24 24",
    "stroke-width": "1.5",
    stroke: "currentColor",
    ...props
  }, [
    h('path', {
      "stroke-linecap": "round",
      "stroke-linejoin": "round",
      d: "M9 12.75L11.25 15 15 9.75M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
    })
  ])
}

const ClockIcon = (props: any) => {
  return h('svg', {
    xmlns: "http://www.w3.org/2000/svg",
    fill: "none",
    viewBox: "0 0 24 24",
    "stroke-width": "1.5",
    stroke: "currentColor",
    ...props
  }, [
    h('path', {
      "stroke-linecap": "round",
      "stroke-linejoin": "round",
      d: "M12 6v6h4.5m4.5 0a9 9 0 11-18 0 9 9 0 0118 0z"
    })
  ])
}

const CashIcon = (props: any) => {
  return h('svg', {
    xmlns: "http://www.w3.org/2000/svg",
    fill: "none",
    viewBox: "0 0 24 24",
    "stroke-width": "1.5",
    stroke: "currentColor",
    ...props
  }, [
    h('path', {
      "stroke-linecap": "round",
      "stroke-linejoin": "round",
      d: "M2.25 18.75a60.07 60.07 0 0115.797 2.101c.727.198 1.453-.342 1.453-1.096V18.75M3.75 4.5v.75A.75.75 0 013 6h-.75m0 0v-.375c0-.621.504-1.125 1.125-1.125H20.25M2.25 6v9m18-10.5v.75c0 .414.336.75.75.75h.75m-1.5-1.5h.375c.621 0 1.125.504 1.125 1.125v9.75c0 .621-.504 1.125-1.125 1.125h-.375m1.5-1.5H21a.75.75 0 00-.75.75v.75m0 0H3.75m0 0h-.375a1.125 1.125 0 01-1.125-1.125V15m1.5 1.5v-.75A.75.75 0 003 15h-.75M15 10.5a3 3 0 11-6 0 3 3 0 016 0zm3 0h.008v.008H18V10.5zm-12 0h.008v.008H6V10.5z"
    })
  ])
}

// Store setup
const referralStore = useReferralStore();
const authStore = useAuthStore();

// Destructure store properties
const { 
  referralCode, 
  loading, 
  error, 
  stats, 
  referralLink 
} = storeToRefs(referralStore);

// Need to keep recentReferrals reactive but also properly typed
const recentReferrals = computed<Referral[]>(() => {
  return referralStore.recentReferrals as Referral[];
});

// User data from auth store with fallback
const user = computed<User>(() => {
  return authStore.user as User || { 
    id: '0',
    fullName: 'Influencer',
    email: 'user@example.com',
    phone: '',
    role: 'influencer' as const,
    profilePicture: '',
    createdAt: '',
    lastLogin: '',
  };
});

// Copy referral code functionality
const copied = ref(false);
const copyReferralCode = () => {
  if (referralCode.value) {
    navigator.clipboard.writeText(referralCode.value);
    copied.value = true;
    setTimeout(() => {
      copied.value = false;
    }, 2000);
  }
};

// Retry fetch function
const retryFetch = async () => {
  console.log('Retrying data fetch...');
  try {
    // Debug auth token first
    console.log('Debugging authentication token:');
    api.debugToken();
    
    await referralStore.fetchDashboardData();
  } catch (err) {
    console.error('Retry failed:', err);
  }
};

// Fetch data on component mount
onMounted(async () => {
  console.log('Dashboard component mounted, fetching data...');
  try {
    // Debug token to see if auth is working
    const tokenValid = api.debugToken();
    console.log('Token valid:', tokenValid);
    
    if (!tokenValid) {
      console.warn('Invalid token - you may need to log in again');
    }
    
    // Always attempt to fetch data anyway
    await referralStore.fetchDashboardData();
    
    console.log('Data fetched successfully');
    console.log('Current store state:', {
      referralCode: referralCode.value,
      stats: stats.value,
      recentReferrals: recentReferrals.value.length > 0 ? 'Has data' : 'Empty'
    });
  } catch (err) {
    console.error('Error in onMounted:', err);
  }
});
</script>