// src/services/pushNotificationManager.ts
// This service handles push notifications for the PWA

// Custom notification options interface to match our usage
interface CustomNotificationOptions {
  title: string;
  body: string;
  icon?: string;
  badge?: string;
  image?: string; // Not in standard NotificationOptions
  tag?: string;
  data?: any;
  actions?: NotificationAction[];
  requireInteraction?: boolean;
  renotify?: boolean;
  silent?: boolean;
  timestamp?: number;
  vibrate?: number[]; // Not in standard NotificationOptions
}

interface NotificationAction {
  action: string;
  title: string;
  icon?: string;
}

class PushNotificationManager {
  private applicationServerKey: string | null = null;
  private _subscription: PushSubscription | null = null;
  private _permissionStatus: NotificationPermission = 'default';
  private readonly SUBSCRIPTION_ENDPOINT = 'https://api.ecomflex.com/api/subscriptions/push';
  
  constructor() {
    this.checkPermission();
  }
  
  /**
   * Initialize the push notification manager
   * @param applicationServerKey VAPID public key
   */
  public async initialize(applicationServerKey: string): Promise<boolean> {
    if (!this.isPushSupported()) {
      console.warn('Push notifications are not supported in this browser');
      return false;
    }
    
    this.applicationServerKey = applicationServerKey;
    await this.checkPermission();
    
    // If we already have permission, check for existing subscription
    if (this._permissionStatus === 'granted') {
      await this.getExistingSubscription();
    }
    
    return true;
  }
  
  /**
   * Check if push notifications are supported
   */
  public isPushSupported(): boolean {
    return 'serviceWorker' in navigator && 
           'PushManager' in window && 
           'Notification' in window;
  }
  
  /**
   * Check the current notification permission status
   */
  public async checkPermission(): Promise<NotificationPermission> {
    if (!('Notification' in window)) {
      this._permissionStatus = 'denied';
      return this._permissionStatus;
    }
    
    this._permissionStatus = Notification.permission;
    return this._permissionStatus;
  }
  
  /**
   * Request permission for push notifications
   */
  public async requestPermission(): Promise<NotificationPermission> {
    if (!('Notification' in window)) {
      this._permissionStatus = 'denied';
      return this._permissionStatus;
    }
    
    try {
      this._permissionStatus = await Notification.requestPermission();
      
      // If permission granted, subscribe to push notifications
      if (this._permissionStatus === 'granted' && this.applicationServerKey) {
        await this.subscribeToPush();
      }
      
      return this._permissionStatus;
    } catch (error) {
      console.error('Error requesting permission:', error);
      return 'denied';
    }
  }
  
  /**
   * Get an existing push subscription if available
   */
  private async getExistingSubscription(): Promise<PushSubscription | null> {
    if (!this.isPushSupported()) {
      return null;
    }
    
    try {
      const registration = await navigator.serviceWorker.ready;
      const subscription = await registration.pushManager.getSubscription();
      
      this._subscription = subscription;
      return subscription;
    } catch (error) {
      console.error('Error getting existing subscription:', error);
      return null;
    }
  }
  
  /**
   * Subscribe to push notifications
   */
  public async subscribeToPush(): Promise<PushSubscription | null> {
    if (!this.isPushSupported() || !this.applicationServerKey) {
      return null;
    }
    
    if (this._permissionStatus !== 'granted') {
      const permission = await this.requestPermission();
      if (permission !== 'granted') {
        return null;
      }
    }
    
    try {
      const registration = await navigator.serviceWorker.ready;
      
      // Check for existing subscription
      let subscription = await registration.pushManager.getSubscription();
      
      // If already subscribed, return the subscription
      if (subscription) {
        this._subscription = subscription;
        return subscription;
      }
      
      // Convert the application server key to Uint8Array
      const key = this.urlBase64ToUint8Array(this.applicationServerKey);
      
      // Use a more compatible way to handle the key by forcing it to be a BufferSource
      const applicationServerKey = key.buffer as BufferSource;
      
      // Subscribe
      subscription = await registration.pushManager.subscribe({
        userVisibleOnly: true,
        applicationServerKey
      });
      
      this._subscription = subscription;
      
      // Send the subscription to the server
      await this.sendSubscriptionToServer(subscription);
      
      return subscription;
    } catch (error) {
      console.error('Error subscribing to push notifications:', error);
      return null;
    }
  }
  
