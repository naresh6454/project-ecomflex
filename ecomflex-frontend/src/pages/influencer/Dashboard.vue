<template>
  <div class="influencer-dashboard">
    <!-- Loading State -->
    <div v-if="isLoading" class="flex items-center justify-center min-h-screen">
      <div class="text-center">
        <div class="animate-spin rounded-full h-16 w-16 border-b-4 border-accent mx-auto mb-4"></div>
        <p class="text-gray-600">Loading your dashboard...</p>
      </div>
    </div>

    <!-- Error State -->
    <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-6 mb-8">
      <div class="flex items-center">
        <svg class="h-6 w-6 text-red-500 mr-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
        </svg>
        <div>
          <h3 class="text-red-800 font-semibold">Error Loading Dashboard</h3>
          <p class="text-red-600 text-sm">{{ error }}</p>
        </div>
      </div>
      <button @click="loadDashboardData" class="mt-4 px-4 py-2 bg-red-600 text-white rounded-lg hover:bg-red-700">
        Retry
      </button>
    </div>

    <!-- Dashboard Content -->
    <div v-else>
      <!-- Page Header -->
      <div class="mb-8">
        <h1 class="text-3xl font-bold text-gray-900 mb-2">Influencer Dashboard</h1>
        <p class="text-gray-600">Welcome back{{ userName ? ', ' + userName : '' }}! Here's an overview of your referral performance.</p>
      </div>

      <!-- Referral Code Card -->
      <div class="bg-white rounded-xl shadow-lg p-6 mb-8 transition-all hover:shadow-xl">
        <div class="flex flex-col md:flex-row justify-between items-start md:items-center">
          <div>
            <h2 class="text-lg font-medium text-gray-700 mb-2">Your Referral Code</h2>
            <div class="flex items-center">
              <div class="bg-gray-100 rounded-lg px-4 py-3 flex items-center">
                <span class="font-mono text-xl font-bold text-gray-800">{{ referralCode || 'Loading...' }}</span>
              </div>
              <button 
                @click="copyReferralCode" 
                class="ml-3 px-4 py-2 bg-accent text-white rounded-lg hover:bg-brand-dark transition-colors duration-200 flex items-center"
                :class="{ 'animate-bounce': isCopied }"
                :disabled="!referralCode"
              >
                <span>{{ isCopied ? 'Copied!' : 'Copy' }}</span>
                <svg v-if="!isCopied" xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5H6a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2v-1M8 5a2 2 0 002 2h2a2 2 0 002-2M8 5a2 2 0 012-2h2a2 2 0 012 2m0 0h2a2 2 0 012 2v3m2 4H10m0 0l3-3m-3 3l3 3" />
                </svg>
                <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
              </button>
              <button 
                @click="shareReferralCode" 
                class="ml-3 px-4 py-2 border border-accent text-accent rounded-lg hover:bg-accent hover:text-white transition-colors duration-200 flex items-center"
                :disabled="!referralCode"
              >
                <span>Share</span>
                <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z" />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>

      <!-- Stats Cards -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <!-- Total Referrals -->
        <div class="bg-white rounded-xl shadow-lg p-6 transition-all hover:shadow-xl">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-600 text-sm">Total Referrals</p>
              <h3 class="text-3xl font-bold text-gray-900 mt-1">{{ stats.totalReferrals || 0 }}</h3>
              <p class="text-sm mt-2">
                <span class="text-green-500 font-medium">+{{ stats.newReferrals || 0 }}</span>
                <span class="text-gray-600 ml-1">this month</span>
              </p>
            </div>
            <div class="bg-accent/10 p-3 rounded-full">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-accent" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Approved Bookings -->
        <div class="bg-white rounded-xl shadow-lg p-6 transition-all hover:shadow-xl">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-600 text-sm">Approved Bookings</p>
              <h3 class="text-3xl font-bold text-gray-900 mt-1">{{ stats.approvedBookings || 0 }}</h3>
              <p class="text-sm mt-2">
                <span class="text-green-500 font-medium">{{ stats.approvalRate || 0 }}%</span>
                <span class="text-gray-600 ml-1">approval rate</span>
              </p>
            </div>
            <div class="bg-green-100 p-3 rounded-full">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-green-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Pending Bookings -->
        <div class="bg-white rounded-xl shadow-lg p-6 transition-all hover:shadow-xl">
          <div class="flex items-center justify-between">
            <div>
              <p class="text-gray-600 text-sm">Pending Bookings</p>
              <h3 class="text-3xl font-bold text-gray-900 mt-1">{{ stats.pendingBookings || 0 }}</h3>
              <p class="text-sm mt-2">
                <span class="text-amber-500 font-medium">{{ stats.pendingAmount || '$0' }}</span>
                <span class="text-gray-600 ml-1">in cashback</span>
              </p>
            </div>
            <div class="bg-amber-100 p-3 rounded-full">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-8 w-8 text-amber-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
          </div>
        </div>
      </div>

      <!-- Latest Products -->
      <div class="mb-8">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-xl font-bold text-gray-900">Products to Promote</h2>
          <a href="/influencer/products" class="text-accent hover:text-brand-dark transition-colors duration-200 flex items-center">
            View All
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </a>
        </div>

        <div v-if="latestProducts.length === 0" class="bg-gray-50 rounded-xl p-8 text-center">
          <p class="text-gray-500">No products available to promote yet.</p>
        </div>

        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <!-- Product Card -->
          <div v-for="product in latestProducts" :key="product.id" class="bg-white rounded-xl shadow-lg overflow-hidden transition-all hover:shadow-xl hover:scale-[1.02] duration-300">
            <div class="relative">
              <img :src="product.image || 'https://via.placeholder.com/400x300?text=Product'" :alt="product.name" class="w-full h-48 object-cover">
              <span class="absolute top-3 right-3 bg-accent text-white text-sm py-1 px-3 rounded-full font-medium">
                ${{ product.price }}
              </span>
            </div>
            <div class="p-5">
              <h3 class="text-lg font-bold text-gray-900 mb-2">{{ product.name }}</h3>
              <p class="text-gray-600 text-sm mb-4 line-clamp-2">{{ product.description }}</p>
              <div class="flex justify-between items-center">
                <span class="text-sm text-gray-500">
                  <span class="font-medium">{{ product.referrals || 0 }}</span> referrals
                </span>
                <button 
                  @click="shareProduct(product)"
                  class="px-3 py-1.5 bg-accent text-white rounded-lg text-sm hover:bg-brand-dark transition-colors duration-200"
                >
                  Share Link
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Bookings -->
      <div class="mb-8">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-xl font-bold text-gray-900">Recent Bookings</h2>
          <a href="/influencer/referrals" class="text-accent hover:text-brand-dark transition-colors duration-200 flex items-center">
            View All
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 ml-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7" />
            </svg>
          </a>
        </div>

        <div v-if="recentBookings.length === 0" class="bg-white rounded-xl shadow-lg p-8 text-center">
          <p class="text-gray-500">No bookings yet. Share your referral code to get started!</p>
        </div>

        <div v-else class="bg-white rounded-xl shadow-lg overflow-hidden">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  User
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Product
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Date
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Status
                </th>
                <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                  Cashback
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="booking in recentBookings" :key="booking.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="flex items-center">
                    <div class="h-10 w-10 rounded-full bg-gray-200 flex items-center justify-center text-gray-500">
                      {{ booking.user?.name?.charAt(0) || '?' }}
                    </div>
                    <div class="ml-4">
                      <div class="text-sm font-medium text-gray-900">{{ booking.user?.name || 'Unknown' }}</div>
                      <div class="text-sm text-gray-500">{{ booking.user?.email || '' }}</div>
                    </div>
                  </div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">{{ booking.product?.name || 'Unknown Product' }}</div>
                  <div class="text-sm text-gray-500">${{ booking.product?.price || 0 }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <div class="text-sm text-gray-900">{{ formatDate(booking.date) }}</div>
                </td>
                <td class="px-6 py-4 whitespace-nowrap">
                  <span 
                    class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
                    :class="{
                      'bg-green-100 text-green-800': booking.status === 'Approved' || booking.status === 'approved',
                      'bg-yellow-100 text-yellow-800': booking.status === 'Pending' || booking.status === 'pending',
                      'bg-red-100 text-red-800': booking.status === 'Rejected' || booking.status === 'rejected'
                    }"
                  >
                    {{ booking.status }}
                  </span>
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  <span 
                    :class="{
                      'text-green-600': booking.cashbackStatus === 'Paid' || booking.cashbackStatus === 'paid',
                      'text-gray-600': booking.cashbackStatus !== 'Paid' && booking.cashbackStatus !== 'paid'
                    }"
                  >
                    {{ booking.cashbackStatus || 'Not Paid' }}
                  </span>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from 'vue';
