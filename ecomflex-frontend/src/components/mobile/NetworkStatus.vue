```vue
<template>
  <div class="network-status">
    <div v-if="shouldShowBanner" class="network-banner" :class="[networkStatus, { 'is-visible': isVisible }]">
      <div class="network-banner-content">
        <div class="network-icon">
          <svg v-if="networkStatus === 'offline'" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="1" y1="1" x2="23" y2="23"></line>
            <path d="M16.72 11.06A10.94 10.94 0 0 1 19 12.55"></path>
            <path d="M5 12.55a10.94 10.94 0 0 1 5.17-2.39"></path>
            <path d="M10.71 5.05A16 16 0 0 1 22.58 9"></path>
            <path d="M1.42 9a15.91 15.91 0 0 1 4.7-2.88"></path>
            <path d="M8.53 16.11a6 6 0 0 1 6.95 0"></path>
            <line x1="12" y1="20" x2="12.01" y2="20"></line>
          </svg>
          <svg v-else-if="networkStatus === 'slow'" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"></path>
            <path d="M22 4L12 14.01l-3-3"></path>
          </svg>
          <svg v-else xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M5 12.55a11 11 0 0 1 14.08 0"></path>
            <path d="M1.42 9a16 16 0 0 1 21.16 0"></path>
            <path d="M8.53 16.11a6 6 0 0 1 6.95 0"></path>
            <line x1="12" y1="20" x2="12.01" y2="20"></line>
          </svg>
        </div>
        <div class="network-message">
          <span v-if="networkStatus === 'offline'">You're offline. Some features may be limited.</span>
          <span v-else-if="networkStatus === 'slow'">Slow connection detected.</span>
          <span v-else>You're back online!</span>
        </div>
        <button v-if="networkStatus === 'online'" @click="dismissBanner" class="dismiss-button">
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>
    </div>
    
    <!-- Network Diagnostics (shown when user clicks on network status) -->
    <div v-if="showDiagnostics" class="network-diagnostics" :class="{ 'is-open': showDiagnostics }">
      <div class="diagnostics-header">
        <h3>Connection Diagnostics</h3>
        <button @click="closeDiagnostics" class="close-button">
          <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <line x1="18" y1="6" x2="6" y2="18"></line>
            <line x1="6" y1="6" x2="18" y2="18"></line>
          </svg>
        </button>
      </div>
      
      <div class="diagnostics-content">
        <div class="diagnostic-item">
          <div class="diagnostic-label">Status:</div>
          <div class="diagnostic-value" :class="networkStatus">
            {{ networkStatus === 'offline' ? 'Offline' : networkStatus === 'slow' ? 'Slow Connection' : 'Online' }}
          </div>
        </div>
        
        <div v-if="networkStatus !== 'offline'" class="diagnostic-item">
          <div class="diagnostic-label">Connection Type:</div>
          <div class="diagnostic-value">
            {{ connectionType }}
          </div>
        </div>
        
        <div v-if="networkStatus !== 'offline'" class="diagnostic-item">
          <div class="diagnostic-label">Network Speed:</div>
          <div class="diagnostic-value">
            {{ formatSpeed(lastSpeed) }}
          </div>
        </div>
        
        <div v-if="networkStatus !== 'offline'" class="diagnostic-item">
          <div class="diagnostic-label">Latency:</div>
          <div class="diagnostic-value" :class="getLatencyClass(lastLatency)">
            {{ lastLatency }}ms
          </div>
        </div>
        
        <div class="diagnostic-history">
          <h4>Connection History</h4>
          <div v-for="(event, index) in connectionHistory" :key="index" class="history-item">
            <div class="history-time">{{ formatTime(event.timestamp) }}</div>
            <div class="history-status" :class="event.status">{{ event.status }}</div>
            <div class="history-details">{{ event.details }}</div>
          </div>
        </div>
        
        <div class="diagnostic-actions">
          <button @click="runNetworkTest" class="test-button" :disabled="isTestRunning">
            <span v-if="isTestRunning">Testing...</span>
            <span v-else>Test Connection</span>
          </button>
          <button @click="syncOfflineData" class="sync-button" :disabled="networkStatus === 'offline' || isSyncing">
            <span v-if="isSyncing">Syncing...</span>
            <span v-else>Sync Offline Data</span>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, onMounted, onBeforeUnmount, watch } from 'vue';

// Define NetworkInformation interface for TypeScript
declare global {
  interface NetworkInformation {
    effectiveType?: string;
    downlink?: number;
    rtt?: number;
    saveData?: boolean;
    type?: string;
  }

  interface Navigator {
    connection?: NetworkInformation;
  }
}

export default defineComponent({
  name: 'NetworkStatus',
  setup() {
    const networkStatus = ref('online'); // 'online', 'offline', 'slow'
    const isVisible = ref(false);
    const lastCheckTime = ref(0);
    const showDiagnostics = ref(false);
    const connectionType = ref('unknown');
    const lastSpeed = ref(0);
    const lastLatency = ref(0);
    const isTestRunning = ref(false);
    const isSyncing = ref(false);
    
    // Keep a history of network events
    const connectionHistory = ref([
      { timestamp: Date.now(), status: 'online', details: 'Initial connection' }
    ]);
    
    // Show banner for certain statuses
    const shouldShowBanner = computed(() => {
      // Always show offline banner
      if (networkStatus.value === 'offline') return true;
      
      // Show online banner only temporarily after reconnecting
      if (networkStatus.value === 'online' && Date.now() - lastCheckTime.value < 5000) return true;
      
      // Show slow connection banner
      if (networkStatus.value === 'slow') return true;
      
      return false;
    });
    
    // Detect network status changes
    const handleOnline = () => {
      networkStatus.value = 'online';
      isVisible.value = true;
      lastCheckTime.value = Date.now();
      addToHistory('online', 'Connection restored');
      
      // Run a quick network test to see if connection is slow
      setTimeout(() => {
        testConnectionSpeed();
      }, 1000);
    };
    
    const handleOffline = () => {
      networkStatus.value = 'offline';
      isVisible.value = true;
      addToHistory('offline', 'Connection lost');
    };
    
    // Test connection speed
    const testConnectionSpeed = async () => {
      if (navigator.onLine === false) return;
      
      try {
        const startTime = Date.now();
        const response = await fetch('https://www.google.com/favicon.ico', { 
          cache: 'no-store',
          headers: { 'Cache-Control': 'no-cache' }
        });
        
        if (!response.ok) throw new Error('Network test failed');
        
        const endTime = Date.now();
        lastLatency.value = endTime - startTime;
        
        // Get the blob to measure size
        const blob = await response.blob();
        const size = blob.size; // in bytes
        const timeTaken = endTime - startTime; // in ms
        
        // Calculate speed in kbps
        lastSpeed.value = (size * 8) / (timeTaken / 1000) / 1024;
        
        // Determine if connection is slow
        if (lastLatency.value > 500 || lastSpeed.value < 500) {
          networkStatus.value = 'slow';
          isVisible.value = true;
          addToHistory('slow', `Latency: ${lastLatency.value}ms, Speed: ${formatSpeed(lastSpeed.value)}`);
        }
        
        // Update connection type
        detectConnectionType();
      } catch (error) {
        console.error('Connection test failed:', error);
        addToHistory('error', 'Connection test failed');
      }
    };
    
    // Detect connection type
    const detectConnectionType = () => {
      if (navigator.connection) {
        const conn = navigator.connection;
        connectionType.value = conn.effectiveType || 'unknown';
      } else {
        connectionType.value = 'unknown';
      }
    };
    
    // Add event to history
    const addToHistory = (status: string, details: string) => {
      connectionHistory.value.unshift({
        timestamp: Date.now(),
        status,
        details
      });
      
      // Keep history limited to 10 items
      if (connectionHistory.value.length > 10) {
        connectionHistory.value.pop();
      }
    };
    
    // Format speed for display
    const formatSpeed = (speed: number): string => {
      if (speed < 1000) {
        return `${Math.round(speed)} Kbps`;
      } else {
        return `${(speed / 1024).toFixed(1)} Mbps`;
      }
    };
    
    // Format time for display
    const formatTime = (timestamp: number): string => {
      const date = new Date(timestamp);
      const hours = date.getHours().toString().padStart(2, '0');
      const minutes = date.getMinutes().toString().padStart(2, '0');
      const seconds = date.getSeconds().toString().padStart(2, '0');
      return `${hours}:${minutes}:${seconds}`;
    };
    
    // Get class for latency display
    const getLatencyClass = (latency: number): string => {
      if (latency < 100) return 'good';
      if (latency < 300) return 'moderate';
      return 'poor';
    };
    
    // Dismiss the online banner
    const dismissBanner = () => {
      isVisible.value = false;
    };
    
    // Open and close diagnostics panel
    const openDiagnostics = () => {
      showDiagnostics.value = true;
    };
    
    const closeDiagnostics = () => {
      showDiagnostics.value = false;
    };
    
    // Run a comprehensive network test
    const runNetworkTest = async () => {
      if (isTestRunning.value) return;
      
      isTestRunning.value = true;
      
      try {
        // Multiple tests for better accuracy
        let totalLatency = 0;
        let totalSpeed = 0;
        const testCount = 3;
        
        for (let i = 0; i < testCount; i++) {
          const startTime = Date.now();
          const response = await fetch(`https://www.google.com/favicon.ico?_=${Date.now()}`, { 
            cache: 'no-store',
            headers: { 'Cache-Control': 'no-cache' }
          });
          
          if (!response.ok) throw new Error('Network test failed');
          
          const endTime = Date.now();
          const latency = endTime - startTime;
          totalLatency += latency;
          
          // Get the blob to measure size
          const blob = await response.blob();
          const size = blob.size; // in bytes
          const timeTaken = endTime - startTime; // in ms
          
          // Calculate speed in kbps
          const speed = (size * 8) / (timeTaken / 1000) / 1024;
          totalSpeed += speed;
        }
        
        // Average results
        lastLatency.value = Math.round(totalLatency / testCount);
        lastSpeed.value = totalSpeed / testCount;
        
        // Update status based on results
        if (lastLatency.value > 500 || lastSpeed.value < 500) {
          networkStatus.value = 'slow';
          isVisible.value = true;
        } else {
          networkStatus.value = 'online';
        }
        
        addToHistory('test', `Test completed: Latency: ${lastLatency.value}ms, Speed: ${formatSpeed(lastSpeed.value)}`);
      } catch (error) {
        console.error('Connection test failed:', error);
        addToHistory('error', 'Comprehensive test failed');
      } finally {
        isTestRunning.value = false;
      }
    };
    
    // Sync offline data
    const syncOfflineData = async () => {
      if (networkStatus.value === 'offline' || isSyncing.value) return;
      
      isSyncing.value = true;
      
      try {
        // Simulate syncing process
        await new Promise(resolve => setTimeout(resolve, 2000));
        
        // In a real app, this would sync IndexedDB data with server
        
        addToHistory('sync', 'Offline data synced successfully');
      } catch (error) {
        console.error('Sync failed:', error);
        addToHistory('error', 'Data sync failed');
      } finally {
        isSyncing.value = false;
      }
    };
    
    // Watch for network status changes
    watch(networkStatus, (newStatus) => {
      // Automatically hide online banner after 5 seconds
      if (newStatus === 'online') {
        setTimeout(() => {
          if (networkStatus.value === 'online') {
            isVisible.value = false;
          }
        }, 5000);
      }
    });
    
    onMounted(() => {
      // Add event listeners
      window.addEventListener('online', handleOnline);
      window.addEventListener('offline', handleOffline);
      
      // Check initial status
      if (navigator.onLine) {
        networkStatus.value = 'online';
        
        // Run initial connection test
        setTimeout(() => {
          testConnectionSpeed();
        }, 1000);
      } else {
        networkStatus.value = 'offline';
        isVisible.value = true;
      }
      
      // Listen for custom events from the service worker
      document.addEventListener('app:online', handleOnline);
      document.addEventListener('app:offline', handleOffline);
      
      // Make the banner clickable to show diagnostics
      const banner = document.querySelector('.network-banner');
      if (banner) {
        banner.addEventListener('click', openDiagnostics);
      }
    });
    
    onBeforeUnmount(() => {
      // Remove event listeners
      window.removeEventListener('online', handleOnline);
      window.removeEventListener('offline', handleOffline);
      document.removeEventListener('app:online', handleOnline);
      document.removeEventListener('app:offline', handleOffline);
      
      const banner = document.querySelector('.network-banner');
      if (banner) {
        banner.removeEventListener('click', openDiagnostics);
      }
    });
    
    return {
      networkStatus,
      isVisible,
      shouldShowBanner,
      showDiagnostics,
      connectionType,
      lastSpeed,
      lastLatency,
      connectionHistory,
      isTestRunning,
      isSyncing,
      dismissBanner,
      closeDiagnostics,
      formatSpeed,
      formatTime,
      getLatencyClass,
      runNetworkTest,
      syncOfflineData
    };
  }
});
</script>

