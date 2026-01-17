```vue
<template>
  <div class="notification-center">
    <!-- Notification Permission Request -->
    <div v-if="showPermissionRequest" class="permission-request">
      <div class="permission-content">
        <div class="permission-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round">
            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
            <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
          </svg>
        </div>
        <div class="permission-text">
          <h3>Stay Updated</h3>
          <p>Enable notifications to get updates about your bookings, cashback, and new products.</p>
        </div>
        <div class="permission-actions">
          <button class="enable-button" @click="requestPermission">
            Enable Notifications
          </button>
          <button class="dismiss-button" @click="dismissPermissionRequest">
            Maybe Later
          </button>
        </div>
      </div>
    </div>
    
    <!-- Notification Drawer (for mobile) -->
    <div class="notification-drawer" :class="{ 'is-open': drawerOpen }">
      <div class="drawer-header">
        <h3>Notifications</h3>
        <button @click="toggleDrawer" class="close-drawer">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>
      
      <div class="drawer-content">
        <div v-if="notifications.length === 0" class="empty-state">
          <div class="empty-icon">
            <svg xmlns="http://www.w3.org/2000/svg" width="48" height="48" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1" stroke-linecap="round" stroke-linejoin="round">
              <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
              <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
            </svg>
          </div>
          <p>No notifications yet</p>
        </div>
        
        <div v-else class="notification-list">
          <div 
            v-for="(notification, index) in notifications" 
            :key="index"
            class="notification-item"
            :class="{ 'unread': !notification.read }"
            @click="markAsRead(notification.id)"
          >
            <div class="notification-badge" v-if="!notification.read"></div>
            <div class="notification-icon" :class="notification.type">
              <svg v-if="notification.type === 'booking'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                <line x1="16" y1="2" x2="16" y2="6"></line>
                <line x1="8" y1="2" x2="8" y2="6"></line>
                <line x1="3" y1="10" x2="21" y2="10"></line>
              </svg>
              
              <svg v-else-if="notification.type === 'cashback'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="12" y1="1" x2="12" y2="23"></line>
                <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
              </svg>
              
              <svg v-else-if="notification.type === 'product'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"></path>
                <polyline points="3.27 6.96 12 12.01 20.73 6.96"></polyline>
                <line x1="12" y1="22.08" x2="12" y2="12"></line>
              </svg>
              
              <svg v-else xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <circle cx="12" cy="12" r="10"></circle>
                <line x1="12" y1="8" x2="12" y2="12"></line>
                <line x1="12" y1="16" x2="12.01" y2="16"></line>
              </svg>
            </div>
            
            <div class="notification-content">
              <div class="notification-title">{{ notification.title }}</div>
              <div class="notification-body">{{ notification.body }}</div>
              <div class="notification-time">{{ formatTime(notification.timestamp) }}</div>
            </div>
            
            <button class="delete-button" @click.stop="deleteNotification(notification.id)">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <polyline points="3 6 5 6 21 6"></polyline>
                <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"></path>
              </svg>
            </button>
          </div>
        </div>
      </div>
      
      <div class="drawer-footer">
        <button 
          v-if="notifications.length > 0" 
          class="clear-all-button"
          @click="clearAllNotifications"
        >
          Clear All
        </button>
        
        <div class="notification-settings">
          <button class="settings-button" @click="openSettings">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <circle cx="12" cy="12" r="3"></circle>
              <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
            </svg>
            Notification Settings
          </button>
        </div>
      </div>
    </div>
    
    <!-- Notification Settings Modal -->
    <div v-if="showSettings" class="settings-modal">
      <div class="modal-overlay" @click="closeSettings"></div>
      <div class="modal-content">
        <div class="modal-header">
          <h3>Notification Settings</h3>
          <button @click="closeSettings" class="close-modal">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <line x1="18" y1="6" x2="6" y2="18"></line>
              <line x1="6" y1="6" x2="18" y2="18"></line>
            </svg>
          </button>
        </div>
        
        <div class="modal-body">
          <div class="settings-section">
            <h4>Notification Types</h4>
            
            <div class="setting-item">
              <div class="setting-label">
                <div class="setting-icon booking">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect>
                    <line x1="16" y1="2" x2="16" y2="6"></line>
                    <line x1="8" y1="2" x2="8" y2="6"></line>
                    <line x1="3" y1="10" x2="21" y2="10"></line>
                  </svg>
                </div>
                <span>Booking Updates</span>
              </div>
              <label class="toggle">
                <input type="checkbox" v-model="settings.bookingUpdates">
                <span class="toggle-slider"></span>
              </label>
            </div>
            
            <div class="setting-item">
              <div class="setting-label">
                <div class="setting-icon cashback">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <line x1="12" y1="1" x2="12" y2="23"></line>
                    <path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"></path>
                  </svg>
                </div>
                <span>Cashback Alerts</span>
              </div>
              <label class="toggle">
                <input type="checkbox" v-model="settings.cashbackAlerts">
                <span class="toggle-slider"></span>
              </label>
            </div>
            
            <div class="setting-item">
              <div class="setting-label">
                <div class="setting-icon product">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"></path>
                    <polyline points="3.27 6.96 12 12.01 20.73 6.96"></polyline>
                    <line x1="12" y1="22.08" x2="12" y2="12"></line>
                  </svg>
                </div>
                <span>New Products</span>
              </div>
              <label class="toggle">
                <input type="checkbox" v-model="settings.newProducts">
                <span class="toggle-slider"></span>
              </label>
            </div>
            
            <div class="setting-item">
              <div class="setting-label">
                <div class="setting-icon system">
                  <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <rect x="2" y="2" width="20" height="8" rx="2" ry="2"></rect>
                    <rect x="2" y="14" width="20" height="8" rx="2" ry="2"></rect>
                    <line x1="6" y1="6" x2="6.01" y2="6"></line>
                    <line x1="6" y1="18" x2="6.01" y2="18"></line>
                  </svg>
                </div>
                <span>System Updates</span>
              </div>
              <label class="toggle">
                <input type="checkbox" v-model="settings.systemUpdates">
                <span class="toggle-slider"></span>
              </label>
            </div>
          </div>
          
          <div class="settings-section">
            <h4>Notification Preferences</h4>
            
            <div class="setting-item">
              <div class="setting-label">
                <span>Push Notifications</span>
              </div>
              <label class="toggle">
                <input type="checkbox" v-model="settings.pushEnabled" :disabled="!isPushSupported">
                <span class="toggle-slider"></span>
              </label>
            </div>
            
            <div class="setting-item">
              <div class="setting-label">
                <span>Sound</span>
              </div>
              <label class="toggle">
                <input type="checkbox" v-model="settings.sound">
                <span class="toggle-slider"></span>
              </label>
            </div>
            
            <div class="setting-item">
              <div class="setting-label">
                <span>Vibration</span>
              </div>
              <label class="toggle">
                <input type="checkbox" v-model="settings.vibration">
                <span class="toggle-slider"></span>
              </label>
            </div>
          </div>
        </div>
        
        <div class="modal-footer">
          <button class="save-button" @click="saveSettings">Save Settings</button>
        </div>
      </div>
    </div>
    
    <!-- Floating Notification Bell -->
    <div v-if="isAuthenticated && !drawerOpen" class="notification-bell" @click="toggleDrawer">
      <div class="bell-icon">
        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9"></path>
          <path d="M13.73 21a2 2 0 0 1-3.46 0"></path>
        </svg>
        <div v-if="unreadCount > 0" class="unread-badge">{{ unreadCount > 9 ? '9+' : unreadCount }}</div>
      </div>
    </div>
    
    <!-- Toast Notification -->
    <transition name="toast">
      <div v-if="isToastVisible" class="toast-notification" :class="toastType">
        <div class="toast-icon">
          <svg v-if="toastType === 'success'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
            <polyline points="22 4 12 14.01 9 11.01"></polyline>
          </svg>
          
          <svg v-else-if="toastType === 'error'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="10"></circle>
            <line x1="15" y1="9" x2="9" y2="15"></line>
            <line x1="9" y1="9" x2="15" y2="15"></line>
          </svg>
          
          <svg v-else-if="toastType === 'warning'" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M10.29 3.86L1.82 18a2 2 0 0 0 1.71 3h16.94a2 2 0 0 0 1.71-3L13.71 3.86a2 2 0 0 0-3.42 0z"></path>
            <line x1="12" y1="9" x2="12" y2="13"></line>
            <line x1="12" y1="17" x2="12.01" y2="17"></line>
          </svg>
          
          <svg v-else xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <circle cx="12" cy="12" r="10"></circle>
            <line x1="12" y1="8" x2="12" y2="12"></line>
            <line x1="12" y1="16" x2="12.01" y2="16"></line>
          </svg>
        </div>
        
        <div class="toast-content">
          <div class="toast-title">{{ toastTitle }}</div>
          <div class="toast-message">{{ toastMessage }}</div>
        </div>
        
        <button class="toast-close" @click="closeToast">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, onMounted, onBeforeUnmount, watch } from 'vue';
import pushNotificationManager from '@/services/pushNotificationManager';

interface Notification {
  id: string;
  type: 'booking' | 'cashback' | 'product' | 'system';
  title: string;
  body: string;
  read: boolean;
  timestamp: number;
  data?: any;
}

interface NotificationSettings {
  bookingUpdates: boolean;
  cashbackAlerts: boolean;
  newProducts: boolean;
  systemUpdates: boolean;
  pushEnabled: boolean;
  sound: boolean;
  vibration: boolean;
}

export default defineComponent({
  name: 'NotificationCenter',
  props: {
    isAuthenticated: {
      type: Boolean,
      default: false
    }
  },
  setup(props) {
    const drawerOpen = ref(false);
    const showSettings = ref(false);
    const showPermissionRequest = ref(false);
    const isToastVisible = ref(false);
    const toastTitle = ref('');
    const toastMessage = ref('');
    const toastType = ref('info');
    const toastTimeout = ref<number | null>(null);
    
    // Mock notifications for demo
    const notifications = ref<Notification[]>([
      {
        id: '1',
        type: 'booking',
        title: 'Booking Confirmed',
        body: 'Your booking for Wireless Earbuds has been confirmed',
        read: false,
        timestamp: Date.now() - 1000 * 60 * 30 // 30 minutes ago
      },
      {
        id: '2',
        type: 'cashback',
        title: 'Cashback Received',
        body: 'You received $59.99 cashback for Wireless Earbuds',
        read: true,
        timestamp: Date.now() - 1000 * 60 * 60 * 2 // 2 hours ago
      },
      {
        id: '3',
        type: 'product',
        title: 'New Product Available',
        body: 'Check out the new Smart WiFi Coffee Maker with cashback',
        read: false,
        timestamp: Date.now() - 1000 * 60 * 60 * 24 // 1 day ago
      }
    ]);
    
    // Notification settings
    const settings = ref<NotificationSettings>({
      bookingUpdates: true,
      cashbackAlerts: true,
      newProducts: true,
      systemUpdates: true,
      pushEnabled: false,
      sound: true,
      vibration: true
    });
    
    // Computed properties
    const unreadCount = computed(() => {
      return notifications.value.filter(notification => !notification.read).length;
    });
    
    const isPushSupported = computed(() => {
      return pushNotificationManager.isPushSupported();
    });
    
    // Check if we should show the permission request
    const checkPermissionRequestStatus = () => {
      if (!isPushSupported.value) {
        return;
      }
      
      if (props.isAuthenticated && pushNotificationManager.permissionStatus === 'default') {
        // Check if the user has dismissed the request recently
        const lastDismissed = localStorage.getItem('notificationPermissionDismissed');
        if (!lastDismissed || Date.now() - parseInt(lastDismissed) > 7 * 24 * 60 * 60 * 1000) {
          showPermissionRequest.value = true;
        }
      }
    };
    
    // Load settings from localStorage
    const loadSettings = () => {
      const savedSettings = localStorage.getItem('notificationSettings');
      if (savedSettings) {
        settings.value = JSON.parse(savedSettings);
      }
      
      // Update push enabled based on current permission
      if (pushNotificationManager.permissionStatus === 'granted') {
        settings.value.pushEnabled = true;
      }
    };
    
    // Save settings to localStorage
    const saveSettings = () => {
      localStorage.setItem('notificationSettings', JSON.stringify(settings.value));
      
      // If push enabled changed, handle subscription
      if (settings.value.pushEnabled && pushNotificationManager.permissionStatus !== 'granted') {
        requestPermission();
      } else if (!settings.value.pushEnabled && pushNotificationManager.permissionStatus === 'granted') {
        pushNotificationManager.unsubscribeFromPush();
      }
      
      showSettings.value = false;
      displayToast('Success', 'Notification settings saved', 'success');
    };
    
    // Toggle notification drawer
    const toggleDrawer = () => {
      drawerOpen.value = !drawerOpen.value;
    };
    
    // Open settings modal
    const openSettings = () => {
      showSettings.value = true;
    };
    
    // Close settings modal
    const closeSettings = () => {
      showSettings.value = false;
    };
    
    // Mark notification as read
    const markAsRead = (id: string) => {
      const notification = notifications.value.find(n => n.id === id);
      if (notification) {
        notification.read = true;
      }
    };
    
    // Delete notification
    const deleteNotification = (id: string) => {
      notifications.value = notifications.value.filter(n => n.id !== id);
    };
    
    // Clear all notifications
    const clearAllNotifications = () => {
      notifications.value = [];
    };
    
    // Format time for display
    const formatTime = (timestamp: number): string => {
      const now = Date.now();
      const diff = now - timestamp;
      
      // Less than a minute
      if (diff < 60 * 1000) {
        return 'Just now';
      }
      
      // Less than an hour
      if (diff < 60 * 60 * 1000) {
        const minutes = Math.floor(diff / (60 * 1000));
        return `${minutes} ${minutes === 1 ? 'minute' : 'minutes'} ago`;
      }
      
      // Less than a day
      if (diff < 24 * 60 * 60 * 1000) {
        const hours = Math.floor(diff / (60 * 60 * 1000));
        return `${hours} ${hours === 1 ? 'hour' : 'hours'} ago`;
      }
      
      // Less than a week
      if (diff < 7 * 24 * 60 * 60 * 1000) {
        const days = Math.floor(diff / (24 * 60 * 60 * 1000));
        return `${days} ${days === 1 ? 'day' : 'days'} ago`;
      }
      
      // Format as date
      const date = new Date(timestamp);
      return date.toLocaleDateString();
    };
    
    // Request notification permission
    const requestPermission = async () => {
      const permission = await pushNotificationManager.requestPermission();
      
      if (permission === 'granted') {
        settings.value.pushEnabled = true;
        saveSettings();
        displayToast('Success', 'Notifications enabled', 'success');
      } else {
        settings.value.pushEnabled = false;
        saveSettings();
        displayToast('Notice', 'Notification permission denied', 'warning');
      }
      
      showPermissionRequest.value = false;
    };
    
    // Dismiss permission request
    const dismissPermissionRequest = () => {
      showPermissionRequest.value = false;
      localStorage.setItem('notificationPermissionDismissed', Date.now().toString());
    };
    
    // Show toast notification
    const displayToast = (title: string, message: string, type: string = 'info') => {
      // Clear any existing timeout
      if (toastTimeout.value) {
        clearTimeout(toastTimeout.value);
      }
      
      toastTitle.value = title;
      toastMessage.value = message;
      toastType.value = type;
      isToastVisible.value = true;
      
      // Auto-hide after 5 seconds
      toastTimeout.value = window.setTimeout(() => {
        closeToast();
      }, 5000);
    };
    
    // Close toast notification
    const closeToast = () => {
      isToastVisible.value = false;
      if (toastTimeout.value) {
        clearTimeout(toastTimeout.value);
        toastTimeout.value = null;
      }
    };
    
    // Add a new notification
    const addNotification = (notification: Omit<Notification, 'id' | 'read' | 'timestamp'>) => {
      const id = `notification_${Date.now()}_${Math.random().toString(36).substr(2, 9)}`;
      
      notifications.value.unshift({
        id,
        ...notification,
        read: false,
        timestamp: Date.now()
      });
      
      // Show toast for the new notification
      displayToast(notification.title, notification.body, 'info');
      
      // Play sound if enabled
      if (settings.value.sound) {
        playNotificationSound();
      }
      
      // Vibrate if enabled and supported
      if (settings.value.vibration && 'vibrate' in navigator) {
        navigator.vibrate([200, 100, 200]);
      }
    };
    
    // Play notification sound
    const playNotificationSound = () => {
      const audio = new Audio('/sounds/notification.mp3');
      audio.play().catch(error => {
        console.error('Error playing notification sound:', error);
      });
    };
    
    // Initialize push notification manager
    const initPushManager = async () => {
      if (isPushSupported.value) {
        // This should be your actual VAPID public key
        const applicationServerKey = 'BPtE2ZobBk3JRwHvK5Nt9LsGUdvE-aV2CvwkKF6i0nuM6-8SjwKHcJa4xVCONMXi1A9tMqWpZAl5NZrfBLOIXLY';
        await pushNotificationManager.initialize(applicationServerKey);
        
        // Update settings based on permission
        if (pushNotificationManager.permissionStatus === 'granted') {
          settings.value.pushEnabled = true;
        }
      }
    };
    
    // Listen for push notifications from service worker
    const setupNotificationListener = () => {
      if ('serviceWorker' in navigator) {
        navigator.serviceWorker.addEventListener('message', event => {
          if (event.data && event.data.type === 'NOTIFICATION') {
            addNotification({
              type: event.data.notificationType || 'system',
              title: event.data.title,
              body: event.data.body,
              data: event.data.data
            });
          }
        });
      }
    };
    
    // Watch for authentication changes
    watch(() => props.isAuthenticated, (newValue) => {
      if (newValue) {
        checkPermissionRequestStatus();
      }
    });
    
    onMounted(() => {
      loadSettings();
      initPushManager();
      setupNotificationListener();
      
      if (props.isAuthenticated) {
        checkPermissionRequestStatus();
      }
    });
    
    onBeforeUnmount(() => {
      if (toastTimeout.value) {
        clearTimeout(toastTimeout.value);
      }
    });
    
    return {
      drawerOpen,
      showSettings,
      showPermissionRequest,
      isToastVisible,
      toastTitle,
      toastMessage,
      toastType,
      notifications,
      settings,
      unreadCount,
      isPushSupported,
      toggleDrawer,
      openSettings,
      closeSettings,
      markAsRead,
      deleteNotification,
      clearAllNotifications,
      formatTime,
      requestPermission,
      dismissPermissionRequest,
      closeToast,
      saveSettings,
      displayToast
    };
  }
});
</script>

<style scoped>
.notification-center {
  z-index: 1000;
}

/* Permission Request */
.permission-request {
  position: fixed;
  bottom: 20px;
  left: 50%;
  transform: translateX(-50%);
  width: 90%;
  max-width: 400px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.15);
  z-index: 1010;
  animation: slideUp 0.3s ease-out;
  overflow: hidden;
}

.permission-content {
  padding: 20px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
}

.permission-icon {
  color: #007BFF;
  margin-bottom: 16px;
}

.permission-text h3 {
  font-size: 18px;
  margin: 0 0 8px 0;
  color: #333;
}

.permission-text p {
  font-size: 14px;
  color: #666;
  margin: 0 0 20px 0;
  line-height: 1.5;
}

.permission-actions {
  display: flex;
  flex-direction: column;
  gap: 12px;
  width: 100%;
}

.enable-button {
  background: #007BFF;
  color: white;
  border: none;
  border-radius: 8px;
  padding: 12px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s ease;
}

.enable-button:hover {
  background: #0056b3;
}

.dismiss-button {
  background: transparent;
  color: #666;
  border: none;
  padding: 8px;
  cursor: pointer;
  transition: color 0.2s ease;
}

.dismiss-button:hover {
  color: #333;
}

/* Notification Drawer */
.notification-drawer {
  position: fixed;
  top: 0;
  right: 0;
  width: 100%;
  max-width: 400px;
  height: 100vh;
  background: white;
  box-shadow: -5px 0 20px rgba(0, 0, 0, 0.15);
  z-index: 1020;
  transform: translateX(100%);
  transition: transform 0.3s ease;
  display: flex;
  flex-direction: column;
}

.notification-drawer.is-open {
  transform: translateX(0);
}

.drawer-header {
  padding: 20px;
  border-bottom: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.drawer-header h3 {
  margin: 0;
  font-size: 20px;
  color: #333;
}

.close-drawer {
  background: transparent;
  border: none;
  color: #666;
  cursor: pointer;
  padding: 5px;
}

.drawer-content {
  flex: 1;
  overflow-y: auto;
  padding: 0 0 20px 0;
}

.empty-state {
  padding: 40px 20px;
  text-align: center;
  color: #999;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.empty-icon {
  margin-bottom: 16px;
  color: #ccc;
}

.notification-list {
  padding: 10px 0;
}

.notification-item {
  position: relative;
  padding: 16px 20px;
  border-bottom: 1px solid #f5f5f5;
  display: flex;
  align-items: flex-start;
  cursor: pointer;
  transition: background 0.2s ease;
}

.notification-item:hover {
  background: #f9f9f9;
}

.notification-item.unread {
  background: #f0f7ff;
}

.notification-item.unread:hover {
  background: #e6f0fa;
}

.notification-badge {
  position: absolute;
  top: 20px;
  left: 12px;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #007BFF;
}

.notification-icon {
  margin-right: 16px;
  padding: 10px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.notification-icon.booking {
  background: rgba(0, 123, 255, 0.1);
  color: #007BFF;
}

.notification-icon.cashback {
  background: rgba(76, 175, 80, 0.1);
  color: #4caf50;
}

.notification-icon.product {
  background: rgba(255, 152, 0, 0.1);
  color: #ff9800;
}

.notification-icon.system {
  background: rgba(158, 158, 158, 0.1);
  color: #9e9e9e;
}

.notification-content {
  flex: 1;
}

.notification-title {
  font-weight: 600;
  margin-bottom: 4px;
  color: #333;
}

.notification-body {
  font-size: 14px;
  color: #666;
  margin-bottom: 8px;
  line-height: 1.4;
}

.notification-time {
  font-size: 12px;
  color: #999;
}

.delete-button {
  background: transparent;
  border: none;
  color: #999;
  padding: 8px;
  cursor: pointer;
  opacity: 0;
  transition: opacity 0.2s ease, color 0.2s ease;
}

.notification-item:hover .delete-button {
  opacity: 1;
}

.delete-button:hover {
  color: #e53935;
}

.drawer-footer {
  padding: 16px 20px;
  border-top: 1px solid #eee;
}

.clear-all-button {
  width: 100%;
  background: transparent;
  border: 1px solid #ddd;
  color: #666;
  padding: 10px;
  border-radius: 8px;
  margin-bottom: 16px;
  cursor: pointer;
  transition: background 0.2s ease, color 0.2s ease;
}

.clear-all-button:hover {
  background: #f5f5f5;
  color: #333;
}

.settings-button {
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  background: transparent;
  border: none;
  color: #007BFF;
  padding: 10px;
  cursor: pointer;
  transition: color 0.2s ease;
}

.settings-button:hover {
  color: #0056b3;
}

/* Notification Bell */
.notification-bell {
  position: fixed;
  bottom: 80px;
  right: 20px;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background: #007BFF;
  box-shadow: 0 4px 10px rgba(0, 123, 255, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  z-index: 1000;
  transition: transform 0.2s ease, background 0.2s ease;
}

.notification-bell:hover {
  transform: translateY(-2px);
  background: #0056b3;
}

.bell-icon {
  position: relative;
  color: white;
}

.unread-badge {
  position: absolute;
  top: -8px;
  right: -8px;
  width: 20px;
  height: 20px;
  border-radius: 50%;
  background: #e53935;
  color: white;
  font-size: 12px;
  font-weight: 600;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Toast Notification */
.toast-notification {
  position: fixed;
  top: 20px;
  right: 20px;
  max-width: 400px;
  width: calc(100% - 40px);
  background: white;
  border-radius: 8px;
  box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
  padding: 16px;
  display: flex;
  align-items: flex-start;
  z-index: 1030;
}

.toast-notification.success {
  border-left: 4px solid #4caf50;
}

.toast-notification.error {
  border-left: 4px solid #e53935;
}

.toast-notification.warning {
  border-left: 4px solid #ff9800;
}

.toast-notification.info {
  border-left: 4px solid #2196f3;
}

.toast-icon {
  margin-right: 12px;
  padding-top: 2px;
}

.toast-notification.success .toast-icon {
  color: #4caf50;
}

.toast-notification.error .toast-icon {
  color: #e53935;
}

.toast-notification.warning .toast-icon {
  color: #ff9800;
}

.toast-notification.info .toast-icon {
  color: #2196f3;
}

.toast-content {
  flex: 1;
}

.toast-title {
  font-weight: 600;
  margin-bottom: 4px;
  color: #333;
}

.toast-message {
  font-size: 14px;
  color: #666;
}

.toast-close {
  background: transparent;
  border: none;
  color: #999;
  padding: 5px;
  cursor: pointer;
  margin-left: 8px;
}

/* Settings Modal */
.settings-modal {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 1040;
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  animation: fadeIn 0.3s ease;
}

.modal-content {
  position: relative;
  width: 90%;
  max-width: 500px;
  max-height: 90vh;
  background: white;
  border-radius: 12px;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.2);
  display: flex;
  flex-direction: column;
  animation: zoomIn 0.3s ease;
  overflow: hidden;
}

.modal-header {
  padding: 20px;
  border-bottom: 1px solid #eee;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h3 {
  margin: 0;
  font-size: 20px;
  color: #333;
}

.close-modal {
  background: transparent;
  border: none;
  color: #666;
  cursor: pointer;
  padding: 5px;
}

.modal-body {
  padding: 20px;
  overflow-y: auto;
  flex: 1;
}

.settings-section {
  margin-bottom: 24px;
}

.settings-section h4 {
  font-size: 16px;
  margin: 0 0 16px 0;
  color: #333;
  padding-bottom: 8px;
  border-bottom: 1px solid #eee;
}

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
}

.setting-label {
  display: flex;
  align-items: center;
}

.setting-icon {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 12px;
}

.setting-icon.booking {
  background: rgba(0, 123, 255, 0.1);
  color: #007BFF;
}

.setting-icon.cashback {
  background: rgba(76, 175, 80, 0.1);
  color: #4caf50;
}

.setting-icon.product {
  background: rgba(255, 152, 0, 0.1);
  color: #ff9800;
}

.setting-icon.system {
  background: rgba(158, 158, 158, 0.1);
  color: #9e9e9e;
}

/* Toggle Switch */
.toggle {
  position: relative;
  display: inline-block;
  width: 48px;
  height: 24px;
}

.toggle input {
  opacity: 0;
  width: 0;
  height: 0;
}

.toggle-slider {
  position: absolute;
  cursor: pointer;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #ccc;
  transition: 0.3s;
  border-radius: 24px;
}

.toggle-slider:before {
  position: absolute;
  content: "";
  height: 18px;
  width: 18px;
  left: 3px;
  bottom: 3px;
  background-color: white;
  transition: 0.3s;
  border-radius: 50%;
}

input:checked + .toggle-slider {
  background-color: #007BFF;
}

input:disabled + .toggle-slider {
  background-color: #ddd;
  cursor: not-allowed;
}

input:checked + .toggle-slider:before {
  transform: translateX(24px);
}

.modal-footer {
  padding: 16px 20px;
  border-top: 1px solid #eee;
  display: flex;
  justify-content: flex-end;
}

.save-button {
  background: #007BFF;
  color: white;
  border: none;
  border-radius: 8px;
  padding: 12px 24px;
  font-weight: 600;
  cursor: pointer;
  transition: background 0.2s ease;
}

.save-button:hover {
  background: #0056b3;
}

/* Animations */
@keyframes slideUp {
  from { opacity: 0; transform: translate(-50%, 20px); }
  to { opacity: 1; transform: translate(-50%, 0); }
}

@keyframes zoomIn {
  from { opacity: 0; transform: scale(0.95); }
  to { opacity: 1; transform: scale(1); }
}

@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

.toast-enter-active, .toast-leave-active {
  transition: transform 0.3s ease, opacity 0.3s ease;
}

.toast-enter-from, .toast-leave-to {
  transform: translateX(100%);
  opacity: 0;
}

/* Media Queries */
@media (max-width: 480px) {
  .notification-drawer {
    max-width: 100%;
  }
  
  .toast-notification {
    width: calc(100% - 20px);
    right: 10px;
    top: 10px;
  }
}
</style>
```