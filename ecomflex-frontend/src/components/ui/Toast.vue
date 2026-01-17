<template>
  <Teleport to="body">
    <div 
      class="fixed top-4 right-4 z-50 flex flex-col gap-2 w-full max-w-sm pointer-events-none"
      aria-live="assertive"
      aria-atomic="true"
    >
      <TransitionGroup 
        name="toast"
        tag="div"
        class="flex flex-col gap-2"
      >
        <div
          v-for="toast in toasts"
          :key="toast.id"
          class="pointer-events-auto w-full overflow-hidden rounded-lg shadow-lg ring-1 bg-white dark:bg-gray-800 transform transition-all duration-300"
          :class="getToastClasses(toast.type)"
          role="alert"
          @mouseenter="pauseAutoClose(toast.id)"
          @mouseleave="resumeAutoClose(toast.id)"
        >
          <div class="p-4 flex items-start">
            <div class="flex-shrink-0">
              <component 
                :is="getIcon(toast.type)" 
                class="h-5 w-5" 
                :class="getIconClasses(toast.type)"
              />
            </div>
            <div class="ml-3 w-0 flex-1">
              <p v-if="toast.title" class="text-sm font-medium text-gray-900 dark:text-white">
                {{ toast.title }}
              </p>
              <p class="text-sm text-gray-500 dark:text-gray-300">
                {{ toast.message }}
              </p>
            </div>
            <div class="ml-4 flex-shrink-0 flex">
              <button
                type="button"
                class="inline-flex rounded-md text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2"
                @click="closeToast(toast.id)"
              >
                <span class="sr-only">Close</span>
                <XMarkIcon class="h-5 w-5" />
              </button>
            </div>
          </div>
          <div 
            v-if="toast.autoClose" 
            class="h-1 bg-gray-200 dark:bg-gray-700"
          >
            <div 
              class="h-full transition-all duration-300 ease-linear"
              :class="getProgressClasses(toast.type)"
              :style="{ width: `${getProgress(toast.id)}%` }"
            ></div>
          </div>
        </div>
      </TransitionGroup>
    </div>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { useNotificationStore } from '@/stores/notification'
import { 
  CheckCircleIcon, 
  ExclamationCircleIcon, 
  ExclamationTriangleIcon,
  InformationCircleIcon,
  XMarkIcon 
} from '@heroicons/vue/24/solid'

const notificationStore = useNotificationStore()
const toasts = computed(() => notificationStore.toasts)

// Progress tracking for auto-closing toasts
const progressMap = ref(new Map())
const timeoutMap = ref(new Map())

// Reset progress when toast is added
onMounted(() => {
  toasts.value.forEach(toast => {
    if (toast.autoClose) {
      progressMap.value.set(toast.id, 100)
      startProgressTimer(toast.id, toast.duration || 5000)
    }
  })
})

// Clear all timeouts when component is unmounted
onUnmounted(() => {
  timeoutMap.value.forEach(timeout => {
    clearInterval(timeout)
  })
})

// Start progress timer for auto-closing toasts
const startProgressTimer = (id: string, duration: number) => {
  const interval = 10 // Update every 10ms for smooth animation
  const totalSteps = duration / interval
  const decrementAmount = 100 / totalSteps
  
  progressMap.value.set(id, 100)
  
  const timer = setInterval(() => {
    const currentProgress = progressMap.value.get(id) || 0
    const newProgress = Math.max(0, currentProgress - decrementAmount)
    
    progressMap.value.set(id, newProgress)
    
    if (newProgress <= 0) {
      clearInterval(timer)
      closeToast(id)
    }
  }, interval)
  
  timeoutMap.value.set(id, timer)
  
  return timer
}

// Get current progress for a toast
const getProgress = (id: string) => {
  return progressMap.value.get(id) || 100
}

// Pause auto-close when hovering
const pauseAutoClose = (id: string) => {
  const timer = timeoutMap.value.get(id)
  if (timer) {
    clearInterval(timer)
  }
}

// Resume auto-close when no longer hovering
const resumeAutoClose = (id: string) => {
  const toast = toasts.value.find(t => t.id === id)
  if (toast?.autoClose) {
    const remainingProgress = progressMap.value.get(id) || 0
    const originalDuration = toast.duration || 5000
    const remainingDuration = (remainingProgress / 100) * originalDuration
    
    if (remainingDuration > 0) {
      const timer = startProgressTimer(id, remainingDuration)
      timeoutMap.value.set(id, timer)
    }
  }
}

// Close a toast
const closeToast = (id: string) => {
  const timer = timeoutMap.value.get(id)
  if (timer) {
    clearInterval(timer)
    timeoutMap.value.delete(id)
  }
  progressMap.value.delete(id)
  notificationStore.removeToast(id)
}

// Get icon component based on toast type
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

// Get icon color classes based on toast type
const getIconClasses = (type: string) => {
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

// Get toast ring classes based on type
const getToastClasses = (type: string) => {
  switch (type) {
    case 'success':
      return 'ring-green-500/10'
    case 'error':
      return 'ring-red-500/10'
    case 'warning':
      return 'ring-yellow-500/10'
    case 'info':
    default:
      return 'ring-accent/10'
  }
}

// Get progress bar color based on toast type
const getProgressClasses = (type: string) => {
  switch (type) {
    case 'success':
      return 'bg-green-500'
    case 'error':
      return 'bg-red-500'
    case 'warning':
      return 'bg-yellow-500'
    case 'info':
    default:
      return 'bg-accent'
  }
}
</script>

<style scoped>
.toast-enter-active,
.toast-leave-active {
  transition: all 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

.toast-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.toast-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
</style>