<style scoped>
.network-status {
  z-index: 1000;
}

.network-banner {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  padding: 10px 16px;
  transform: translateY(-100%);
  transition: transform 0.3s ease-in-out;
  z-index: 1010;
  font-size: 14px;
  cursor: pointer;
}

.network-banner.is-visible {
  transform: translateY(0);
}

.network-banner.offline {
  background-color: #ffcc00;
  color: #333;
}

.network-banner.slow {
  background-color: #ff9800;
  color: #333;
}

.network-banner.online {
  background-color: #4caf50;
  color: white;
}

.network-banner-content {
  display: flex;
  align-items: center;
  max-width: 1200px;
  margin: 0 auto;
}

.network-icon {
  margin-right: 10px;
  display: flex;
  align-items: center;
}

.network-message {
  flex: 1;
}

.dismiss-button {
  background: transparent;
  border: none;
  color: inherit;
  cursor: pointer;
  padding: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0.7;
  transition: opacity 0.2s ease;
}

.dismiss-button:hover {
  opacity: 1;
}

/* Network Diagnostics */
.network-diagnostics {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%) scale(0.9);
  width: 90%;
  max-width: 400px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
  z-index: 1020;
  opacity: 0;
  pointer-events: none;
  transition: all 0.3s ease;
  overflow: hidden;
}

