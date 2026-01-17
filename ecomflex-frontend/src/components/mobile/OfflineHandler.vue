<template>
  <div class="offline-handler">
    <div v-if="isOffline" class="offline-banner" :class="{ 'expanded': showDetails }">
      <div class="offline-header" @click="toggleDetails">
        <div class="offline-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="1" y1="1" x2="23" y2="23"></line>
            <path d="M16.72 11.06A10.94 10.94 0 0 1 19 12.55"></path>
            <path d="M5 12.55a10.94 10.94 0 0 1 5.17-2.39"></path>
            <path d="M10.71 5.05A16 16 0 0 1 22.58 9"></path>
            <path d="M1.42 9a15.91 15.91 0 0 1 4.7-2.88"></path>
            <path d="M8.53 16.11a6 6 0 0 1 6.95 0"></path>
            <line x1="12" y1="20" x2="12.01" y2="20"></line>
          </svg>
        </div>
        <div class="offline-title">You're offline</div>
        <div class="offline-chevron" :class="{ 'rotated': showDetails }">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <polyline points="6 9 12 15 18 9"></polyline>
          </svg>
        </div>
      </div>
      
      <div v-if="showDetails" class="offline-details">
        <p>While offline, you can still:</p>
        <ul>
          <li>View previously loaded products</li>
          <li>Check your cached booking history</li>
          <li>Browse saved referral codes</li>
        </ul>
        
        <p class="offline-limitations">Some features like uploading proofs, submitting forms, or viewing new products require an internet connection.</p>
        
        <div class="offline-actions">
          <button class="retry-button" @click="checkConnection">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M21 2v6h-6"></path>
              <path d="M3 12a9 9 0 0 1 15-6.7L21 8"></path>
              <path d="M3 12a9 9 0 0 0 15 6.7l3-2.7"></path>
            </svg>
            Try again
          </button>
        </div>
      </div>
    </div>
    
    <div v-if="isOffline && currentRoute === '/'" class="offline-homepage">
      <div class="offline-content">
        <img src="/images/offline-illustration.svg" alt="Offline" class="offline-illustration" />
        <h2>No internet connection</h2>
        <p>Please check your connection and try again</p>
        <button class="retry-button large" @click="checkConnection">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M21 2v6h-6"></path>
            <path d="M3 12a9 9 0 0 1 15-6.7L21 8"></path>
            <path d="M3 12a9 9 0 0 0 15 6.7l3-2.7"></path>
          </svg>
          Retry
        </button>
        
        <div class="offline-cached-sections">
          <h3>Available offline</h3>
          
          <div v-if="cachedBookings.length > 0" class="cached-section">
            <h4>My Bookings</h4>
            <router-link to="/my-bookings" class="view-link">
              View
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="9 18 15 12 9 6"></polyline>
              </svg>
            </router-link>
          </div>
          
          <div v-if="cachedProducts.length > 0" class="cached-section">
            <h4>Cached Products</h4>
            <router-link to="/home" class="view-link">
              View
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="9 18 15 12 9 6"></polyline>
              </svg>
            </router-link>
          </div>
          
          <div v-if="userRole === 'influencer' && cachedReferrals.length > 0" class="cached-section">
            <h4>My Referrals</h4>
            <router-link to="/influencer/referrals" class="view-link">
              View
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="9 18 15 12 9 6"></polyline>
              </svg>
            </router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, onBeforeUnmount, computed } from 'vue';
import { useRoute } from 'vue-router';