import { useAuthStore } from '../../stores/auth';
import referralService from '../../services/referral.service';
import api from '../../services/api';

// Auth store
const authStore = useAuthStore();

// State
const isLoading = ref(true);
const error = ref<string | null>(null);
const referralCode = ref<string>('');
const isCopied = ref(false);
const userName = computed(() => authStore.user?.fullName || authStore.user?.email?.split('@')[0]);

// Stats data
const stats = ref({
  totalReferrals: 0,
  newReferrals: 0,
  approvedBookings: 0,
  approvalRate: 0,
  pendingBookings: 0,
  pendingAmount: '$0'
});

// Products and bookings
const latestProducts = ref<any[]>([]);
const recentBookings = ref<any[]>([]);

// Format date for display
const formatDate = (date: Date | string) => {
  const dateObj = typeof date === 'string' ? new Date(date) : date;
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(dateObj);
};

// Copy referral code function
const copyReferralCode = () => {
  if (!referralCode.value) return;
  
  navigator.clipboard.writeText(referralCode.value);
  isCopied.value = true;
  setTimeout(() => {
    isCopied.value = false;
  }, 2000);
};

// Share referral code function
const shareReferralCode = async () => {
  if (!referralCode.value) return;
  
  const shareData = {
    title: 'Ecomflex Referral',
    text: `Use my referral code ${referralCode.value} to get 100% cashback on your product purchases!`,
    url: `${window.location.origin}/referral/${referralCode.value}`
  };

  try {
    if (navigator.share) {
      await navigator.share(shareData);
    } else {
      // Fallback if Web Share API is not available
      copyReferralCode();
    }
  } catch (err) {
    console.error('Error sharing:', err);
    copyReferralCode();
  }
};