.network-diagnostics.is-open {
  opacity: 1;
  pointer-events: auto;
  transform: translate(-50%, -50%) scale(1);
}

.diagnostics-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 16px 20px;
  border-bottom: 1px solid #eee;
}

.diagnostics-header h3 {
  margin: 0;
  font-size: 18px;
  color: #333;
}

.close-button {
  background: transparent;
  border: none;
  color: #666;
  cursor: pointer;
  padding: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.2s ease;
}

.close-button:hover {
  color: #333;
}

.diagnostics-content {
  padding: 20px;
  max-height: 60vh;
  overflow-y: auto;
}

.diagnostic-item {
  display: flex;
  margin-bottom: 16px;
  padding-bottom: 16px;
  border-bottom: 1px solid #eee;
}

.diagnostic-label {
  flex: 1;
  font-weight: 500;
  color: #666;
}

.diagnostic-value {
  font-weight: 600;
}

.diagnostic-value.offline {
  color: #e53935;
}

.diagnostic-value.slow {
  color: #ff9800;
}

.diagnostic-value.online {
  color: #4caf50;
}

.diagnostic-value.good {
  color: #4caf50;
}

.diagnostic-value.moderate {
  color: #ff9800;
}

.diagnostic-value.poor {
  color: #e53935;
}

