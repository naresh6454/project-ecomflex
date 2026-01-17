<template>
  <button
    :class="[
      'btn',
      `btn-${variant}`,
      { 'w-full': block },
      { 'opacity-50 cursor-not-allowed': disabled },
      { 'ripple': ripple },
      { 'animate-pulse-gold': loading && variant === 'accent' },
      { 'animate-pulse': loading && variant !== 'accent' },
      sizeClasses,
      className
    ]"
    :disabled="disabled || loading"
    @click="handleClick"
  >
    <!-- Loading Spinner -->
    <span v-if="loading" class="mr-2">
      <svg class="animate-spin h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
        <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
        <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
      </svg>
    </span>
    
    <!-- Left Icon -->
    <span v-if="leftIcon && !loading" class="mr-2">
      <slot name="left-icon"></slot>
    </span>
    
    <!-- Button Content -->
    <span>
      <slot>{{ label }}</slot>
    </span>
    
    <!-- Right Icon -->
    <span v-if="rightIcon" class="ml-2">
      <slot name="right-icon"></slot>
    </span>
  </button>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  variant?: 'primary' | 'secondary' | 'accent' | 'danger' | 'link'
  size?: 'sm' | 'md' | 'lg'
  block?: boolean
  disabled?: boolean
  loading?: boolean
  leftIcon?: boolean
  rightIcon?: boolean
  ripple?: boolean
  label?: string
  className?: string
}

const props = withDefaults(defineProps<Props>(), {
  variant: 'primary',
  size: 'md',
  block: false,
  disabled: false,
  loading: false,
  leftIcon: false,
  rightIcon: false,
  ripple: true,
  label: '',
  className: ''
})

const emit = defineEmits<{
  (e: 'click', event: MouseEvent): void
}>()

// Calculate size classes based on the size prop
const sizeClasses = computed(() => {
  switch (props.size) {
    case 'sm':
      return 'px-3 py-1.5 text-sm'
    case 'lg':
      return 'px-6 py-3 text-lg'
    default: // md
      return 'px-4 py-2'
  }
})

// Handle button click
const handleClick = (event: MouseEvent) => {
  if (!props.disabled && !props.loading) {
    emit('click', event)
  }
}
</script>