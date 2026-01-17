<template>
  <router-view v-slot="{ Component, route }">
    <transition
      :name="transitionName"
      :mode="mode"
      :duration="duration"
      @before-leave="beforeLeave"
      @enter="enter"
      @after-enter="afterEnter"
    >
      <component
        :is="Component"
        :key="getTransitionKey(route)"
        v-if="Component"
        class="page-wrapper"
      />
    </transition>
  </router-view>
</template>

<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'

// Props with defaults
const props = withDefaults(defineProps<{
  name?: string
  mode?: 'in-out' | 'out-in' | 'default'
  duration?: number | { enter: number, leave: number }
  alwaysRender?: boolean
  useCustomTransition?: boolean
  noTransition?: boolean
}>(), {
  name: 'page',
  mode: 'out-in',
  duration: undefined,
  alwaysRender: false,
  useCustomTransition: false,
  noTransition: false
})

// Setup route and transition state
const route = useRoute()
const previousRoutePath = ref('')
const transitionDirection = ref('forward')
const hasInitialized = ref(false)

// Custom transition name based on navigation direction
const transitionName = computed(() => {
  if (props.noTransition) return '';
  
  if (props.useCustomTransition) {
    // Use direction-based transitions
    if (hasInitialized.value) {
      return transitionDirection.value === 'forward' ? 'slide-left' : 'slide-right'
    } else {
      return 'fade' // Use fade for initial page load
    }
  }
  
  return props.name
})

// Detect navigation direction based on route depth/path
watch(() => route.path, (newPath, oldPath) => {
  if (oldPath && newPath) {
    hasInitialized.value = true
    
    // Determine direction based on route path
    const newDepth = newPath.split('/').filter(Boolean).length
    const oldDepth = oldPath.split('/').filter(Boolean).length
    
    if (newDepth > oldDepth) {
      transitionDirection.value = 'forward'
    } else if (newDepth < oldDepth) {
      transitionDirection.value = 'backward'
    } else {
      // Same depth, compare path segments to determine direction
      const newSegments = newPath.split('/').filter(Boolean)
      const oldSegments = oldPath.split('/').filter(Boolean)
      
      // Find the first different segment
      for (let i = 0; i < Math.min(newSegments.length, oldSegments.length); i++) {
        if (newSegments[i] !== oldSegments[i]) {
          // Try to determine if it's a "forward" or "backward" navigation
          // This is a simplification - in a real app you might use more complex logic
          transitionDirection.value = 'forward'
          break
        }
      }
    }
    
    previousRoutePath.value = oldPath
  }
})

// Generate unique key for transition based on route
const getTransitionKey = (route: any) => {
  if (props.alwaysRender) {
    // Force re-render on every navigation
    return route.fullPath
  }
  
  // For nested routes, use the fullPath to ensure unique transitions
  return route.name ? route.name.toString() : route.fullPath
}

// Transition hooks for custom animations
const beforeLeave = (_el: Element) => {
  // Save scroll position or other state if needed
  window.scrollTo(0, 0)
}

const enter = (_el: Element) => {
  // Entrance animations could be implemented here
  // console.log('Component entering', _el)
}

const afterEnter = (_el: Element) => {
  // Cleanup or post-transition actions could be implemented here
  // console.log('Component entered', _el)
}
</script>

<style scoped>
/* Define page transition animations */
.page-enter-active,
.page-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
}

.page-enter-from {
  opacity: 0;
  transform: translateX(20px);
}

.page-leave-to {
  opacity: 0;
  transform: translateX(-20px);
}

/* Slide transitions */
.slide-left-enter-active,
.slide-left-leave-active,
.slide-right-enter-active,
.slide-right-leave-active {
  transition: opacity 0.3s ease, transform 0.3s ease;
  position: absolute;
  width: 100%;
}

.slide-left-enter-from {
  opacity: 0;
  transform: translateX(30px);
}

.slide-left-leave-to {
  opacity: 0;
  transform: translateX(-30px);
}

.slide-right-enter-from {
  opacity: 0;
  transform: translateX(-30px);
}

.slide-right-leave-to {
  opacity: 0;
  transform: translateX(30px);
}

/* Fade transition */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

/* Page wrapper to ensure proper layout during transitions */
.page-wrapper {
  width: 100%;
  min-height: 100%;
}
</style>