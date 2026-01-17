import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export type NotificationType = 'success' | 'error' | 'info' | 'warning'

export interface Notification {
  id: string
  type: NotificationType
  message: string
  title?: string
  autoClose?: boolean
  duration?: number
  createdAt: Date
  read: boolean
}

export const useNotificationStore = defineStore('notification', () => {
  // State
  const notifications = ref<Notification[]>([])
  const toasts = ref<Notification[]>([])
  
  // Getters
  const unreadCount = computed(() => {
    return notifications.value.filter(notification => !notification.read).length
  })
  
  const hasUnread = computed(() => unreadCount.value > 0)
  
  // Actions
  const addNotification = (notification: Omit<Notification, 'id' | 'createdAt' | 'read'>) => {
    const id = `notification-${Date.now()}-${Math.random().toString(36).substring(2, 9)}`
    const newNotification: Notification = {
      id,
      createdAt: new Date(),
      read: false,
      autoClose: false,
      ...notification
    }
    
    notifications.value.unshift(newNotification)
    
    // Keep only the latest 50 notifications
    if (notifications.value.length > 50) {
      notifications.value = notifications.value.slice(0, 50)
    }
    
    return id
  }
  
  const showToast = (notification: Omit<Notification, 'id' | 'createdAt' | 'read'>) => {
    const id = `toast-${Date.now()}-${Math.random().toString(36).substring(2, 9)}`
    const newToast: Notification = {
      id,
      createdAt: new Date(),
      read: false,
      autoClose: true,
      duration: 5000, // Default 5 seconds
      ...notification
    }
    
    toasts.value.push(newToast)
    
    if (newToast.autoClose) {
      setTimeout(() => {
        removeToast(id)
      }, newToast.duration)
    }
    
    // Also add to notifications list
    if (notification.type !== 'info') {
      addNotification({
        type: notification.type,
        message: notification.message,
        title: notification.title
      })
    }
    
    return id
  }
  
  const removeToast = (id: string) => {
    const index = toasts.value.findIndex(toast => toast.id === id)
    if (index !== -1) {
      toasts.value.splice(index, 1)
    }
  }
  
  const markAsRead = (id: string) => {
    const notification = notifications.value.find(n => n.id === id)
    if (notification) {
      notification.read = true
    }
  }
  
  const markAllAsRead = () => {
    notifications.value.forEach(notification => {
      notification.read = true
    })
  }
  
  const clearNotifications = () => {
    notifications.value = []
  }
  
  return {
    // State
    notifications,
    toasts,
    
    // Getters
    unreadCount,
    hasUnread,
    
    // Actions
    addNotification,
    showToast,
    removeToast,
    markAsRead,
    markAllAsRead,
    clearNotifications
  }
})