export default defineComponent({
  name: 'OfflineHandler',
  props: {
    userRole: {
      type: String,
      default: 'public'
    }
  },
  setup(props) {
    const route = useRoute();
    const isOffline = ref(!navigator.onLine);
    const showDetails = ref(false);
    
    // Mock cached data for demo purposes
    // In a real app, this would come from IndexedDB or localStorage
    const cachedBookings = ref([
      { id: 1, productName: 'Wireless Earbuds', status: 'Approved', date: '2023-08-15' },
      { id: 2, productName: 'Smart Watch', status: 'Pending', date: '2023-09-02' }
    ]);
    
    const cachedProducts = ref([
      { id: 1, name: 'Bluetooth Speaker', asin: 'B07X123YZ', price: '$49.99' },
      { id: 2, name: 'Phone Holder', asin: 'B08Z456AB', price: '$19.99' },
      { id: 3, name: 'Wireless Charger', asin: 'B09C789DE', price: '$29.99' }
    ]);
    
    const cachedReferrals = ref([
      { id: 1, code: 'INFL123', uses: 5, conversions: 3 },
      { id: 2, code: 'SUMMER21', uses: 12, conversions: 8 }
    ]);
    
    const currentRoute = computed(() => route.path);
    
    // Toggle details section
    const toggleDetails = () => {
      showDetails.value = !showDetails.value;
    };
    
    // Check connection
    const checkConnection = () => {
      if (navigator.onLine) {
        isOffline.value = false;
        window.location.reload();
      } else {
        // Show toast or feedback that still offline
        const offlineBanner = document.querySelector('.offline-banner');
        if (offlineBanner) {
          offlineBanner.classList.add('shake');
          setTimeout(() => {
            offlineBanner.classList.remove('shake');
          }, 500);
        }
      }
    };
    
    // Handle online/offline events
    const handleOnline = () => {
      isOffline.value = false;
    };
    
    const handleOffline = () => {
      isOffline.value = true;
      showDetails.value = true; // Auto-expand when going offline
    };
    
    // Custom event handlers for PWA
    const handleAppOnline = () => {
      isOffline.value = false;
    };
    
    const handleAppOffline = () => {
      isOffline.value = true;
    };
    
    onMounted(() => {
      // Check initial status
      isOffline.value = !navigator.onLine;
      
      // Add event listeners
      window.addEventListener('online', handleOnline);
      window.addEventListener('offline', handleOffline);
      document.addEventListener('app:online', handleAppOnline);
      document.addEventListener('app:offline', handleAppOffline);
    });
    
    onBeforeUnmount(() => {
      // Remove event listeners
      window.removeEventListener('online', handleOnline);
      window.removeEventListener('offline', handleOffline);
      document.removeEventListener('app:online', handleAppOnline);
      document.removeEventListener('app:offline', handleAppOffline);
    });
    
    return {
      isOffline,
      showDetails,
      cachedBookings,
      cachedProducts,
      cachedReferrals,
      currentRoute,
      toggleDetails,
      checkConnection,
      userRole: props.userRole
    };
  }
});
</script>

<style scoped>
.offline-handler {
  width: 100%;
}

.offline-banner {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  background-color: #ffcc00;
  color: #333;
  z-index: 100;
  transition: all 0.3s ease;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.offline-header {
  display: flex;
  align-items: center;
  padding: 12px 16px;
  cursor: pointer;
}

.offline-icon {
  margin-right: 12px;
}

.offline-title {
  flex: 1;
  font-weight: 500;
}

.offline-chevron {
  transition: transform 0.3s ease;
}

.offline-chevron.rotated {
  transform: rotate(180deg);
}

.offline-details {
  padding: 0 16px 16px;
  border-top: 1px solid rgba(0, 0, 0, 0.1);
  animation: slideDown 0.3s ease;
}

.offline-details p {
  margin: 12px 0 8px;
  font-weight: 500;
}

.offline-details ul {
  list-style-type: none;
  padding: 0;
  margin: 0 0 16px;
}

.offline-details li {
  padding: 6px 0 6px 24px;
  position: relative;
}

.offline-details li:before {
  content: "✓";
  position: absolute;
  left: 0;
  color: #007BFF;
}

.offline-limitations {
  font-size: 13px;
  color: #666;
  margin-top: 16px;
}

.offline-actions {
  display: flex;
  justify-content: flex-end;
  margin-top: 16px;
}

.retry-button {
  display: flex;
  align-items: center;
  gap: 8px;
  background: #007BFF;
  color: white;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.retry-button:hover {
  background: #0056b3;
  transform: translateY(-1px);
}

.retry-button.large {
  padding: 12px 24px;
  font-size: 16px;
}

/* Full offline homepage */
.offline-homepage {
  padding: 24px 16px;
  min-height: calc(100vh - 64px); /* Account for bottom nav */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  background-color: var(--bg-color, #F0F8FF);
}

.offline-content {
  max-width: 400px;
  margin: 0 auto;
}

.offline-illustration {
  width: 180px;
  height: auto;
  margin-bottom: 24px;
}

.offline-homepage h2 {
  font-size: 24px;
  margin-bottom: 12px;
  color: #333;
}

.offline-homepage p {
  font-size: 16px;
  margin-bottom: 24px;
  color: #666;
}

.offline-cached-sections {
  margin-top: 40px;
  width: 100%;
  text-align: left;
}

.offline-cached-sections h3 {
  font-size: 18px;
  margin-bottom: 16px;
  color: #333;
  border-bottom: 1px solid rgba(0, 0, 0, 0.1);
  padding-bottom: 8px;
}

.cached-section {
  background: white;
  border-radius: 12px;
  padding: 16px;
  margin-bottom: 16px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.cached-section h4 {
  font-size: 16px;
  margin: 0;
  color: #333;
}

.view-link {
  display: flex;
  align-items: center;
  color: #007BFF;
  text-decoration: none;
  font-weight: 500;
  gap: 4px;
}

/* Animations */
@keyframes slideDown {
  from { opacity: 0; transform: translateY(-10px); }
  to { opacity: 1; transform: translateY(0); }
}

.shake {
  animation: shake 0.5s cubic-bezier(.36,.07,.19,.97) both;
}

@keyframes shake {
  0%, 100% { transform: translateX(0); }
  10%, 30%, 50%, 70%, 90% { transform: translateX(-5px); }
  20%, 40%, 60%, 80% { transform: translateX(5px); }
}
</style>