<template>
  <div class="relative">
    <!-- Notification Bell Button -->
    <button
      type="button"
      class="relative rounded-full p-1 text-gray-600 hover:text-gray-800 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 dark:text-gray-300 dark:hover:text-white"
      @click="toggleDropdown"
      ref="notificationButton"
      aria-haspopup="true"
      :aria-expanded="isOpen"
    >
      <span class="sr-only">View notifications</span>
      <BellIcon class="h-6 w-6" />
      
      <!-- Notification Badge -->
      <span
        v-if="hasUnread"
        class="absolute top-0 right-0 block h-2.5 w-2.5 rounded-full bg-accent ring-2 ring-white dark:ring-gray-900"
        aria-hidden="true"
      ></span>
    </button>

    <!-- Notification Dropdown -->
    <transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="transform opacity-0 scale-95"
      enter-to-class="transform opacity-100 scale-100"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="transform opacity-100 scale-100"
      leave-to-class="transform opacity-0 scale-95"
    >
      <div
        v-if="isOpen"
        class="absolute right-0 z-50 mt-2 w-80 origin-top-right rounded-lg bg-white py-1 shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none dark:bg-gray-800 dark:ring-gray-700"
        role="menu"
        aria-orientation="vertical"
        :aria-labelledby="notificationButton"
        ref="dropdown"
      >
        <div class="border-b border-gray-200 px-4 py-2 dark:border-gray-700">
          <div class="flex items-center justify-between">
            <h3 class="text-sm font-medium text-gray-900 dark:text-white">Notifications</h3>
            <button
              v-if="unreadCount > 0"
              class="text-xs font-medium text-accent hover:text-accent-dark focus:outline-none"
              @click="markAllAsRead"
            >
              Mark all as read
            </button>
          </div>
        </div>

        <div class="max-h-80 overflow-y-auto py-1" role="none">
          <div v-if="notifications.length === 0" class="px-4 py-6 text-center">
            <BellSlashIcon class="mx-auto h-8 w-8 text-gray-400" />
            <p class="mt-2 text-sm text-gray-500 dark:text-gray-400">
              No notifications yet
            </p>
          </div>

          <transition-group name="notification-list">
            <div
              v-for="notification in notifications"
              :key="notification.id"
              class="relative px-4 py-3 hover:bg-gray-50 dark:hover:bg-gray-700"
              :class="{ 'bg-accent/5': !notification.read }"
              @click="handleNotificationClick(notification)"
            >
              <div class="flex items-start">
                <div class="flex-shrink-0">
                  <component
                    :is="getIcon(notification.type)"
                    class="h-5 w-5"
                    :class="getIconClass(notification.type)"
                  />
                </div>
                <div class="ml-3 w-0 flex-1">
                  <p
                    v-if="notification.title"
                    class="text-sm font-medium text-gray-900 dark:text-white"
                  >
                    {{ notification.title }}
                  </p>
                  <p
                    class="text-sm text-gray-500 dark:text-gray-400"
                    :class="{ 'font-semibold': !notification.read }"
                  >
                    {{ notification.message }}
                  </p>
                  <p class="mt-1 text-xs text-gray-400 dark:text-gray-500">
                    {{ formatDate(notification.createdAt) }}
                  </p>
                </div>
                <div class="ml-4 flex flex-shrink-0">
                  <button
                    type="button"
                    class="inline-flex rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 dark:bg-transparent dark:hover:text-gray-300"
                    @click.stop="removeNotification(notification.id)"
                  >
                    <span class="sr-only">Remove notification</span>
                    <XMarkIcon class="h-4 w-4" />
                  </button>
                </div>
              </div>
              
              <!-- Unread indicator dot -->
              <span
                v-if="!notification.read"
                class="absolute left-0 top-1/2 h-1.5 w-1.5 -translate-x-1/2 -translate-y-1/2 transform rounded-full bg-accent"
                aria-hidden="true"
              ></span>
            </div>
          </transition-group>
        </div>

        <div class="border-t border-gray-200 px-4 py-2 text-center dark:border-gray-700">
          <button
            v-if="notifications.length > 0"
            class="text-xs font-medium text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-300"
            @click="clearAllNotifications"
          >
            Clear all notifications
          </button>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useNotificationStore } from '@/stores/notification'
