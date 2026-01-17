<template>
  <div class="min-h-screen flex flex-col">
    <!-- PWA Update Handler -->
    <PWAUpdateHandler />
    
    <!-- Page transitions -->
    <router-view v-slot="{ Component }">
      <transition
        name="page"
        mode="out-in"
        enter-active-class="transition-all duration-300 ease-out"
        enter-from-class="opacity-0 translate-y-4"
        enter-to-class="opacity-100 translate-y-0"
        leave-active-class="transition-all duration-200 ease-in"
        leave-from-class="opacity-100 translate-y-0"
        leave-to-class="opacity-0 translate-y-4"
      >
        <component :is="Component" />
      </transition>
    </router-view>
    
    <!-- Network Status Toast -->
    <transition name="fade">
      <div v-if="showOfflineToast" class="fixed bottom-4 right-4 bg-amber-500 text-white px-4 py-3 rounded-lg shadow-lg flex items-center">
        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
        </svg>
        <span>You are currently offline</span>
        <button @click="hideOfflineToast" class="ml-4 text-white">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
          </svg>
        </button>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import PWAUpdateHandler from './components/mobile/PWAUpdateHandler.vue';

// Network status handling
const showOfflineToast = ref(false);
let offlineToastTimeout: number | null = null;

// Hide the offline toast
const hideOfflineToast = () => {
  showOfflineToast.value = false;
  if (offlineToastTimeout) {
    clearTimeout(offlineToastTimeout);
    offlineToastTimeout = null;
  }
};

// Handle network status changes
const handleOnlineStatus = () => {
  if (!navigator.onLine) {
    showOfflineToast.value = true;
    
    // Auto-hide after 5 seconds
    if (offlineToastTimeout) {
      clearTimeout(offlineToastTimeout);
    }
    
    offlineToastTimeout = window.setTimeout(() => {
      showOfflineToast.value = false;
      offlineToastTimeout = null;
    }, 5000);
  } else {
    // We're back online
    if (showOfflineToast.value) {
      // If the offline toast is showing, hide it and show online toast
      hideOfflineToast();
    }
  }
};

// Custom event handlers for PWA events
const handlePWAOnline = () => {
  if (showOfflineToast.value) {
    hideOfflineToast();
  }
};

const handlePWAOffline = () => {
  showOfflineToast.value = true;
  
  // Auto-hide after 5 seconds
  if (offlineToastTimeout) {
    clearTimeout(offlineToastTimeout);
  }
  
  offlineToastTimeout = window.setTimeout(() => {
    showOfflineToast.value = false;
    offlineToastTimeout = null;
  }, 5000);
};

onMounted(() => {
  // Set initial status
  if (!navigator.onLine) {
    showOfflineToast.value = true;
    
    // Auto-hide after 5 seconds
    offlineToastTimeout = window.setTimeout(() => {
      showOfflineToast.value = false;
      offlineToastTimeout = null;
    }, 5000);
  }
  
  // Add event listeners
  window.addEventListener('online', handleOnlineStatus);
  window.addEventListener('offline', handleOnlineStatus);
  
  // Listen for custom PWA events
  document.addEventListener('pwa:online', handlePWAOnline);
  document.addEventListener('pwa:offline', handlePWAOffline);
});

onBeforeUnmount(() => {
  // Remove event listeners
  window.removeEventListener('online', handleOnlineStatus);
  window.removeEventListener('offline', handleOnlineStatus);
  document.removeEventListener('pwa:online', handlePWAOnline);
  document.removeEventListener('pwa:offline', handlePWAOffline);
  
  // Clear any timeouts
  if (offlineToastTimeout) {
    clearTimeout(offlineToastTimeout);
  }
});
</script>

<style>
/* Any global styles not covered by Tailwind can go here */
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.3s ease;
}
.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

/* Define PWA-specific animations */
@keyframes slide-in-right {
  from {
    transform: translateX(100%);
  }
  to {
    transform: translateX(0);
  }
}

@keyframes slide-in-up {
  from {
    transform: translateY(100%);
  }
  to {
    transform: translateY(0);
  }
}

/* Custom styles for PWA install button glow effect */
.install-btn-glow {
  box-shadow: 0 0 15px 5px rgba(255, 215, 0, 0.3);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% {
    box-shadow: 0 0 15px 5px rgba(255, 215, 0, 0.3);
  }
  50% {
    box-shadow: 0 0 25px 10px rgba(255, 215, 0, 0.5);
  }
  100% {
    box-shadow: 0 0 15px 5px rgba(255, 215, 0, 0.3);
  }
}
</style>