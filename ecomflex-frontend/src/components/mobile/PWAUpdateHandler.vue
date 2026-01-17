<!-- src/components/mobile/PWAUpdateHandler.vue -->
<template>
  <div class="pwa-update-handler">
    <!-- Update Available Banner -->
    <transition name="slide-down">
      <div v-if="showUpdateBanner" class="fixed top-0 left-0 right-0 bg-white shadow-lg z-50">
        <div class="max-w-7xl mx-auto px-4 py-3 flex items-center justify-between">
          <div class="flex items-center">
            <div class="text-brand-600 mr-3">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 animate-spin-slow" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
            </div>
            <div>
              <p class="font-medium text-gray-900">A new version is available!</p>
              <p class="text-sm text-gray-500">Update now for the latest features and improvements</p>
            </div>
          </div>
          <div class="flex space-x-2">
            <button 
              @click="updateNow" 
              class="bg-brand-600 text-white px-4 py-2 rounded-md text-sm font-medium hover:bg-brand-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500"
            >
              Update Now
            </button>
            <button 
              @click="dismissUpdate" 
              class="bg-white text-gray-700 px-4 py-2 rounded-md text-sm font-medium hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-brand-500 border border-gray-300"
            >
              Later
            </button>
          </div>
        </div>
      </div>
    </transition>
    
    <!-- Installing Update Overlay -->
    <div v-if="isUpdating" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white p-6 rounded-lg shadow-xl max-w-md w-full">
        <div class="flex flex-col items-center">
          <div class="text-brand-600 mb-4">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-12 w-12 animate-spin" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-gray-900 mb-1">Updating Ecomflex</h3>
          <p class="text-gray-600 mb-4">This will only take a moment...</p>
          <div class="w-full bg-gray-200 rounded-full h-2.5">
            <div class="bg-brand-600 h-2.5 rounded-full" :style="{ width: `${updateProgress}%` }"></div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- PWA Install Prompt -->
    <transition name="slide-up">
      <div v-if="showInstallPrompt" class="fixed bottom-0 inset-x-0 pb-2 sm:pb-5 z-40">
        <div class="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
          <div class="p-2 rounded-lg bg-brand-600 shadow-lg sm:p-3">
            <div class="flex items-center justify-between flex-wrap">
              <div class="flex-1 flex items-center">
                <span class="flex p-2 rounded-lg bg-brand-800">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18h.01M8 21h8a2 2 0 002-2V5a2 2 0 00-2-2H8a2 2 0 00-2 2v14a2 2 0 002 2z" />
                  </svg>
                </span>
                <p class="ml-3 font-medium text-white truncate">
                  <span class="md:hidden">Install Ecomflex app for better experience!</span>
                  <span class="hidden md:inline">Install Ecomflex on your device for a faster and better experience!</span>
                </p>
              </div>
              <div class="order-3 mt-2 flex-shrink-0 w-full sm:order-2 sm:mt-0 sm:w-auto">
                <button 
                  @click="installPWA" 
                  class="flex items-center justify-center px-4 py-2 border border-transparent rounded-md shadow-sm text-sm font-medium text-brand-600 bg-white hover:bg-brand-50"
                >
                  Install Now
                </button>
              </div>
              <div class="order-2 flex-shrink-0 sm:order-3 sm:ml-2">
                <button 
                  @click="dismissInstallPrompt" 
                  type="button" 
                  class="-mr-1 flex p-2 rounded-md hover:bg-brand-500 focus:outline-none focus:ring-2 focus:ring-white"
                >
                  <span class="sr-only">Dismiss</span>
                  <svg class="h-6 w-6 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </transition>
    
    <!-- Offline Status Banner -->
    <transition name="slide-down">
      <div v-if="isOffline" class="fixed top-0 left-0 right-0 bg-amber-500 text-white z-50">
        <div class="max-w-7xl mx-auto px-4 py-2 flex items-center">
          <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <p>You're currently offline. Some features may be limited.</p>
        </div>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, onBeforeUnmount } from 'vue';

