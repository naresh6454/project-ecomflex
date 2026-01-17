<template>
  <component
    :is="tag"
    :class="[wrapperClass, { 'overflow-hidden': !overflowDuringTransition }]"
  >
    <transition
      :name="name"
      :mode="mode"
      :appear="appear"
      :css="css"
      :duration="duration"
      v-bind="$attrs"
      @before-enter="onBeforeEnter"
      @enter="onEnter"
      @after-enter="onAfterEnter"
      @enter-cancelled="onEnterCancelled"
      @before-leave="onBeforeLeave"
      @leave="onLeave"
      @after-leave="onAfterLeave"
      @leave-cancelled="onLeaveCancelled"
    >
      <slot></slot>
    </transition>
  </component>
</template>

<script setup lang="ts">
import { onMounted } from 'vue'

const props = withDefaults(defineProps<{
  name?: string
  mode?: 'in-out' | 'out-in' | 'default'
  appear?: boolean
  css?: boolean
  duration?: number | { enter: number, leave: number }
  tag?: string
  wrapperClass?: string
  overflowDuringTransition?: boolean
  stagger?: number
  staggerChildren?: boolean
}>(), {
  name: 'fade',
  mode: 'out-in',
  appear: false,
  css: true,
  duration: undefined,
  tag: 'div',
  wrapperClass: '',
  overflowDuringTransition: false,
  stagger: 50,
  staggerChildren: false
})

const emit = defineEmits([
  'before-enter',
  'enter',
  'after-enter',
  'enter-cancelled',
  'before-leave',
  'leave',
  'after-leave',
  'leave-cancelled'
])

// Define transition hook handlers
const onBeforeEnter = (el: Element) => {
  emit('before-enter', el)
  
  // Apply stagger to children if enabled
  if (props.staggerChildren) {
    const children = Array.from(el.children)
    children.forEach((child, index) => {
      if (child instanceof HTMLElement) {
        child.style.transitionDelay = `${index * props.stagger}ms`
        child.style.opacity = '0'
        child.style.transform = 'translateY(10px)'
      }
    })
  }
}

const onEnter = (el: Element, done: () => void) => {
  emit('enter', el, done)
  
  // Apply stagger to children if enabled
  if (props.staggerChildren) {
    const children = Array.from(el.children)
    
    // Ensure transition properties are applied
    requestAnimationFrame(() => {
      children.forEach((child) => {
        if (child instanceof HTMLElement) {
          child.style.transition = 'opacity 0.3s ease, transform 0.3s ease'
          child.style.opacity = '1'
          child.style.transform = 'translateY(0)'
        }
      })
    })
    
    // Call done when last child's transition completes
    if (children.length > 0) {
      const lastChild = children[children.length - 1] as HTMLElement
      const transitionEndHandler = () => {
        lastChild.removeEventListener('transitionend', transitionEndHandler)
        done()
      }
      lastChild.addEventListener('transitionend', transitionEndHandler)
    } else {
      done()
    }
  } else {
    done()
  }
}

const onAfterEnter = (el: Element) => {
  emit('after-enter', el)
  
  // Clean up stagger styles
  if (props.staggerChildren) {
    const children = Array.from(el.children)
    children.forEach((child) => {
      if (child instanceof HTMLElement) {
        child.style.transitionDelay = ''
      }
    })
  }
}

const onEnterCancelled = (el: Element) => {
  emit('enter-cancelled', el)
}

const onBeforeLeave = (el: Element) => {
  emit('before-leave', el)
  
  // Apply stagger to children if enabled
  if (props.staggerChildren) {
    const children = Array.from(el.children)
    children.forEach((child, index) => {
      if (child instanceof HTMLElement) {
        // Reverse order for leave animation
        const reverseIndex = children.length - 1 - index
        child.style.transitionDelay = `${reverseIndex * props.stagger}ms`
      }
    })
  }
}

const onLeave = (el: Element, done: () => void) => {
  emit('leave', el, done)
  
  // Apply stagger to children if enabled
  if (props.staggerChildren) {
    const children = Array.from(el.children)
    
    // Ensure transition properties are applied
    requestAnimationFrame(() => {
      children.forEach((child) => {
        if (child instanceof HTMLElement) {
          child.style.transition = 'opacity 0.3s ease, transform 0.3s ease'
          child.style.opacity = '0'
          child.style.transform = 'translateY(10px)'
        }
      })
    })
    
    // Call done when last child's transition completes
    if (children.length > 0) {
      const firstChild = children[0] as HTMLElement
      const transitionEndHandler = () => {
        firstChild.removeEventListener('transitionend', transitionEndHandler)
        done()
      }
      firstChild.addEventListener('transitionend', transitionEndHandler)
    } else {
      done()
    }
  } else {
    done()
  }
}

const onAfterLeave = (el: Element) => {
  emit('after-leave', el)
  
  // Clean up stagger styles
  if (props.staggerChildren) {
    const children = Array.from(el.children)
    children.forEach((child) => {
      if (child instanceof HTMLElement) {
        child.style.transitionDelay = ''
        child.style.opacity = ''
        child.style.transform = ''
      }
    })
  }
}

const onLeaveCancelled = (el: Element) => {
  emit('leave-cancelled', el)
}

// Apply default class for transition
onMounted(() => {
  // Add any initialization if needed
})
</script>

<style scoped>
/* Basic transition classes */
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s ease-out;
}

.slide-up-enter-from,
.slide-up-leave-to {
  opacity: 0;
  transform: translateY(20px);
}

.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.3s ease-out;
}

.slide-down-enter-from,
.slide-down-leave-to {
  opacity: 0;
  transform: translateY(-20px);
}

.scale-enter-active,
.scale-leave-active {
  transition: all 0.3s ease-out;
}

.scale-enter-from,
.scale-leave-to {
  opacity: 0;
  transform: scale(0.9);
}

.scale-up-enter-active,
.scale-up-leave-active {
  transition: all 0.3s ease-out;
}

.scale-up-enter-from,
.scale-up-leave-to {
  opacity: 0;
  transform: scale(0.95) translateY(10px);
}
</style>