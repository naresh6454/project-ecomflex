<template>
  <div class="influencer-referrals">
    <!-- Loading State -->
    <div v-if="isLoading" class="flex items-center justify-center py-12">
      <div class="text-center">
        <div class="inline-flex animate-spin rounded-full h-12 w-12 border-b-2 border-accent mb-4"></div>
        <p class="text-gray-600">Loading referral data...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
      <p class="text-red-700">{{ error }}</p>
      <button 
        @click="loadReferralData" 
        class="mt-3 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700"
      >
        Retry
      </button>
    </div>

    <!-- Main Content -->
    <div v-else>
    <!-- Page Header -->
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-900 mb-2">Referral Tracking</h1>
      <p class="text-gray-600">Track and manage all your referrals and their performance.</p>
    </div>

    <!-- Referral Stats Overview -->
    <div class="bg-white rounded-xl shadow-lg p-6 mb-8">
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <!-- Total Referrals -->
        <div class="border-r border-gray-200 pr-6 last:border-r-0 last:pr-0">
          <h3 class="text-gray-600 text-sm mb-1">Total Referrals</h3>
          <div class="flex items-baseline">
            <span class="text-3xl font-bold text-gray-900 mr-2">{{ stats.totalReferrals }}</span>
            <span class="text-sm text-green-500">
              <span class="font-medium">+{{ stats.newReferrals }}</span> this month
            </span>
          </div>
        </div>

        <!-- Conversion Rate -->
        <div class="border-r border-gray-200 pr-6 last:border-r-0 last:pr-0">
          <h3 class="text-gray-600 text-sm mb-1">Conversion Rate</h3>
          <div class="flex items-baseline">
            <span class="text-3xl font-bold text-gray-900 mr-2">{{ stats.conversionRate }}%</span>
            <span class="text-sm" :class="stats.conversionTrend > 0 ? 'text-green-500' : 'text-red-500'">
              <span class="font-medium">{{ stats.conversionTrend > 0 ? '+' : '' }}{{ stats.conversionTrend }}%</span> vs last month
            </span>
          </div>
        </div>

        <!-- Total Earnings -->
        <div>
          <h3 class="text-gray-600 text-sm mb-1">Total Earnings</h3>
          <div class="flex items-baseline">
            <span class="text-3xl font-bold text-gray-900 mr-2">${{ stats.totalEarnings }}</span>
            <span class="text-sm text-green-500">
              <span class="font-medium">${{ stats.pendingEarnings }}</span> pending
            </span>
          </div>
        </div>
      </div>
    </div>

    <!-- Monthly Performance -->
    <div class="bg-white rounded-xl shadow-lg p-6 mb-8">
      <h2 class="text-xl font-bold text-gray-900 mb-6">Monthly Performance</h2>
      
      <div class="space-y-6">
        <!-- Current Month -->
        <div>
          <div class="flex justify-between items-center mb-2">
            <div>
              <h3 class="text-base font-medium text-gray-700">{{ currentMonth }}</h3>
              <p class="text-sm text-gray-500">{{ monthlyTarget.current }} / {{ monthlyTarget.target }} referrals</p>
            </div>
            <span class="text-sm font-medium" :class="monthlyTarget.progress >= 50 ? 'text-green-600' : 'text-amber-500'">
              {{ monthlyTarget.progress }}%
            </span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2.5">
            <div class="bg-accent h-2.5 rounded-full" :style="{ width: `${monthlyTarget.progress}%` }"></div>
          </div>
        </div>

        <!-- Past Months -->
        <div v-for="(month, index) in pastMonths" :key="index">
          <div class="flex justify-between items-center mb-2">
            <div>
              <h3 class="text-base font-medium text-gray-700">{{ month.name }}</h3>
              <p class="text-sm text-gray-500">{{ month.referrals }} / {{ month.target }} referrals</p>
            </div>
            <span class="text-sm font-medium" :class="month.progress >= 100 ? 'text-green-600' : month.progress >= 50 ? 'text-amber-500' : 'text-gray-600'">
              {{ month.progress }}%
            </span>
          </div>
          <div class="w-full bg-gray-200 rounded-full h-2.5">
            <div 
              class="h-2.5 rounded-full" 
              :class="month.progress >= 100 ? 'bg-green-500' : 'bg-accent'"
              :style="{ width: `${month.progress}%` }"
            ></div>
          </div>
        </div>
      </div>
    </div>

    <!-- Referral Code Sharing -->
    <div class="bg-white rounded-xl shadow-lg p-6 mb-8">
      <h2 class="text-xl font-bold text-gray-900 mb-4">Share Your Referral Code</h2>
      
      <div class="flex flex-col md:flex-row items-start md:items-center justify-between">
        <div class="mb-4 md:mb-0">
          <p class="text-gray-600 mb-2">Your unique referral code:</p>
          <div class="flex items-center">
            <div class="bg-gray-100 rounded-lg px-4 py-3 mr-3">
              <span class="font-mono text-xl font-bold text-gray-800">{{ referralCode }}</span>
            </div>
            <button 
              @click="copyReferralCode" 
              class="px-3 py-2 bg-accent text-white rounded-lg hover:bg-brand-dark transition-colors duration-200 flex items-center"
              :class="{ 'animate-bounce': isCopied }"
            >
              <span>{{ isCopied ? 'Copied!' : 'Copy' }}</span>
              <svg v-if="!isCopied" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
            </button>
          </div>
        </div>

        <div class="flex space-x-3">
          <button 
            @click="shareViaWhatsApp"
            class="px-3 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition-colors duration-200 flex items-center"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="mr-2" viewBox="0 0 16 16">
              <path d="M13.601 2.326A7.854 7.854 0 0 0 7.994 0C3.627 0 .068 3.558.064 7.926c0 1.399.366 2.76 1.057 3.965L0 16l4.204-1.102a7.933 7.933 0 0 0 3.79.965h.004c4.368 0 7.926-3.558 7.93-7.93A7.898 7.898 0 0 0 13.6 2.326zM7.994 14.521a6.573 6.573 0 0 1-3.356-.92l-.24-.144-2.494.654.666-2.433-.156-.251a6.56 6.56 0 0 1-1.007-3.505c0-3.626 2.957-6.584 6.591-6.584a6.56 6.56 0 0 1 4.66 1.931 6.557 6.557 0 0 1 1.928 4.66c-.004 3.639-2.961 6.592-6.592 6.592zm3.615-4.934c-.197-.099-1.17-.578-1.353-.646-.182-.065-.315-.099-.445.099-.133.197-.513.646-.627.775-.114.133-.232.148-.43.05-.197-.1-.836-.308-1.592-.985-.59-.525-.985-1.175-1.103-1.372-.114-.198-.011-.304.088-.403.087-.088.197-.232.296-.346.1-.114.133-.198.198-.33.065-.134.034-.248-.015-.347-.05-.099-.445-1.076-.612-1.47-.16-.389-.323-.335-.445-.34-.114-.007-.247-.007-.38-.007a.729.729 0 0 0-.529.247c-.182.198-.691.677-.691 1.654 0 .977.71 1.916.81 2.049.098.133 1.394 2.132 3.383 2.992.47.205.84.326 1.129.418.475.152.904.129 1.246.08.38-.058 1.171-.48 1.338-.943.164-.464.164-.86.114-.943-.049-.084-.182-.133-.38-.232z"/>
            </svg>
            Share on WhatsApp
          </button>
          
          <button 
            @click="shareViaTwitter"
            class="px-3 py-2 bg-blue-400 text-white rounded-lg hover:bg-blue-500 transition-colors duration-200 flex items-center"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" class="mr-2" viewBox="0 0 16 16">
              <path d="M5.026 15c6.038 0 9.341-5.003 9.341-9.334 0-.14 0-.282-.006-.422A6.685 6.685 0 0 0 16 3.542a6.658 6.658 0 0 1-1.889.518 3.301 3.301 0 0 0 1.447-1.817 6.533 6.533 0 0 1-2.087.793A3.286 3.286 0 0 0 7.875 6.03a9.325 9.325 0 0 1-6.767-3.429 3.289 3.289 0 0 0 1.018 4.382A3.323 3.323 0 0 1 .64 6.575v.045a3.288 3.288 0 0 0 2.632 3.218 3.203 3.203 0 0 1-.865.115 3.23 3.23 0 0 1-.614-.057 3.283 3.283 0 0 0 3.067 2.277A6.588 6.588 0 0 1 .78 13.58a6.32 6.32 0 0 1-.78-.045A9.344 9.344 0 0 0 5.026 15z"/>
            </svg>
            Share on Twitter
          </button>
          
          <button 
            @click="generateReferralLink"
            class="px-3 py-2 border border-accent text-accent rounded-lg hover:bg-accent hover:text-white transition-colors duration-200 flex items-center"
          >
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
            </svg>
            Get Referral Link
          </button>
        </div>
      </div>
    </div>

    <!-- Referrals Table -->
    <div class="mb-8">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-xl font-bold text-gray-900">Your Referrals</h2>
        
        <div class="flex">
          <div class="relative mr-4">
            <input 
              type="text" 
              placeholder="Search referrals..." 
              class="pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-accent/30 focus:border-accent"
            >
            <div class="absolute left-3 top-2.5 text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>

          <select class="px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-accent/30 focus:border-accent">
            <option value="all">All Time</option>
            <option value="month">This Month</option>
            <option value="week">This Week</option>
            <option value="day">Today</option>
          </select>
        </div>
      </div>

      <div class="bg-white rounded-xl shadow-lg overflow-hidden">
        <div class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  User
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Referred Date
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Products Booked
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Total Value
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Status
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Earnings
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-if="referrals.length === 0" class="hover:bg-gray-50">
                <td colspan="6" class="px-6 py-8 text-center text-gray-500">
                  No referrals yet. Start sharing your referral code to earn rewards!
                </td>
              </tr>
              <tr v-for="referral in referrals" :key="referral.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="h-10 w-10 rounded-full bg-gray-200 flex items-center justify-center text-gray-500">
                      {{ (referral.user?.name || referral.name || 'U').charAt(0).toUpperCase() }}
                    </div>
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900">{{ referral.user?.name || referral.name || 'Unknown' }}</div>
                      <div class="text-sm text-gray-500">{{ referral.user?.email || referral.email || 'N/A' }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">{{ formatDate(referral.referral_date || referral.date) }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">{{ referral.products_booked || referral.productsBooked || 0 }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">${{ (referral.total_value || referral.totalValue || 0).toFixed(2) }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span 
                    class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                    :class="{
                      'bg-green-100 text-green-800': (referral.status || '').toLowerCase() === 'active',
                      'bg-yellow-100 text-yellow-800': (referral.status || '').toLowerCase() === 'pending',
                      'bg-red-100 text-red-800': (referral.status || '').toLowerCase() === 'inactive',
                      'bg-gray-100 text-gray-800': !referral.status
                    }"
                  >
                    {{ referral.status || 'Pending' }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm font-medium" :class="(referral.earnings || referral.totalValue || 0) > 0 ? 'text-green-600' : 'text-gray-500'">
                    ${{ (referral.earnings || 0).toFixed(2) }}
                  </div>
                </td>
              </tr>
            </tbody>
          </table>
        </div>

        <!-- Pagination -->
        <div class="px-6 py-4 bg-gray-50 border-t border-gray-200 flex items-center justify-between">
          <div class="flex-1 flex justify-between sm:hidden">
            <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
              Previous
            </button>
            <button class="ml-3 relative inline-flex items-center px-4 py-2 border border-gray-300 text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50">
              Next
            </button>
          </div>
          <div class="hidden sm:flex-1 sm:flex sm:items-center sm:justify-between">
            <div>
              <p class="text-sm text-gray-700">
                Showing <span class="font-medium">1</span> to <span class="font-medium">10</span> of <span class="font-medium">{{ referrals.length }}</span> results
              </p>
            </div>
            <div>
              <nav class="relative z-0 inline-flex rounded-md shadow-sm -space-x-px" aria-label="Pagination">
                <button class="relative inline-flex items-center px-2 py-2 rounded-l-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
                  <span class="sr-only">Previous</span>
                  <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                    <path fill-rule="evenodd" d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z" clip-rule="evenodd" />
                  </svg>
                </button>
                <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50">
                  1
                </button>
                <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-accent text-sm font-medium text-white">
                  2
                </button>
                <button class="relative inline-flex items-center px-4 py-2 border border-gray-300 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50">
                  3
                </button>
                <button class="relative inline-flex items-center px-2 py-2 rounded-r-md border border-gray-300 bg-white text-sm font-medium text-gray-500 hover:bg-gray-50">
                  <span class="sr-only">Next</span>
                  <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                    <path fill-rule="evenodd" d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z" clip-rule="evenodd" />
                  </svg>
                </button>
              </nav>
            </div>
          </div>
        </div>
      </div>
    </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import referralService from '../../services/referral.service';
import { useAuthStore } from '../../stores/auth';

// Get authenticated user info
const authStore = useAuthStore();

// Referral code handling
const referralCode = ref('');
const isCopied = ref(false);
const isLoading = ref(true);
const error = ref('');

// Stats data
const stats = ref({
  totalReferrals: 0,
  newReferrals: 0,
  conversionRate: 0,
  conversionTrend: 0,
  totalEarnings: 0,
  pendingEarnings: 0
});

// Monthly performance data
const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
const currentMonth = computed(() => {
  const date = new Date();
  return months[date.getMonth()];
});

const monthlyTarget = ref({
  current: 0,
  target: 30,
  progress: 0
});

const pastMonths = ref<Array<any>>([]);

// Referrals data
const referrals = ref<Array<any>>([]);

// Search and filter
const searchQuery = ref('');
const timeFilter = ref('all');

// Format date
const formatDate = (date: Date | string) => {
  const dateObj = typeof date === 'string' ? new Date(date) : date;
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(dateObj);
};

// Copy referral code
const copyReferralCode = () => {
  navigator.clipboard.writeText(referralCode.value);
  isCopied.value = true;
  setTimeout(() => {
    isCopied.value = false;
  }, 2000);
};

// Share via WhatsApp
const shareViaWhatsApp = () => {
  const text = `Get 100% cashback on products! Use my referral code ${referralCode.value} on Ecomflex. ${window.location.origin}/referral/${referralCode.value}`;
  window.open(`https://wa.me/?text=${encodeURIComponent(text)}`, '_blank');
};

// Share via Twitter
const shareViaTwitter = () => {
  const text = `Get 100% cashback on products! Use my referral code ${referralCode.value} on Ecomflex.`;
  window.open(`https://twitter.com/intent/tweet?text=${encodeURIComponent(text)}&url=${encodeURIComponent(`${window.location.origin}/referral/${referralCode.value}`)}`, '_blank');
};

// Generate shareable referral link
const generateReferralLink = () => {
  const link = `${window.location.origin}/referral/${referralCode.value}`;
  navigator.clipboard.writeText(link);
  isCopied.value = true;
  setTimeout(() => {
    isCopied.value = false;
  }, 2000);
};

// Load referral data from backend
const loadReferralData = async () => {
  try {
    isLoading.value = true;
    error.value = '';

    // Fetch referral code
    try {
      const codeResponse = await referralService.getReferralCode();
      referralCode.value = codeResponse.data?.referral_code || 'N/A';
    } catch (err) {
      console.warn('Failed to load referral code:', err);
    }

    // Fetch referral statistics
    try {
      const statsResponse = await referralService.getReferralStats();
      if (statsResponse.data) {
        stats.value = {
          totalReferrals: statsResponse.data.total_referrals || 0,
          newReferrals: statsResponse.data.new_referrals_this_month || 0,
          conversionRate: statsResponse.data.conversion_rate || 0,
          conversionTrend: statsResponse.data.conversion_trend || 0,
          totalEarnings: statsResponse.data.total_earnings || 0,
          pendingEarnings: statsResponse.data.pending_earnings || 0
        };

        // Update monthly target
        monthlyTarget.value = {
          current: statsResponse.data.new_referrals_this_month || 0,
          target: 30,
          progress: Math.round(((statsResponse.data.new_referrals_this_month || 0) / 30) * 100)
        };
      }
    } catch (err) {
      console.warn('Failed to load referral stats:', err);
    }

    // Fetch referrals
    try {
      const referralsResponse = await referralService.getReferrals();
      if (referralsResponse.data && Array.isArray(referralsResponse.data)) {
        referrals.value = referralsResponse.data;
      } else if (referralsResponse.data?.referrals) {
        referrals.value = referralsResponse.data.referrals;
      }
    } catch (err) {
      console.warn('Failed to load referrals:', err);
    }
  } catch (err) {
    error.value = 'Failed to load referral data. Please try again.';
    console.error('Error loading referral data:', err);
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  // Check if user is authenticated and is an influencer
  if (!authStore.user || authStore.user.role !== 'influencer') {
    error.value = 'You must be logged in as an influencer to view referrals.';
    isLoading.value = false;
    return;
  }

  loadReferralData();
});
</script>

<style scoped>
/* Custom CSS (most styling is done with Tailwind utility classes) */
.bg-accent {
  background-color: #FFD700;
}

.text-accent {
  color: #FFD700;
}

.border-accent {
  border-color: #FFD700;
}

.hover\:bg-brand-dark:hover {
  background-color: #E9C200;
}

.hover\:text-brand-dark:hover {
  color: #E9C200;
}

.focus\:border-accent:focus {
  border-color: #FFD700;
}

.focus\:ring-accent\/30:focus {
  --tw-ring-color: rgba(255, 215, 0, 0.3);
}

/* Animation for button copy */
.animate-bounce {
  animation: bounce 0.5s;
}

@keyframes bounce {
  0%, 100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-10px);
  }
}

/* For tables on small screens */
@media (max-width: 768px) {
  .overflow-x-auto {
    -webkit-overflow-scrolling: touch;
  }
}
</style>