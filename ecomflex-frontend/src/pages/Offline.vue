<template>
  <div class="offline-page">
    <div 
      class="max-w-3xl mx-auto text-center px-6 py-12"
      v-motion
      :initial="{ opacity: 0, y: 20 }"
      :enter="{ opacity: 1, y: 0, transition: { duration: 700 } }">
      
      <!-- Offline Icon -->
      <div class="flex justify-center mb-8">
        <div class="w-24 h-24 rounded-full bg-accent/10 flex items-center justify-center">
          <svg class="w-12 h-12 text-accent" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636a9 9 0 010 12.728m0 0l-2.829-2.829m2.829 2.829L21 21M15.536 8.464a5 5 0 010 7.072m0 0l-2.829-2.829m-4.243 2.829a4.978 4.978 0 01-1.414-2.83m-1.414 5.658a9 9 0 01-2.167-9.238m7.824 2.167a1 1 0 111.414 1.414m-1.414-1.414L3 3m8.293 8.293l1.414 1.414"></path>
          </svg>
        </div>
      </div>
      
      <!-- Title -->
      <h1 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">You're Offline</h1>
      
      <!-- Subtitle -->
      <p class="text-lg text-gray-600 mb-8">
        It looks like you've lost your internet connection. Don't worry, some features are still available offline.
      </p>
      
      <!-- Status Indicator -->
      <div 
        class="bg-gray-50 rounded-xl shadow-md p-4 mb-8 max-w-md mx-auto"
        v-motion
        :initial="{ opacity: 0 }"
        :enter="{ opacity: 1, transition: { delay: 300, duration: 500 } }">
        <div class="flex items-center mb-2">
          <div class="w-3 h-3 rounded-full bg-red-500 mr-2 pulse-animation"></div>
          <span class="text-sm font-medium text-gray-700">Network Status: Offline</span>
        </div>
        <p class="text-xs text-gray-500">
          We'll automatically reconnect when your internet is back.
        </p>
      </div>
      
      <!-- Available Features Card -->
      <div 
        class="bg-white rounded-xl shadow-lg p-6 mb-8"
        v-motion
        :initial="{ opacity: 0, scale: 0.95 }"
        :enter="{ opacity: 1, scale: 1, transition: { delay: 500, duration: 500 } }">
        <h3 class="text-lg font-medium text-gray-900 mb-4">Available While Offline</h3>
        
        <div class="grid md:grid-cols-2 gap-4 text-left">
          <div class="p-3 rounded-lg bg-gray-50">
            <div class="flex items-center mb-2">
              <div class="w-8 h-8 rounded-full bg-green-100 flex items-center justify-center mr-3">
                <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <span class="text-sm font-medium text-gray-800">View My Bookings</span>
            </div>
            <p class="text-xs text-gray-500 ml-11">
              You can view your previously loaded bookings
            </p>
          </div>
          
          <div class="p-3 rounded-lg bg-gray-50">
            <div class="flex items-center mb-2">
              <div class="w-8 h-8 rounded-full bg-green-100 flex items-center justify-center mr-3">
                <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <span class="text-sm font-medium text-gray-800">Your Referral Code</span>
            </div>
            <p class="text-xs text-gray-500 ml-11">
              Your referral code is available to share
            </p>
          </div>
          
          <div class="p-3 rounded-lg bg-gray-50">
            <div class="flex items-center mb-2">
              <div class="w-8 h-8 rounded-full bg-green-100 flex items-center justify-center mr-3">
                <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <span class="text-sm font-medium text-gray-800">Profile Information</span>
            </div>
            <p class="text-xs text-gray-500 ml-11">
              Your cached profile details are accessible
            </p>
          </div>
          
          <div class="p-3 rounded-lg bg-gray-50">
            <div class="flex items-center mb-2">
              <div class="w-8 h-8 rounded-full bg-green-100 flex items-center justify-center mr-3">
                <svg class="w-5 h-5 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"></path>
                </svg>
              </div>
              <span class="text-sm font-medium text-gray-800">Cached Products</span>
            </div>
            <p class="text-xs text-gray-500 ml-11">
              Previously loaded products are viewable
            </p>
          </div>
        </div>
      </div>
      
      <!-- Actions -->
      <div
        v-motion
        :initial="{ opacity: 0 }"
        :enter="{ opacity: 1, transition: { delay: 700, duration: 500 } }">
        <button 
          @click="checkConnection"
          class="px-6 py-3 bg-accent text-white rounded-lg hover:bg-accent/80 focus:outline-none focus:ring-4 focus:ring-accent/30 transition-all duration-300 shadow-md"
          :disabled="isChecking">
          <span v-if="isChecking" class="flex items-center">
            <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
              <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
              <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
            </svg>
            Checking Connection...
          </span>
          <span v-else>Try Again</span>
        </button>
        
        <div class="mt-4">
          <router-link 
            to="/" 
            class="text-sm text-accent hover:text-accent/80 transition-colors duration-300">
            Go to Home Page
          </router-link>
        </div>
      </div>
      
      <!-- Tip -->
      <div 
        class="mt-12 text-sm text-gray-500 max-w-md mx-auto"
        v-motion
        :initial="{ opacity: 0 }"
        :enter="{ opacity: 1, transition: { delay: 900, duration: 500 } }">
        <p>Tip: If you're experiencing connection issues, try these steps:</p>
        <ul class="text-left mt-2 space-y-1 ml-6 list-disc">
          <li>Check your Wi-Fi or mobile data connection</li>
          <li>Try switching between Wi-Fi and mobile data</li>
          <li>Restart your router or device</li>
          <li>Move to an area with better signal strength</li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { useMotion } from '@vueuse/motion';

const router = useRouter();
const isChecking = ref(false);
const connectionCheckInterval = ref(null);

const checkConnection = async () => {
  isChecking.value = true;
  
  try {
    // Try to fetch a small file to check for internet connectivity
    const response = await fetch('/api/ping', { 
      method: 'GET',
      cache: 'no-store',
      headers: {
        'Cache-Control': 'no-cache',
        'Pragma': 'no-cache'
      }
    });
    
    if (response.ok) {
      // Connection is back, redirect to home
      router.push('/');
    } else {
      // Still offline
      isChecking.value = false;
    }
  } catch (error) {
    // Network error, still offline
    isChecking.value = false;
  }
};

// Setup automatic connection check
const setupConnectionCheck = () => {
  // Check connection every 30 seconds
  connectionCheckInterval.value = setInterval(() => {
    // Only check if not already checking
    if (!isChecking.value) {
      checkConnection();
    }
  }, 30000);
  
  // Also listen for online event
  window.addEventListener('online', handleOnline);
};

// Handle online event
const handleOnline = () => {
  // Redirect to home when connection is restored
  router.push('/');
};

// Clean up
onUnmounted(() => {
  if (connectionCheckInterval.value) {
    clearInterval(connectionCheckInterval.value);
  }
  window.removeEventListener('online', handleOnline);
});

// Initialize
onMounted(() => {
  setupConnectionCheck();
});
</script>

<style scoped>
.offline-page {
  @apply min-h-screen bg-gray-50 flex items-center justify-center;
}

.pulse-animation {
  animation: pulse 2s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

@keyframes pulse {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0.5;
  }
}
</style>