.diagnostic-history {
  margin-top: 20px;
}

.diagnostic-history h4 {
  font-size: 16px;
  margin: 0 0 16px;
  color: #333;
}

.history-item {
  display: flex;
  padding: 10px 0;
  font-size: 13px;
  border-bottom: 1px solid #f5f5f5;
}

.history-time {
  width: 80px;
  color: #666;
}

.history-status {
  width: 80px;
  font-weight: 600;
  margin-right: 10px;
}

.history-status.online {
  color: #4caf50;
}

.history-status.offline {
  color: #e53935;
}

.history-status.slow {
  color: #ff9800;
}

.history-status.error {
  color: #e53935;
}

.history-status.test {
  color: #2196f3;
}

.history-status.sync {
  color: #9c27b0;
}

.history-details {
  flex: 1;
  color: #333;
}

.diagnostic-actions {
  display: flex;
  gap: 10px;
  margin-top: 20px;
}

.test-button, .sync-button {
  flex: 1;
  background: #007BFF;
  color: white;
  border: none;
  border-radius: 6px;
  padding: 12px;
  font-weight: 500;
  cursor: pointer;
  transition: background 0.2s ease;
}

.test-button:hover, .sync-button:hover {
  background: #0056b3;
}

.test-button:disabled, .sync-button:disabled {
  background: #ccc;
  cursor: not-allowed;
}

.sync-button {
  background: #4caf50;
}

.sync-button:hover {
  background: #388e3c;
}
</style>
```