import {
  BellIcon,
  BellSlashIcon,
  CheckCircleIcon,
  ExclamationCircleIcon,
  ExclamationTriangleIcon,
  InformationCircleIcon,
  XMarkIcon
} from '@heroicons/vue/24/outline'

const notificationStore = useNotificationStore()

// State
const isOpen = ref(false)
const notificationButton = ref<HTMLElement | null>(null)
const dropdown = ref<HTMLElement | null>(null)

// Computed properties
const notifications = computed(() => notificationStore.notifications)
const hasUnread = computed(() => notificationStore.hasUnread)
const unreadCount = computed(() => notificationStore.unreadCount)

// Methods
const toggleDropdown = () => {
  isOpen.value = !isOpen.value
}

const markAllAsRead = () => {
  notificationStore.markAllAsRead()
}

const handleNotificationClick = (notification: any) => {
  if (!notification.read) {
    notificationStore.markAsRead(notification.id)
  }
  // Additional handling (routing, etc.) can be added here
}

const removeNotification = (id: string) => {
  const notification = notifications.value.find(n => n.id === id)
  if (notification) {
    notificationStore.markAsRead(id) // Ensure it's marked as read
    
    // Remove from the list (this would need to be added to the store)
    const index = notifications.value.findIndex(n => n.id === id)
    if (index !== -1) {
      notifications.value.splice(index, 1)
    }
  }
}

const clearAllNotifications = () => {
  notificationStore.clearNotifications()
  isOpen.value = false
}

// Format notification date
const formatDate = (date: Date) => {
  if (!(date instanceof Date)) {
    date = new Date(date)
  }
  
  const now = new Date()
  const diffMs = now.getTime() - date.getTime()
  const diffSecs = Math.floor(diffMs / 1000)
  const diffMins = Math.floor(diffSecs / 60)
  const diffHours = Math.floor(diffMins / 60)
  const diffDays = Math.floor(diffHours / 24)
  
  if (diffSecs < 60) {
    return 'Just now'
  } else if (diffMins < 60) {
    return `${diffMins}m ago`
  } else if (diffHours < 24) {
    return `${diffHours}h ago`
  } else if (diffDays < 7) {
    return `${diffDays}d ago`
  } else {
    return date.toLocaleDateString()
  }
}

// Get appropriate icon for notification type
const getIcon = (type: string) => {
  switch (type) {
    case 'success':
      return CheckCircleIcon
    case 'error':
      return ExclamationCircleIcon
    case 'warning':
      return ExclamationTriangleIcon
    case 'info':
    default:
      return InformationCircleIcon
  }
}

// Get icon color class based on notification type
const getIconClass = (type: string) => {
  switch (type) {
    case 'success':
      return 'text-green-500'
    case 'error':
      return 'text-red-500'
    case 'warning':
      return 'text-yellow-500'
    case 'info':
    default:
      return 'text-accent'
  }
}

// Close dropdown when clicking outside
const handleClickOutside = (event: MouseEvent) => {
  if (
    isOpen.value &&
    notificationButton.value &&
    dropdown.value &&
    !notificationButton.value.contains(event.target as Node) &&
    !dropdown.value.contains(event.target as Node)
  ) {
    isOpen.value = false
  }
}

// Close dropdown on escape key
const handleEscapeKey = (event: KeyboardEvent) => {
  if (isOpen.value && event.key === 'Escape') {
    isOpen.value = false
  }
}

// Lifecycle hooks
onMounted(() => {
  document.addEventListener('mousedown', handleClickOutside)
  document.addEventListener('keydown', handleEscapeKey)
})

onUnmounted(() => {
  document.removeEventListener('mousedown', handleClickOutside)
  document.removeEventListener('keydown', handleEscapeKey)
})
</script>

<style scoped>
.notification-list-enter-active,
.notification-list-leave-active {
  transition: all 0.3s ease;
}
.notification-list-enter-from,
.notification-list-leave-to {
  opacity: 0;
  transform: translateY(-10px);
}
</style>