<template>
  <div 
    :class="[
      'animate-pulse rounded overflow-hidden',
      shapeClasses,
      heightClasses,
      widthClasses,
      colorClasses,
      radiusClasses,
      className
    ]"
  >
    <!-- Custom SVG shape pattern if specified -->
    <svg 
      v-if="pattern"
      class="w-full h-full text-gray-200 dark:text-gray-700"
      viewBox="0 0 100 100"
      preserveAspectRatio="none"
    >
      <!-- Circle pattern -->
      <template v-if="pattern === 'circle'">
        <circle cx="50" cy="50" r="45" fill="currentColor" />
      </template>
      
      <!-- Wave pattern -->
      <template v-else-if="pattern === 'wave'">
        <path
          d="M0 50 Q25 30, 50 50 T100 50 V100 H0 Z"
          fill="currentColor"
        />
      </template>
      
      <!-- Stripe pattern -->
      <template v-else-if="pattern === 'stripe'">
        <rect x="0" y="0" width="100" height="20" fill="currentColor" />
        <rect x="0" y="40" width="100" height="20" fill="currentColor" />
        <rect x="0" y="80" width="100" height="20" fill="currentColor" />
      </template>
      
      <!-- Dots pattern -->
      <template v-else-if="pattern === 'dots'">
        <circle cx="20" cy="20" r="10" fill="currentColor" />
        <circle cx="50" cy="20" r="10" fill="currentColor" />
        <circle cx="80" cy="20" r="10" fill="currentColor" />
        <circle cx="20" cy="50" r="10" fill="currentColor" />
        <circle cx="50" cy="50" r="10" fill="currentColor" />
        <circle cx="80" cy="50" r="10" fill="currentColor" />
        <circle cx="20" cy="80" r="10" fill="currentColor" />
        <circle cx="50" cy="80" r="10" fill="currentColor" />
        <circle cx="80" cy="80" r="10" fill="currentColor" />
      </template>
      
      <!-- Default line pattern -->
      <template v-else>
        <rect x="0" y="0" width="100" height="100" fill="currentColor" />
      </template>
    </svg>
    
    <!-- Shimmer effect overlay -->
    <div 
      v-if="shimmer"
      class="absolute inset-0 golden-shimmer-overlay"
    ></div>
    
    <slot></slot>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

// Props with default values
const props = withDefaults(defineProps<{
  variant?: 'text' | 'rect' | 'circle' | 'avatar' | 'button' | 'card' | 'image'
  width?: string | number
  height?: string | number
  radius?: string
  color?: 'default' | 'dark' | 'light' | 'accent' | 'golden'
  lines?: number
  pattern?: '' | 'circle' | 'wave' | 'stripe' | 'dots'
  shimmer?: boolean
  className?: string
}>(), {
  variant: 'rect',
  width: 'full',
  height: '20',
  radius: '',
  color: 'default',
  lines: 1,
  pattern: '',
  shimmer: true,
  className: ''
})

// Computed classes based on props
const shapeClasses = computed(() => {
  switch (props.variant) {
    case 'text':
      return 'h-4'
    case 'circle':
      return 'rounded-full aspect-square'
    case 'avatar':
      return 'rounded-full aspect-square'
    case 'button':
      return 'h-10 rounded-md'
    case 'card':
      return 'relative rounded-lg shadow-sm'
    case 'image':
      return 'aspect-video rounded-md'
    case 'rect':
    default:
      return 'relative'
  }
})

const heightClasses = computed(() => {
  if (props.height === 'auto') return 'h-auto'
  if (typeof props.height === 'number' || /^\d+$/.test(props.height as string)) {
    return `h-[${props.height}px]`
  }
  return `h-${props.height}`
})

const widthClasses = computed(() => {
  if (props.width === 'auto') return 'w-auto'
  if (typeof props.width === 'number' || /^\d+$/.test(props.width as string)) {
    return `w-[${props.width}px]`
  }
  return `w-${props.width}`
})

const colorClasses = computed(() => {
  switch (props.color) {
    case 'dark':
      return 'bg-gray-300 dark:bg-gray-700'
    case 'light':
      return 'bg-gray-100 dark:bg-gray-800'
    case 'accent':
      return 'bg-accent/20'
    case 'golden':
      return 'bg-gradient-to-r from-accent/10 to-accent/30'
    case 'default':
    default:
      return 'bg-gray-200 dark:bg-gray-700'
  }
})

const radiusClasses = computed(() => {
  if (!props.radius) {
    return '' // Use the default from variant
  }
  
  return `rounded-${props.radius}`
})
</script>

<style scoped>
/* Add any custom styling here */
</style>