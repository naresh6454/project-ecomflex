<template>
  <div class="min-h-screen bg-brand-light pt-16 pb-12">
    <Navbar />
    
    <main class="container mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div class="max-w-4xl mx-auto">
        <div class="bg-white rounded-xl shadow-lg p-6 mb-8 dark:bg-gray-800">
          <h1 class="text-2xl font-bold text-gray-900 mb-6 dark:text-white">
            Notification System Demo
          </h1>
          
          <div class="space-y-8">
            <!-- Toast Notifications Demo -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                Toast Notifications
              </h2>
              
              <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
                <button
                  @click="showToast('success', 'Success Message', 'Operation completed successfully!')"
                  class="px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-green-500 focus:ring-offset-2 transition-all duration-200"
                >
                  Success Toast
                </button>
                
                <button
                  @click="showToast('error', 'Error Message', 'Something went wrong. Please try again.')"
                  class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 transition-all duration-200"
                >
                  Error Toast
                </button>
                
                <button
                  @click="showToast('warning', 'Warning Message', 'Please review your information before proceeding.')"
                  class="px-4 py-2 bg-yellow-500 text-white rounded-lg hover:bg-yellow-600 focus:outline-none focus:ring-2 focus:ring-yellow-500 focus:ring-offset-2 transition-all duration-200"
                >
                  Warning Toast
                </button>
                
                <button
                  @click="showToast('info', 'Information', 'This is an informational message.')"
                  class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                >
                  Info Toast
                </button>
              </div>
              
              <div class="mt-4">
                <label class="flex items-center space-x-2">
                  <input
                    type="checkbox"
                    v-model="autoClose"
                    class="h-4 w-4 text-accent focus:ring-accent border-gray-300 rounded"
                  />
                  <span class="text-sm text-gray-700 dark:text-gray-300">Auto-close after</span>
                </label>
                
                <div v-if="autoClose" class="mt-2 flex items-center space-x-2">
                  <input
                    type="range"
                    v-model="duration"
                    min="1000"
                    max="10000"
                    step="1000"
                    class="w-48 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700"
                  />
                  <span class="text-sm text-gray-700 dark:text-gray-300">{{ duration / 1000 }}s</span>
                </div>
              </div>
            </section>
            
            <!-- Animation Loader Demo -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                Animated Loaders
              </h2>
              
              <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                <div class="flex flex-col items-center justify-center p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Spinner</h3>
                  <AnimatedLoader type="spinner" />
                </div>
                
                <div class="flex flex-col items-center justify-center p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Pulse</h3>
                  <AnimatedLoader type="pulse" />
                </div>
                
                <div class="flex flex-col items-center justify-center p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Progress ({{ loaderProgress }}%)</h3>
                  <AnimatedLoader
                    type="progress"
                    :progress="loaderProgress"
                    showPercentage
                    class="w-full"
                  />
                  <input
                    type="range"
                    v-model="loaderProgress"
                    min="0"
                    max="100"
                    class="w-full mt-3 h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700"
                  />
                </div>
                
                <div class="flex flex-col items-center justify-center p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Shimmer</h3>
                  <AnimatedLoader type="shimmer" size="lg" />
                </div>
              </div>
              
              <div class="mt-6">
                <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Loader Options</h3>
                
                <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1 dark:text-gray-300">Size</label>
                    <select
                      v-model="loaderSize"
                      class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md dark:bg-gray-800 dark:border-gray-600 dark:text-white"
                    >
                      <option value="xs">Extra Small</option>
                      <option value="sm">Small</option>
                      <option value="md">Medium</option>
                      <option value="lg">Large</option>
                      <option value="xl">Extra Large</option>
                    </select>
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1 dark:text-gray-300">Color</label>
                    <select
                      v-model="loaderColor"
                      class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md dark:bg-gray-800 dark:border-gray-600 dark:text-white"
                    >
                      <option value="accent">Accent (Gold)</option>
                      <option value="primary">Primary</option>
                      <option value="success">Success</option>
                      <option value="error">Error</option>
                      <option value="warning">Warning</option>
                      <option value="info">Info</option>
                    </select>
                  </div>
                  
                  <div>
                    <label class="block text-sm font-medium text-gray-700 mb-1 dark:text-gray-300">Thickness</label>
                    <select
                      v-model="loaderThickness"
                      class="block w-full pl-3 pr-10 py-2 text-base border-gray-300 focus:outline-none focus:ring-accent focus:border-accent sm:text-sm rounded-md dark:bg-gray-800 dark:border-gray-600 dark:text-white"
                    >
                      <option :value="2">Thin</option>
                      <option :value="4">Medium</option>
                      <option :value="8">Thick</option>
                    </select>
                  </div>
                </div>
                
                <div class="mt-4 p-6 bg-gray-100 rounded-lg flex items-center justify-center dark:bg-gray-800">
                  <AnimatedLoader
                    :type="loaderType"
                    :size="loaderSize"
                    :color="loaderColor"
                    :thickness="loaderThickness"
                    :progress="loaderProgress"
                    :indeterminate="loaderIndeterminate"
                    showPercentage
                    label="Custom Loader"
                  />
                </div>
                
                <div class="mt-4 grid grid-cols-2 sm:grid-cols-4 gap-4">
                  <button
                    v-for="type in loaderTypes"
                    :key="type"
                    @click="loaderType = type"
                    class="px-4 py-2 bg-white border border-gray-300 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200 dark:bg-gray-700 dark:border-gray-600 dark:text-white dark:hover:bg-gray-600"
                    :class="{ 'ring-2 ring-accent ring-offset-2': loaderType === type }"
                  >
                    {{ type.charAt(0).toUpperCase() + type.slice(1) }}
                  </button>
                </div>
                
                <div class="mt-4">
                  <label class="flex items-center space-x-2">
                    <input
                      type="checkbox"
                      v-model="loaderIndeterminate"
                      class="h-4 w-4 text-accent focus:ring-accent border-gray-300 rounded"
                    />
                    <span class="text-sm text-gray-700 dark:text-gray-300">Indeterminate progress</span>
                  </label>
                </div>
              </div>
            </section>
            
            <!-- Fullscreen loader demo -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                Fullscreen Loader
              </h2>
              
              <button
                @click="showFullscreenLoader"
                class="px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
              >
                Show Fullscreen Loader for 3 seconds
              </button>
            </section>
            
            <!-- Button ripple effect demo -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                Button Ripple Effect
              </h2>
              
              <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
                <button
                  @click="handleRippleClick"
                  class="px-6 py-3 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200 relative overflow-hidden"
                >
                  Click for Ripple Effect
                </button>
                
                <button
                  @click="handleRippleClick"
                  class="px-6 py-3 bg-white border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200 relative overflow-hidden dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:bg-gray-700"
                >
                  Light Button
                </button>
                
                <button
                  @click="handleRippleClick"
                  class="px-6 py-3 bg-gray-900 text-white rounded-lg hover:bg-gray-800 focus:outline-none focus:ring-2 focus:ring-gray-900 focus:ring-offset-2 transition-all duration-200 relative overflow-hidden dark:bg-gray-700"
                >
                  Dark Button
                </button>
              </div>
            </section>
            
            <!-- Animation composable demo -->
            <section>
              <h2 class="text-lg font-semibold text-gray-800 mb-4 dark:text-gray-200">
                Animation Composable
              </h2>
              
              <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
                <div class="p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Element Animations</h3>
                  
                  <div class="space-y-4">
                    <div ref="animationTarget" class="p-4 bg-gray-100 rounded-lg dark:bg-gray-800">
                      <p class="text-gray-700 dark:text-gray-300">This element will be animated</p>
                    </div>
                    
                    <div class="flex flex-wrap gap-2">
                      <button
                        @click="runAnimation('fadeIn')"
                        class="px-3 py-1 text-sm bg-accent text-white rounded-md hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      >
                        Fade In
                      </button>
                      
                      <button
                        @click="runAnimation('fadeOut')"
                        class="px-3 py-1 text-sm bg-accent text-white rounded-md hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      >
                        Fade Out
                      </button>
                      
                      <button
                        @click="runAnimation('slideIn', 'top')"
                        class="px-3 py-1 text-sm bg-accent text-white rounded-md hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      >
                        Slide In (Top)
                      </button>
                      
                      <button
                        @click="runAnimation('slideIn', 'right')"
                        class="px-3 py-1 text-sm bg-accent text-white rounded-md hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      >
                        Slide In (Right)
                      </button>
                      
                      <button
                        @click="runAnimation('scale')"
                        class="px-3 py-1 text-sm bg-accent text-white rounded-md hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      >
                        Scale
                      </button>
                      
                      <button
                        @click="runAnimation('pulseGlow')"
                        class="px-3 py-1 text-sm bg-accent text-white rounded-md hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                      >
                        Pulse Glow
                      </button>
                      
                      <button
                        @click="resetAnimation"
                        class="px-3 py-1 text-sm bg-gray-500 text-white rounded-md hover:bg-gray-600 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2 transition-all duration-200"
                      >
                        Reset
                      </button>
                    </div>
                  </div>
                </div>
                
                <div class="p-4 bg-white rounded-lg shadow dark:bg-gray-700">
                  <h3 class="text-sm font-medium text-gray-700 mb-3 dark:text-gray-300">Count Up Animation</h3>
                  
                  <div class="text-center p-6 bg-gray-100 rounded-lg dark:bg-gray-800">
                    <span ref="counterElement" class="text-4xl font-bold text-accent">0</span>
                  </div>
                  
                  <div class="mt-4 space-y-2">
                    <div class="flex justify-between">
                      <label class="text-sm text-gray-700 dark:text-gray-300">Target Value:</label>
                      <span class="text-sm font-medium text-gray-900 dark:text-white">{{ counterValue }}</span>
                    </div>
                    
                    <input
                      type="range"
                      v-model="counterValue"
                      min="0"
                      max="10000"
                      step="100"
                      class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700"
                    />
                    
                    <div class="flex justify-between">
                      <label class="text-sm text-gray-700 dark:text-gray-300">Duration (ms):</label>
                      <span class="text-sm font-medium text-gray-900 dark:text-white">{{ counterDuration }}</span>
                    </div>
                    
                    <input
                      type="range"
                      v-model="counterDuration"
                      min="500"
                      max="5000"
                      step="100"
                      class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700"
                    />
                    
                    <button
                      @click="startCountUp"
                      class="w-full px-4 py-2 bg-accent text-white rounded-lg hover:bg-accent-dark focus:outline-none focus:ring-2 focus:ring-accent focus:ring-offset-2 transition-all duration-200"
                    >
                      Start Count Up
                    </button>
                  </div>
                </div>
              </div>
            </section>
          </div>
        </div>
      </div>
    </main>
    
    <!-- Fullscreen Loader Modal -->
    <Teleport to="body">
      <transition
        enter-active-class="transition ease-out duration-300"
        enter-from-class="opacity-0"
        enter-to-class="opacity-100"
        leave-active-class="transition ease-in duration-200"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
      >
        <div v-if="showingFullscreenLoader" class="fixed inset-0 z-50 flex items-center justify-center">
          <AnimatedLoader
            :type="loaderType"
            size="xl"
            color="accent"
            :thickness="loaderThickness"
            fullscreen
            overlay
            label="Loading..."
          />
        </div>
      </transition>
    </Teleport>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Navbar from '@/components/layouts/Navbar.vue'