export default defineComponent({
  name: 'PWAUpdateHandler',
  setup() {
    const showUpdateBanner = ref(false);
    const showInstallPrompt = ref(false);
    const isUpdating = ref(false);
    const isOffline = ref(false);
    const updateProgress = ref(0);
    let deferredPrompt: any = null;
    let progressInterval: number | null = null;
    
    // Check if the app needs an update
    const checkForUpdates = () => {
      if ('serviceWorker' in navigator) {
        navigator.serviceWorker.ready.then((registration) => {
          registration.addEventListener('updatefound', () => {
            if (registration.installing) {
              const newWorker = registration.installing;
              
              newWorker.addEventListener('statechange', () => {
                if (newWorker.state === 'installed' && navigator.serviceWorker.controller) {
                  // New content is available, show update banner
                  showUpdateBanner.value = true;
                }
              });
            }
          });
          
          // Check for updates every hour
          setInterval(() => {
            registration.update().catch(console.error);
          }, 60 * 60 * 1000);
        });
      }
    };
    
    // Update the application now
    const updateNow = () => {
      if ('serviceWorker' in navigator) {
        isUpdating.value = true;
        
        // Simulate progress for better UX
        updateProgress.value = 0;
        progressInterval = window.setInterval(() => {
          updateProgress.value += Math.random() * 10;
          if (updateProgress.value >= 100) {
            updateProgress.value = 100;
            clearInterval(progressInterval as number);
            
            // Force page reload to activate the new service worker
            window.location.reload();
          }
        }, 300);
        
        navigator.serviceWorker.ready.then((registration) => {
          if (registration.waiting) {
            // Send message to service worker to skip waiting
            registration.waiting.postMessage({ type: 'SKIP_WAITING' });
          }
        });
      }
    };
    
    // Dismiss the update banner
    const dismissUpdate = () => {
      showUpdateBanner.value = false;
      
      // Save the dismissal time to avoid showing again too soon
      localStorage.setItem('updateDismissedAt', Date.now().toString());
    };
    
    // Install the PWA
    const installPWA = () => {
      if (deferredPrompt) {
        // Show the installation prompt
        deferredPrompt.prompt();
        
        deferredPrompt.userChoice.then((choiceResult: { outcome: string }) => {
          if (choiceResult.outcome === 'accepted') {
            console.log('User accepted the install prompt');
          } else {
            console.log('User dismissed the install prompt');
          }
          
          // Clear the deferred prompt
          deferredPrompt = null;
          showInstallPrompt.value = false;
        });
      }
    };
    
    // Dismiss the install prompt
    const dismissInstallPrompt = () => {
      showInstallPrompt.value = false;
      
      // Save the dismissal time to localStorage
      localStorage.setItem('installPromptDismissedAt', Date.now().toString());
    };
    
    // Check if the user is offline
    const handleOnlineStatusChange = () => {
      isOffline.value = !navigator.onLine;
    };
    
    onMounted(() => {
      checkForUpdates();
      
      // Check if we should show the update banner based on last dismissal
      const lastDismissed = localStorage.getItem('updateDismissedAt');
      if (lastDismissed) {
        const dismissedTime = parseInt(lastDismissed);
        const now = Date.now();
        
        // Only show if the last dismissal was more than 24 hours ago
        if (now - dismissedTime > 24 * 60 * 60 * 1000) {
          showUpdateBanner.value = false; // Initially false, will be set to true if an update is found
        }
      }
      
      // Listen for the beforeinstallprompt event
      window.addEventListener('beforeinstallprompt', (e) => {
        // Prevent Chrome 67 and earlier from automatically showing the prompt
        e.preventDefault();
        
        // Stash the event so it can be triggered later
        deferredPrompt = e;
        
        // Check if we should show the install prompt based on last dismissal
        const lastDismissed = localStorage.getItem('installPromptDismissedAt');
        if (lastDismissed) {
          const dismissedTime = parseInt(lastDismissed);
          const now = Date.now();
          
          // Only show if the last dismissal was more than 7 days ago
          if (now - dismissedTime > 7 * 24 * 60 * 60 * 1000) {
            showInstallPrompt.value = true;
          }
        } else {
          // Show the install prompt if it's never been dismissed
          showInstallPrompt.value = true;
        }
      });
      
      // Listen for the appinstalled event
      window.addEventListener('appinstalled', () => {
        // Hide the install prompt
        showInstallPrompt.value = false;
        deferredPrompt = null;
        
        // Log install to analytics
        console.log('PWA was installed');
      });
      
      // Listen for online/offline events
      window.addEventListener('online', handleOnlineStatusChange);
      window.addEventListener('offline', handleOnlineStatusChange);
      
      // Initial online status check
      isOffline.value = !navigator.onLine;
      
      // Listen for service worker controlling the page
      // DISABLED - This was causing reload loops
      // navigator.serviceWorker.addEventListener('controllerchange', () => {
      //   if (!isUpdating.value) {
      //     // This is called when the service worker controlling this page changes,
      //     // for example after a new service worker has skipped waiting and become active.
      //     window.location.reload();
      //   }
      // });
    });
    
    onBeforeUnmount(() => {
      // Clean up event listeners
      window.removeEventListener('online', handleOnlineStatusChange);
      window.removeEventListener('offline', handleOnlineStatusChange);
      
      // Clear any intervals
      if (progressInterval !== null) {
        clearInterval(progressInterval);
      }
    });
    
    return {
      showUpdateBanner,
      showInstallPrompt,
      isUpdating,
      isOffline,
      updateProgress,
      updateNow,
      dismissUpdate,
      installPWA,
      dismissInstallPrompt
    };
  }
});
</script>

<style scoped>
/* Slide Down Animation */
.slide-down-enter-active, .slide-down-leave-active {
  transition: transform 0.3s ease, opacity 0.3s ease;
}
.slide-down-enter-from, .slide-down-leave-to {
  transform: translateY(-100%);
  opacity: 0;
}

/* Slide Up Animation */
.slide-up-enter-active, .slide-up-leave-active {
  transition: transform 0.3s ease, opacity 0.3s ease;
}
.slide-up-enter-from, .slide-up-leave-to {
  transform: translateY(100%);
  opacity: 0;
}

/* Slow Spin Animation */
@keyframes spin-slow {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
.animate-spin-slow {
  animation: spin-slow 2s linear infinite;
}
</style>