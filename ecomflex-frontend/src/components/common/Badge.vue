<template>
  <span 
    class="badge"
    :class="[
      variantClass,
      sizeClass,
      { 'animate-pulse': pulse }
    ]"
  >
    <slot></slot>
  </span>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  variant?: 'blue' | 'green' | 'yellow' | 'red' | 'gold' | 'gray'
  size?: 'sm' | 'md' | 'lg'
  pulse?: boolean
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'blue',
  size: 'md',
  pulse: false
})

// Compute variant classes
const variantClass = computed(() => {
  switch (props.variant) {
    case 'blue':
      return 'badge-blue'
    case 'green':
      return 'badge-green'
    case 'yellow':
      return 'badge-yellow'
    case 'red':
      return 'badge-red'
    case 'gold':
      return 'bg-amber-100 text-amber-800'
    case 'gray':
      return 'bg-gray-100 text-gray-800'
    default:
      return 'badge-blue'
  }
})

// Compute size classes
const sizeClass = computed(() => {
  switch (props.size) {
    case 'sm':
      return 'text-xs px-2 py-0.5'
    case 'lg':
      return 'text-sm px-3 py-1'
    default: // md
      return 'text-xs px-2.5 py-0.5'
  }
})
</script>