import AnimatedLoader from '@/components/ui/AnimatedLoader.vue'
import { useNotificationStore } from '@/stores/notification'
import { useAnimation } from '@/composables/useAnimation'

// Store and composables
const notificationStore = useNotificationStore()
const { fadeIn, fadeOut, slideIn, scale, pulseGlow, countUp, buttonRipple } = useAnimation()

// Toast settings
const autoClose = ref(true)
const duration = ref(5000)

// Loader types
type LoaderType = 'spinner' | 'pulse' | 'progress' | 'shimmer'
type LoaderSize = 'xs' | 'sm' | 'md' | 'lg' | 'xl'
type LoaderColor = 'accent' | 'primary' | 'success' | 'error' | 'warning' | 'info'
type LoaderThickness = 2 | 4 | 8

// Loader types array for iteration
const loaderTypes: LoaderType[] = ['spinner', 'pulse', 'progress', 'shimmer']

// Loader settings with proper type annotations
const loaderType = ref<LoaderType>('spinner')
const loaderSize = ref<LoaderSize>('md')
const loaderColor = ref<LoaderColor>('accent')
const loaderThickness = ref<LoaderThickness>(4)
const loaderProgress = ref(60)
const loaderIndeterminate = ref(false)
const showingFullscreenLoader = ref(false)

