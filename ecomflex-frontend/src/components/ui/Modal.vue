<template>
  <Teleport to="body">
    <transition
      name="modal-fade"
      @before-enter="beforeEnter"
      @after-enter="afterEnter"
      @before-leave="beforeLeave"
      @after-leave="afterLeave"
    >
      <div 
        v-if="modelValue" 
        class="fixed inset-0 z-50 overflow-y-auto"
        aria-labelledby="modal-title"
        role="dialog" 
        aria-modal="true"
        @keydown.esc="closeModal"
        ref="modalEl"
      >
        <!-- Backdrop -->
        <div 
          class="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity"
          @click="closeOnBackdrop ? closeModal() : null"
          :class="{ 'backdrop-blur-sm': blur }"
        ></div>

        <!-- Modal Container -->
        <div 
          class="flex min-h-full items-center justify-center p-4 text-center sm:p-0"
          :class="{ 'sm:items-end': position === 'bottom', 'sm:items-start': position === 'top' }"
        >
          <!-- Modal Content -->
          <div 
            ref="modalContent"
            class="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 w-full dark:bg-gray-800"
            :class="[
              sizeClasses,
              position === 'bottom' ? 'sm:rounded-b-none' : '',
              position === 'top' ? 'sm:rounded-t-none' : '',
              roundedClasses,
              withGoldenAccent ? 'ring-2 ring-accent/50' : ''
            ]"
            @click.stop
          >
            <!-- Close Button (if showClose is true) -->
            <button
              v-if="showClose"
              type="button"
              class="absolute top-3 right-3 rounded-md bg-white text-gray-400 hover:text-gray-500 focus:outline-none focus:ring-2 focus:ring-accent dark:bg-gray-800 dark:text-gray-500 dark:hover:text-gray-400"
              @click="closeModal"
              data-no-ripple
            >
              <span class="sr-only">Close</span>
              <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
              </svg>
            </button>

            <!-- Header Slot -->
            <div v-if="$slots.header || title" :class="['px-6 py-4', headerClass]">
              <slot name="header">
                <h3 
                  class="text-lg font-medium leading-6 text-gray-900 dark:text-white" 
                  id="modal-title"
                >
                  {{ title }}
                </h3>
              </slot>
            </div>

            <!-- Content Slot -->
            <div :class="['px-6 py-4', bodyClass]">
              <slot></slot>
            </div>

            <!-- Footer Slot -->
            <div 
              v-if="$slots.footer || showFooter" 
              :class="[
                'px-6 py-4 flex justify-end space-x-3 border-t border-gray-200 dark:border-gray-700', 
                footerClass
              ]"
            >
              <slot name="footer">
                <button
                  v-if="showCancel"
                  type="button"
                  class="inline-flex justify-center rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-medium text-gray-700 shadow-sm hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 dark:bg-gray-700 dark:border-gray-600 dark:text-gray-200 dark:hover:bg-gray-600"
                  @click="cancelModal"
                >
                  {{ cancelText }}
                </button>
                <button
                  v-if="showConfirm"
                  type="button"
                  class="inline-flex justify-center rounded-md border border-transparent bg-accent px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2"
                  @click="confirmModal"
                  :disabled="confirmDisabled"
                >
                  {{ confirmText }}
                </button>
              </slot>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </Teleport>
</template>

<script setup lang="ts">
import { ref, watch, computed, onMounted, onUnmounted, nextTick } from 'vue'
import { useAnimation } from '@/composables/useAnimation'

const props = withDefaults(defineProps<{
  modelValue: boolean
  title?: string
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | 'full'
  position?: 'center' | 'top' | 'bottom'
  closeOnBackdrop?: boolean
  closeOnEsc?: boolean
  showClose?: boolean
  rounded?: 'none' | 'sm' | 'md' | 'lg' | 'xl' | '2xl' | '3xl'
  animationVariant?: 'fade' | 'zoom' | 'slide-up' | 'slide-down'
  headerClass?: string
  bodyClass?: string
  footerClass?: string
  showFooter?: boolean
  showCancel?: boolean
  showConfirm?: boolean
  cancelText?: string
  confirmText?: string
  confirmDisabled?: boolean
  withGoldenAccent?: boolean
  blur?: boolean
}>(), {
  title: '',
  size: 'md',
  position: 'center',
  closeOnBackdrop: true,
  closeOnEsc: true,
  showClose: true,
  rounded: 'lg',
  animationVariant: 'zoom',
  headerClass: '',
  bodyClass: '',
  footerClass: '',
  showFooter: false,
  showCancel: true,
  showConfirm: true,
  cancelText: 'Cancel',
  confirmText: 'Confirm',
  confirmDisabled: false,
  withGoldenAccent: false,
  blur: false
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'confirm'): void
  (e: 'cancel'): void
  (e: 'opened'): void
  (e: 'closed'): void
}>()

