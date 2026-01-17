```vue
<template>
  <div class="bottom-navigation" :class="{ 'hidden': !isMobile }">
    <div class="bottom-nav-container">
      <router-link 
        v-for="item in navItems" 
        :key="item.to"
        :to="item.to" 
        class="bottom-nav-item"
        :class="{ 'active': isActive(item.to) }"
        @click="handleClick(item)"
      >
        <div class="bottom-nav-icon">
          <component :is="item.icon" />
        </div>
        <span class="bottom-nav-label">{{ item.label }}</span>
        <div class="nav-indicator" v-if="isActive(item.to)"></div>
      </router-link>
    </div>
    
    <!-- Install PWA Prompt -->
    <div v-if="showInstallPrompt" class="install-prompt" ref="installPrompt">
      <div class="install-prompt-content">
        <div class="install-prompt-icon">
          <img src="/icons/icon-96x96.png" alt="Ecomflex logo" />
        </div>
        <div class="install-prompt-text">
          <h3>Add to Home Screen</h3>
          <p>Install Ecomflex for a better experience</p>
        </div>
        <button class="install-button" @click="installPWA">
          Install
        </button>
        <button class="dismiss-button" @click="dismissInstallPrompt">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>
    </div>
    
    <!-- Update Available Prompt -->
    <div v-if="updateAvailable" class="update-prompt" ref="updatePrompt">
      <div class="update-prompt-content">
        <div class="update-prompt-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M21 2v6h-6"></path>
            <path d="M3 12a9 9 0 0 1 15-6.7L21 8"></path>
            <path d="M3 12a9 9 0 0 0 15 6.7l3-2.7"></path>
          </svg>
        </div>
        <div class="update-prompt-text">
          <h3>Update Available</h3>
          <p>Refresh to get the latest version</p>
        </div>
        <button class="update-button" @click="updateApp">
          Update Now
        </button>
      </div>
    </div>
    
    <!-- Offline Alert -->
    <div v-if="isOffline" class="offline-alert" ref="offlineAlert">
      <div class="offline-alert-content">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <line x1="1" y1="1" x2="23" y2="23"></line>
          <path d="M16.72 11.06A10.94 10.94 0 0 1 19 12.55"></path>
          <path d="M5 12.55a10.94 10.94 0 0 1 5.17-2.39"></path>
          <path d="M10.71 5.05A16 16 0 0 1 22.58 9"></path>
          <path d="M1.42 9a15.91 15.91 0 0 1 4.7-2.88"></path>
          <path d="M8.53 16.11a6 6 0 0 1 6.95 0"></path>
          <line x1="12" y1="20" x2="12.01" y2="20"></line>
        </svg>
        <span>You're offline. Some features may be limited.</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, onMounted, onBeforeUnmount, h } from 'vue';
import { useRoute } from 'vue-router';

// Define custom window properties for TypeScript
declare global {
  interface Window {
    showInstallPrompt?: () => void;
    updateApplication?: () => void;
  }
}

// Define the navigation items interface
interface NavItem {
  to: string;
  label: string;
  icon: any;
  role?: string[];
}

export default defineComponent({
  name: 'BottomNavigation',
  props: {
    userRole: {
      type: String,
      default: 'public'
    }
  },
  setup(props) {
    const route = useRoute();
    const isMobile = ref(window.innerWidth < 768);
    const showInstallPrompt = ref(false);
    const updateAvailable = ref(false);
    const isOffline = ref(!navigator.onLine);
    
    // Create icons for the navigation
    const HomeIcon = () => h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      width: '24',
      height: '24',
      viewBox: '0 0 24 24',
      fill: 'none',
      stroke: 'currentColor',
      'stroke-width': '2',
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round'
    }, [
      h('path', { d: 'M3 9l9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z' }),
      h('polyline', { points: '9 22 9 12 15 12 15 22' })
    ]);
    
    const BookingsIcon = () => h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      width: '24',
      height: '24',
      viewBox: '0 0 24 24',
      fill: 'none',
      stroke: 'currentColor',
      'stroke-width': '2',
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round'
    }, [
      h('path', { d: 'M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2' }),
      h('rect', { x: '8', y: '2', width: '8', height: '4', rx: '1', ry: '1' }),
      h('path', { d: 'M9 14l2 2 4-4' })
    ]);
    
    const ReferralsIcon = () => h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      width: '24',
      height: '24',
      viewBox: '0 0 24 24',
      fill: 'none',
      stroke: 'currentColor',
      'stroke-width': '2',
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round'
    }, [
      h('path', { d: 'M17 21v-2a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v2' }),
      h('circle', { cx: '9', cy: '7', r: '4' }),
      h('path', { d: 'M23 21v-2a4 4 0 0 0-3-3.87' }),
      h('path', { d: 'M16 3.13a4 4 0 0 1 0 7.75' })
    ]);
    
    const ProfileIcon = () => h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      width: '24',
      height: '24',
      viewBox: '0 0 24 24',
      fill: 'none',
      stroke: 'currentColor',
      'stroke-width': '2',
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round'
    }, [
      h('path', { d: 'M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2' }),
      h('circle', { cx: '12', cy: '7', r: '4' })
    ]);
    
    const DashboardIcon = () => h('svg', {
      xmlns: 'http://www.w3.org/2000/svg',
      width: '24',
      height: '24',
      viewBox: '0 0 24 24',
      fill: 'none',
      stroke: 'currentColor',
      'stroke-width': '2',
      'stroke-linecap': 'round',
      'stroke-linejoin': 'round'
    }, [
      h('rect', { x: '3', y: '3', width: '7', height: '7' }),
      h('rect', { x: '14', y: '3', width: '7', height: '7' }),
      h('rect', { x: '14', y: '14', width: '7', height: '7' }),
      h('rect', { x: '3', y: '14', width: '7', height: '7' })
    ]);

    // Define navigation items based on user role
    const navItems = computed(() => {
      const items: NavItem[] = [
        { to: '/home', label: 'Home', icon: HomeIcon },
        { to: '/my-bookings', label: 'Bookings', icon: BookingsIcon },
      ];
      
      // Add role-specific navigation items
      if (props.userRole === 'influencer') {
        items.push({ to: '/influencer/referrals', label: 'Referrals', icon: ReferralsIcon });
      } else if (props.userRole === 'admin') {
        items.push({ to: '/admin/dashboard', label: 'Dashboard', icon: DashboardIcon });
      }
      
      // Add profile for all users
      items.push({ to: '/profile', label: 'Profile', icon: ProfileIcon });
      
      return items;
    });
    
    // Check if a route is active
    const isActive = (path: string) => {
      if (path === '/home' && route.path === '/') return true;
      return route.path.startsWith(path);
    };
    
    // Handle navigation item click
    const handleClick = (item: NavItem) => {
      // Add any click animations or tracking here
      console.log(`Navigating to ${item.to}`);
    };
    
    // Resize handler
    const handleResize = () => {
      isMobile.value = window.innerWidth < 768;
    };
    
    // PWA installation
    const installPWA = () => {
      if (window.showInstallPrompt) {
        window.showInstallPrompt();
        showInstallPrompt.value = false;
      }
    };
    
    const dismissInstallPrompt = () => {
      showInstallPrompt.value = false;
      // Save to localStorage to prevent showing again for some time
      localStorage.setItem('installPromptDismissed', Date.now().toString());
    };
    
    // Update the application
    const updateApp = () => {
      if (window.updateApplication) {
        window.updateApplication();
        updateAvailable.value = false;
      }
    };
    
    // Lifecycle hooks
    onMounted(() => {
      // Add event listeners
      window.addEventListener('resize', handleResize);
      
      // Listen for PWA installability
      document.addEventListener('app:installable', () => {
        // Check if the prompt was recently dismissed
        const dismissedTime = localStorage.getItem('installPromptDismissed');
        if (!dismissedTime || Date.now() - parseInt(dismissedTime) > 7 * 24 * 60 * 60 * 1000) {
          showInstallPrompt.value = true;
        }
      });
      
      // Listen for updates available
      document.addEventListener('app:updateAvailable', () => {
        updateAvailable.value = true;
      });
      
      // Listen for online/offline events
      document.addEventListener('app:online', () => {
        isOffline.value = false;
      });
      
      document.addEventListener('app:offline', () => {
        isOffline.value = true;
      });
      
      // Initial online status check
      isOffline.value = !navigator.onLine;
    });
    
    onBeforeUnmount(() => {
      // Remove event listeners
      window.removeEventListener('resize', handleResize);
      document.removeEventListener('app:installable', () => {});
      document.removeEventListener('app:updateAvailable', () => {});
      document.removeEventListener('app:online', () => {});
      document.removeEventListener('app:offline', () => {});
    });
    
    return {
      isMobile,
      navItems,
      isActive,
      handleClick,
      showInstallPrompt,
      updateAvailable,
      isOffline,
      installPWA,
      dismissInstallPrompt,
      updateApp
    };
  }
});
</script>

<style scoped>
.bottom-navigation {
  position: fixed;
  bottom: 0;
  left: 0;
  width: 100%;
  z-index: 50;
  transition: transform 0.3s ease-in-out;
}

.bottom-navigation.hidden {
  display: none;
}

.bottom-nav-container {
  display: flex;
  justify-content: space-around;
  align-items: center;
  background: white;
  border-top: 1px solid rgba(0, 123, 255, 0.1);
  height: 64px;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.05);
}

.bottom-nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  flex: 1;
  height: 100%;
  color: #6c757d;
  text-decoration: none;
  position: relative;
  transition: all 0.2s ease;
  padding: 8px 0;
}

.bottom-nav-item.active {
  color: #007BFF;
}

.bottom-nav-icon {
  margin-bottom: 4px;
  position: relative;
}

.bottom-nav-label {
  font-size: 12px;
  font-weight: 500;
}

.nav-indicator {
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  width: 40px;
  height: 3px;
  background-color: #007BFF;
  border-radius: 3px;
  transition: all 0.3s ease;
}

/* PWA Install Prompt */
.install-prompt {
  position: fixed;
  bottom: 76px;
  left: 50%;
  transform: translateX(-50%);
  width: 90%;
  max-width: 360px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 60;
  animation: slideUp 0.3s ease-out;
}

.install-prompt-content {
  display: flex;
  align-items: center;
  padding: 16px;
  position: relative;
}

.install-prompt-icon {
  margin-right: 12px;
}

.install-prompt-icon img {
  width: 48px;
  height: 48px;
  border-radius: 8px;
}

.install-prompt-text {
  flex: 1;
}

.install-prompt-text h3 {
  font-size: 16px;
  margin: 0 0 4px 0;
  color: #333;
}

.install-prompt-text p {
  font-size: 13px;
  margin: 0;
  color: #666;
}

.install-button {
  background: #007BFF;
  color: white;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  margin-left: 12px;
  transition: background 0.2s ease;
}

.install-button:hover {
  background: #0056b3;
}

.dismiss-button {
  position: absolute;
  top: 8px;
  right: 8px;
  background: transparent;
  border: none;
  color: #999;
  cursor: pointer;
  padding: 4px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Update Prompt */
.update-prompt {
  position: fixed;
  top: 16px;
  left: 50%;
  transform: translateX(-50%);
  width: 90%;
  max-width: 360px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 60;
  animation: slideDown 0.3s ease-out;
}

.update-prompt-content {
  display: flex;
  align-items: center;
  padding: 16px;
}

.update-prompt-icon {
  margin-right: 12px;
  color: #007BFF;
  animation: rotate 2s linear infinite;
}

.update-prompt-text {
  flex: 1;
}

.update-prompt-text h3 {
  font-size: 16px;
  margin: 0 0 4px 0;
  color: #333;
}

.update-prompt-text p {
  font-size: 13px;
  margin: 0;
  color: #666;
}

.update-button {
  background: #007BFF;
  color: white;
  border: none;
  border-radius: 20px;
  padding: 8px 16px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  margin-left: 12px;
  transition: background 0.2s ease;
}

.update-button:hover {
  background: #0056b3;
}

/* Offline Alert */
.offline-alert {
  position: fixed;
  top: 16px;
  left: 50%;
  transform: translateX(-50%);
  width: 90%;
  max-width: 360px;
  background: #ffcc00;
  color: #333;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  z-index: 60;
  animation: slideDown 0.3s ease-out;
}

.offline-alert-content {
  display: flex;
  align-items: center;
  padding: 12px 16px;
}

.offline-alert-content svg {
  margin-right: 12px;
}

/* Animations */
@keyframes slideUp {
  from { opacity: 0; transform: translate(-50%, 20px); }
  to { opacity: 1; transform: translate(-50%, 0); }
}

@keyframes slideDown {
  from { opacity: 0; transform: translate(-50%, -20px); }
  to { opacity: 1; transform: translate(-50%, 0); }
}

@keyframes rotate {
  from { transform: rotate(0deg); }
  to { transform: rotate(360deg); }
}
</style>
```