  /**
   * Unsubscribe from push notifications
   */
  public async unsubscribeFromPush(): Promise<boolean> {
    if (!this.isPushSupported()) {
      return false;
    }
    
    try {
      const registration = await navigator.serviceWorker.ready;
      const subscription = await registration.pushManager.getSubscription();
      
      if (subscription) {
        const result = await subscription.unsubscribe();
        
        if (result) {
          // Remove the subscription from the server
          await this.removeSubscriptionFromServer(subscription);
          this._subscription = null;
        }
        
        return result;
      }
      
      return false;
    } catch (error) {
      console.error('Error unsubscribing from push notifications:', error);
      return false;
    }
  }
  
  /**
   * Send the subscription to the server
   * @param subscription The push subscription
   */
  private async sendSubscriptionToServer(subscription: PushSubscription): Promise<boolean> {
    try {
      const userId = localStorage.getItem('userId') || 'anonymous';
      
      const response = await fetch(this.SUBSCRIPTION_ENDPOINT, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('token') || ''}`
        },
        body: JSON.stringify({
          subscription,
          userId,
          platform: 'web',
          timestamp: Date.now()
        })
      });
      
      return response.ok;
    } catch (error) {
      console.error('Error sending subscription to server:', error);
      return false;
    }
  }
  
  /**
   * Remove the subscription from the server
   * @param subscription The push subscription
   */
  private async removeSubscriptionFromServer(subscription: PushSubscription): Promise<boolean> {
    try {
      const userId = localStorage.getItem('userId') || 'anonymous';
      const endpoint = subscription.endpoint;
      
      const response = await fetch(`${this.SUBSCRIPTION_ENDPOINT}/${encodeURIComponent(endpoint)}`, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('token') || ''}`
        },
        body: JSON.stringify({
          userId,
          platform: 'web'
        })
      });
      
      return response.ok;
    } catch (error) {
      console.error('Error removing subscription from server:', error);
      return false;
    }
  }
  
  /**
   * Show a local notification
   * @param options Notification options
   */
  public async showNotification(options: CustomNotificationOptions): Promise<boolean> {
    if (!this.isPushSupported() || this._permissionStatus !== 'granted') {
      return false;
    }
    
    try {
      const registration = await navigator.serviceWorker.ready;
      
      // Extract non-standard properties from the options
      const { image, vibrate, ...standardOptions } = options;
      
      // Create the notification with standard options
      await registration.showNotification(options.title, {
        ...standardOptions,
        icon: options.icon || '/icons/icon-192x192.png',
        badge: options.badge || '/icons/badge-96x96.png',
      });
      
      return true;
    } catch (error) {
      console.error('Error showing notification:', error);
      return false;
    }
  }
  
  /**
   * Close all notifications
   */
  public async closeAllNotifications(): Promise<void> {
    if (!('Notification' in window)) {
      return;
    }
    
    const notifications = await this.getActiveNotifications();
    notifications.forEach(notification => notification.close());
  }
  
  /**
   * Get all active notifications
   */
  public async getActiveNotifications(): Promise<Notification[]> {
    if (!this.isPushSupported()) {
      return [];
    }
    
    try {
      const registration = await navigator.serviceWorker.ready;
      
      if (registration.getNotifications) {
        return await registration.getNotifications();
      }
      
      return [];
    } catch (error) {
      console.error('Error getting active notifications:', error);
      return [];
    }
  }
  
  /**
   * Get the current subscription
   */
  public get subscription(): PushSubscription | null {
    return this._subscription;
  }
  
  /**
   * Get the permission status
   */
  public get permissionStatus(): NotificationPermission {
    return this._permissionStatus;
  }
  
  /**
   * Convert a URL-safe base64 string to a Uint8Array
   * @param base64String URL-safe base64 string
   */
  private urlBase64ToUint8Array(base64String: string): Uint8Array {
    const padding = '='.repeat((4 - (base64String.length % 4)) % 4);
    const base64 = (base64String + padding)
      .replace(/-/g, '+')
      .replace(/_/g, '/');
    
    const rawData = window.atob(base64);
    const outputArray = new Uint8Array(rawData.length);
    
    for (let i = 0; i < rawData.length; ++i) {
      outputArray[i] = rawData.charCodeAt(i);
    }
    
    return outputArray;
  }
}

// Create and export a singleton instance
const pushNotificationManager = new PushNotificationManager();
export default pushNotificationManager;