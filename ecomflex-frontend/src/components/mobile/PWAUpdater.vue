<template>
  <div class="pwa-updater">
    <!-- Update Available Toast -->
    <div 
      v-if="showUpdateToast" 
      class="update-toast"
      :class="{ 'slide-in': showUpdateToast }"
    >
      <div class="update-toast-content">
        <div class="update-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M21 2v6h-6"></path>
            <path d="M3 12a9 9 0 0 1 15-6.7L21 8"></path>
            <path d="M3 12a9 9 0 0 0 15 6.7l3-2.7"></path>
          </svg>
        </div>
        <div class="update-message">
          <h3>Update Available</h3>
          <p>A new version of Ecomflex is available</p>
        </div>
        <div class="update-actions">
          <button class="update-now-btn" @click="updateNow">
            Update Now
          </button>
          <button class="dismiss-btn" @click="dismissUpdate">
            Later
          </button>
        </div>
      </div>
    </div>
    
    <!-- Install PWA Prompt -->
    <div 
      v-if="showInstallPrompt" 
      class="install-prompt"
      :class="{ 'slide-in': showInstallPrompt }"
    >
      <div class="install-prompt-content">
        <div class="install-icon">
          <img src="/icons/icon-96x96.png" alt="Ecomflex logo" />
        </div>
        <div class="install-message">
          <h3>Add to Home Screen</h3>
          <p>Install Ecomflex for a better experience</p>
        </div>
        <div class="install-actions">
          <button class="install-btn" @click="installPWA">
            Install
          </button>
          <button class="dismiss-btn" @click="dismissInstall">
            Not Now
          </button>
        </div>
      </div>
    </div>
    
    <!-- Refreshing Overlay (shown during update installation) -->
    <div v-if="isRefreshing" class="refreshing-overlay">
      <div class="refreshing-content">
        <div class="refreshing-spinner">
          <svg class="circular" viewBox="25 25 50 50">
            <circle class="path" cx="50" cy="50" r="20" fill="none" stroke-width="4" stroke-miterlimit="10"/>
          </svg>
        </div>
        <h3>Updating Ecomflex</h3>
        <p>This will only take a moment...</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, onBeforeUnmount } from 'vue';

// Define custom window properties
declare global {
  interface Window {
    updateApplication?: () => void;
    showInstallPrompt?: () => void;
  }
}

export default defineComponent({
  name: 'PWAUpdater',
  setup() {
    const showUpdateToast = ref(false);
    const showInstallPrompt = ref(false);
    const isRefreshing = ref(false);
    
    // Handler for app update available
    const handleUpdateAvailable = () => {
      showUpdateToast.value = true;
    };
    
    // Handler for app installable
    const handleInstallable = () => {
      // Check if we've previously dismissed the prompt
      const dismissed = localStorage.getItem('installPromptDismissed');
      const dismissedTime = dismissed ? parseInt(dismissed) : 0;
      
      // Only show the prompt if it hasn't been dismissed in the last 7 days
      if (!dismissed || (Date.now() - dismissedTime > 7 * 24 * 60 * 60 * 1000)) {
        showInstallPrompt.value = true;
      }
    };
    
    // Update the app now
    const updateNow = () => {
      if (typeof window.updateApplication === 'function') {
        isRefreshing.value = true;
        
        // Small delay to show the loading screen before actually updating
        setTimeout(() => {
          window.updateApplication?.();
        }, 1000);
      }
    };
    
    // Dismiss the update toast
    const dismissUpdate = () => {
      showUpdateToast.value = false;
    };
    
    // Install the PWA
    const installPWA = () => {
      if (typeof window.showInstallPrompt === 'function') {
        window.showInstallPrompt();
      }
      showInstallPrompt.value = false;
    };
    
    // Dismiss the install prompt
    const dismissInstall = () => {
      showInstallPrompt.value = false;
      // Remember that we dismissed the prompt
      localStorage.setItem('installPromptDismissed', Date.now().toString());
    };
    
    // Listen for app installed event
    const handleAppInstalled = () => {
      showInstallPrompt.value = false;
    };
    
    onMounted(() => {
      // Add event listeners
      document.addEventListener('app:updateAvailable', handleUpdateAvailable);
      document.addEventListener('app:installable', handleInstallable);
      document.addEventListener('app:installed', handleAppInstalled);
    });
    
    onBeforeUnmount(() => {
      // Remove event listeners
      document.removeEventListener('app:updateAvailable', handleUpdateAvailable);
      document.removeEventListener('app:installable', handleInstallable);
      document.removeEventListener('app:installed', handleAppInstalled);
    });
    
    return {
      showUpdateToast,
      showInstallPrompt,
      isRefreshing,
      updateNow,
      dismissUpdate,
      installPWA,
      dismissInstall
    };
  }
});
</script>