// Refs
const modalEl = ref<HTMLElement | null>(null)
const modalContent = ref<HTMLElement | null>(null)
const { scale } = useAnimation()

// Computed classes based on size prop
const sizeClasses = computed(() => {
  switch (props.size) {
    case 'xs':
      return 'sm:max-w-xs'
    case 'sm':
      return 'sm:max-w-sm'
    case 'md':
      return 'sm:max-w-md'
    case 'lg':
      return 'sm:max-w-lg'
    case 'xl':
      return 'sm:max-w-xl'
    case 'full':
      return 'sm:max-w-full sm:m-4'
    default:
      return 'sm:max-w-md'
  }
})

// Computed classes based on rounded prop
const roundedClasses = computed(() => {
  switch (props.rounded) {
    case 'none':
      return 'rounded-none'
    case 'sm':
      return 'rounded-sm'
    case 'md':
      return 'rounded-md'
    case 'lg':
      return 'rounded-lg'
    case 'xl':
      return 'rounded-xl'
    case '2xl':
      return 'rounded-2xl'
    case '3xl':
      return 'rounded-3xl'
    default:
      return 'rounded-lg'
  }
})

// Close the modal
const closeModal = () => {
  emit('update:modelValue', false)
}

// Confirm button handler
const confirmModal = () => {
  emit('confirm')
  if (!props.confirmDisabled) {
    closeModal()
  }
}

// Cancel button handler
const cancelModal = () => {
  emit('cancel')
  closeModal()
}

// Handle keyboard events for accessibility
const handleKeyDown = (event: KeyboardEvent) => {
  if (props.closeOnEsc && event.key === 'Escape') {
    closeModal()
  }
}

// Focus trap and a11y
const setupFocusTrap = () => {
  if (!modalEl.value) return
  
  // Focus the first focusable element or the modal itself
  const focusableElements = modalEl.value.querySelectorAll<HTMLElement>(
    'button, [href], input, select, textarea, [tabindex]:not([tabindex="-1"])'
  )
  
  if (focusableElements.length > 0) {
    focusableElements[0].focus()
  } else {
    modalEl.value.focus()
  }
}

// Prevent body scroll when modal is open
const lockBodyScroll = () => {
  document.body.style.overflow = 'hidden'
  document.body.style.paddingRight = '15px' // Prevent layout shift
}

const unlockBodyScroll = () => {
  document.body.style.overflow = ''
  document.body.style.paddingRight = ''
}

// Lifecycle hooks
onMounted(() => {
  document.addEventListener('keydown', handleKeyDown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeyDown)
  unlockBodyScroll()
})

// Animation hooks
const beforeEnter = () => {
  lockBodyScroll()
}

const afterEnter = () => {
  setupFocusTrap()
  emit('opened')
  
  // Add entrance animation using the Animation API
  if (modalContent.value) {
    switch (props.animationVariant) {
      case 'zoom':
        scale(modalContent.value, 0.9, 1, { duration: 250, timing: 'ease-out' })
        break
      // Other animations can be added here
    }
  }
}

const beforeLeave = () => {
  // Can add exit animations here
}

const afterLeave = () => {
  unlockBodyScroll()
  emit('closed')
}

// Watch for changes to modelValue
watch(() => props.modelValue, (newVal) => {
  if (newVal) {
    nextTick(() => {
      setupFocusTrap()
    })
  }
})
</script>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.2s ease;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}

/* Custom animation for slide-up variant */
.slide-up-enter-active,
.slide-up-leave-active {
  transition: all 0.3s ease-out;
}

.slide-up-enter-from,
.slide-up-leave-to {
  opacity: 0;
  transform: translateY(30px);
}

/* Custom animation for slide-down variant */
.slide-down-enter-active,
.slide-down-leave-active {
  transition: all 0.3s ease-out;
}

.slide-down-enter-from,
.slide-down-leave-to {
  opacity: 0;
  transform: translateY(-30px);
}
</style>