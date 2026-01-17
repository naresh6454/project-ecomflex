<template>
  <div 
    :class="[
      'relative flex items-center justify-center',
      { 'h-full w-full': fullscreen }
    ]"
  >
    <div 
      v-if="overlay && fullscreen" 
      class="absolute inset-0 bg-gray-900/20 backdrop-blur-sm"
    ></div>
    
    <div 
      class="relative flex flex-col items-center justify-center"
      :class="{ 'z-10': overlay && fullscreen }"
    >
      <!-- Spinner type loader -->
      <div 
        v-if="type === 'spinner'"
        class="animate-spin rounded-full border-t-transparent"
        :class="[
          sizeClasses,
          `border-${thickness}`,
          colorClass
        ]"
      ></div>
      
      <!-- Pulse type loader -->
      <div
        v-else-if="type === 'pulse'"
        class="relative flex"
      >
        <div 
          v-for="(_, index) in 3" 
          :key="index"
          class="animate-pulse rounded-full mx-1"
          :class="[
            sizeClasses.replace('h-', 'h-2/3 ').replace('w-', 'w-2/3 '),
            colorClass,
            `animate-delay-${index * 100}`
          ]"
          :style="{ animationDelay: `${index * 0.15}s` }"
        ></div>
      </div>
      
      <!-- Progress type loader -->
      <div
        v-else-if="type === 'progress'"
        class="w-full max-w-md bg-gray-200 rounded-full overflow-hidden dark:bg-gray-700"
      >
        <div
          class="h-2 rounded-full transition-all duration-300"
          :class="colorClass"
          :style="{ width: `${progress}%` }"
        ></div>
      </div>
      
      <!-- Golden shimmer type loader -->
      <div
        v-else-if="type === 'shimmer'"
        class="relative overflow-hidden rounded-lg"
        :class="[sizeClasses]"
      >
        <div class="absolute inset-0 bg-gray-200 dark:bg-gray-700"></div>
        <div 
          class="absolute inset-0 animate-shimmer"
          style="background: linear-gradient(90deg, transparent, rgba(255, 215, 0, 0.4), transparent); background-size: 200% 100%;"
        ></div>
      </div>
      
      <!-- Label below loader if provided -->
      <p 
        v-if="label" 
        class="mt-3 text-sm font-medium text-gray-600 dark:text-gray-300"
      >
        {{ label }}
      </p>
      
      <!-- Progress percentage if showing progress -->
      <p 
        v-if="type === 'progress' && showPercentage" 
        class="mt-1 text-xs font-medium text-gray-500 dark:text-gray-400"
      >
        {{ Math.round(progress) }}%
      </p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'

const props = withDefaults(defineProps<{
  type?: 'spinner' | 'pulse' | 'progress' | 'shimmer'
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl'
  color?: 'accent' | 'primary' | 'success' | 'error' | 'warning' | 'info'
  thickness?: 2 | 4 | 8
  fullscreen?: boolean
  overlay?: boolean
  label?: string
  progress?: number
  showPercentage?: boolean
  indeterminate?: boolean
}>(), {
  type: 'spinner',
  size: 'md',
  color: 'accent',
  thickness: 4,
  fullscreen: false,
  overlay: false,
  progress: 0,
  showPercentage: false,
  indeterminate: false
})

// Computed size classes based on size prop
const sizeClasses = computed(() => {
  switch (props.size) {
    case 'xs':
      return 'h-4 w-4'
    case 'sm':
      return 'h-6 w-6'
    case 'md':
      return 'h-10 w-10'
    case 'lg':
      return 'h-16 w-16'
    case 'xl':
      return 'h-24 w-24'
    default:
      return 'h-10 w-10'
  }
})

// Computed color class based on color prop
const colorClass = computed(() => {
  switch (props.color) {
    case 'primary':
      return 'border-primary bg-primary'
    case 'success':
      return 'border-green-500 bg-green-500'
    case 'error':
      return 'border-red-500 bg-red-500'
    case 'warning':
      return 'border-yellow-500 bg-yellow-500'
    case 'info':
      return 'border-blue-500 bg-blue-500'
    case 'accent':
    default:
      return 'border-accent bg-accent'
  }
})

// For indeterminate progress
const indeterminateProgress = ref(0)
const indeterminateDirection = ref(1)

// Update indeterminate progress animation
if (props.indeterminate && props.type === 'progress') {
  const updateIndeterminate = () => {
    indeterminateProgress.value += indeterminateDirection.value

    if (indeterminateProgress.value >= 95) {
      indeterminateDirection.value = -1
    } else if (indeterminateProgress.value <= 5) {
      indeterminateDirection.value = 1
    }
    
    requestAnimationFrame(updateIndeterminate)
  }
  
  requestAnimationFrame(updateIndeterminate)
}

// Computed actual progress (either from props or indeterminate)
const actualProgress = computed(() => {
  return props.indeterminate ? indeterminateProgress.value : props.progress
})
</script>

<style scoped>
@keyframes shimmer {
  0% {
    background-position: -200% 0;
  }
  100% {
    background-position: 200% 0;
  }
}

.animate-shimmer {
  animation: shimmer 2s infinite;
}

.animate-delay-100 {
  animation-delay: 0.1s;
}

.animate-delay-200 {
  animation-delay: 0.2s;
}
</style>