// Animation demo refs
const animationTarget = ref<HTMLElement | null>(null)
const currentAnimation = ref<any>(null)

// Counter animation
const counterElement = ref<HTMLElement | null>(null)
const counterValue = ref(1000)
const counterDuration = ref(2000)
let counterInstance: any = null

// Show toast notification
const showToast = (type: 'success' | 'error' | 'warning' | 'info', title: string, message: string) => {
  notificationStore.showToast({
    type,
    title,
    message,
    autoClose: autoClose.value,
    duration: autoClose.value ? duration.value : undefined
  })
}

// Show fullscreen loader
const showFullscreenLoader = () => {
  showingFullscreenLoader.value = true
  
  setTimeout(() => {
    showingFullscreenLoader.value = false
    
    notificationStore.showToast({
      type: 'success',
      message: 'Operation completed successfully!'
    })
  }, 3000)
}

// Handle ripple effect click
const handleRippleClick = (event: MouseEvent) => {
  buttonRipple(event, {
    color: 'rgba(255, 215, 0, 0.3)', // Golden ripple
    duration: 600
  })
}

// Run animation on target element
const runAnimation = (animationType: string, direction?: string) => {
  if (!animationTarget.value) return
  
  // Reset previous animation if any
  resetAnimation()
  
  // Run the requested animation
  switch (animationType) {
    case 'fadeIn':
      currentAnimation.value = fadeIn(animationTarget.value, { duration: 800 })
      break
    case 'fadeOut':
      currentAnimation.value = fadeOut(animationTarget.value, { duration: 800 })
      break
    case 'slideIn':
      currentAnimation.value = slideIn(
        animationTarget.value,
        direction as any || 'bottom',
        '30px',
        { duration: 800 }
      )
      break
    case 'scale':
      currentAnimation.value = scale(animationTarget.value, 0.5, 1, { duration: 800 })
      break
    case 'pulseGlow':
      currentAnimation.value = pulseGlow(
        animationTarget.value,
        'rgba(255, 215, 0, 0.3)',
        { duration: 1500, iterations: 3 }
      )
      break
  }
}

// Reset animation
const resetAnimation = () => {
  if (currentAnimation.value) {
    currentAnimation.value.cancel()
    currentAnimation.value = null
  }
}

// Start count up animation
const startCountUp = () => {
  if (!counterElement.value) return
  
  // Stop previous animation if running
  if (counterInstance) {
    counterInstance.stop()
  }
  
  // Start new animation
  counterInstance = countUp(counterElement.value, counterValue.value, {
    duration: counterDuration.value,
    decimals: 0,
    startValue: 0,
    easing: 'ease-out'
  })
}

// Initialize demo
onMounted(() => {
  // Initialize with fade in animation
  if (animationTarget.value) {
    animationTarget.value.style.opacity = '0'
    setTimeout(() => {
      runAnimation('fadeIn')
    }, 500)
  }
  
  // Demo notification
  setTimeout(() => {
    notificationStore.showToast({
      type: 'info',
      title: 'Welcome to the Animation Demo',
      message: 'Try out different animations and notification styles'
    })
  }, 1000)
})
</script>