<template>
  <Teleport to="body">
    <!-- Backdrop Overlay -->
    <transition
      enter-active-class="transition ease-out duration-200"
      enter-from-class="opacity-0"
      enter-to-class="opacity-100"
      leave-active-class="transition ease-in duration-150"
      leave-from-class="opacity-100"
      leave-to-class="opacity-0"
    >
      <div 
        v-if="modelValue" 
        class="fixed inset-0 bg-black bg-opacity-50 backdrop-blur-sm z-40 transition-opacity"
        @click="closeOnBackdrop ? close() : null"
        aria-hidden="true"
      ></div>
    </transition>

    <!-- Modal Dialog -->
    <transition
      enter-active-class="transition ease-out duration-300"
      enter-from-class="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
      enter-to-class="opacity-100 translate-y-0 sm:scale-100"
      leave-active-class="transition ease-in duration-200"
      leave-from-class="opacity-100 translate-y-0 sm:scale-100"
      leave-to-class="opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
    >
      <div 
        v-if="modelValue" 
        class="fixed inset-0 z-50 overflow-y-auto"
        role="dialog"
        aria-modal="true"
        :aria-labelledby="id + '-title'"
      >
        <div class="flex min-h-screen items-center justify-center p-4 text-center sm:p-0">
          <div 
            class="relative transform overflow-hidden rounded-2xl bg-white text-left shadow-xl transition-all sm:my-8"
            :class="sizeClass"
          >
            <!-- Close Button -->
            <button 
              v-if="showClose"
              @click="close" 
              type="button" 
              class="absolute right-2 top-2 z-10 rounded-full p-1 text-gray-400 hover:bg-gray-100 hover:text-gray-500 focus:outline-none"
              aria-label="Close"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
              </svg>
            </button>

            <!-- Modal Header -->
            <div v-if="title" class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
              <h3 
                :id="id + '-title'"
                class="text-lg font-medium leading-6 text-gray-900"
              >
                {{ title }}
              </h3>
              <div v-if="description" class="mt-2">
                <p class="text-sm text-gray-500">{{ description }}</p>
              </div>
            </div>

            <!-- Modal Content -->
            <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4" :class="{ 'pt-0': title }">
              <slot></slot>
            </div>

            <!-- Modal Footer -->
            <div v-if="$slots.footer" class="bg-gray-50 px-4 py-3 sm:flex sm:flex-row-reverse sm:px-6">
              <slot name="footer"></slot>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </Teleport>
</template>

<script setup lang="ts">
import { computed, watch, onMounted, onBeforeUnmount } from 'vue'

interface Props {
  modelValue: boolean
  title?: string
  description?: string
  size?: 'sm' | 'md' | 'lg' | 'xl' | 'full'
  closeOnBackdrop?: boolean
  showClose?: boolean
  id?: string
}

const props = withDefaults(defineProps<Props>(), {
  size: 'md',
  closeOnBackdrop: true,
  showClose: true,
  id: 'modal'
})

const emit = defineEmits<{
  (e: 'update:modelValue', value: boolean): void
  (e: 'close'): void
}>()

// Compute modal size class
const sizeClass = computed(() => {
  switch (props.size) {
    case 'sm':
      return 'sm:max-w-md w-full'
    case 'lg':
      return 'sm:max-w-2xl w-full'
    case 'xl':
      return 'sm:max-w-4xl w-full'
    case 'full':
      return 'sm:max-w-[90%] w-full sm:h-[90%]'
    default: // md
      return 'sm:max-w-lg w-full'
  }
})

// Close the modal
const close = () => {
  emit('update:modelValue', false)
  emit('close')
}

// Handle escape key press to close modal
const handleEscKey = (event: KeyboardEvent) => {
  if (event.key === 'Escape' && props.modelValue) {
    close()
  }
}

// Watch for changes in modelValue to manage body scrolling
watch(() => props.modelValue, (isOpen) => {
  if (isOpen) {
    document.body.classList.add('overflow-hidden')
  } else {
    document.body.classList.remove('overflow-hidden')
  }
})

// Lifecycle hooks for event listeners
onMounted(() => {
  document.addEventListener('keydown', handleEscKey)
})

onBeforeUnmount(() => {
  document.removeEventListener('keydown', handleEscKey)
  document.body.classList.remove('overflow-hidden')
})
</script>