<style scoped>
.pwa-updater {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  z-index: 1000;
}

/* Update Toast */
.update-toast {
  position: fixed;
  top: 16px;
  left: 50%;
  transform: translateX(-50%) translateY(-100%);
  width: 90%;
  max-width: 400px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
  pointer-events: auto;
  transition: transform 0.3s ease-out;
  z-index: 1010;
}

.update-toast.slide-in {
  transform: translateX(-50%) translateY(0);
}

.update-toast-content {
  display: flex;
  padding: 16px;
  align-items: center;
}

.update-icon {
  color: #007BFF;
  margin-right: 16px;
}

.update-icon svg {
  animation: rotate 2s linear infinite;
}

.update-message {
  flex: 1;
}

.update-message h3 {
  font-size: 16px;
  margin: 0 0 4px 0;
  color: #333;
}

.update-message p {
  font-size: 14px;
  margin: 0;
  color: #666;
}

.update-actions {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-left: 16px;
}

.update-now-btn {
  background: #007BFF;
  color: white;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s ease;
}

.update-now-btn:hover {
  background: #0056b3;
}

.dismiss-btn {
  background: transparent;
  color: #666;
  border: none;
  padding: 4px 8px;
  font-size: 14px;
  cursor: pointer;
  transition: color 0.2s ease;
}

.dismiss-btn:hover {
  color: #333;
}

/* Install Prompt */
.install-prompt {
  position: fixed;
  bottom: 16px;
  left: 50%;
  transform: translateX(-50%) translateY(100%);
  width: 90%;
  max-width: 400px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
  pointer-events: auto;
  transition: transform 0.3s ease-out;
  z-index: 1010;
}

.install-prompt.slide-in {
  transform: translateX(-50%) translateY(0);
}

.install-prompt-content {
  display: flex;
  padding: 16px;
  align-items: center;
}

.install-icon {
  margin-right: 16px;
}

.install-icon img {
  width: 48px;
  height: 48px;
  border-radius: 8px;
}

.install-message {
  flex: 1;
}

.install-message h3 {
  font-size: 16px;
  margin: 0 0 4px 0;
  color: #333;
}

.install-message p {
  font-size: 14px;
  margin: 0;
  color: #666;
}

.install-actions {
  display: flex;
  flex-direction: column;
  gap: 8px;
  margin-left: 16px;
}

.install-btn {
  background: #007BFF;
  color: white;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s ease;
}

.install-btn:hover {
  background: #0056b3;
}

/* Refreshing Overlay */
.refreshing-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(255, 255, 255, 0.9);
  display: flex;
  justify-content: center;
  align-items: center;
  pointer-events: auto;
  z-index: 1020;
  animation: fadeIn 0.3s ease;
}

.refreshing-content {
  text-align: center;
  padding: 24px;
}

.refreshing-spinner {
  margin-bottom: 24px;
}

.refreshing-content h3 {
  font-size: 20px;
  margin: 0 0 8px 0;
  color: #333;
}

.refreshing-content p {
  font-size: 16px;
  margin: 0;
  color: #666;
}

/* SVG Spinner */
.circular {
  height: 50px;
  width: 50px;
  animation: rotate 2s linear infinite;
}

.path {
  stroke: #007BFF;
  stroke-linecap: round;
  animation: dash 1.5s ease-in-out infinite;
}

/* Animations */
@keyframes rotate {
  100% {
    transform: rotate(360deg);
  }
}

@keyframes dash {
  0% {
    stroke-dasharray: 1, 150;
    stroke-dashoffset: 0;
  }
  50% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -35;
  }
  100% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -124;
  }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}
</style>