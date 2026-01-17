// src/pwa/register.ts
// Service worker registration for PWA functionality

export function registerServiceWorker() {
  // Note: Service Worker registration is now handled by VitePWA plugin
  // This function remains for compatibility but doesn't manually register
  
  if ('serviceWorker' in navigator) {
    // Handle service worker update events
    navigator.serviceWorker.ready.then(registration => {
      // Add event listener for new service workers
      registration.addEventListener('updatefound', () => {
        const newWorker = registration.installing;
        if (newWorker) {
          newWorker.addEventListener('statechange', () => {
            if (newWorker.state === 'installed' && navigator.serviceWorker.controller) {
              // Notify the app that a new version is available
              document.dispatchEvent(new CustomEvent('pwa:update-available'));
            }
          });
        }
      });
    });
    
    // Listen for the controllerchange event
    navigator.serviceWorker.addEventListener('controllerchange', () => {
      console.log('New service worker controller, page will reload...');
      // Only reload if we didn't trigger the navigation ourselves
      // Commented out to prevent reload loops
      // if (!window.__pwaUpdating) {
      //   window.location.reload();
      // }
    });
    
    // Handle network status changes
    window.addEventListener('online', () => {
      document.dispatchEvent(new CustomEvent('pwa:online'));
    });
    
    window.addEventListener('offline', () => {
      document.dispatchEvent(new CustomEvent('pwa:offline'));
    });
  }
}

// Function to update the service worker
export function updateServiceWorker() {
  if ('serviceWorker' in navigator) {
    // Flag to prevent duplicate reloads
    window.__pwaUpdating = true;
    
    navigator.serviceWorker.ready.then(registration => {
      if (registration.waiting) {
        // Send skip-waiting message to the waiting service worker
        registration.waiting.postMessage({ type: 'SKIP_WAITING' });
      }
    });
  }
}

// Function to check for installability
export function checkInstallability(): Promise<boolean> {
  return new Promise(resolve => {
    if ('serviceWorker' in navigator && 'BeforeInstallPromptEvent' in window) {
      window.addEventListener('beforeinstallprompt', (e: Event) => {
        // Prevent the default prompt
        e.preventDefault();
        
        // Store the event for later use
        window.__installPromptEvent = e;
        
        // Notify the app that it's installable
        document.dispatchEvent(new CustomEvent('pwa:installable'));
        
        resolve(true);
      });
      
      // If the app is already installed, the beforeinstallprompt event won't fire
      window.addEventListener('appinstalled', () => {
        window.__installPromptEvent = null;
        document.dispatchEvent(new CustomEvent('pwa:installed'));
        resolve(false);
      });
    } else {
      resolve(false);
    }
  });
}

// Function to show the install prompt
export function showInstallPrompt() {
  if (window.__installPromptEvent) {
    window.__installPromptEvent.prompt();
    
    // Wait for the user to respond to the prompt
    window.__installPromptEvent.userChoice.then((choiceResult: { outcome: string }) => {
      if (choiceResult.outcome === 'accepted') {
        console.log('User accepted the install prompt');
        document.dispatchEvent(new CustomEvent('pwa:installed'));
      } else {
        console.log('User dismissed the install prompt');
      }
      
      // Clear the install prompt event
      window.__installPromptEvent = null;
    });
  }
}

// Initialize PWA functionality
export function initPWA() {
  registerServiceWorker();
  checkInstallability();
  
  // Add global types
  window.__pwaUpdating = false;
  window.__installPromptEvent = null;
}

// Declare global types
declare global {
  interface Window {
    __pwaUpdating: boolean;
    __installPromptEvent: any;
  }
}

export default initPWA;