// Share product function
const shareProduct = async (product: any) => {
  const shareData = {
    title: `Check out ${product.name}`,
    text: `Get ${product.name} with 100% cashback using my referral code ${referralCode.value}!`,
    url: `${window.location.origin}/products/${product.id}?ref=${referralCode.value}`
  };

  try {
    if (navigator.share) {
      await navigator.share(shareData);
    } else {
      const productUrl = `${window.location.origin}/products/${product.id}?ref=${referralCode.value}`;
      await navigator.clipboard.writeText(productUrl);
      alert('Product link copied to clipboard!');
    }
  } catch (err) {
    console.error('Error sharing product:', err);
  }
};

// Load all dashboard data
const loadDashboardData = async () => {
  isLoading.value = true;
  error.value = null;

  try {
    console.log('🔄 Loading dashboard data...');
    console.log('📋 Current user:', authStore.user);
    console.log('🔑 Access token exists:', !!authStore.accessToken);

    // Fetch referral code and user profile
    const [codeResponse, statsResponse, productsResponse, referralsResponse] = await Promise.allSettled([
      referralService.getReferralCode(),
      referralService.getReferralStats(),
      api.products.getAllProducts({ limit: 6 }),
      referralService.getReferrals()
    ]);

    // Handle referral code
    if (codeResponse.status === 'fulfilled') {
      console.log('✅ Referral code response:', codeResponse.value);
      // Handle different response formats
      if (codeResponse.value.data?.referral_code) {
        referralCode.value = codeResponse.value.data.referral_code;
      } else if (codeResponse.value.referral_code) {
        referralCode.value = codeResponse.value.referral_code;
      } else if (codeResponse.value.code) {
        referralCode.value = codeResponse.value.code;
      } else {
        // Fallback: use user info from auth store
        referralCode.value = authStore.user?.email?.split('@')[0]?.toUpperCase() || 'LOADING';
      }
    } else {
      console.error('❌ Failed to load referral code:', codeResponse.reason);
      referralCode.value = authStore.user?.email?.split('@')[0]?.toUpperCase() || 'ERROR';
    }

    // Handle stats
    if (statsResponse.status === 'fulfilled') {
      console.log('✅ Stats response:', statsResponse.value);
      const statsData = statsResponse.value.data || statsResponse.value;
      stats.value = {
        totalReferrals: statsData.total_referrals || statsData.totalReferrals || 0,
        newReferrals: statsData.new_referrals || statsData.newReferrals || 0,
        approvedBookings: statsData.approved_bookings || statsData.approvedBookings || 0,
        approvalRate: statsData.approval_rate || statsData.approvalRate || 0,
        pendingBookings: statsData.pending_bookings || statsData.pendingBookings || 0,
        pendingAmount: statsData.pending_amount || statsData.pendingAmount || '$0'
      };
    } else {
      console.error('❌ Failed to load stats:', statsResponse.reason);
    }

    // Handle products
    if (productsResponse.status === 'fulfilled') {
      console.log('✅ Products response:', productsResponse.value);
      const productsData = productsResponse.value.data?.data || productsResponse.value.data?.products || productsResponse.value.data || [];
      latestProducts.value = productsData.slice(0, 6).map((product: any) => ({
        id: product.id,
        name: product.name || product.title,
        description: product.description,
        price: product.price || product.amount,
        image: product.image || product.imageUrl || product.image_url,
        referrals: product.referral_count || product.referrals || 0
      }));
    } else {
      console.error('❌ Failed to load products:', productsResponse.reason);
    }

    // Handle referrals/bookings
    if (referralsResponse.status === 'fulfilled') {
      console.log('✅ Referrals response:', referralsResponse.value);
      const referralsData = referralsResponse.value.data?.data || referralsResponse.value.data?.referrals || referralsResponse.value.data || [];
      recentBookings.value = referralsData.slice(0, 5).map((referral: any) => ({
        id: referral.id,
        user: {
          name: referral.user_name || referral.userName || referral.customer_name || 'Unknown User',
          email: referral.user_email || referral.userEmail || referral.customer_email || ''
        },
        product: {
          name: referral.product_name || referral.productName || 'Unknown Product',
          price: referral.product_price || referral.productPrice || 0
        },
        date: referral.created_at || referral.createdAt || referral.booking_date || new Date(),
        status: referral.status || 'pending',
        cashbackStatus: referral.cashback_status || referral.cashbackStatus || 'Not Paid'
      }));
    } else {
      console.error('❌ Failed to load referrals:', referralsResponse.reason);
    }

    console.log('✅ Dashboard data loaded successfully');
  } catch (err: any) {
    console.error('❌ Error loading dashboard data:', err);
    error.value = err.response?.data?.message || err.message || 'Failed to load dashboard data. Please try again.';
  } finally {
    isLoading.value = false;
  }
};

// Initialize on mount
onMounted(() => {
  console.log('🚀 Dashboard component mounted');
  console.log('👤 Logged in user:', authStore.user);
  console.log('🔐 Is authenticated:', authStore.isAuthenticated);
  
  loadDashboardData();
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

.bg-accent\/10 {
  background-color: rgba(255, 215, 0, 0.1);
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

/* Line clamp for product descriptions */
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;  /* ← ADD THIS LINE */
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style>